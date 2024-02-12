package statusbar

import (
	"fmt"

	"github.com/12yanogden/colors"
	shutters "github.com/12yanogden/spinners/shutters"
)

type StatusBar struct {
	Msg        string
	isSpinning bool
	Spinner    shutters.Shutters
}

func (bar *StatusBar) Start(msg string) {
	// Build the struct
	bar.Msg = msg
	bar.isSpinning = true
	bar.Spinner = shutters.New(&bar.isSpinning)

	// Play the wave animation
	go spin(bar)
}

// Play the wave animation
func spin(bar *StatusBar) {
	for bar.isSpinning {
		fmt.Printf("\r[ %s ] %s", bar.Spinner.Play(), bar.Msg)
	}
}

// Print a PASSED status
func (bar *StatusBar) Pass() {
	bar.revealStatus("PASSED", "GREEN")
}

// Print a FAILED status
func (bar *StatusBar) Fail() {
	bar.revealStatus("FAILED", "RED")
}

// Print the status given
func (bar *StatusBar) revealStatus(status string, colorKey string) {
	bar.isSpinning = false

	fmt.Printf("\r[ %s%s%s ] %s\n",
		colors.COLORS[colorKey],
		status,
		colors.COLORS["RESET"],
		bar.Msg,
	)
}
