package password

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"ocserv/pkg/config"
	"time"
)

// MakeHash create hash from password with salt string
func MakeHash(password, salt string) string {
	app := config.GetApp()
	saltPassword := fmt.Sprintf("%s%s%s", password, salt, app.SecretKey)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(saltPassword), 10)
	if err != nil {
		return ""
	}
	return string(hashedPassword)
}

// Compare check password with hash
func Compare(password, salt, hashedPassword string) bool {
	app := config.GetApp()
	saltPassword := fmt.Sprintf("%s%s%s", password, salt, app.SecretKey)
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(saltPassword))
	return err == nil
}

// CreateRandom create ocserv password
func CreateRandom(length int) string {
	if length < 1 {
		length = 8
	}
	const validChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789*@$!"
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	password := make([]byte, length)
	for i := 0; i < length; i++ {
		password[i] = validChars[r.Intn(len(validChars))]
	}
	return string(password)
}
