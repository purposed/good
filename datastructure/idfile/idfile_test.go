package idfile_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/purposed/good/datastructure/idfile"
)

func createTempFile() (string, error) {
	tempFile, err := ioutil.TempFile("", "idfile")
	if err != nil {
		return "", fmt.Errorf("could not create temp file: %s", err.Error())
	}
	defer tempFile.Close()

	return tempFile.Name(), os.Remove(tempFile.Name())
}

func Test_IDFile(t *testing.T) {
	tempFilePath, err := createTempFile()
	if err != nil {
		t.Error(err.Error())
		return
	}

	defer func() {
		// Cleanup temp file.
		if _, err := os.Stat(tempFilePath); err == nil {
			if err := os.Remove(tempFilePath); err != nil {
				panic(err)
			}
		}
	}()

	// Create a blank IDFile, make sure creation doesn't fail.
	mf, err := idfile.New(tempFilePath)
	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
		return
	}

	// Make sure that the initial ID is correct.
	nextID := mf.NextID()
	if nextID != 0 {
		t.Errorf("incorrect initial nextID: %d", nextID)
		return
	}

	// Make sure both internal maps are initially empty.
	if _, ok := mf.ResolveValue("hello"); ok {
		t.Errorf("IDFile was supposed to be empty")
		return
	}

	if _, err := mf.ResolveID(0); err == nil {
		t.Errorf("IDFile was supposed to be empty")
		return
	}

	// Try adding values to the map.
	if id := mf.AddValue("hello"); id != 0 {
		t.Errorf("got unexpected id: %d instead of 0", id)
		return
	}
	if id := mf.AddValue("world"); id != 1 {
		t.Errorf("got unexpected id: %d instead of 1", id)
		return
	}

	// Try fetching back those values.
	val, err := mf.ResolveID(0)
	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
		return
	}

	if val != "hello" {
		t.Errorf("wrong value: %s", val)
		return
	}

	eID, ok := mf.ResolveValue("hello")
	if !ok {
		t.Errorf("value not saved properly")
		return
	}
	if eID != 0 {
		t.Errorf("returned wrong id: %d", eID)
		return
	}

	// Commit the file & reload it.
	if err := mf.Commit(); err != nil {
		t.Errorf("could not commit: %s", err.Error())
		return
	}

	mf2, err := idfile.New(tempFilePath)
	if err != nil {
		t.Errorf("reloading failed: %s", err.Error())
		return
	}

	val, err = mf2.ResolveID(1)
	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
		return
	}

	if val != "world" {
		t.Errorf("bad value: %s", val)
		return
	}

	eID, ok = mf2.ResolveValue("world")
	if !ok {
		t.Errorf("value not saved properly")
		return
	}
	if eID != 1 {
		t.Errorf("returned wrong id: %d", eID)
		return
	}

	// Test re-adding an existing value.
	if rID := mf2.AddValue("world"); rID != 1 {
		t.Errorf("re-adding value yielded different ID")
		return
	}
}
