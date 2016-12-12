package slack

import (
	"fmt"
	"testing"
)

func TestGetAvatarFieldForEmoji(t *testing.T) {
	expected := "icon_emoji"
	avatarField := getAvatarField(":scream:")

	if avatarField != expected {
		t.Errorf("Expected %s, got %s", expected, avatarField)
	}
}

func TestGetAvatarFieldForUrl(t *testing.T) {
	expected := "icon_url"
	avatarField := getAvatarField("http://foo.bar/barz.jpg")

	if avatarField != expected {
		t.Errorf("Expected %s, got %s", expected, avatarField)
	}
}

func TestBuildPayload(t *testing.T) {
	token := "token"
	channel := "general"
	username := "slacker"
	avatar := ":scream:"
	message := "test"

	expectedAvatar := "%3Ascream%3A"
	expected := fmt.Sprintf(
		"as_user=false&channel=%v&icon_emoji=%v&parse=full&text=%v&token=%v&username=%v",
		channel, expectedAvatar, message, token, username,
	)
	payload := buildPayload(token, channel, username, avatar, message)

	if payload != expected {
		t.Errorf("Expected %s, got %s", expected, payload)
	}
}
