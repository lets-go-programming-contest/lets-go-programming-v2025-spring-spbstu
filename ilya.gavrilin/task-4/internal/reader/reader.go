package reader

import (
	"fmt"
	"math/rand"
)

var paths = []string{"home", "about"}

func GenerateRandomURL() string {
	path := paths[rand.Intn(len(paths))]
	param := rand.Intn(2)
	return fmt.Sprintf("/%s?id=%d", path, param)
}