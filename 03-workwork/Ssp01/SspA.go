package ssp

import (
	"fmt"
	"gorm.io/gorm"
)

// Student 学生模型
type Student struct {
	ID    uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name  string `gorm:"type:varchar(100);not null" json:"name"`
	Age   int    `gorm:"not null" json:"age"`
	Grade string `gorm:"type:varchar(50);not null" json:"grade"`
}

// TableName 指定表名
func (Student) TableName() string {
	return "students"
}

// CreateStudentsTable 创建students表
func CreateStudentsTable(db *gorm.DB) error {
	err := db.AutoMigrate(&Student{})
	if err != nil {
		return fmt.Errorf("创建表失败: %v", err)
	}
	fmt.Println("✅ students表创建成功")
	return nil
}

// InsertStudent 插入学生记录
func InsertStudent(db *gorm.DB, name string, age int, grade string) error {
	student := Student{
		Name:  name,
		Age:   age,
		Grade: grade,
	}

	result := db.Create(&student)
	if result.Error != nil {
		return fmt.Errorf("插入学生记录失败: %v", result.Error)
	}
	fmt.Printf("✅ 成功插入学生: %s, 年龄: %d, 年级: %s (ID: %d)\n", name, age, grade, student.ID)
	return nil
}

// QueryStudentsByAge 查询年龄大于指定值的学生
func QueryStudentsByAge(db *gorm.DB, minAge int) error {
	var students []Student
	
	result := db.Where("age > ?", minAge).Find(&students)
	if result.Error != nil {
		return fmt.Errorf("查询失败: %v", result.Error)
	}

	fmt.Printf("📋 年龄大于 %d 岁的学生信息:\n", minAge)
	fmt.Println("ID\t姓名\t年龄\t年级")
	fmt.Println("------------------------")
	
	if len(students) == 0 {
		fmt.Println("没有找到符合条件的学生")
		return nil
	}

	for _, student := range students {
		fmt.Printf("%d\t%s\t%d\t%s\n", student.ID, student.Name, student.Age, student.Grade)
	}

	return nil
}

// UpdateStudentGrade 更新学生年级
func UpdateStudentGrade(db *gorm.DB, name, newGrade string) error {
	result := db.Model(&Student{}).Where("name = ?", name).Update("grade", newGrade)
	if result.Error != nil {
		return fmt.Errorf("更新学生年级失败: %v", result.Error)
	}

	if result.RowsAffected > 0 {
		fmt.Printf("✅ 成功更新学生 %s 的年级为: %s\n", name, newGrade)
	} else {
		fmt.Printf("⚠️  未找到姓名为 %s 的学生\n", name)
	}
	return nil
}

// DeleteStudentsByAge 删除年龄小于指定值的学生
func DeleteStudentsByAge(db *gorm.DB, maxAge int) error {
	result := db.Where("age < ?", maxAge).Delete(&Student{})
	if result.Error != nil {
		return fmt.Errorf("删除学生记录失败: %v", result.Error)
	}

	fmt.Printf("✅ 成功删除 %d 条年龄小于 %d 岁的学生记录\n", result.RowsAffected, maxAge)
	return nil
}

// ShowAllStudents 显示所有学生信息
func ShowAllStudents(db *gorm.DB) error {
	var students []Student
	
	result := db.Order("id").Find(&students)
	if result.Error != nil {
		return fmt.Errorf("查询失败: %v", result.Error)
	}

	fmt.Println("📋 所有学生信息:")
	fmt.Println("ID\t姓名\t年龄\t年级")
	fmt.Println("------------------------")
	
	if len(students) == 0 {
		fmt.Println("表中暂无学生数据")
		return nil
	}

	for _, student := range students {
		fmt.Printf("%d\t%s\t%d\t%s\n", student.ID, student.Name, student.Age, student.Grade)
	}

	return nil
}

// RunCRUDDemo 运行CRUD演示
func RunCRUDDemo(db *gorm.DB) {
	fmt.Println("🚀 开始CRUD操作演示")
	//fmt.Println("==================================================")

	// 1. 创建表
	fmt.Println("\n1️⃣ 创建students表")
	err := CreateStudentsTable(db)
	if err != nil {
		panic(err)
	}

	// 2. 插入学生记录
	fmt.Println("\n2️⃣ 插入学生记录")
	err = InsertStudent(db, "张三", 20, "三年级")
	if err != nil {
		panic(err)
	}

	// 插入更多测试数据
	testStudents := []struct {
		name  string
		age   int
		grade string
	}{
		{"李四", 19, "二年级"},
		{"王五", 16, "一年级"},
		{"赵六", 22, "四年级"},
		{"孙七", 14, "一年级"},
	}

	for _, student := range testStudents {
		err = InsertStudent(db, student.name, student.age, student.grade)
		if err != nil {
			panic(err)
		}
	}

	// 显示所有学生
	fmt.Println("\n📊 当前所有学生:")
	err = ShowAllStudents(db)
	if err != nil {
		panic(err)
	}

	// 3. 查询年龄大于18的学生
	fmt.Println("\n3️⃣ 查询年龄大于18岁的学生")
	err = QueryStudentsByAge(db, 18)
	if err != nil {
		panic(err)
	}

	// 4. 更新张三的年级
	fmt.Println("\n4️⃣ 更新张三的年级")
	err = UpdateStudentGrade(db, "张三", "四年级")
	if err != nil {
		panic(err)
	}

	// 显示更新后的所有学生
	fmt.Println("\n📊 更新后的所有学生:")
	err = ShowAllStudents(db)
	if err != nil {
		panic(err)
	}

	// 5. 删除年龄小于15的学生
	fmt.Println("\n5️⃣ 删除年龄小于15岁的学生")
	err = DeleteStudentsByAge(db, 15)
	if err != nil {
		panic(err)
	}

	// 显示删除后的所有学生
	fmt.Println("\n📊 删除后的所有学生:")
	err = ShowAllStudents(db)
	if err != nil {
		panic(err)
	}

	fmt.Println("\n✅ CRUD操作演示完成!")
}
