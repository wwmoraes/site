package vcard

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io"
	"net/url"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var ntagBytes = map[string]int{
	"203": 144, //nolint:gomnd
	"210": 48,  //nolint:gomnd
	"212": 128, //nolint:gomnd
	"213": 144, //nolint:gomnd
	"215": 504, //nolint:gomnd
	"216": 888, //nolint:gomnd
	"223": 144, //nolint:gomnd
	"224": 144, //nolint:gomnd
	"424": 256, //nolint:gomnd
	"426": 916, //nolint:gomnd
}

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vcard",
		Short: "generates a vCard URL hash",
		Long:  "compresses and encodes a vcf file content as a safe URL to share",
		RunE:  vCard,
		Args:  cobra.ExactArgs(1),
	}

	return cmd
}

func compress(r io.Reader) ([]byte, error) {
	var gzData bytes.Buffer

	gw := gzip.NewWriter(&gzData)

	_, err := io.Copy(gw, r)
	if err != nil {
		return nil, err
	}

	err = gw.Close()
	if err != nil {
		return nil, err
	}

	return gzData.Bytes(), nil
}

func base64encode(data []byte) (string, error) {
	var encodedStr strings.Builder

	base64Encode := base64.NewEncoder(base64.RawStdEncoding, &encodedStr)

	_, err := base64Encode.Write(data)
	if err != nil {
		return "", err
	}

	err = base64Encode.Close()
	if err != nil {
		return "", err
	}

	return encodedStr.String(), nil
}

func yesOrNo(yes bool) string {
	if yes {
		return "✅"
	}

	return "❌"
}

func vCard(cmd *cobra.Command, args []string) error {
	fd, err := os.Open(args[0])
	if err != nil {
		return err
	}
	defer fd.Close()

	gzData, err := compress(fd)
	if err != nil {
		return err
	}

	encodedStr, err := base64encode(gzData)
	if err != nil {
		return err
	}

	finalStr := url.QueryEscape(encodedStr)
	finalStrLen := len(finalStr)

	for name, size := range ntagBytes {
		cmd.Printf("NTAG %s: %s\n", name, yesOrNo(finalStrLen <= size))
	}

	cmd.Printf("https://artero.dev/vcard#%s", finalStr)

	return nil
}
