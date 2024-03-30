package template_extract

import (
	"fmt"
	"strings"
	"text/template"
	"text/template/parse"
)

type TemplateDataExtractor interface {
	// ExtractPlaceHolders processes the template and extract raw, clean and mapped placeholders.
	ExtractPlaceHolders()
	// GetRawData retrieves raw placeholders data.
	GetRawData() []string
	// GetCleanData retrieves clean placeholders data.
	GetCleanData() []string
	// GetMapData retrieves mapped placeholders data.
	GetMapData() map[string]interface{}

	listTemplateNodes(node parse.Node)
	clean()
	populatePlaceHoldersMap()
}

// templateDataExtractor encapsulates the process of extracting placeholders
// from a Go template and provides methods to manipulate and retrieve data.
type templateDataExtractor struct {
	tmpl              *template.Template
	rawPlaceHolders   []string
	cleanPlaceHolders []string
	placeHolderMap    map[string]interface{}
}

// NewTemplateDataExtractor creates a new templateDataExtractor struct with the provided template and returns
// TemplateDataExtractor interface and an error if parameter *template.Template is nil.
func NewTemplateDataExtractor(t *template.Template) (TemplateDataExtractor, error) {
	if t == nil {
		return nil, fmt.Errorf("nil template")
	}

	td := templateDataExtractor{
		tmpl:              t,
		rawPlaceHolders:   []string{},
		cleanPlaceHolders: []string{},
		placeHolderMap:    make(map[string]interface{})}

	return &td, nil
}

// ExtractPlaceHolders processes the template and extract raw, clean and mapped placeholders.
func (t *templateDataExtractor) ExtractPlaceHolders() {
	t.listTemplateNodes(t.tmpl.Tree.Root)
	t.clean()
	t.populatePlaceHoldersMap()
}

// listTemplateNodes recursively traverses the template nodes and extracts raw placeholders.
func (t *templateDataExtractor) listTemplateNodes(node parse.Node) {
	if node.Type() == parse.NodeAction {
		t.rawPlaceHolders = append(t.rawPlaceHolders, node.String())
	}
	if ln, ok := node.(*parse.ListNode); ok {
		for _, n := range ln.Nodes {
			t.listTemplateNodes(n)
		}
	}
}

// clean remove duplicated placeholders and extract actual placeholders.
func (t *templateDataExtractor) clean() {
	unique := make(map[string]bool)
	var cRaw []string
	for _, s := range t.rawPlaceHolders {
		if _, ok := unique[s]; !ok {
			unique[s] = true
			cRaw = append(cRaw, s)
		}
	}
	for _, s := range cRaw {
		sn := strings.Split(s, ".")
		if len(sn) == 2 {
			sr := strings.TrimSuffix(strings.TrimSpace(sn[1]), "}}")
			t.cleanPlaceHolders = append(t.cleanPlaceHolders, sr)
		}
	}
}

// populatePlaceHoldersMap creates a map with unique placeholders as keys.
func (t *templateDataExtractor) populatePlaceHoldersMap() {
	for _, r := range t.cleanPlaceHolders {
		t.placeHolderMap[r] = ""
	}
}

// GetRawData retrieves raw placeholders data.
func (t *templateDataExtractor) GetRawData() []string {
	return t.rawPlaceHolders
}

// GetCleanData retrieves clean placeholders data.
func (t *templateDataExtractor) GetCleanData() []string {
	return t.cleanPlaceHolders
}

// GetMapData retrieves mapped placeholder data.
func (t *templateDataExtractor) GetMapData() map[string]interface{} {
	return t.placeHolderMap
}
