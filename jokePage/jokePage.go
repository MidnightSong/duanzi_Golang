package jokePage

import (
	"errors"
	"net/http"
	"regexp"
	"strings"
)

func GetJoke(url string) (title, content string, err error) {
	result, e := JokePageGet(url)
	//开始爬取页面内容
	if err != nil {
		err = e
		return
	}
	//取关键信息
	//取段子标题		<h1> 标题  </h1>   只取一个
	compile1 := regexp.MustCompile(`<h1>(?s:(.*?))</h1>`)
	if compile1 == nil {
		err = errors.New("GetJoke regexp title compile error")
		return
	}
	tmpTitle := compile1.FindAllStringSubmatch(result, 1) //1，代表只过滤第1个
	for _, data := range tmpTitle {
		title = data[1]
		title = strings.Replace(title, "\t", "", -1) //去掉多余字符
		break
	}

	//取段子内容  <div class="content-txt pt10">  内容  <a id="prev"
	compile2 := regexp.MustCompile(`<div class="content-txt pt10">(?s:(.*?))<a id="prev"`)
	if compile2 == nil {
		err = errors.New("GetJoke regexp content compile error")
		return
	}
	tmpContent := compile2.FindAllStringSubmatch(result, -1) //-1代表过滤所有内容
	for _, data := range tmpContent {
		content = data[1]
		content = strings.Replace(content, "\t", "", -1)     //去掉多余字符
		content = strings.Replace(content, "<br />", "", -1) //去掉多余字符
		content = strings.Replace(content, "&nbsp;", "", -1) //去掉多余字符
		break
	}
	return
}

/***
获取页面内容
parm 获取页面的URL
return 网页内容<body>
*/
func JokePageGet(url string) (result string, err error) {
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
