package game_utilities

import (
	"context"
	"strings"

	"github.com/Gophercraft/core/bnet/pb/bgs/protocol"
	game_utilities_v1 "github.com/Gophercraft/core/bnet/pb/bgs/protocol/game_utilities/v1"
	"github.com/Gophercraft/core/bnet/rpc/codes"
	"github.com/Gophercraft/core/bnet/rpc/status"
	"github.com/Gophercraft/core/bnet/util"
	"github.com/Gophercraft/core/home/provider/bnet/rpc/authentication"
	"github.com/Gophercraft/phylactery/database"
)

type ClientRequestHandlerFunc func(service *Service, ctx context.Context, parameters ClientRequestParameters, response *game_utilities_v1.ClientResponse) (err error)

type Service struct {
	game_utilities_v1.UnimplementedGameUtilitiesServiceServer
	home_db                 *database.Container
	client_request_handlers map[string]ClientRequestHandlerFunc
}

func New(home_db *database.Container) (service *Service) {
	service = new(Service)
	service.home_db = home_db
	service.client_request_handlers = map[string]ClientRequestHandlerFunc{
		"Command_RealmListTicketRequest_v1_b9": handle_realmlist_ticket_request_v1,
		"Command_LastCharPlayedRequest_v1_b9":  handle_last_char_played_request_v1,
		"Command_RealmListRequest_v1_b9":       handle_realmlist_request_v1,
		"Command_RealmJoinRequest_v1_b9":       handle_realm_join_request_v1,
	}
	return
}

func (service *Service) ProcessClientRequest(ctx context.Context, request *game_utilities_v1.ClientRequest) (response *game_utilities_v1.ClientResponse, err error) {
	var (
		command_attribute *protocol.Attribute
		parameters        = make(map[string]*protocol.Variant)
	)

	for i := range request.Attribute {
		attribute := request.Attribute[i]
		if strings.HasPrefix(attribute.GetName(), "Command_") {
			command_attribute = attribute
		} else {
			parameters[attribute.GetName()] = attribute.GetValue()
		}
	}

	// a client request must have a command
	if command_attribute == nil {
		err = status.Errorf(codes.ERROR_RPC_MALFORMED_REQUEST, "client request lacks a Command attribute")
		return
	}

	command_name := command_attribute.GetName()

	// lookup client handler
	client_request_handler, ok := service.client_request_handlers[command_name]
	if !ok {
		err = status.Errorf(codes.ERROR_RPC_NOT_IMPLEMENTED, "client request handler for '%s' is not implemented", command_name)
		return
	}

	response = new(game_utilities_v1.ClientResponse)

	err = client_request_handler(service, ctx, parameters, response)
	return
}

func (service *Service) GetAllValuesForAttribute(ctx context.Context, request *game_utilities_v1.GetAllValuesForAttributeRequest) (response *game_utilities_v1.GetAllValuesForAttributeResponse, err error) {
	auth_details := authentication.GetDetails(ctx)
	if auth_details == nil {
		err = status.Errorf(codes.ERROR_DENIED, "no auth details")
		return
	}

	response = new(game_utilities_v1.GetAllValuesForAttributeResponse)

	switch request.GetAttributeKey() {
	case "Command_RealmListRequest_v1_b9":
		var attributevalue = new(protocol.Variant)
		util.Set(&attributevalue.StringValue, NewRealmAddress(1, 1, 0).String())
		response.AttributeValue = append(response.AttributeValue, attributevalue)
	// TODO implement
	default:
		err = status.Errorf(codes.ERROR_RPC_NOT_IMPLEMENTED, "don't know how to get values for attribute key '%s'", request.GetAttributeKey())
	}

	return
}
