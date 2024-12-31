#!/bin/bash
git fetch --tags 
latest_tag=$(git tag --sort=committerdate | grep -o 'v.*' | sort -r | head -1)

if [[ -z "$latest_tag" ]]; then
  echo "v0.1.0"
  exit 0
fi

major=$(echo "$latest_tag" | cut -d'.' -f1 | sed 's/v//')
minor=$(echo "$latest_tag" | cut -d'.' -f2)
patch=$(echo "$latest_tag" | cut -d'.' -f3)

if git log --oneline -1 | grep -q "BREAKING CHANGE"; then
  major=$((major + 1))
  minor=0
  patch=0
elif git log --oneline -1 | grep -q "feat"; then
  minor=$((minor + 1))
  patch=0
else
  patch=$((patch + 1))
fi

echo "v${major}.${minor}.${patch}"