package output

import (
	"fmt"
	"github.com/tidwall/pretty"
)

// Print is used to print output json
func Print(b []byte, raw bool) {
	if !raw {
		result := pretty.Color(b, nil)
		fmt.Println(string(result))
	}

	if raw {
		fmt.Println(string(b))
	}
}
