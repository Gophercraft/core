package rest

import (
	"github.com/Gophercraft/core/bnet/pb/login"
	"github.com/Gophercraft/core/bnet/util"
)

func make_form_input(input_id, input_type, label string, max_length uint32) (input *login.FormInput) {
	input = new(login.FormInput)
	util.Set(&input.InputId, input_id)
	util.Set(&input.Type, input_type)
	util.Set(&input.Label, label)

	if max_length > 0 {
		util.Set(&input.MaxLength, max_length)
	}

	return
}

func make_login_form_inputs() (inputs *login.FormInputs) {
	inputs = new(login.FormInputs)
	util.Set(&inputs.Type, login.FormType_LOGIN_FORM)
	inputs.Inputs = []*login.FormInput{
		make_form_input("account_name", "text", "E-mail", 320),
		make_form_input("password", "password", "Password", 128),
		make_form_input("log_in_submit", "submit", "Log In", 0),
	}
	return
}

func get_login_form_input(login_form *login.LoginForm, input_id string) (input string) {
	for _, value := range login_form.Inputs {
		if value.GetInputId() == input_id {
			input = value.GetValue()
			break
		}
	}

	return
}
