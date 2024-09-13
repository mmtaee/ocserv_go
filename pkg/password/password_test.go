package password

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	password string
	salt     string
)

func init() {
	password = "password"
	salt = "123"
}

func TestCreatePassword(t *testing.T) {
	hash := MakeHash(password, salt)
	assert.NotEmpty(t, hash)
}

func TestCheckHash(t *testing.T) {
	hash := MakeHash(password, salt)
	assert.NotEmpty(t, hash)
	check := Compare(password, salt, hash)
	assert.True(t, check)
}

func TestCreateRandom(t *testing.T) {
	passwd := CreateRandom(8)
	assert.NotEmpty(t, passwd)
	assert.Equal(t, len(passwd), 8)
}

func BenchmarkCreateHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MakeHash(fmt.Sprintf("%s%d", password, i), salt)
	}
}

func BenchmarkCreateRandom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CreateRandom(8)
	}
}
