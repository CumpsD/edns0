# Log UDP requests

ChatGPT generated code.

## Can you write a traefik middleware plugin which logs UDP requests?

```go
package main

import (
    "fmt"
    "log"
    "net"
    "net/http"
    "time"

    "github.com/traefik/traefik/v2/pkg/plugins"
)

type UDPLogger struct {
    next http.Handler
}

func NewUDPLogger(next http.Handler) (http.Handler, error) {
    return &UDPLogger{next: next}, nil
}

func (u *UDPLogger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // Log the request using UDP
    logMessage := fmt.Sprintf("%s - %s - %s - %s", time.Now().Format(time.RFC3339), r.Method, r.URL.Path, r.RemoteAddr)
    err := sendLogMessageUDP(logMessage)
    if err != nil {
        log.Printf("Failed to log request via UDP: %s", err)
    }

    // Pass the request to the next middleware or handler
    u.next.ServeHTTP(w, r)
}

func sendLogMessageUDP(message string) error {
    conn, err := net.Dial("udp", "127.0.0.1:514") // Change the address and port to your preferred logging destination
    if err != nil {
        return err
    }
    defer conn.Close()

    _, err = fmt.Fprint(conn, message)
    if err != nil {
        return err
    }

    return nil
}

func main() {
    // Register the UDPLogger middleware as a Traefik plugin
    plugins.Register(&plugins.Plugin{
        Name: "UDPLogger",
        // Use the MiddlewareFunc to create instances of the UDPLogger middleware
        MiddlewareFunc: func(next http.Handler, config *plugins.Configuration, name string) (http.Handler, error) {
            return NewUDPLogger(next)
            },
    })
}

```

## Can you change the sendLogMessageUDP to write to Stdout?


```go
package main

import (
    "fmt"
    "log"
    "net/http"
    "time"

    "github.com/traefik/traefik/v2/pkg/plugins"
)

type UDPLogger struct {
    next http.Handler
}

func NewUDPLogger(next http.Handler) (http.Handler, error) {
    return &UDPLogger{next: next}, nil
}

func (u *UDPLogger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // Log the request to stdout
    logMessage := fmt.Sprintf("%s - %s - %s - %s", time.Now().Format(time.RFC3339), r.Method, r.URL.Path, r.RemoteAddr)
    log.Println(logMessage)

    // Pass the request to the next middleware or handler
    u.next.ServeHTTP(w, r)
}

func main() {
    // Register the UDPLogger middleware as a Traefik plugin
    plugins.Register(&plugins.Plugin{
        Name: "UDPLogger",
        // Use the MiddlewareFunc to create instances of the UDPLogger middleware
        MiddlewareFunc: func(next http.Handler, config *plugins.Configuration, name string) (http.Handler, error) {
            return NewUDPLogger(next)
        },
    })
}
```
