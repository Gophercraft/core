package datapack

import (
	"archive/zip"
	"io"
)

func add_file_to_zip(zip_writer *zip.Writer, filepath string, file io.Reader) error {
	header := &zip.FileHeader{}

	header.Name = filepath
	header.Method = zip.Deflate

	writer, err := zip_writer.CreateHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, file)
	return err
}
