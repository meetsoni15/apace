package apace

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"math/rand"
)

// default character set
const (
	chars         = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	defaultDigits = 999999
)

// Seeders seed
// The seed is a value which initializes the random number generator.
// Random number generators produce values by performing some operation on a previous value.
// When the algorithm starts, the seed is the initial value on which the generator operates.
// The most important and difficult part of the generators is to provide a seed that is close to a truly random number.
func init() {
	var b [8]byte
	_, err := crypto_rand.Read(b[:])
	if err != nil {
		panic("cannot seed math/rand package with cryptographically secure random number generator")
	}
	rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
}

// Generate random string
// It takes 2 arguments one is mandatory that is string
// Another one is optional character set
// If want to generate string with particular character set then supply on 2nd argument
func RandomString(l int, charSet ...string) (string, error) {
	var charPool = chars
	if len(charSet) > 0 {
		if charSet[0] != "" {
			charPool = charSet[0]
		}
	}
	bytes := make([]byte, l)

	_, err := crypto_rand.Read(bytes)
	if err != nil {
		return "", err
	}

	for i, b := range bytes {
		bytes[i] = charPool[b%byte(len(charPool))]
	}
	return string(bytes), nil
}

// RandomArray
// It will generate random array of type string and int/int64
// It required 2 parameter length of array, provide character set
func RandomArray(l int, digits any) []any {
	arr := make([]any, l)

	switch digits.(type) {
	case string:
		for i := 0; i < l; i++ {
			str, _ := RandomString(l, digits.(string))
			arr[i] = str
		}
	case int64:
		for i := 0; i < l; i++ {
			arr[i] = rand.Int63n(digits.(int64))
		}
	case int:
		for i := 0; i < l; i++ {
			arr[i] = rand.Intn(digits.(int))
		}
	default:
		for i := 0; i < l; i++ {
			arr[i] = rand.Intn(defaultDigits)
		}
	}

	return arr
}

// RandInt
func RandomInt(min, max int) int {
	return min + rand.Intn(max-min)
}
