/* 2018-12-22 (cc) <paul4hough@gmail.com>
   receive /error http requests
*/
package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func errorHandler(
	w http.ResponseWriter,
	r *http.Request ) {

	log.Debug("error recvd")

	errorRecvd.Inc()

	log.Error("received")
}
