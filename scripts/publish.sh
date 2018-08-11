#!/bin/bash
GITC=$(git rev-list --count HEAD)
TAG=$(git describe --tags --abbrev=0)
IFS='.' read -ra vers <<< "$TAG"
MAJOR="${vers[0]}"
MINOR="${vers[1]}"

if [ $1 = "mi" ]; then
    MINOR=$(($MINOR+1))
fi

if [ $1 = "ma" ]; then
    MAJOR=$(($MAJOR+1))
fi

git tag "$MAJOR.$MINOR.$GITC"
git push --tags