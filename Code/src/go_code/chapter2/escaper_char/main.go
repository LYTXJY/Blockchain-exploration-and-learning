package main

import "fmt" //fmt包中提供格式化，输出，输入得函数

func main(){
	// 演示转义字符的使用
	// 1.\t	:制表符1
	fmt.Println("1.:制表符1")
	fmt.Println("韩梅梅\t马小跳")//

	// 2.\n	:换行符
	fmt.Println("2.:换行符")
	fmt.Println("王一博\n肖战")
	
	// 3.\\	:表示一个\,多用来表示地址时使用

	fmt.Println("3.:表示一个\\")
	fmt.Println("D:\\za_xiang\\bilibili_video\\BilibiliDownload\\75631089")

	// 4.\"	:表示一个",多用来想在输出中见到“”号，防止与Println("")中的引号冲突
	fmt.Println("4.:表示一个\"")
	fmt.Println("\"为中华崛起而读书\"，周恩来如是说")
	// 5.\r	:回车符号，回车不是换行，回车字符后面的内容输出到句子首部

	fmt.Println("5.:表示回车字符")
	fmt.Println("本是男儿身，不是女娇娥\r程蝶衣我啊")

}