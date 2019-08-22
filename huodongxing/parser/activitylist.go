package parser

import (
	"fmt"
	"log"
	"regexp"
	"reptile/engine"
	"strconv"
)

const (
	activityListRe = `<a href="(/event/[0-9]+)`
	urlPrefix      = "https://bj.huodongxing.com"
	urlListPrefix  = "https://bj.huodongxing.com/events?orderby=o&status=1&tag=%E8%81%9A%E4%BC%9A%E4%BA%A4%E5%8F%8B&city=%E5%8C%97%E4%BA%AC&page="
)

var (
	//pageNextRe = regexp.MustCompile(`<a href="javascript:;" class="layui-laypage-next"\s+data-page="([0-9]+)">下一页</a>`)
	pageNextRe = regexp.MustCompile(`elem: 'pagination',\s+count:\s([0-9]+),\s+limit:\s([0-9]+),\s+curr:\s([0-9]+),`)
)

func ParseActivityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(activityListRe)
	matches := re.FindAllSubmatch(contents, -1)
	//log.Printf("activity list len is %d", len(matches))
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, string(m[1]))
		//log.Printf("activity url is %s", string(m[0]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        urlPrefix + string(m[1]),
			ParserFunc: ParseActivity,
		})
	}
	//log.Printf("activity list is %s", string(contents))
	matches = pageNextRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		//log.Println("-----",string(m[0]))
		count,err:=strconv.Atoi(string(m[1]))
		if err!=nil {
			log.Printf("page next add err : %v",err)
		}
		limit,err:=strconv.Atoi(string(m[2]))
		if err!=nil {
			log.Printf("page next add err : %v",err)
		}
		curr,err:=strconv.Atoi(string(m[3]))
		if err!=nil {
			log.Printf("page next add err : %v",err)
		}
		if count/limit>curr {
			result.Requests = append(result.Requests, engine.Request{
				Url:         urlListPrefix+fmt.Sprint(curr+1),
				ParserFunc: ParseActivityList,
			})
		}
	}
	return result
}
