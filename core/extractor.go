package core

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"compress/bzip2"
	"compress/gzip"
	"context"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/Matir/redcache/log"
)

type Extractor interface {
	Name() string
	Identifier(*bytes.Buffer) bool
	Extract(*bytes.Buffer, string) (io.ReadCloser, error)
}

type TarExtractor string
type ZipExtractor string
type GzipExtractor string
type BzipExtractor string

var extractors = []Extractor{
	TarExtractor("tar"),
	ZipExtractor("zip"),
	GzipExtractor("gzip"),
	BzipExtractor("bzip2"),
}

// Extract a single file from an Archive
// Do not depend on the source ReadCloser outliving the file
func ExtractFromArchive(ctx context.Context, src io.ReadCloser, path string) (io.ReadCloser, error) {
	// Read entire thing into a buffer so we can seek.
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, src); err != nil {
		logger.WithFields(log.Fields{
			"err": err,
		}).Error("Error reading archive")
		return nil, err
	}

	return extractFromBuffer(&buf, path)
}

func extractFromBuffer(buf *bytes.Buffer, path string) (io.ReadCloser, error) {
	for _, e := range extractors {
		if e.Identifier(bytes.NewBuffer(buf.Bytes())) {
			logger.WithFields(log.Fields{
				"format": e.Name(),
			}).Info("Identified archive file.")
			return e.Extract(buf, path)
		}
	}

	return nil, errors.New("no valid extractor found")
}

func pathStringsMatch(a, b string) bool {
	return strings.TrimLeft(a, "/") == strings.TrimLeft(b, "/")
}

// Plain tar implementation
func (t TarExtractor) Name() string {
	return string(t)
}

func (TarExtractor) Identifier(buf *bytes.Buffer) bool {
	magic := []byte("ustar\x00")
	magicstart := 257
	return bytes.Equal(buf.Bytes()[magicstart:magicstart+len(magic)], magic)
}

func (TarExtractor) Extract(buf *bytes.Buffer, path string) (io.ReadCloser, error) {
	rdr := tar.NewReader(buf)
	for {
		hdr, err := rdr.Next()
		if err == io.EOF {
			// done here, guess it's not found!
			logger.WithFields(log.Fields{
				"path": path,
			}).Error("Entry not found in archive")
			return nil, fmt.Errorf("entry %v not found in archive", path)
		}
		if err != nil {
			logger.WithFields(log.Fields{
				"err":  err,
				"path": path,
			}).Error("Error reading archive.")
			return nil, fmt.Errorf("error reading archive: %w", err)
		}
		if pathStringsMatch(hdr.Name, path) {
			// We have a match!
			var destBuf bytes.Buffer
			if n, err := io.Copy(&destBuf, rdr); err != nil {
				return nil, err
			} else {
				if n == 0 {
					return nil, fmt.Errorf("no bytes read")
				}
				return io.NopCloser(&destBuf), nil
			}
		}
	}
}

// Zip File Implementation
func (z ZipExtractor) Name() string {
	return string(z)
}

func (ZipExtractor) Identifier(buf *bytes.Buffer) bool {
	magic := []byte("PK\x03\x04")
	return bytes.Equal(buf.Bytes()[:len(magic)], magic)
}

func (ZipExtractor) Extract(buf *bytes.Buffer, path string) (io.ReadCloser, error) {
	bufrdr := bytes.NewReader(buf.Bytes())
	ziprdr, err := zip.NewReader(bufrdr, int64(bufrdr.Len()))
	if err != nil {
		logger.WithFields(log.Fields{
			"err": err,
		}).Error("Error reading zipfile.")
		return nil, err
	}
	fp, err := ziprdr.Open(strings.TrimLeft(path, "/"))
	if err != nil {
		logger.WithFields(log.Fields{
			"err":  err,
			"path": path,
		}).Error("Error reading zipfile.")
		return nil, err
	}
	var destBuf bytes.Buffer
	if _, err := io.Copy(&destBuf, fp); err != nil {
		logger.WithFields(log.Fields{
			"err":  err,
			"path": path,
		}).Error("Error reading zipfile.")
		return nil, err
	}
	return io.NopCloser(&destBuf), nil
}

// Gzip compression layer
func (g GzipExtractor) Name() string {
	return string(g)
}

func (GzipExtractor) Identifier(buf *bytes.Buffer) bool {
	magic := []byte("\x1f\x8b\x08")
	return bytes.Equal(buf.Bytes()[:len(magic)], magic)
}

func (GzipExtractor) Extract(buf *bytes.Buffer, path string) (io.ReadCloser, error) {
	gzrdr, err := gzip.NewReader(buf)
	if err != nil {
		logger.WithField("err", err).Error("Error opening gzip.")
		return nil, fmt.Errorf("error reading gzip: %w", err)
	}
	var tmpbuf bytes.Buffer
	if _, err := io.Copy(&tmpbuf, gzrdr); err != nil {
		logger.WithField("err", err).Error("Error reading gzip.")
		return nil, fmt.Errorf("error reading gzip: %w", err)
	}
	return extractFromBuffer(&tmpbuf, path)
}

// Bzip compression layer
func (b BzipExtractor) Name() string {
	return string(b)
}

func (BzipExtractor) Identifier(buf *bytes.Buffer) bool {
	magic := []byte("BZh")
	return bytes.Equal(buf.Bytes()[:len(magic)], magic)
}

func (BzipExtractor) Extract(buf *bytes.Buffer, path string) (io.ReadCloser, error) {
	rdr := bzip2.NewReader(buf)
	var tmpbuf bytes.Buffer
	if _, err := io.Copy(&tmpbuf, rdr); err != nil {
		logger.WithField("err", err).Error("Error opening bzip2.")
		return nil, fmt.Errorf("error reading bzip2: %w", err)
	}
	return extractFromBuffer(&tmpbuf, path)
}
