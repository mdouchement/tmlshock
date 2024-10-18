package main

import (
	"fmt"
	"os"

	"github.com/mdouchement/tmlshock/v2/cmd"
	"github.com/mdouchement/tmlshock/v2/cmd/clock"
	"github.com/mdouchement/tmlshock/v2/cmd/stopwatch"
	"github.com/mdouchement/tmlshock/v2/cmd/timer"
	"github.com/spf13/cobra"
)

func main() {
	c := &cobra.Command{
		Use:     "tmlshock",
		Short:   "A terminal ttl clock timer and stopwatch build by golang",
		Example: "tmlshock [command] [option]",
		Version: cmd.Version(),
	}
	c.AddCommand(clock.Command())
	c.AddCommand(stopwatch.Command())
	c.AddCommand(timer.Command())
	c.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Version tmlshock",
		Args:  cobra.NoArgs,
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Println(c.Version)
		},
	})

	if err := c.Execute(); err != nil {
		os.Exit(1)
	}
}
