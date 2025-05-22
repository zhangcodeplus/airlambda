package models

import (
	"database/sql"
	"time"
)

type Article struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Status    string    `json:"status"` // active, expired
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ArticleModel struct {
	DB *sql.DB
}

func NewArticleModel(db *sql.DB) *ArticleModel {
	return &ArticleModel{DB: db}
}

func (m *ArticleModel) GetAll() ([]Article, error) {
	rows, err := m.DB.Query("SELECT id, title, content, status, created_at, updated_at FROM articles")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []Article
	for rows.Next() {
		var a Article
		err := rows.Scan(&a.ID, &a.Title, &a.Content, &a.Status, &a.CreatedAt, &a.UpdatedAt)
		if err != nil {
			return nil, err
		}
		articles = append(articles, a)
	}
	return articles, nil
}

func (m *ArticleModel) Create(title, content string) error {
	_, err := m.DB.Exec(
		"INSERT INTO articles (title, content, status, created_at, updated_at) VALUES (?, ?, 'active', NOW(), NOW())",
		title, content,
	)
	return err
}

func (m *ArticleModel) UpdateStatus(id int64, status string) error {
	_, err := m.DB.Exec(
		"UPDATE articles SET status = ?, updated_at = NOW() WHERE id = ?",
		status, id,
	)
	return err
}

func (m *ArticleModel) Delete(id int64) error {
	_, err := m.DB.Exec("DELETE FROM articles WHERE id = ?", id)
	return err
}
