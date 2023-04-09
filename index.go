package main

import (
	"net/http"
	"text/template"

	"github.com/uptrace/bunrouter"
)

var indexTmpl = `
<html>
  <h1>Welcome</h1>
  <ul>
    <li><a href="/products">Products List</a></li>
	<li><a href="/products/new">New Product</a></li>
  </ul>
</html>
`

func indexTemplate() *template.Template {
	return template.Must(template.New("index").Parse(indexTmpl))
}

func indexHandler(w http.ResponseWriter, req bunrouter.Request) error {
	return indexTemplate().Execute(w, nil)
}
