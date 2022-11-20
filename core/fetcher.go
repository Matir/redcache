package core

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	config_pb "github.com/Matir/redcache/config/pb"
	"github.com/Matir/redcache/log"
)

const (
	USER_AGENT = "redcache/fetcher Go/1.x"
)

var logger = log.LoggerForModule("core")

// Fetches (and extracts) a tool source
func FetchToolSource(ctx context.Context, src *config_pb.ToolSource) (io.ReadCloser, error) {
	if src == nil {
		return nil, fmt.Errorf("invalid nil ToolSource")
	}
	if src.SourcePath == "" {
		logger.Error("Missing tool source for fetch!")
		return nil, fmt.Errorf("empty Source Path when fetching")
	}
	u, err := url.Parse(src.SourcePath)
	if err != nil {
		return nil, fmt.Errorf("error parsing source path: %w", err)
	}
	var rdr io.ReadCloser
	if u.Scheme == "" {
		rdr, err = fetchLocalSource(ctx, src)
	} else {
		rdr, err = fetchNetSource(ctx, src, u)
	}
	if err != nil {
		return nil, err
	}
	if src.ArchivePath != "" {
		// Extract path
		defer rdr.Close()
		return ExtractFromArchive(ctx, rdr, src.ArchivePath)
	}
	return rdr, nil
}

func fetchLocalSource(ctx context.Context, src *config_pb.ToolSource) (io.ReadCloser, error) {
	fp, err := os.Open(src.SourcePath)
	if err != nil {
		logger.WithFields(log.Fields{
			"path": src.SourcePath,
			"err":  err,
		}).Error("Failed opening local source.")
		return nil, fmt.Errorf("failed opening local source: %w", err)
	}
	return fp, nil
}

func fetchNetSource(ctx context.Context, src *config_pb.ToolSource, u *url.URL) (io.ReadCloser, error) {
	logger.WithFields(log.Fields{
		"url": u,
	}).Info("Fetching from Net")
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", USER_AGENT)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.WithFields(log.Fields{
			"url": u,
			"err": err,
		}).Error("Error fetching from Net")
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		logger.WithFields(log.Fields{
			"url":    u,
			"status": resp.StatusCode,
		}).Error("HTTP Response was not OK")
		return nil, fmt.Errorf("non-200 response: %s", resp.Status)
	}
	return resp.Body, nil
}
