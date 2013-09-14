package main

import (
	"github.com/daneharrigan/hipchat"
	"strings"
)

// Returns the appropriate reply message for a given ping
func replyMessage(message hipchat.Message) (reply, kind string) {
	if strings.Contains(message.Body, "logo") {
		return "<img src='" + LOGO_URL + "'/>", "html"
	} else {
		return "Hello, " + name(message.From), "text"
	}
}