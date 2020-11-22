all: bin/cli bin/server

bin/cli:
	go build -o ./bin/cli ./cli

bin/server:
	go build -o ./bin/server

clean:
	rm ./bin/server
	rm ./bin/cli