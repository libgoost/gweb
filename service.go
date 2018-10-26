package gweb

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"sync"
)

type GoostService struct {
	cfg  *Config
	pool *sync.Pool
	tmpl *template.Template
}

func NewService(cfg *Config) *GoostService {
	g := &GoostService{cfg: cfg}
	g.pool = &sync.Pool{New: func() interface{} { return &RepoInfo{} }}
	g.tmpl = template.Must(template.New("").Parse(tmpl))
	return g
}

func (g *GoostService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.Redirect(w, r, g.cfg.RepoRoot, http.StatusTemporaryRedirect)
		return
	}
	g.showRepo(w, r)
}

type RepoInfo struct {
	RealRepo, Path string
}

func (g *GoostService) newRepoInfo(r *http.Request) (ri *RepoInfo) {
	ri = g.pool.Get().(*RepoInfo)
	ri.Path = r.URL.Path[1:]
	ri.RealRepo = fmt.Sprintf("%s%s", g.cfg.RepoRoot,
		strings.Replace(ri.Path, "/", "-", -1))
	return
}

func (g *GoostService) showRepo(w http.ResponseWriter, r *http.Request) {
	ri := g.newRepoInfo(r)
	err := g.tmpl.Execute(w, ri)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	g.pool.Put(ri)
}

const tmpl = `<html>
<head>
<title>Goost {{.Path}}</title>
<meta name="go-import" content="goost.org/{{ .Path }} git {{ .RealRepo }}">
<meta name="go-source" content="goost.org/{{ .Path }} _ {{ .RealRepo}}/tree/master{/dir} {{ .RealRepo}}/blob/master{/dir}/{file}#L{line}">
</head>
<body>
	<h3>{{ .Path }}</h3>
	<p><a href="https://godoc.org/goost.org/{{.Path}}">godoc</a></p>
	<p><pre>go get goost.org/{{ .Path }}</pre></p>
</body>
</html>
`
