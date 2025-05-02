package passw

import (
	"github.com/alexedwards/argon2id"
	"runtime"
)

var params = &argon2id.Params{
	Memory:      64 * 1024,
	Iterations:  1,
	Parallelism: uint8(runtime.NumCPU()),
	SaltLength:  16,
	KeyLength:   32,
}

func CreateHash(password string) (string, error) {
	hash, err := argon2id.CreateHash(password, params)
	if err != nil {
		return "", err
	}
	return hash, nil
}
