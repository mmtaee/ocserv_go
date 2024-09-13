package token

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreateToken(t *testing.T) {
	var (
		userId   = uint(1)
		expireAt = time.Now().Add(time.Hour * 24).Unix()
	)
	token := Create(userId, expireAt)
	assert.NotEmpty(t, token)
}

func BenchmarkCreateToken(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Create(uint(i+1), time.Now().Unix())
	}
}
