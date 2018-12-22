/* 2018-12-22 (cc) <paul4hough@gmail.com>
   receive /warn http requests
*/
package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func warnHandler(
	w http.ResponseWriter,
	r *http.Request ) {

	log.Debug("warn recvd")

	warnRecvd.Inc()

	log.Warn("received")
}
