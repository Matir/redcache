package server

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Matir/redcache/config"
	"github.com/Matir/redcache/fetcher"
	"github.com/Matir/redcache/log"
)

const (
	SERVER_NAME = "RedCache/1.x"
)

var logger = log.LoggerForPackage("server")

type CloseLenReader interface {
	io.Reader
	io.Closer
	Len() int
}

type CacheServer struct {
	listenAddr string
	toolMap    map[string]config.Tool
	serveIndex bool
	rootPrefix string
	fetchCache *fetcher.FetchCache
}

func NewCacheServer(cfg *config.Config) (*CacheServer, error) {
	toolMap, err := cfg.GetToolMap()
	if err != nil {
		return nil, err
	}
	rv := &CacheServer{
		listenAddr: cfg.ListenAddr,
		toolMap:    toolMap,
		serveIndex: !cfg.HideIndex,
		rootPrefix: cfg.RootPath,
	}
	if !strings.HasSuffix(rv.rootPrefix, "/") {
		rv.rootPrefix += "/"
	}
	return rv, nil
}

func (srv *CacheServer) ListenAndServe() error {
	logger.WithField("addr", srv.listenAddr).Info("Starting HTTP Server")
	return http.ListenAndServe(srv.listenAddr, srv)
}

func (srv *CacheServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Add("Server", SERVER_NAME)
	reqPath := r.URL.Path
	logger.WithFields(log.Fields{
		"path":        reqPath,
		"remote_addr": r.RemoteAddr,
	}).Info("Request")
	if !strings.HasPrefix(reqPath, srv.rootPrefix) {
		// error
		logger.WithFields(log.Fields{
			"path":        reqPath,
			"remote_addr": r.RemoteAddr,
			"prefix":      srv.rootPrefix,
		}).Error("Missing prefix.")
		srv.Serve404(ctx, w, r)
		return
	}
	reqPath = strings.TrimLeft(reqPath, srv.rootPrefix)
	if tool, ok := srv.toolMap[reqPath]; ok {
		srv.ServeTool(ctx, w, r, tool)
		return
	} else {
		srv.Serve404(ctx, w, r)
		return
	}
}

func (srv *CacheServer) Serve404(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	// TODO: actual 404 page
}

func (srv *CacheServer) ServeTool(ctx context.Context, w http.ResponseWriter, r *http.Request, tool config.Tool) {
	if srv.fetchCache == nil {
		logger.Error("Fetch Cache is nil!")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	rdr, err := srv.fetchCache.FetchTool(ctx, tool)
	if err != nil {
		logger.WithFields(log.Fields{
			"err": err,
		}).Error("Unable to serve tool.")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer rdr.Close()
	if lenrdr, ok := rdr.(CloseLenReader); ok {
		w.Header().Add("Content-Length", fmt.Sprintf("%d", lenrdr.Len()))
	}
	if n, err := io.Copy(w, rdr); err != nil {
		logger.WithFields(log.Fields{
			"len_copied": n,
			"err":        err,
		}).Error("Failed copying to client.")
	}
}
