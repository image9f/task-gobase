package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Student struct {
	ID    int
	Name  string
	Age   int
	Grade string
}

func main() {

	db, err := gorm.Open(sqlite.Open("students.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Student{})

	fmt.Println("students db init success")

	//s := []Student{
	//	{1, "aaa", 10, "一年级"},
	//	{2, "bbb", 15, "五年级"},
	//	{3, "ccc", 18, "六年级"},
	//}
	//result := db.Create(&s)
	//if result.Error != nil {
	//	panic(result.Error.Error())
	//}

	//插入张三学生信息
	//stu1 := Student{4, "张三", 20, "三年级"}
	//result := db.Create(&stu1)
	//if result.Error != nil {
	//	fmt.Println("create student fail")
	//}
	//fmt.Println("students db insert 张三 success")

	//查找年龄>18的学生
	//var students []Student
	//db.Debug().Where("age > ?", 18).Find(&students)
	//for _, student := range students {
	//	fmt.Println(student)
	//}

	//更新张三为四年级
	//result := db.Model(&Student{}).Where("name = ?", "张三").Update("grade", "四年级")
	//if result.Error != nil {
	//	fmt.Println("update error", result.Error)
	//}
	//var test Student
	//db.Where("name=?", "张三").Find(&test)
	//fmt.Println(test)

	//删除小于15岁的学生
	result := db.Where("Age < ?", 15).Delete(&Student{})
	if result.Error != nil {
		fmt.Println("delete students failed")
	}

}
