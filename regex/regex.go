package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
)


var (
	emailRe= regexp.MustCompile(`/event/[0-9]+`)
	pageNextRe =regexp.MustCompile(`<a href="javascript:;" class="layui-laypage-next"\s+data-page="([0-9]+)">下一页</a>`)
)

func main() {
	bytes, err := ioutil.ReadFile("regex/hdx.html")
	if err!=nil{
		log.Printf("err:%v",err)
		return
	}
	match := pageNextRe.FindAllSubmatch(bytes,-1)
    for _,m:=range match{
    	fmt.Println(string(m[1]))
	}

}
