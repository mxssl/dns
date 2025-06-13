package queries

import (
	"fmt"
	"io"
	"net/http"

	"github.com/mxssl/dns/output"
)

// APIURL is DNS-LG endpoint
var APIURL = "http://www.dns-lg.com"

// Resolver is what resolver will we use
var Resolver string

// GetQ is used for http queries
func GetQ(resolver string, queryType string, domain string, raw bool) {
	url := fmt.Sprintf("%s/%s/%s/%s", APIURL, resolver, domain, queryType)

	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		if err := response.Body.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	content, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	output.Print(content, raw)
}
