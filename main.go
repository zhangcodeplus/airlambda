package main

import (
	"database/sql"
	"log"
	"net/http"

	"airlambda/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db *sql.DB
var cfg *config.Config

func initDB() {
	var err error
	// 加载配置文件
	cfg, err = config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	// 连接数据库
	db, err = sql.Open("postgres", cfg.GetDSN())
	if err != nil {
		log.Fatal(err)
	}

	// 测试数据库连接
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// 创建文章表
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS articles (
			id SERIAL PRIMARY KEY,
			title TEXT NOT NULL,
			content TEXT NOT NULL,
			status TEXT NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// 初始化数据库
	initDB()

	r := gin.Default()

	// 配置 CORS
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = cfg.Server.CORS.AllowedOrigins
	corsConfig.AllowMethods = cfg.Server.CORS.AllowedMethods
	corsConfig.AllowHeaders = cfg.Server.CORS.AllowedHeaders
	r.Use(cors.New(corsConfig))

	// API 路由
	api := r.Group("/api")
	{
		// 文章相关路由
		api.GET("/articles", getArticles)
		api.POST("/articles", createArticle)
		api.PUT("/articles/:id/status", updateArticleStatus)
		api.DELETE("/articles/:id", deleteArticle)
	}

	// 启动服务器
	serverAddr := ":" + cfg.Server.Port
	if err := r.Run(serverAddr); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

// 获取所有文章
func getArticles(c *gin.Context) {
	rows, err := db.Query("SELECT id, title, content, status, created_at, updated_at FROM articles ORDER BY created_at DESC")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var articles []map[string]interface{}
	for rows.Next() {
		var article struct {
			ID        int64
			Title     string
			Content   string
			Status    string
			CreatedAt string
			UpdatedAt string
		}
		err := rows.Scan(&article.ID, &article.Title, &article.Content, &article.Status, &article.CreatedAt, &article.UpdatedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		articles = append(articles, map[string]interface{}{
			"id":         article.ID,
			"title":      article.Title,
			"content":    article.Content,
			"status":     article.Status,
			"created_at": article.CreatedAt,
			"updated_at": article.UpdatedAt,
		})
	}
	c.JSON(http.StatusOK, articles)
}

// 创建新文章
func createArticle(c *gin.Context) {
	var article struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Exec(
		"INSERT INTO articles (title, content, status) VALUES ($1, $2, 'active')",
		article.Title, article.Content,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "文章创建成功"})
}

// 更新文章状态
func updateArticleStatus(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		Status string `json:"status"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Exec(
		"UPDATE articles SET status = $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2",
		req.Status, id,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "状态更新成功"})
}

// 删除文章
func deleteArticle(c *gin.Context) {
	id := c.Param("id")
	_, err := db.Exec("DELETE FROM articles WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "文章删除成功"})
}
