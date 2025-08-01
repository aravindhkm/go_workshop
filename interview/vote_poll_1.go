package interview

import (
	"fmt"
	"sync"
)

// PollRoom holds poll details
type PollRoom struct {
	Question   string
	Candidates []string
	Votes      map[string]int
	VotedUsers map[string]bool
	mu         sync.Mutex
}

// NewPollRoom creates a new PollRoom
func NewPollRoom(question string, candidates []string) *PollRoom {
	return &PollRoom{
		Question:   question,
		Candidates: candidates,
		Votes:      make(map[string]int),
		VotedUsers: make(map[string]bool),
	}
}

// Vote allows a user to vote for an option
func (pr *PollRoom) Vote(user string, candidate string) error {
	pr.mu.Lock()
	defer pr.mu.Unlock()

	if pr.VotedUsers[user] {
		return fmt.Errorf("user %s has already voted", user)
	}

	valid := false
	for _, cnd := range pr.Candidates {
		if cnd == candidate {
			valid = true
			break
		}
	}
	if !valid {
		return fmt.Errorf("invalid candidate: %s", candidate)
	}

	pr.Votes[candidate]++
	pr.VotedUsers[user] = true
	return nil
}

// ShowResults displays the current vote tally
func (pr *PollRoom) ShowResults() {
	fmt.Printf("\n📊 Results for: %s\n", pr.Question)
	for _, opt := range pr.Candidates {
		fmt.Printf("%s: %d votes\n", opt, pr.Votes[opt])
	}
}

func VotePollOne() {
	// Setup poll
	poll := NewPollRoom("Who should be the team lead?", []string{"Alice", "Bob", "Charlie"})

	// Simulate voting
	users := map[string]string{
		"user1": "Alice",
		"user2": "Bob",
		"user3": "Alice",
		"user4": "Charlie",
		"user5": "Alice",
		"user6": "Charlie",
	}

	for user, vote := range users {
		err := poll.Vote(user, vote)
		if err != nil {
			fmt.Printf("❌ %s: %v\n", user, err)
		} else {
			fmt.Printf("✅ %s voted for %s\n", user, vote)
		}
	}

	// Show final results
	poll.ShowResults()
}
