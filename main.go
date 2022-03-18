package main

import (
	"os"
	"time"
	"net/http"
	"github.com/lightstep/otel-launcher-go/launcher"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/trace"
)

var tracer trace.Tracer


func main() {
    otelLauncher := launcher.ConfigureOpentelemetry(
		launcher.WithServiceName("laci-lightstep"),
		launcher.WithAccessToken(os.Getenv("LS_ACCESS_TOKEN")),
	)
    defer otelLauncher.Shutdown()

	wrappedHandler := otelhttp.NewHandler(http.HandlerFunc(helloHandler), "/")
	http.Handle("/", wrappedHandler)
	http.ListenAndServe(":8080", nil)
}

// Example HTTP Handler
func helloHandler(w http.ResponseWriter, req *http.Request) {
	cxt := req.Context()
	span := trace.SpanFromContext(cxt)
	defer span.End()
	
	time.Sleep(time.Second)
	w.Write([]byte("Hello Lightstep!!!"))
}