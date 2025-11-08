package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// Employee 员工结构体
type Employee struct {
	ID         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

func main() {
	// 连接到 SQLite 数据库
	db, err := sqlx.Connect("sqlite3", "employees.db")
	if err != nil {
		log.Fatalf("无法连接到数据库: %v", err)
	}
	defer db.Close()

	// 创建 employees 表
	schema := `
	CREATE TABLE IF NOT EXISTS employees (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		department TEXT,
		salary REAL
	);
	`
	_, err = db.Exec(schema)
	if err != nil {
		log.Fatalf("无法创建表: %v", err)
	}

	// 插入示例数据
	fmt.Println("--- 插入示例数据 ---")
	insertStmt := `INSERT INTO employees (name, department, salary) VALUES (?, ?, ?)`
	_, err = db.Exec(insertStmt, "张三", "技术部", 8000.00)
	if err != nil {
		log.Printf("插入数据失败: %v", err)
	}
	_, err = db.Exec(insertStmt, "李四", "技术部", 9500.00)
	if err != nil {
		log.Printf("插入数据失败: %v", err)
	}
	_, err = db.Exec(insertStmt, "王五", "销售部", 7000.00)
	if err != nil {
		log.Printf("插入数据失败: %v", err)
	}
	_, err = db.Exec(insertStmt, "赵六", "市场部", 8800.00)
	if err != nil {
		log.Printf("插入数据失败: %v", err)
	}
	_, err = db.Exec(insertStmt, "钱七", "技术部", 12000.00)
	if err != nil {
		log.Printf("插入数据失败: %v", err)
	}
	fmt.Println("示例数据插入完成。")

	// --- 查询所有部门为 "技术部" 的员工 ---
	fmt.Println("\n--- 查询所有部门为 \"技术部\" 的员工 ---")
	var techEmployees []Employee
	err = db.Select(&techEmployees, "SELECT id, name, department, salary FROM employees WHERE department = ?", "技术部")
	if err != nil {
		log.Fatalf("查询技术部员工失败: %v", err)
	}
	for _, emp := range techEmployees {
		fmt.Printf("ID: %d, 姓名: %s, 部门: %s, 工资: %.2f\n", emp.ID, emp.Name, emp.Department, emp.Salary)
	}

	// --- 查询工资最高的员工 ---
	fmt.Println("\n--- 查询工资最高的员工 ---")
	var highestPaidEmployee Employee
	err = db.Get(&highestPaidEmployee, "SELECT id, name, department, salary FROM employees ORDER BY salary DESC LIMIT 1")
	if err != nil {
		log.Fatalf("查询工资最高的员工失败: %v", err)
	}
	fmt.Printf("ID: %d, 姓名: %s, 部门: %s, 工资: %.2f\n", highestPaidEmployee.ID, highestPaidEmployee.Name, highestPaidEmployee.Department, highestPaidEmployee.Salary)
}
