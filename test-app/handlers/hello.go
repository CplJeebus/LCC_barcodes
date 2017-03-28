package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cloudfoundry-samples/test-app/helpers"
)

type Hello struct {
	Time time.Time
}

func (p *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	index, _ := helpers.FetchIndex()
	name,err :=  url.QueryUnescape(r.URL.Query().Get("name"))

	styledTemplate.Execute(w, Body{Body: fmt.Sprintf(`
<div class="hello">
	Cycle	
</div>

<div class="my-index">Cycle 1</div>

<div class="mid-color">This is stuff: %s</div>
<div class="bottom-color"></div>
    `, name})
}
