# gin-trace
trace middleware for gin

```	import "github.com/techjanitor/gin-trace"```

https://github.com/gin-gonic/gin

https://godoc.org/golang.org/x/net/trace



```
	router := gin.New()

    // Add the trace middleware to your router or group
	router.Use(trace.Trace())
	
    // Add the controller to view the default request handler
	router.GET("/debug/requests", trace.TraceController)
```
