package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char c = 'a';
unsigned char uc = 253;
short s = 233;
int i = 100;
long lt = 11112222;
float f = 3.14;
double db = 3.15;
void * ptr;
const char * pc = "huangjian";
char ac[20] = "char array";

int f1() {
	return 200;
}

const char * f2() {
	return "I'm f2";
}

char * f3() {
	char * pc = (char*)malloc(10*sizeof(char));
	pc[0] = 'a';
	pc[1] = 'b';
	pc[2] = 'c';
	pc[3] = '\0';
	return pc;
}

void f4(char ** out) {
	char * pc = (char*)malloc(10*sizeof(char));
	pc[0] = 'd';
	pc[1] = 'e';
	pc[2] = 'f';
	pc[3] = '\0';
	*out = pc;
}

void printI(void *i) {
	printf("print i = %d\n", *(int *)i);
}

struct ImgInfo {
	char *imgPath;
	int format;
	unsigned int width;
	unsigned int height;
};
void printStruct(struct ImgInfo *imgInfo) {
	fprintf(stdout, "imgPath = %s\n", imgInfo->imgPath);
	fprintf(stdout, "format = %d\n", imgInfo->format);
	fprintf(stdout, "width = %d\n", imgInfo->width);
	fprintf(stdout, "height = %d\n", imgInfo->height);
}

*/
import "C"
import (
	"fmt"
	"reflect"
	"unsafe"
)

func go2c() {
	fmt.Println("-------------------------Go to C-------------------------")
	fmt.Println(C._Bool(true))
	fmt.Println(C._Bool(false))
	fmt.Println(C.char('Y'))
	fmt.Printf("%c\n", C.char('Y'))
	fmt.Println(C.uchar('C'))
	fmt.Println(C.short(254))
	fmt.Println(C.long(11112222))

	var goi int = 2
	cpi := unsafe.Pointer(&goi)
	C.printI(cpi)

	C.ptr = unsafe.Pointer(nil)
	fmt.Println(C.ptr)
}

func c2go() {
	fmt.Println("-------------------------C to Go-------------------------")

	fmt.Printf("C.c = %c, %v\n", C.c, reflect.TypeOf(C.c))
	fmt.Printf("C.uc = %v\n", C.uc)
	fmt.Printf("C.s = %v\n", C.s)
	fmt.Printf("C.i = %v\n", C.i)
	fmt.Printf("C.lt = %v\n", C.lt)
	fmt.Printf("C.f = %v, %v, %v, %v\n", C.f, reflect.TypeOf(C.f), float32(C.f), reflect.TypeOf(float32(C.f)))
	fmt.Printf("C.db = %v, %v, %v, %v\n", C.db, reflect.TypeOf(C.db), float64(C.db), reflect.TypeOf(float64(C.db)))

	pc := C.GoString(C.pc)
	fmt.Printf("C.pc = %v, %v, %v, %v\n", C.pc, reflect.TypeOf(C.pc), pc, reflect.TypeOf(pc))

	var charray []byte
	for i := range C.ac {
		if C.ac[i] != 0 {
			charray = append(charray, byte(C.ac[i]))
		}
	}
	fmt.Printf("C.ac = %v, %v, %v, %v, %v\n", C.ac, reflect.TypeOf(C.ac), charray, reflect.TypeOf(charray), string(charray))

	// function
	fmt.Printf("C.f1() = %v\n", C.f1())

	f2ret := C.GoString(C.f2())
	fmt.Printf("C.f2() = %v\n", f2ret)

	f3ret := C.f3()
	f3retStr := C.GoString(f3ret)
	fmt.Printf("C.f3() = %v\n", f3retStr)
	C.free(unsafe.Pointer(f3ret))

	var f4out *C.char
	C.f4(&f4out)
	f4outStr := C.GoString(f4out)
	fmt.Printf("C.f4() = %v\n", f4outStr)
	C.free(unsafe.Pointer(f4out))

	imgInfo := C.struct_ImgInfo{
		imgPath: C.CString("xx.jpg"),
		format:  0,
		width:   500,
		height:  400,
	}
	defer C.free(unsafe.Pointer(imgInfo.imgPath))
	fmt.Printf("imgInfo = %v, %v\n", imgInfo, reflect.TypeOf(imgInfo))
	C.printStruct(&imgInfo)
}

func main() {
	go2c()
	c2go()
}
