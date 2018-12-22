/* 2018-12-21 (cc) <paul4hough@gmail.com>
   handle ticket http requests
*/
package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func ticketHandler(
	w http.ResponseWriter,
	r *http.Request ) {

	// inc metrics counter
	ticketRecvd.Inc()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	// debug - format json
	var buf bytes.Buffer
	if err := json.Indent(&buf, b, " >", "  "); err != nil {
		panic(err)
	}
	log.Debug("req body\n",buf.String())

	createTicket(TicketAlertType, b)

	log.Debug("ticket req complete")
}
