package stoplightelements

import (
	"bytes"
	"embed"
	_ "embed"
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"

	"github.com/pb33f/libopenapi"
)

//go:embed template.html
var Templates embed.FS

// TemplateData is the data used to render the HTML template
type TemplateData struct {
	Title             string // Title of the page
	ApiDescriptionUrl string // URL to the OpenAPI specification
}

// HtmlFromSpec generates the HTML page for a Stoplight Elements viewer of
// the OpenAPI specification at the given path
func HtmlFromSpec(specPath string, apiDescriptionUrl string) []byte {
	spec, err := os.ReadFile(specPath)
	if err != nil {
		log.Fatal(err)
	}
	document, err := libopenapi.NewDocument(spec)
	if err != nil {
		log.Fatal(err)
	}

	var data TemplateData
	data.ApiDescriptionUrl = apiDescriptionUrl

	// build model depending on OpenAPI version
	// and fill necessary template data
	version := document.GetVersion()
	if strings.HasPrefix(version, "3.") {
		docModel, errors := document.BuildV3Model()
		checkModelErrors(errors)
		data.Title = docModel.Model.Info.Title
	} else {
		docModel, errors := document.BuildV2Model()
		checkModelErrors(errors)
		data.Title = docModel.Model.Info.Title
	}

	t, err := template.ParseFS(Templates, "template.html")
	if err != nil {
		log.Fatal(err)
	}
	buf := bytes.Buffer{}
	t.Execute(&buf, data)

	return buf.Bytes()
}

func checkModelErrors(errors []error) {
	if len(errors) > 0 {
		for i := range errors {
			fmt.Println(errors[i])
		}
		log.Fatal("Could not create model from specification")
	}
}
