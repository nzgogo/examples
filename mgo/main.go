package main

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/nzgogo/micro/db"
	"log"
	"reflect"
	"fmt"
)

const (
	url    = "mongodb://qiekai:pbklYdJ4WrP7X3GZ@gogotest-shard-00-00-ntrrw.mongodb.net:27017,gogotest-shard-00-01-ntrrw.mongodb.net:27017,gogotest-shard-00-02-ntrrw.mongodb.net:27017/test?ssl=true&replicaSet=GOGOTEST-shard-0&authSource=admin"
	TESTDB = "test"
)

var (
	ProdCatC *db.MicroCollect
	ProductC *db.MicroCollect
	ItemC    *db.MicroCollect
	SpecifC  *db.MicroCollect
)

func main() {
	// Create Mgo Session
	mgoSession := db.NewMongoDB(url)

	log.Println("Connecting to MongoDB...")
	if err := mgoSession.Connect(); err != nil {
		log.Fatal("Failed to connect to MongoDB...", err.Error())
	}
	defer mgoSession.Close()

	// Connect to specific Database
	mgoDB := mgoSession.DB(TESTDB)
	//ProdCatC = mgoDB.C("ProductCategory")
	//ProductC = mgoDB.C("Product")
	//SpecifC = mgoDB.C("Specification")
	ItemC = mgoDB.C("Item")
	ItemC.DropCollection()
	ItemC = mgoDB.C("Item")
	// insertion
	if err := ItemC.Insert(items[0], items[1], items[2], items[3]); err != nil {
		log.Fatal("Insert does not pass " + err.Error())
	}

	// Count
	if cnt, err := ItemC.Count(); err != nil || cnt != len(items) {
		log.Fatal("Count[1] does not pass")
	}

	// find
	query := bson.M{"stock_reset": false}
	results := []bson.M{}
	result := bson.M{}

	if err := ItemC.Find(query).All(&results); err != nil {
		log.Fatal("Find[1] not pass " + err.Error())
	} else if len(results) != 2 {
		log.Fatal("Find[2] not pass")
	}


	//if err := ItemC.Find(item).One(&); err != nil {
	//	log.Fatal("Find not pass " + err.Error())
	//}


	if err := ItemC.FindId(results[0]["_id"]).One(&result); err != nil {
		log.Fatal("FindId[1] not pass " + err.Error())
	} else if !reflect.DeepEqual(results[0], result) {
		ItemC.DropCollection()
		log.Println(results[0])
		log.Println(result)
		log.Fatal("FindId[2] not pass. not equal ")
	}

	// remove
	selector := bson.M{}
	b, _ := bson.Marshal(items[0])
	bson.Unmarshal(b, selector)
	//log.Println(selector)
	if err := ItemC.Remove(selector); err != nil {
		log.Fatal("Remove[1] does not pass " + err.Error())
	}
	if err := ItemC.Find(selector).One(result); err != mgo.ErrNotFound {
		log.Fatal("Remove[2] not pass")
	}
	if err := ItemC.FindWithTrash(selector).One(result); err != nil {
		log.Fatal("Remove[3] not pass " + err.Error())
	}
	if err := ItemC.RemoveId(result["_id"]); err != mgo.ErrNotFound {
		log.Fatal("RemoveId[1] does not pass ")
	}
	if err := ItemC.UpdateWithTrash(result, selector); err != nil {
		log.Fatal("Remove[4] not pass " + err.Error())
	}
	if err := ItemC.RemoveId(result["_id"]); err != nil {
		log.Fatal("RemoveId[2] does not pass ")
	}
	if err := ItemC.Find(selector).One(result); err != mgo.ErrNotFound {
		log.Fatal("RemoveId[3] not pass ")
	}
	if err := ItemC.FindWithTrash(selector).One(result); err != nil {
		log.Fatal("RemoveId[4] not pass " + err.Error())
	}
	if err := ItemC.UpdateWithTrash(result, selector); err != nil {
		log.Fatal("Remove[4] not pass")
	}
	if cnt, err := ItemC.Count(); err != nil {
		log.Fatal("Count[2] does not pass " + err.Error())
	} else if cnt != len(items) {
		log.Fatal("RemoveId[5] not pass")
	}
	if info, err := ItemC.RemoveAll(nil); err != nil {
		log.Fatal("RemoveAll[1] does not pass " + err.Error())
	} else if info.Updated != len(items) {
		log.Fatal("RemoveAll[2] does not pass ")
	}
	if cnt, err := ItemC.Count(); err != nil {
		log.Fatal("Count[3] does not pass " + err.Error())
	} else if cnt != 0 {
		log.Fatal("RemoveAll[3] does not pass ")
	}

	// update
	if err := ItemC.FindWithTrash(nil).All(&results); err != nil {
		log.Fatal("FindWithTrash not pass " + err.Error())
	} else if len(results) != len(items) {
		log.Fatal("FindWithTrash not pass")
	}
	for i := 0; i < len(items)-1; i++ {
		if err := ItemC.Update(results[i], items[i]); err != mgo.ErrNotFound {
			log.Fatal("Update[1] not pass")
		}
		if err := ItemC.UpdateId(results[i]["_id"], items[i]); err != mgo.ErrNotFound {
			log.Fatal("UpdateId[1] not pass")
		}
		b, _ := bson.Marshal(items[i])
		bson.Unmarshal(b, result)
		if err := ItemC.UpdateWithTrash(results[i], result); err != nil {
			log.Fatal("UpdateWithTrash[1] not pass " + err.Error())
		}
		result["newField"] = "newField"
		if err := ItemC.Update(bson.M{"_id": results[i]["_id"]}, result); err != nil {
			log.Fatal("Update[2] not pass " + err.Error())
		}
		if err := ItemC.UpdateId(results[i]["_id"], result); err != nil {
			log.Fatal("UpdateId[2] not pass " + err.Error())
		}
	}
	if cnt, err := ItemC.Count(); err != nil {
		log.Fatal("Count[4] does not pass " + err.Error())
	} else if cnt != (len(items)-1) {
		log.Fatal("UpdateWithTrash[2] does not pass ")
	}

	//ItemC.UpdateParts()
	//
	if info,err := ItemC.UpdateAll(bson.M{"stock_reset_cycle":12},bson.M{"$set":bson.M{"stock_reset_cycle":13}}); err != nil {
		log.Fatal("UpdateAll[1] does not pass " + err.Error())
	} else if info.Updated != (len(items)-1) {
		log.Fatal("UpdateAll[2] does not pass ")
	}

	// IncrementUpdate
	//if info, err := ItemC.IncrementUpdateAll(); err != nil {
	//	log.Fatal("IncrementUpdateAll does not pass " + err.Error())
	//} else if info.Updated != len(items) {
	//	log.Fatal("IncrementUpdateAll does not pass. ChangedInfo does not match " )
	//}
	//// increment update
	//if err := ItemC.IncrementUpdate(); err !=nil {
	//	log.Fatal("IncrementUpdate does not pass " + err.Error())
	//}
	//if err := ItemC.IncrementUpdateId(); err != nil {
	//	log.Fatal("IncrementUpdateId does not pass " + err.Error())
	//}

	// IncreUpsert
	//ItemC.IncreUpsert()
	//ItemC.IncreUpsertId()

	//ItemC.DropCollection()
	fmt.Println("All tests passed")

}
