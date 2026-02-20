
package hm

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	phone := os.Args[1]

	phone = strings.ReplaceAll(phone, " ", "")
	phone = strings.ReplaceAll(phone, "(", "")
	phone = strings.ReplaceAll(phone, ")", "")
	phone = strings.ReplaceAll(phone, "-", "")

	fmt.Println(phone)
}