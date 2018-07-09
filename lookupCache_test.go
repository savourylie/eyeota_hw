package main

import (
	"fmt"
	"reflect"
	"testing"

	"./lookupcache"
)

func TestGetSegmentForOrgAndKeyTypeAndEmpty(t *testing.T) {
	filename := "./data/test.json"
	lc := initCache(filename)

	s := lc.GetSegmentForOrgAndKey("org1", "paramName1")
	sType := reflect.TypeOf(s)

	fmt.Println(s)
	// Check type
	if sType != reflect.TypeOf([]lookupcache.SegmentConfig{}) {
		t.Errorf("Type is not correct.")
	}

	// Check length
	if len(s) != 0 {
		t.Errorf("The array is not empty!")
	}

}

func TestGetSegmentForOrgAndKeyAndValOneElement(t *testing.T) {
	filename := "./data/test.json"
	lc := initCache(filename)

	s := lc.GetSegmentForOrgAndKeyAndVal("org1", "paramName1", "paramVal1")
	sType := reflect.TypeOf(s)

	// Check type
	if sType != reflect.TypeOf([]lookupcache.SegmentConfig{}) {
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

func TestGetSegmentForOrgAndKeyAndValMultiple(t *testing.T) {
	filename := "./data/test.json"
	lc := initCache(filename)

	paramValSlice := []string{"paramVal2", "paramVal3", "paramVal4", "paramVal5"}

	for _, pv := range paramValSlice {
		s := lc.GetSegmentForOrgAndKeyAndVal("org1", "paramName1", pv)

		// Check length
		if len(s) != 1 {
			t.Errorf("The array is not empty!")
		}

		if s[0].Id != "intr.edu" {
			t.Errorf("ValueError!")
		}
	}

}

func TestGetSegmentForOrgAndKeyAndVal6(t *testing.T) {
	filename := "./data/test.json"
	lc := initCache(filename)

	target := []string{"dem.infg.m", "intr.heal", "dem.infg.f"}
	s := lc.GetSegmentForOrgAndKeyAndVal("org1", "paramName1", "paramVal6")

	// Check length
	if len(s) != 3 {
		t.Errorf("The array is not empty!")
	}

	result := []string{}

	for _, seg := range s {
		result = append(result, seg.Id)
	}

	for i := range result {
		if result[i] != target[i] {
			t.Errorf("ValueError!")
		}
	}
}

func TestGetSegmentForOrgAndKey_testedu(t *testing.T) {
	filename := "./data/test.json"
	lc := initCache(filename)

	s := lc.GetSegmentForOrgAndKey("org1", "testedu")

	// fmt.Println(s)
	// Check length
	if len(s) != 1 {
		t.Errorf("The number of element is not 1!")
	}

	if s[0].Id != "n277" {
		t.Errorf("ValueError!")
	}
}

func TestGetSegmentForOrgAndKey_testeduEmptyString(t *testing.T) {
	filename := "./data/test.json"
	lc := initCache(filename)

	s := lc.GetSegmentForOrgAndKeyAndVal("org1", "testedu", "")

	// fmt.Println(s)
	// Check length
	if len(s) != 1 {
		t.Errorf("The number of element is not 1!")
	}

	if s[0].Id != "n277" {
		t.Errorf("ValueError!")
	}
}

func TestGetSegmentForOrgAndKey_testeduRandom(t *testing.T) {
	filename := "./data/test.json"
	lc := initCache(filename)

	s := lc.GetSegmentForOrgAndKeyAndVal("org1", "testedu", "DonaldTrump")

	// fmt.Println(s)
	// Check length
	if len(s) != 0 {
		t.Errorf("The array is not empty!")
	}
}

func TestGetSegmentForOrgAndKey_genFemale(t *testing.T) {
	filename := "./data/test.json"
	lc := initCache(filename)

	s := lc.GetSegmentForOrgAndKeyAndVal("org1", "gen", "Female")

	// fmt.Println(s)
	// Check length
	if len(s) != 1 {
		t.Errorf("The number of element is not 1!")
	}

	if s[0].Id != "dem.g.f" {
		t.Errorf("ValueError!")
	}
}

func TestGetSegmentForOrgAndKey_genMale(t *testing.T) {
	filename := "./data/test.json"
	lc := initCache(filename)

	s := lc.GetSegmentForOrgAndKeyAndVal("org1", "gen", "Male")

	// fmt.Println(s)
	// Check length
	if len(s) != 1 {
		t.Errorf("The number of element is not 1!")
	}

	if s[0].Id != "dem.g.m" {
		t.Errorf("ValueError!")
	}
}

func TestGetSegmentForOrgAndKey_genEmpty(t *testing.T) {
	filename := "./data/test.json"
	lc := initCache(filename)

	s := lc.GetSegmentForOrgAndKeyAndVal("org1", "gen", "JFK")

	// fmt.Println(s)
	// Check length
	if len(s) != 0 {
		t.Errorf("The array is not empty!")
	}
}
