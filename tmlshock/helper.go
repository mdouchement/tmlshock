package tmlshock

import (
	"fmt"
	"strings"
	"time"
)

type TickMsg struct {
	Time time.Time
}

func FormatTime(t time.Time, use12hFormat bool) string {
	hour := t.Hour()
	if use12hFormat {
		hour = hour % 12
		if hour == 0 {
			hour = 12
		}
	}

	return fmt.Sprintf("%02d:%02d:%02d", hour, t.Minute(), t.Second())
}

func FormatDuration(d time.Duration, hour, decisecond bool) string {
	parts := make([]string, 0, 4)

	if hour {
		h := d / time.Hour
		d -= h * time.Hour
		parts = append(parts, fmt.Sprintf("%02d:", h))
	}

	m := d / time.Minute
	d -= m * time.Minute
	parts = append(parts, fmt.Sprintf("%02d:", m))

	s := d / time.Second
	parts = append(parts, fmt.Sprintf("%02d", s))

	if decisecond {
		d = d - s*time.Second
		parts = append(parts, fmt.Sprintf(".%d", d.Milliseconds()/100))
	}

	return strings.Join(parts, "")
}
