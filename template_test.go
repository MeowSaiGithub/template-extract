package template_extract

import (
	"reflect"
	"testing"
	"text/template"
)

type TestCase struct {
	name          string
	templateStr   string
	expectedRaw   []string
	expectedClean []string
	expectedMap   map[string]interface{}
}

func TestT_ExtractArrayPlaceHolders(t *testing.T) {
	tcs := []TestCase{
		{
			name:          "Test 1",
			templateStr:   "Hello, {{.Name}}! Your number is {{.Number}}.",
			expectedRaw:   []string{"{{.Name}}", "{{.Number}}"},
			expectedClean: []string{"Name", "Number"},
			expectedMap:   map[string]interface{}{"Name": "", "Number": ""},
		},
		{
			name:          "Test 2",
			templateStr:   "Hello, {{.Name}}! Your number is {{.Number}}. {{.Name}} is again. {{.Number}}",
			expectedRaw:   []string{"{{.Name}}", "{{.Number}}", "{{.Name}}", "{{.Number}}"},
			expectedClean: []string{"Name", "Number"},
			expectedMap:   map[string]interface{}{"Name": "", "Number": ""},
		},
		{
			name:          "Test 3",
			templateStr:   "{{$Name := .Name}} {{$Number := .Number}} Hello, {{$Name}}! Your number is {{$Number}}.",
			expectedRaw:   []string{"{{$Name := .Name}}", "{{$Number := .Number}}", "{{$Name}}", "{{$Number}}"},
			expectedClean: []string{"Name", "Number"},
			expectedMap:   map[string]interface{}{"Name": "", "Number": ""},
		},
		{
			name:          "Test 4",
			templateStr:   "There is no placeholder here :(.",
			expectedRaw:   []string{},
			expectedClean: []string{},
			expectedMap:   map[string]interface{}{},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {

			tmpl, _ := template.New("test").Parse(tc.templateStr)
			testStruct, _ := NewTemplateDataExtractor(tmpl)
			testStruct.ExtractPlaceHolders()
			if !reflect.DeepEqual(testStruct.GetRawData(), tc.expectedRaw) {
				t.Errorf("expected rawPlaceHolders to be %v, got %v", tc.expectedRaw, testStruct.GetRawData())
			}
			if !reflect.DeepEqual(testStruct.GetCleanData(), tc.expectedClean) {
				t.Errorf("expected rawPlaceHolders to be %v, got %v", tc.expectedClean, testStruct.GetCleanData())
			}
			if !reflect.DeepEqual(testStruct.GetMapData(), tc.expectedMap) {
				t.Errorf("expected rawPlaceHolders to be %v, got %v", tc.expectedMap, testStruct.GetMapData())
			}
		})
	}
}
