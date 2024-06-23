package wizard

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Gophercraft/core/home/protocol/pb/admin"
)

func (c *Configurator) TakeHomeBackup(path string) (backup_file string, err error) {
	if !c.Connected() {
		err = fmt.Errorf("wizard: must be connected to a home server to take a backup")
		return
	}

	admin_service_client := admin.NewAdminServiceClient(c.home_connection)

	var (
		stream admin.AdminService_TakeBackupClient
		file   *os.File
	)

	stream, err = admin_service_client.TakeBackup(context.Background(), &admin.TakeBackupRequest{})
	if err != nil {
		return
	}

	download_started := false

	var download *admin.FileDownload

	var (
		file_size      uint64
		bytes_received uint64
	)

	defer func() {
		if download_started {
			file.Close()
		}
	}()

	for {
		// if download is
		if download_started {
			if bytes_received >= file_size {
				return
			}
		}

		download, err = stream.Recv()
		if err != nil {
			return
		}

		switch data := download.GetData().(type) {
		case *admin.FileDownload_Header:
			if download_started {
				err = fmt.Errorf("wizard: server sent two file headers (?)")
				return
			}
			// if user didn't ask for a specific filename
			if path == "" {
				// make up a filepath based on what server sent
				path = filepath.Clean(data.Header.Name)
			}
			// check if file already exists
			if _, err = os.Stat(path); err == nil {
				err = fmt.Errorf("wizard: backup file '%s' already exists", path)
				return
			}
			backup_file = path
			file_size = data.Header.Size
			file, err = os.Create(path)
			if err != nil {
				return
			}
			download_started = true
		case *admin.FileDownload_Chunk:
			var i int
			i, err = file.Write(data.Chunk.Data)
			if err != nil {
				return
			}

			bytes_received += uint64(i)
		}
	}
}
