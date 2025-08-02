package services

import "math/rand"

func Roll(sides int, bonus int) int {
	return rand.Intn(sides) + bonus
}
