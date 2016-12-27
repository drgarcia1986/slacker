package slack

import (
	"io/ioutil"
	"net/http"
	"testing"

	httpmock "gopkg.in/jarcoal/httpmock.v1"
)

const okResponse = `{"ok": true, "error": ""}`

func TestSuccessfulPostMessage(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", slackUrl,
		httpmock.NewStringResponder(201, okResponse),
	)

	client := New("token")
	err := client.PostMessage("c", "u", "a", "m")
	if err != nil {
		t.Errorf("Error on Post Message %v", err)
	}
}

func TestFailedPostMessage(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", slackUrl,
		httpmock.NewStringResponder(201, `{"ok": false, "error": "foo"}`),
	)

	client := New("token")
	err := client.PostMessage("c", "u", "a", "m")
	if err.Error() != "foo" {
		t.Errorf("Expected foo, got %s", err)
	}
}

func TestPostMessageCheckPayload(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	var payloadRequested string
	httpmock.RegisterResponder("POST", slackUrl,
		func(req *http.Request) (*http.Response, error) {
			payload, _ := ioutil.ReadAll(req.Body)
			payloadRequested = string(payload)
			return httpmock.NewStringResponse(201, okResponse), nil
		},
	)

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
