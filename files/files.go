package files

import (
	"strings"

	"github.com/pkg/errors"
	"upspin.io/upspin"
)

type Accesser interface {
	Glob(pattern string) ([]*upspin.DirEntry, error)
}

type Server struct {
	Accesser Accesser
}

func (s *Server) List(path string) ([]string, error) {
	path = formatPath(path)

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

func formatPath(path string) string {
	path = path + "*"
	path = strings.TrimPrefix(path, "/")
	return path
}
