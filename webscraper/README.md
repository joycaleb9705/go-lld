# Concurrent URL Meta Fetcher

A simple, concurrent web scraping tool to practice low level design, concurrency primitives, and error handling.

## Project Spec
Objective: Build a CLI tool that takes a list of URLs, fetches the HTML for each concurrently, parses the `<title>` tag from the HTML, and returns the results.

### Functional Requirements
1. Input: The tool should accept a predefined slice of URLs in `main.go`.
2. Concurrent Fetching: Fetch all URLS simultaneously using goroutines.
3. Data Parsing: For each fetched HTML body, extract content within `<title>...</title>` tags (proper HTML parser).
4. Results Aggregation: Collect all results into a slice and print them at the end.
6. Output: For each URL, print the URL, the Title found, and the time taken to fetch it.

### Technical Details
1. Concurrency
- Use `sync.WaitGroup` to wait for all fetching goroutines to complete.
- Use a buffered channel to collect results safely from goroutines.

2. Networking / Parsing
- Use `net/http` for fetching.
- Implement `context.WithTimeout` to ensure a slow server doesn't hang the tool.
- Use `golang.org/x/net/html` to parse the HTML and find the `<title>` tag.