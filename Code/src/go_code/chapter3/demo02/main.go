// 变量的使用规则，有三种情形
package main
import "fmt"
func main(){
	// 情形一，定义变量后，仅声明了变量类型，没有进行下一步的赋值操作
	// 则,变量的值为默认值(不同类型的变量,默认值不一样)
	var a int
	fmt.Println("a = ", a)
	var b int
	fmt.Println("b = ", b)
	// 情形二,定义变量的值,但不声明变量是什么类型,
	// 交给编译器去弄明白这是什么变量,这叫变量的类型推导(python java支持)
	var c = 10.1
	fmt.Println("c = ", c)
	// 情形三,简写变量的定义与赋值
	// var name string
	// name = "tom"
	name := "tom"
	fmt.Println("name = ", name)
}