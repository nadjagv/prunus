package servis

import "net/smtp"

var from = "projekatxml@gmail.com"
var lozinka = "wahkfnrnxjhojwij" //wahkfnrnxjhojwij

var host = "smtp.gmail.com"
var port = "587"
var address = host + ":" + port

func PosaljiMejl(subject string, body string, mejlAdresa string) error {
	message := []byte("Subject:" + subject + "\n\n" + body)

	to := []string{mejlAdresa}

	auth := smtp.PlainAuth("", from, lozinka, host)

	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		return err
	}

	return nil
}
