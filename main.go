package main

import (
	"fmt"
	. "百度贴吧爬虫/controller"
)

/**
主函数
*/
func main() {
	//爬取的起始页和终止页面
	var start, end int
	fmt.Print("输入起始页信息:")
	fmt.Scan(&start)
	fmt.Print("输入结束页面:")
	fmt.Scan(&end)
	//掉用开始的函数
	StartWorking(start, end)

}
