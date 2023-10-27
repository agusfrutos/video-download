package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			url := r.FormValue("url")
			cmd := exec.Command("youtube-dl", "-x", "--audio-format", "mp3", url)
			stdout, err := cmd.StdoutPipe()
			if err != nil {
				fmt.Fprintln(w, err)
				return
			}
			if err := cmd.Start(); err != nil {
				fmt.Fprintln(w, err)
				return
			}
			if _, err := io.Copy(os.Stdout, stdout); err != nil {
				fmt.Fprintln(w, err)
				return
			}
			if err := cmd.Wait(); err != nil {
				fmt.Fprintln(w, err)
				return
			}
		}
		fmt.Fprintf(w, `
<!DOCTYPE html>
<html>
<head>
    <title>Descargador de videos en línea</title>
    <script src="https://unpkg.com/htmx.org/dist/htmx.js"></script>
</head>
<body>
    <h1>Descargador de videos en línea</h1>
    <form hx-post="/" hx-target="#result">
        <label for="url">URL del video:</label><br>
        <input type="text" id="url" name="url"><br><br>
        <button type="submit">Descargar</button>
    </form>
    <div id="result"></div>
</body>
</html>
`)
	})

	http.ListenAndServe(":8080", nil)
}
