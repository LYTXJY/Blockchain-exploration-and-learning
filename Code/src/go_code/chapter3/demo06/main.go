package main
// import "fmt"
// import "unsafe"
import(
	"fmt"
	"unsafe"
)


func main(){

	var i int8 = -128
	fmt.Println("i = ", i)
	var j int8 = 127
	fmt.Println("j = ", j)

	var i1 int16 = -129
	fmt.Println("i1 = ", i1)
	var j1 int16 = 128
	fmt.Println("j1 = ", j1)

	var i2 uint8 = 0
	fmt.Println("i2 = ", i2)
	var j2 uint8 = 255
	fmt.Println("j2 = ", j2)

	var i3 int = -123123
	fmt.Println("i3 =", i3)

	var j3 byte = 244
	fmt.Println("j3 =", j3)

	// 整型使用细节

	var n1 = 100
	fmt.Printf("n1 的类型是 %T", n1)
	
	fmt.Println("\n")

	var n2 int64 = -321321321
	fmt.Printf("n2 的类型是 %T ,  n2 占用的字节数是：%d",n2, unsafe.Sizeof(n2))




}