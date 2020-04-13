include .env

init:
	go get github.com/fatih/color
	go get -u github.com/wcharczuk/go-chart
	go build -o $(PROJECTNAME)

clean: 
	rm $(PROJECTNAME)

all: init