package slash

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {

	token := r.FormValue("token")
	if token == "YOUR_TOKEN" {

		command := r.FormValue("command")
		user := r.FormValue("user_name")
		text := r.FormValue("text")

		if command == "/webex" {

			var jsonStr = []byte(nil)

			if text == "help" {
				jsonStr = []byte(`{"response_type":"ephemeral", "text": "How to use /webex slash command", "attachments":[{"text": "/webex to use a link with your Slack username\n/webex [username] to use a link with specified username"}]}`)
			} else {
				if text != "" {
					user = text
				}
				jsonStr = []byte(`{"response_type":"in_channel", "text": "` + user + ` has set up a WebEx room to join", "attachments":[{"text":"https://companyName.webex.com/join/` + user + `","color":"#3AA3E3"}]}`)
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonStr)

		} else {
			fmt.Fprint(w, "I do not understand your command.")
		}
	}
}
