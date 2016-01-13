package main

import (
    "fmt"
    "strconv"
    "time"
	"net/http"
	"io/ioutil"
	"net/url"
	"github.com/bitly/go-simplejson"
)

/*
headers = [
    {'User-Agent': 'Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:34.0) Gecko/20100101 Firefox/34.0'},
    {'User-Agent': 'Mozilla/5.0 (Windows; U; Windows NT 6.1; en-US; rv:1.9.1.6) Gecko/20091201 Firefox/3.5.6'},
    {'User-Agent': 'Mozilla/5.0 (Windows NT 6.2) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.12 Safari/535.11'},
    {'User-Agent': 'Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Trident/6.0)'},
    {'User-Agent': 'Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:40.0) Gecko/20100101 Firefox/40.0'},
    {'User-Agent': 'Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/44.0.2403.89 Chrome/44.0.2403.89 Safari/537.36'}
]
*/
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

func GetSearchResult(keyword string, city string){
    paramurl := "http://www.lagou.com/custom/search.json?"
    paramKey := "&positionName=" + keyword
    paramCity := "&city=" + city
    paramPage := "pageNo="

//打印验证URL
//  url = url + paramPage + paramCity + paramKey
//  fmt.Println(url)

    for i := 1 ;i <=1 ; i++{
        time.Sleep(1*time.Second)
        result := MyGetFunction(paramurl + paramPage + strconv.Itoa(i) + paramCity + paramKey)
        //fmt.Printf("%s", result)
        js, err := simplejson.NewJson(result)
        if err != nil {
            panic(err.Error())
        }
        //fmt.Println(js)

		var nodes = make(map[string]interface{})
		nodes, _ = js.Map()
		/*
		for key,_ := range nodes {
			fmt.Println(nodes[key])
		}
		*/
		fmt.Printf("%#v\n",nodes["content"].Get("data"))

    }

}

func main() {
    GetSearchResult("python" ,"上海" )
}