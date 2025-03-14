package s3testing

import (
	"fmt"
	"math/rand"

	"github.com/matthew188/aws-sdk-go-v2/internal/sdkio"
)

var randBytes = func() []byte {
	b := make([]byte, 10*sdkio.MebiByte)

	if _, err := rand.Read(b); err != nil {
		panic(fmt.Sprintf("failed to read random bytes, %v", err))
	}
	return b
}()

// GetTestBytes returns a pseudo-random []byte of length size
func GetTestBytes(size int) []byte {
	if len(randBytes) >= size {
		return randBytes[:size]
	}

	b := append(randBytes, GetTestBytes(size-len(randBytes))...)
	return b
}
