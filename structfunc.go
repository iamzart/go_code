package main

import "fmt"

/*图书管理系统*/

type Book struct {
	Title  string
	Author string
	Price  float64
}

func (b Book) Display() {
	fmt.Println(b.Title, b.Author, b.Price)
}

type Library struct {
	Booklist []Book //Booklist是一个切片，里边存放的是结构体
}

func (l *Library) Addbook(b Book) {
	l.Booklist = append(l.Booklist, b)
} //绑定的结构体是图书馆，这个方法的参数是结构体（图书）

func (l Library) Showbook() {
	if len(l.Booklist) == 0 {
		fmt.Println("图书馆目前没有书籍")
		return
	}
	for _, book := range l.Booklist { //book相当于l.Booklist里面的每一个元素，就是结构体嘛
		book.Display()
	}
}

//查找
//返回结构体类型指针
func (l Library) Findbook(title string) *Book {
	//先遍历
	for _, book := range l.Booklist {
		if title == book.Title {
			fmt.Println("找到书籍： ")
			book.Display()
			return &book //找到就退出
		}
	}
	fmt.Println("没找到")
	return nil
}

//删除
func (l *Library) Removebook(title string) {
	//先遍历
	for i, book := range l.Booklist {
		if title == book.Title {
			fmt.Println("该书籍是： ")
			book.Display()
			//删除 从i开始删除一个
			l.Booklist = append(l.Booklist[:i], l.Booklist[i+1:]...)
			return
		}
	}
	fmt.Println("没找到")
}
func main() {
	//var b Book
	/*b = Book{
		Title:  "我与地坛",
		Author: "史铁生",
		Price:  59.0,
	}
	c := Book{
		Title:  "小王子",
		Author: "圣埃克苏佩里",
		Price:  39.0,
	}
	l.Addbook(b)
	l.Addbook(c)
	l.Showbook()
	var bookname string
	fmt.Println("请输入想要查找的书名： ")
	fmt.Scanln(&bookname)
	l.Findbook(bookname)*/
	var l Library
	var title, author string
	var price float64
	//终于明白为什么要把书名分开写，因为把他们都分开可以被提示的输入，更现实，更有模有样，而且书名为0直接退出

	fmt.Println("******************************************")
	fmt.Println("************欢迎使用图书管理系统************")
	fmt.Println("******************************************")
	fmt.Println("*******1.添加书籍********2.查找书籍********")
	fmt.Println("******3.删除书籍*******4.显示所有书籍******")
	fmt.Println("***************5.退出系统****************")
	fmt.Println("*****************************************")
	var choice int
	for {
		fmt.Println("请输入你的选择： ")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			fmt.Println("添加书籍")
			for {
				fmt.Printf("输入书籍信息,书名为0停止查找\n")
				fmt.Println("书名：")
				fmt.Scanln(&title)
				if title == "0" {
					break
				}
				fmt.Println("作者： ")
				fmt.Scanln(&author)
				fmt.Println("价格： ")
				fmt.Scanln(&price)
				fmt.Println("已录入")
				l.Addbook(Book{Title: title, Author: author, Price: price})

			}
		case 2:
			fmt.Println("查找书籍")
			//查找
			var findname string
			for {

				fmt.Println("请输入想要查找的书名： ,停止查找时输入0")
				fmt.Scanln(&findname)
				if findname == "0" {
					break
				}
				l.Findbook(findname)
			}
		case 3:
			fmt.Println("删除书籍")
			//删除
			var delname string
			for {
				fmt.Println("请输入要删除的书名： ,停止删除时请输入0")
				fmt.Scanln(&delname)
				if delname == "0" {
					break
				}
				l.Removebook(delname)
			}
			fmt.Println("删除后： ")
			l.Showbook()
		case 4:
			fmt.Println("显示所有书籍")
			//展示
			//用l的方法展示的是结构体里包含的切片的所有元素
			l.Showbook()
		case 5:
			fmt.Println("退出系统")
			return
		default:
			fmt.Println("输入错误,请重新选择")

		}
	}
}
