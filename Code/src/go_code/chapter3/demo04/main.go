// // 多个变量的使用规则，有三种情形
// package main
// import "fmt"

// var(
// 	v1 int
// 	v2 string
// )


// func main(){
// 	var i int
// 	i = 10
// 	i = 50

// 	i1 := 60

// 	var i2 float32 

// 	fmt.Println("i2 = ",i2)

// 	fmt.Println("i1 = ", i1, "i = ",i)

// 	v1 = 100
// 	// v2 = "tom"
// 	fmt.Println("v1 = ", 100, "v2 = ", v2)
// }


package main
import "fmt"
func main(){
	var i int = 2131321
	var j float32 = 2321.321321
	var k string = "tom_nb"

	fmt.Println("i = ",i,"\n","j = ",j,"\n","k = ",k)
}