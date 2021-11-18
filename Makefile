MAIN_FILE=src/main.go
PROG=game
LINUX_PLATFORM=linux
DARWIN_PLATFORM=darwin

build:
	go build -o $(PROG) $(MAIN_FILE)

run:
	go run $(MAIN_FILE)

compile:
	echo "Compile pour tous les OS et Platforms"
	GOOS=$(LINUX_PLATFORM) GOARCH=amd64 go build -o $(PROG)-$(LINUX_PLATFORM) $(MAIN_FILE)
	GOOS=$(DARWIN_PLATFORM) GOARCH=amd64 go build -o $(PROG)-$(DARWIN_PLATFORM) $(MAIN_FILE)



