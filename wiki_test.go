package main

import (
	"bytes"
	"testing"
)

func TestSaveLoadPage(t *testing.T) {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	if p1.Title != p2.Title || !bytes.Equal(p1.Body, p2.Body) {
		t.Errorf("loaded page: %v\n expected: %v", p2, p1)
	}
}
