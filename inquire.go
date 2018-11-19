package main

import (
	"fmt"
	"strings"
	"os/exec"
	"net/http"
)

func sendmail(reply_to string, subject string, body string) {
	cmd := exec.Command("/usr/sbin/sendmail", "jj@rtts.eu")
	email := "To: jj@rtts.eu\nReply-To: " + reply_to + "\nSubject: " + subject + "\n\n" + body
	cmd.Stdin = strings.NewReader(email)
	cmd.Run()
}

func mailhandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}
	reply_to := r.FormValue("email")
	body := r.FormValue("message")
	redirect := r.FormValue("redirect")

	// Bad request, log IP address and return 400
	if reply_to == "" || body == "" || redirect == "" {
                ip := r.Header.Get("X-Real-IP")
                if ip == "" {
			ip = r.RemoteAddr
		}
		fmt.Println("Bad request from " + ip)
		http.Error(w, "Get off my lawn!", 400)

	// Good request, send mail and return a 302 redirect
	} else {
		fmt.Println("Valid inquiry from " + reply_to)
		sendmail(reply_to, "Contactformulier", body)
		http.Redirect(w, r, redirect, http.StatusFound)
	}
}

func main() {
	fmt.Println("Listening for inquiries at port 8008")
	http.HandleFunc("/", mailhandler)
	http.ListenAndServe(":8008", nil)
}
