package crawlerlib 

import (
    "net/url"
    "net/http"
    "io/ioutil"
    "fmt"
    "strconv"
    "github.com/bitly/go-simplejson"
)


func MyGetFunction(getUrl string ,)([]byte){
    u, _ := url.Parse(getUrl)
    q := u.Query()
    u.RawQuery = q.Encode()
    res, err := http.Get(u.String());
    if err != nil { 
        return nil
    }
    result, err := ioutil.ReadAll(res.Body) 
    res.Body.Close() 
    if err != nil { 
        return nil
    } 

    return result
}


func GetPageNumber(keyword string ,city string ) (int){
    fmt.Printf("\n***************\n我想知道一共有多少\n")
    paramurl := "http://www.lagou.com/custom/search.json?"
    paramKey := "&positionName=" + keyword
    paramCity := "&city=" + city
    paramPage := "pageNo="

    result := MyGetFunction(paramurl + paramPage + strconv.Itoa(1) + paramCity + paramKey)
    
    js, err := simplejson.NewJson(result)
    if err != nil {
        panic(err.Error())
    }
//取长度，并且取出内容
    pageCount,_ := js.Get("content").Get("data").Get("page").Get("totalCount").Int()

    return pageCount

}