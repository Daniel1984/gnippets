package cfg

import (
	"errors"
	"flag"
	"net/http"
	"net/url"
	"os"
)

type Flagset struct {
	Target   string
	Method   string
	Attempts int
}

func Parse() (*Flagset, error) {
	if len(os.Args) < 2 {
		return nil, errors.New("'atack' command is required")
	}

	fs := &Flagset{}

	// Subcommands
	atackCommand := flag.NewFlagSet("atack", flag.ExitOnError)

	// atack subcommand flag pointers
	atackCommand.StringVar(&fs.Target, "target", "", "Website URL / API endpoint you want to target")
	atackCommand.StringVar(&fs.Method, "method", http.MethodGet, "request method <GET|POST>, default is GET")
	atackCommand.IntVar(&fs.Attempts, "attempts", 1, "number of times you want to atack the Target")

	switch os.Args[1] {
	case "atack":
		atackCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		return nil, errors.New("wrong arguments")
	}

	if atackCommand.Parsed() {
		if fs.Target == "" {
			atackCommand.PrintDefaults()
			return nil, errors.New("wrong arguments")
		}

		u, err := url.Parse(fs.Target)
		if err != nil || len(u.Host) == 0 {
			atackCommand.PrintDefaults()
			return nil, errors.New("wrong arguments")
		}

		if fs.Attempts < 0 {
			atackCommand.PrintDefaults()
			return nil, errors.New("wrong arguments")
		}

		if fs.Method == "" || fs.Method != http.MethodGet && fs.Method != http.MethodPost {
			atackCommand.PrintDefaults()
			return nil, errors.New("wrong arguments")
		}
	}

	return fs, nil
}
