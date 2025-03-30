package handlers

import (
	"net/http"
)

func HandleCommand(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}
	text, token, userID, err := parseRequest(r)
	if err != nil {
		http.Error(w, "Bad form", http.StatusBadRequest)
		return
	}
	e := validateToken(token)
	if e != nil {
		http.Error(w, e.Error(), http.StatusUnauthorized)
		return
	}
	fullSlash := "/votebot " + text

	cmd, argList := parseCmdArgs(text)
	msg, eph := routeCommand(cmd, argList, userID, text)
	finalMsg := fullSlash + "\n" + msg
	respond(w, finalMsg, eph)
}
