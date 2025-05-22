---
title: "如何使用 Typora 编辑文章"
date: 2024-02-21T10:00:00+08:00
categories: ["挖矿"]
subcategories: ["手机挖矿"]
tags: ["教程", "Typora", "Markdown"]
thumbnail: "/images/uploads/typora.png"
---

# 使用 Typora 编辑文章

Typora 是一个所见即所得的 Markdown 编辑器，让您可以专注于写作。

## 基本功能

### 1. 标题
使用 `#` 符号创建标题，例如：
# 一级标题
## 二级标题
### 三级标题

### 2. 文本格式化
- **粗体**：使用 `**文本**`
- *斜体*：使用 `*文本*`
- ~~删除线~~：使用 `~~文本~~`
- `代码`：使用 `` `代码` ``

### 3. 列表
无序列表：
- 项目 1
- 项目 2
- 项目 3

有序列表：
1. 第一步
2. 第二步
3. 第三步

### 4. 链接和图片
[链接文本](https://example.com)
![图片描述](/images/uploads/example.jpg)

### 5. 代码块
```python
def hello_world():
    print("Hello, World!")
```

### 6. 表格
| 标题 1 | 标题 2 | 标题 3 |
|--------|--------|--------|
| 内容 1 | 内容 2 | 内容 3 |
| 内容 4 | 内容 5 | 内容 6 |

### 7. 引用
> 这是一段引用文本
> 可以包含多行

## 使用技巧

1. 使用 `Ctrl + B` 快速加粗
2. 使用 `Ctrl + I` 快速斜体
3. 使用 `Ctrl + K` 快速插入链接
4. 使用 `Ctrl + Shift + I` 快速插入图片
5. 使用 `Ctrl + Shift + Q` 快速插入引用

## 图片管理

1. 将图片拖拽到编辑器中
2. 图片会自动复制到 `static/images/uploads` 目录
3. 图片路径会自动更新

## 文章元数据

在文章开头使用 `---` 包围的 YAML 格式元数据：

```yaml
---
title: "文章标题"
date: 2024-02-21T10:00:00+08:00
categories: ["分类"]
subcategories: ["子分类"]
tags: ["标签1", "标签2"]
thumbnail: "/images/uploads/图片.jpg"
---
```

## 导出选项

Typora 支持导出为：
- PDF
- HTML
- Word
- LaTeX
- 等多种格式

## 主题设置

1. 打开 Typora 设置
2. 选择"外观"
3. 选择喜欢的主题
4. 可以自定义 CSS

## 快捷键

- `Ctrl + N`：新建文件
- `Ctrl + S`：保存文件
- `Ctrl + P`：打印/导出
- `Ctrl + /`：切换源代码模式
- `Ctrl + Shift + L`：显示/隐藏侧边栏 