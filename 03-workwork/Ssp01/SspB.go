package ssp

import (
	"fmt"
	"gorm.io/gorm"
)

// Account 账户模型
type Account struct {
	ID      uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Balance float64 `gorm:"type:decimal(10,2);not null;default:0" json:"balance"`
}

// TableName 指定表名
func (Account) TableName() string {
	return "accounts"
}

// Transaction 交易记录模型
type Transaction struct {
	ID            uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	FromAccountID uint    `gorm:"not null" json:"from_account_id"`
	ToAccountID   uint    `gorm:"not null" json:"to_account_id"`
	Amount        float64 `gorm:"type:decimal(10,2);not null" json:"amount"`
}

// TableName 指定表名
func (Transaction) TableName() string {
	return "transactions"
}

// CreateBankTables 创建银行相关表
func CreateBankTables(db *gorm.DB) error {
	err := db.AutoMigrate(&Account{}, &Transaction{})
	if err != nil {
		return fmt.Errorf("创建银行表失败: %v", err)
	}
	fmt.Println("✅ 银行表创建成功 (accounts, transactions)")
	return nil
}

// CreateAccount 创建账户
func CreateAccount(db *gorm.DB, balance float64) (*Account, error) {
	account := Account{
		Balance: balance,
	}

	result := db.Create(&account)
	if result.Error != nil {
		return nil, fmt.Errorf("创建账户失败: %v", result.Error)
	}

	fmt.Printf("✅ 成功创建账户: ID=%d, 余额=%.2f\n", account.ID, account.Balance)
	return &account, nil
}

// GetAccountByID 根据ID获取账户信息
func GetAccountByID(db *gorm.DB, accountID uint) (*Account, error) {
	var account Account
	result := db.First(&account, accountID)
	if result.Error != nil {
		return nil, fmt.Errorf("获取账户失败: %v", result.Error)
	}
	return &account, nil
}

// TransferMoney 转账操作（使用事务）
func TransferMoney(db *gorm.DB, fromAccountID, toAccountID uint, amount float64) error {
	fmt.Printf("🔄 开始转账: 账户%d → 账户%d, 金额: %.2f\n", fromAccountID, toAccountID, amount)

	// 开始事务
	tx := db.Begin()
	if tx.Error != nil {
		return fmt.Errorf("开始事务失败: %v", tx.Error)
	}

	// 确保事务结束时进行清理
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 1. 检查转出账户是否存在并获取余额
	var fromAccount Account
	if err := tx.First(&fromAccount, fromAccountID).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("转出账户不存在: %v", err)
	}

	// 2. 检查转入账户是否存在
	var toAccount Account
	if err := tx.First(&toAccount, toAccountID).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("转入账户不存在: %v", err)
	}

	// 3. 检查余额是否足够
	if fromAccount.Balance < amount {
		tx.Rollback()
		return fmt.Errorf("余额不足: 账户%d当前余额%.2f, 需要%.2f", fromAccountID, fromAccount.Balance, amount)
	}

	// 4. 扣除转出账户余额
	if err := tx.Model(&fromAccount).Update("balance", fromAccount.Balance-amount).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("扣除转出账户余额失败: %v", err)
	}

	// 5. 增加转入账户余额
	if err := tx.Model(&toAccount).Update("balance", toAccount.Balance+amount).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("增加转入账户余额失败: %v", err)
	}

	// 6. 记录交易信息
	transaction := Transaction{
		FromAccountID: fromAccountID,
		ToAccountID:   toAccountID,
		Amount:        amount,
	}

	if err := tx.Create(&transaction).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("记录交易信息失败: %v", err)
	}

	// 7. 提交事务
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("提交事务失败: %v", err)
	}

	fmt.Printf("✅ 转账成功! 交易ID: %d\n", transaction.ID)
	fmt.Printf("   账户%d余额: %.2f → %.2f\n", fromAccountID, fromAccount.Balance, fromAccount.Balance-amount)
	fmt.Printf("   账户%d余额: %.2f → %.2f\n", toAccountID, toAccount.Balance, toAccount.Balance+amount)

	return nil
}

