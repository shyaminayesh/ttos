default: run

run:
	sudo go run ttos.go

dist:
	go build -o ttos -ldflags "-s -w" -trimpath ttos.go
