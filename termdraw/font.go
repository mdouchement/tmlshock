package termdraw

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type ClockFont struct {
	charset map[rune]string
	style   lipgloss.Style
}

func NewFont(background, foreground, colon lipgloss.TerminalColor) *ClockFont {
	f := new(ClockFont)

	styleON := lipgloss.NewStyle().
		Background(foreground). // We print spaces so we need to invert bg/fg
		Foreground(background)

	styleOFF := lipgloss.NewStyle().
		Background(background).
		Foreground(foreground)

	styleColon := styleON
	if colon != FlagColorUndefined {
		styleColon = lipgloss.NewStyle().
			Background(colon). // We print spaces so we need to invert bg/fg
			Foreground(background)
	}

	f.style = styleOFF
	f.charset = map[rune]string{
		'0': ClockFontZero(styleOFF, styleON),
		'1': ClockFontOne(styleOFF, styleON),
		'2': ClockFontTwo(styleOFF, styleON),
		'3': ClockFontThree(styleOFF, styleON),
		'4': ClockFontFour(styleOFF, styleON),
		'5': ClockFontFive(styleOFF, styleON),
		'6': ClockFontSix(styleOFF, styleON),
		'7': ClockFontSeven(styleOFF, styleON),
		'8': ClockFontEight(styleOFF, styleON),
		'9': ClockFontNine(styleOFF, styleON),
		':': ClockFontColon(styleOFF, styleColon),
		'.': ClockFontDot(styleOFF, styleColon),
		' ': ClockFontSpace(styleOFF, styleON),
	}
	return f
}

func (f *ClockFont) Format(time string) string {
	var runes []string
	time = strings.Join(strings.Split(time, ""), " ") // insert space between chars

	for _, r := range time {
		f, ok := f.charset[r]
		if !ok {
			continue
		}
		runes = append(runes, f)
	}

	return lipgloss.JoinHorizontal(lipgloss.Left, runes...)
}

func (f *ClockFont) Text(text string) string {
	return f.style.Render(text)
}

func (f *ClockFont) Zero() string {
	return f.charset['0']
}

func (f *ClockFont) One() string {
	return f.charset['1']
}

func (f *ClockFont) Two() string {
	return f.charset['2']
}

func (f *ClockFont) Three() string {
	return f.charset['3']
}

func (f *ClockFont) Four() string {
	return f.charset['4']
}

func (f *ClockFont) Five() string {
	return f.charset['5']
}

func (f *ClockFont) Six() string {
	return f.charset['6']
}

func (f *ClockFont) Seven() string {
	return f.charset['7']
}

func (f *ClockFont) Eight() string {
	return f.charset['8']
}

func (f *ClockFont) Nine() string {
	return f.charset['9']
}

func (f *ClockFont) Colon() string {
	return f.charset[':']
}

func (f *ClockFont) Dot() string {
	return f.charset['.']
}

func (f *ClockFont) Space() string {
	return f.charset[' ']
}

//
//
//
//
//
//

const (
	SpaceSquare = "  "
	SpaceLine   = SpaceSquare + SpaceSquare + SpaceSquare
)

func ClockFontZero(off, on lipgloss.Style) string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		[]string{
			on.Render(SpaceLine),
			on.Render(SpaceSquare) + off.Render(SpaceSquare) + on.Render(SpaceSquare),
			on.Render(SpaceSquare) + off.Render(SpaceSquare) + on.Render(SpaceSquare),
			on.Render(SpaceSquare) + off.Render(SpaceSquare) + on.Render(SpaceSquare),
			on.Render(SpaceLine),
		}...,
	)
}

func ClockFontOne(off, on lipgloss.Style) string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		[]string{
			off.Render(SpaceSquare+SpaceSquare) + on.Render(SpaceSquare),
			off.Render(SpaceSquare+SpaceSquare) + on.Render(SpaceSquare),
			off.Render(SpaceSquare+SpaceSquare) + on.Render(SpaceSquare),
			off.Render(SpaceSquare+SpaceSquare) + on.Render(SpaceSquare),
			off.Render(SpaceSquare+SpaceSquare) + on.Render(SpaceSquare),
		}...,
	)
}

