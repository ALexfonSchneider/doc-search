package entity

type NoDocumentInCacheErr struct {
	Key string
}

func (d *NoDocumentInCacheErr) Error() string {
	return "There is no documents in cache"
}

func (d *NoDocumentInCacheErr) Is(err error) bool {
	return d.Error() == err.Error()
}

type DocumentNotFoundErr struct {
	Id string
}

func (d *DocumentNotFoundErr) Error() string {
	return "Document not found"
}

func (d *DocumentNotFoundErr) Is(err error) bool {
	return d.Error() == err.Error()
}

type DocumentExistsErr struct {
	Id string
}

func (d *DocumentExistsErr) Error() string {
	return "Document already exist"
}

func (d *DocumentExistsErr) Is(err error) bool {
	return d.Error() == err.Error()
}
