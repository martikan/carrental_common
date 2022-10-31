.PHONY: clean build test sec

clean:
	rm *.out
	go clean

test:
	go test -v ./... -race -coverprofile=coverage.out -covermode=atomic
