// Define and load configurations for the redcache system.
package config

//go:generate mkdir -p pb
//go:generate protoc --go_out=pb --go_opt=paths=source_relative config.proto

import (
	"fmt"
	"io"
	"strings"

	"google.golang.org/protobuf/proto"

	"github.com/Matir/redcache/config/pb"
)

type Config struct {
	pb.Config
}

type Tool struct {
	*pb.Tool
}

func LoadConfigFromReader(r io.Reader) (*Config, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := proto.Unmarshal(data, &cfg.Config); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func (cfg *Config) GetToolMap() (map[string]Tool, error) {
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
