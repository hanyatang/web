package main

import (
	"fmt"
	"os"
)
type studay struct {
	name string
	yuwen int
	shuxue int
	yingyu int
}
var allstuday = make([]*studay,0,200)
func show() {
	fmt.Println("欢  迎  来  到  学  生  管  理  系  统")
    fmt.Println("1.添加学生数据")
    fmt.Println("2.修改学生数据")
	fmt.Println("3.显示学生数据")
	fmt.Println("4.删除学生数据")
    fmt.Println("5.退出系统")
}
func newstuday(name string,yuwen int,shuxue int,yingyu int) * studay{
	return &studay{
		name : name,
		yuwen : yuwen,
		shuxue : shuxue,
		yingyu : yingyu,
	}
}
func addstuday() {
	var (
		name string
		yuwen int
		shuxue int
		yingyu int	
	)
	fmt.Println("请根据提示输入相关内容")
    fmt.Print("请输入姓名:")
    fmt.Scanln(&name)
    fmt.Print("请输入语文成绩:")
    fmt.Scanln(&yuwen)
    fmt.Print("请输入数学成绩:")
    fmt.Scanln(&shuxue)
    fmt.Print("请输入英语成绩:")
    fmt.Scanln(&yingyu)
	fmt.Println(name,yuwen,shuxue,yingyu)
	book :=newstuday(name,yuwen,shuxue,yingyu)
	for _,b := range allstuday{
		if b.name == book.name{
			fmt.Println("该学生已存在")
			fmt.Print("\n")
			fmt.Print("\n")
			return
			
		}
	}

	allstuday = append(allstuday,book)
	fmt.Println("添加数据成功")
}
func useinput() *studay {
	var (
		name string
		yuwen int
		shuxue int
		yingyu int
	)
	fmt.Print("请输入姓名:")
    fmt.Scanln(&name)
    fmt.Print("请输入语文成绩:")
    fmt.Scanln(&yuwen)
    fmt.Print("请输入数学成绩:")
    fmt.Scanln(&shuxue)
    fmt.Print("请输入英语成绩:")
	fmt.Scanln(&yingyu)
	book :=newstuday(name,yuwen,shuxue,yingyu)
	return book
}
func updatestuday() {
	book := useinput()
	for index,b := range allstuday {
		if b.name == book.name{
			allstuday[index] = book
			fmt.Printf("%s的信息已经更新完毕",b.name)
			fmt.Print("\n") 
			fmt.Print("\n") 
			fmt.Print("\n") 
			return
		}
	}
	fmt.Printf("同学%s不存在",book.name)
	fmt.Print("\n")
	fmt.Print("\n")
	fmt.Print("\n")
}
func showstuday() {
	if len(allstuday) == 0 {
		fmt.Println("没有信息")
	}else{
		for _,b :=range allstuday{
			fmt.Printf("%s的成绩为：语文%d,数学%d,英语%d.",b.name,b.yuwen,b.shuxue,b.yingyu)
			fmt.Print("\n")
			fmt.Print("\n")
		}
		}
}
func delet() {
	
	var input int
	fmt.Println("学生信息如下,请选择删除序号几")
	for index,key := range allstuday{
		fmt.Println(index,key)
	}
		
		fmt.Scanln(&input)
		
		allstuday = append(allstuday[:input],allstuday[input+1:]...)
		fmt.Println("删除成功")
		
		}
	
		
	
		
	
		
	
	

	
	

func main() {
	
	for {
		show()
		var open int
		fmt.Scanln(&open)
	
		switch open {
		case 1:
			addstuday()
		case 2:
			updatestuday()
		case 3:
			showstuday()
		case 4:
			delet()
			fmt.Println(&allstuday)
		case 5:
			fmt.Println("欢迎下次在使用")
			os.Exit(0)
		}
	}
}