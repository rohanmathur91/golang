.PHONY: clean build run

clean:
	rm -f bin

build: clean
	@echo "Building..."
	go build -o bin .

run: build 
	@echo "Running..." 
	./bin 