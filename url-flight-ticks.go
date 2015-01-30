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
	filename := path[1] 

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
	fp.WriteString("\txmlns:mobile=\"http://www.baidu.com/schemas/sitemap-mobile/1/\">\n")

	for _, v := range v.Loc {
		tempstring, _ := url.Parse(strings.Trim(v.Url, "\t"))

		path := strings.Split(tempstring.Path, "/")
		fmt.Println(path[2])
		//fmt.Println("SSSSSSSS" + v.Url)


		//low price
		//sHtml := strings.Split(path[2], "-")
		//fmt.Println(len(sHtml[6]))

		//inter_schedule_searchresult
		//sHtmlInter := strings.Split(path[3], "-")
		//fmt.Println(sHtmlInter)

		/* Sitemap_schedule1*/
		
		fmt.Println("Sitemap_schedule1")
		fp.WriteString("\t\t<url>\n\t\t\t<loc>")
		fp.WriteString("http://m.ctrip.com/html5/Flight/schedule/" + path[2])
		fp.WriteString("</loc>\n\t\t<mobile:mobile/>\n\t\t</url>\n")
		
		/*Sitemap_searchresult
		fmt.Println("Sitemap_searchresult")
		fp.WriteString("\t\t<url>\n\t\t\t<loc>")
		fp.WriteString("http://m.ctrip.com/html5/Flight/" + path[2])
		fp.WriteString("</loc>\n\t\t<mobile:mobile/>\n\t\t</url>\n")
		*/

		/* inter_schedule_searchresult 
		if len(sHtmlInter[0]) == 3 {
			fmt.Println("inter_schedule_searchresult")
			fp.WriteString("\t\t<url>\n\t\t\t<loc>")
			fp.WriteString("http://m.ctrip.com/html5/Flight/schedule/" + sHtmlInter[0] + "." + sHtmlInter[1])
			fp.WriteString("</loc>\n\t\t<mobile:mobile/>\n\t\t</url>\n")
		}
		*/

		
		/* low price
		if  len(sHtml[6]) == 3 && len(sHtml[7]) == 1 {
			//lowprice
			fmt.Println("In Lowprice")
			fp.WriteString("\t\t<url>\n\t\t\t<loc>")
			fp.WriteString("http://m.ctrip.com/html5/Flight/" + sHtml[1] + "-" +sHtml[2]+"-day-1.html")
			fp.WriteString("</loc>\n\t\t<mobile:mobile/>\n\t\t</url>\n")
		}
		*/
		//fmt.Println(sHtml[5])

		/* 详情页 sitemap baidu 
		if len(sHtml) == 4 {

			fp.WriteString("\t\t<url>\n\t\t\t<loc>")
			fp.WriteString("http://m.ctrip.com/html5/Flight/" + path[2])
			fp.WriteString("</loc>\n\t\t<mobile:mobile/>\n\t\t</url>\n")

			fp.WriteString("\t\t<url>\n\t\t\t<loc>")
			fp.WriteString("http://m.ctrip.com/html5/Flight/" + sHtml[0] + "-" + sHtml[1] + "-d.html")
			fp.WriteString("</loc>\n\t\t<mobile:mobile/>\n\t\t</url>\n")

		}*/
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

	GetUrl("http://flights.ctrip.com/Sitemap_schedule1.xml", header)
	//GetUrl("http://flights.ctrip.com/Sitemap_schedule1.xml", header)

	//GetUrl("http://flights.ctrip.com/inter_schedule_searchresult.xml", header)

	//GetUrl("http://flights.ctrip.com/Sitemap_lowprice.xml", header)

	//GetUrl("http://flights.ctrip.com/dflights-sitemap.xml", header)

	//GetUrl("http://flights.ctrip.com/dflights-baidu.xml", header)

}
