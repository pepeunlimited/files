package validator

type FileUploadServerValidator struct {}

func NewFileUploadServerValidator() FileUploadServerValidator {
	return FileUploadServerValidator{}
}

func (FileUploadServerValidator) UploadFile() error {
	return nil
}