package handlers

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

//	"github.com/cloudfoundry-samples/test-app/helpers"
)

type Hello struct {
	Time time.Time
}

func (p *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	name,err :=  url.QueryUnescape(r.URL.Query().Get("name"))
	_ = err 
	styledTemplate.Execute(w, Body{Body: fmt.Sprintf(`
<div class="hello">
	Test App	
</div>

<div class="mid-color">Hello  %s this is a bit crap but getting there.</div>
<div class="bottom-color"></div>
    `,name)})
}
