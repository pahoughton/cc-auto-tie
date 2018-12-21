/* 2018-12-20 (cc) <paul4hough@gmail.com>
   copied from github.com/prometheus/alertmanager/examples/webhook/echo.go
*/
package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/",alertHandler)
	log.Fatal(http.ListenAndServe(":5001",nil))
}
