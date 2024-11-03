package Int

import (
	"fmt"
	"strconv"
	_ "strings"

	"golang.org/x/example/hello/reverse"
)

// Int returns the decimal reversal of the integer i.
func Int(i int) int {
	i, _ = strconv.Atoi(reverse.String(strconv.Itoa(i)))
	return i
}

func main() {
	fmt.Println("Int main")
}
