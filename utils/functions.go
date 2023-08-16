package utils

import (
	"crypto/sha256"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"log"
	"os"
	"runtime"
)

var (
	ChecksumFailed = errors.New("checksum validation failed")
)

func GetOSPlatform() string {
	return fmt.Sprintf("%s-%s", runtime.GOOS, runtime.GOARCH)
}

func ValidateChecksum(filePath, checksum string) error {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = f.Close()
	}()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		err = fmt.Errorf("copying file content: %w", err)
		return err
	}

	if fmt.Sprintf("%x", h.Sum(nil)) != checksum {
		return ChecksumFailed
	}
	return nil
}
