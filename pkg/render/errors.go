package render

import "errors"

var (
	errTemplateNotFound = errors.New("could not find the given template from template cache")
)
