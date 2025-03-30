package handlers

import (
    "fmt"

    "vote_bot/db"
    "vote_bot/poll"
)

func replacePoll(p *poll.Poll) (interface{}, error) {
    return db.Conn.Replace("polls", []interface{}{
        p.ID,
        p.OwnerID,
        p.Question,
        p.Options,
        p.Votes,
        p.IsActive,
    })
}

func replaceDelete(id string) (interface{}, error) {
    return db.Conn.Delete("polls", "primary", []interface{}{id})
}

func getPollByID(id string) (*poll.Poll, error) {
    r, e := db.Conn.Select("polls", "primary", 0, 1, 0, []interface{}{id})
    if e != nil {
        return nil, e
    }
    if len(r.Tuples()) == 0 {
        return nil, fmt.Errorf("poll %s not found", id)
    }
    t := r.Tuples()[0]
    p := &poll.Poll{
        ID:       t[0].(string),
        OwnerID:  t[1].(string),
        Question: t[2].(string),
        Options:  toStringSlice(t[3]),
        Votes:    toMapStringSlice(t[4]),
        IsActive: t[5].(bool),
    }
    return p, nil
}

func toStringSlice(i interface{}) []string {
    arr, ok := i.([]interface{})
    if !ok {
        return nil
    }
    var out []string
    for _, v := range arr {
        s, _ := v.(string)
        out = append(out, s)
    }
    return out
}

func toMapStringSlice(i interface{}) map[string][]string {
    m, ok := i.(map[interface{}]interface{})
    if !ok {
        return nil
    }
    out := make(map[string][]string)
    for k, v := range m {
        ks, _ := k.(string)
        vs, ok := v.([]interface{})
        if !ok {
            continue
        }
        var tmp []string
        for _, x := range vs {
            xs, _ := x.(string)
            tmp = append(tmp, xs)
        }
        out[ks] = tmp
    }
    return out
}
