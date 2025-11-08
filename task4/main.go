package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

var (
	db     *gorm.DB
	jwtKey = []byte("your_secret_key") // JWT 密钥
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username" binding:"required"`
	Password string `gorm:"not null" json:"password" binding:"required"`
	Email    string `gorm:"unique;not null" json:"email" binding:"required,email"`
}

type Post struct {
	gorm.Model
	Title   string `gorm:"not null" json:"title" binding:"required"`
	Content string `gorm:"not null" json:"content" binding:"required"`
	UserID  uint   `json:"user_id"`
	User    User   `json:"user"`
}

type Comment struct {
	gorm.Model
	Content string `gorm:"not null" json:"content" binding:"required"`
	UserID  uint   `json:"user_id"`
	User    User   `json:"user"`
	PostID  uint   `json:"post_id"`
	Post    Post   `json:"post"`
}

func init() {
	db, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(User{}, Post{}, Comment{})

}

func Register(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)

	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}

	var storedUser User
	if err := db.Where("username = ?", user.Username).First(&storedUser).Error; err != nil {
		c.Error(err)
		return
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		c.Error(err)
		return
	}

	// 生成 JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       storedUser.ID,
		"username": storedUser.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.Error(err)
		return
	}

	// 记录登录成功日志
	log.Printf("[INFO] 用户登录成功: %s", storedUser.Username)
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func CreatePost(c *gin.Context) {
	userID, _ := c.Get("userID")
	var post Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}

	post.UserID = userID.(uint)
	if err := db.Create(&post).Error; err != nil {
		c.Error(err)
		return
	}

	log.Printf("[INFO] 用户 %d 创建了文章: %s", userID, post.Title)
	c.JSON(http.StatusCreated, gin.H{"message": "Post created successfully", "post": post})
}

func CreateComment(c *gin.Context) {
	userID, _ := c.Get("userID")
	postID := c.Param("post_id")

	var post Post
	if err := db.First(&post, postID).Error; err != nil {
		c.Error(err)
		return
	}
	var comment Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}
	comment.UserID = userID.(uint)
	comment.PostID = post.ID
	if err := db.Create(&comment).Error; err != nil {
		c.Error(err)
		return
	}

	log.Printf("用户%d对文章:%s发表了评论:%s", userID, post.ID, comment.Content)
	c.JSON(http.StatusCreated, gin.H{"message": "Comment created successfully", "comment": comment})
}

func GetPost(c *gin.Context) {
	id := c.Param("id")

	var post Post
	if err := db.Preload("User").First(&post, id).Error; err != nil {
		c.Error(err)
		return
	}

	log.Printf("获取文章内容: %s", post.Title)
	c.JSON(http.StatusOK, gin.H{"post": post})
}

func GetPosts(c *gin.Context) {
	var posts []Post
	if err := db.Preload("User").Find(&posts).Error; err != nil {
		c.Error(err)
		return
	}
	log.Printf("获取所有文章列表")
	c.JSON(http.StatusOK, gin.H{"posts": posts})
}

func GetComments(c *gin.Context) {
	postID := c.Param("post_id")

	var post Post
	if err := db.First(&post, postID).Error; err != nil {
		c.Error(err)
		return
	}
	var comments []Comment
	if err := db.Where("post_id=?", postID).Preload("User").Find(&comments).Error; err != nil {
		c.Error(err)
		return
	}

	log.Printf("获取文章 %s 的评论列表", postID)
	c.JSON(http.StatusOK, gin.H{"comments": comments})
}

func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	userID, _ := c.Get("userID")

	var post Post
	if err := db.First(&post, id).Error; err != nil {
		c.Error(err)
		return
	}

	if post.UserID != userID.(uint) {
		c.Error(fmt.Errorf("无权限修改文章")).SetType(gin.ErrorTypePrivate)
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "无权限修改文章"})
		return
	}
	if err := c.ShouldBindJSON(&post); err != nil {
		c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}
	if err := db.Save(&post).Error; err != nil {
		c.Error(err)
		return
	}

	log.Printf("用户%d更新了文章%s", userID, post.Title)
	c.JSON(http.StatusOK, gin.H{"message": "Post updated successfully", "post": post})
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")
	userID, _ := c.Get("userID")

	var post Post
	if err := db.First(&post, id).Error; err != nil {
		c.Error(err)
		return
	}

	if post.UserID != userID.(uint) {
		c.Error(fmt.Errorf("无权删除此文章")).SetType(gin.ErrorTypePrivate)
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "无权删除此文章"})
		return
	}

	if err := db.Delete(&post).Error; err != nil {
		c.Error(err)
		return
	}

	log.Printf("用户 %d 删除了文章: %s (ID: %s)", userID, post.Title, id)
	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.Error(fmt.Errorf("Authorization header required"))
			c.Abort()
			return
		}

		tokenString = tokenString[len("Bearer "):] // 移除 "Bearer " 前缀

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtKey, nil
		})

		if err != nil {
			c.Error(err)
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("userID", claims["id"])
			c.Next()
		} else {
			c.Error(fmt.Errorf("invalid token"))
			c.Abort()
		}
	}
}

func main() {

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Blog System",
		})
	})
	r.POST("/register", Register)
	r.POST("/login", Login)

	r.GET("/posts", GetPosts)
	r.GET("/posts/:id/coments", GetComments)
	r.GET("/posts/:id", GetPost)

	auth := r.Group("/auth")
	auth.Use(AuthMiddleware())
	{
		auth.GET("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Authenticated!"})
		})
		auth.POST("/posts", CreatePost)
		auth.PUT("/posts/:id", UpdatePost)
		auth.DELETE("/posts/:id", DeletePost)
		auth.POST("/posts/:post_id/comments", CreateComment)
	}

	r.Run(":8080")

}
