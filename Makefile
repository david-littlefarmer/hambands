run:
	@go run ./main.go

rerun:
	@go run github.com/goware/rerun/cmd/rerun -watch ./ -ignore bin -run 'GOGC=off go build -o ./bin/hambands ./main.go && ./bin/hambands'
