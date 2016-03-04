package trace

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/trace"
)

// Trace will gather information from the request and also add the trace methods to handlers
func Trace() gin.HandlerFunc {
	return func(c *gin.Context) {

		// a new trace
		tr := trace.New(c.HandlerName(), c.Request.URL.Path)

		// access the trace methods from subsequent handlers
		c.Set("trace", tr)

		c.Next()

		// finish the trace
		tr.Finish()

	}
}

// TraceController returns the default trace requests page
// example handler: r.GET("/debug/requests", trace.RequestsController)
func TraceController(c *gin.Context) {
	// render the requests page
	trace.Render(c.Writer, c.Request, false)
	return
}
