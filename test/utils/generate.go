package e2e

import (
	"math/rand"
	"time"
)

func GeneratePassportData() int {
	rand.Seed(time.Now().UnixNano())
	min, max := 1111111111, 9999999999
	number := rand.Intn(max-min) + min
	return number
}
