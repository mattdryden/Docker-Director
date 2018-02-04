package main
import (
  "net/http"
  "fmt"
  "sync"
  "strings"
  "os/exec"
)

func viewAction(w http.ResponseWriter, r *http.Request) {
	instance := r.URL.Query().Get("instance")
	url := "";
	switch instance {
		case "plex":
			url = "http://192.168.1.37:32400"
		case "sonarr":
			url = "http://192.168.1.37:8989"
		case "rutorrent":
			url = "http://192.168.1.37"
		case "sabnzbd":
			url = "http://192.168.1.37:8080"
		default:
			url = "/"
		}
		if(url != "") {
			http.Redirect(w, r, url, 302)
		}
}

func generateMenu(w http.ResponseWriter, r *http.Request) {
	msg := r.URL.Query().Get("msg")
	err := r.URL.Query().Get("err")
	response := ``

	if (msg != "") {
		response = `<strong class="msg">` + msg + `</strong>`
	} else if(err != "") {
		response = `<strong class="err">` + err + `</strong>`
	}

  message := `
    <html>
      <head>
        <title>Docker Director</title>
        <style type="text/css">
				/* http://meyerweb.com/eric/tools/css/reset/
				   v2.0 | 20110126
				   License: none (public domain)
				*/

				html, body, div, span, applet, object, iframe,
				h1, h2, h3, h4, h5, h6, p, blockquote, pre,
				a, abbr, acronym, address, big, cite, code,
				del, dfn, em, img, ins, kbd, q, s, samp,
				small, strike, strong, sub, sup, tt, var,
				b, u, i, center,
				dl, dt, dd, ol, ul, li,
				fieldset, form, label, legend,
				table, caption, tbody, tfoot, thead, tr, th, td,
				article, aside, canvas, details, embed,
				figure, figcaption, footer, header, hgroup,
				menu, nav, output, ruby, section, summary,
				time, mark, audio, video {
					margin: 0;
					padding: 0;
					border: 0;
					font-size: 100%;
					font: inherit;
					vertical-align: baseline;
				}
				/* HTML5 display-role reset for older browsers */
				article, aside, details, figcaption, figure,
				footer, header, hgroup, menu, nav, section {
					display: block;
				}
				body {
					line-height: 1;
				}
				ol, ul {
					list-style: none;
				}
				blockquote, q {
					quotes: none;
				}
				blockquote:before, blockquote:after,
				q:before, q:after {
					content: '';
					content: none;
				}
				table {
					border-collapse: collapse;
					border-spacing: 0;
				}
        ol {
          margin-bottom: 2rem;
        }
				body {
					margin: 1rem;
				}
        body > ol > li > span {
          font-weight: bold;
          font-size: 1.2rem;
					margin-bottom: .5rem;
					display: block;
        }
        body > ol > li {
          font-size: 1.2rem;
					display: block;
        }
				.msg, .err {
					display: block;
					border-radius: 0.25rem;
					color: #FFF;
					padding: 1rem;
					text-align: center;
					margin-bottom: 1rem;
				}
				.msg {
					background: green;
				}
				.err {
					background: red;
				}

        </style>
      </head>
      <body>
			` + response + `
        <ol>
          <li><span>Sonarr</span>
            <ol>
						<li><a href="/reset?instance=sonarr">Restart</a></li>
						<li><a href="/view?instance=sonarr">View</a></li>
              </ol>
            </li>
          <li><span>Sabnzbd</span>
            <ol>
              <li><a href="/reset?instance=sabnzbd">Restart</a></li>
							<li><a href="/view?instance=sabnzbd">View</a></li>
              </ol>
            </li>
          <li><span>ruTorrent</span>
            <ol>
              <li><a href="/reset?instance=rutorrent">Restart</a></li>
							<li><a href="/view?instance=rutorrent">View</a></li>
              </ol>
            </li>
          <li><span>Plex</span>
            <ol>
              <li><a href="/reset?instance=plex">Restart</a></li>
							<li><a href="/view?instance=plex">View</a></li>
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
	instance := r.URL.Query().Get("instance")
	out, err := exec.Command("sh","-c","docker restart " + instance).Output()



	if err != nil {
    // s := fmt.Sprintf("%s", err)
		http.Redirect(w, r, `/?err=Failed to restart`, 302)
  } else {
		s := fmt.Sprintf("%s", out)
		http.Redirect(w, r, `/?msg=Success:` + s, 302)
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
	http.HandleFunc("/view", viewAction)

  if err := http.ListenAndServe(":999", nil); err != nil {
    panic(err)
  }
}
