/* 2018-12-22 (cc) <paul4hough@gmail.com>
   receive /info http requests
*/
package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func infoHandler(
	w http.ResponseWriter,
	r *http.Request ) {

	log.Debug("info recvd")

	infoRecvd.Inc()

	log.Info("received")
}
