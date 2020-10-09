package simple

import (
	"fmt"
	"testing"
)

func Test_reverseString(t *testing.T) {
	bytes := reverseString([]byte("12345"))
	fmt.Println(string(bytes))
}
