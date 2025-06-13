package queries

import (
	"fmt"
	"io"
	"net/http"

	"github.com/mxssl/dns/output"
)

// GetRptr is used for dns ptr requests
func GetRptr(resolver string, ip string, raw bool) {
	url := fmt.Sprintf("%s/%s/x/%s", APIURL, resolver, ip)

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
