package slack

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

const okResponse = `{"ok": true, "error": ""}`

func TestSuccessfulPostMessage(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, okResponse)
	}))
	defer ts.Close()

	slackUrl = ts.URL
	client := New("token")
	err := client.PostMessage("c", "u", "a", "m")
	if err != nil {
		t.Errorf("Error on Post Message %v", err)
	}
}

func TestFailedPostMessage(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"ok": false, "error": "foo"}`)
	}))
	defer ts.Close()

	slackUrl = ts.URL
	client := New("token")
	err := client.PostMessage("c", "u", "a", "m")
	if err.Error() != "foo" {
		t.Errorf("Expected foo, got %s", err)
	}
}

func TestPostMessageCheckPayload(t *testing.T) {
	var payloadRequested string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload, _ := ioutil.ReadAll(r.Body)
		payloadRequested = string(payload)
		fmt.Fprintln(w, okResponse)
	}))
	defer ts.Close()

	slackUrl = ts.URL
	client := New("t")
	err := client.PostMessage("c", "u", "a", "m")
	if err != nil {
		t.Errorf("Error on Post Message %v", err)
	}
	expectedPayload := buildPayload("t", "c", "u", "a", "m")
	if payloadRequested != expectedPayload {
		t.Errorf("Expected %s, got %s", expectedPayload, payloadRequested)
	}
}
