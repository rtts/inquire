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

