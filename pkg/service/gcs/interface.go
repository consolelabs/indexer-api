package gcs

type IService interface {
	UploadObject(object, filePath string) (string, error)
}
