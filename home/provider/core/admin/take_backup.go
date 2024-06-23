package admin

import (
	"errors"
	"io"
	"os"
	"path/filepath"

	"github.com/Gophercraft/core/home/protocol/pb/admin"
)

func (provider *admin_provider) TakeBackup(take_backup_request *admin.TakeBackupRequest, stream admin.AdminService_TakeBackupServer) (err error) {
	_, err = provider.validate_admin_access(stream.Context(), take_backup_request.Credential)
	if err != nil {
		return
	}

	backup_directory := filepath.Join(os.TempDir(), "gophercraft_home_backups")

	if err = os.MkdirAll(backup_directory, 0755); err != nil {
		return
	}

	var (
		backup_path string
		fi          os.FileInfo
		file        *os.File
	)

	backup_path, err = provider.home_db.TakeBackup(backup_directory)
	if err != nil {
		return
	}

	fi, err = os.Stat(backup_path)
	if err != nil {
		return
	}

	size := fi.Size()

	// Send file header
	var file_header admin.FileHeader
	file_header.Name = filepath.Base(backup_path)
	file_header.Size = uint64(size)

	if err = stream.Send(&admin.FileDownload{
		Data: &admin.FileDownload_Header{
			Header: &file_header,
		},
	}); err != nil {
		return
	}

	const chunk_size int64 = 128000

	chunk_data := make([]byte, chunk_size)

	file, err = os.Open(backup_path)
	if err != nil {
		return
	}

	for filepart := int64(0); filepart < size; {
		// offset_low := filepart
		// offset_hi := filepart + chunk_size
		// if offset_hi > size {
		// 	offset_hi = size
		// }
		var chunk_read int
		chunk_read, err = file.Read(chunk_data)
		if err != nil {
			if errors.Is(err, io.EOF) {
				// TODO: free old backups?
				err = nil
				return
			}

			return
		}

		if err = stream.Send(&admin.FileDownload{
			Data: &admin.FileDownload_Chunk{
				Chunk: &admin.FileChunk{
					Data: chunk_data[:chunk_read],
				},
			},
		}); err != nil {
			return
		}

		filepart += int64(chunk_read)
	}

	return
}
