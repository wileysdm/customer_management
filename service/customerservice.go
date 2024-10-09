package service

import (
	"customermanage/model"
	"fmt"
)

type Customerservice struct {
	customers   []model.Customer
	customerNum int
}

func Newcustomerservice() *Customerservice {
	customerservice := &Customerservice{}
	customerservice.customerNum = 1
	customer := model.Newcustomer(1, "wes", 12, "male", "1232134", "4576fdd45ef")
	customerservice.customers = append(customerservice.customers, customer)
	return customerservice
}

func (this *Customerservice) List() []model.Customer {
	return this.customers
}

func (this *Customerservice) Add(customer model.Customer) bool {
	this.customerNum++
	customer.Id = this.customerNum
	this.customers = append(this.customers, customer)
	return true
}

func (this *Customerservice) Findid(id int) int {
	index := -1
	for i := 0; i < len(this.customers); i++ {
		if this.customers[i].Id == id {
			index = i

		}
	}
	return index

}

func (this *Customerservice) Delete(id int) bool {
	index := this.Findid(id)
	if index == -1 {
		return false
	}
	this.customers = append(this.customers[:index], this.customers[index+1:]...)
	return true
}

func (this *Customerservice) Update(id int) bool {
	index := this.Findid(id)
	if index == -1 {
		return false
	}
	fmt.Printf("姓名(%v):", this.customers[index].Name)
	name := ""
	fmt.Scan(&name)
	this.customers[index].Name = name
	fmt.Println()
	fmt.Printf("性别(%v):", this.customers[index].Gender)
	gender := ""
	fmt.Scan(&gender)
	this.customers[index].Gender = gender
	fmt.Println()

	fmt.Printf("年龄(%v):", this.customers[index].Age)
	age := 0
	fmt.Scan(&age)
	this.customers[index].Age = age
	fmt.Println()

	fmt.Printf("电话(%v):", this.customers[index].Phone)
	phone := ""
	fmt.Scan(&phone)
	this.customers[index].Phone = phone
	fmt.Println()

	fmt.Printf("邮箱(%v):", this.customers[index].Email)
	email := ""
	fmt.Scan(&email)
	this.customers[index].Email = email
	return true

}
