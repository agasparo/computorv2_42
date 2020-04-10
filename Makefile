include .env

init:
	go get github.com/fatih/color
	go build -o $(PROJECTNAME)

clean: 
	rm $(PROJECTNAME)

all: init