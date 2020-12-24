package history

import (
	"fmt"

	"github.com/nikoksr/nsh/internal/command"
	orderedmap "github.com/wk8/go-ordered-map"
)

type CommandHistory struct {
	store *orderedmap.OrderedMap
}

func NewCommandHistory() *CommandHistory {
	return &CommandHistory{store: orderedmap.New()}
}

func (ch CommandHistory) Append(cmd command.Command) {
	ch.store.Set(cmd.Name, struct{}{})
}

func (ch CommandHistory) GetLastCommand() string {
	if ch.store.Len() < 1 {
		return ""
	}
	entry := ch.store.Newest()

	return fmt.Sprintf("%v", entry.Key)
}
