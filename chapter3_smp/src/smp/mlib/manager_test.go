package mlib

import "testing"

func TestOps(t *testing.T) {
	mm := NewMusicManager()
	if mm == nil {
		t.Error("New MusicManager failed.")
	}

	if mm.Len() != 0 {
		t.Error("New MusicManager failed,not empty")
	}

	m0 := &MusicEntry{
		"1",
		"My Heart Will Go On",
		"Celion Dion",
		"http://qbox.me/24501234",
		"MP3",
	}
	mm.Add(m0)

	if mm.Len() != 1 {
		t.Error("MusicManager.Add() failed.") //t.Error()会自动结束进程的运行
	}

	m := mm.Find(m0.Name)
	if m == nil {
		t.Error("MusicManager.Find failed.")
	}

	if m.Id != m0.Id || m.Artist != m0.Artist ||
		m.Name != m0.Name || m.Source != m0.Source || m.Type != m0.Type {
		t.Error("MusicManager.Find() failed.Found item mismatch")
	}

	m, err := mm.Get(0)
	if m == nil {
		t.Error("MusicManager.Get() failed.", err)
	}

	m = mm.Remove(0)

	if m == nil || mm.Len() != 0 {
		t.Error("MusicManager.Remove() failed.")
	}

	//======================================================
	m1 := &MusicEntry{
		"1",
		"Take me to your heart",
		"Celion Dion",
		"http://qbox.me/24501234",
		"MP3",
	}
	mm.Add(m1)
	if mm.Len() != 1 {
		t.Error("Add 'Take me to your heart' failure")
	}

	m = mm.RemoveByName("Take me to your heart")

	if m == nil || mm.Len() != 0 {
		t.Error("MusicManager.RemoveByName() failed")
	}

}
