.PHONY: clean build run

clean:
	rm -f bin

build: clean
	go build -o bin .

run: build 
	./bin 