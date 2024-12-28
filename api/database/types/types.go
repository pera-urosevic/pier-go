package types

type DatabaseIndex struct {
	Database   string `json:"database"`
	Collection string `json:"collection"`
	Facets     string `json:"facets"`
}

type AutocompleteFields struct {
	Selects []string `json:"selects"`
	Tags    []string `json:"tags"`
}

type Autocompletes = map[string][]string

type Doc = map[string]interface{}
