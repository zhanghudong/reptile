package main

import (
	"reptile/engine"
	"reptile/huodongxing/parser"
)




func main() {
	url := "https://bj.huodongxing.com/events?orderby=o&status=1&tag=%E8%81%9A%E4%BC%9A%E4%BA%A4%E5%8F%8B&city=%E5%8C%97%E4%BA%AC"
	//url := "https://bj.huodongxing.com//event/2506703473311"
	engine.Run(
		engine.Request{
			Url:url,
			ParserFunc:parser.ParseActivityList,
		},
	)

}

