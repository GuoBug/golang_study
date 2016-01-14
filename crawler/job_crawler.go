package main

import (
    "fmt"
    "strconv"
    "time"
    "net/http"
    "io/ioutil"
    "io"
    "errors"
    "strings"
    "net/url"
    "os"
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
func CheckFileIsExist(fileName string) (bool) {
    var exist = true;
    if _, err := os.Stat(fileName); os.IsNotExist(err) {
        exist = false; 
    }
    return exist;
}

func RecordInfo(fileName string ,record string) {

//检测，有则追加，没有则新建
    fp := new(os.File)
    var err error = errors.New("")
    if CheckFileIsExist(fileName) == false {
        fmt.Println("*************False****************")
        fp, err = os.Create(fileName)
        if err != nil {
            panic(err)
        }
    }else{
        fmt.Println("*************True****************")
        fp, err = os.OpenFile(fileName, os.O_APPEND, 0666)
        if err != nil {
            panic(err)
        }
    }

    n, writeErr := io.WriteString(fp, record)
    if writeErr != nil {
        panic (writeErr)
    }
    defer fp.Close()
    fmt.Printf("本次输入 **%d** 个字节,%s", n,record)

    fp.Close()
}

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

func GetSearchResult(keyword string ,city string ,jobsCount int){
//URL参数
    paramurl := "http://www.lagou.com/custom/search.json?"
    paramKey := "&positionName=" + keyword
    paramCity := "&city=" + city
    paramPage := "pageNo="

//文件名相关
    strTime := time.Now().Format("2006-01-02")
    fileName := paramKey +  strTime + ".csv"

    pagesCount := jobsCount/15
    fmt.Println(pagesCount)

// 具体json
    jobJson := simplejson.New()
//打印验证URL
//  url = url + paramPage + paramCity + paramKey
//  fmt.Println(url)

    for i := 1 ;i <=pagesCount ; i++{
        time.Sleep(1*time.Second)
        result := MyGetFunction(paramurl + paramPage + strconv.Itoa(i) + paramCity + paramKey)
        //fmt.Printf("%s", result)
        js, err := simplejson.NewJson(result)
        if err != nil {
            panic(err.Error())
        }
        //fmt.Println(js)

//取长度，并且取出内容

        jobsCount,_ := js.Get("content").Get("data").Get("page").Get("pageSize").Int()

        for count := 0 ; count < jobsCount ; count ++ {
//            fmt.Println(js.Get("content").Get("data").Get("page").Get("result").GetIndex(i))
//获取具体工作内容的json
            jobJson = js.Get("content").Get("data").Get("page").Get("result").GetIndex(count)

            jobName ,_ := jobJson.Get("positionName").Bytes()
            jobCity ,_ := jobJson.Get("city").Bytes()
            jobCreatTime ,_ := jobJson.Get("createTime").Bytes()
            jobsalary ,_ := jobJson.Get("salary").Bytes()
            jobCompany ,_ := jobJson.Get("companyName").Bytes()

//开始写文件
            s := []string{string(jobName),string(jobCity),string(jobCreatTime),string(jobsalary),string(jobCompany)}
            RecordInfo(fileName,strings.Join(s, "\t"))
//            fmt.Printf("%s\t%s\t%s\t%s\t%s\n",jobName,jobCity,jobCreatTime,jobsalary,jobCompany)
        }

    }
}

func GetPageNumber(keyword string ,city string ) (int){
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

func main() {

    jobName := os.Args[1]
    workCity := os.Args[2]

//获取总数
    var jobNums int  
    //fmt.Printf("***********%d***************\n",len(os.Args))
    //fmt.Println(os.Args[3]) 
    if len(os.Args) == 4 {
        jobNums ,err := strconv.Atoi(os.Args[3])   
        if err != nil {
            panic(err)
        }
        fmt.Printf("您需要总共%d,条职位信息。\n",jobNums)
    }else{
        jobNums = GetPageNumber(jobName , workCity)
        fmt.Printf("尚未输入总数，将获取全部%n个职位信息\n",jobNums)
    }

    //创建文件
    strTime := time.Now().Format("2006-01-02")
    fileName := jobName +  strTime + ".csv"
    //fmt.Printf("fileName:%s", fileName)
    RecordInfo(fileName,"")

    GetSearchResult(jobName , workCity ,jobNums)
}
