package fcm

import (
	"errors"
	"time"
)

var (
	// ErrInvalidMessage occurs if push notification message is nil.
	ErrInvalidMessage = errors.New("message is invalid")

	// ErrInvalidTarget occurs if message topic is empty.
	ErrInvalidTarget = errors.New("topic is invalid or registration ids are not set")

	// ErrToManyRegIDs occurs when registration ids more then 1000.
	ErrToManyRegIDs = errors.New("too many registrations ids")

	// ErrInvalidTimeToLive occurs if TimeToLive more then 2419200.
	ErrInvalidTimeToLive = errors.New("messages time-to-live is invalid")
)

// Notification specifies the predefined, user-visible key-value pairs of the
// notification payload.
type Notification struct {
	// The notification's title.
	Title string `json:"title,omitempty"`

	// The notification's body text.
	Body string `json:"body,omitempty"`
}

// Message represents list of targets, options, and payload for HTTP JSON
// messages.
type Message struct {

	// The identifier of the message sent, in the format of projects/*/messages/{message_id}.
	Name string `json:"name,omitempty"`

	// Registration token to send a message to.
	Token string `json:"token,omitempty"`

	// Topic name to send a message to, e.g. "weather". Note: "/topics/" prefix should not be provided.
	Topic string `json:"topic,omitempty"`

	// Condition to send a message to, e.g. "'foo' in topics && 'bar' in topics".
	Condition string `json:"condition,omitempty"`

	// Apple Push Notification Service specific options.
	Apns map[string]string `json:"apns,omitempty"`

	// Webpush protocol options.
	Webpush map[string]string `json:"webpush,omitempty"`

	// Android specific options for messages sent through FCM connection server.
	Android *AndroidConfig `json:"android,omitempty"`

	// Basic notification template to use across all platforms.
	Notification *Notification `json:"notification,omitempty"`

	// Arbitrary key/value payload.
	// An object containing a list of "key": value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Data map[string]interface{} `json:"data,omitempty"`
}

// Validate returns an error if the message is not well-formed.
func (msg *Message) Validate() error {
	if msg == nil {
		return ErrInvalidMessage
	}

	if msg.Android != nil && msg.Android.TTL != "" {
		if _, err := time.ParseDuration(msg.Android.TTL); err != nil {
			return ErrInvalidTimeToLive
		}
	}

	return nil
}
