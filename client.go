package fcm

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"google.golang.org/api/option"
	"time"
)

const (
	// DefaultTimeout duration in second
	DefaultTimeout time.Duration = 30 * time.Second
)

// Client abstracts the interaction between the application server and the
// FCM server via HTTP protocol. The developer must obtain an API key from the
// Google APIs Console page and pass it to the `Client` so that it can
// perform authorized requests on the application server's behalf.
// To send a message to one or more devices use the Client's Send.
//
// If the `HTTP` field is nil, a zeroed http.Client will be allocated and used
// to send messages.
type Client struct {
	firebaseClient *messaging.Client
	timeout        time.Duration
}

// NewClient creates new Firebase Cloud Messaging Client based on API key and
// with default endpoint and http httpClient.
func NewClient(projectId string, credentialsFile string) (*Client, error) {

	ctx := context.Background()
	app, err := firebase.NewApp(ctx, &firebase.Config{ProjectID: projectId}, option.WithCredentialsFile(credentialsFile))
	if err != nil {
		return nil, err
	}

	firebaseClient, err := app.Messaging(ctx)
	if err != nil {
		return nil, err
	}

	c := &Client{
		firebaseClient: firebaseClient,
		timeout:        DefaultTimeout,
	}

	return c, nil
}

// SendWithContext sends a message to the FCM server without retrying in case of service
// unavailability. A non-nil error is returned if a non-recoverable error
// occurs (i.e. if the response status is not "200 OK").
// Behaves just like regular send, but uses external context.
func (c *Client) SendWithContext(ctx context.Context, messages []*messaging.Message) (*messaging.BatchResponse, error) {

	return c.send(ctx, messages)
}

// Send sends a message to the FCM server without retrying in case of service
// unavailability. A non-nil error is returned if a non-recoverable error
// occurs (i.e. if the response status is not "200 OK").
func (c *Client) Send(messages []*messaging.Message) (*messaging.BatchResponse, error) {

	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	return c.SendWithContext(ctx, messages)
}

// send sends a request.
func (c *Client) send(ctx context.Context, messages []*messaging.Message) (*messaging.BatchResponse, error) {

	return c.firebaseClient.SendEach(ctx, messages)
}
