package trace

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/trace"
	"net/http"
)

// Trace will gather information from the request and also add the trace methods to handlers
func Trace() gin.HandlerFunc {
	return func(c *gin.Context) {

		if gin.IsDebugging() {
			// a new trace
			tr := trace.New(c.HandlerName(), c.Request.URL.Path)

			c.Next()

			if len(c.Errors) != 0 {
				// loop through errors
				for _, err := range c.Errors {
					tr.LazyLog(fmt.Sprintf("Error: %s\nMeta: %s", err.Err, err.Meta), false)
				}
				tr.SetError()
			}

			// finish the trace
			tr.Finish()

			return
		}

		c.Next()

	}
}

// TraceController returns the default trace requests page
// example handler: r.GET("/debug/requests", trace.TraceController)
func TraceController(c *gin.Context) {

	if gin.IsDebugging() {
		// render the requests page
		trace.Render(c.Writer, c.Request, false)
	}

	c.String(http.StatusNotFound, "Not found")

	return
}
