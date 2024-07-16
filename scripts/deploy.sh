#! /bin/bash
set -e

git restore . && git clean -f -d
# rebase branch
echo "Rebase branch"
git fetch && git rebase

echo "Checking golang version"
/usr/local/go/bin/go version

echo "Install go dependencies"
/usr/local/go/bin/go mod tidy -v

echo "Build Vidinfra"
/usr/local/go/bin/go build -o bin/vidinfra cmd/main.go
echo "Change permission for Vidinfra"
sudo chmod +x bin/vidinfra
echo "Restart Vidinfra service"
sudo systemctl restart api.vidinfra
sudo systemctl status api.vidinfra

echo "Restart Vidinfra Worker service"
sudo systemctl restart worker.vidinfra
sudo systemctl status worker.vidinfra

echo "Restart nginx service"
sudo systemctl restart nginx
echo "🎉 Deploy Vidinfra successfully"
