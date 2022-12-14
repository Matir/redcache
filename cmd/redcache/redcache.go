package main

import (
	"bytes"
	"context"
	_ "embed"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Matir/redcache/config"
	"github.com/Matir/redcache/fetcher"
	"github.com/Matir/redcache/log"
	"github.com/Matir/redcache/server"
)

var (
	logLevelFlag   = flag.String("log-level", "info", "Log level for logging.")
	logFileFlag    = flag.String("log-file", "", "Log file for logging (default stderr)")
	configFileFlag = flag.String("config", "", "Path to config file.")
	listenAddrFlag = flag.String("listen-addr", "", "Listen address (overrides config)")
	cacheDirFlag   = flag.String("cache-dir", "", "Cache directory (overrides config)")
	preCacheFlag   = flag.Bool("precache", false, "Download tools to cache directory and exit.")

	// Should this be somewhere else?
	//go:embed config.asciipb
	embeddedConfig []byte
)

func main() {
	flag.Parse()

	loglevel, err := log.ParseLevel(*logLevelFlag)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unknown log level %s\n", *logLevelFlag)
		os.Exit(1)
	}
	log.SetLevel(loglevel)
	if *logFileFlag != "" {
		logfp, err := os.Create(*logFileFlag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening log file (%s): %s\n", *logFileFlag, err)
			os.Exit(1)
		}
		defer logfp.Close()
		log.SetOutput(logfp)
	}
	logger := log.LoggerForPackage("main")

	var mainConfig *config.Config
	if *configFileFlag != "" {
		mainConfig, err = LoadConfigFromFile(*configFileFlag)
		if err == nil {
			logger.WithField("file", *configFileFlag).Info("Using config file.")
		}
	} else {
		mainConfig, err = LoadEmbeddedConfig()
		if err == nil {
			logger.Info("Using embedded config")
		}
	}
	if err != nil {
		logger.WithField("err", err).Error("Error loading config.")
		fmt.Fprintf(os.Stderr, "Error loading config: %s\n", err)
		os.Exit(1)
	}

	// Build the fetcher
	if *cacheDirFlag != "" {
		mainConfig.CacheDir = *cacheDirFlag
	}
	cacheDir := expandUser(mainConfig.CacheDir)
	fetchImpl := fetcher.NewFetchCache(cacheDir)

	if *preCacheFlag {
		runPrecache(logger, mainConfig, fetchImpl)
		return
	}

	// Override listen addr
	if *listenAddrFlag != "" {
		mainConfig.ListenAddr = *listenAddrFlag
	}

	srv, err := server.NewCacheServer(mainConfig, fetchImpl)
	if err != nil {
		logger.WithField("err", err).Error("Error building server.")
		fmt.Fprintf(os.Stderr, "Error building server: %s\n", err)
		os.Exit(1)
	}
	if err := srv.ListenAndServe(); err != nil {
		logger.WithField("err", err).Error("Server shutting down.")
	}
}

func LoadConfigFromFile(filename string) (*config.Config, error) {
	fp, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fp.Close()
	return config.LoadConfigFromReader(fp)
}

func LoadEmbeddedConfig() (*config.Config, error) {
	buf := bytes.NewReader(embeddedConfig)
	return config.LoadConfigFromReader(buf)
}

// If the path starts with ~/, we expand the current user's homedir.
// This does not support ~name style.
func expandUser(path string) string {
	pieces := strings.Split(path, string(filepath.Separator))
	if pieces[0] != "~" {
		return path
	}
	homedir, err := os.UserHomeDir()
	if err != nil {
		return path
	}
	pieces[0] = homedir
	return filepath.Join(pieces...)
}

func runPrecache(logger *log.Entry, cfg *config.Config, fetchImpl *fetcher.FetchCache) {
	ctx := context.Background()
	ct := 0
	for _, tool := range cfg.Tool {
		logger.WithField("tool", tool.Name).Info("Caching tool.")
		rdr, err := fetchImpl.FetchTool(ctx, config.Tool{Tool: tool})
		if err == nil {
			ct++
			rdr.Close()
		}
	}
	logger.WithField("count", ct).Info("Cached tools.")
}
