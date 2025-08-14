package main

import (
	"fmt"
	"time"
)

// Person 基础结构体
type Person struct {
	Name string
	Age  int
}

// Person 的方法
func (p Person) GetName() string {
	return p.Name
}

func (p Person) GetAge() int {
	return p.Age
}

func (p Person) IsAdult() bool {
	return p.Age >= 18
}

func (p Person) String() string {
	return fmt.Sprintf("Person{Name: %s, Age: %d}", p.Name, p.Age)
}

// Employee 员工结构体，组合了 Person
type Employee struct {
	Person      // 匿名嵌入，实现组合
	EmployeeID  string
	Department  string
	Salary      float64
	HireDate    time.Time
}

// Employee 的构造函数
func NewEmployee(name string, age int, employeeID, department string, salary float64) *Employee {
	return &Employee{
		Person: Person{
			Name: name,
			Age:  age,
		},
		EmployeeID: employeeID,
		Department: department,
		Salary:     salary,
		HireDate:   time.Now(),
	}
}

// Employee 实现 PrintInfo 方法
func (e Employee) PrintInfo() {
	fmt.Println("=== 员工信息 ===")
	fmt.Printf("姓名: %s\n", e.Name)           // 直接访问Person的字段
	fmt.Printf("年龄: %d\n", e.Age)            // 直接访问Person的字段
	fmt.Printf("员工ID: %s\n", e.EmployeeID)
	fmt.Printf("部门: %s\n", e.Department)
	fmt.Printf("薪资: %.2f\n", e.Salary)
	fmt.Printf("入职时间: %s\n", e.HireDate.Format("2006-01-02 15:04:05"))
	fmt.Printf("是否成年: %t\n", e.IsAdult())  // 调用Person的方法
	fmt.Println("================")
}

// Employee 的其他方法
func (e Employee) GetEmployeeID() string {
	return e.EmployeeID
}

func (e Employee) GetDepartment() string {
	return e.Department
}

func (e Employee) GetSalary() float64 {
	return e.Salary
}

func (e Employee) GetHireDate() time.Time {
	return e.HireDate
}

func (e Employee) WorkYears() int {
	return int(time.Since(e.HireDate).Hours() / 24 / 365)
}

func (e Employee) String() string {
	return fmt.Sprintf("Employee{Name: %s, Age: %d, ID: %s, Dept: %s, Salary: %.2f}", 
		e.Name, e.Age, e.EmployeeID, e.Department, e.Salary)
}

// Manager 经理结构体，组合了 Employee
type Manager struct {
	Employee           // 匿名嵌入Employee
	TeamSize          int
	ManagementLevel   string
}

// Manager 的构造函数
func NewManager(name string, age int, employeeID, department string, salary float64, teamSize int, level string) *Manager {
	return &Manager{
		Employee: Employee{
			Person: Person{
				Name: name,
				Age:  age,
			},
			EmployeeID: employeeID,
			Department: department,
			Salary:     salary,
			HireDate:   time.Now(),
		},
		TeamSize:        teamSize,
		ManagementLevel: level,
	}
}

// Manager 重写 PrintInfo 方法
func (m Manager) PrintInfo() {
	fmt.Println("=== 经理信息 ===")
	fmt.Printf("姓名: %s\n", m.Name)
	fmt.Printf("年龄: %d\n", m.Age)
	fmt.Printf("员工ID: %s\n", m.EmployeeID)
	fmt.Printf("部门: %s\n", m.Department)
	fmt.Printf("薪资: %.2f\n", m.Salary)
	fmt.Printf("入职时间: %s\n", m.HireDate.Format("2006-01-02 15:04:05"))
	fmt.Printf("团队规模: %d人\n", m.TeamSize)
	fmt.Printf("管理级别: %s\n", m.ManagementLevel)
	fmt.Printf("工作年限: %d年\n", m.WorkYears())
	fmt.Printf("是否成年: %t\n", m.IsAdult())
	fmt.Println("================")
}

// Manager 特有的方法
func (m Manager) GetTeamSize() int {
	return m.TeamSize
}

