package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"strconv"
	"os/exec"
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"io"
)

type ErrForRedirect struct {
	Id      string
	Message string
}

func main() {

	fp, err := os.Create("BaiduURL.txt")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	fpWeb, err := os.Create("BaiduWEB.txt")
	if err != nil {
		panic(err)
	}
	defer fpWeb.Close()

	for i :=36001 ; i <= 100000;i++ {
		
		sInUrl  := "http://m.bilibili.com/video/av" + strconv.Itoa(i) +".html"
		//get host
		sHost := strings.Split(sInUrl, "/")

		//fmt.Println(sHost[2])

		header := map[string]string{
			"Host":            sHost[2],
			"User-Agent":      "Mozilla/5.0 (iPhone; U; CPU iPhone OS 4_0 like Mac OS X; en-us;) AppleWebKit/532.9 (KHTML, like Gecko) Version/4.0.5 Mobile/8A293 Safari/6531.22.7",
			"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
			"Accept-Encoding": "gzip,deflate,sdch",
			"Cache-Control":   "max-age=0",
			"Connection":      "keep-alive",
			"Accept-Language": "zh-CN,zh;q=0.8,en;q=0.6,zh-TW;q=0.4,ja;q=0.2",
		}
		// fmt.Println(strings.Trim(sInUrl, ""))

		req, err := http.NewRequest("", sInUrl, nil)
		if err != nil {
		}

		for key, value := range header {
			//fmt.Println(key + ":" + value + "\n")
			req.Header.Add(key, value)
		}

		//fmt.Println(req.Header)

		//client
		client := http.Client{CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return err
		}}

		//fmt.Println(req.UserAgent())

		resp, err := client.Do(req)
		if err != nil {
		}
		var reader io.ReadCloser

		if resp.Header.Get("Content-Encoding") == "gzip" {
			reader, err = gzip.NewReader(resp.Body)
			if err != nil {
				return
			}
		} else {
			reader = resp.Body
		}

		body, err := ioutil.ReadAll(reader)
		if err != nil {

		}
		//fmt.Println(string(body))

		//fmt.Println("response:\n" + resp.Status + resp.Header.Get("Location"))
		if resp.Status == "200 OK" {
			if strings.Count(string(body),"视频不见了哟") == 1{
				//fmt.Println(i)
			}else{
				fp.WriteString(sInUrl + "\n")
				fpWeb.WriteString("http://www.bilibili.com/video/av" + strconv.Itoa(i) + "/\n")
			}
		}
		resp.Body.Close()

		if i % 2000 == 0{
		
			fp.Close()
			var out bytes.Buffer //缓冲字节
			cmd := exec.Command("/bin/sh", "-c", `curl -H 'Content-Type:text/plain' --data-binary @BaiduURL.txt "http://data.zz.baidu.com/urls?site=m.bilibili.com&token=gvEuIzrUYLunH7GK&type=original"`)
			fmt.Printf("%s\n", out.String())

		    fp, err = os.Create("BaiduURL.txt")
			if err != nil {
				panic(err)
			}
			defer fp.Close()

			fpWeb.Close()
			cmd = exec.Command("/bin/sh", "-c", `curl -H 'Content-Type:text/plain' --data-binary @BaiduWEB.txt "http://data.zz.baidu.com/urls?site=www.bilibili.com&token=gvEuIzrUYLunH7GK&type=original"`)
		    cmd.Stdout = &out
		    err = cmd.Run() //运行指令 ，做判断
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("%s\n", out.String())

		    fpWeb, err = os.Create("BaiduWEB.txt")
			if err != nil {
				panic(err)
			}
			defer fpWeb.Close()

		}

	}
	fp.Close()
	fpWeb.Close()

}
