// 多个变量的使用规则，有三种情形
package main
import "fmt"

// 全局变量的声明与赋值
// 方式一
var v1 = 100
var v2 = "tom ke lu si "
var v3 = 52.10
// 方式二
var(
	v4 = 1000
	v5 = "jack love"
	v6 = 741.41
)


func main(){
	// 情形一，
	var n1, n2, n3 int
	fmt.Println("n1 =",n1,"n2 = ", n2, "n3 = ", n3)

	// // 情形二,
	var n21, n22, n23 = 100, "tom", 888.11
	fmt.Println("n21 =",n21,"n22 = ", n22, "n23 = ", n23)
	// // 情形三,
	n11, n12, n13 := 100, "tom", 888.11
	fmt.Println("n11 =",n11,"n12 = ", n12, "n13 = ", n13)
	
	fmt.Println("v1 =",v1,"v2 = ", v2, "v3 = ", v3)
	
	fmt.Println("v4 =",v4,"v5 = ", v5, "v6 = ", v6)
}