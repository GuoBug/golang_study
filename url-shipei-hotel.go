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

	path := strings.Split(urlmap.Path, "/")
	stringNum := strings.Split(path[3], ".")
	filename := path[2] + "/" + stringNum[0] + ".xml"
	fmt.Println(filename)
	os.Mkdir(path[2], 0777)

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
		tempstring, _ := url.Parse(v.Url)

		path := strings.Split(tempstring.Path, "/")
		fmt.Println(path[1])
		fmt.Println("SSSSSSSS" + v.Url)

		sHtml := strings.Split(path[2], ".")

		/* 详情页 */
		if path[1] == "hotel" && sHtml[0] != path[2] {
			fp.WriteString("<url><loc><![CDATA[" + v.Url + "]]></loc>\n")
			fp.WriteString("<data>\n<display>\n<html5_url><![CDATA[")
			fp.WriteString("http://m.ctrip.com/html5/Hotel/HotelDetail/" + path[2])
			fp.WriteString("]]>\n</html5_url>\n</display>\n</data>\n</url>\n")

		} else if path[1] == "pic" {

			fp.WriteString("<url><loc><![CDATA[" + v.Url + "]]></loc>\n")
			fp.WriteString("<data>\n<display>\n<html5_url><![CDATA[")
			fp.WriteString(strings.Replace(v.Url, "hotels.ctrip.com", "m.ctrip.com/html5/Hotel/HotelDetail", 1))
			fp.WriteString("]]>\n</html5_url>\n</display>\n</data>\n</url>\n")
		} else if path[1] == "map" {

			fp.WriteString("<url><loc><![CDATA[" + v.Url + "]]></loc>\n")
			fp.WriteString("<data>\n<display>\n<html5_url><![CDATA[")
			fp.WriteString(strings.Replace(v.Url, "hotels.ctrip.com", "m.ctrip.com/html5/Hotel/HotelDetail", 1))
			fp.WriteString("]]>\n</html5_url>\n</display>\n</data>\n</url>\n")

		} else if path[1] == "hotel" && path[2] == "dianping" {

			fp.WriteString("<url><loc><![CDATA[" + v.Url + "]]></loc>\n")
			fp.WriteString("<data>\n<display>\n<html5_url><![CDATA[")
			fp.WriteString(strings.Replace(v.Url, "hotels.ctrip.com/hotel", "m.ctrip.com/html5/Hotel/HotelDetail", 1))
			fp.WriteString("]]>\n</html5_url>\n</display>\n</data>\n</url>\n")

		} else if path[1] == "international" {

			fp.WriteString("<url><loc><![CDATA[" + v.Url + "]]></loc>\n")
			fp.WriteString("<data>\n<display>\n<html5_url><![CDATA[")
			fp.WriteString(strings.Replace(v.Url, "hotels.ctrip.com/international", "m.ctrip.com/html5/oversea", 1))
			fp.WriteString("]]>\n</html5_url>\n</display>\n</data>\n</url>\n")

		} else {

			fp.WriteString("<url><loc><![CDATA[" + v.Url + "]]></loc>\n")
			fp.WriteString("<data>\n<display>\n<html5_url><![CDATA[")
			fp.WriteString(strings.Replace(v.Url, "hotels.ctrip.com", "m.ctrip.com/html5", 1))
			fp.WriteString("]]>\n</html5_url>\n</display>\n</data>\n</url>\n")

		}
	}
	fp.WriteString("</urlset>")

	fp.Close()
}

func main() {

	header := map[string]string{"Host": "ctrip.com",
		"User-Agent": "Mozilla/5.0 (Linux; Android 4.2.1; en-us; Nexus 4 Build/JOP40D) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166 Mobile Safari/535.19",
	}

	url := "http://hotels.ctrip.com/sitemap/sitemap.xml"

	spider := &Spider{url, header}

	html := spider.get_html_header()

	v := sitemapindex{}

	err := xml.Unmarshal(html, &v)
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range v.Sitemap {

		GetUrl(v.Url, header)

	}
}
