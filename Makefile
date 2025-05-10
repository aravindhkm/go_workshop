# Define the commit message (this can be passed as a variable)
COMMIT_MSG = "."

# Define the current branch
CURRENT_BRANCH = $(shell git rev-parse --abbrev-ref HEAD)

# Target to stage changes
add:
	git add .

# Target to commit changes with the specified commit message
commit: add
	git commit -m $(COMMIT_MSG)

# Target to push the changes to the current branch
push: commit
	git push origin $(CURRENT_BRANCH)

# Default target: run all steps (add, commit, push)
all: push

run-%:
	@args="$(wordlist 2, 99, $(MAKECMDGOALS))"; \
	go run main.go $$args

# Prevent make from thinking 'workout' and '1' are targets
%:
	@: