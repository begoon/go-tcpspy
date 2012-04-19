all: format run

file = gotcpspy

fmt_flags = -w=true -tabs=false -tabwidth=2

format:
	gofmt $(fmt_flags) $(file).go

run:
	go run $(file).go $(args)

