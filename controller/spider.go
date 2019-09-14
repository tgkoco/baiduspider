package controller

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

/**
爬取数据的控制信息
*/

const Url = "https://tieba.baidu.com/f?kw=golang&ie=utf-8&pn="

//这个是开始爬取,输入开始页面和结束页面
func StartWorking(start, end int) {
	//定义一个管道,晒先主子的同步
	page := make(chan int)
	fmt.Printf("正在爬取%d页到第%d页...\n", start, end)
	path, err := os.Getwd()
	if err != nil {
		fmt.Println("读取文件的路径错误", err)
	}
	for i := start; i <= end; i++ {
		//使用go调用单页爬取的函数
		go spiderPageOne(i, path, page)
	}
	//这时就会快速执行主构成 子构成没有执行
	//主子构成的同步
	for i := start; i <= end; i++ {
		fmt.Println("爬取完成第%d页面", <-page)
	}

}

//根据url进行爬取信息(使用的是get)
func httpgetByUrl(url string) (result string, error error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		error = err1 //返回错误信息
	}
	defer resp.Body.Close() //关闭数据请求
	//循环调用 网页数据 传给调用者
	bufr := make([]byte, 4096) //创加一个接受信息的数组
	for {
		n, err2 := resp.Body.Read(bufr)
		if n == 0 {
			fmt.Println("读取完成", url)
			break
		}
		if err2 != nil && err2 != io.EOF { //判断是否读取结束
			error = err2
			return
		}
		//累加读取的信息内容
		result += string(bufr[:n])
	}
	return
}

//爬取单个网页的函数第一参数是第几页,第二个是文件保存的路径
func spiderPageOne(i int, path string, page chan int) {
	//拼接url地址
	url := Url + strconv.Itoa((i-1)*50)
	//调用函数获取网页的信息
	result, err := httpgetByUrl(url)
	//打印错误继续爬取内容
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println("网页的数据信息",result)
	//网页数据的保存---->文件类型
	//保存的名称
	file, err := os.Create(path + "/data/golang百度贴吧" + "第" + strconv.Itoa(i) + "页" + ".html")
	if err != nil {
		fmt.Println("创建文件发生错误信息:", err)
	}
	//向文件中写入数据
	file.WriteString(result)
	//关闭文件不能等到执行完毕执行,因为每次循环都有一个文件创建
	file.Close()
	//page添加数据
	page <- i
}
