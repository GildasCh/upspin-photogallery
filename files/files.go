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

func (s *Server) List(path string) ([]string, error) {
	path = formatDirPath(path)

	entries, err := s.Accesser.Glob(path)
	if err != nil {
		return nil, errors.Wrapf(err, "could not Glob path %q", path)
	}

	filenames := []string{}
	for _, entry := range entries {
		filenames = append(filenames, string(entry.Name))
	}

	return filenames, nil
}

func (s *Server) Get(path string) (io.Reader, error) {
	upath := formatFilePath(path)

	f, err := s.Accesser.Open(upath)
	if err != nil {
		return nil, errors.Wrapf(err, "could not Open path %q", path)
	}

	return f, nil
}

func formatDirPath(path string) string {
	path = path + "*"
	path = strings.TrimPrefix(path, "/")
	return path
}

func formatFilePath(path string) upspin.PathName {
	path = strings.TrimPrefix(path, "/")
	return upspin.PathName(path)
}
