// Define and load configurations for the redcache system.
package config

//go:generate mkdir -p pb
//go:generate protoc --go_out=pb --go_opt=paths=source_relative config.proto

import (
	"fmt"
	"io"
	"sort"
	"strings"

	"google.golang.org/protobuf/encoding/prototext"

	"github.com/Matir/redcache/config/pb"
)

type Config struct {
	pb.Config
}

type Tool struct {
	*pb.Tool
}

var (
	platformNames = map[pb.Tool_Platform]string{
		pb.Tool_PLATFORM_UNKNOWN: "(Unknown)",
		pb.Tool_PLATFORM_ANY:     "Any",
		pb.Tool_PLATFORM_LINUX:   "Linux",
		pb.Tool_PLATFORM_WINDOWS: "Windows",
		pb.Tool_PLATFORM_OSX:     "OS X",
	}
	archNames = map[pb.Tool_Architecture]string{
		pb.Tool_ARCH_UNKNOWN: "(Unknown)",
		pb.Tool_ARCH_ANY:     "Any",
		pb.Tool_ARCH_X86:     "x86",
		pb.Tool_ARCH_X64:     "x86_64",
		pb.Tool_ARCH_ARM:     "ARM",
	}
)

type ToolMap map[string]Tool

func LoadConfigFromReader(r io.Reader) (*Config, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := prototext.Unmarshal(data, &cfg.Config); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func (cfg *Config) GetToolMap() (ToolMap, error) {
	res := make(map[string]Tool)
	for _, t := range cfg.Tool {
		wrapped := Tool{t}
		for _, p := range t.Path {
			p = strings.TrimLeft(p, "/")
			if _, ok := res[p]; ok {
				return nil, fmt.Errorf("duplicate path: %v", p)
			}
			res[p] = wrapped
		}
	}
	return res, nil
}

func (t Tool) GetCacheName() string {
	if len(t.Path) < 1 {
		// This is unservable anyway
		return ""
	}
	return strings.ReplaceAll(strings.TrimLeft(strings.TrimRight(t.Path[0], "/"), "/"), "/", "_")
}

func (t Tool) String() string {
	return t.Name
}

func (t Tool) GetPlatformName() string {
	if v, ok := platformNames[t.Platform]; ok {
		return v
	}
	return t.Platform.String()
}

func (t Tool) GetArchName() string {
	if v, ok := archNames[t.Arch]; ok {
		return v
	}
	return t.Arch.String()
}

func (t ToolMap) Iterate() toolMapIterator {
	tmi := make([]string, len(t))
	i := 0
	for k := range t {
		tmi[i] = k
		i++
	}
	rv := toolMapIterator{
		keys: tmi,
		tm:   t,
	}
	sort.Sort(rv)
	return rv
}

type toolMapIterator struct {
	keys    []string
	tm      ToolMap
	started bool
}

func (tmi toolMapIterator) Len() int {
	return len(tmi.keys)
}

func (tmi toolMapIterator) Less(i, j int) bool {
	return tmi.keys[i] < tmi.keys[j]
}

func (tmi toolMapIterator) Swap(i, j int) {
	tmi.keys[i], tmi.keys[j] = tmi.keys[j], tmi.keys[i]
}

func (tmi *toolMapIterator) Next() bool {
	if len(tmi.keys) == 0 {
		return false
	}
	if !tmi.started {
		tmi.started = true
		return true
	}
	tmi.keys = tmi.keys[1:]
	return len(tmi.keys) > 0
}

func (tmi toolMapIterator) Item() (string, Tool) {
	k := tmi.keys[0]
	return k, tmi.tm[k]
}
