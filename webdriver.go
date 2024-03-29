package webdriver

import (
	"errors"
	"fmt"
	"net"
	"time"
)

// probe d.Port until get a reply or timeout is up
func probePort(port int, timeout time.Duration) error {
	address := fmt.Sprintf("127.0.0.1:%d", port)
	now := time.Now()
	for {
		if conn, err := net.Dial("tcp", address); err == nil {
			if err = conn.Close(); err != nil {
				return err
			}
			break
		}
		if time.Since(now) > timeout {
			return errors.New("start failed: timeout expired")
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}
