/* 2018-12-22 (cc) <paul4hough@gmail.com>
   create ticket(s) for received alerts

   need to save ticket id for automation to update
*/
package main

import (
	log     "github.com/sirupsen/logrus"
)

func createTicket(at AlertType, rBody []byte) {

	log.Debug("create ticket")

	ticketCreated.Inc()

}
