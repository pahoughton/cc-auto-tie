/* 2018-12-20 (cc) <paul4hough@gmail.com>
   copied from github.com/prometheus/alertmanager/examples/webhook/echo.go
*/
package main

import (
	"net/http"
	"os"
	"path/filepath"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
	log     "github.com/sirupsen/logrus"

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

	ansibleRecvd = proma.NewCounter(
		prom.CounterOpts{
			Name:      "ansible_received_total",
			Help:      "number of ansible tasks received",
		})
	emmetRecvd = proma.NewCounter(
		prom.CounterOpts{
			Name:      "emmet_received_total",
			Help:      "number of emmet flows received",
		})
	ticketRecvd = proma.NewCounter(
		prom.CounterOpts{
			Name:      "ticket_received_total",
			Help:      "number of tickets requests received",
		})
	unsupRecvd = proma.NewCounter(
		prom.CounterOpts{
			Name:      "unsupported_received_total",
			Help:      "number of unsupported request received",
		})
)

func main() {

	app.Version("0.0.2")
	kingpin.MustParse(app.Parse(os.Args[1:]))

	log.SetLevel(log.TraceLevel)
	log.Info("Starting ",os.Args[0])

	http.HandleFunc("/ansible",ansibleHandler)
	http.HandleFunc("/emmet",emmetHandler)
	http.HandleFunc("/ticket",ticketHandler)
	http.HandleFunc("/",defaultHandler)

	http.Handle("/metrics", promh.Handler())

	log.Fatal(http.ListenAndServe(*listenAddr,nil))
}
