package main

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
)

/* copy from tonyhb
URL :https://gist.github.com/tonyhb/5819315
*/

type Person struct {
	Name string
	Age  uint
}

func main() {
	betty := Person{
		Name: "Betty",
		Age:  82,
	}

	urlValues := structToMap(&betty)
	fmt.Printf("%#v", urlValues)
	// Response...
	// url.Values{"Name":[]string{"Betty"}, "Age":[]string{"82"}}
	// It works! Play with it @ http://play.golang.org/p/R8jhzfI_ac
}

func structToMap(i interface{}) (values url.Values) {
	values = url.Values{}
	iVal := reflect.ValueOf(i).Elem()
	typ := iVal.Type()
	for i := 0; i < iVal.NumField(); i++ {
		f := iVal.Field(i)
		// You ca use tags here...
		// tag := typ.Field(i).Tag.Get("tagname")
		// Convert each type into a string for the url.Values string map
		var v string
		switch f.Interface().(type) {
		case int, int8, int16, int32, int64:
			v = strconv.FormatInt(f.Int(), 10)
		case uint, uint8, uint16, uint32, uint64:
			v = strconv.FormatUint(f.Uint(), 10)
		case float32:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 32)
		case float64:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 64)
		case []byte:
			v = string(f.Bytes())
		case string:
			v = f.String()
		}
		values.Set(typ.Field(i).Name, v)
	}
	return
}
