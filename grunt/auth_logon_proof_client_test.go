package grunt_test

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/Gophercraft/core/grunt"
	"github.com/davecgh/go-spew/spew"
)

func TestAuth(t *testing.T) {
	const client_proof_hex = `01429d3dbafc13bcd95daa3020f7860d947ab6788413c277284ee74da342d504372acfe08e2d92470c4279dd5b479229700a3303fb738cbd9fad90defa934d110ace3aacd346a4a51e0000`

	client_proof_bytes, _ := hex.DecodeString(client_proof_hex)
	reader := bytes.NewReader(client_proof_bytes)

	var client_proof grunt.AuthLogonProof_Client
	err := grunt.ReadAuthLogonProof_Client(reader, &client_proof)
	if err != nil {
		panic(err)
	}

	fmt.Println(spew.Sdump(&client_proof))
}
