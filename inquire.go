package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

func sendmail(reply_to string, subject string, body string) {
	cmd := exec.Command("/usr/sbin/sendmail", "jj@rtts.eu")
	email := "To: jj@rtts.eu\nReply-To: " + reply_to + "\nSubject: " + subject + "\n\n" + body
	cmd.Stdin = strings.NewReader(email)
	cmd.Run()
}

func mailhandler(w http.ResponseWriter, r *http.Request) {
	reply_to := r.FormValue("email")
	body := r.FormValue("message")
	redirect := r.FormValue("redirect")
	if reply_to == "" || body == "" || redirect == "" {
		log.Println("Bad request from " + r.RemoteAddr)
		http.Error(w, "Get off my lawn!", 400)
	} else {
		log.Println("Valid inquiry from " + reply_to)
		sendmail(reply_to, "Contactformulier", body)
		http.Redirect(w, r, redirect, http.StatusFound)
	}
}

func main() {
	fmt.Println("Listening for inquiries at port 8008")
	http.HandleFunc("/", mailhandler)
	log.Fatal(http.ListenAndServe(":8008", nil))
}