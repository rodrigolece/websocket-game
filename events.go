package main

import (
    "fmt"
)

type wsEvent struct{
    //The json package only accesses the exported fields of struct types
    Event   string `json:"event"`
    Id      string `json:"id"`
    Pos     vector `json:"pos"`
    Vel     vector `json:"vel"`
}

func identityEvent(id string) []byte {
    s := fmt.Sprintf(`{"event":"identity", "id":"%s"}`, id)
    return []byte(s)
}

func createPlayerEvent(p *player) []byte {
    s := fmt.Sprintf(`{"event":"createPlayer", "id":"%s", "pos":%v, "vel":%v}`,
        p.id, p.pos, p.vel)
    // fmt.Println(s)
    return []byte(s)
}

func destroyPlayerEvent(p *player) []byte {
    s := fmt.Sprintf(`{"event":"destroyPlayer", "id":"%s"}`, p.id)
    return []byte(s)
}

func updateEvent(p * player) []byte {
    s := fmt.Sprintf(`{"event":"update", "id":"%s", "pos":%v, "vel":%v}`,
        p.id, p.pos, p.vel)
    return []byte(s)
}
