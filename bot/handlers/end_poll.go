package handlers

import "fmt"

func handleEnd(cmd, userID string, args []string) string {
    if len(args) < 1 {
        return fmtEndUsage(cmd)
    }

    pid := args[0]
    e := endPoll(userID, pid)
    if e != nil {
        return fmt.Sprintf("[%s] Error: %v", cmd, e)
    }

    return fmt.Sprintf("[%s] Poll %s ended", cmd, pid)
}

func endPoll(userID, pollID string) error {
    p, e := getPollByID(pollID)
    if e != nil {
        return e
    }

    if p.OwnerID != userID {
        return fmt.Errorf("only owner %s can end poll", p.OwnerID)
    }

    if !p.IsActive {
        return fmt.Errorf("poll already ended")
    }

    p.IsActive = false
    _, err := replacePoll(p)
	
    return err
}

func fmtEndUsage(cmd string) string {
    return fmt.Sprintf("[%s] Usage: /votebot end <pollID>", cmd)
}
