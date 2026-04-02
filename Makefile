all: todo

todo: internal/*/*.go cmd/*.go
	go build -o todo cmd/main.go

.PHONY: clean
clean:
	$(RM) todo

.PHONY: run
run: todo
	./todo

.PHONY: test_local
test_local:
	go test ./... -v
	
.PHONY: test_CI
test_CI:
	go test ./... -json > TestResults.json