package collection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAblums(t *testing.T) {
	c := &Collection{
		Images: []string{
			"user.test@gmail.com/root/pic1.jpg",
			"user.test@gmail.com/root/pic2.jpg",
			"user.test@gmail.com/root/a subdir/another pic.jpg",
			"user.test@gmail.com/root/a subdir/something.jpg",
			"user.test@gmail.com/root/another subdir/something else.jpg",
			"user.test@gmail.com/root/a subdir/pic1.jpg",
			"user.test@gmail.com/root/pic3.jpg",
			"user.test@gmail.com/root/a sub/subdir/something.jpg",
		}}

	expected := []string{
		"user.test@gmail.com/root",
		"user.test@gmail.com/root/a subdir",
		"user.test@gmail.com/root/another subdir",
		"user.test@gmail.com/root/a sub/subdir",
	}

	actual := c.Albums()

	assert.Equal(t, len(expected), len(actual))
	for _, e := range expected {
		assert.Contains(t, actual, e)
	}
}
