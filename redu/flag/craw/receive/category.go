package receive

type Category struct {
	ID       string     `json:"ID"`
	Level    string     `json:"level"`
	ParentID string     `json:"parent_id"`
	Name     string     `json:"name"`
	Status   string     `json:"status"`
	Image    string     `json:"image"`
	Num      int        `json:"num"`
	Childer  []Category `json:"childer" gorm:"foreignkey:ParentID"`
}
