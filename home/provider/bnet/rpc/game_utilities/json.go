package game_utilities

import (
	"encoding/json"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var marshal_options = protojson.MarshalOptions{
	UseProtoNames: true,
}

func json_marshal(message proto.Message) ([]byte, error) {
	return marshal_options.Marshal(message)
}

func json_unmarshal(data []byte, message any) error {
	return json.Unmarshal(data, message)
}
