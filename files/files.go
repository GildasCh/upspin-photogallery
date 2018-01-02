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
		if !entry.IsDir() {
			f, err := s.Accesser.Open(entry.Name)
			if err != nil {
				// we cannot access this file
				continue
			}
			f.Close()

			filenames = append(filenames, string(entry.Name))
		}

		sub, err := s.List(string(entry.Name))
		if err != nil {
			continue
		}
		filenames = append(filenames, sub...)
	}

	return filenames, nil
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
