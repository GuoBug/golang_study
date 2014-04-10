package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

//定义新的数据类型
type Spider struct {
	url    string
	header map[string]string
}

//定义 Spider的方法
func (keyword Spider) get_html_header() string {

	client := &http.Client{}

	req, err := http.NewRequest("GET", keyword.url, nil)
	if err != nil {
	}

	for key, value := range keyword.header {

		req.Header.Add(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	}
	return string(body)

}

func FileW(keyword string, header map[string]string) {

	url := "http://suggestion.baidu.com/su?json=1&wd=" + keyword

	spider := &Spider{url, header}

	html := spider.get_html_header()
	fmt.Println(html)
	rp1 := regexp.MustCompile("\\[.*?\\]") //找出【】中的内容
	find_txt := rp1.FindAllString(html, -1)
	//fmt.Println(find_txt)
	no_head := strings.Replace(find_txt[0], "[\"", "", 1)
	pur_word := strings.Replace(no_head, "\"]", "", 1)
	end := strings.Split(pur_word, "\",\"")

	f, err := os.OpenFile("reci.xls", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rp2 := regexp.MustCompile(":\".*?\"")
	fmt.Println(rp2)
	find_txt = rp2.FindAllString(html, -1)
	pur_word = strings.Replace(find_txt[0], "\"", "", 2)
	pur_word = strings.Replace(pur_word, ":", "", 1)

	f.WriteString(pur_word)
	for _, v := range end {
		f.WriteString("\t" + v)
	}
	f.WriteString("\n")

	f.Close()
}

func main() {
	header := map[string]string{"Host": "suggestion.baidu.com",
		"Referer":    "http://www.baidu.com/",
		"DNT":        "1",
		"User-Agent": "Mozilla/5.0 (Linux; Android 4.2.1; en-us; Nexus 4 Build/JOP40D) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166 Mobile Safari/535.19",
		"Cookie":     "BAIDUID=DBFCD2FF9684376495DF4D6A5B506EF0:FG=1; BDUSS=kZIYU9Gczl5ZlpaNUwxVmQ0R2VKZHNwNUJtQmd4UExtdlNDM1pjTm05dUdZZEZTQVFBQUFBJCQAAAAAAAAAAAEAAABFYtAlZ3VvYnVnYm95AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIbUqVKG1KlSU; NBID=6CA7F4A0ACE9D82536B66BDD95EE9A23:FG=1; BAIDU_DUP_lcr=https://www.google.com.hk/; cflag=65535:1; H_PS_PSSID=5750_5094_1452_5225_5287_5722_5823_5849_4261_5831_4760_5659_5857_5824",
	}

	FileW("科技", header)
	FileW("旅游", header)
	FileW("飞机", header)
	FileW("哥", header)
}
