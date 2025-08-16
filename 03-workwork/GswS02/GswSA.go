package gsw

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
)

// Employee 员工模型
type Employee struct {
	ID         int    `db:"id" json:"id"`
	Name       string `db:"name" json:"name"`
	Department string `db:"department" json:"department"`
	Salary     int    `db:"salary" json:"salary"`
}

// CreateEmployeesTable 创建employees表
func CreateEmployeesTable(db *sqlx.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS employees (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		department VARCHAR(50) NOT NULL,
		salary INT NOT NULL
	)`

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("创建employees表失败: %v", err)
	}
	fmt.Println("✅ employees表创建成功")
	return nil
}

// InsertEmployee 插入员工记录
func InsertEmployee(db *sqlx.DB, name, department string, salary int) error {
	query := `INSERT INTO employees (name, department, salary) VALUES (?, ?, ?)`
	
	_, err := db.Exec(query, name, department, salary)
	if err != nil {
		return fmt.Errorf("插入员工记录失败: %v", err)
	}
	
	fmt.Printf("✅ 成功插入员工: %s, 部门: %s, 工资: %d\n", name, department, salary)
	return nil
}

// QueryEmployeesByDepartment 查询指定部门的所有员工
func QueryEmployeesByDepartment(db *sqlx.DB, department string) ([]Employee, error) {
	query := `SELECT id, name, department, salary FROM employees WHERE department = ?`
	
	var employees []Employee
	err := db.Select(&employees, query, department)
	if err != nil {
		return nil, fmt.Errorf("查询员工失败: %v", err)
	}
	
	return employees, nil
}

// QueryHighestSalaryEmployee 查询工资最高的员工
func QueryHighestSalaryEmployee(db *sqlx.DB) (*Employee, error) {
	query := `SELECT id, name, department, salary FROM employees ORDER BY salary DESC LIMIT 1`
	
	var employee Employee
	err := db.Get(&employee, query)
	if err != nil {
		return nil, fmt.Errorf("查询最高工资员工失败: %v", err)
	}
	
	return &employee, nil
}

// ShowAllEmployees 显示所有员工信息
func ShowAllEmployees(db *sqlx.DB) error {
	query := `SELECT id, name, department, salary FROM employees ORDER BY id`
	
	var employees []Employee
	err := db.Select(&employees, query)
	if err != nil {
		return fmt.Errorf("查询所有员工失败: %v", err)
	}
	
	fmt.Println("📋 所有员工信息:")
	fmt.Println("ID\t姓名\t\t部门\t\t工资")
	fmt.Println("----------------------------------------")
	
	if len(employees) == 0 {
		fmt.Println("暂无员工数据")
		return nil
	}
	
	for _, emp := range employees {
		fmt.Printf("%d\t%s\t\t%s\t\t%d\n", emp.ID, emp.Name, emp.Department, emp.Salary)
	}
	
	return nil
}

// RunSqlxDemo 运行Sqlx演示
func RunSqlxDemo(db *sqlx.DB) {
	fmt.Println("👥 Sqlx员工查询系统演示")
	fmt.Println("==================================================")

	// 1. 创建employees表
	fmt.Println("\n1️⃣ 创建employees表")
	err := CreateEmployeesTable(db)
	if err != nil {
		panic(err)
	}

	// 2. 插入测试员工数据
	fmt.Println("\n2️⃣ 插入测试员工数据")
	testEmployees := []struct {
		name       string
		department string
		salary     int
	}{
		{"张三", "技术部", 8000},
		{"李四", "技术部", 9000},
		{"王五", "技术部", 7500},
		{"赵六", "销售部", 6000},
		{"孙七", "销售部", 7000},
		{"周八", "技术部", 12000},
		{"吴九", "人事部", 5000},
		{"郑十", "技术部", 8500},
	}

	for _, emp := range testEmployees {
		err = InsertEmployee(db, emp.name, emp.department, emp.salary)
		if err != nil {
			panic(err)
		}
	}

	// 显示所有员工
	fmt.Println("\n📊 当前所有员工:")
	err = ShowAllEmployees(db)
	if err != nil {
		panic(err)
	}

	// 3. 查询技术部的所有员工
	fmt.Println("\n3️⃣ 查询技术部的所有员工")
	techEmployees, err := QueryEmployeesByDepartment(db, "技术部")
	if err != nil {
		panic(err)
	}

	fmt.Printf("📋 技术部员工信息 (共%d人):\n", len(techEmployees))
	fmt.Println("ID\t姓名\t\t部门\t\t工资")
	fmt.Println("----------------------------------------")
	
	for _, emp := range techEmployees {
		fmt.Printf("%d\t%s\t\t%s\t\t%d\n", emp.ID, emp.Name, emp.Department, emp.Salary)
	}

	// 4. 查询工资最高的员工
	fmt.Println("\n4️⃣ 查询工资最高的员工")
	highestSalaryEmployee, err := QueryHighestSalaryEmployee(db)
	if err != nil {
		panic(err)
	}

	fmt.Println("🏆 工资最高的员工:")
	fmt.Printf("ID: %d\n", highestSalaryEmployee.ID)
	fmt.Printf("姓名: %s\n", highestSalaryEmployee.Name)
	fmt.Printf("部门: %s\n", highestSalaryEmployee.Department)
	fmt.Printf("工资: %d\n", highestSalaryEmployee.Salary)

	// 5. 查询其他部门的员工
	fmt.Println("\n5️⃣ 查询销售部的员工")
	salesEmployees, err := QueryEmployeesByDepartment(db, "销售部")
	if err != nil {
		panic(err)
	}

	fmt.Printf("📋 销售部员工信息 (共%d人):\n", len(salesEmployees))
	fmt.Println("ID\t姓名\t\t部门\t\t工资")
	fmt.Println("----------------------------------------")
	
	for _, emp := range salesEmployees {
		fmt.Printf("%d\t%s\t\t%s\t\t%d\n", emp.ID, emp.Name, emp.Department, emp.Salary)
	}

	fmt.Println("\n✅ Sqlx员工查询演示完成!")
}
