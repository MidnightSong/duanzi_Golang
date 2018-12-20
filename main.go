package main

import (
	"duanzi/mainPage"
	"fmt"
)

/**主页
https://www.pengfu.com/xiaohua_1.html
https://www.pengfu.com/xiaohua_2.html

**
<h1 class="dp-b"><a href="  一个段子的url  "
*/

func main() {
	var start, end int
	fmt.Printf("输入起始页（>=1）:")
	fmt.Scan(&start)
	fmt.Printf("输入终止页（>=起始页）：")
	fmt.Scan(&end)

	mainPage.DoWork(start, end)
}
