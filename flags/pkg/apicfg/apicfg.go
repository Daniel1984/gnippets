package apicfg

import (
	"errors"
	"flag"
	"net/http"
	"os"
)

type Flagset struct {
	Token    string
	Method   string
	Resource string
}

func Parse() (*Flagset, error) {
	//os.Args = []string{"cmd", "list"}

	if len(os.Args) < 2 {
		return nil, errors.New("api subcommand is required")
	}

	fs := &Flagset{}

	// Subcommands
	apiCommand := flag.NewFlagSet("api", flag.ExitOnError)

	// api subcommand flag pointers
	apiCommand.StringVar(&fs.Token, "token", "", "Api token (Required)")
	apiCommand.StringVar(&fs.Method, "method", http.MethodGet, "request method <GET|POST|PUT|PATCH|DELETE>, default is GET")
	apiCommand.StringVar(&fs.Resource, "resource", "", "api resource you want to call <users/|payments/|orders> (required)")

	switch os.Args[1] {
	case "api":
		apiCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		return nil, errors.New("wrong arguments")
	}

	if apiCommand.Parsed() {
		if fs.Token == "" {
			apiCommand.PrintDefaults()
			return nil, errors.New("wrong arguments")
		}

		if len(fs.Resource) == 0 { // function to check against valid req methods
			apiCommand.PrintDefaults()
			return nil, errors.New("wrong arguments")
		}
	}

	return fs, nil
}
