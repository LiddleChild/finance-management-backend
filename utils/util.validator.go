package utils

import (
	"fmt"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var validate *validator.Validate
var translator ut.Translator

func GetValidator() *validator.Validate {
	if validate == nil {
		en := en.New()
		uni := ut.New(en, en)
	
		translator, _ = uni.GetTranslator("en")
	
		validate = validator.New()
		en_translations.RegisterDefaultTranslations(validate, translator)
	}

	return validate
}

func TranslateError(err error) (errs []error) {
  if err == nil {
    return nil
  }

  validatorErrs := err.(validator.ValidationErrors)
  for _, e := range validatorErrs {
    translatedErr := fmt.Errorf(e.Translate(translator))
    errs = append(errs, translatedErr)
  }

  return errs
}

func ErrorsToString(err []error) (str []string) {
	for _, i := range err {
		str = append(str, i.Error())
	}

	return str
}