package controller

import (
	"github.com/HiLittleCat/core"

	"encoding/json"
	"net/http"
	"net/url"
	"testing"
)

const (
	URL = "http://127.0.0.1:9000/User"
)

func TestCreate(t *testing.T) {
	data := make(url.Values)
	data["mobile"] = []string{"189XXXX8340"}
	data["name"] = []string{"perfect"}
	resp, err := http.PostForm(URL+"/Create", data)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	var dst struct{ Salutation string }
	if err := json.NewDecoder(resp.Body).Decode(&dst); err != nil {
		t.Fatal(err)
	}
}

func TestGet(t *testing.T) {
	resp, err := http.Get(URL + "?_id=59cc99b1ea4005ac61be70c2")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	var dst struct{ Salutation string }
	if err := json.NewDecoder(resp.Body).Decode(&dst); err != nil {
		t.Fatal(err)
	}
}

func TestGetPage(t *testing.T) {
	resp, err := http.Get(URL + "/Page?page=1&pageCount=10")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	var dst core.ResFormat
	if err := json.NewDecoder(resp.Body).Decode(&dst); err != nil {
		t.Fatal(err)
	}
	if dst.Ok != true {
		t.Fatal(dst.Message)
	}
}
