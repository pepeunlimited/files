package validator

type FileServerValidator struct {}

func NewFileServerValidator() FileServerValidator {
	return FileServerValidator{}
}

func (FileServerValidator) CreateFile() error {
	return nil
}
