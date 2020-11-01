// +build linux

package libgeoip

import (
	"fmt"
	"go/build"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadPipeline(t *testing.T) {
	r := build.Default.ReleaseTags
	if len(r) < 2 {
		fmt.Println(r, "len <2!")
		t.SkipNow()
	}
	version := r[1]

	if version != "1.5" {
		fmt.Println("version != 1.5, skipped")
		t.SkipNow()

	}
	env, err := GetGoEnv()
	assert.Nil(t, err)
	modPath, found := env["GOMOD"]
	assert.True(t, found, "GOMOD not set")
	currentDir := filepath.Dir(modPath)
	fmt.Println(currentDir)
	path := strings.Join([]string{currentDir, "examples", "pipeline.so"}, string(os.PathSeparator))
	pipeline, err := LoadPipeline(path)
	if err != nil {
		t.Fatal(err)
	}

	_, err = pipeline.Init()

	if err != nil {
		t.Fatal(err)
	}

}
