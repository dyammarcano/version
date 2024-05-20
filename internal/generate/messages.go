package generate

import "errors"

const (
	GeneratingVersionFile  = "generating version file: %s"
	ErrorClosingFile       = "error closing file: %g"
	ErrorCreatingFile      = "error creating file: %w"
	ErrorEncodingJSON      = "error encoding json: %w"
	ErrorGettingTag        = "error getting tag: %w"
	ErrorGettingTags       = "error getting tags: %w"
	ErrorParsingTemplate   = "error parsing template: %w"
	ErrorExecutingTemplate = "error executing template: %w"
)

func OperationCanceled() error {
	return errors.New("operation canceled")
}
