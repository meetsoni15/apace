package apace

import (
	"fmt"
	"math/rand"
)

// Generate IP address
func RandomIPv4() string {
	// Generate four random integers between 0 and 255
	octet1 := rand.Intn(256)
	octet2 := rand.Intn(256)
	octet3 := rand.Intn(256)
	octet4 := rand.Intn(256)

	// Combine the four integers into a dotted IPv4 address
	return fmt.Sprintf("%d.%d.%d.%d", octet1, octet2, octet3, octet4)
}
