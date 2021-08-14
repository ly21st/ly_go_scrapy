package utils

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func ZipMultiFile(srcFileList []string, dstZip string) error {
	zipFile, err := os.Create(dstZip)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	archive := zip.NewWriter(zipFile)
	defer archive.Close()

	for _, srcFile := range srcFileList {
		if srcFile == " " || srcFile == "" {
			continue
		}
		ZipHelper(srcFile, archive)
	}
	return nil
}

// srcFile could be a single file or a directory
func ZipHelper(srcFile string, archive *zip.Writer) error {
	var err error

	filepath.Walk(srcFile, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		//header.Name = strings.TrimPrefix(path, filepath.Dir(srcFile)+"/")
		header.Name = filepath.Base(path)
		// header.Name = path
		header.Method = zip.Deflate

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(writer, file)

		return err
	})

	return err
}
