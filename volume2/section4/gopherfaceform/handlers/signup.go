package handlers

import (
	"net/http"

	"github.com/EngineerKamesh/gofullstack/volume2/section4/gopherfaceform/validationkit"
)

type SignUpForm struct {
	FieldNames []string
	Fields     map[string]string
	Errors     map[string]string
}

// DisplaySignUpForm displays the Sign Up form
func DisplaySignUpForm(w http.ResponseWriter, r *http.Request, s *SignUpForm) {
	RenderTemplate(w, "./templates/signupform.html", s)
}

func DisplayConfirmation(w http.ResponseWriter, r *http.Request, s *SignUpForm) {
	RenderUnsafeTemplate(w, "./templates/signupconfirmation.html", s)
}

func PopulateFormFields(r *http.Request, s *SignUpForm) {

	for _, fieldName := range s.FieldNames {
		s.Fields[fieldName] = r.FormValue(fieldName)
	}

}

// ValidateSignUpForm validates the Sign Up form's fields
func ValidateSignUpForm(w http.ResponseWriter, r *http.Request, s *SignUpForm) {

	PopulateFormFields(r, s)
	// Check if username was filled out
	if r.FormValue("username") == "" {
		s.Errors["usernameError"] = "The username field is required."
	}

	// Check if first name was filled out
	if r.FormValue("firstName") == "" {
		s.Errors["firstNameError"] = "The first name field is required."
	}

	// Check if last name was filled out
	if r.FormValue("lastName") == "" {
		s.Errors["lastNameError"] = "The last name field is required."
	}

	// Check if e-mail address was filled out
	if r.FormValue("email") == "" {
		s.Errors["emailError"] = "The e-mail address field is required."
	}

	// Check if e-mail address was filled out
	if r.FormValue("password") == "" {
		s.Errors["passwordError"] = "The password field is required."
	}

	// Check if e-mail address was filled out
	if r.FormValue("confirmPassword") == "" {
		s.Errors["confirmPasswordError"] = "The confirm password field is required."
	}

	// Check username syntax
	if validationkit.CheckUsernameSyntax(r.FormValue("username")) == false {

		usernameErrorMessage := "The username entered has an improper syntax."
		if _, ok := s.Errors["usernameError"]; ok {
			s.Errors["usernameError"] += " " + usernameErrorMessage
		} else {
			s.Errors["usernameError"] = usernameErrorMessage
		}
	}

	// Check e-mail address syntax
	if validationkit.CheckEmailSyntax(r.FormValue("email")) == false {
		emailErrorMessage := "The e-mail address entered has an improper syntax."
		if _, ok := s.Errors["usernameError"]; ok {
			s.Errors["emailError"] += " " + emailErrorMessage
		} else {
			s.Errors["emailError"] = emailErrorMessage
		}
	}

	// Check if passord and confirm password field values match
	if r.FormValue("password") != r.FormValue("confirmPassword") {
		s.Errors["confirmPasswordError"] = "The password and confirm pasword fields do not match."
	}

	if len(s.Errors) > 0 {
		DisplaySignUpForm(w, r, s)
	} else {
		ProcessSignUpForm(w, r, s)
	}

}

// ProcessSignUpForm
func ProcessSignUpForm(w http.ResponseWriter, r *http.Request, s *SignUpForm) {

	// If we reached this point, that indicates that we had a successful form submission.
	// Later, we will include form processing logic here, in this case that would be
	// inserting the information from the form as an entry into the database.

	// Display form confirmation message
	DisplayConfirmation(w, r, s)

}

func SignUpHandler(w http.ResponseWriter, r *http.Request) {

	s := SignUpForm{}
	s.FieldNames = []string{"username", "firstName", "lastName", "email"}
	s.Fields = make(map[string]string)
	s.Errors = make(map[string]string)

	switch r.Method {

	case "GET":
		DisplaySignUpForm(w, r, &s)
	case "POST":
		ValidateSignUpForm(w, r, &s)
	default:
		DisplaySignUpForm(w, r, &s)
	}

}
