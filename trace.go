package trace

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/trace"
)

// Trace the request and adds some helper methods to subsequent handlers
func Trace() gin.HandlerFunc {
	return func(c *gin.Context) {

		// a new trace
		tr := trace.New(c.HandlerName(), c.Request.URL.Path)

		c.Next()

		// finish the trace
		tr.Finish()

	}
}
