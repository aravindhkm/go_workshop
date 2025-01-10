# Define the commit message (this can be passed as a variable)
COMMIT_MSG = "Update changes"

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

# Mark 'test' as a phony target, so make will always execute it
.PHONY: test
test:
	npx hardhat test

.PHONY: call
call:
	npx hardhat run scripts/contractCall.ts --network bscTestnet

.PHONY: id
id:
	npx hardhat run scripts/userId.ts --network bscTestnet

# Mark 'cto' as a phony target, so make will always execute it
.PHONY: cto
cto:
	clear && npx hardhat test test/cto.ts

# Mark 'deploy' as a phony target, so make will always execute it
.PHONY: deploy
deploy:
	npx hardhat run scripts/deploy.ts --network bscTestnet
