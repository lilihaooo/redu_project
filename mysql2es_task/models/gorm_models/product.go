package models

type Product struct {
	ID                uint    `json:"id"`
	BeginTime         int     `json:"begin_time"`
	EndTime           int     `json:"end_time"`
	Status            string  `json:"status"`
	XingxuanCredit    string  `json:"xingxuan_credit"`
	ProductCategoryId string  `json:"product_category_id"`
	SalesAll          string  `json:"sales_all"`
	Sales24           string  `json:"sales_24"` //  0   2.5w+
	Platform          string  `json:"platform"`
	Title             string  `json:"title"`
	Commission        string  `json:"commission"`
	BottomPrice       float64 `json:"bottom_price"`
	Url               string  `json:"url"`
	PlatformProductId string  `json:"platform_product_id"`
	SpecimenCount     string  `json:"specimen_count"`
	RewardBeginTime   string  `json:"reward_begin_time"`
	RewardEndTime     string  `json:"reward_end_time"`
	TrueSales         string  `json:"true_sales"`
	CommissionValue   float64 `json:"commission_value"`
	Remaining         int     `json:"remaining"`
	ExperienceScore   float64 `json:"experience_score"` // 体验分
	IsFake            int     `json:"is_fake"`          // 假数据
}
