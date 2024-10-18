package tmlshock

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mdouchement/tmlshock/v2/termdraw"
)

type Stopwatch struct {
	opts     *sOption
	font     *termdraw.ClockFont
	width    int
	height   int
	ref      time.Time
	duration time.Duration
	pause    bool
}

func NewStopwatch(options ...StopwatchOption) *Stopwatch {
	s := &Stopwatch{
		opts: sOptionDefaults(),
	}

	for _, opt := range options {
		opt(s.opts)
	}

	if s.opts.colonColor == termdraw.FlagColorUndefined {
		s.opts.colonColor = s.opts.color
	}

	s.font = termdraw.NewFont(s.opts.backgroundColor, s.opts.color, s.opts.colonColor)
	return s
}

func (s *Stopwatch) Init() tea.Cmd {
	s.ref = time.Now()
	return s.tick()
}

func (s *Stopwatch) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		s.width = msg.Width
		s.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "r":
			s.ref = time.Now()
			return s, nil
		case " ": // Space bar
			s.pause = !s.pause
			return s, nil
		case "ctrl+c", "q", "Q":
			return s, tea.Quit
		}
	case TickMsg:
		if !s.pause {
			s.duration = msg.Time.Sub(s.ref)
		}
		return s, s.tick()
	}

	return s, nil
}

func (s *Stopwatch) View() string {
	if s.width == 0 {
		return "loading..."
	}

	duration := FormatDuration(s.duration, s.opts.showHour, s.opts.showDecisecond)
	return lipgloss.Place(
		s.width,
		s.height,
		lipgloss.Center,
		lipgloss.Center,
		s.font.Format(duration),
	)
}

func (s *Stopwatch) tick() tea.Cmd {
	d := time.Second
	if s.opts.showDecisecond {
		d = 100 * time.Millisecond
	}

	return tea.Tick(d, func(t time.Time) tea.Msg {
		return TickMsg{
			Time: t,
		}
	})
}
