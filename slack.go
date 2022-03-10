package p

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/Code-Hex/takopi"
	"github.com/nlopes/slack"
)

var verificationToken string

// グローバル変数を定義して、init で取得処理を書くと次回の関数呼び出し時にも再利用される。
// https://cloud.google.com/functions/docs/bestpractices/tips#use_global_variables_to_reuse_objects_in_future_invocations
// https://qiita.com/takasp/items/368d9da90520c3ea48dc
func init() {
	verificationToken = os.Getenv("VERIFICATION_TOKEN")
}

func TakopiCommand(w http.ResponseWriter, r *http.Request) {
	s, err := slack.SlashCommandParse(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !s.ValidateToken(verificationToken) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	switch s.Command {
	case "/takopi":
		params := &slack.Msg{
			ResponseType: "in_channel",
			Text:         takopi.Do(strings.TrimSpace(s.Text)),
		}
		b, err := json.Marshal(params)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	default:
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
