package statusbar

import (
	"fmt"
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

	// Build the struct
	bar.Msg = msg
	bar.isWaving = true
	bar.Wave = &wave
	bar.Reveal = &reveal

	// Initialize the animations
	wave.Init(&bar.isWaving)
	reveal.Init(revealFrames, 10)

	// Play the wave animation
	go playWave(bar)
}

// Play the wave animation
func playWave(bar *StatusBar) {
	for bar.isWaving {
		fmt.Printf("\r[%s] %s", bar.Wave.Play(), bar.Msg)
		time.Sleep(10 * time.Millisecond)
	}
}

// Print a PASSED status
func (bar *StatusBar) Pass() {
	revealStatus(bar, "PASSED", "GREEN")
}

// Print a FAILED status
func (bar *StatusBar) Fail() {
	revealStatus(bar, "FAILED", "RED")
}

// Print the status given
func revealStatus(bar *StatusBar, status string, colorKey string) {
	bar.isWaving = false

	if !hasStatusFrames(bar, status) {
		bar.Reveal.AddFrames(calcStatusFrames(status, colorKey))
	}

	for range slices.Indexes(bar.Reveal.Frames) {
		fmt.Printf("\r%s %s", bar.Reveal.Play(), bar.Msg)
		time.Sleep(30 * time.Millisecond)
	}

	fmt.Println()
}

// Return true if the last frame has the status given. Else, false.
func hasStatusFrames(bar *StatusBar, status string) bool {
	return bar.Reveal.Frames[len(bar.Reveal.Frames)-1] == "[ "+status+" ]"
}

// Build the frames for the reveal with the status and color given
func calcStatusFrames(status string, colorKey string) []string {
	status = str.Center(status, len(status)+2)
	frames := []string{}

	for i := range intslices.Seq(0, len(status)) {
		frames = append(frames, "["+colors.COLORS[colorKey]+status[:i]+colors.COLORS["RESET"]+"]")
	}

	return frames
}
