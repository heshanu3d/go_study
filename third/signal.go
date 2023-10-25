package third

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func CaptureCtrlC() {
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	defer func() {
		fmt.Println("ctrl c interrupt")
		// signal.Stop(c)
		cancel()
	}()
	go func() {
		select {
		case <-c:
			cancel()
		case <-ctx.Done():
		}
	}()
	fmt.Println("press ctrl c in 5s")
	time.Sleep(5 * time.Second)
	fmt.Println("didn't capture ctrl c in 5s, now exit")
}