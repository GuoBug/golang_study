package crawlerlib 

import (
    "strings"
    "github.com/opesun/goquery"
)

func GetQuery(getUrl string ,) (string){

    x, err := goquery.ParseUrl(getUrl)
    if err != nil {
        panic(err)
        return "ERROR! IN Get Query."
    }
    //x.Find(".info").Print()
    detail := x.Find(".content").Html()

//格式化处理
    detail = strings.Replace(detail,"<br/>", "\n", -1)
    detail = strings.Replace(detail,"</p>", "\n", -1)
    detail = strings.Replace(detail,"<p>", "\n", -1)

    return detail
}