package tmlshock

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mdouchement/tmlshock/v2/termdraw"
)

type Timer struct {
	opts     *tOption
	font     *termdraw.ClockFont
	width    int
	height   int
	hour     bool
	ref      time.Time
	duration time.Duration
}

func NewTimer(d time.Duration, options ...TimerOption) *Timer {
	t := &Timer{
		opts:     tOptionDefaults(),
		hour:     d >= time.Hour,
		duration: d,
	}

	for _, opt := range options {
		opt(t.opts)
	}

	if t.opts.colonColor == termdraw.FlagColorUndefined {
		t.opts.colonColor = t.opts.color
	}

	t.font = termdraw.NewFont(t.opts.backgroundColor, t.opts.color, t.opts.colonColor)
	return t
}

func (t *Timer) Init() tea.Cmd {
	t.ref = time.Now().Add(t.duration)
	if !t.opts.showDecisecond {
		t.duration -= time.Second // Compensate ticker cold start
	}
	return t.tick()
}

func (t *Timer) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		t.width = msg.Width
		t.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "Q":
			return t, tea.Quit
		}
	case TickMsg:
		t.duration = t.ref.Sub(msg.Time)
		return t, t.tick()
	}

	return t, nil
}

func (t *Timer) View() string {
	if t.width == 0 {
		return "loading..."
	}

	duration := FormatDuration(t.duration, t.opts.showHour && t.hour, t.opts.showDecisecond)
	return lipgloss.Place(
		t.width,
		t.height,
		lipgloss.Center,
		lipgloss.Center,
		t.font.Format(duration),
	)
}

func (t *Timer) tick() tea.Cmd {
	d := time.Second
	if t.opts.showDecisecond {
		d = 100 * time.Millisecond
	}

	return tea.Tick(d, func(t time.Time) tea.Msg {
		return TickMsg{
			Time: t,
		}
	})
}
