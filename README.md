# jscope

Lightweight JavaScript reconnaissance tool written in Go.

## Overview

jscope is a reconnaissance utility designed to collect and analyze JavaScript resources from web applications.

The primary goal is to build a modular framework for JavaScript reconnaissance, endpoint discovery, and asset analysis.

Current implementation focuses on:

- downloading target pages
- parsing HTML
- extracting JavaScript assets
- building a foundation for future analysis modules

---

## Current Status

Version: v0.1.0

Current capabilities:

- Download HTML from a target website
- Parse HTML content
- Extract JavaScript file URLs from script tags
- Display discovered JavaScript assets

Project status:

- Go module initialized
- Git repository initialized
- GitHub repository connected
- Fetcher implemented
- HTML parser implemented
- First successful scans completed

---

## Architecture

Current project structure:

```text
jscope/
│
├── cmd/
│   └── jscope/
│       └── main.go
│
├── internal/
│   ├── crawler/
│   │   └── crawler.go
│   │
│   ├── fetcher/
│   │   └── fetcher.go
│   │
│   └── types/
│       └── types.go
│
├── go.mod
├── go.sum
└── README.md
```

---

## Components

### main

Responsibilities:

- command line argument handling
- execution pipeline
- output presentation

Pipeline:

```text
Target
 ↓
Fetcher
 ↓
Crawler
 ↓
Results
```

---

### fetcher

Responsibilities:

- perform HTTP requests
- download HTML pages
- future JavaScript downloading

API:

```go
func Download(url string) (string, error)
```

Flow:

```text
URL
 ↓
HTTP GET
 ↓
Response
 ↓
HTML Content
```

---

### crawler

Responsibilities:

- parse HTML
- locate script tags
- extract JavaScript URLs

API:

```go
func GetScriptURLs(content string) ([]string, error)
```

Input:

```html
/app.jsscript>
/vendor.jsscript>
```

Output:

```text
/app.js
/vendor.js
```

---

### types

Shared data structures used across the project.

Currently reserved for future expansion.

Planned structures:

```go
type JSFile struct {}
type Finding struct {}
type ScanResult struct {}
```

---

## Current Workflow

Current execution flow:

```text
Target URL
      ↓
fetcher.Download()
      ↓
HTML
      ↓
crawler.GetScriptURLs()
      ↓
Script URLs
      ↓
Console Output
```

---

## Example Usage

Run:

```bash
go run ./cmd/jscope github.com
```

Example output:

```text
[*] Downloaded 576665 bytes
[*] Found 8 scripts

https://github.githubassets.com/assets/high-contrast-cookie.js
https://github.githubassets.com/assets/environment.js
https://github.githubassets.com/assets/github-elements.js
...
```

---

## Design Philosophy

The project follows several principles:

### Simple

Keep dependencies minimal.

### Modular

Each package should perform one task only.

Example:

```text
fetcher
 ↓
download data

crawler
 ↓
extract data

extractor
 ↓
analyze data
```

### Extensible

New functionality should be added through isolated modules.

---

## Development Roadmap

### Stage 1

Download discovered JavaScript files.

Pipeline:

```text
Site
 ↓
HTML
 ↓
Scripts
 ↓
Download JS
```

---

### Stage 2

Endpoint extraction.

Examples:

```text
/api/user
/api/login
/api/admin
/graphql
```

Goal:

Extract hidden backend endpoints referenced inside JavaScript.

---

### Stage 3

URL extraction.

Examples:

```text
https://api.target.com
https://cdn.target.com
wss://chat.target.com
```

Goal:

Discover additional attack surface.

---

### Stage 4

Interesting keyword detection.

Keywords:

```text
admin
debug
internal
staging
production
test
graphql
token
auth
secret
```

Goal:

Locate potentially sensitive functionality.

---

### Stage 5

Secret discovery.

Potential findings:

```text
API_KEY
TOKEN
JWT
SECRET
CLIENT_ID
```

Goal:

Identify accidentally exposed 
