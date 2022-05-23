package t

import (
	"context"
	"errors"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/errgroup"
)

func TestWrapError(t *testing.T) {
	assert.False(t, true)
	err := errors.New("you are wrong")
	err = WrapError("oneWrap", err)
	err = WrapError("twoWrap", err)
	fmt.Print(err)
}

func TestGetErrorStack(t *testing.T) {
	assert.False(t, true)
	err := errors.New("I am an error")
	fmt.Println(GetErrorStack("real an error", err))
}

func TestErrgroup() {
	eg, ctx := errgroup.WithContext(context.Background())
	for i := 0; i < 100; i++ {
		i := i
		eg.Go(func() error {
			time.Sleep(2 * time.Second)
			select {
			case <-ctx.Done():
				fmt.Println("Canceled:", i)
				return nil
			default:
				fmt.Println("End:", i)
				return nil
			}
		})
	}
	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}
}
