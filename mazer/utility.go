package mazer

import (
	"math/rand"
	"time"
)

func Roll(amount, size int) int {
	rand.Seed(time.Now().UnixNano())
	var sum = 0
	for i := 0; i < amount; i++ {
		sum += rand.Intn(size) + 1
	}
	return sum
}