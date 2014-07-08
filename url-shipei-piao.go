package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Spider struct {
	url    string
	header map[string]string
}

type sitemap struct {
	XMLName xml.Name `xml:"sitemap"`
	Url     string   `xml:"loc"`
	time    string   `xml:"lastmod"`
}

type urlset struct {
	XMLName     xml.Name    `xml:"urlset"`
	Loc         []Urlstruct `xml:"url"`
	Description string      `xml:",innerxml"`
}

type Urlstruct struct {
	XMLName xml.Name `xml:"url"`
	Url     string   `xml:"loc"`
}

type sitemapindex struct {
	XMLName     xml.Name  `xml:"sitemapindex"`
	Version     string    `xml:"version,attr"`
	Sitemap     []sitemap `xml:"sitemap"`
	Description string    `xml:",innerxml"`
}

func (keyword Spider) get_html_header() []byte {

	client := &http.Client{}

	req, err := http.NewRequest("GET", keyword.url, nil)
	if err != nil {
	}

	for key, value := range keyword.header {

		req.Header.Add(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
		return []byte("err")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
		return body
	}

	defer resp.Body.Close()

	return body

}

func GetUrl(u string, s map[string]string) {
	urlmap, err := url.Parse(u)
	if err != nil {
		panic(err)
	}
	fmt.Println(urlmap.Path)

	path := strings.Split(urlmap.Path, "/")
	filename := path[1] + ".xml"

	fp, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	spider := &Spider{u, s}

	html := spider.get_html_header()

	v := urlset{}

	err = xml.Unmarshal(html, &v)
	if err != nil {
		panic(err)
	}

	fp.WriteString("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<urlset>\n")

	for _, v := range v.Loc {
		tempstring, _ := url.Parse(strings.Trim(v.Url, "\t"))

		path := strings.Split(tempstring.Path, "/")
		//fmt.Println(path)
		//fmt.Println("SSSSSSSS" + v.Url)

		fp.WriteString("<url><loc><![CDATA[" + v.Url + "]]></loc>\n")
		fp.WriteString("<data>\n<display>\n<html5_url><![CDATA[")
		fp.WriteString("http://m.ctrip.com/html5/ticket/dest/" + path[2])
		fp.WriteString("]]>\n</html5_url>\n</display>\n</data>\n</url>\n")

	}
	fp.WriteString("</urlset>")

	fp.Close()
}

func main() {

	header := map[string]string{"Host": "ctrip.com",
		"User-Agent": "Mozilla/5.0 (Linux; Android 4.2.1; en-us; Nexus 4 Build/JOP40D) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166 Mobile Safari/535.19",
	}

	GetUrl("http://piao.ctrip.com/sitemap/piao/piaodetail.xml", header)

}
