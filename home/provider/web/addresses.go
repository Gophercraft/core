package web

import (
	api_models "github.com/Gophercraft/core/home/service/web/models"
)

func (provider *service_provider) GetServiceAddresses(user_info *api_models.UserInfo) (service_addresses *api_models.ServiceAddresses, err error) {
	service_addresses = &api_models.ServiceAddresses{
		Addresses: make(map[string]string),
	}

	for service_id, address := range provider.config.ServiceAddresses {
		service_addresses.Addresses[service_id.String()] = address
	}

	return
}
