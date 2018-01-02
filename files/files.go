package files

import (
	"io"
	"strings"

	"github.com/pkg/errors"
	"upspin.io/upspin"
)

type Accesser interface {
	Glob(pattern string) ([]*upspin.DirEntry, error)
	Open(name upspin.PathName) (upspin.File, error)
}

type Server struct {
	Accesser Accesser
}

func createPattern(path string) string {
	pattern := strings.TrimPrefix(path, "/")

	if strings.Contains(pattern, "*") {
		return pattern
	}
	return strings.TrimSuffix(pattern, "/") + "/*"
}

func (s *Server) List(path string) ([]string, error) {
	pattern := createPattern(path)

	entries, err := s.Accesser.Glob(pattern)
	if err != nil {
		return nil, errors.Wrapf(err, "could not Glob pattern %q", pattern)
	}

	filenames := []string{}
	for _, entry := range entries {
		filenames = append(filenames, string(entry.Name))
	}

	return filenames, nil
}

func isImage(filename string) bool {
	filename = strings.ToLower(filename)
	return strings.HasSuffix(filename, ".jpg") ||
		strings.HasSuffix(filename, ".jpeg") ||
		strings.HasSuffix(filename, ".png") ||
		strings.HasSuffix(filename, ".gif") ||
		strings.HasSuffix(filename, ".bmp") ||
		strings.HasSuffix(filename, ".webp")
}

func (s *Server) ListImages(path string) ([]string, error) {
	filenames, err := s.List(path)

	filtered := []string{}
	for _, f := range filenames {
		if isImage(f) {
			filtered = append(filtered, f)
		}
	}

	return filtered, err
}

func formatFilePath(path string) upspin.PathName {
	path = strings.TrimPrefix(path, "/")
	return upspin.PathName(path)
}

func (s *Server) Get(path string) (io.Reader, error) {
	upath := formatFilePath(path)

	f, err := s.Accesser.Open(upath)
	if err != nil {
		return nil, errors.Wrapf(err, "could not Open path %q", path)
	}

	return f, nil
}
