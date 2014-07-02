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
	stringNum := strings.Split(path[1], ".")
	filename := stringNum[0] + ".xml"

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

	fp.WriteString("<?xml version=\"1.0\" encoding=\"UTF-8\" ?>\n")
	fp.WriteString("\t<urlset xmlns=\"http://www.sitemaps.org/schemas/sitemap/0.9\"\n")
	fp.WriteString("\txmlns:mobile=\"http://www.google.com/schemas/sitemap-mobile/1.0\">\n")

	for _, v := range v.Loc {
		tempstring, _ := url.Parse(v.Url)

		path := strings.Split(tempstring.Path, "/")
		fmt.Println(path[1])
		fmt.Println("SSSSSSSS" + v.Url)

		sHtml := strings.Split(path[2], ".")

		/* 详情页 */
		if len(sHtml) == 3 {

			fp.WriteString("\t\t<url>\n\t\t\t<loc>")
			fp.WriteString("http://m.ctrip.com/html5/Flight/Schedule/" + path[2])
			fp.WriteString("</loc>\n\t\t<mobile:mobile/>\n\t\t</url>\n")

		} else if len(sHtml) == 2 {
			fp.WriteString("\t\t<url>\n\t\t\t<loc>")
			fp.WriteString("http://m.ctrip.com/html5/Flight/Schedule/" + path[2])
			fp.WriteString("</loc>\n\t\t<mobile:mobile/>\n\t\t</url>\n")
		}
	}
	fp.WriteString("</urlset>")

	fp.Close()
}

func main() {

	header := map[string]string{"Host": "ctrip.com",
		"User-Agent": "Mozilla/5.0 (Linux; Android 4.2.1; en-us; Nexus 4 Build/JOP40D) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166 Mobile Safari/535.19",
	}

	url := "http://flights.ctrip.com/Sitemap_schedule1.xml"

	spider := &Spider{url, header}

	html := spider.get_html_header()

	v := urlset{}

	err := xml.Unmarshal(html, &v)
	if err != nil {
		fmt.Println(err)
	}

	GetUrl(url, header)

	GetUrl("http://flights.ctrip.com/Sitemap_schedule2.xml", header)

}
