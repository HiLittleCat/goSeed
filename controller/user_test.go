package controller

import (
	"encoding/json"
	"net/http"
	"testing"
)

const (
	URL   = "http://127.0.0.1:9000/User"
	API_N = 1000
)

func TestGet(t *testing.T) {
	resp, err := http.Get(URL + "?_id=59cc99b1ea4005ac61be70c2")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	var dst struct{ Salutation string }
	if err := json.NewDecoder(resp.Body).Decode(&dst); err != nil {
		t.Fatal(err)
	}
}

func TestGetAll(t *testing.T) {
	resp, err := http.Get(URL + "/All")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
}
