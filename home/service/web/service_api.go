package web

func (service *Service) mount_api() {
	service.web_serve_mux.HandleFunc("GET /api/v1/version", service.handle_get_version_info)
	service.web_serve_mux.HandleFunc("GET /api/v1/service_addresses", service.handle_get_service_addresses)

	service.web_serve_mux.HandleFunc("GET /api/v1/captcha/{captcha_id}", service.handle_get_captcha)

	service.web_serve_mux.HandleFunc("GET /api/v1/register", service.handle_get_registration_challenge)
	service.web_serve_mux.HandleFunc("POST /api/v1/register", service.handle_post_registration_request)

	service.web_serve_mux.HandleFunc("GET /api/v1/login", service.handle_get_login_challenge)
	service.web_serve_mux.HandleFunc("POST /api/v1/login", service.handle_post_login_request)
	service.web_serve_mux.HandleFunc("GET /api/v1/logout", service.handle_get_logout_request)

	service.web_serve_mux.HandleFunc("GET /api/v1/credential", service.handle_get_credential_status)

	service.web_serve_mux.HandleFunc("GET /api/v1/2fa/methods", service.handle_get_2fa_methods)
	service.web_serve_mux.HandleFunc("POST /api/v1/2fa/authenticate", service.handle_post_authenticate_credential)
	service.web_serve_mux.HandleFunc("POST /api/v1/2fa/enroll", service.handle_post_2fa_enroll)
	service.web_serve_mux.HandleFunc("POST /api/v1/2fa/send_email_code", service.handle_post_send_2fa_email_code)

	service.web_serve_mux.HandleFunc("GET /api/v1/account", service.handle_get_account_status)

	service.web_serve_mux.HandleFunc("PUT /api/v1/game_account", service.handle_put_new_game_account)
	service.web_serve_mux.HandleFunc("POST /api/v1/game_account/{id}/rename", service.handle_post_rename_game_account)
	service.web_serve_mux.HandleFunc("POST /api/v1/game_account/{id}/activate", service.handle_post_activate_game_account)
	service.web_serve_mux.HandleFunc("DELETE /api/v1/game_account/{id}", service.handle_delete_game_account)

	service.web_serve_mux.HandleFunc("GET /api/v1/realm/status", service.handle_get_realm_status)
}
