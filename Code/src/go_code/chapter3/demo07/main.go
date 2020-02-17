package main
import(
	"fmt"
	"unsafe"
)
func main(){
	var b = false
	fmt.Println("b = ", b)
	fmt.Printf("一个布尔值所占用的字节数： %d", unsafe.Sizeof(b))
}