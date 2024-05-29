build:
	cd '$(CURDIR)/sql/schema/' && goose postgres postgresql://root:password@localhost:5432/chat-app up
	cd '$(CURDIR)/cmd/' && go build -o ../bin/chat-server

run: build
	./bin/chat-server

test:
	go test -v ./...
