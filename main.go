package main
import (
  "net/http"
  "fmt"
  "sync"
  "strings"
  "os/exec"
)

func generateMenu(w http.ResponseWriter, r *http.Request) {
  message := `
    <html>
      <head>
        <title>Docker Director</title>
        <style type="text/css">
        ol {
          margin-bottom: 2rem;
        }
        body > ol > li > span {
          font-weight: bold;
          font-size: 1.2rem;
        }

        body > ol > li {
          font-size: 1.3rem;
        }
        </style>
      </head>
      <body>
        <ol>
          <li><span>Sonarr</span>
            <ol>
              <li>Restart</li>
              </ol>
            </li>
          <li><span>Sabnzbd</span>
            <ol>
              <li>Restart</li>
              </ol>
            </li>
          <li><span>ruTorrent</span>
            <ol>
              <li>Restart</li>
              </ol>
            </li>
          <li><span>Plex</span>
            <ol>
              <li>Restart</li>
              </ol>
            </li>
        </ol>
      </body>
    </html>
  `
  w.Write([]byte(message))
}

func exe_cmd(cmd string, wg *sync.WaitGroup) {
  fmt.Println("command is ",cmd)
  // splitting head => g++ parts => rest of the command
  parts := strings.Fields(cmd)
  head := parts[0]
  parts = parts[1:len(parts)]

  out, err := exec.Command(head,parts...).Output()
  if err != nil {
    fmt.Printf("%s", err)
  }
  fmt.Printf("%s", out)
  wg.Done() // Need to signal to waitgroup that this goroutine is done
}

func main() {
  http.HandleFunc("/", generateMenu)
  if err := http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }
}
