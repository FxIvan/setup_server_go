package util

import (
	"math/rand"

	mongodb_model "github.com/fxivan/set_up_server/microservice/internal/adapter/storage/mogodb/model"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandSeq(n int, amount int) []string {
	var items []string

	for i := 0; i <= amount; i++ {
		b := make([]rune, n)
		for i := range b {
			b[i] = letters[rand.Intn(len(letters))]
		}
		items = append(items, string(b))
	}

	return items
}

func SearchCode(code string, allCode []mongodb_model.CodeCoupon) (bool, string) {
	for _, v := range allCode {
		if v.Code == code {
			return true, code
		} else {
			return false, code
		}
	}
	return false, code
}
