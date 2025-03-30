package handlers

import "fmt"

func handleDelete(cmd, userID string, args []string) string {
    if len(args) < 1 {
        return fmtDeleteUsage(cmd)
    }
    pid := args[0]
    e := deletePoll(userID, pid)
    if e != nil {
        return fmt.Sprintf("[%s] Error: %v", cmd, e)
    }
    return fmt.Sprintf("/votebote [%s] Poll %s deleted", cmd, pid)
}

func deletePoll(userID, pollID string) error {
    p, e := getPollByID(pollID)
    if e != nil {
        return e
    }
    if p.OwnerID != userID {
        return fmt.Errorf("only owner %s can delete poll", p.OwnerID)
    }
    _, err := replaceDelete(p.ID)
    return err
}

func fmtDeleteUsage(cmd string) string {
    return fmt.Sprintf("[%s] Usage: /votebot delete <pollID>", cmd)
}
