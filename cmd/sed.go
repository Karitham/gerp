package cmd

import (
	"bytes"
	"io"
	"os"
	"regexp"

	"github.com/urfave/cli/v2"
)

type SedArgs struct {
	Input   string
	Match   string
	Replace string
	Output  string
}

func (s *SedArgs) Cmd(c *cli.Context) error {
	match, err := regexp.Compile(s.Match)
	if err != nil {
		return err
	}
	replace := []byte(s.Replace)

	input := os.Stdin
	if s.Input != "" {
		input, err = os.Open(s.Input)
		if err != nil {
			return err
		}
	}

	content, err := io.ReadAll(input)
	if err != nil {
		return err
	}

	output := os.Stdout
	if s.Output != "" {
		output, err = os.OpenFile(s.Output, os.O_RDWR, os.ModePerm)
		if err != nil {
			return err
		}
	}

	_, err = io.Copy(output, bytes.NewBuffer(match.ReplaceAll(content, replace)))
	return err
}
