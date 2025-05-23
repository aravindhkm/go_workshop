package interview

import (
	"fmt"
	"math"
)

// MinHandMoves returns the minimum number of hand moves needed
func MinHandMoves(startThumb int, notes []int) int {
	const maxNote = 100 // Assumption: max note value (can be adjusted based on problem)

	n := len(notes)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, maxNote+1)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}

	// Initialize with the first note using the starting thumb position
	for thumb := startThumb - 4; thumb <= startThumb; thumb++ {
		if thumb >= 1 && thumb+4 <= maxNote {
			if notes[0] >= thumb && notes[0] <= thumb+4 {
				dp[0][thumb] = 0
			}
		}
	}

	for i := 1; i < n; i++ {
		for thumb := 1; thumb <= maxNote-4; thumb++ {
			if dp[i-1][thumb] == math.MaxInt32 {
				continue
			}
			// Try to play current note with this hand position or move
			for newThumb := 1; newThumb <= maxNote-4; newThumb++ {
				if notes[i] >= newThumb && notes[i] <= newThumb+4 {
					moveCost := 0
					if newThumb != thumb {
						moveCost = 1
					}
					dp[i][newThumb] = min(dp[i][newThumb], dp[i-1][thumb]+moveCost)
				}
			}
		}
	}

	// Find the minimum in the last row
	minMoves := math.MaxInt32
	for thumb := 1; thumb <= maxNote-4; thumb++ {
		minMoves = min(minMoves, dp[n-1][thumb])
	}

	return minMoves
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Example usage
func PianoOne() {
	startThumb := 10
	notes := []int{10, 11, 12, 17, 18, 22}

	fmt.Println("Minimum hand moves:", MinHandMoves(startThumb, notes))
}


// Let's write a simple 2D Dynamic Programming solution in Go for this problem.


// 🧠 Problem Understanding:
// A piano player uses one hand with 5 fingers to play.

// The thumb's initial position is given (i.e., the key it's placed on).

// The player can reach 5 consecutive notes from the thumb position using one hand.

// The player must play a series of notes in order.

// The goal is to minimize the number of hand movements (i.e., shifting the hand left or right).

// 🟢 Step-by-step Simulation:
// ✅ Initial hand position:
// Thumb is at 10, so hand covers keys: 10 11 12 13 14

// This can play:

// ✅ 11

// ✅ 13

// ❌ 15 (not covered)

// So to play 15, the hand must move.

// 🟢 First move:
// Move thumb to 13 → hand covers: 13 14 15 16 17

// This covers:

// ✅ 15

// ✅ 17

// ❌ 19

// Still one more move needed.

// 🟢 Second move:
// Move thumb to 16 → hand covers: 16 17 18 19 20

// This covers:

// ✅ 19

// So we moved the hand:

// From 10 → 13 (1st move)

// From 13 → 16 (2nd move)

// 🧮 Total Moves: 2

// 🧠 Why Dynamic Programming?
// Instead of trying every possible combination of hand movements manually, the DP table helps:

// Tracks all possible hand positions after playing each note.

// For each note, considers all hand positions that could play it.

// Carries forward the minimum number of moves needed to reach that point.

// This is why the 2D DP approach is both simple and optimal for this.

// Let me know if you want a visual representation or if you'd like to walk through the DP table for this input!