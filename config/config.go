// Define and load configurations for the redcache system.
package config

//go:generate mkdir -p pb
//go:generate protoc --go_out=pb --go_opt=paths=source_relative config.proto

import (
	"io"

	"google.golang.org/protobuf/proto"

	"github.com/Matir/redcache/config/pb"
)

type Config struct {
	pb.Config
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
