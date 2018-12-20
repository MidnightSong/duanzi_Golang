package wirteFile

import (
	"fmt"
	"os"
	"strconv"
)

func JokeToFile(i int, fileTile []string, fileContent []string) {
	//创建文件
	file, e := os.Create(strconv.Itoa(i) + ".txt")
	if e != nil {
		fmt.Println("file create error", e)
	}
	defer file.Close()
	//写文件
	n := len(fileTile)
	for i = 0; i < n; i++ {
		file.WriteString("标题：" + fileTile[i] + "\n")
		file.WriteString(fileContent[i])
		file.WriteString("\n------------------------------------\n")
	}
}