// ShowAllAccounts 显示所有账户信息
func ShowAllAccounts(db *gorm.DB) error {
	var accounts []Account
	
	result := db.Order("id").Find(&accounts)
	if result.Error != nil {
		return fmt.Errorf("查询账户失败: %v", result.Error)
	}

	fmt.Println("📋 所有账户信息:")
	fmt.Println("ID\t余额")
	fmt.Println("----------------")
	
	if len(accounts) == 0 {
		fmt.Println("暂无账户数据")
		return nil
	}

	for _, account := range accounts {
		fmt.Printf("%d\t%.2f\n", account.ID, account.Balance)
	}

	return nil
}

// ShowAllTransactions 显示所有交易记录
func ShowAllTransactions(db *gorm.DB) error {
	var transactions []Transaction
	
	result := db.Order("id").Find(&transactions)
	if result.Error != nil {
		return fmt.Errorf("查询交易记录失败: %v", result.Error)
	}

	fmt.Println("📋 所有交易记录:")
	fmt.Println("ID\t转出账户\t转入账户\t金额")
	fmt.Println("----------------------------------------")
	
	if len(transactions) == 0 {
		fmt.Println("暂无交易记录")
		return nil
	}

	for _, transaction := range transactions {
		fmt.Printf("%d\t%d\t\t%d\t\t%.2f\n", 
			transaction.ID, 
			transaction.FromAccountID, 
			transaction.ToAccountID, 
			transaction.Amount)
	}

	return nil
}

// RunBankDemo 运行银行转账演示
func RunBankDemo(db *gorm.DB) {
	fmt.Println("🏦 银行转账系统演示")
	fmt.Println("==================================================")

	// 1. 创建银行表
	fmt.Println("\n1️⃣ 创建银行表")
	err := CreateBankTables(db)
	if err != nil {
		panic(err)
	}

	// 2. 创建测试账户
	fmt.Println("\n2️⃣ 创建测试账户")
	accountA, err := CreateAccount(db, 500.00) // 账户A，余额500元
	if err != nil {
		panic(err)
	}

	accountB, err := CreateAccount(db, 200.00) // 账户B，余额200元
	if err != nil {
		panic(err)
	}

	accountC, err := CreateAccount(db, 50.00)  // 账户C，余额50元
	if err != nil {
		panic(err)
	}

	// 显示初始账户状态
	fmt.Println("\n📊 初始账户状态:")
	err = ShowAllAccounts(db)
	if err != nil {
		panic(err)
	}

	// 3. 成功转账演示
	fmt.Println("\n3️⃣ 成功转账演示")
	err = TransferMoney(db, accountA.ID, accountB.ID, 100.00)
	if err != nil {
		panic(err)
	}

	// 显示转账后的账户状态
	fmt.Println("\n📊 转账后的账户状态:")
	err = ShowAllAccounts(db)
	if err != nil {
		panic(err)
	}

	// 显示交易记录
	fmt.Println("\n📋 交易记录:")
	err = ShowAllTransactions(db)
	if err != nil {
		panic(err)
	}

	// 4. 余额不足转账演示
	fmt.Println("\n4️⃣ 余额不足转账演示")
	err = TransferMoney(db, accountC.ID, accountA.ID, 100.00)
	if err != nil {
		fmt.Printf("❌ 转账失败（预期）: %v\n", err)
	} else {
		fmt.Println("⚠️  转账成功（意外）")
	}

	// 显示最终账户状态（应该没有变化）
	fmt.Println("\n📊 最终账户状态（余额不足转账后）:")
	err = ShowAllAccounts(db)
	if err != nil {
		panic(err)
	}

	// 5. 再次成功转账演示
	fmt.Println("\n5️⃣ 再次成功转账演示")
	err = TransferMoney(db, accountB.ID, accountC.ID, 50.00)
	if err != nil {
		panic(err)
	}

	// 显示最终状态
	fmt.Println("\n📊 最终账户状态:")
	err = ShowAllAccounts(db)
	if err != nil {
		panic(err)
	}

	fmt.Println("\n📋 最终交易记录:")
	err = ShowAllTransactions(db)
	if err != nil {
		panic(err)
	}

	fmt.Println("\n✅ 银行转账演示完成!")
}
