FILE = back
DIR = ./out

build:
	go build -o out/"$(FILE)"

run:
	clear
	go run server.go
