box.cfg{
    listen = 3301
}

box.schema.create_space('polls', { if_not_exists = true })
box.space.polls:create_index('primary', {
    type = 'hash',
    parts = {1, 'string'},
    if_not_exists = true
})
