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

	prom  "github.com/prometheus/client_golang/prometheus"
	proma "github.com/prometheus/client_golang/prometheus/promauto"
	promh "github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	app = kingpin.New(filepath.Base(os.Args[0]),
		"http dumper service")

	listenAddr = app.Flag("laddr","listen address").
		Short('l').
		Default(":5001").
		Envar("MAUL_LISTEN").
		String()

	metSpace = "cca"
	metSubSys = "maul"

	alertsRecvd = proma.NewCounter(
		prom.CounterOpts{
			Namespace: metSpace,
			Subsystem: metSubSys,
			Name:      "alert_received_total",
			Help:      "number of alerts received",
		})
)

func main() {
	app.Version("0.0.1")
	kingpin.MustParse(app.Parse(os.Args[1:]))

	http.Handle("/metrics", promh.Handler())
	http.HandleFunc("/",alertHandler)

	log.Fatal(http.ListenAndServe(*listenAddr,nil))
}
