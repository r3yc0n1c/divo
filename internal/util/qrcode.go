package util

import (
	"fmt"
	"os"

	qrterminal "github.com/mdp/qrterminal/v3"
)

func PrintQRCode(url string) {
	config := qrterminal.Config{
		HalfBlocks: true,
		Level:      qrterminal.L,
		Writer:     os.Stdout,
	}

	qrterminal.GenerateWithConfig(url, config)
	
	fmt.Print("\n[Q] Scan the above QR code to open the link on your phone...\n\n")
}
