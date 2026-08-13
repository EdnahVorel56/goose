package goose

import (
	"context"
	"io"
)

type Message struct {
	Value string
}

type ConsumerGroup struct {
	messages chan Message
}

func NewConsumerGroup(bufferSize int) *ConsumerGroup {
	return &ConsumerGroup{
		messages: make(chan Message, bufferSize),
	}
}

func (c *ConsumerGroup) FetchMessage(ctx context.Context) (Message, error) {
	if err := ctx.Err(); err != nil {
		return Message{}, err
	}

	select {
	case <-ctx.Done():
		return Message{}, ctx.Err()
	default:
	}

	select {
	case <-ctx.Done():
		return Message{}, ctx.Err()
	case msg, ok := <-c.messages:
		if !ok {
			return Message{}, io.EOF
		}
		if err := ctx.Err(); err != nil {
			c.requeue(msg)
			return Message{}, err
		}
		return msg, nil
	}
}

func (c *ConsumerGroup) requeue(msg Message) {
	select {
	case c.messages <- msg:
	default:
	}
}

func (c *ConsumerGroup) Put(msg Message) {
	c.messages <- msg
}

func (c *ConsumerGroup) Close() {
	close(c.messages)
}