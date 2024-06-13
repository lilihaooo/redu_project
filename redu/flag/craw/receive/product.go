package receive

import (
	"fmt"
	"math"
	"redu/model"
	"strconv"
)

type ProductReceive struct {
	ShopID               string
	Shop                 model.Shop            `json:"shop"`          // 迁移表时去
	ProductImages        []*model.ProductImage `json:"productImages"` // 迁移表时去掉
	BeginTime            int                   `json:"begin_time"`
	EndTime              int                   `json:"end_time"`
	ID                   string                `json:"id"`
	ShopCreaditLogistics string                `json:"shop_creadit_logistics"`
	ShopCreaditProduct   string                `json:"shop_creadit_product"`
	ShopCreaditService   string                `json:"shop_creadit_service"`
	Status               string                `json:"status"`
	XingxuanCredit       string                `json:"xingxuan_credit"`
	//SelfSupport          int                    `json:"self_support"`
	ProductCategoryId string `json:"product_category_id"`
	//SellNum           string `json:"sell_num"`
	//IsNewcomers       int         `json:"is_newcomers"`
	SalesAll    string      `json:"sales_all"`
	Sales24Hour string      `json:"sales24hour"`
	Sales24     interface{} `json:"sales_24"` // todo 0   2.5w+
	//IsGuarantee       string      `json:"is_guarantee"`
	Platform          string      `json:"platform"`
	Title             string      `json:"title"`
	Commission        string      `json:"commission"`
	BottomPrice       float64     `json:"bottom_price"`
	Url               string      `json:"url"`
	PlatformProductId string      `json:"platform_product_id"`
	SpecimenCount     string      `json:"specimen_count"`
	RewardBeginTime   interface{} `json:"reward_begin_time"`
	RewardEndTime     interface{} `json:"reward_end_time"`
	TrueSales         string      `json:"true_sales"`
	//memberSelectionCart
	//ProductTag           []string `json:"product_tag"`
	//ParseWord            []string `json:"parse_word"`
	IsDisplayreward string `json:"is_displayreward"`
	//IsNewyeartd     string  `json:"is_newyeartd"`
	CommissionValue float64 `json:"commission_value"`
	Remaining       int     `json:"remaining"`
}

func (p *ProductReceive) ProductReceive2model() model.Product {
	product := model.Product{}
	// 将接收者的值赋值给商品model
	product.ShopID = p.ShopID
	product.Shop = p.Shop
	product.ProductImages = p.ProductImages
	//product.BeginTime = p.BeginTime
	//product.EndTime = p.EndTime
	num, _ := strconv.ParseUint(p.ID, 10, 0) // 将id: str2uint
	product.ID = uint(num)
	//product.ShopCreaditLogistics = p.ShopCreaditLogistics
	//product.ShopCreaditProduct = p.ShopCreaditProduct
	//product.ShopCreaditService = p.ShopCreaditService
	product.Status = p.Status
	product.XingxuanCredit = p.XingxuanCredit
	//product.SelfSupport = p.SelfSupport
	product.ProductCategoryId = p.ProductCategoryId
	//product.SellNum = p.SellNum
	//product.IsNewcomers = p.IsNewcomers
	product.SalesAll = p.SalesAll
	//product.Sales24Hour = p.Sales24Hour
	// 将 i 转换为 string 类型
	sales24Str := ""
	switch v := p.Sales24.(type) {
	case int:
		// 将 int 转换为 string
		sales24Str = fmt.Sprintf("%d", v)
	case float64:
		// 将小数截断为整数
		truncatedValue := math.Trunc(v)
		// 将截断后的浮点型转换为 string，保留整数
		sales24Str = fmt.Sprintf("%v", truncatedValue)
	case string:
		// 值已经是 string 类型
		sales24Str = v
	default:
		// 如果值不是上述类型，则将其转换为 string "unknown"
		sales24Str = "unknown"
	}
	product.Sales24 = sales24Str
	//rewardBeginStr := ""
	//switch v := p.RewardBeginTime.(type) {
	//case int:
	//	// 将 int 转换为 string
	//	rewardBeginStr = fmt.Sprintf("%d", v)
	//case float64:
	//	// 将小数截断为整数
	//	truncatedValue := math.Trunc(v)
	//	// 将截断后的浮点型转换为 string，保留整数
	//	rewardBeginStr = fmt.Sprintf("%v", truncatedValue)
	//case string:
	//	// 值已经是 string 类型
	//	rewardBeginStr = v
	//default:
	//	// 如果值不是上述类型，则将其转换为 string "unknown"
	//	rewardBeginStr = "unknown"
	//}
	//rewardEndStr := ""
	//switch v := p.RewardEndTime.(type) {
	//case int:
	//	// 将 int 转换为 string
	//	rewardEndStr = fmt.Sprintf("%d", v)
	//case float64:
	//	// 将小数截断为整数
	//	truncatedValue := math.Trunc(v)
	//	// 将截断后的浮点型转换为 string，保留整数
	//	rewardEndStr = fmt.Sprintf("%v", truncatedValue)
	//case string:
	//	// 值已经是 string 类型
	//	rewardEndStr = v
	//default:
	//	// 如果值不是上述类型，则将其转换为 string "unknown"
	//	rewardEndStr = "unknown"
	//}
	//product.RewardBeginTime = rewardBeginStr
	//product.RewardEndTime = rewardEndStr
	//product.IsGuarantee = p.IsGuarantee
	product.Platform = p.Platform
	product.Title = p.Title
	product.Commission = p.Commission
	product.BottomPrice = p.BottomPrice
	//product.Url = p.Url
	//product.PlatformProductId = p.PlatformProductId
	//product.SpecimenCount = p.SpecimenCount
	product.TrueSales = p.TrueSales
	//product.IsDisplayreward = p.IsDisplayreward
	//product.IsNewyeartd = p.IsNewyeartd
	product.CommissionValue = p.CommissionValue
	product.Remaining = p.Remaining
	return product
}
