package main

import "testing"

func TestCheckMatch(t *testing.T) {
	checkPath := "/sample"
	userPath := "/sample"

	t.Logf("checkPath : %v", checkPath)
	t.Logf("userPath : %v", userPath)

	var match = checkMatch(map[string]interface{}{
		"path":   checkPath,
		"target": "",
	}, userPath)

	if match == nil {
		t.Error("Wrong, its must be not nil")
	}
}

func TestCheckMatch2(t *testing.T) {
	checkPath := "/sample"
	userPath := "/sample/"

	t.Logf("checkPath : %v", checkPath)
	t.Logf("userPath : %v", userPath)

	var match = checkMatch(map[string]interface{}{
		"path":   checkPath,
		"target": "",
	}, userPath)

	if match == nil {
		t.Error("Wrong, its must be not nil")
	}
}

func TestCheckMatch3(t *testing.T) {
	checkPath := "/sample"
	userPath := "/SAMPLE"

	t.Logf("checkPath : %v", checkPath)
	t.Logf("userPath : %v", userPath)

	var match = checkMatch(map[string]interface{}{
		"path":   checkPath,
		"target": "",
	}, userPath)

	if match == nil {
		t.Error("Wrong, its must be not nil")
	}
}

func TestCheckMatch4(t *testing.T) {
	checkPath := "/sample"
	userPath := "/SAMPLE/a"

	t.Logf("checkPath : %v", checkPath)
	t.Logf("userPath : %v", userPath)

	var match = checkMatch(map[string]interface{}{
		"path":   checkPath,
		"target": "",
	}, userPath)

	if match != nil {
		t.Error("Wrong, its must be nil")
	}
}

func TestCheckMatch5(t *testing.T) {
	checkPath := "/sample/:wedew"
	userPath := "/SAMPLE/anjay"

	t.Logf("checkPath : %v", checkPath)
	t.Logf("userPath : %v", userPath)

	var match = checkMatch(map[string]interface{}{
		"path":   checkPath,
		"target": "",
	}, userPath)

	if match == nil {
		t.Error("Wrong, its must be not nil")
	}
}

func TestCheckMatch6(t *testing.T) {
	checkPath := "/sample/:wedew/:wedew"
	userPath := "/SAMPLE/anjay/wtf"

	t.Logf("checkPath : %v", checkPath)
	t.Logf("userPath : %v", userPath)

	var match = checkMatch(map[string]interface{}{
		"path":   checkPath,
		"target": "",
	}, userPath)

	if match == nil {
		t.Error("Wrong, its must be not nil")
	}
}

func TestCheckMatch7(t *testing.T) {
	checkPath := "/sample/:wedew/:wedew"
	userPath := "/SAMPLE/anjay"

	t.Logf("checkPath : %v", checkPath)
	t.Logf("userPath : %v", userPath)

	var match = checkMatch(map[string]interface{}{
		"path":   checkPath,
		"target": "",
	}, userPath)

	if match != nil {
		t.Error("Wrong, its must be nil")
	}
}
