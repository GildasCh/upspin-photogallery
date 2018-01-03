package collection

import (
	"fmt"
	"strings"
)

type Collection struct {
	Images []string
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

func New(filenames []string) *Collection {
	images := []string{}
	for _, f := range filenames {
		if isImage(f) {
			images = append(images, f)
		}
	}

	fmt.Println(images)

	return &Collection{Images: images}
}

func (c *Collection) Albums() []string {
	tmp := map[string]struct{}{}

	for _, i := range c.Images {
		splitted := strings.Split(i, "/")
		tmp[strings.Join(splitted[:len(splitted)-1], "/")] = struct{}{}
	}

	uniques := []string{}

	for u, _ := range tmp {
		uniques = append(uniques, u)
	}

	return uniques
}
