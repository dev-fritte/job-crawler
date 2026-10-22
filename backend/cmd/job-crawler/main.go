package main

import (
	"errors"
	"fmt"
	"io"
	"maps"
	"os"
	"slices"
)

var version = "dev"
var paramsMap = map[string]string{
	"crawl":   "Crawl job postings from the web",
	"api":     "Run the API server",
	"worker":  "Run the worker to process jobs",
	"version": "Print the version",
	"help":    "Print this help message",
}

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)

		usage(os.Stderr)

		os.Exit(2)
	}
}

func run(args []string) error {
	if len(args) < 1 {
		return errors.New("missing command")
	}

	switch args[0] {
	case "crawl":
		fmt.Println("crawl: not yet implemented")
	case "api":
		fmt.Println("api: not yet implemented")
	case "worker":
		fmt.Println("worker: not yet implemented")
	case "version":
		fmt.Println("job-crawler:", version)
	case "help":
		usage(os.Stdout)
	default:
		return fmt.Errorf("unsupported command %q", args[0])
	}

	return nil
}

func usage(w io.Writer) {
	sortedKeys := slices.Sorted(maps.Keys(paramsMap))

	for _, paramKey := range sortedKeys {
		fmt.Fprintf(w, "%-8s %s\n", paramKey, paramsMap[paramKey])
	}
}