func (m Manager) GetManagementLevel() string {
	return m.ManagementLevel
}

func (m Manager) String() string {
	return fmt.Sprintf("Manager{Name: %s, ID: %s, Dept: %s, TeamSize: %d, Level: %s}", 
		m.Name, m.EmployeeID, m.Department, m.TeamSize, m.ManagementLevel)
}

// 演示组合的通用函数
func printPersonInfo(p Person) {
	fmt.Printf("人员信息: %v\n", p)
	fmt.Printf("姓名: %s, 年龄: %d, 是否成年: %t\n", p.GetName(), p.GetAge(), p.IsAdult())
}

func printEmployeeInfo(e Employee) {
	fmt.Printf("员工信息: %v\n", e)
	fmt.Printf("员工ID: %s, 部门: %s, 薪资: %.2f\n", e.GetEmployeeID(), e.GetDepartment(), e.GetSalary())
}

func main() {
	fmt.Println("=== Go组合模式示例 ===")
	fmt.Println("使用组合方式创建Person和Employee结构体")
	fmt.Println()

	// 创建Person实例
	person := Person{
		Name: "张三",
		Age:  25,
	}
	fmt.Println("创建Person实例:")
	printPersonInfo(person)
	fmt.Println()

	// 创建Employee实例（使用构造函数）
	employee := NewEmployee("李四", 30, "EMP001", "技术部", 8000.0)
	fmt.Println("创建Employee实例:")
	employee.PrintInfo()
	fmt.Println()

	// 创建Manager实例
	manager := NewManager("王五", 35, "MGR001", "技术部", 15000.0, 8, "高级经理")
	fmt.Println("创建Manager实例:")
	manager.PrintInfo()
	fmt.Println()

	// 演示组合的层次结构
	fmt.Println("=== 组合层次结构演示 ===")
	
	// 创建不同类型的员工
	employees := []interface{}{
		person,
		employee,
		manager,
	}
	
	for i, emp := range employees {
		fmt.Printf("对象 %d:\n", i+1)
		switch v := emp.(type) {
		case Person:
			fmt.Printf("  类型: Person\n")
			fmt.Printf("  信息: %v\n", v)
		case Employee:
			fmt.Printf("  类型: Employee\n")
			fmt.Printf("  信息: %v\n", v)
		case Manager:
			fmt.Printf("  类型: Manager\n")
			fmt.Printf("  信息: %v\n", v)
		}
		fmt.Println()
	}

	// 演示方法调用
	fmt.Println("=== 方法调用演示 ===")
	
	// Person的方法
	fmt.Printf("Person方法调用:\n")
	fmt.Printf("  GetName(): %s\n", person.GetName())
	fmt.Printf("  GetAge(): %d\n", person.GetAge())
	fmt.Printf("  IsAdult(): %t\n", person.IsAdult())
	fmt.Println()

	// Employee的方法（包括继承的Person方法）
	fmt.Printf("Employee方法调用:\n")
	fmt.Printf("  GetName(): %s\n", employee.GetName())        // 继承自Person
	fmt.Printf("  GetAge(): %d\n", employee.GetAge())          // 继承自Person
	fmt.Printf("  IsAdult(): %t\n", employee.IsAdult())        // 继承自Person
	fmt.Printf("  GetEmployeeID(): %s\n", employee.GetEmployeeID())
	fmt.Printf("  GetDepartment(): %s\n", employee.GetDepartment())
	fmt.Printf("  WorkYears(): %d\n", employee.WorkYears())
	fmt.Println()

	// Manager的方法（包括继承的Employee和Person方法）
	fmt.Printf("Manager方法调用:\n")
	fmt.Printf("  GetName(): %s\n", manager.GetName())         // 继承自Person
	fmt.Printf("  GetEmployeeID(): %s\n", manager.GetEmployeeID()) // 继承自Employee
	fmt.Printf("  GetTeamSize(): %d\n", manager.GetTeamSize())
	fmt.Printf("  GetManagementLevel(): %s\n", manager.GetManagementLevel())
	fmt.Printf("  WorkYears(): %d\n", manager.WorkYears())     // 继承自Employee
}
