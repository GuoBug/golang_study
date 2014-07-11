package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func GetUrl(u string) {

	fp, err := os.Create("tuan_shipei.xml")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	from, err := os.Open(u)
	if err != nil {
		panic(err)
	}
	defer from.Close()

	fp.WriteString("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<urlset>\n")

	inputread := bufio.NewReader(from)
	for {
		input, ferr := inputread.ReadString('\n')
		if ferr == io.EOF {
			break
		}
		path := strings.Split(input, "/")
		fmt.Println(path[5])
		fp.WriteString("<url><loc><![CDATA[")
		fp.WriteString("http://tuan.ctrip.com/group/" + path[5] + "/")
		fp.WriteString("]]></loc>\n")
		fp.WriteString("<data>\n<display>\n<html5_url><![CDATA[" + input + "]]>\n</html5_url>\n</display>\n</data>\n</url>\n")
	}
	fp.WriteString("</urlset>")

	fp.Close()
}

func main() {

	GetUrl("tuan.xml")

}
