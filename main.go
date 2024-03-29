package main

import (
	"fmt"
	"log"

	installation "installer.io/installation/src/Installation"
	yamlReader "installer.io/installation/src/configuration"
)

func main() {
	config := yamlReader.ReadFunction()

	if config.WordLists["SecList"] {
		fmt.Println("[+] Cloning SecList...")
		if err := installation.Clone("https://github.com/danielmiessler/SecLists", "wordlist"); err != nil {
			log.Fatalf("Error cloning SecList: %v", err)
		}
	} else {
		fmt.Println("[+] SecList is not available")
	}

	if config.WordLists["FuzzDB"] {
		fmt.Println("[+] Cloning FuzzDB Wordlist...")
		if err := installation.Clone("https://github.com/fuzzdb-project/fuzzdb", "wordlist"); err != nil {
			log.Fatalf("Error cloning FuzzDB: %v", err)
		}
	} else {
		fmt.Println("[+] FuzzDB is not available")
	}

	for tool, available := range config.Tools {
		switch tool {
		case "Subfinder":
			if available {
				fmt.Println("[+] Subfinder is available. Start Downloading it...")
				if err := installation.GoInstall("github.com/projectdiscovery/subfinder/v2/cmd/subfinder@latest"); err != nil {
					log.Fatalf("Error installing Subfinder: %v", err)
				}
			} else {
				fmt.Println("[+] Subfinder is not available")
			}
		case "Nuclei":
			if available {
				fmt.Println("[+] Nuclei is available. Start Downloading it...")
				if err := installation.GoInstall("github.com/projectdiscovery/nuclei/v2/cmd/nuclei@latest"); err != nil {
					log.Fatalf("Error installing Nuclei: %v", err)
				}
			} else {
				fmt.Println("[+] Nuclei is not available")
			}
		case "Dnsx":
			if available {
				fmt.Println("[+] Dnsx is available. Start Downloading it...")
				if err := installation.GoInstall("github.com/projectdiscovery/dnsx/cmd/dnsx@latest"); err != nil {
					log.Fatalf("Error installing Dnsx: %v", err)
				}
			} else {
				fmt.Println("[+] Dnsx is not available")
			}
		case "MassDns":
			if available {
				fmt.Println("[+] MassDns is available. Start Downloading it...")
				if err := installation.Clone("https://github.com/blechschmidt/massdns.git", "tools"); err != nil {
					log.Fatalf("Error cloning MassDns: %v", err)
				}
			} else {
				fmt.Println("[+] MassDns is not available")
			}
		case "Dirsearch":
			if available {
				fmt.Println("[+] Dirsearch is available. Start Downloading it...")
				if err := installation.Clone("https://github.com/maurosoria/dirsearch.git", "tools"); err != nil {
					log.Fatalf("Error cloning Dirsearch: %v", err)
				}
			} else {
				fmt.Println("[+] Dirsearch is not available")
			}
		case "ffuf":
			if available {
				fmt.Println("[+] ffuf is available. Start Downloading it...")
				if err := installation.GoInstall("github.com/ffuf/ffuf/v2@latest"); err != nil {
					log.Fatalf("Error installing ffuf: %v", err)
				}
			} else {
				fmt.Println("[+] ffuf is not available")
			}
		case "katna":
			if available {
				fmt.Println("[+] katna is available. Start Downloading it...")
				if err := installation.GoInstall("github.com/projectdiscovery/katana/cmd/katana@latest"); err != nil {
					log.Fatalf("Error installing katna: %v", err)
				}
			} else {
				fmt.Println("[+] katna is not available")
			}
		case "Waymore":
			if available {
				fmt.Println("[+] Waymore is available. Start Downloading it...")
				if err := installation.Clone("https://github.com/xnl-h4ck3r/waymore.git", "tools"); err != nil {
					log.Fatalf("Error cloning Waymore: %v", err)
				}
			} else {
				fmt.Println("[+] Waymore is not available")
			}
		case "Waybackurl":
			if available {
				fmt.Println("[+] Waybackurl is available. Start Downloading it...")
				if err := installation.GoInstall("github.com/tomnomnom/waybackurls@latest"); err != nil {
					log.Fatalf("Error installing Waybackurl: %v", err)
				}
			} else {
				fmt.Println("[+] Waybackurl is not available")
			}
		case "amass":
			if available {
				fmt.Println("[+] Amass is available. Start Downloading it...")
				if err := installation.GoInstall("github.com/owasp-amass/amass/v4/...@master"); err != nil {
					log.Fatalf("Error installing amass: %v", err)
				}
			} else {
				fmt.Println("[+] Amass is not available")
			}
		case "httpx":
			if available {
				fmt.Println("[+] httpx is available. Start Downloading it...")
				if err := installation.GoInstall("github.com/projectdiscovery/httpx/cmd/httpx@latest"); err != nil {
					log.Fatalf("Error installing httpx: %v", err)
				}
			} else {
				fmt.Println("[+] httpx is not available")
			}
		case "gua":
			if available {
				fmt.Println("[+] gua is available. Start Downloading it...")
				if err := installation.GoInstall("github.com/lc/gau/v2/cmd/gau@latest"); err != nil {
					log.Fatalf("Error installing gua: %v", err)
				}
			} else {
				fmt.Println("[+] gua is not available")
			}
		case "x8":
			if available {
				fmt.Println("[+] x8 is available. Start Downloading it...")
				if err := installation.Clone("https://github.com/Sh1Yo/x8.git", "tools"); err != nil {
					log.Fatalf("Error cloning x8: %v", err)
				}
			} else {
				fmt.Println("[+] x8 is not available")
			}
		case "Dalfox":
			if available {
				fmt.Println("[+] Dalfox is available. Start Downloading it...")
				if err := installation.GoInstall("github.com/hahwul/dalfox/v2@latest"); err != nil {
					log.Fatalf("Error installing Dalfox: %v", err)
				}
			} else {
				fmt.Println("[+] Dalfox is not available")
			}
		case "anew":
			if available {
				fmt.Println("[+] anew is available. Start Downloading it...")
				if err := installation.GoInstall("github.com/tomnomnom/anew@latest"); err != nil {
					log.Fatalf("Error installing anew: %v", err)
				}
			} else {
				fmt.Println("[+] anew is not available")
			}
		default:
			fmt.Printf("[+] Unknown tool: %s\n", tool)
		}
	}
}
