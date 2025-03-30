package handlers

import (
    "encoding/json"
    "net/http"
    "strings"
)

func parseRequest(r *http.Request) (string, string, string, error) {
    e := r.ParseForm()
    if e != nil {
        return "", "", "", e
    }
    t := r.FormValue("text")
    token := r.FormValue("token")
    user := r.FormValue("user_id")
    return t, token, user, nil
}

func parseCmdArgs(text string) (string, []string) {
    parts := strings.Fields(text)
    if len(parts) == 0 {
        return "", nil
    }
    return strings.ToLower(parts[0]), parts[1:]
}

func respond(w http.ResponseWriter, text string, ephemeral bool) {
    rt := "in_channel"
    if ephemeral {
        rt = "ephemeral"
    }
    m := map[string]string{
        "response_type": rt,
        "text":          text,
    }
    b, _ := json.Marshal(m)
    w.Header().Set("Content-Type", "application/json")
    w.Write(b)
}
