package test

import (
	"firebase.google.com/go/v4/messaging"
	"github.com/huangyuhui-github/go-fcm"
	"testing"
)

func TestSample(t *testing.T) {

	messages := []*messaging.Message{
		{
			Notification: &messaging.Notification{
				Title: "FCM Title",
				Body:  "FCM Body",
			},
			Token: "cAy6n8YIRLKmhZF1x4e_D5:APA91bF1HakkzuIkJ4D_8nqqdoVX5Cg-o4jw37nUwHW6yVhrgni66NiiZbLIAN33h3CLcYY6Hq9erb40zcLAgzq0oEqQrQ_gSu7QutxqhjUwaSQY5GNYp_zOSc8Mejkb4MpzslzndY5D",
			Android: &messaging.AndroidConfig{
				Notification: &messaging.AndroidNotification{
					DefaultVibrateTimings: true,
				},
			},

			Data: map[string]string{
				"message": "这是透传消息",
			},
		},
	}

	client, err := fcm.NewClient("linen-marking-386905", "./credentials.json")
	if err != nil {
		t.Fatal(err)
	}

	// Send the message and receive the response without retries.
	resp, err := client.Send(messages)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v\n", resp.Responses)
}
