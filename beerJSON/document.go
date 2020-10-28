package beerJSON

import (
	"encoding/json"
	"fmt"
	"github.com/beerproto/beerjson.go"
	"github.com/go-playground/statics/static"
	"github.com/xeipuuv/gojsonschema"
	"go/build"
	"io/ioutil"
	"log"
	"path/filepath"
	"github.com/hashicorp/go-multierror"
)

type Document struct {
	Beer *beerjson.Beerjson `json:"beerjson"`
}


func Load(doc string) (*Document, error) {
	data, err := ioutil.ReadFile(doc)
	if err != nil {
		return nil, err
	}

	abs,err := filepath.Abs(doc)
	if err != nil {
		return nil, err
	}

	Validate(abs)

	beer := &beerjson.Beerjson{}
	str :=  &Document{ Beer: beer}

	err = json.Unmarshal(data, &str)
	if err != nil {
		return nil, err
	}

	return str, nil
}

func Validate(doc string) error {
	gopath := build.Default.GOPATH + ""
	pkgPath := "/schema"
	pkg := gopath + pkgPath

	config := &static.Config{
		UseStaticFiles: true,
		AbsPkgPath:     pkg,
		FallbackToDisk: true,
	}

	assets, err := newStaticAssets(config)
	if err != nil {
		log.Println(err)
	}

	schemaLoader := gojsonschema.NewReferenceLoaderFileSystem("file:///schema/beer.json", assets.FS())
	documentLoader := gojsonschema.NewReferenceLoader("file://" + doc)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return err
	}

	 if !result.Valid()  {
		for _, desc := range result.Errors() {
			err = multierror.Append(err, fmt.Errorf(desc.String()))
		}
	}
	return err
}