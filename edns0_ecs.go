package edns0_ecs

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Config struct {
	Prefix string
}

func CreateConfig() *Config {
	return &Config{
		Prefix: "EDNS0",
	}
}

type EdnsEcsPlugin struct {
	next   http.Handler
	name   string
	config *Config
	dnslog *log.Logger
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	dnslog := log.New(io.Discard, config.Prefix+": ", log.Lmsgprefix|log.Ldate|log.Ltime)
	dnslog.SetOutput(os.Stdout)

	dnslog.Println("EDNS0/ECS Plugin loaded")

	return &EdnsEcsPlugin{
		next:   next,
		name:   name,
		dnslog: dnslog,
	}, nil
}

func (a *EdnsEcsPlugin) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.dnslog.Println("Start of plugin handling the request")

	a.dnslog.Printf(
		"%s - %s - %s - %s - %s",
		time.Now().Format(time.RFC3339),
		r.Method,
		r.URL.Path,
		r.RemoteAddr,
		r.Proto)

	a.dnslog.Println("End of plugin handling the request")

	a.next.ServeHTTP(w, r)
}
