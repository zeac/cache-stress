package cachestress

import "github.com/gorilla/mux"
import "net/http"
import "html/template"
import "strconv"

func init() {
  r := mux.NewRouter();
  r.HandleFunc("/", indexPage)
  r.HandleFunc("/index.js", indexJavaScript)
  r.HandleFunc("/cache/{expire}/{size}/{seed}", cachedPage)
  http.Handle("/", r)
}

func indexPage(w http.ResponseWriter, r *http.Request) {
  t, _ := template.ParseFiles("index.html")

  data := struct {
    Title string
  }{
    Title: "Test",
  }
  t.Execute(w, data)
}

const CONTENT = "Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."

func cachedPage(w http.ResponseWriter, r *http.Request) {
  p := mux.Vars(r)
  expire := p["expire"]
  size, err := strconv.Atoi(p["size"])
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
  content := []byte(CONTENT)
  min := func(a int, b int) int {
    if a > b {
      return b
    }
    return a
  }

  w.Header().Set("cache-control", "private, max-age=" + expire)
  for size >= 0 {
    n := min (size, len(content))
    w.Write(content[:n])
    size = size - len(content)
  }
}

func indexJavaScript(w http.ResponseWriter, r *http.Request) {
  http.ServeFile(w, r, "index.js")
}

