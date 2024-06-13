package job

import "canal_kafka_job/models"

type ProductMassage struct {
	Data     []models.Product `json:"data"`
	IsDdl    bool             `json:"isDdl"`
	Database string           `json:"database"`
	Table    string           `json:"table"`
	Type     string           `json:"type"`
}
