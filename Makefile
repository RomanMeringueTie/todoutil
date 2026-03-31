all: todo

todo: internal/*/*.go cmd/*.go
	go build -o todo cmd/main.go cmd/output.go

.PHONY: clean
clean:
	$(RM) todo

.PHONY: run
run: todo
	./todo

.PHONY: test
test:
	go test ./...
	
