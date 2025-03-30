package handlers

import (
    "fmt"
    "strings"

    "vote_bot/db"
)

func handleVote(cmd, userID string, args []string, text string) string {
    if len(args) < 1 {
        return fmtVoteUsage(cmd)
    }

    pid := args[0]
    opt := parseVoteOption(text, cmd, pid)
    if opt == "" {
        return fmtVoteUsage(cmd)
    }

    e := votePoll(userID, pid, opt)
    if e != nil {
        return fmt.Sprintf("[%s] Error: %v", cmd, e)
    }

    return fmt.Sprintf("[%s] Voted for %s in poll %s", cmd, opt, pid)
}

func parseVoteOption(text, cmd, pid string) string {
    t := safeCut(text, cmd)
    t = strings.TrimSpace(t)
    t = strings.TrimPrefix(t, pid)
    t = strings.TrimSpace(t)
    a := parseQuotedArgs(t)
    if len(a) < 1 {
        return ""
    }

    return a[0]
}

func votePoll(userID, pollID, option string) error {
    p, e := getPollByID(pollID)
    if e != nil {
        return e
    }

    if !p.IsActive {
        return fmt.Errorf("poll not active")
    }

    f := false
    for _, o := range p.Options {
        if o == option {
            f = true
            break
        }
    }
    if !f {
        return fmt.Errorf("option not found: %s", option)
    }

    p.Votes[option] = append(p.Votes[option], userID)
    _, err := db.Conn.Replace("polls", []interface{}{p.ID, p.OwnerID, p.Question, p.Options, p.Votes, p.IsActive})
	
    return err
}

func fmtVoteUsage(cmd string) string {
    return fmt.Sprintf("[%s] Usage: /votebot vote <pollID> \"Option\"", cmd)
}
