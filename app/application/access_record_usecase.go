package application

import (
	"WorkerPlace/app/domain/entity"
	"WorkerPlace/app/domain/repository"
	"fmt"
)

type AccessRecordUseCase struct {
	accessRepository repository.AccessRecordRepository
}

func NewAccessRecordUseCase(accessRecordRepository repository.AccessRecordRepository) *AccessRecordUseCase {
	return &AccessRecordUseCase{accessRepository: accessRecordRepository}
}

func (c AccessRecordUseCase) CreateOrUpdateAccessRecord(record *entity.AccessRecord) {
	document, err := c.accessRepository.GetAccessRecordDocument(record.Id)
	if err != nil {
		document = entity.NewAccessRecordDocument(record)
		err = c.accessRepository.SaveAccessRecordDocument(document)
		if err != nil {
			fmt.Println("unable to save access record")
		}
	}else{
		UpdateAccessRecord(record, document)
		err = c.accessRepository.UpdateAccessRecordDocument(document)
		if err != nil {
			fmt.Println("unable to update access record")
		}
	}

	fmt.Printf("document %s saved\n", document.Id)
}

func UpdateAccessRecord(record *entity.AccessRecord, document *entity.AccessRecordDocument)  {
	document.CurrentPlace = record.Place
	document.LastAccessTime = record.AccessTime

	lastIndex := len(document.Places)-1
	document.Places[lastIndex].OutDate = record.AccessTime
	document.Places[lastIndex].DurationInSecond = document.Places[lastIndex].OutDate.Sub(document.Places[lastIndex].InDate).Seconds()

	document.Places = append(document.Places, entity.NewAccessPlace(record) )
}
