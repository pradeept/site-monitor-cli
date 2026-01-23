/*
This file is used as a playground. Will be trashed later.
*/
package main

import (
	"fmt"
	"os"
	"strings"
)

// STORE: SQlit
// Requests: net/http
// CLI : flag
// STD out: log / fmt

// add website (website string, time time.Seconds)
// remove website (website string)
// list website stats () []websites

// LIVE VIEW: live refresh of websites request status
// S. NO. | Website Name | Request history | total requests | fail rate | success rate

func main() {

	os.Setenv("secret", "This is my little secret")
	os.Setenv("aws-key", "Zuckerburg took it, I don't have it now!")

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0], "=", pair[1])
	}
}
