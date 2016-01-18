package main

import (
    "./crawlerlib"
    "strconv"
    "time"
    "strings"
    "os"
    "fmt"
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

func GetSearchResult(keyword string ,city string ,jobsCount int){

    fmt.Printf("\n*************************************************\n我开始了，W(￣_￣)W …………\n")
//URL参数
    paramurl := "http://www.lagou.com/custom/search.json?"
    paramKey := "&positionName=" + keyword
    paramCity := "&city=" + city
    paramPage := "pageNo="

//文件名相关
    strTime := time.Now().Format("2006-01-02")
    fileName := keyword +  strTime + ".csv"

    pagesCount := jobsCount/15
    fmt.Println(pagesCount)

// 具体json
    jobJson := simplejson.New()
//打印验证URL
//  url = url + paramPage + paramCity + paramKey
//  fmt.Println(url)

    for i := 1 ;i <=pagesCount ; i++{
        time.Sleep(1*time.Second)
        result := crawlerlib.MyGetFunction(paramurl + paramPage + strconv.Itoa(i) + paramCity + paramKey)
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
            jobID ,_ := jobJson.Get("positionId").Int()

//            fmt.Println(strconv.Itoa(jobID))
//            break
            
//获取job 详情
            jobURL := "http://www.lagou.com/center/job_" + strconv.Itoa(jobID) + ".html?m=1"
            fmt.Println(jobURL)
            jobDetail := crawlerlib.GetQuery(jobURL)

//开始写文件
            s := []string{string(jobName),strconv.Itoa(jobID),string(jobCity),string(jobCreatTime),string(jobsalary),string(jobCompany),string(jobDetail)}
            crawlerlib.RecordInfo(fileName,strings.Join(s, "\t"))
            crawlerlib.RecordInfo(fileName,"\n")
//            fmt.Printf("%s\t%s\t%s\t%s\t%s\n",jobName,jobCity,jobCreatTime,jobsalary,jobCompany)
        }

    }
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
        jobNums = crawlerlib.GetPageNumber(jobName , workCity)
        fmt.Printf("尚未输入总数，将获取全部%n个职位信息\n",jobNums)
    }

    //创建文件
    strTime := time.Now().Format("2006-01-02")
    fileName := jobName +  strTime + ".csv"
    //fmt.Printf("fileName:%s", fileName)
    crawlerlib.CreateFile(fileName)

    GetSearchResult(jobName , workCity ,jobNums)

}
