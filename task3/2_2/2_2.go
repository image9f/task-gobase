package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Book struct {
	ID     uint `gorm:"primaryKey;autoIncrement"`
	Title  string
	Author string
	Price  float64
}

func main() {
	db, err := gorm.Open(sqlite.Open("book.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	db.AutoMigrate(&Book{})

	//初始化book
	books := []Book{
		{Title: "Go编程", Author: "张三", Price: 45.50},
		{Title: "GORM实战", Author: "李四", Price: 65.00},
		{Title: "Web开发", Author: "王五", Price: 72.80},
		{Title: "数据结构", Author: "赵六", Price: 39.90},
		{Title: "算法导论", Author: "钱七", Price: 99.99},
	}
	for _, book := range books {
		db.FirstOrCreate(&book, Book{Title: book.Title, Author: book.Author, Price: book.Price})
	}

	//查询价格大于50的书
	var bookfind []Book
	err = db.Where("price > ?", 50).Find(&bookfind).Error
	if err != nil {
		fmt.Println("查询价格>50的书失败", err)
	}

	for _, book := range bookfind {
		fmt.Println(book)
	}

}
