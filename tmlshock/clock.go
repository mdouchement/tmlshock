package tmlshock

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mdouchement/tmlshock/v2/termdraw"
)

type Clock struct {
	opts   *cOption
	font   *termdraw.ClockFont
	cut    int
	width  int
	height int
	time   time.Time
}

func NewClock(options ...ClockOption) *Clock {
	c := &Clock{
		opts: cOptionDefaults(),
		cut:  5,
		time: time.Now(), // Just to avoid displaying uintialized clock during startup.
	}

	for _, opt := range options {
		opt(c.opts)
	}

	if c.opts.colonColor == termdraw.FlagColorUndefined {
		c.opts.colonColor = c.opts.color
	}

	c.font = termdraw.NewFont(c.opts.backgroundColor, c.opts.color, c.opts.colonColor)
	if c.opts.showSecond {
		c.cut = 8
	}

	return c
}

func (c *Clock) Init() tea.Cmd {
	return c.tick()
}

func (c *Clock) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		c.width = msg.Width
		c.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "Q":
			return c, tea.Quit
		}
	case TickMsg:
		c.time = msg.Time
		return c, c.tick()
	}

	return c, nil
}

func (c *Clock) View() string {
	if c.width == 0 {
		return "loading..."
	}

	formatted := FormatTime(c.time.In(c.opts.timezone), c.opts.hour12)
	formatted = c.font.Format(formatted[:c.cut])
	if c.opts.showDate {
		date := c.time.Format(c.opts.dateLayout)
		formatted = lipgloss.JoinVertical(
			lipgloss.Center,
			formatted,
			"",
			c.font.Text(date),
		)
	}

	return lipgloss.Place(
		c.width,
		c.height,
		lipgloss.Center,
		lipgloss.Center,
		formatted,
	)
}

func (c *Clock) tick() tea.Cmd {
	const minimum = time.Second

	var d time.Duration
	if !c.opts.showSecond {
		// Since seconds are not shown, just refresh the less possible.
		d = time.Duration(60-time.Now().Second()) * time.Second
	}

	if d < minimum {
		d = minimum
	}

	return tea.Tick(d, func(t time.Time) tea.Msg {
		return TickMsg{
			Time: t,
		}
	})
}