func ClockFontTwo(off, on lipgloss.Style) string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		[]string{
			on.Render(SpaceLine),
			off.Render(SpaceSquare+SpaceSquare) + on.Render(SpaceSquare),
			on.Render(SpaceLine),
			on.Render(SpaceSquare) + off.Render(SpaceSquare+SpaceSquare),
			on.Render(SpaceLine),
		}...,
	)
}

func ClockFontThree(off, on lipgloss.Style) string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		[]string{
			on.Render(SpaceLine),
			off.Render(SpaceSquare+SpaceSquare) + on.Render(SpaceSquare),
			on.Render(SpaceLine),
			off.Render(SpaceSquare+SpaceSquare) + on.Render(SpaceSquare),
			on.Render(SpaceLine),
		}...,
	)
}

func ClockFontFour(off, on lipgloss.Style) string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		[]string{
			on.Render(SpaceSquare) + off.Render(SpaceSquare) + on.Render(SpaceSquare),
			on.Render(SpaceSquare) + off.Render(SpaceSquare) + on.Render(SpaceSquare),
			on.Render(SpaceLine),
			off.Render(SpaceSquare+SpaceSquare) + on.Render(SpaceSquare),
			off.Render(SpaceSquare+SpaceSquare) + on.Render(SpaceSquare),
		}...,
	)
}

func ClockFontFive(off, on lipgloss.Style) string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		[]string{
			on.Render(SpaceLine),
			on.Render(SpaceSquare) + off.Render(SpaceSquare+SpaceSquare),
			on.Render(SpaceLine),
			off.Render(SpaceSquare+SpaceSquare) + on.Render(SpaceSquare),
			on.Render(SpaceLine),
		}...,
	)
}

func ClockFontSix(off, on lipgloss.Style) string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		[]string{
			on.Render(SpaceLine),
			on.Render(SpaceSquare) + off.Render(SpaceSquare+SpaceSquare),
			on.Render(SpaceLine),
			on.Render(SpaceSquare) + off.Render(SpaceSquare) + on.Render(SpaceSquare),
			on.Render(SpaceLine),
		}...,
	)
}

func ClockFontSeven(off, on lipgloss.Style) string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		[]string{
			on.Render(SpaceLine),
			off.Render(SpaceSquare+SpaceSquare) + on.Render(SpaceSquare),
			off.Render(SpaceSquare+SpaceSquare) + on.Render(SpaceSquare),
			off.Render(SpaceSquare+SpaceSquare) + on.Render(SpaceSquare),
			off.Render(SpaceSquare+SpaceSquare) + on.Render(SpaceSquare),
		}...,
	)
}

func ClockFontEight(off, on lipgloss.Style) string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		[]string{
			on.Render(SpaceLine),
			on.Render(SpaceSquare) + off.Render(SpaceSquare) + on.Render(SpaceSquare),
			on.Render(SpaceLine),
			on.Render(SpaceSquare) + off.Render(SpaceSquare) + on.Render(SpaceSquare),
			on.Render(SpaceLine),
		}...,
	)
}

func ClockFontNine(off, on lipgloss.Style) string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		[]string{
			on.Render(SpaceLine),
			on.Render(SpaceSquare) + off.Render(SpaceSquare) + on.Render(SpaceSquare),
			on.Render(SpaceLine),
			off.Render(SpaceSquare+SpaceSquare) + on.Render(SpaceSquare),
			off.Render(SpaceSquare+SpaceSquare) + on.Render(SpaceSquare),
		}...,
	)
}

func ClockFontColon(off, on lipgloss.Style) string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		[]string{
			off.Render(SpaceSquare),
			on.Render(SpaceSquare),
			off.Render(SpaceSquare),
			on.Render(SpaceSquare),
			off.Render(SpaceSquare),
		}...,
	)
}

func ClockFontDot(off, on lipgloss.Style) string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		[]string{
			off.Render(SpaceSquare),
			off.Render(SpaceSquare),
			off.Render(SpaceSquare),
			off.Render(SpaceSquare),
			on.Render(SpaceSquare),
		}...,
	)
}

func ClockFontSpace(off, on lipgloss.Style) string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		[]string{
			off.Render(SpaceSquare),
			off.Render(SpaceSquare),
			off.Render(SpaceSquare),
			off.Render(SpaceSquare),
			off.Render(SpaceSquare),
		}...,
	)
}
