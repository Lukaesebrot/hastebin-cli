package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/Lukaesebrot/hastebin-cli/config"
	"github.com/Lukaesebrot/hastebin-cli/service"
	"github.com/atotto/clipboard"
	"github.com/spf13/viper"
)

func main() {
	// Initialize the current configuration file
	err := config.Initialize()
	if err != nil {
		panic(err)
	}

	// Read from stdin
	reader := bufio.NewReader(os.Stdin)
	var builder strings.Builder
	for {
		byt, err := reader.ReadByte()
		if err != nil {
			if err != io.EOF {
				panic(err)
			}
			break
		}
		builder.WriteByte(byt)
	}

	// Paste the read string to the hastebin instance
	pasteURL, err := service.Paste(builder.String(), viper.GetString("instance"))
	if err != nil {
		panic(err)
	}

	// Copy the paste URL if auto-clip is enabled
	if viper.GetBool("autoClip") {
		err = clipboard.WriteAll(pasteURL)
		if err != nil {
			fmt.Println(err)
		}
	}

	// Print out the paste URL
	fmt.Println(pasteURL)
}
