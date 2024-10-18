package clock

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mdouchement/tmlshock/v2/tmlshock"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	var date, hour, second bool
	var color, colonColor, layout string

	c := &cobra.Command{
		Use:     "clock",
		Aliases: []string{"c"},
		Short:   "Start a clock",
		Example: "tmlshock clock [command options]",
		Args:    cobra.NoArgs,
		RunE: func(_ *cobra.Command, _ []string) error {
			var options []tmlshock.ClockOption
			options = append(options, tmlshock.WithClockShowDate(date))
			options = append(options, tmlshock.WithClock12hFormat(hour))
			options = append(options, tmlshock.WithClockShowSecond(!second))
			options = append(options, tmlshock.WithClockDateLayout(layout))
			options = append(options, tmlshock.WithClockColor(color))
			// options = append(options, tmlshock.WithClockBackgroundColor(backgroundColor))
			if colonColor != "" {
				options = append(options, tmlshock.WithClockColonColor(colonColor))
			}

			clock := tmlshock.NewClock(options...)
			p := tea.NewProgram(clock, tea.WithAltScreen())
			_, err := p.Run()
			return err
		},
	}

	// c.Flags().StringVarP(&backgroundColor, "background-color", "b", "none", "Set the background color (https://github.com/mdouchement/tmlshock?tab=readme-ov-file#color)")
	c.Flags().StringVarP(&color, "color", "c", "FF2400", "Set the string color (https://github.com/mdouchement/tmlshock?tab=readme-ov-file#color)")
	c.Flags().StringVarP(&colonColor, "colon-color", "C", "", "Set the colon color (https://github.com/mdouchement/tmlshock?tab=readme-ov-file#color)")
	c.Flags().StringVarP(&layout, "date-layout", "l", "2006/01/02", "Set the clock date (https://pkg.go.dev/time#pkg-constants)")
	c.Flags().BoolVarP(&date, "date", "d", false, "Set the clock with date")
	c.Flags().BoolVarP(&hour, "hour12", "H", false, "Set the clock in 12h display")
	c.Flags().BoolVarP(&second, "no-second", "s", false, "Set the clock without second")
	return c
}
