package main

import (
	"errors"
	"regexp"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func ValidateRegexWithFormSubmitDisable(regexpstr, reason string, submit_btn *widget.Button) fyne.StringValidator {
	expression, err := regexp.Compile(regexpstr)
	if err != nil {
		fyne.LogError("Regexp did not compile", err)
		return nil
	}

	err = errors.New(reason)
	return func(text string) error {
		if expression != nil && !expression.MatchString(text) {
			submit_btn.Disable()
      return err
		}
    
    submit_btn.Enable()
		return nil // Nothing to validate with, same as having no validator.
	}
}
