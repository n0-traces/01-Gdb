package gsw

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
)

// Book 书籍模型
type Book struct {
	ID     int     `db:"id" json:"id"`
	Title  string  `db:"title" json:"title"`
	Author string  `db:"author" json:"author"`
	Price  float64 `db:"price" json:"price"`
}

// CreateBooksTable 创建books表
func CreateBooksTable(db *sqlx.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS books (
		id INT AUTO_INCREMENT PRIMARY KEY,
		title VARCHAR(200) NOT NULL,
		author VARCHAR(100) NOT NULL,
		price DECIMAL(10,2) NOT NULL
	)`

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("创建books表失败: %v", err)
	}
	fmt.Println("✅ books表创建成功")
	return nil
}

// InsertBook 插入书籍记录
func InsertBook(db *sqlx.DB, title, author string, price float64) error {
	query := `INSERT INTO books (title, author, price) VALUES (?, ?, ?)`
	
	_, err := db.Exec(query, title, author, price)
	if err != nil {
		return fmt.Errorf("插入书籍记录失败: %v", err)
	}
	
	fmt.Printf("✅ 成功插入书籍: %s, 作者: %s, 价格: %.2f元\n", title, author, price)
	return nil
}

// QueryBooksByPrice 查询价格大于指定值的书籍
func QueryBooksByPrice(db *sqlx.DB, minPrice float64) ([]Book, error) {
	query := `SELECT id, title, author, price FROM books WHERE price > ? ORDER BY price DESC`
	
	var books []Book
	err := db.Select(&books, query, minPrice)
	if err != nil {
		return nil, fmt.Errorf("查询书籍失败: %v", err)
	}
	
	return books, nil
}

// QueryBooksByAuthor 查询指定作者的书籍
func QueryBooksByAuthor(db *sqlx.DB, author string) ([]Book, error) {
	query := `SELECT id, title, author, price FROM books WHERE author = ? ORDER BY price DESC`
	
	var books []Book
	err := db.Select(&books, query, author)
	if err != nil {
		return nil, fmt.Errorf("查询作者书籍失败: %v", err)
	}
	
	return books, nil
}

// QueryExpensiveBooks 查询价格最高的书籍
func QueryExpensiveBooks(db *sqlx.DB, limit int) ([]Book, error) {
	query := `SELECT id, title, author, price FROM books ORDER BY price DESC LIMIT ?`
	
	var books []Book
	err := db.Select(&books, query, limit)
	if err != nil {
		return nil, fmt.Errorf("查询最贵书籍失败: %v", err)
	}
	
	return books, nil
}

// QueryBooksByPriceRange 查询价格在指定范围内的书籍
func QueryBooksByPriceRange(db *sqlx.DB, minPrice, maxPrice float64) ([]Book, error) {
	query := `SELECT id, title, author, price FROM books WHERE price BETWEEN ? AND ? ORDER BY price ASC`
	
	var books []Book
	err := db.Select(&books, query, minPrice, maxPrice)
	if err != nil {
		return nil, fmt.Errorf("查询价格范围书籍失败: %v", err)
	}
	
	return books, nil
}

// GetAveragePrice 获取所有书籍的平均价格
func GetAveragePrice(db *sqlx.DB) (float64, error) {
	query := `SELECT AVG(price) as avg_price FROM books`
	
	var avgPrice float64
	err := db.Get(&avgPrice, query)
	if err != nil {
		return 0, fmt.Errorf("查询平均价格失败: %v", err)
	}
	
	return avgPrice, nil
}

// ShowAllBooks 显示所有书籍信息
func ShowAllBooks(db *sqlx.DB) error {
	query := `SELECT id, title, author, price FROM books ORDER BY id`
	
	var books []Book
	err := db.Select(&books, query)
	if err != nil {
		return fmt.Errorf("查询所有书籍失败: %v", err)
	}
	
	fmt.Println("📚 所有书籍信息:")
	fmt.Println("ID\t书名\t\t\t作者\t\t价格(元)")
	fmt.Println("--------------------------------------------------------")
	
	if len(books) == 0 {
		fmt.Println("暂无书籍数据")
		return nil
	}
	
	for _, book := range books {
		fmt.Printf("%d\t%-20s\t%-10s\t%.2f\n", book.ID, book.Title, book.Author, book.Price)
	}
	
	return nil
}

// RunBookDemo 运行书籍查询演示
func RunBookDemo(db *sqlx.DB) {
	fmt.Println("📚 Sqlx书籍查询系统演示")
	fmt.Println("==================================================")

	// 1. 创建books表
	fmt.Println("\n1️⃣ 创建books表")
	err := CreateBooksTable(db)
	if err != nil {
		panic(err)
	}

	// 2. 插入测试书籍数据
	fmt.Println("\n2️⃣ 插入测试书籍数据")
	testBooks := []struct {
		title  string
		author string
		price  float64
	}{
		{"Go语言实战", "张三", 89.00},
		{"Python编程", "李四", 65.50},
		{"Java核心技术", "王五", 120.00},
		{"数据结构与算法", "赵六", 78.00},
		{"数据库系统概论", "孙七", 95.00},
		{"计算机网络", "周八", 85.00},
		{"操作系统", "吴九", 110.00},
		{"软件工程", "郑十", 72.00},
		{"机器学习", "张三", 150.00},
		{"深度学习", "李四", 180.00},
		{"算法导论", "王五", 200.00},
		{"设计模式", "赵六", 88.00},
	}

	for _, book := range testBooks {
		err = InsertBook(db, book.title, book.author, book.price)
		if err != nil {
			panic(err)
		}
	}

	// 显示所有书籍
	fmt.Println("\n📊 当前所有书籍:")
	err = ShowAllBooks(db)
	if err != nil {
		panic(err)
	}

	// 3. 查询价格大于50元的书籍
	fmt.Println("\n3️⃣ 查询价格大于50元的书籍")
	expensiveBooks, err := QueryBooksByPrice(db, 50.0)
	if err != nil {
		panic(err)
	}

	fmt.Printf("📋 价格大于50元的书籍 (共%d本):\n", len(expensiveBooks))
	fmt.Println("ID\t书名\t\t\t作者\t\t价格(元)")
	fmt.Println("--------------------------------------------------------")
	
	for _, book := range expensiveBooks {
		fmt.Printf("%d\t%-20s\t%-10s\t%.2f\n", book.ID, book.Title, book.Author, book.Price)
	}

	// 4. 查询指定作者的书籍
	fmt.Println("\n4️⃣ 查询作者'张三'的书籍")
	authorBooks, err := QueryBooksByAuthor(db, "张三")
	if err != nil {
		panic(err)
	}

	fmt.Printf("📋 作者'张三'的书籍 (共%d本):\n", len(authorBooks))
	fmt.Println("ID\t书名\t\t\t作者\t\t价格(元)")
	fmt.Println("--------------------------------------------------------")
	
	for _, book := range authorBooks {
		fmt.Printf("%d\t%-20s\t%-10s\t%.2f\n", book.ID, book.Title, book.Author, book.Price)
	}

	// 5. 查询最贵的3本书
	fmt.Println("\n5️⃣ 查询最贵的3本书")
	topBooks, err := QueryExpensiveBooks(db, 3)
	if err != nil {
		panic(err)
	}

	fmt.Println("🏆 最贵的3本书:")
	fmt.Println("ID\t书名\t\t\t作者\t\t价格(元)")
	fmt.Println("--------------------------------------------------------")
	
	for _, book := range topBooks {
		fmt.Printf("%d\t%-20s\t%-10s\t%.2f\n", book.ID, book.Title, book.Author, book.Price)
	}

	// 6. 查询价格在80-120元之间的书籍
	fmt.Println("\n6️⃣ 查询价格在80-120元之间的书籍")
	rangeBooks, err := QueryBooksByPriceRange(db, 80.0, 120.0)
	if err != nil {
		panic(err)
	}

	fmt.Printf("📋 价格在80-120元之间的书籍 (共%d本):\n", len(rangeBooks))
	fmt.Println("ID\t书名\t\t\t作者\t\t价格(元)")
	fmt.Println("--------------------------------------------------------")
	
	for _, book := range rangeBooks {
		fmt.Printf("%d\t%-20s\t%-10s\t%.2f\n", book.ID, book.Title, book.Author, book.Price)
	}

	// 7. 获取平均价格
	fmt.Println("\n7️⃣ 获取所有书籍的平均价格")
	avgPrice, err := GetAveragePrice(db)
	if err != nil {
		panic(err)
	}

	fmt.Printf("📊 所有书籍的平均价格: %.2f元\n", avgPrice)

	fmt.Println("\n✅ Sqlx书籍查询演示完成!")
}
