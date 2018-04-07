package service

import (
	"testing"
	"time"
)

func TestNewItem(t *testing.T) {
	tm := time.Now()
	tmstr := tm.Format("2006-01-02 03:04:05 PM")
	item := NewItem("play", false, tmstr)
	if item.Label != "play" || item.IsFinished || item.CreateAt != tmstr {
		t.Fail()
	}
}
