package rand_test

import (
	"fmt"

	"github.com/fsmarc/sRand/rand"
)

func ExampleNew() {
	rnd := rand.New(rand.NewSource(1)) //use another int64 as Seed.
	fmt.Println(rnd.Stringn(10))
}

func ExampleDefaultCharset() {
	fmt.Println(rand.DefaultCharset)
	//Output:
	//abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789
}

func ExampleNewWithCharset() {
	rnd := rand.NewWithCharset(rand.NewSource(1), "Hello World!") //use another int64 as Seed.
	fmt.Println(rnd.Char())
}

func ExampleBase64UrlCharset() {
	rnd := rand.NewWithCharset(rand.NewSource(1), rand.Base64UrlCharset)
	fmt.Println(rnd.Stringn(11))
}
