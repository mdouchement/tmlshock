package tmlshock

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/mdouchement/tmlshock/v2/termdraw"
)

type (
	TimerOption func(*tOption)

	tOption struct {
		color           lipgloss.TerminalColor
		colonColor      lipgloss.TerminalColor
		backgroundColor lipgloss.TerminalColor
		showDecisecond  bool
		showHour        bool
	}
)

func tOptionDefaults() *tOption {
	return &tOption{
		color:           termdraw.FlagColor("red"),
		colonColor:      termdraw.FlagColorUndefined,
		backgroundColor: termdraw.FlagColorUndefined,
		showHour:        true,
		showDecisecond:  false,
	}
}

func WithTimerColor(color string) TimerOption {
	return func(o *tOption) {
		o.color = termdraw.FlagColor(color)
	}
}

func WithTimerColonColor(color string) TimerOption {
	return func(o *tOption) {
		o.colonColor = termdraw.FlagColor(color)
	}
}

func WithTimerBackgroundColor(color string) TimerOption {
	return func(o *tOption) {
		o.backgroundColor = termdraw.FlagColor(color)
	}
}

func WithTimerShowHour(b bool) TimerOption {
	return func(o *tOption) {
		o.showHour = b
	}
}

func WithTimerShowDecisecond(b bool) TimerOption {
	return func(o *tOption) {
		o.showDecisecond = b
	}
}
