package main
import (
	"fmt"
)

func main(){
	var n1 int32 = 100

	var n2 float32 = float32(n1)

	var n3 = "zhe shi go"

	fmt.Println("n1 = ", n1, "n2 = ", n2)
	fmt.Printf("n3 = %c \n", n3)
	fmt.Printf("n1的类型是：%T ,n2的类型是：%T\n", n1, n2)
	fmt.Printf("n1的值是：%v ,n2的值是：%v", n1, n2)
	
}