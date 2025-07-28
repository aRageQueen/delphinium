# delphinium
a CLI tool that takes a target (domain, org name, IP, etc.), gathers open source events over time (breaches, DNS changes, WHOIS updates, blog/news mentions), and outputs a chronological timeline in Markdown

## Features

- Generate timelines for domains or org names
- Pulls data from:
  - WHOIS records (via WhoisXMLAPI)
  - Public breach info (coming soon)
  - Blog/news mentions (coming soon)
- Output to **Markdown** or **HTML** (coming soon)
- Sorts and formats events chronologically

## Structure
```python
delphinium/
├── delphinium.go           # CLI entrypoint
├── sources/
│   ├── whois.go            # WHOIS data retrieval
│   ├── breaches.go         # Breach history
│   └── news.go             # Blog/news mentions (RSS)
├── builder/
│   └── timeline.go         # Event merging and formatting
├── internal/
│   └── types.go            # Event struct and helpers
├── output/
│   └── markdown.go         # Markdown rendering
├── go.mod
└── README.md
```
