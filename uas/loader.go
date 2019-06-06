package uas

import "encoding/json"

type Pat struct {
	Pattern   string   `json:"pattern"`
	URL       string   `json:"url"`
	Instances []string `json:"instances"`
}

func Load() []Pat {
	return load()
}

func load() []Pat {
	j := FSMustByte(false, "/deps/crawler-user-agents/crawler-user-agents.json")
	var r []Pat
	_ = json.Unmarshal(j, &r)
	return r
}
