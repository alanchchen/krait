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

func exitFunc(s *Shell) *cobra.Command {
	return &cobra.Command{
		Use:   "exit",
		Short: "Exit the shell",
		Long:  `Exit the shell`,
		Run: func(cmd *cobra.Command, args []string) {
			s.Stop()
		},
	}
}

func clearFunc(s *Shell) *cobra.Command {
	return &cobra.Command{
		Use:   "clear",
		Short: "Clear screen",
		Long:  `Clear screen`,
		Run: func(cmd *cobra.Command, args []string) {
			s.ClearScreen()
		},
	}
}

func addDefaultFuncs(s *Shell) {
	s.RegisterInterrupt(interruptFunc(s))
}

func interruptFunc(s *Shell) *cobra.Command {
	interruptCmd := &cobra.Command{
		Use: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			s.interruptCount++
			if s.interruptCount >= 2 {
				return ExitErr("Interrupted")
			}
			return nil
		},
	}
	interruptCmd.SetUsageFunc(func(c *cobra.Command) error {
		return nil
	})
	interruptCmd.SetHelpFunc(func(c *cobra.Command, args []string) {})
	return interruptCmd
}
