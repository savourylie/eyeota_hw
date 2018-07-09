package main

import (
	"reflect"
	"testing"
)

func TestGetSegmentForTypeAndEmpty(t *testing.T) {

	s := getSegmentFor("org1", "paramName1")
	sType := reflect.TypeOf(s)

	// Check type
	if sType != reflect.TypeOf([]SegmentConfig{}) {
		t.Errorf("Type is not correct.")
	}

	// Check length
	if len(s) != 0 {
		t.Errorf("The array is not empty!")
	}

}

func TestGetSegmentForValOneElement(t *testing.T) {
	s := getSegmentFor("org1", "paramName1", "paramVal1")
	sType := reflect.TypeOf(s)

	// Check type
	if sType != reflect.TypeOf([]SegmentConfig{}) {
		t.Errorf("Type is not correct.")
	}

	// Check length
	if len(s) != 1 {
		t.Errorf("The array is not empty!")
	}

	if s[0].Id != "seg_1234" {
		t.Errorf("ValueError!")
	}
}
