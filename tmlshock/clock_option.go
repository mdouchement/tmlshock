package tmlshock

import (
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/mdouchement/tmlshock/v2/termdraw"
)

type (
	ClockOption func(*cOption)

	cOption struct {
		timezone        *time.Location
		color           lipgloss.TerminalColor
		colonColor      lipgloss.TerminalColor
		backgroundColor lipgloss.TerminalColor
		hour12          bool
		showSecond      bool
		showDate        bool
		dateLayout      string
	}
)

func cOptionDefaults() *cOption {
	return &cOption{
		timezone:        time.Now().Location(),
		color:           termdraw.FlagColor("red"),
		colonColor:      termdraw.FlagColorUndefined,
		backgroundColor: termdraw.FlagColorUndefined,
		showSecond:      true,
		showDate:        false,
		dateLayout:      "2006/01/02",
	}
}

func WithClockTimezone(tz *time.Location) ClockOption {
	return func(o *cOption) {
		o.timezone = tz
	}
}

func WithClockColor(color string) ClockOption {
	return func(o *cOption) {
		o.color = termdraw.FlagColor(color)
	}
}

func WithClockColonColor(color string) ClockOption {
	return func(o *cOption) {
		o.colonColor = termdraw.FlagColor(color)
	}
}

func WithClockBackgroundColor(color string) ClockOption {
	return func(o *cOption) {
		o.backgroundColor = termdraw.FlagColor(color)
	}
}

func WithClockShowSecond(b bool) ClockOption {
	return func(o *cOption) {
		o.showSecond = b
	}
}

func WithClockShowDate(b bool) ClockOption {
	return func(o *cOption) {
		o.showDate = b
	}
}

func WithClockDateLayout(layout string) ClockOption {
	return func(o *cOption) {
		o.dateLayout = layout
	}
}

func WithClock12hFormat(b bool) ClockOption {
	return func(o *cOption) {
		o.hour12 = b
	}
}
