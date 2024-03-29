package main

import (
	"fmt"

	installation "installer.io/installation/src/Installation"
	yamlReader "installer.io/installation/src/configuration"
)

func main() {
	config := yamlReader.ReadFunction()
	for tool, available := range config.Tools {
		switch tool {
		case "Subfinder":
			if available {
				fmt.Println("[+]Subfinder is available.Start Downloading it")
				installation.GoInstall("github.com/projectdiscovery/subfinder/v2/cmd/subfinder@latest")
			} else {
				fmt.Println("[+]Subfinder is not available")
			}
		case "Nuclei":
			if available {
				fmt.Println("[+]Nuclei is available.Start Downloading it")
				installation.GoInstall("github.com/projectdiscovery/nuclei/v2/cmd/nuclei@latest")
			} else {
				fmt.Println("[+]Nuclei is not available")
			}
		case "Dnsx":
			if available {
				fmt.Println("[+]Dnsx is available.Start Downloading it")
				installation.GoInstall("github.com/projectdiscovery/dnsx/cmd/dnsx@latest")
			} else {
				fmt.Println("[+]Dnsx is not available")
			}

		case "MassDns":
			if available {
				fmt.Println("[+]MassDns available.Start Downloading it")
				installation.Clone("https://github.com/blechschmidt/massdns.git", "tools")
			} else {
				fmt.Println("[+]MassDns is not available")
			}

		case "Dirsearch":
			if available {
				fmt.Println("[+]Dirsearch is available.Start Downloading it")
				installation.Clone("https://github.com/maurosoria/dirsearch.git", "tools")
			} else {
				fmt.Println("[+]Dirsearch is not available")
			}

		case "ffuf":
			if available {
				fmt.Println("[+]ffuf is available.Start Downloading it")
				installation.GoInstall("github.com/ffuf/ffuf/v2@latest")
			} else {
				fmt.Println("[+]ffuf is not available")
			}

		case "katna":
			if available {
				fmt.Println("[+]katna is available.Start Downloading it")
				installation.GoInstall("github.com/projectdiscovery/katana/cmd/katana@latest")
			} else {
				fmt.Println("[+]katna is not available")
			}
		case "Waymore":
			if available {
				fmt.Println("[+] Waymore available.Start Downloading it")
				installation.Clone("https://github.com/xnl-h4ck3r/waymore.git", "tools")
			}
		case "Waybackurl":
			if available {
				fmt.Println("[+] WybackUrl is available Start Downloading it")
				installation.GoInstall("github.com/tomnomnom/waybackurls@latest")
			} else {
				fmt.Println("[+] Waybackurl is not available")

			}

		case "amass":
			if available {
				fmt.Println("[+]Amass is available Start Downloading it")
				installation.GoInstall("github.com/owasp-amass/amass/v4/...@master")
			} else {
				fmt.Println("[+]Amass is not available")
			}
		case "httpx":
			if available {
				fmt.Println("[+]httpx is available Start Downloading it")
				installation.GoInstall("github.com/projectdiscovery/httpx/cmd/httpx@latest")
			} else {
				fmt.Println("[+]httpx is not available")
			}

		case "gua":
			if available {
				fmt.Println("[+]gua is available Start Downloading it")
				installation.GoInstall("github.com/lc/gau/v2/cmd/gau@latest")
			} else {
				fmt.Println("[+]gua is not available")
			}

		case "x8":
			if available {
				fmt.Println("[+]x8 is available Start Downloading it")
				installation.Clone("https://github.com/Sh1Yo/x8.git", "tools")
			} else {
				fmt.Println("[+]x8 is not available")
			}

		case "Dalfox":
			if available {
				fmt.Println("[+]Dalfox is available start Downloading it")
				installation.GoInstall("github.com/hahwul/dalfox/v2@latest")
			} else {
				fmt.Println("[+]Dalfox is not available")
			}
		case "anew":
			if available {
				fmt.Println("[+]anew is available start Downloading it")
				installation.GoInstall("github.com/tomnomnom/anew@latest")
			} else {
				fmt.Println("[+]anew is not available")
			}
		}
		if config.WordLists["SecList"] {
			fmt.Printf("[+] Clone SecList\n\n")
			installation.Clone("https://github.com/danielmiessler/SecLists", "wordlist")
		} else {
			fmt.Println("[+]SecList is not available")
		}
		if config.WordLists["FuzzDB"] {
			fmt.Printf("[+] Clone Fuzzdb Wordlsit\n\n")
			installation.Clone("https://github.com/fuzzdb-project/fuzzdb", "wordlist")
		} else {
			fmt.Println("[+]Fuzzdb is not available")
		}
	}
}
