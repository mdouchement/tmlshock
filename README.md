This a hard fork of [tmlshock](https://github.com/yorukot/tmlshock) which supports 24-bit (RGB) colors.\
It started by a simple fix which evolves in playing around [bubbletea](https://github.com/charmbracelet/bubbletea) and [lipgloss](https://github.com/charmbracelet/lipgloss).

It happens that bubbletea/lipgloss are easier to use than `github.com/nsf/termbox-go` but with lesser perforrmance when displaying millisecond (some tearing appears).\
As a result in this fork the Timer & Stopwatch only display decisecond which is quick enough for human.

## About tmlshock
Terminal ttl clock, including customizable clock timer and stopwatch

![feature](/.resources/feature.png)

> [!NOTE]
> This image comes from the command `zellij -l .resources/zellij_tmlshock.layout.kdl`

# Contents

- [Install](#install)
- [Setting](#settings)
  * [**Color**](#color)
  * [**Date format**](#date-format)
  * [**12-hour format**](#12-hour-format)
  * [**Timer duration format**](#timer-duration-format)

# Install

See [release](https://github.com/mdouchement/tmlshock/releases)
or run `go install github.com/mdouchement/tmlshock/v2@latest`.

# Settings

## **Color**

You can use color codes (down below), use color names or hex color codes

> [!NOTE]
> Hexa color need a compatible terminal\
> - alacritty, foot, iTerm, kitty, Konsole, st, tmux, vte-based, wezterm, Ghostty, Windows Terminal\
> - See https://github.com/muesli/termenv/tree/master?tab=readme-ov-file#color-support

![color](/.resources/color.png)

```
black
blue
cyan
dark-gray
green
light-green
light-blue
light-cyan
light-gray
light-magenta
light-red
light-yellow
magenta
light-red
light-yellow
red
white
yellow
```

Example:
```bash
tmlshock clock -c red
tmlshock clock -c 1
tmlshock clock -c ff0000
```

## **Date format**

tmlshock uses Golang's [time format](https://pkg.go.dev/time#pkg-constants)

To use a custom date format just enter `-l 02/01/2006`(DD/MM/YYYY)

**Example**
```bash
tmlshock clock -d -l 02/01/2006
```

## **12-hour format**

> [!NOTE]
> This option only for clock.

To use a 12-hour format just enter `-H`

**Example**
```bash
tmlshock clock -H
```

## **Timer duration format**

> [!NOTE]
> This option only for timer.

tmlshock uses Golang's [duration](https://pkg.go.dev/time#ParseDuration) format.

**Example**
```bash
tmlshock timer 4h42m10s
tmlshock timer 2m
tmlshock timer 1h4s
```
