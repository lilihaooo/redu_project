package model

type Shop struct {
	ID               string `json:"id"`
	SellNum          string `json:"sell_num"`
	Name             string `json:"name"`
	Logo             string `json:"logo"`
	CreaditLogistics string `json:"creadit_logistics"`
	CreaditProduct   string `json:"creadit_product"`
	CreaditService   string `json:"creadit_service"`
}
