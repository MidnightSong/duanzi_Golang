package mainPage

import (
	"duanzi/jokePage"
	"duanzi/wirteFile"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
)

//工作函数
func DoWork(start, end int) {
	fmt.Printf("准备爬去%d页到%d页的网址\n", start, end)
	page := make(chan int)
	for i := start; i <= end; i++ {
		//定义一个函数，爬主页面
		go MainPage(i, page)
	}

	for i := start; i <= end; i++ {
		fmt.Printf("第%d页爬取完成\n", <-page)
	}
}

//爬取主页面
func MainPage(i int, page chan int) {
	url := "https://www.pengfu.com/xiaohua_" + strconv.Itoa(i) + ".html"
	fmt.Printf("正在爬取第%d个网页：%s\n", i, url)

	//开始爬取页面内容
	result, err := MainPageGet(url)
	if err != nil {
		fmt.Println("MainPageGet error:", err)
		return
	}

	/*提取内容，去掉无用的数据

	<h1 class="dp-b"><a href=" 一个段子的url "
	通过正则表达式获取关键信息       ?s: 代表处理\n
	*/
	compile := regexp.MustCompile(`<h1 class="dp-b"><a href="(?s:(.*?))"`)
	if compile == nil {
		fmt.Println("正则表达式compile出错")
		return
	}
	jokesURL := compile.FindAllStringSubmatch(result, -1) //-1代表查找所有内容,返回字符串切片
	/* jokesURL :=
	[[<h1 class="dp-b"><a href="https://www.pengfu.com/content_1854668_1.html" https://www.pengfu.com/content_1854668_1.html]
		[<h1............................
	二维切片
	*/
	fileTitle := make([]string, 0)
	fileContent := make([]string, 0)
	//提取网址
	for _, data := range jokesURL {
		//开始爬取每一个笑话、段子
		title, content, err := jokePage.GetJoke(data[1])
		if err != nil {
			fmt.Println("GetJoke error :", err)
		}
		fileTitle = append(fileTitle, title)
		fileContent = append(fileContent, content)
	}

	//写入文件
	wirteFile.JokeToFile(i, fileTitle, fileContent)
	//同步
	page <- i

}

//获取主页面内容
func MainPageGet(url string) (result string, err error) {
	resp, err1 := http.Get(url) //发送get请求
	if err != nil {
		err = err1
		return
	}
	defer resp.Body.Close()

	buf := make([]byte, 4*1024)
	for {
		n, _ := resp.Body.Read(buf)
		if n == 0 { //n==0说明读取完毕
			break
		}
		//把结果存到result,[;n]重要，代表每次读多少存多少
		result += string(buf[:n])
	}
	return
}
