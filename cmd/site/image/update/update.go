package update

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/wwmoraes/site/pkg/exifer"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "modifies images in-place",
		Args:  cobra.ArbitraryArgs,
		RunE:  update,
	}

	return cmd
}

func update(cmd *cobra.Command, args []string) error {
	handler := exifer.NewMetaMediaHandler()

	for _, imagePath := range args {
		metaFD, err := os.OpenFile(imagePath+".json", os.O_RDONLY, 0o400)
		if errors.Is(err, os.ErrNotExist) {
			cmd.Printf("tags file for %q not found, skipping update\n", imagePath)

			continue
		}

		if err != nil {
			return fmt.Errorf("failed to open tags file: %w", err)
		}

		tags, err := loadTags(metaFD)
		if err != nil {
			return fmt.Errorf("failed to load tags: %w", err)
		}

		builder, err := tags.ExifIFDBuilder()
		if err != nil {
			return fmt.Errorf("failed to convert tags to EXIF IFD builder: %w", err)
		}

		fd, err := os.OpenFile(imagePath, os.O_RDWR, 0o600)
		if err != nil {
			return err
		}
		defer fd.Close()

		info, err := fd.Stat()
		if err != nil {
			return fmt.Errorf("failed to stat image file: %w", err)
		}

		cmd.Printf("updating %q\n", imagePath)

		err = handler.UpdateImageWithEXIF(fd, int(info.Size()), builder)
		if err != nil {
			return fmt.Errorf("failed to update image EXIF: %w", err)
		}
	}

	return nil
}
