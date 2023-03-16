# option

```
import "github.com/powerpuffpenguin/easygo/option"
```

go does not support optional parameters, but we can simulate optional parameters by using interfaces, but before go1.8, because templates are not supported, you have to manually write a lot of similar codes to simulate optional parameters every time, now through templates you only need Simply call Option to support optional parameters

```
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
```

Now you can call the New function with an optional parameter
```
func Example() {
	New()
	New(ServerPort(123))
	New(ServerAddr("0.0.0.0"))
	New(ServerPort(123))
	New(ServerAddr("0.0.0.0"), ServerPort(123))
}
```