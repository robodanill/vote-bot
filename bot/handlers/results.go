package handlers

import "fmt"

func handleResults(cmd string, args []string) string {
    if len(args) < 1 {
        return fmtResultsUsage(cmd)
    }
    pid := args[0]
    p, e := getPollByID(pid)
    if e != nil {
        return fmt.Sprintf("[%s] Error: %v", cmd, e)
    }
    return formatResults(cmd, p.ID, p.Question, p.Options, p.Votes, p.IsActive)
}

func formatResults(cmd, id, question string, opts []string, votes map[string][]string, active bool) string {
    s := fmt.Sprintf("[%s] Poll %s\nQuestion: %s\n", cmd, id, question)
    c := map[string]int{}
    for k, arr := range votes {
        c[k] = len(arr)
    }
    for _, opt := range opts {
        s += fmt.Sprintf("- %s: %d\n", opt, c[opt])
    }
    if active {
        s += "Poll is active"
    } else {
        s += "Poll is ended"
    }
    return s
}

func fmtResultsUsage(cmd string) string {
    return fmt.Sprintf("[%s] Usage: /votebot results <pollID>", cmd)
}
