package main

import (
	"errors"
	"fmt"
	"io"
	"maps"
	"os"
	"slices"

	"github.com/dev-fritte/job-crawler/internal/domain"
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
	case "demo":
		runDemo()
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

func runDemo() {

	job := domain.Job{
		ExternalID:  "12345",
		Source:      "example.com",
		URL:         "https://example.com/job/12345",
		Title:       "Software Engineer",
		Description: "We are looking for a Software Engineer to join our team.",
		Company: domain.Company{
			Name:    "Example Inc.",
			Website: "https://exampleInc.com",
		},
		Salary: &domain.Salary{
			Currency: domain.CurrencyEUR,
			Period:   domain.SalaryPeriodYearly,
			Min:      60000,
			Max:      90000,
		},
	}

	job.AddTag("golang")

	job.AddTag("  go  ")
	job.AddTag("  go  ")

	fmt.Println("Has Tag", job.HasTag("goLang"))
	fmt.Printf("Job: %+v\n", job)

	fmt.Println("Validate:", job.Validate())
	fmt.Println("Validate empty job:", domain.Job{}.Validate())
}
