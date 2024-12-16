package gojawalker

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/dop251/goja"
)

func TestWalk(t *testing.T) {
	filepath.WalkDir("./testdata", func(filepath string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		fmt.Printf("Filepath: %s\n", filepath)
		content, err := ioutil.ReadFile(filepath)
		if err != nil {
			t.Fatalf("fail to read file: %s, error: %s", filepath, err)
		}

		ast_tree, err := goja.Parse("test.js", string(content))
		if err != nil {
			fmt.Println(ast_tree, err)
			t.Fail()
		}

		visitor := ExampleVisitor{}
		Walk(visitor, ast_tree)

		return nil
	})

}
