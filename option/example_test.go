package option_test

import (
	"github.com/powerpuffpenguin/easygo/option"
)

// Special template
type ServerOption option.Option[serverOptions]

// Defines the properties that optional parameters contain
type serverOptions struct {
	port int
	addr string
}

// Specify default values for parameters
var defaultServerOptions = serverOptions{
	port: 80,
	addr: "127.0.0.1",
}

// Create an optional parameter specifying the port
func ServerPort(port int) option.Option[serverOptions] {
	return option.New(func(opts *serverOptions) {
		opts.port = port
	})
}

// Create an optional parameter specifying the addr
func ServerAddr(addr string) option.Option[serverOptions] {
	return option.New(func(opts *serverOptions) {
		opts.addr = addr
	})
}

type Server struct {
	opts serverOptions
}

// Functions that support optional parameters
func New(opt ...ServerOption) *Server {
	// Set parameters to default values
	opts := defaultServerOptions
	for _, o := range opt {
		o.Apply(&opts) // Set incoming parameters
	}
	return &Server{
		opts: opts, // use parameters
	}
}

func Example() {
	New()
	New(ServerPort(123))
	New(ServerAddr("0.0.0.0"))
	New(ServerPort(123))
	New(ServerAddr("0.0.0.0"), ServerPort(123))
}
