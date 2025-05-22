<template>
  <div class="app">
    <el-container>
      <el-header>
        <h1>文章管理系统</h1>
      </el-header>
      
      <el-main>
        <el-button type="primary" @click="showAddDialog">添加文章</el-button>
        
        <el-table :data="articles" style="width: 100%; margin-top: 20px">
          <el-table-column prop="id" label="ID" width="80" />
          <el-table-column prop="title" label="标题" />
          <el-table-column prop="status" label="状态" width="120">
            <template #default="scope">
              <el-tag :type="scope.row.status === 'active' ? 'success' : 'warning'">
                {{ scope.row.status === 'active' ? '活跃' : '已过期' }}
              </el-tag>
            </template>
          </el-table-column>
          <el-table-column prop="created_at" label="创建时间" width="180" />
          <el-table-column label="操作" width="200">
            <template #default="scope">
              <el-button
                size="small"
                :type="scope.row.status === 'active' ? 'warning' : 'success'"
                @click="updateStatus(scope.row)"
              >
                {{ scope.row.status === 'active' ? '标记过期' : '标记活跃' }}
              </el-button>
              <el-button
                size="small"
                type="danger"
                @click="deleteArticle(scope.row)"
              >
                删除
              </el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-main>
    </el-container>

    <!-- 添加文章对话框 -->
    <el-dialog v-model="dialogVisible" title="添加文章">
      <el-form :model="newArticle" label-width="80px">
        <el-form-item label="标题">
          <el-input v-model="newArticle.title" />
        </el-form-item>
        <el-form-item label="内容">
          <el-input
            v-model="newArticle.content"
            type="textarea"
            :rows="4"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" @click="addArticle">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import axios from 'axios'
import { ElMessage, ElMessageBox } from 'element-plus'

export default {
  name: 'App',
  setup() {
    const articles = ref([])
    const dialogVisible = ref(false)
    const newArticle = ref({
      title: '',
      content: ''
    })

    const fetchArticles = async () => {
      try {
        const response = await axios.get('http://localhost:3000/api/articles')
        articles.value = response.data
      } catch (error) {
        ElMessage.error('获取文章列表失败')
      }
    }

    const showAddDialog = () => {
      newArticle.value = { title: '', content: '' }
      dialogVisible.value = true
    }

    const addArticle = async () => {
      try {
        await axios.post('http://localhost:3000/api/articles', newArticle.value)
        ElMessage.success('添加文章成功')
        dialogVisible.value = false
        fetchArticles()
      } catch (error) {
        ElMessage.error('添加文章失败')
      }
    }

    const updateStatus = async (article) => {
      try {
        const newStatus = article.status === 'active' ? 'expired' : 'active'
        await axios.put(`http://localhost:3000/api/articles/${article.id}/status`, {
          status: newStatus
        })
        ElMessage.success('更新状态成功')
        fetchArticles()
      } catch (error) {
        ElMessage.error('更新状态失败')
      }
    }

    const deleteArticle = async (article) => {
      try {
        await ElMessageBox.confirm('确定要删除这篇文章吗？', '警告', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        })
        
        await axios.delete(`http://localhost:3000/api/articles/${article.id}`)
        ElMessage.success('删除文章成功')
        fetchArticles()
      } catch (error) {
        if (error !== 'cancel') {
          ElMessage.error('删除文章失败')
        }
      }
    }

    onMounted(() => {
      fetchArticles()
    })

    return {
      articles,
      dialogVisible,
      newArticle,
      showAddDialog,
      addArticle,
      updateStatus,
      deleteArticle
    }
  }
}
</script>

<style>
.app {
  padding: 20px;
}

.el-header {
  background-color: #409EFF;
  color: white;
  line-height: 60px;
  text-align: center;
}

.el-main {
  padding: 20px;
}

.dialog-footer {
  display: flex;
  justify-content: flex-end;
  gap: 10px;
}
</style> 