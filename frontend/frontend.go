package frontend

import (
	"embed"
	"io/fs"
	"net/http"
	"path/filepath"
	"strings"
)

//go:embed build/**
var buildDir embed.FS

func Handler() http.Handler {
	stripped, err := fs.Sub(buildDir, "build")
	if err != nil {
		panic(err)
	}

	fileServer := http.FileServer(http.FS(stripped))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/app")

		if path == "" || path == "/" {
			w.Header().Set("Content-Type", "text/html")
			index, err := fs.ReadFile(stripped, "index.html")
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			w.Write(index)
			return
		}

		// Serve assets from the build directory
		if strings.HasPrefix(path, "/_app/") || filepath.Ext(path) != "" {
			r.URL.Path = path
			fileServer.ServeHTTP(w, r)
			return
		}

		// For all other paths, serve index.html to let SvelteKit handle routing
		w.Header().Set("Content-Type", "text/html")
		index, err := fs.ReadFile(stripped, "index.html")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.Write(index)
	})
}
