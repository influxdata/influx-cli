package secret

import (
	"os"
	"strings"

	"github.com/tcnksm/go-input"
)

// vt100EscapeCodes
var (
	KeyEscape = byte(27)
	KeyReset  = []byte{KeyEscape, '[', '0', 'm'}
	ColorCyan = []byte{KeyEscape, '[', '3', '6', 'm'}
)

func promptWithColor(s string, color []byte) []byte {
	bb := append(color, []byte(s)...)
	return append(bb, KeyReset...)
}

func getSecret(ui *input.UI) (secret string) {
	var err error
	query := string(promptWithColor("Please type your secret", ColorCyan))
	for {
		secret, err = ui.Ask(query, &input.Options{
			Required:  true,
			HideOrder: true,
			Hide:      true,
			Mask:      false,
		})
		switch err {
		case input.ErrInterrupted:
			os.Exit(1)
		default:
			if secret = strings.TrimSpace(secret); secret == "" {
				continue
			}
		}
		break
	}
	return secret
}
