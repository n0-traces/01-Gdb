package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gaa "workwork/Ga03"
)

func main() {

	// GORM数据库连接
	gormDB, err := gorm.Open(mysql.Open("root:root@tcp(192.168.x.x:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}

	/*
		// Sqlx数据库连接
		sqlxDB, err := sqlx.Connect("mysql", "root:root@tcp(192.168.x.x:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local")
		if err != nil {
			panic(err)
		}
		defer sqlxDB.Close()
	*/

	//fmt.Println("🎓 学生管理系统 - CRUD操作演示")
	//fmt.Println("==================================================")

	// 调用Ssp01中的CRUD演示函数
	//ssp.RunCRUDDemo(gormDB)

	//fmt.Println("🏦 银行转账系统演示")
	//fmt.Println("==================================================")

	// 调用Ssp01中的银行转账演示函数
	//ssp.RunBankDemo(gormDB)

	//fmt.Println("👥 Sqlx员工查询系统演示")
	//fmt.Println("==================================================")

	// 调用GswS02中的Sqlx演示函数
	//gsw.RunSqlxDemo(sqlxDB)

	//fmt.Println("\n" + strings.Repeat("=", 60))
	//fmt.Println("📚 Sqlx书籍查询系统演示")
	//fmt.Println("==================================================")

	// 调用GswS02中的书籍查询演示函数
	//gsw.RunBookDemo(sqlxDB)

	/*
		fmt.Println("📚 在main中映射Book结构体演示")
		fmt.Println("==================================================")

		// 1. 在main中直接创建Book结构体实例
		fmt.Println("\n1️⃣ 创建Book结构体实例:")
		book1 := gsw.Book{
			ID:     1,
			Title:  "Go语言实战",
			Author: "张三",
			Price:  89.00,
		}

		book2 := gsw.Book{
			ID:     2,
			Title:  "Python编程",
			Author: "李四",
			Price:  65.50,
		}

		book3 := gsw.Book{
			ID:     3,
			Title:  "Java核心技术",
			Author: "王五",
			Price:  120.00,
		}

		// 2. 显示创建的Book实例
		fmt.Println("📋 创建的Book实例:")
		fmt.Println("ID\t书名\t\t\t作者\t\t价格(元)")
		fmt.Println("--------------------------------------------------------")
		fmt.Printf("%d\t%-20s\t%-10s\t%.2f\n", book1.ID, book1.Title, book1.Author, book1.Price)
		fmt.Printf("%d\t%-20s\t%-10s\t%.2f\n", book2.ID, book2.Title, book2.Author, book2.Price)
		fmt.Printf("%d\t%-20s\t%-10s\t%.2f\n", book3.ID, book3.Title, book3.Author, book3.Price)

		// 3. 创建Book切片
		fmt.Println("\n2️⃣ 创建Book切片:")
		books := []gsw.Book{book1, book2, book3}

		fmt.Printf("📚 Book切片包含 %d 本书:\n", len(books))
		for i, book := range books {
			fmt.Printf("  %d. %s (作者: %s, 价格: %.2f元)\n", i+1, book.Title, book.Author, book.Price)
		}

		// 4. 在main中进行简单的查询逻辑
		fmt.Println("\n3️⃣ 在main中进行查询逻辑:")

		// 查询价格大于50元的书籍
		fmt.Println("🔍 查询价格大于50元的书籍:")
		expensiveBooks := []gsw.Book{}
		for _, book := range books {
			if book.Price > 50.0 {
				expensiveBooks = append(expensiveBooks, book)
			}
		}

		if len(expensiveBooks) > 0 {
			fmt.Printf("找到 %d 本价格大于50元的书籍:\n", len(expensiveBooks))
			for _, book := range expensiveBooks {
				fmt.Printf("  - %s (%.2f元)\n", book.Title, book.Price)
			}
		} else {
			fmt.Println("没有找到价格大于50元的书籍")
		}

		// 查询指定作者的书籍
		fmt.Println("\n🔍 查询作者'张三'的书籍:")
		authorBooks := []gsw.Book{}
		for _, book := range books {
			if book.Author == "张三" {
				authorBooks = append(authorBooks, book)
			}
		}

		if len(authorBooks) > 0 {
			fmt.Printf("找到 %d 本作者'张三'的书籍:\n", len(authorBooks))
			for _, book := range authorBooks {
				fmt.Printf("  - %s (%.2f元)\n", book.Title, book.Price)
			}
		} else {
			fmt.Println("没有找到作者'张三'的书籍")
		}

		// 5. 计算平均价格
		fmt.Println("\n4️⃣ 计算平均价格:")
		totalPrice := 0.0
		for _, book := range books {
			totalPrice += book.Price
		}
		avgPrice := totalPrice / float64(len(books))
		fmt.Printf("📊 所有书籍的平均价格: %.2f元\n", avgPrice)

		fmt.Println("\n✅ 在main中映射Book结构体演示完成!")
		fmt.Println("\n" + strings.Repeat("=", 60))

		fmt.Println("\n✅ Book结构体创建成功!")
		fmt.Println("💡 要运行完整的数据库演示，请执行: go run main.go")

		fmt.Println("📚 Sqlx书籍查询系统演示")
		fmt.Println("==================================================")

		// 调用GswS02中的书籍查询演示函数
		gsw.RunBookDemo(sqlxDB)
	*/

	// 博客系统GORM演示
	fmt.Println("📝 博客系统GORM演示")
	fmt.Println("==================================================")

	// 调用Ga03中的博客系统演示函数
	gaa.RunBlogDemo(gormDB)

}
