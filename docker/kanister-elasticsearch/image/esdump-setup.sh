#! /bin/sh

apk add --update bash
apk add --update nodejs
apk add --update npm
apk add --update curl

npm install -g npm yo grunt-cli bower express

# check locations and packages are correct
which node
node -v
which npm
npm -v
npm ls -g --depth=0
npm install elasticdump -g
