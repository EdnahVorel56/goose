package goose

import (
	"context"
	"testing"
	"time"
)

func TestFetchMessage_PreCanceledContext(t *testing.T) {
	cg := NewConsumerGroup(10)
	msg := Message{Value: "test-message"}
	cg.Put(msg)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := cg.FetchMessage(ctx)
	if err != context.Canceled {
		t.Errorf("expected context.Canceled, got %v", err)
	}

	fetched, err := cg.FetchMessage(context.Background())
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if fetched.Value != "test-message" {
		t.Errorf("expected 'test-message', got %v", fetched.Value)
	}
}

func TestFetchMessage_TimeoutDuringFetch(t *testing.T) {
	cg := NewConsumerGroup(10)

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	_, err := cg.FetchMessage(ctx)
	if err != context.DeadlineExceeded {
		t.Errorf("expected context.DeadlineExceeded, got %v", err)
	}
}

func TestFetchMessage_ContextCanceledDuringFetch(t *testing.T) {
	cg := NewConsumerGroup(10)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(50 * time.Millisecond)
		cancel()
	}()

	_, err := cg.FetchMessage(ctx)
	if err != context.Canceled {
		t.Errorf("expected context.Canceled, got %v", err)
	}
}