# Define the current branch
CURRENT_BRANCH = $(shell git rev-parse --abbrev-ref HEAD)

# Generate a commit message like GitLens (last commit diff summary)
COMMIT_MSG = $(shell git diff --cached --name-status | awk '{print $$2}' | paste -sd, -)

# Target to stage changes
add:
	git add .

# Target to commit changes with the specified commit message
commit: add
	git commit -m "$(COMMIT_MSG)"

# Target to push the changes to the current branch
push: commit
	git push -f origin $(CURRENT_BRANCH)

# Default target: run all steps (add, commit, push)
all: push

# Run a Go file with arguments like `make run-arg1 arg2`
run-%:
	@args="$(wordlist 2, 99, $(MAKECMDGOALS))"; \
	go run main.go $$args

# Prevent make from treating plain arguments as targets
%:
	@:
