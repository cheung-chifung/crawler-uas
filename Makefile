PAT_JSON=./deps/crawler-user-agents/crawler-user-agents.json
$(PAT_JSON):
	(cd deps/crawler-user-agents && git submodule update)
	[[ -e "./deps/crawler-user-agents/crawler-user-agents.jso" ]] || echo "file not found"

uas/generated.go: $(PAT_JSON)
	esc -o uas/generated.go -pkg uas $(PAT_JSON)

all: uas/generated.go
