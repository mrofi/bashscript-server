package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Read environment variables
	scriptsDir := os.Getenv("SCRIPT_DIR")
	if scriptsDir == "" {
		scriptsDir = "/app/scripts"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// HTTP handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		scriptName := strings.TrimPrefix(r.URL.Path, "/")

		if scriptName == "" {
			http.Error(w,
				"Usage: curl -sL https://<host>/<script>.sh | bash -s -- [args]",
				http.StatusBadRequest)
			return
		}

		if !strings.HasSuffix(scriptName, ".sh") {
			scriptName += ".sh"
		}

		scriptPath := filepath.Join(scriptsDir, scriptName)

		if _, err := os.Stat(scriptPath); os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}

		w.Header().Set("Content-Type", "text/x-shellscript")
		http.ServeFile(w, r, scriptPath)

		log.Printf("Served script: %s from %s\n", scriptName, r.RemoteAddr)
	})

	log.Printf("🚀 bashscript-server serving %s on port %s\n", scriptsDir, port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
