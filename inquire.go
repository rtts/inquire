package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"
)

var ALLOWED = map[string]bool{
	"jj@rtts.eu":                     true,
	"jj@returntothesource.nl":        true,
	"info@superformosa.nl":           true,
	"fraukevandenbogaard@gmail.com":  true,
	"frauke@studiomanna.nl":          true,
}

func sendmail(to string, reply_to string, subject string, body string) {
	cmd := exec.Command("/usr/sbin/sendmail", to)
	email := "To: " + to + "\nReply-To: " + reply_to + "\nSubject: " + subject + "\n\n" + body
	cmd.Stdin = strings.NewReader(email)
	cmd.Run()
}

func mailhandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}
	to := r.FormValue("to")
	reply_to := r.FormValue("reply_to")
	subject := r.FormValue("subject")
	redirect := r.FormValue("redirect")

	// Spammers will never figure this out }:‑)
	body := r.FormValue("get_off_my_lawn")
	spam_detect := r.FormValue("message")

	// Bad request, log IP address but still return a "success" redirect
	if to == "" || reply_to == "" || subject == "" || body == "" || redirect == "" || spam_detect != "" || !ALLOWED[to] {
		ip := r.Header.Get("X-Real-IP")
		if ip == "" {
			ip = r.RemoteAddr
		}
		fmt.Println("Bad request from " + ip)
		http.Redirect(w, r, redirect, http.StatusFound)

	// Good request, send mail and return a 302 redirect
	} else {
		fmt.Println("Valid inquiry from " + reply_to)
		sendmail(to, reply_to, subject, body)
		http.Redirect(w, r, redirect, http.StatusFound)
	}
}

func main() {
	fmt.Println("Listening for inquiries at port 8008")
	http.HandleFunc("/", mailhandler)
	http.ListenAndServe(":8008", nil)
}
