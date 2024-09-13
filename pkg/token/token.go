package token

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"ocserv/pkg/config"
	"time"
)

// Create a token for login admin or staffs
func Create(userId uint, expireAt int64) string {
	app := config.GetApp()
	now := time.Now().Unix()
	str := fmt.Sprintf("%d%v%s%v", now, userId, app.SecretKey, expireAt)
	hashed := sha256.New()
	hashed.Write([]byte(str))
	hash := hashed.Sum(nil)
	hashHex := hex.EncodeToString(hash)
	return hashHex
}
