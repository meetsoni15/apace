## Apace - random string/int/array generator utility for the GoLang

This package contains a random string, array and int generator. It is a wrapper around the crypto/rand package. For array currently it supports int, int64 and string.

### Usage
```
go get github.com/meetsoni15/apace
```

```go
package main

import (
	"fmt"
	"log"

	"github.com/meetsoni15/apace"
)

func main() {
	// String generator with default character set
	for i := 0; i < 5; i++ {
		str, err := apace.RandomString(15)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(str)
	}

    // String generator with custom default character set
	for i := 0; i < 5; i++ {
		str, err := apace.RandomString(15,"ABCDEFGHIJKLMNOPQRSTUVWXYZ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(str)
	}

	// Int generator
	for i := 0; i < 5; i++ {
		num := apace.RandomInt(5, 100000)
		fmt.Println(num)
	}

	// Int Array generator
	for i := 0; i < 5; i++ {
		// Here we have provided argument
		// Length = 5 and digits in range upto 9999
		arr := apace.RandomArray(5, 9999)
		fmt.Println(arr)
	}

	///String Array generator
	for i := 0; i < 5; i++ {
		// Here we have provided argument
		// Length = 5 and digits in range upto 9999
		arr := apace.RandomArray(5, "AbcdDMoPP")
		fmt.Println(arr)
	}

}

// Output:
/*
 
 //String generator with default character set  
 
M3G7fTFykemo76y
RQeR9pxm5Wzfwi7
59V861ML0GVoMWW
RiRnmiAlHO1BCvc
9yEu3mDEyxy66GE
 
 //String generator with custom character set  
 
TGHFQPDXNARBXSG
ZXLMRJQXHWIPPKG
NPENIDSDNTXDSWM
DFWZLWNIALPBKSE
RMIKJJZWTBHLSWF
 
 //Int generator  
 
41546
67653
28204
45842
58885
 
 //Int Array generator  
 
[1170 3486 1747 432 7920]
[8349 9941 7560 2167 7737]
[6662 5560 8827 708 7551]
[5014 8662 6734 4369 4844]
[352 4378 7547 3965 6602]
 
 //String Array generator  
 
[dAdbP bbdMD PdbcM oDAoP bdDMP]
[PPAPc cdobP oPoPc cPdbb AMMbc]
[cAPcP AbPoc obDod bDAPD DoPPP]
[PooMP PPDcd APMbD PAMbD dcDPD]
[DPoPd ADPdb bbDPd cAMAo cAPPP]
*/

```

### Benchmark
```go
go test -bench=.
goos: darwin
goarch: arm64
pkg: github.com/meetsoni15/apace
BenchmarkRandomString-8          1860854               643.1 ns/op            64 B/op          2 allocs/op
BenchmarkRandomInt-8            67795495                17.66 ns/op            0 B/op          0 allocs/op
BenchmarkRandomArray-8            475621              2388 ns/op            2589 B/op        100 allocs/op
PASS
ok      github.com/meetsoni15/apace     4.455s
```
