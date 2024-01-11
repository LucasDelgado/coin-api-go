package bd

import "golang.org/x/crypto/bcrypt"

func EncryptPass(pass string) (string, error) {
	costo := 6
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)

	return string(bytes), err
}
