# template-extract
A Library for extracting placeholders from the templates.

### Getting Library
Install the library with

```go get -u "github.com/MeowSaiGithub/template-extract```

Import the library

```go
import te "github.com/MeowSaiGithub/template-extract"
```
### How to use

```go
        t, err := template.New("test").Parse("Hello {{.Name}}")
	if err != nil {
		panic(err)
	}
	td, _ := te.NewTemplateDataExtractor(t)
	td.ExtractPlaceHolders()
	fmt.Println(td.GetRawData())
	fmt.Println(td.GetCleanData())
	fmt.Println(td.GetMapData())
```

Raw Data is every action placeholders such as ``{{.Name}}``, ``{{$name := .Name}}``.

Clean Data is only ``Name`` from both data of above. It removed duplicated data and remove un-necessary strings.

Map data is ``map["Name": ""]``. This can be used for json marshalling and unmarshalling the request data.

The original code snippet can be found here - https://stackoverflow.com/questions/40584612/how-to-get-a-map-or-list-of-template-actions-from-a-parsed-template.
and Credit to all those answers.