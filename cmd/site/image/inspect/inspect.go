package inspect

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/dsoprea/go-exif/v3"
	"github.com/spf13/cobra"
	"github.com/wwmoraes/site/pkg/exifer"
	"github.com/wwmoraes/site/pkg/functional"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "inspect",
		Short: "shows image EXIF metadata",
		Args:  cobra.ArbitraryArgs,
		RunE:  inspect,
	}

	return cmd
}

func inspect(cmd *cobra.Command, args []string) error {
	pwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current working directory: %w", err)
	}

	fsys := os.DirFS(pwd)

	logger := func(err error) {
		cmd.PrintErrln(err)
	}

	bufferSize := runtime.NumCPU()

	imagePaths := functional.LogErrors(getImagePaths(fsys, args, bufferSize), bufferSize, logger)
	imageTags := functional.LogErrors(getImageTags(imagePaths, bufferSize), bufferSize, logger)

	for tags := range imageTags {
		for key, value := range tags {
			cmd.Printf("%20s: %s\n", key, value)
		}

		//nolint:mnd // 2 * 40 = 80 columns
		cmd.Println(strings.Repeat("- ", 40))
	}

	return nil
}

func getImagePaths(fsys fs.FS, patterns []string, buffer int) <-chan *functional.Result[string] {
	imagePaths := make(chan *functional.Result[string], buffer)

	go func() {
		defer func() {
			close(imagePaths)
		}()

		for _, pattern := range patterns {
			matches, err := doublestar.Glob(fsys, pattern, doublestar.WithFilesOnly())
			if err != nil {
				functional.SendError(imagePaths, fmt.Errorf("failed to parse pattern: %w", err))

				return
			}

			for _, match := range matches {
				functional.SendValue(imagePaths, strings.TrimSuffix(match, ".json"))
			}
		}
	}()

	return imagePaths
}

//nolint:funlen // TODO refactor getImageTags
func getImageTags(imagePaths <-chan string, buffer int) <-chan *functional.Result[map[string]any] {
	imageTags := make(chan *functional.Result[map[string]any], buffer)

	go func() {
		parser := exifer.NewMetaMediaParser()

		for imagePath := range imagePaths {
			fd, err := os.OpenFile(imagePath, os.O_RDWR, 0o600)
			if err != nil {
				functional.SendError(imageTags, err)

				continue
			}
			defer fd.Close()

			info, err := fd.Stat()
			if err != nil {
				functional.SendError(imageTags, fmt.Errorf("failed to stat image file: %w", err))

				continue
			}

			media, err := parser.Parse(fd, int(info.Size()))
			if err != nil {
				functional.SendError(imageTags, fmt.Errorf("failed to parse image: %w", err))
			}

			ifd, _, err := media.Exif()
			if errors.Is(err, exif.ErrNoExif) {
				continue
			}

			if err != nil {
				functional.SendError(imageTags, fmt.Errorf("failed to get EXIF from image: %w", err))

				continue
			}

			tags := make(map[string]any, 2)

			tags["Directory"] = filepath.Dir(imagePath)
			tags["File Name"] = filepath.Base(imagePath)

			err = ifd.EnumerateTagsRecursively(func(i *exif.Ifd, ite *exif.IfdTagEntry) error {
				value, err := ite.Value()
				if err != nil {
					return err
				}

				tags[ite.TagName()] = fmt.Sprint(value)

				return nil
			})
			if err != nil {
				functional.SendError(imageTags, fmt.Errorf("failed to enumerate tags: %w", err))

				continue
			}

			functional.SendValue(imageTags, tags)
		}

		close(imageTags)
	}()

	return imageTags
}
