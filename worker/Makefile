.PHONY: build run

build:
	go build -o worker

run: build dbstruct
	./worker

dbstruct:
	chmod +x dbstruct.sh
	./dbstruct.sh
