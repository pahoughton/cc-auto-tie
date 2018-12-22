/* 2018-12-20 (cc) <paul4hough@gmail.com>
   FIXME what is this for?
*/
package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func ansibleHandler(
	w http.ResponseWriter,
	r *http.Request ) {

	log.Debug("ansible recvd")

	ansibleRecvd.Inc()

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

	createTicket(AnsibleAlertType, b)

	log.Debug("ansible complete")
}
