
all: amd64 arm64

amd64:
	go build -tags="li" -o amd64 main.go

arm64:
	go build -tags="an" -o arm64 main.go

clean:
	rm -rf amd64 arm64
