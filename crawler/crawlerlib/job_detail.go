package crawlerlib 

import (
    "regexp"
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
//移除HTML
    r := regexp.MustCompile("<.*?>")
    detail = r.ReplaceAllString(detail, " ")
//移除换行
    r = regexp.MustCompile("[ \f\n\r\t\v]")
    detail = r.ReplaceAllString(detail, "")

    return detail
}