package parser

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"regexp"
	"reptile/engine"
	"strconv"
	"strings"
)

var (
	timePeopleNoRe = regexp.MustCompile(`([0-9]+年[0-9]+月[0-9]+日 [0-9:]+).+[^\S]([0-9]+年[0-9]+月[0-9]+日 [0-9:]+)\s+(\S+)`)
	positionRe     = regexp.MustCompile(`var position = "([0-9.]+),([0-9.]+)"`)
)

func ParseActivity(contents []byte) engine.ParseResult {
	result := engine.ParseResult{}
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(contents))
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("#container-lg").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title s.
		content, _ := s.Find("#event_desc_page").Html() //活动内容
		title := s.Find(".detail-pp").Text()            //title
		avatar, _ := s.Find("img").Attr("src")          //title
		timePeopleNo := s.Find(".address-info-wrap").Eq(0).Text()
		tpMatches := timePeopleNoRe.FindAllStringSubmatch(timePeopleNo, -1)
		var beginTime string
		var endTime string
		var peopleNo string
		if len(tpMatches)==1&& len(tpMatches[0])==4 {
			beginTime = tpMatches[0][1]
			endTime = tpMatches[0][2]
			peopleNo = tpMatches[0][3]
		}
		address := s.Find(".link-a-hover").Text()
		location := s.Find(".aside script").Text()
		lMatches := positionRe.FindAllStringSubmatch(location, -1)
		var longitude float64
		var latitude float64
		if len(lMatches) == 1 && len(lMatches[0]) == 3 {
			longitude, err = strconv.ParseFloat(lMatches[0][1], 64)
			latitude, err = strconv.ParseFloat(lMatches[0][2], 64)
		}
		content = strings.ReplaceAll(content, `src="/file/ue/`, `src="http://wimg.huodongxing.com/file/ue/`)

		fmt.Printf("content : %s\n",  content)
		fmt.Printf("title : %s\n",  title)
		fmt.Printf("avatar : %s\n",  avatar)
		fmt.Printf("address : %s\n",  address)
		fmt.Printf("longitude : %f\n",  longitude)
		fmt.Printf("latitude : %f\n",  latitude)
		fmt.Printf("beginTime : %s\n",  beginTime)
		fmt.Printf("endTime : %s\n",  endTime)
		fmt.Printf("peopleNo : %s\n", peopleNo)

	})
	return result
}
