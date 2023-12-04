binary_file=aoc

build:
	go build -o $(binary_file) .

clean:
	rm -f $(binary_file)

run_all:
	./$(binary_file)
