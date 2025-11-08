# Go Gin Blog API

这是一个使用 Go 语言和 Gin 框架构建的简单博客 API。它提供了用户注册、登录、文章管理和评论功能。

## 功能

- 用户认证 (注册、登录、JWT)
- 文章管理 (创建、读取、更新、删除)
- 评论管理 (创建、读取)
- 统一错误处理

## 技术栈

- **后端**: Go (Gin 框架)
- **数据库**: SQLite (通过 GORM)
- **认证**: JWT (JSON Web Tokens)
- **密码加密**: bcrypt

## 运行环境

- Go 1.16+ (推荐最新稳定版本)

## 依赖安装

1.  **克隆仓库**:

    ```bash
    git clone <仓库地址>
    cd blog
    ```

2.  **安装 Go 模块依赖**:

    ```bash
    go mod tidy
    ```

    这将下载所有必要的 Go 模块，例如 Gin、GORM、JWT 和 bcrypt。

## 启动方式

1.  **运行应用程序**:

    ```bash
    go run main.go
    ```

    应用程序将在 `http://localhost:8080` 端口上启动。

## API 接口文档

以下是可用的 API 接口及其功能：

### 用户认证

-   **注册用户**
    -   **URL**: `/register`
    -   **方法**: `POST`
    -   **请求体**: `application/json`
        ```json
        {
            "username": "testuser",
            "password": "password123",
            "email": "test@example.com"
        }
        ```
    -   **响应**: `application/json`
        ```json
        {
            "message": "User registered successfully"
        }
        ```

-   **登录用户**
    -   **URL**: `/login`
    -   **方法**: `POST`
    -   **请求体**: `application/json`
        ```json
        {
            "username": "testuser",
            "password": "password123"
        }
        ```
    -   **响应**: `application/json`
        ```json
        {
            "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
        }
        ```

### 文章管理

-   **创建文章** (需要认证)
    -   **URL**: `/auth/posts`
    -   **方法**: `POST`
    -   **请求头**: `Authorization: Bearer <JWT_TOKEN>`
    -   **请求体**: `application/json`
        ```json
        {
            "title": "我的第一篇文章",
            "content": "这是我的第一篇文章内容。"
        }
        ```
    -   **响应**: `application/json`
        ```json
        {
            "message": "Post created successfully",
            "post": {
                "ID": 1,
                "CreatedAt": "2023-10-27T10:00:00Z",
                "UpdatedAt": "2023-10-27T10:00:00Z",
                "DeletedAt": null,
                "title": "我的第一篇文章",
                "content": "这是我的第一篇文章内容。",
                "user_id": 1
            }
        }
        ```

-   **获取所有文章**
    -   **URL**: `/posts`
    -   **方法**: `GET`
    -   **响应**: `application/json`
        ```json
        {
            "posts": [
                {
                    "ID": 1,
                    "CreatedAt": "2023-10-27T10:00:00Z",
                    "UpdatedAt": "2023-10-27T10:00:00Z",
                    "DeletedAt": null,
                    "title": "我的第一篇文章",
                    "content": "这是我的第一篇文章内容。",
                    "user_id": 1,
                    "user": {
                        "ID": 1,
                        "CreatedAt": "2023-10-27T09:50:00Z",
                        "UpdatedAt": "2023-10-27T09:50:00Z",
                        "DeletedAt": null,
                        "username": "testuser",
                        "email": "test@example.com"
                    }
                }
            ]
        }
        ```

-   **获取单篇文章**
    -   **URL**: `/posts/:id` (例如: `/posts/1`)
    -   **方法**: `GET`
    -   **响应**: `application/json`
        ```json
        {
            "post": {
                "ID": 1,
                "CreatedAt": "2023-10-27T10:00:00Z",
                "UpdatedAt": "2023-10-27T10:00:00Z",
                "DeletedAt": null,
                "title": "我的第一篇文章",
                "content": "这是我的第一篇文章内容。",
                "user_id": 1,
                "user": {
                        "ID": 1,
                        "CreatedAt": "2023-10-27T09:50:00Z",
                        "UpdatedAt": "2023-10-27T09:50:00Z",
                        "DeletedAt": null,
                        "username": "testuser",
                        "email": "test@example.com"
                    }
            }
        }
        ```

-   **更新文章** (需要认证，且只能更新自己的文章)
    -   **URL**: `/auth/posts/:id` (例如: `/auth/posts/1`)
    -   **方法**: `PUT`
    -   **请求头**: `Authorization: Bearer <JWT_TOKEN>`
    -   **请求体**: `application/json`
        ```json
        {
            "title": "更新后的文章标题",
            "content": "这是更新后的文章内容。"
        }
        ```
    -   **响应**: `application/json`
        ```json
        {
            "message": "Post updated successfully",
            "post": {
                "ID": 1,
                "CreatedAt": "2023-10-27T10:00:00Z",
                "UpdatedAt": "2023-10-27T10:30:00Z",
                "DeletedAt": null,
                "title": "更新后的文章标题",
                "content": "这是更新后的文章内容。",
                "user_id": 1
            }
        }
        ```

-   **删除文章** (需要认证，且只能删除自己的文章)
    -   **URL**: `/auth/posts/:id` (例如: `/auth/posts/1`)
    -   **方法**: `DELETE`
    -   **请求头**: `Authorization: Bearer <JWT_TOKEN>`
    -   **响应**: `application/json`
        ```json
        {
            "message": "Post deleted successfully"
        }
        ```

### 评论管理

-   **创建评论** (需要认证)
    -   **URL**: `/auth/posts/:post_id/comments` (例如: `/auth/posts/1/comments`)
    -   **方法**: `POST`
    -   **请求头**: `Authorization: Bearer <JWT_TOKEN>`
    -   **请求体**: `application/json`
        ```json
        {
            "content": "这是一条评论。"
        }
        ```
    -   **响应**: `application/json`
        ```json
        {
            "message": "Comment created successfully",
            "comment": {
                "ID": 1,
                "CreatedAt": "2023-10-27T11:00:00Z",
                "UpdatedAt": "2023-10-27T11:00:00Z",
                "DeletedAt": null,
                "content": "这是一条评论。",
                "user_id": 1,
                "post_id": 1
            }
        }
        ```

-   **获取文章评论**
    -   **URL**: `/posts/:id/comments` (例如: `/posts/1/comments`)
    -   **方法**: `GET`
    -   **响应**: `application/json`
        ```json
        {
            "comments": [
                {
                    "ID": 1,
                    "CreatedAt": "2023-10-27T11:00:00Z",
                    "UpdatedAt": "2023-10-27T11:00:00Z",
                    "DeletedAt": null,
                    "content": "这是一条评论。",
                    "user_id": 1,
                    "post_id": 1,
                    "user": {
                        "ID": 1,
                        "CreatedAt": "2023-10-27T09:50:00Z",
                        "UpdatedAt": "2023-10-27T09:50:00Z",
                        "DeletedAt": null,
                        "username": "testuser",
                        "email": "test@example.com"
                    }
                }
            ]
        }
        ```