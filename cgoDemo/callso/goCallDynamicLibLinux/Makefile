
all: targetLinux

targetLinux:
	go build -o main-amd64 main.go
	mkdir -p ../bin/linux/
	mv main-amd64 ../bin/linux/
	cp ../cppLib/linux/dynamicLib/libadd.so ../bin/linux/
