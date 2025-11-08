package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Comment struct {
	ID      uint `gorm:"primary_key;auto_increment"`
	Content string
	PostID  uint
	UserID  uint
	User    User `gorm:"foreignkey:UserID"`
}
type Post struct {
	ID       uint `gorm:"primary_key;auto_increment"`
	Title    string
	Content  string
	UserID   uint
	Comments []Comment `gorm:"foreignkey:PostID"`
	States   string    `gorm:"default: 有评论"`
}
type User struct {
	ID        uint `gorm:"primary_key;auto_increment"`
	Name      string
	Account   string
	Posts     []Post
	PostCount uint
}

type MaxComment struct {
	Post
	Count uint
}

// 在创建POST后更新user的postsCount
func (p *Post) AfterCreate(tx *gorm.DB) error {
	var user User
	if err := tx.First(&user, "id=?", p.UserID).Error; err != nil {
		return err
	}
	user.PostCount++
	if err := tx.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

// 删除评论后更新状态
func (c *Comment) AfterDelete(tx *gorm.DB) error {
	var count int64
	tx.Model(&c).Where("post_id=?", c.PostID).Count(&count)
	if count == 0 {
		tx.Model(&Post{}).Where("id=?", c.PostID).Update("states", "无评论")
	}
	return nil
}

func main() {

	/*
		模型定义
	*/

	db, err := gorm.Open(sqlite.Open("blogs.db"), &gorm.Config{})
	if err != nil {
		fmt.Errorf("failed to connect database")
	}

	db.AutoMigrate(&User{}, &Post{}, &Comment{})

	//users := []User{
	//	{Name: "John", Account: "111111"},
	//	{Name: "Jane", Account: "222222"},
	//	{Name: "Doe", Account: "333333"},
	//}
	//db.Create(&users)
	//
	//posts := []Post{
	//	{Title: "bookA", Content: "xxxxxxxxx", UserID: users[0].ID},
	//	{Title: "bookB", Content: "yyyyyyyyy", UserID: users[0].ID},
	//	{Title: "bookC", Content: "ZZZZZZZZZ", UserID: users[1].ID},
	//	{Title: "bookD", Content: "AAAAAAAAA", UserID: users[2].ID},
	//}
	//db.Create(&posts)
	//
	//comments := []Comment{
	//	{Content: "good", PostID: posts[0].ID, UserID: users[0].ID},
	//	{Content: "good good", PostID: posts[0].ID, UserID: users[0].ID},
	//	{Content: "general", PostID: posts[2].ID, UserID: users[1].ID},
	//	{Content: "general general", PostID: posts[3].ID, UserID: users[2].ID},
	//}
	//db.Create(&comments)
	//
	///*
	//	关联查询
	//*/
	//
	////查询某个用户发布的所有文章及其对应的评论信息
	//fmt.Println("查询某个用户发布的所有文章及其对应的评论信息")
	//var user User
	//db.Preload("Posts.Comments").First(&user, "name = ?", "John")
	//if user.Name != "John" {
	//	fmt.Println("user's name is not John")
	//} else {
	//	fmt.Println("user's name is John")
	//	if len(user.Posts) == 0 {
	//		fmt.Println("John's posts is empty")
	//	} else {
	//		fmt.Println("John's posts length is ", len(user.Posts))
	//		for _, post := range user.Posts {
	//			if len(post.Comments) == 0 {
	//				fmt.Println("  post's post is ", post)
	//				fmt.Println("    post's comments is empty")
	//			} else {
	//				fmt.Println("  post's post is ", post)
	//				for _, comment := range post.Comments {
	//					fmt.Println("    post's comments is ", comment)
	//				}
	//			}
	//		}
	//	}
	//}
	//
	////查询评论数量最多的文章信息
	//fmt.Println("查询评论数量最多的文章信息")
	//var maxComment MaxComment
	//db.Table("posts").
	//	Select("posts.*,count(comments.id) as count").
	//	Joins("left join comments on comments.post_id = posts.id").
	//	Group("posts.id").
	//	Order("count desc").
	//	Limit(1).
	//	Scan(&maxComment)
	//
	//if maxComment.Count == 0 {
	//	fmt.Println("the max comment count is 0")
	//} else {
	//	fmt.Printf("the max comment count (post: %s, userid: %d, count: %d \n ",
	//		maxComment.Title, maxComment.UserID, maxComment.Count)
	//}

	/*
		钩子函数
	*/
	fmt.Println("钩子函数测试")
	var states Post
	db.First(&states, 3) //posts[0].ID
	fmt.Printf("before, title: %s ,id:%d, state:%s \n ", states.Title, states.ID, states.States)

	//删除评论
	var del []Comment
	db.Where("post_id = ?", 3).Find(&del)
	fmt.Println("del:", del)
	for _, comment := range del {
		db.Delete(&comment)
	}

	db.First(&states, 3)
	fmt.Printf("after, title: %s ,id:%d, state:%s \n ", states.Title, states.ID, states.States)
}
