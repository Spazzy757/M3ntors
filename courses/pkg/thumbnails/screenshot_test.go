package thumbnails

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestScreenShotGenerate(t *testing.T) {
	assert := require.New(t)
	t.Run("generate screen shot", func(t *testing.T) {
		path, err := os.Getwd()
		assert.NoError(err)
		s := NewScreenshot(context.Background(), "test")
		s.LoadTime = 0
		s.Quality = 1
		s.OutputDir = "."
		err = s.Generate(fmt.Sprintf("file://%v/../../testdata/test.html", path))
		assert.NoError(err)
		_, err = os.Stat("./test.jpeg")
		assert.NoError(err)
		err = os.Remove("./test.jpeg")
		assert.NoError(err)
	})
}
