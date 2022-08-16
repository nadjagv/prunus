package util

import "golang.org/x/crypto/bcrypt"

func Hesiraj(lozinka string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(lozinka), bcrypt.DefaultCost)

}
