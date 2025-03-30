package poll

type Poll struct {
    ID       string
    OwnerID  string
    Question string
    Options  []string
    Votes    map[string][]string
    IsActive bool
}
