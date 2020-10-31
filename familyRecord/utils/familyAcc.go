package utils

import "fmt"

// FamilyAccount 家庭账户结构体
type FamilyAccount struct {
	// 声明必须的字段
	// 接收用户输入的功能选项
	key string
	// 控制循环是否继续
	loop bool
	// 初始余额
	balance float64
	// 每次收支金额
	money float64
	// 收支说明
	note string
	// 记录是否有收支
	flag bool
	// 收支详情记录
	details string
}

// 封装记账相关的方法

// showDetails 显示明细
func (acc *FamilyAccount) showDetails() {
	fmt.Println("------------------当前收支明细记录-----------------")
	if !acc.flag {
		fmt.Println("当前没有收支明细，请先添加")
	} else {
		fmt.Println(acc.details)
	}
}

// registerIncoming  登记收入
func (acc *FamilyAccount) registerIncoming() {
	fmt.Println("本次收入金额：")
	fmt.Scanln(&acc.money)
	acc.balance += acc.money
	fmt.Println("本次收入说明：")
	fmt.Scanln(&acc.note)
	acc.details += fmt.Sprintf("\n收入\t\t%v\t\t%v\t\t%v", acc.balance, acc.money, acc.note)
	acc.flag = true
}

// registerPurchases 登记支出
func (acc *FamilyAccount) registerPurchases() {
	fmt.Println("本次支出金额：")
	fmt.Scanln(&acc.money)
	// 支出需要检查余额
	if acc.money > acc.balance {
		fmt.Println("余额不足，操作失败....")
	} else {
		acc.balance -= acc.money
		fmt.Println("本次支出说明：")
		fmt.Scanln(&acc.note)
		acc.details += fmt.Sprintf("\n支出\t\t%v\t\t%v\t\t%v", acc.balance, acc.money, acc.note)
		acc.flag = true
	}
}

// exit 退出软件
func (acc *FamilyAccount) exit() {
	fmt.Println("请确定是否要退出？ y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("输入指令有误，请重新输入。 y/n")
	}
	if choice == "y" {
		acc.loop = false
	}
}

// NewAccount 编写要给工厂模式的构造方法，返回一个构造体的实例
func NewAccount() *FamilyAccount {
	return &FamilyAccount{
		key:     "",
		loop:    true,
		balance: 10000.0,
		money:   0.0,
		note:    "",
		flag:    false,
		details: "收支\t\t账户金额\t\t收支金额\t\t说明",
	}
}

// MainMenu 显示主菜单
func (acc *FamilyAccount) MainMenu() {
	for {
		fmt.Println("\n---------------------家庭收支软件-------------------")
		fmt.Println("                      1.收支明细")
		fmt.Println("                      2.登记收入")
		fmt.Println("                      3.登记支出")
		fmt.Println("                      4.退出软件")
		fmt.Println("请选择功能（1-4）：")

		fmt.Scanln(&acc.key)
		switch acc.key {
		case "1":
			acc.showDetails()
		case "2":
			acc.registerIncoming()
		case "3":
			acc.registerPurchases()
		case "4":
			acc.exit()
		default:
			fmt.Println("请输入正确的选项")
		}
		if !acc.loop {
			fmt.Println("已退出记账软件...")
			break
		}
	}
}
