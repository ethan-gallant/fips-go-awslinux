package main

import (
	"crypto/boring"
	_ "crypto/tls/fipsonly" // This is the wrong way, it will be available regardless of whether FIPS is enabled
	"fmt"
)

func main() {
	// This is the right way to check if FIPS is enabled
	if boring.Enabled() {
		fmt.Println("FIPS is enabled")
	} else {
		fmt.Println("FIPS NOT ENABLED")
	}
}
