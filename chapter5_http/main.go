// chapter5_http project main.go
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Fprintln(os.Stderr, "http error:%s", err.Error())
		os.Exit(1)
	}
	defer resp.Body.Close()

	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, "copy to stdout error:%s", err.Error())
		os.Exit(1)
	}

}
