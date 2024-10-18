package stopwatch

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mdouchement/tmlshock/v2/tmlshock"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	var hour, decisecond bool
	var color, colonColor string

	c := &cobra.Command{
		Use:     "stopwatch",
		Aliases: []string{"s"},
		Short:   "Start a stopwatch",
		Example: "tmlshock stopwatch [command options]",
		Args:    cobra.NoArgs,
		RunE: func(_ *cobra.Command, args []string) error {
			var options []tmlshock.StopwatchOption
			options = append(options, tmlshock.WithStopwatchColor(color))
			options = append(options, tmlshock.WithStopwatchShowDecisecond(decisecond))
			options = append(options, tmlshock.WithStopwatchShowHour(!hour))
			if colonColor != "" {
				options = append(options, tmlshock.WithStopwatchColonColor(colonColor))
			}

			stopwatch := tmlshock.NewStopwatch(options...)
			p := tea.NewProgram(stopwatch, tea.WithAltScreen())
			_, err := p.Run()
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
