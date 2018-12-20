/* 2018-12-20 (cc) <paul4hough@gmail.com>
   copied from github.com/prometheus/alertmanager/examples/webhook/echo.go
*/
package main

import (
        "bytes"
        "encoding/json"
        "io/ioutil"
        "log"
        "net/http"
)

func alertHandler(
	w http.ResponseWriter,
	r *http.Request ) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	var buf bytes.Buffer
	if err := json.Indent(&buf, b, " >", "  "); err != nil {
		panic(err)
	}
	log.Println(buf.String())
}

func main() {
	http.HandleFunc("/",alertHandler)
	log.Fatal(http.ListenAndServe(":5001",nil))
}
