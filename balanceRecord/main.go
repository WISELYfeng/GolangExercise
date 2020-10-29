package main

import (
	"fmt"
)

// 编写家庭收支软件，实现记录收入与支出、退出的功能
func main() {
	// 接收用户输入的功能选项
	key := ""

	// 控制循环是否继续
	loop := true
	// 初始余额
	balance := 10000.0
	// 每次收支金额
	money := 0.0
	// 收支说明
	note := ""
	// 记录是否有收支
	flag := false
	// 收支详情记录
	details := "收支\t\t账户金额\t\t收支金额\t\t说明"
	// 显示主菜单
	for {
		fmt.Println("\n---------------------家庭收支软件-------------------")
		fmt.Println("                      1.收支明细")
		fmt.Println("                      2.登记收入")
		fmt.Println("                      3.登记支出")
		fmt.Println("                      4.退出软件")
		fmt.Println("请选择功能（1-4）：")

		fmt.Scanln(&key)
		switch key {
		case "1":
			if !flag {
				fmt.Println("当前没有收支明细，请先添加")
			} else {
				fmt.Println("------------------当前收支明细记录-----------------")
				fmt.Println(details)
			}
		case "2":
			fmt.Println("本次收入金额：")
			fmt.Scanln(&money)
			balance += money
			fmt.Println("本次收入说明：")
			fmt.Scanln(&note)
			details += fmt.Sprintf("\n收入\t\t%v\t\t%v\t\t%v", balance, money, note)
			flag = true
		case "3":
			fmt.Println("本次支出金额：")
			fmt.Scanln(&money)
			// 支出需要检查余额
			if money > balance {
				fmt.Println("余额不足，操作失败....")
				break
			}
			balance -= money
			fmt.Println("本次支出说明：")
			fmt.Scanln(&note)
			details += fmt.Sprintf("\n支出\t\t%v\t\t%v\t\t%v", balance, money, note)
			flag = true
		case "4":
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
				loop = false
			}
		default:
			fmt.Println("请输入正确的选项")
		}

		if !loop {
			break
		}
	}
	fmt.Println("已退出记账软件...")
}
