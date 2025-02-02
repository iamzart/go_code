package main

import "fmt"

type Product struct {
	num   int
	name  string
	price float64
}

func (p Product) Proshow() {
	fmt.Printf("商品数量： %v\n商品名称为： %v\n商品价格为：%v\n", p.num, p.name, p.price)
	fmt.Println("***************************************")
}

type cart struct {
	products []Product //结构体切片，后边在购物车里边加东西
}

type Shoppingcart interface {
	Addproduct(p Product)
	Totalammount() float64
}

func (c *cart) Addproduct(p Product) {
	c.products = append(c.products, p) //给切片里边加p，加结构体类型的商品
}

func (c *cart) Totalammount() float64 {
	var total float64
	//因为所有商品都在切片里，所以遍历切片
	for _, amount := range c.products {
		total += amount.price * float64(amount.num)
	}
	return total
}

func main() {

	fmt.Println("欢迎来到购物车！")
	fmt.Println("*********************************")
	fmt.Println("***********1.添加商品*************")
	fmt.Println("***********2.查看购物车***********")
	fmt.Println("***********3.结算购物车***********")
	fmt.Println("***********4.退出购物车***********")
	fmt.Println("*********************************")
	mycart := &cart{}
	var num int
	var name string
	var price float64
	var op int
	fmt.Println("请输入你的操作：")
	for {
		fmt.Scanln(&op)
		switch op {
		case 1:
			for {
				fmt.Print("请输入商品数量：")
				fmt.Scanln(&num)
				// 如果数量为 0，则退出循环
				if num == 0 {
					break
				}
				fmt.Print("请输入商品名称：")
				fmt.Scanln(&name)

				fmt.Print("请输入商品价格：")
				fmt.Scanln(&price)
				p := Product{num, name, price}
				mycart.Addproduct(p) //调用购物车结构体的添加商品方法，把商品添加到购物车里边
				fmt.Println("商品添加成功！")
			}
		case 2:
			fmt.Println("查看购物车：")
			for _, v := range mycart.products {
				v.Proshow() //这个切片里都是结构体，然而show方法绑定的就是结构体
			}

		case 3:
			fmt.Println("结算购物车：")
			total := mycart.Totalammount()
			fmt.Printf("总价为：%v\n", total)
		case 4:
			fmt.Println("退出购物车：")
			return
		default:
			fmt.Println("输入错误，请重新输入：")
		}
	}
}
