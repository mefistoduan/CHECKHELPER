package main

import (
	"fmt"
	_ "fmt"
	_ "github.com/go-vgo/robotgo"
	"os"
	"strings"
)

func main() {
	print(" $$$$$$\\  $$\\   $$\\ $$$$$$$$\\  $$$$$$\\  $$\\   $$\\ $$\\   $$\\ $$$$$$$$\\ $$\\       $$$$$$$\\  $$$$$$$$\\ $$$$$$$\\  \n")
	print("$$  __$$\\ $$ |  $$ |$$  _____|$$  __$$\\ $$ | $$  |$$ |  $$ |$$  _____|$$ |      $$  __$$\\ $$  _____|$$  __$$\\ \n")
	print("$$ /  \\__|$$ |  $$ |$$ |      $$ /  \\__|$$ |$$  / $$ |  $$ |$$ |      $$ |      $$ |  $$ |$$ |      $$ |  $$ |\n")
	print("$$ |      $$$$$$$$ |$$$$$\\    $$ |      $$$$$  /  $$$$$$$$ |$$$$$\\    $$ |      $$$$$$$  |$$$$$\\    $$$$$$$  |\n")
	print("$$ |      $$  __$$ |$$  __|   $$ |      $$  $$<   $$  __$$ |$$  __|   $$ |      $$  ____/ $$  __|   $$  __$$< \n")
	print("$$ |  $$\\ $$ |  $$ |$$ |      $$ |  $$\\ $$ |\\$$\\  $$ |  $$ |$$ |      $$ |      $$ |      $$ |      $$ |  $$ |\n")
	print("\\$$$$$$  |$$ |  $$ |$$$$$$$$\\ \\$$$$$$  |$$ | \\$$\\ $$ |  $$ |$$$$$$$$\\ $$$$$$$$\\ $$ |      $$$$$$$$\\ $$ |  $$ |\n")
	print(" \\______/ \\__|  \\__|\\________| \\______/ \\__|  \\__|\\__|  \\__|\\________|\\________|\\__|      \\________|\\__|  \\__|\n")
	println("自动处理考勤文件系统(CHECKHELPER) V1.0 2021-1-9")
	println("auther:Mefisto 838560574@qq.com")
	println("点击数字后开始转换，1是保留前1个月，依此类推")

	//one := robotgo.AddEvents("1")
	//if one {
	//	fmt.Println("转换已开始...")
	filetext := GetKeyWordsList()
	//	//遍历并删除指定内容
	forAndDel(filetext, 1)
	//}
}

//遍历并删除指定内容
func forAndDel(text string, month int) {
	//目标内容
	goal := "2021-01"

	if strings.ContainsAny(text, goal) {
		num := strings.Index(text, goal) - len(goal)
		content := text[num : len(text)-1]
		println(content)
		//createFile(content)
	} else {
		println("没有找到目标内容")
	}
}

//获取关键字数组
func GetKeyWordsList() string {
	f, err := os.OpenFile("1_attlog.dat", os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("文件打开失败: ", err)
		return ""
	}
	defer f.Close()
	/**
	  切片在使用前要申请内存空间,如果不知道文件的大小，尽量多给点空间最好是4K的倍数
	  如果在读取大文件的情况下，我们应该循环读取，当然，我的笔记里有读取大文件的实战案例。
	*/
	temp := make([]byte, 1024*2000)

	_, _ = f.Read(temp)

	return string(temp)
}

//生成内容
func createFile(text string) {
	fileName := "1_attlog.txt"
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("文件已转换成功,名字是" + fileName)
		_, _ = f.Write([]byte(text))
	}
}
