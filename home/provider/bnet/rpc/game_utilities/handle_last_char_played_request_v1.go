package game_utilities

import (
	"context"

	game_utilities_v1 "github.com/Gophercraft/core/bnet/pb/bgs/protocol/game_utilities/v1"
)

func handle_last_char_played_request_v1(service *Service, ctx context.Context, parameters ClientRequestParameters, response *game_utilities_v1.ClientResponse) (err error) {
	sub_region := parameters.Get("Command_LastCharPlayedRequest_v1_b9")

	if sub_region != nil {
		// todo: implement
	}

	return
}
