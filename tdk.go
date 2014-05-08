package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
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

func FileW(keyword string, header map[string]string) (string, string, string) {

	//pc	url := "http://suggestion.baidu.com/su?json=1&wd=" + keyword
	url := keyword

	spider := &Spider{url, header}

	html := spider.get_html_header()

	rp1 := regexp.MustCompile("<title>.*?</title>") //找出【】中的内容
	find_txt := rp1.FindAllString(html, -1)

	rp2 := regexp.MustCompile("name=\"keywords\".*?/>") //找出【】中的内容
	find_key := rp2.FindAllString(html, -1)

	rp3 := regexp.MustCompile("name=\"description\".*?/>") //找出【】中的内容
	find_description := rp3.FindAllString(html, -1)

	return find_txt[0], find_key[0], find_description[0]
}

func GetKeyWord(c chan string) {
	fmt.Println("GKW start")

	f, err := os.Open("key.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buff := bufio.NewReader(f)

	fmt.Println(">>>>>>>")

	for {
		w, e := buff.ReadString('\n')
		if e != nil || io.EOF == e {
			fmt.Println(w)
			break
		}
		time.Sleep(time.Second * 1)

		w = strings.Replace(w, "\n", "", 1)
		fmt.Println(w)
		c <- w
	}

	close(c)
}

func main() {
	fmt.Println("start")
	header := map[string]string{
		"Host":       "flights.ctrip.com",
		"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/34.0.1847.131 Safari/537.36",
	}

	key_word_channel := make(chan string)

	go GetKeyWord(key_word_channel)

	f, err := os.OpenFile("tdk.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for w := range key_word_channel {
		//传入后处理
		key_title, key_word, key_des := FileW(w, header)
		time.Sleep(time.Second * 2)
		//
		//		fmt.Println(key_word)
		f.WriteString(key_title + "\t" + key_word + "\t" + key_des + "\n")
	}

	f.Close()
}
