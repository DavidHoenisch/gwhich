/*
Copyright © 2025 DavidHoenisch

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// genConfigCmd represents the genConfig command
var genConfigCmd = &cobra.Command{
	Use:   "genConfig",
	Short: "generates a config file in the XDG home folder",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		configHome, _ := os.UserConfigDir()
		configFilePath := fmt.Sprintf("%s/%s",
			configHome, "gwhich")

		_, err := os.Stat(configFilePath)
		if errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(configFilePath, 0744)
			if err != nil {
				log.Println(err)
			}
		}

		_, err = os.Stat(fmt.Sprint(configFilePath, "/config.toml"))
		if errors.Is(err, os.ErrNotExist) {
			_, err := os.Create(fmt.Sprint(configFilePath, "/config.toml"))
			if err != nil {
				log.Println(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(genConfigCmd)
}
