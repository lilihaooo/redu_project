package product

import (
	"redu/global"
	"redu/model"
)

func (Product) MakeCategoryTree() (tree []model.Category, err error) {
	var m model.Category
	categories, err := m.GetAll()
	if err != nil {
		return nil, err
	}
	tree = buildTree(categories, "0")
	return tree, nil
}

func buildTree(all []model.Category, parentID string) (tree []model.Category) {
	for _, menu := range all {
		if menu.ParentID == parentID {
			subMenu := buildTree(all, menu.ID)
			menu.Childer = subMenu
			tree = append(tree, menu)
		}
	}
	return
}

func (Product) GetIDsByLevel1(level1ID string) (level3IDs []string) {
	// 找到所有二级分类
	var category model.Category
	var level2IDs []string
	err := global.DB.Model(category).Where("parent_id = ?", level1ID).Select("id").Scan(&level2IDs).Error
	if err != nil {
		global.Logrus.Error(err)
		return nil
	}
	err = global.DB.Model(category).Where("parent_id in ?", level2IDs).Select("id").Scan(&level3IDs).Error
	if err != nil {
		global.Logrus.Error(err)
		return nil
	}
	return
}

func (Product) GetIDsByLevel2(level2ID string) (level3IDs []string) {
	// 找到所有三级分类
	var category model.Category
	err := global.DB.Model(category).Where("parent_id = ?", level2ID).Select("id").Scan(&level3IDs).Error
	if err != nil {
		global.Logrus.Error(err)
		return nil
	}
	return
}
