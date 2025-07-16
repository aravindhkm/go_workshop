#!/bin/bash

DIFF=$(git diff --cached)

COMMIT_MSG=$(curl -s https://api.openai.com/v1/chat/completions \
  -H "Authorization: Bearer $OPENAI_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "gpt-4",
    "messages": [
      {"role": "system", "content": "You are a helpful assistant that writes concise Git commit messages."},
      {"role": "user", "content": "Generate a Git commit message based on this diff:\n'"$DIFF"'"}
    ]
  }' | jq -r '.choices[0].message.content')

echo "$COMMIT_MSG"
