package detection

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/Gophercraft/core/version"
	pe "github.com/Velocidex/go-pe"
)

// failing a hash-lookup
// attempt to read PE file metadata
func detect_exe_build(path string) (build version.Build, err error) {
	// open fd
	var (
		file *os.File
	)
	file, err = os.Open(path)
	if err != nil {
		return
	}

	// create hash of file
	sha256_hash := sha256.New()
	if _, err = io.Copy(sha256_hash, file); err != nil {
		return
	}
	sha256_digest := sha256_hash.Sum(nil)
	hash_string := hex.EncodeToString(sha256_digest)

	var found bool
	build, found = binary_hashes[hash_string]
	if found {
		// if a build is associated with this file hash,
		// return it
		return
	}

	// otherwise, we can see if this PE file has metadata
	// as to what the version is

	file.Seek(0, io.SeekStart)

	pe_file, err := pe.NewPEFile(file)
	if err != nil {
		return 0, err
	}

	vinfo := pe_file.VersionInformation["FileVersion"]
	elements := strings.Split(vinfo, ", ")

	tail := elements[len(elements)-1]

	i, err := strconv.ParseInt(tail, 0, 32)
	if err != nil {
		return 0, err
	}

	return version.Build(i), nil
}
