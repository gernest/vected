package cover

import (
	"encoding/json"
	"fmt"
	"go/token"
	"sync"

	"github.com/gopherjs/gopherjs/js"
)

const (
	Key = "mad_coverage_stats"
)

type Profile struct {
	FileName string
	Mode     string
	Blocks   []*ProfileBlock
}

// ProfileBlock represents a single block of profiling data.
type ProfileBlock struct {
	StartPosition  *token.Position
	EndPosition    *token.Position
	NumStmt, Count int
}

var state = &State{
	files: make(map[string]*CoverStats),
}

type State struct {
	files map[string]*CoverStats
	mu    sync.RWMutex
}

type CoverStats struct {
	Profile *Profile
	mu      sync.RWMutex
}

func (c *CoverStats) Mark(p *ProfileBlock) int {
	c.mu.Lock()
	c.Profile.Blocks = append(c.Profile.Blocks, p)
	idx := len(c.Profile.Blocks) - 1
	c.mu.Unlock()
	return idx
}

func (c *CoverStats) Hit(idx int, pos *token.Position) {
	c.mu.Lock()
	b := c.Profile.Blocks[idx]
	b.Count++
	c.mu.Unlock()
}

func Mark(numStmt int, start, end *token.Position) int {
	state.mu.Lock()
	defer state.mu.Unlock()
	p := &ProfileBlock{
		StartPosition: start,
		EndPosition:   end,
		NumStmt:       numStmt,
	}
	fileName := p.StartPosition.Filename
	f, ok := state.files[fileName]
	if !ok {
		f = &CoverStats{Profile: &Profile{
			FileName: fileName,
		}}
		state.files[fileName] = f
		return f.Mark(p)

	}
	return f.Mark(p)
}

func Hit(idx int, pos *token.Position) {
	state.mu.Lock()
	defer state.mu.Unlock()
	if f, ok := state.files[pos.Filename]; ok {
		f.Hit(idx, pos)
	}
}

func Stats() []*CoverStats {
	state.mu.RLock()
	var o []*CoverStats
	for _, v := range state.files {
		o = append(o, v)
	}
	state.mu.RUnlock()
	return o
}

// JSON marshals current coverage state to json string.
func JSON() string {
	b, _ := json.MarshalIndent(Stats(), "", "  ")
	return string(b)
}

// Dump calls console.info("coverage",$COVERAGE_DATA_IN_JSON)
func Dump() (err error) {
	defer func() {
		if perr := recover(); perr != nil {
			err = fmt.Errorf("%v", perr)
		}
	}()
	js.Global.Get("console").Call("info", "coverage", JSON())
	return
}
