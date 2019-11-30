package queries

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mxssl/dns/output"
)

// GetRptr is used for dns ptr requests
func GetRptr(resolver string, ip string, raw bool) {
	url := fmt.Sprintf(APIURL + "/" + resolver + "/x/" + ip)

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
