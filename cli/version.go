package cli

import (
	"fmt"
)

const Version string = "1.4.1"

func PrintVersion() error {
	_, err := fmt.Printf("Dogestry %s\n", Version)
	return err
}

func (cli *DogestryCli) CmdVersion(args ...string) error {
	return PrintVersion()
}
