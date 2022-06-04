package caesarchipher_test

import (
	"fmt"

	"github.com/steveanlorn/learning-go/testing/example_test/caesarchipher"
)

func ExampleEncrypt() {
	chiperText := caesarchipher.Encrypt("zebra", 2)
	fmt.Println(chiperText)
	// Output: bgdtc
}

func ExampleDecrypt() {
	chiperText := caesarchipher.Decrypt("bgdtc", 2)
	fmt.Println(chiperText)
	// Output: zebra
}
