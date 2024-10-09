package main

import (
	"customermanage/model"
	"customermanage/service"
	"fmt"
)

type customerview struct {
	key  int
	loop bool

	customerservice *service.Customerservice
}

func (this *customerview) list() {
	customers := this.customerservice.List()
	fmt.Println("-----------客户列表------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].Getinfo())
	}

	fmt.Println("-----------客户列表完成-------------")
}

func (this *customerview) add() {

	fmt.Println("-----------添加客户------------")
	name := ""
	fmt.Println("姓名:")

	fmt.Scan(&name)

	fmt.Println("性别:")
	gender := ""
	fmt.Scan(&gender)
	fmt.Println("年龄:")
	age := 0
	fmt.Scan(&age)
	fmt.Println("电话号码:")
	phone := ""
	fmt.Scan(&phone)
	fmt.Println("邮箱:")
	email := ""
	fmt.Scan(&email)

	customer := model.Newcustomer2(name, age, gender, phone, email)
	// this.customerservice.Add(customer)

	if this.customerservice.Add(customer) {
		fmt.Println("------------客户添加完成-----------")
	} else {
		fmt.Println("------------客户添加失败-----------")
	}
}
func (this *customerview) delete() {
	fmt.Println("------------删除客户-------------")
	fmt.Println("    请输入待删除的客户编号(-1退出):")
	id := -1
	fmt.Scan(&id)
	if id == -1 {
		return
	}
	for {
		fmt.Println("  请确认是否要删除(y/n):")

		ch := ""
		fmt.Scan(&ch)
		if ch == "y" || ch == "Y" {
			break
		} else if ch == "n" || ch == "N" {
			return
		} else {
			fmt.Println("    输入字符不正确,请重新输入")
		}
	}
	if this.customerservice.Delete(id) {
		fmt.Println("          删除完成         ")
	} else {
		fmt.Println("        该ID号不存在,删除失败")
	}
}

func (this *customerview) update() {
	fmt.Println("     请输入你要修改的客户编号:")
	id := 0
	fmt.Scan(&id)
	if this.customerservice.Update(id) {
		fmt.Println("         修改成功！")
	} else {
		fmt.Println("       修改失败")
	}

}

func (this *customerview) mainview() {
	this.loop = true
	for {
		fmt.Println("\n------------客户信息管理软件------------")
		fmt.Println("              1.添加客户              ")
		fmt.Println("              2.修改客户              ")
		fmt.Println("              3.删除客户              ")
		fmt.Println("              4.客户列表              ")
		fmt.Println("              5.退    出              ")
		fmt.Println()
		fmt.Print("            请选择(1-5): ")
		fmt.Scan(&this.key)
		switch this.key {
		case 1:
			this.add()
		case 2:
			this.update()
		case 3:
			this.delete()
		case 4:
			this.list()
		case 5:
			for {
				fmt.Println("         请确认是否退出(y/n)")
				chh := ""
				fmt.Scan(&chh)
				if chh == "y" || chh == "Y" {
					this.loop = false
					break
				} else if chh == "n" || chh == "N" {
					this.loop = true
					break
				} else {
					fmt.Println("          请输入正确的值")
				}
			}
		default:
			fmt.Println("error")
		}
		if !this.loop {
			break
		}
	}
	fmt.Println("--------------exit--------------")
}

func main() {
	cv := customerview{}
	cv.customerservice = service.Newcustomerservice()
	cv.mainview()
}

