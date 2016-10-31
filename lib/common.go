package lib

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func YesOrNo(message string) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(message + " [yN]")
	str, _ := reader.ReadString('\n')
	lower := strings.ToLower(str)
	return strings.HasPrefix(lower, "y")
}
