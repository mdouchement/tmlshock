package timer

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mdouchement/tmlshock/v2/tmlshock"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	var hour, decisecond bool
	var color, colonColor string

	c := &cobra.Command{
		Use:     "timer",
		Aliases: []string{"t"},
		Short:   "Start a timer",
		Example: "tmlshock timer [command options] duration\ntmlshock timer 4h42m10s\ntmlshock timer 2m\ntmlshock timer 1h4s",
		Args:    cobra.ExactArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			var options []tmlshock.TimerOption
			options = append(options, tmlshock.WithTimerColor(color))
			options = append(options, tmlshock.WithTimerShowDecisecond(decisecond))
			options = append(options, tmlshock.WithTimerShowHour(!hour))
			if colonColor != "" {
				options = append(options, tmlshock.WithTimerColonColor(colonColor))
			}

			duration, err := time.ParseDuration(args[0])
			if err != nil {
				return fmt.Errorf("could not parse timer duration: %w", err)
			}

			timer := tmlshock.NewTimer(duration, options...)
			p := tea.NewProgram(timer, tea.WithAltScreen())
			_, err = p.Run()
			return err
		},
	}

	// c.Flags().StringVarP(&backgroundColor, "background-color", "b", "none", "Set the background color (https://github.com/mdouchement/tmlshock?tab=readme-ov-file#color)")
	c.Flags().StringVarP(&color, "color", "c", "FF2400", "Set the string color (https://github.com/mdouchement/tmlshock?tab=readme-ov-file#color)")
	c.Flags().StringVarP(&colonColor, "colon-color", "C", "", "Set the colon color (https://github.com/mdouchement/tmlshock?tab=readme-ov-file#color)")
	c.Flags().BoolVarP(&hour, "hour", "H", false, "Disable hour")
	c.Flags().BoolVarP(&decisecond, "decisecond", "d", false, "Show decisecond")
	return c
}
