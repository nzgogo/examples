package main

import (
	"time"
	"github.com/globalsign/mgo/bson"
)

var (
	vfalse = false
	vtrue  = true

	items = []Item{
		{
			Id:				bson.NewObjectId(),
			NameEn:          "nianweilapin",
			NameCn:          "年味儿腊拼",
			ContractPrice:   62,
			SalePrice:       55,
			Stock:           3,
			DefaultStock:    5,
			StockReset:      &vfalse,
			StockResetCycle: 12,
			StockResetTime:  time.Now(),
			SpecificationCategories: []SpecificationCategory {
				{
					NameEn         :"Flavor",
					NameCn         :"口味",
					Required       :true,
					Multiple       : true,
					MinQuantity    :2,
					MaxQuantity    :10,
					IsLocked       :true,
					SpecificationsId :[]bson.ObjectId{
						bson.ObjectIdHex("5ad5d8411a80a30a43066832"),
						bson.ObjectIdHex("5ad5d8411a80a30a43066835"),
					},
				},
			},
		},
		{
			Id:				bson.NewObjectId(),
			NameEn:          "55nianchaunjialuya",
			NameCn:          "55年传家卤鹅",
			ContractPrice:   70,
			SalePrice:       60,
			Stock:           3,
			DefaultStock:    5,
			StockReset:      &vfalse,
			StockResetCycle: 12,
			StockResetTime:  time.Now(),
			SpecificationCategories: []SpecificationCategory {
				{
					NameEn         :"Size",
					NameCn         :"分量",
					Required       :true,
					Multiple       : true,
					MinQuantity    :2,
					MaxQuantity    :10,
					IsLocked       :true,
					SpecificationsId :[]bson.ObjectId{
						bson.ObjectIdHex("5ad5d8411a80a30a43066832"),
						bson.ObjectIdHex("5ad5d8411a80a30a43066835"),
					},
				},
			},
		},
		{
			NameEn:          "huobianzi",
			NameCn:          "火边子牛肉蔬果沙拉",
			ContractPrice:   70,
			SalePrice:       55,
			Stock:           5,
			DefaultStock:    5,
			StockReset:      &vtrue,
			StockResetCycle: 12,
			StockResetTime:  time.Now(),
			SpecificationCategories: []SpecificationCategory {
				{
					NameEn         :"Flavor",
					NameCn         :"口味",
					Required       :true,
					Multiple       : true,
					MinQuantity    :2,
					MaxQuantity    :10,
					IsLocked       :true,
					SpecificationsId :[]bson.ObjectId{
						bson.ObjectIdHex("5ad5d8411a80a30a43066832"),
						bson.ObjectIdHex("5ad5d8411a80a30a43066835"),
					},
				},
			},
		},
		{
			NameEn:          "coke",
			NameCn:          "可乐",
			ContractPrice:   5,
			SalePrice:       3,
			Stock:           300,
			DefaultStock:    500,
			StockReset:      &vtrue,
			StockResetCycle: 12,
			StockResetTime:  time.Now(),
			SpecificationCategories: []SpecificationCategory {
				{
					NameEn         :"Size",
					NameCn         :"分量",
					Required       :true,
					Multiple       : true,
					MinQuantity    :1,
					MaxQuantity    :1,
					IsLocked       :true,
					//SpecificationsId :[]bson.ObjectId{
					//	bson.ObjectIdHex("5ad5d8411a80a30a43066832"),
					//	bson.ObjectIdHex("5ad5d8411a80a30a43066835"),
					//},
				},
			},
		},
	}

	specs = []Specification{
		{
			NameEn:      "",
			NameCn:      "小份",
			Price:       2,
			IsAvailable: false,
			IsLocked:    true,
		},
		{
			NameEn:      "big",
			NameCn:      "大份",
			Price:       3,
			IsAvailable: true,
			IsLocked:    false,
		},
		{
			NameEn:      "very spicy",
			NameCn:      "大辣",
			Price:       2,
			IsAvailable: true,
			IsLocked:    false,
		},
		{
			NameEn:      "little spicy",
			NameCn:      "小辣",
			Price:       3,
			IsAvailable: true,
			IsLocked:    false,
		},
	}

)
