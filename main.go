package main

import (
	configuration "installer.io/installation/src/configuration"
)

func main() {
	// fmt.Printf("[+] Clone SecList\n\n")
	// installation.Clone("https://github.com/danielmiessler/SecLists", "wordlist")

	// fmt.Printf("[+] Clone Fuzzdb Wordlsit\n\n")
	// installation.Clone("https://github.com/fuzzdb-project/fuzzdb", "wordlist")

	// fmt.Printf("[+] Installing Waymore")
	// installation.Clone("https://github.com/xnl-h4ck3r/waymore", "tools")

	configuration.ReadFunction()
}
