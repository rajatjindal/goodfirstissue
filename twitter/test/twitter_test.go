package test

import (
	"errors"
	"github.com/rajatjindal/goodfirstissue/twitter"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	// Given
	tokens := &twitter.Tokens{
		ConsumerKey:   "test_consumer_key",
		ConsumerToken: "test_consumer_token",
		Token:         "test_token",
		TokenSecret:   "test_token_secret",
	}

	// When
	client, err := twitter.NewClient(tokens)

	// Then
	assert.NoError(t, err, "NewClient returned an error")
	assert.NotNil(t, client, "NewClient returned a client with a nil twitter field")
}

func TestTweetSuccess(t *testing.T) {
	// Given
	tokens := &twitter.Tokens{
		ConsumerKey:   "your_consumer_key",
		ConsumerToken: "your_consumer_token",
		Token:         "your_token",
		TokenSecret:   "your_token_secret",
	}

	c, _ := twitter.NewClient(tokens)

	msg := "Hello, Twitter!"

	// When
	err := c.Tweet(msg)

	// Then
	assert.NoError(t, err, "Unexpected error")
}

func TestTweetError(t *testing.T) {
	// Given
	tokens := &twitter.Tokens{
		ConsumerKey:   "your_consumer_key",
		ConsumerToken: "your_consumer_token",
		Token:         "your_token",
		TokenSecret:   "your_token_secret",
	}

	c, _ := twitter.NewClient(tokens)

	msg := ""

	// When
	err := c.Tweet(msg)

	// Then
	assert.Error(t, err, "Expected an error")
	assert.True(t, errors.Is(err, twitter.ErrTweetFailed), "Error should be of type ErrTweetFailed")
}
