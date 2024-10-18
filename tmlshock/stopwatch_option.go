package tmlshock

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/mdouchement/tmlshock/v2/termdraw"
)

type (
	StopwatchOption func(*sOption)

	sOption struct {
		color           lipgloss.TerminalColor
		colonColor      lipgloss.TerminalColor
		backgroundColor lipgloss.TerminalColor
		showDecisecond  bool
		showHour        bool
	}
)

func sOptionDefaults() *sOption {
	return &sOption{
		color:           termdraw.FlagColor("red"),
		colonColor:      termdraw.FlagColorUndefined,
		backgroundColor: termdraw.FlagColorUndefined,
		showHour:        true,
		showDecisecond:  false,
	}
}

func WithStopwatchColor(color string) StopwatchOption {
	return func(o *sOption) {
		o.color = termdraw.FlagColor(color)
	}
}

func WithStopwatchColonColor(color string) StopwatchOption {
	return func(o *sOption) {
		o.colonColor = termdraw.FlagColor(color)
	}
}

func WithStopwatchBackgroundColor(color string) StopwatchOption {
	return func(o *sOption) {
		o.backgroundColor = termdraw.FlagColor(color)
	}
}

func WithStopwatchShowHour(b bool) StopwatchOption {
	return func(o *sOption) {
		o.showHour = b
	}
}

func WithStopwatchShowDecisecond(b bool) StopwatchOption {
	return func(o *sOption) {
		o.showDecisecond = b
	}
}
