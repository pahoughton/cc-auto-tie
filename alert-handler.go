/* 2018-12-20 (cc) <paul4hough@gmail.com>
   FIXME what is this for?
*/
package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func alertHandler(
	w http.ResponseWriter,
	r *http.Request ) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	var buf bytes.Buffer
	if err := json.Indent(&buf, b, " >", "  "); err != nil {
		panic(err)
	}
	log.Println(buf.String())
}
