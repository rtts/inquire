Inquire
=======

Remember the old days when you would add a contact form to your
homepage with something like this code?

    <FORM METHOD ="post" ACTION="/cgi-bin/mail-a-form">
    <P>Van (e-mailadres): <INPUT TYPE="text" NAME="from"><BR>
    <INPUT TYPE="hidden" NAME="subject" VALUE="Formulier verstuurd vanaf paginanaam">
    <INPUT TYPE="hidden" NAME="missing" VALUE="http://loginnaam.home.xs4all.nl/">
    <INPUT TYPE="hidden" NAME="nextpage" VALUE="http://loginnaam.home.xs4all.nl/">
    <INPUT TYPE="hidden" NAME="to" VALUE="loginnaam@xs4all.nl">
    <TEXTAREA NAME="veld3" ROWS="5" COLS="40">Dit is een veld waarin veel tekst kan worden ingevuld</TEXTAREA><BR>
    <INPUT TYPE="submit" VALUE="Verstuur">
    <INPUT TYPE="reset" VALUE="Wis">
    </FORM>

Well, first of all, these days aren't as you might think! The HTML
sample code above was downloaded in November 2018 from
https://www.xs4all.nl/service/diensten/hosting-en-homepage/gebruiken/shared-webhosting/cgi-scripts/mailaform.htm

Inquire does more or less the same thing as `mail-a-form` above,
except that it doesn't let the caller define the "To" address for
obvious reasons. After starting the server with you can embed a
contact form like this:

    <form action="https://inquire.rtts.eu/" method="post">
      <input type="email" name="email">
      <input type="hidden" name="redirect" value="https://rtts.eu">
      <textarea name="message"></textarea>
      <button>Send</button>
    </form>

If you want to change the "To" address or email subject, please edit
the source code and recompile with `go build`

Installation
------------

This couldn't be easier! Just copy the binary anywhere you'd like and
run it. The server wil listen on port 8008 and log to standard
output. Of course, you should use some kind of process monitor like
systemd in production environments. Here's a sample systemd service
file:

    [Unit]
    Description = Inquire

    [Service]
    ExecStart = /opt/inquire/inquire
    User = www-data
    Restart = always

    [Install]
    WantedBy = multi-user.target

Also, here's an example configuration to use nginx as a reverse proxy:

    server {
      server_name inquire.rtts.eu;
      listen 80;
      listen 443 ssl;
      ssl_certificate inquire.rtts.eu.chained.crt;
      ssl_certificate_key inquire.rtts.eu.pem;

      location / {
        proxy_set_header X-Real-IP $remote_addr;
        proxy_pass http://127.0.0.1:8008;
      }

      location /.well-known/acme-challenge {
        alias /etc/nginx/challenges/inquire.rtts.eu;
      }
    }
