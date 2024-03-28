package main

import (
	"fmt"

	installation "installer.io/installation/src"
)

func main() {
	// fmt.Printf("[+] Clone SecList\n\n")
	// installation.Clone("https://github.com/danielmiessler/SecLists", "wordlist")

	// fmt.Printf("[+] Clone Fuzzdb Wordlsit\n\n")
	// installation.Clone("https://github.com/fuzzdb-project/fuzzdb", "wordlist")

	fmt.Printf("[+] Installing Waymore")
	installation.Clone("https://github.com/xnl-h4ck3r/waymore", "tools")

}
