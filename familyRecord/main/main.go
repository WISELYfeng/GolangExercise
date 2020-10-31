package main

import (
	"go_code/chapter12/familyRecord/utils"
)

// 编写家庭收支软件，实现记录收入与支出、退出的功能
func main() {
	utils.NewAccount().MainMenu()
}
