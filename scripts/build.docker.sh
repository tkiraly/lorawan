GITC=$(git rev-list --count HEAD)
TAG=$(git describe --tags --abbrev=0)
IFS='.' read -ra vers <<< "$TAG"
MAJOR="${vers[0]}"
MINOR="${vers[1]}"

docker build -t registry.gitlab.com/loriot/nawarol:$MAJOR.$MINOR.$GITC .
docker push registry.gitlab.com/loriot/nawarol:$MAJOR.$MINOR.$GITC