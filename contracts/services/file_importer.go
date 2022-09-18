package services

type FileImportServiceInterface interface{
	ProcessFile(filePath string) (string, []error)
}
