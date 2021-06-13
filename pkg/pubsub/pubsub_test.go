package pubsub

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPublish(t *testing.T) {
	err := testService.Publish(context.TODO(), "channel", "hello world")
	assert.NoError(t, err)
}

func TestSubscribe(t *testing.T) {
	t.Parallel()
	channel := "testchannel"
	val := time.Now().String()
	reply := make(chan []byte)
	err := testService.Subscribe(context.TODO(), channel, reply)
	assert.NoError(t, err)

	err = testService.Publish(context.TODO(), channel, val)
	assert.NoError(t, err)
	t.Run("message", func(t *testing.T) {
		msg := <-reply
		fmt.Println("Recieved message: ", string(msg))
	})
}
