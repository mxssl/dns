package queries

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

var APIURL = "http://www.dns-lg.com"

var Resolver string

func GetQ(resolver string, queryType string, domain string) {
	url := fmt.Sprintf(APIURL + "/" + resolver + "/" + domain + "/" + queryType)

	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(content))
}
