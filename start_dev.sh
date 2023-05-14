apk add --update nodejs npm
npm install nodemon
node_modules/nodemon/bin/nodemon.js --exec go run main.go --signal SIGTERM 