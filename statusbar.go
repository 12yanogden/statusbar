package statusbar

import (
	"fmt"
	"strings"
	"time"

	"github.com/12yanogden/colors"
	"github.com/12yanogden/intslices"
	"github.com/12yanogden/reel"
	"github.com/12yanogden/slices"
	wave "github.com/12yanogden/spinners/wave"
	"github.com/12yanogden/str"
)

type StatusBar struct {
	Msg      string
	isWaving bool
	Wave     *(wave.Wave)
	Reveal   *(reel.Reel)
}

func (bar *StatusBar) Start(msg string) {
	var wave wave.Wave
	var reveal reel.Reel

	revealFrames := []string{
		"[        ]",
		"[       ]",
		"[      ]",
		"[     ]",
		"[    ]",
		"[   ]",
		"[  ]",
		"[ ]",
		"[]",
	}

	bar.isWaving = true
	bar.Msg = msg
	bar.Wave = &wave
	bar.Reveal = &reveal

	wave.Init(&bar.isWaving)
	reveal.Init(revealFrames, 10)

	go playWave(bar)
}

func playWave(bar *StatusBar) {
	for bar.isWaving {
		fmt.Printf("\r[%s] %s", bar.Wave.Play(), bar.Msg)
		time.Sleep(10 * time.Millisecond)
	}
}

func (bar *StatusBar) Pass() {
	revealStatus(bar, "PASSED", "GREEN")
}

func (bar *StatusBar) Fail() {
	revealStatus(bar, "FAILED", "RED")
}

func revealStatus(bar *StatusBar, status string, colorKey string) {
	bar.isWaving = false

	if !hasStatusFrames(bar, status) {
		bar.Reveal.AddFrames(calcStatusFrames(status, colorKey))
	}

	for range slices.Indexes(bar.Reveal.Frames) {
		fmt.Printf("\r%s %s", bar.Reveal.Play(), bar.Msg)
		time.Sleep(50 * time.Millisecond)
	}

	fmt.Println()
}

func hasStatusFrames(bar *StatusBar, status string) bool {
	return strings.Contains(bar.Reveal.Frames[len(bar.Reveal.Frames)-1], status)
}

func calcStatusFrames(status string, colorKey string) []string {
	status = str.Center(status, len(status)+2)
	frames := []string{}

	for i := range intslices.Seq(0, len(status)) {
		frames = append(frames, "["+colors.COLORS[colorKey]+status[:i]+colors.COLORS["RESET"]+"]")
	}

	return frames
}
