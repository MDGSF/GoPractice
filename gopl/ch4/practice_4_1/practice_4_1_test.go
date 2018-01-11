package main

import "testing"

func Test_getOneBit(t *testing.T) {

	if ret := getOneBit(byte(8), uint(0)); ret != 0 {
		t.Errorf("getOneBit() 8 0 ret = %v", ret)
	}
	if ret := getOneBit(byte(8), uint(1)); ret != 0 {
		t.Errorf("getOneBit() 8 1 ret = %v", ret)
	}
	if ret := getOneBit(byte(8), uint(2)); ret != 0 {
		t.Errorf("getOneBit() 8 2 ret = %v", ret)
	}
	if ret := getOneBit(byte(8), uint(3)); ret != 1 {
		t.Errorf("getOneBit() 8 3 ret = %v", ret)
	}

	if ret := getOneBit(byte(7), uint(0)); ret != 1 {
		t.Errorf("getOneBit() 7 0 ret = %v", ret)
	}
	if ret := getOneBit(byte(7), uint(1)); ret != 1 {
		t.Errorf("getOneBit() 7 1 ret = %v", ret)
	}
	if ret := getOneBit(byte(7), uint(2)); ret != 1 {
		t.Errorf("getOneBit() 7 2 ret = %v", ret)
	}
	if ret := getOneBit(byte(7), uint(3)); ret != 0 {
		t.Errorf("getOneBit() 7 3 ret = %v", ret)
	}

}
