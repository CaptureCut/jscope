package types

type JSFile struct {
	URL     string
	Content string
}

type Finding struct {
	Type  string
	Value string
}

type ScanResult struct {
	Target   string
	JSFiles  []JSFile
	Findings []Finding
}