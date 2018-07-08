package queries

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func GetRptr(resolver string, ip string) {
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

	fmt.Println(string(content))
}
