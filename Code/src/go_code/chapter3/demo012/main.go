package main
import(
	"fmt"
	"unsafe"
)

func main(){

	var n1 int64 = 100
	fmt.Printf("n1的类型是： %T\n", n1)

	var n2 int8 = int8(n1)
	fmt.Printf("n2的类型是： %T \n", n2)

	var n3 float32 = 321.5554444444
	fmt.Printf("n3的类型是： %T \n", n3)
	fmt.Printf("n3的值是： %f \n", n3)

	var n4 float64 = float64(n3)
	fmt.Printf("n4的类型是： %T \n", n4)
	fmt.Printf("n4的值是： %f \n", n4)

	var n5 int = 007

// 这里发生了个小问题？
// int = 607与int = 0607返回值竟然不一样
// 这里报错 malformed malformed malformed malformed malformed malformed interge octal constant :畸形的整型八进制变量，哈哈，被发现了，
// 我去看看python报不报错，同样有的，invaild token

	fmt.Printf("n5的类型是： %T \n", n5)
	fmt.Printf("n5的值是： %d \n", n5)
	fmt.Printf("n5占用的字节数（注意注意不是比特数）：是： %d \n", unsafe.Sizeof(n5))

	var n6 float64 = float64(n5)
	fmt.Printf("n6的类型是： %T \n", n6)
	fmt.Printf("n6的值是： %f \n", n6)

	
	var n7 float64 = 454.6666
	fmt.Printf("n7的类型是： %T \n", n7)
	fmt.Printf("n7的值是： %f \n", n7)

	
	var n8 int64 = int64(n7)
	fmt.Printf("n8的类型是： %T \n", n8)
	fmt.Printf("n8的值是： %d \n", n8)

	fmt.Printf("n7的类型是%T",n7)






}