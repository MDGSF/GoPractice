ANDROIDGCC=/usr/local/ndk-toolchain-arm64/bin/aarch64-linux-android-g++

all: android dynamicLib staticLib

android:
	$(ANDROIDGCC) -fPIC -shared -o libadd.so add.cpp
	mkdir -p ../cppLib/android/dynamicLib
	mv libadd.so ../cppLib/android/dynamicLib

dynamicLib:
	g++ -fPIC -shared -o libadd.so add.cpp
	mkdir -p ../cppLib/linux/dynamicLib
	mv libadd.so ../cppLib/linux/dynamicLib

staticLib:
	g++ -Wall -c -o add.o add.cpp && ar rcs libadd.a add.o
	rm -rf add.o
	mkdir -p ../cppLib/linux/staticLib
	mv libadd.a ../cppLib/linux/staticLib
