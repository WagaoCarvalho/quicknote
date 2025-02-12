package validations

import "regexp"

type FormValidator struct {
	FieldErrors map[string]string
	Flash       string
}

func (fv *FormValidator) Valid() bool {
	return len(fv.FieldErrors) == 0
}

func (fv *FormValidator) AddFieldError(field, message string) {
	if fv.FieldErrors == nil {
		fv.FieldErrors = make(map[string]string)
	}
	fv.FieldErrors[field] = message
}

func (fv *FormValidator) IsEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}
