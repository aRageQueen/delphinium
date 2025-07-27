# delphinium
a CLI tool that takes a target (domain, org name, IP, etc.), gathers open source events over time (breaches, DNS changes, WHOIS updates, blog/news mentions), and outputs a chronological timeline in Markdown or HTML

## Structure
```python
delphinium/
├── delphinium.go           # CLI entrypoint
├── sources/
│   ├── whois.go            # WHOIS data retrieval
│   ├── breaches.go         # Breach history
│   └── news.go             # Blog/news mentions (RSS)
├── builder/
│   └── delphinium.go       # Event merging and formatting
├── internal/
│   └── types.go            # Event struct and helpers
├── output/
│   └── markdown.go         # Markdown rendering
├── go.mod
└── README.md
```