package main
import(
	"fmt"
	
)

func main(){
	var b string = "加油家源"
	fmt.Println(b)

	// 字符串使用细节
	str1 := "hello " + "world!"
	str1 += "haha"
	fmt.Println(str1)

	str2 := "hello " + "world!" +"hello " + 
	"world!" +"hello " + "world!"+"hello " + 
	"world!"+"hello " + "world!"
	fmt.Println(str2)

	// 显示源代码，使用反引号

	str3 :=`package main
	import(
		"fmt"
		
	)
	
	func main(){
		var b string = "加油家源"
		fmt.Println(b)
	
		// 字符串使用细节
		str1 := "hello " + "world!"
		str1 += "haha"
		fmt.Println(str1)
	
		str2 := "hello " + "world!" +"hello " + 
		"world!" +"hello " + "world!"+"hello " + 
		"world!"+"hello " + "world!"
		fmt.Println(str2)
	
		
	}`
	fmt.Println(str3)

}