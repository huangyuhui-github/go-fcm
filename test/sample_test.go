package test

import (
	"github.com/appleboy/go-fcm"
	"testing"
)

func TestSample(t *testing.T) {

	msg := &fcm.Message{
		Token: "cAy6n8YIRLKmhZF1x4e_D5:APA91bF1HakkzuIkJ4D_8nqqdoVX5Cg-o4jw37nUwHW6yVhrgni66NiiZbLIAN33h3CLcYY6Hq9erb40zcLAgzq0oEqQrQ_gSu7QutxqhjUwaSQY5GNYp_zOSc8Mejkb4MpzslzndY5D",
		Data: map[string]interface{}{
			"foo": "bar",
		},
		Notification: &fcm.Notification{
			Title: "FCM title message .",
			Body:  "FCM body message.",
		},
	}

	client, err := fcm.NewClient("linen-marking-386905", "./credentials.json")
	if err != nil {
		t.Fatal(err)
	}

	// Send the message and receive the response without retries.
	response, err := client.Send(msg)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v\n", response)
}
