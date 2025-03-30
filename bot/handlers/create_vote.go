package handlers

import (
	"fmt"

	"vote_bot/db"
	"vote_bot/poll"

	"github.com/google/uuid"
)

func handleCreate(cmd, userID, text string) string {
	sub := safeCut(text, cmd)
	a := parseQuotedArgs(sub)
	if len(a) < 2 {
		return fmtCreateError(cmd)
	}

	q := a[0]
	opts := a[1:]
	id, err := createPoll(userID, q, opts)
	if err != nil {
		return fmt.Sprintf("[%s] Error: %v", cmd, err)
	}

	return fmt.Sprintf("[%s] Created poll\nID: %s\nQuestion: %s\nOptions: %v", cmd, id, q, opts)
}

func createPoll(ownerID, question string, options []string) (string, error) {
	id := uuid.New().String()
	p := poll.Poll{ID: id, OwnerID: ownerID, Question: question, Options: options, Votes: map[string][]string{}, IsActive: true}
	
	_, err := db.Conn.Insert("polls", []interface{}{p.ID, p.OwnerID, p.Question, p.Options, p.Votes, p.IsActive})

	if err != nil {
		return "", err
	}

	return id, nil
}

func parseQuotedArgs(s string) []string {
	var r []string
	var c []rune
	var inQ bool
	var slash bool
	for _, x := range s {
		switch x {
		case '"':
			if !slash {
				inQ = !inQ
			} else {
				c = append(c, x)
			}
			slash = false
		case '\\':
			if slash {
				c = append(c, '\\')
				slash = false
			} else {
				slash = true
			}
		default:
			if slash {
				c = append(c, '\\')
				slash = false
			}
			if inQ {
				c = append(c, x)
			} else {
				if x == ' ' || x == '\t' || x == '\n' {
					if len(c) > 0 {
						r = append(r, string(c))
						c = nil
					}
				} else {
					c = append(c, x)
				}
			}
		}
	}
	if len(c) > 0 {
		r = append(r, string(c))
	}
	return r
}

func fmtCreateError(cmd string) string {
	return fmt.Sprintf("[%s] Need question and at least one option\nExample: /votebot create \"Title?\" \"Opt1\" \"Opt2\"", cmd)
}
