package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

type Robot struct {
	roboName string
}

var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func (r Robot) Name() (string, error) {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, 2)
	randomDigit := 100 + rand.Intn(999-100)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return fmt.Sprintf("%s%d", string(b), randomDigit), nil
}
