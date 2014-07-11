package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	//"io/ioutil"
)

type ErrForRedirect struct {
	Id      string
	Message string
}

func main() {

	from, err := os.Open("url.txt")
	if err != nil {
		panic(err)
	}
	defer from.Close()

	inputread := bufio.NewReader(from)

	fp, err := os.Create("跳转.txt")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	for {
		input, ferr := inputread.ReadString('$')
		if ferr == io.EOF {
			break
		}
		sInUrl := strings.Trim(input, "$")
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

		fmt.Println(req.UserAgent())

		resp, err := client.Do(req)
		if err != nil {
		}
		//body, err := ioutil.ReadAll(resp.Body)
		if err != nil {

		}

		fmt.Println("response:\n" + resp.Status + resp.Header.Get("Location"))
		fp.WriteString(sInUrl + "\t" + resp.Header.Get("Location") + "\t" + resp.Status + "\n")

		resp.Body.Close()
	}
	from.Close()
	fp.Close()

}
