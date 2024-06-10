/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/spf13/cobra"
)

var (
	ByteFlag bool
	LineFlag bool
	WordFlag bool
	CharFlag bool
)

const BUFFER_SIZE = 64 * 1024

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ccwc",
	Short: "ccwc - a command which counts bytes, lines, words & characters",
	Long:  `ccwc - a command which counts bytes, lines, words & characters either from a file or from input`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		file := os.Stdin
		fileName := ""
		if len(args) > 0 {
			fileName = args[0]
			f, err := os.Open(fileName)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			defer f.Close()
			file = f
		}

		if cmd.Flags().NFlag() == 0 {
			ByteFlag = true
			LineFlag = true
			WordFlag = true
		}

		output := ""
		byteCount := 0
		lineCount := 0
		wordCount := 0
		charCount := 0

		buffer := make([]byte, BUFFER_SIZE)
		r := bufio.NewReader(file)

		currentBufferStartedWithSpace := false
		lastBufferEndedWithSpace := true
		for {
			bytesRead, err := r.Read(buffer)
			if err == io.EOF {
				break
			}

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			byteCount += bytesRead
			lineCount += bytes.Count(buffer[:bytesRead], []byte("\n"))
			charCount += bytes.Count(buffer[:bytesRead], []byte("")) - 1

			wordCount += len(bytes.Fields(buffer[:bytesRead]))
			currentBufferStartedWithSpace = unicode.IsSpace(rune(buffer[0]))
			if !currentBufferStartedWithSpace && !lastBufferEndedWithSpace {
				wordCount -= 1
			}
			lastBufferEndedWithSpace = unicode.IsSpace(rune(buffer[bytesRead-1]))
		}

		if LineFlag {
			output += " " + strconv.Itoa(lineCount)
		}

		if WordFlag {
			output += " " + strconv.Itoa(wordCount)
		}

		if ByteFlag {
			output += " " + strconv.Itoa(byteCount)
		}

		if CharFlag {
			output += " " + strconv.Itoa(charCount)
		}

		output += " " + fileName
		output = strings.TrimSpace(output)
		fmt.Println(output)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().BoolVarP(&ByteFlag, "byte", "c", false, "count bytes")
	rootCmd.PersistentFlags().BoolVarP(&LineFlag, "line", "l", false, "count lines")
	rootCmd.PersistentFlags().BoolVarP(&WordFlag, "word", "w", false, "count words")
	rootCmd.PersistentFlags().BoolVarP(&CharFlag, "char", "m", false, "count characters")
}
