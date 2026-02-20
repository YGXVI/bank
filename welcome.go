package hm

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	name := os.Args[1]
	code := os.Args[2]

	message := "Салом, {username}! Ваш код доступа: {code}."

	message = strings.ReplaceAll(message, "{username}", name)
	message = strings.ReplaceAll(message, "{code}", code)

	fmt.Println(message)
}