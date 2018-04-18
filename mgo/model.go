package main

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

// Collection ProductCategory
type ProductCategory struct {
	Id            bson.ObjectId  `json:"id" bson:"_id,omitempty"`
	NameEn        string         `json:"nameEn,omitempty" bson:"name_en,omitempty"`
	NameCn        string         `json:"nameCn,omitempty" bson:"name_cn,omitempty"`
	DescriptionEn string         `json:"descriptionEn,omitempty" bson:"description_en,omitempty"`
	DescriptionCn string         `json:"descriptionCn,omitempty" bson:"description_cn,omitempty"`
	Icon          string         `json:"icon,omitempty" bson:"icon,omitempty"`
	IconGrey      string         `json:"iconGrey,omitempty" bson:"icon_grey,omitempty"`
	MinQuantity   int64          `json:"minQuantity,omitempty" bson:"min_quantity,omitempty"`
	ProductsInfos []*ProductInfo `json:"productsInfos,omitempty" bson:"products_infos,omitempty"`
}
type ProductInfo struct {
	ProductID bson.ObjectId `json:"productID,omitempty" bson:"product_id,omitempty"`
	Sort      int64         `json:"sort,omitempty" bson:"sort,omitempty"`
}

// Collection Product
type Product struct {
	Id                bson.ObjectId      `json:"id" bson:"_id,omitempty"`
	NameEn            string             `json:"nameEn,omitempty" bson:"name_en,omitempty"`
	NameCn            string             `json:"nameCn,omitempty" bson:"name_cn,omitempty"`
	DescriptionEn     string             `json:"descriptionEn,omitempty" bson:"description_en,omitempty"`
	DescriptionCn     string             `json:"descriptionCn,omitempty" bson:"description_cn,omitempty"`
	Img               string             `json:"img,omitempty" bson:"img,omitempty"`
	PriceFixed        int64              `json:"priceFixed,omitempty" bson:"price_fixed,omitempty"`
	PriceRate         int64              `json:"priceRate,omitempty" bson:"price_rate,omitempty"`
	PackingFee        int64              `json:"packingFee,omitempty" bson:"packing_fee,omitempty"`
	MinQuantity       int64              `json:"minQuantity,omitempty" bson:"min_quantity,omitempty"`
	MonthlySales      int64              `json:"monthlySales,omitempty" bson:"monthly_sales,omitempty"`
	Rating            int64              `json:"rating,omitempty"`
	ProductCategories []*ProductCategory `json:"productCategories,omitempty" bson:"product_categories,omitempty"`
	Items             []*Item            `json:"items,omitempty" bson:"items,omitempty"`
	//ProductCategoriesId []bson.ObjectId `json:"productCategoriesId" bson:"product_categories_id,omitempty"`
	//ItemsId             []bson.ObjectId `json:"itemId" bson:"item_id,omitempty"`

}

// Collection Item
type Item struct {
	Id                      bson.ObjectId           `json:"id" bson:"_id,omitempty"`
	NameEn                  string                  `json:"nameEn,omitempty" bson:"name_en,omitempty"`
	NameCn                  string                  `json:"nameCn,omitempty" bson:"name_cn,omitempty"`
	DescriptionEn           string                  `json:"descriptionEn,omitempty" bson:"description_en,omitempty"`
	DescriptionCn           string                  `json:"descriptionCn,omitempty" bson:"description_cn,omitempty"`
	Img                     string                  `json:"img,omitempty" bson:"img,omitempty"`
	ContractPrice           int64                   `json:"contractPrice,omitempty" bson:"contract_price,omitempty"`
	SalePrice               int64                   `json:"salePrice,omitempty" bson:"sale_price,omitempty"`
	Stock                   int64                   `json:"stock,omitempty" bson:"stock,omitempty"`
	DefaultStock            int64                   `json:"defaultStock,omitempty" bson:"default_stock,omitempty"`
	StockReset              *bool                   `json:"stockReset,omitempty" bson:"stock_reset,omitempty"`
	StockResetCycle         int64                   `json:"stockResetCycle,omitempty" bson:"stock_reset_cycle,omitempty"`
	StockResetTime          time.Time               `json:"stockResetTime,omitempty" bson:"stock_reset_time,omitempty"`
	SpecificationCategories []SpecificationCategory `json:"specificationCategories,omitempty" bson:"specification_categories,omitempty"`
}
type SpecificationCategory struct {
	NameEn           string          `json:"nameEn,omitempty" bson:"name_en,omitempty"`
	NameCn           string          `json:"nameCn,omitempty" bson:"name_cn,omitempty"`
	DescriptionEn    string          `json:"descriptionEn,omitempty" bson:"description_en,omitempty"`
	DescriptionCn    string          `json:"descriptionCn,omitempty" bson:"description_cn,omitempty"`
	Required         bool            `json:"required,omitempty" bson:"required,omitempty"`
	Multiple         bool            `json:"multiple,omitempty" bson:"multiple,omitempty"`
	MinQuantity      int64           `json:"minQuantity,omitempty" bson:"min_quantity,omitempty"`
	MaxQuantity      int64           `json:"maxQuantity,omitempty" bson:"max_quantity,omitempty"`
	Sort             int64           `json:"sort,omitempty" bson:"sort,omitempty"`
	IsLocked         bool            `json:"isLocked,omitempty" bson:"is_locked,omitempty"`
	SpecificationsId []bson.ObjectId `json:"specificationsId,omitempty" bson:"specifications_id,omitempty"`
}

// Collection Specification
type Specification struct {
	Id          bson.ObjectId   `json:"id" bson:"_id,omitempty"`
	NameEn      string          `json:"nameEn,omitempty" bson:"name_en,omitempty"`
	NameCn      string          `json:"nameCn,omitempty" bson:"name_cn,omitempty"`
	Price       int64           `json:"price,omitempty" bson:"price,omitempty"`
	IsAvailable bool            `json:"isAvailable,omitempty" bson:"is_available,omitempty"`
	IsLocked    bool            `json:"isLocked,omitempty" bson:"is_locked,omitempty"`
	Items       []bson.ObjectId `json:"itemsId,omitempty" bson:"items_id,omitempty"`
}
