GITC=$(git rev-list --count HEAD)
TAG=$(git describe --tags --abbrev=0)
IFS='.' read -ra vers <<< "$TAG"
MAJOR="${vers[0]}"
MINOR="${vers[1]}"

go build -o nawarol -ldflags "-X 'gitlab.com/loriot/nawarol/cmd.version=$MAJOR.$MINOR.$GITC'" main.go
