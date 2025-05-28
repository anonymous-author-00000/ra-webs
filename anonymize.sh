# this script is used to anonymize the ra-webs repository
# memo: docker run --volume ./:/app -it ubuntu:latest bash

apt-get update && apt-get install -y git
git clone git@github.com:anonymous-author-00000/ra-webs-private.git
rm -rf ra-webs-private/ra-webs/.git 
cp -r ra-webs-private/* .
rm -rf ./ra-webs-private

git config --local user.name "Anonymous Author"
git config --local user.email "213727166+anonymous-author-00000@users.noreply.github.com"

git config --global --add safe.directory /app

git add .
git commit --amend -m "anonymize" 
git push -f origin main

