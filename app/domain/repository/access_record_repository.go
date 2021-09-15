package repository

import "WorkerPlace/app/domain/entity"

type AccessRecordRepository interface {
	GetAccessRecordDocument(idDocument string) (*entity.AccessRecordDocument,error)
	SaveAccessRecordDocument(record *entity.AccessRecordDocument) error
	UpdateAccessRecordDocument(record *entity.AccessRecordDocument) error
}
