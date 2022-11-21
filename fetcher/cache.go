package fetcher

import (
	"bytes"
	"context"
	"embed"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"sync"

	"github.com/Matir/redcache/config"
	"github.com/Matir/redcache/log"
)

const (
	embedCacheDirName = "cachedir"
)

var (
	//go:embed cachedir/*
	cacheEmbedFS  embed.FS
	listCacheOnce sync.Once
)

type FetchCache struct {
	cacheDir   string
	cacheFs    fs.FS
	writeCache bool
}

func NewFetchCache(cacheDir string) *FetchCache {
	logger.WithField("cachedir", cacheDir).Info("Using local cache dir.")
	return &FetchCache{
		cacheDir:   cacheDir,
		cacheFs:    os.DirFS(cacheDir),
		writeCache: true,
	}
}

func (c *FetchCache) FetchTool(ctx context.Context, tool config.Tool) (io.ReadCloser, error) {
	listCache()
	cacheName := tool.GetCacheName()
	// check embeds
	if data, err := cacheEmbedFS.ReadFile(path.Join(embedCacheDirName, cacheName)); err == nil {
		// wrap in a buffer and return
		logger.WithFields(log.Fields{
			"tool":      tool.String(),
			"cacheName": cacheName,
		}).Info("Serving from embed.")
		return io.NopCloser(bytes.NewBuffer(data)), nil
	}
	// check cachedir
	if fp, err := c.cacheFs.Open(cacheName); err == nil {
		logger.WithFields(log.Fields{
			"tool":      tool.String(),
			"cacheName": cacheName,
		}).Info("Serving from FS cache.")
		return fp, nil
	} else if !errors.Is(err, fs.ErrNotExist) {
		logger.WithFields(log.Fields{
			"cacheName": cacheName,
			"err":       err,
		}).Error("Error opening from cache.")
	}
	// source files
	rdr, err := c.fetchToolNoCache(ctx, tool)
	if err != nil {
		logger.WithFields(log.Fields{
			"tool": tool.Name,
		}).Error("No valid source found.")
		return nil, err
	}
	if !c.writeCache {
		return rdr, nil
	}
	// Write rdr to cache with copy
	if err := os.MkdirAll(c.cacheDir, fs.ModeDir|0755); err != nil {
		logger.WithField("err", err).Error("Failed making cache dir.")
		return nil, err
	}
	fp, err := os.Create(filepath.Join(c.cacheDir, cacheName))
	if err != nil {
		logger.WithFields(log.Fields{
			"err": err,
		}).Error("Failed writing cache file.")
		return rdr, nil
	}
	defer rdr.Close()
	defer fp.Close()
	var buf bytes.Buffer
	mw := io.MultiWriter(&buf, fp)
	if _, err := io.Copy(mw, rdr); err != nil {
		logger.WithFields(log.Fields{
			"err": err,
		}).Error("Error copying readers")
		return nil, fmt.Errorf("error copying reader: %w", err)
	}
	return io.NopCloser(&buf), nil
}

func (c *FetchCache) fetchToolNoCache(ctx context.Context, tool config.Tool) (io.ReadCloser, error) {
	for _, src := range tool.Source {
		if rdr, err := FetchToolSource(ctx, src); err != nil {
			logger.WithFields(log.Fields{
				"tool": tool.Name,
				"err":  err,
				"src":  src.SourcePath,
			}).Warn("Could not fetch source.")
		} else {
			return rdr, nil
		}
	}
	return nil, fmt.Errorf("no source to download")
}

func listCache() {
	listCacheOnce.Do(func() {
		cacheInitLog := log.LoggerForPackage("cache")
		dirEnts, err := cacheEmbedFS.ReadDir(embedCacheDirName)
		if err != nil {
			cacheInitLog.WithField("err", err).Error("Error getting directory list from cache embed dir.")
			return
		}
		for _, de := range dirEnts {
			cacheInitLog.WithFields(log.Fields{
				"name": de.Name(),
				"mode": de.Type().String(),
			}).Info("Cache Embed Entry")
		}
	})
}
