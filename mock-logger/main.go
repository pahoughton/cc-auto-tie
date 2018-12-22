/* 2018-12-22 (cc) <paul4hough@gmail.com>
   mock log service
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

	laddr = app.Flag("laddr","listen address").
		Short('l').
		Default(":5002").
		Envar("MOCK_LOG_SERVICE_LISTEN").
		String()

	logfn = app.Flag("log-fn","log filename").
		Short('f').
		Default("logger.log").
		Envar("MOCK_LOG_FILENAME").
		String()


	infoRecvd = proma.NewCounter(
		prom.CounterOpts{
			Name:      "info_received_total",
			Help:      "number of info requests received",
		})
	warnRecvd = proma.NewCounter(
		prom.CounterOpts{
			Name:      "warn_received_total",
			Help:      "number of warning requests received",
		})
	errorRecvd = proma.NewCounter(
		prom.CounterOpts{
			Name:      "error_received_total",
			Help:      "number of error requests received",
		})
	unsupRecvd = proma.NewCounter(
		prom.CounterOpts{
			Name:      "unsupported_received_total",
			Help:      "number of unsupported request received",
		})
)

func main() {

	app.Version("0.0.1")
	kingpin.MustParse(app.Parse(os.Args[1:]))

	if logfn != nil {
		logf, err := os.OpenFile(*logfn,
			os.O_APPEND|os.O_CREATE|os.O_WRONLY,
			0666)
		if  err != nil {
			panic(err)
		}
		log.SetOutput( logf )
	}

	log.SetLevel(log.TraceLevel)
	log.Info(os.Args[0]," listening on ",*laddr)

	http.HandleFunc("/info",infoHandler)
	http.HandleFunc("/warn",warnHandler)
	http.HandleFunc("/error",errorHandler)
	http.HandleFunc("/",defaultHandler)

	http.Handle("/metrics", promh.Handler())

	log.Fatal(http.ListenAndServe(*laddr,nil))
}
