package inputs

import (
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/Baipyrus/AoC-25/internal/registry"
	"github.com/ktr0731/go-fuzzyfinder"
)

type File struct {
	Path    string
	Content string
}

func GetChallenge(challenges []registry.Challenge) func(string) {
	idx, err := fuzzyfinder.Find(
		challenges,
		func(i int) string {
			return challenges[i].Name
		},
		fuzzyfinder.WithPromptString("Select Challenge: "))
	if err != nil {
		log.Fatal(err)
	}

	return challenges[idx].Exec
}

func GetInput(path string) string {
	var files []File

	err := filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		name := info.Name()
		isInput := strings.HasSuffix(name, ".txt")
		if !isInput {
			return nil
		}

		file, err := os.Open(path)

		if err != nil {
			return err
		}
		defer file.Close()

		bytes, err := io.ReadAll(file)
		if err != nil {
			return err
		}

		files = append(
			files,
			File{
				path,
				string(bytes),
			})

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	idx, err := fuzzyfinder.Find(
		files,
		func(i int) string {
			return files[i].Path
		},
		fuzzyfinder.WithPromptString("Select Input File: "),
		fuzzyfinder.WithPreviewWindow(func(i, _, _ int) string {
			if i == -1 {
				return ""
			}
			return files[i].Content
		}))
	if err != nil {
		log.Fatal(err)
	}

	return files[idx].Content
}
