// MIT License
//
// Copyright (c) 2016 Alan Chen
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package krait

import "github.com/spf13/cobra"

// NewConsoleCommand returns a command to enter an interactive console
func NewConsoleCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "console",
		Short: "Enter an interactive console",
		Long:  `Enter an interactive console`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		Run: enterConsole,
	}
}

func enterConsole(cmd *cobra.Command, args []string) {
	shell := New()
	commands := cmd.Root().Commands()
	newRootCommand := &cobra.Command{
		Use: "root",
	}

	for _, c := range commands {
		if cmd != c {
			newRootCommand.AddCommand(c)
		}
	}

	newRootCommand.SetUsageTemplate(`{{ if .HasAvailableSubCommands}}Available Commands:{{range .Commands}}{{if .IsAvailableCommand}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}
`)
	newRootCommand.AddCommand(exitFunc(shell))
	newRootCommand.AddCommand(clearFunc(shell))

	shell.Register(newRootCommand)
	shell.Start()
}
