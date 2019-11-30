package queries

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mxssl/dns/output"
)

// APIURL is DNS-LG endpoint
var APIURL = "http://www.dns-lg.com"

// Resolver is what resolver will we use
var Resolver string

// GetQ is used for http queries
func GetQ(resolver string, queryType string, domain string, raw bool) {
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

	output.Print(content, raw)
}
