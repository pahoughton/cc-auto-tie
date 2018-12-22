/* 2018-12-21 (cc) <paul4hough@gmail.com>
   handle ticket http requests
*/
package main

import (
	"fmt"
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

	resp := fmt.Sprintf(`<!DOCTYPE html>
<html>
<body>
<h2> Ticket Created </h2>
<p><b> payload </b></p>
<pre>
%s
</pre>
</body>
</html>
`,b)
	log.Info("ticket created")
	log.Debug(resp)
	w.WriteHeader(200)
	w.Write([]byte(resp))
}
