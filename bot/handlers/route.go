package handlers

import (
	"fmt"
	"strings"
	"vote_bot/config"
)

func routeCommand(cmd string, args []string, userID, text string) (string, bool) {
	switch cmd {
	case "create":
		return handleCreate(cmd, userID, text), false
	case "vote":
		return handleVote(cmd, userID, args, text), false
	case "results":
		return handleResults(cmd, args), false
	case "end":
		return handleEnd(cmd, userID, args), false
	case "delete":
		return handleDelete(cmd, userID, args), false
	case "help":
		return handleHelp(cmd), true
	default:
		return "[votebot] Unknown command. Use /votebot help", true
	}
}

func validateToken(token string) error {
	if config.MattermostToken != "" && token != config.MattermostToken {
		return fmt.Errorf("unauthorized")
	}

	return nil
}

func safeCut(s, prefix string) string {
	x := len(prefix)
	if len(s) <= x {
		return ""
	}
	
	return strings.TrimSpace(s[x:])
}
