package main
import (
  "net/http"
  "fmt"
  "sync"
  "strings"
  "os/exec"
)

func generateMenu(w http.ResponseWriter, r *http.Request) {
	msg := r.URL.Query().Get("msg")

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
				.msg {
					display: block;
					border-radius: 0.25rem;
					color: #FFF;
					background: green;
					padding: 1rem;
					text-align: center;
				}

        </style>
      </head>
      <body>
			<strong class="msg">` + msg + `</strong>
        <ol>
          <li><span>Sonarr</span>
            <ol>
              <li><a href="/reset?instance=sonarr">Restart</a></li>
              </ol>
            </li>
          <li><span>Sabnzbd</span>
            <ol>
              <li><a href="/reset?instance=sabnzbd">Restart</a></li>
              </ol>
            </li>
          <li><span>ruTorrent</span>
            <ol>
              <li><a href="/reset?instance=rutorrent">Restart</a></li>
              </ol>
            </li>
          <li><span>Plex</span>
            <ol>
              <li><a href="/reset?instance=plex">Restart</a></li>
              </ol>
            </li>
        </ol>
      </body>
    </html>
  `
  w.Write([]byte(message))
}

func resetAction(w http.ResponseWriter, r *http.Request) {
	// message := "Break my stride"
	// w.Write([]byte(message))
	// instance := r.URL.Query().Get("instance")
	out, err := exec.Command("sh","-c","ls").Output()



	if err != nil {
    s := fmt.Sprintf("%s", err)
  } else {
		s := fmt.Sprintf("%s", out)
		http.Redirect(w, r, `/?msg=` +  s + ` restarted successfully`, 301)
	}



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
	http.HandleFunc("/reset", resetAction)

  if err := http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }
}
