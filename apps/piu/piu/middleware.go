package piu

import (
	"fmt"
	"time"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		// Start timer
		t := time.Now()
		// Process request
		c.Next()
		// Calculate resolution time
		fmt.Printf("[%d] %s in %v \n", c.StatusCode, c.Request.RequestURI, time.Since(t))
	}
}
