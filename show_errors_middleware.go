package traffic

import (
  "net/http"
  "runtime/debug"
  "html/template"
)

type ShowErrorsMiddleware struct {}

func (middleware ShowErrorsMiddleware) RenderError(w http.ResponseWriter, r *http.Request, err interface{}, stack []byte) {
  html := `
  <html>
    <head>
      <title>Traffic Panic</title>
      <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
      <style>
      html, body{ padding: 0; margin: 0; }
      h1 { background: #C52F24; color: white; padding: 2px 10px; margin: 0 0 10px 0; }
      #error { color: #C52F24; font-size: 18px; }
      #container { padding: 0 10px; }
      </style>
    </head>
  <body>
    <header>
      <h1>Error</h1>
    </header>

    <div id="container">
      <p id="error">{{ .Error }}</p>
      <pre id="stack">{{ .Stack }}</pre>
    </div>
  </body>
  </html>
  `

  data := struct {
    Error interface{}
    Stack string
  }{
    err,
    string(stack),
  }

  w.Header().Add("Content-Type", "text/html")
  tpl := template.Must(template.New("ErrorPage").Parse(html))
  tpl.Execute(w, data)
}

func (middleware ShowErrorsMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next NextMiddlewareFunc) (http.ResponseWriter, *http.Request) {
  defer func() {
    if err := recover(); r != nil {
      middleware.RenderError(w, r, err, debug.Stack())
    }
  }()

  if nextMiddleware := next(); nextMiddleware != nil {
    w, r = nextMiddleware.ServeHTTP(w, r, next)
  }

  return w, r
}
