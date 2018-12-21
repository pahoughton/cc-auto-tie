/* 2018-12-20 (cc) <paul4hough@gmail.com>
   copied from github.com/prometheus/alertmanager/examples/webhook/echo.go
*/
package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New(filepath.Base(os.Args[0]),
		"http dumper service")

	listenAddr = app.Flag("laddr","listen address").
		Short('l').
		Default(":5001").
		Envar("MAUL_LISTEN").
		String()

)

func main() {
	app.Version("0.0.1")
	kingpin.MustParse(app.Parse(os.Args[1:]))

	http.HandleFunc("/",alertHandler)
	log.Fatal(http.ListenAndServe(*listenAddr,nil))
}
