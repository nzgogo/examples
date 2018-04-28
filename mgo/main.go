package main

import (
	"github.com/nzgogo/mgo/bson"
	"github.com/nzgogo/mgo"
	"log"
	"reflect"
	"fmt"
)

var (
	ItemC    *mgo.GCollect
)

func main() {
	// Create Mgo Session
	mgoSession := mgo.NewMongoDB(url)

	log.Println("Connecting to MongoDB...")
	if err := mgoSession.Connect(); err != nil {
		log.Fatal("Failed to connect to MongoDB...", err.Error())
	}
	defer mgoSession.Close()

	// Connect to specific Database
	mgoDB := mgoSession.DB(TESTDB)
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
	// remove
	selector := bson.M{}
	b, _ := bson.Marshal(items[0])
	bson.Unmarshal(b, selector)
	delete(selector,"specification_categories")

	results := []bson.M{}
	result := bson.M{}

	if err := ItemC.Collection.Find(selector).One(&result); err != nil {
		log.Fatal("Find[4] not pass " + err.Error())
	}

	query := bson.M{"stock_reset": false}
	if err := ItemC.Find(query).All(&results); err != nil {
		log.Fatal("Find[1] not pass " + err.Error())
	} else if len(results) != 2 {
		log.Fatal("Find[2] not pass")
	}

	if err := ItemC.Find(query).One(&result); err != nil {
		log.Fatal("Find[3] not pass " + err.Error())
	}

	if err := ItemC.FindId(results[0]["_id"]).One(&result); err != nil {
		log.Fatal("FindId[1] not pass " + err.Error())
	} else if !reflect.DeepEqual(results[0], result) {
		ItemC.DropCollection()
		log.Println(results[0])
		log.Println(result)
		log.Fatal("FindId[2] not pass. not equal ")
	}

	// remove
	b, _ = bson.Marshal(items[0])
	bson.Unmarshal(b, selector)
	delete(selector,"specification_categories")

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
	delete(result,"specification_categories")
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
	delete(result,"specification_categories")
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
		delete(results[i],"specification_categories")
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

	if info,err := ItemC.UpdateAll(bson.M{"stock_reset_cycle":12},bson.M{"$set":bson.M{"stock_reset_cycle":13}}); err != nil {
		log.Fatal("UpdateAll[1] does not pass " + err.Error())
	} else if info.Updated != (len(items)-1) {
		log.Fatal("UpdateAll[2] does not pass ")
	}

	// Upsert
	if err := ItemC.Collection.Find(bson.M{"name_en":"coke"}).One(&result); err !=nil {
		log.Fatal("Upsert[1] does not pass " + err.Error())
	}
	delete(result,"delete_at")
	delete(result,"_id")
	if _, err := ItemC.Upsert(bson.M{"name_en":"coke"}, result); err != nil {
		log.Fatal("Upsert[2] does not pass " + err.Error())
	}
	if err := ItemC.Find(bson.M{"name_en":"coke"}).One(&result); err !=nil {
		log.Fatal("Upsert[3] does not pass " + err.Error())
	}
	id:=result["_id"]
	delete(result,"_id")
	result["name_en"]="sprite"
	if _, err := ItemC.UpsertId(id, result); err != nil {
		log.Fatal("Upsert[4] does not pass " + err.Error())
	}
	if err := ItemC.FindId(id).One(&result); err !=nil {
		log.Fatal("Upsert[5] does not pass " + err.Error())
	}

	// IncrementUpdate
	if err := ItemC.Find(bson.M{"newField":"newField"}).One(&result); err !=nil {
		log.Fatal("IncrementUpdate[1] does not pass " + err.Error())
	}
	update := bson.M{"$set": bson.M{"contract_price": 1}}
	delete(result,"specification_categories")
	if err := ItemC.IncrementUpdate(result,update); err !=nil {
		log.Fatal("IncrementUpdate[2] does not pass " + err.Error())
	}
	if err := ItemC.Find(bson.M{"newField":"newField"}).One(&result); err !=nil {
		log.Fatal("IncrementUpdate[1] does not pass " + err.Error())
	}
	update = bson.M{"$set": bson.M{"contract_price": 2}}
	if err := ItemC.IncrementUpdateId(result["_id"],update); err != nil {
		log.Fatal("IncrementUpdateId[1] does not pass " + err.Error())
	}

	// IncreUpsert
	//ItemC.IncreUpsert()
	//ItemC.IncreUpsertId()

	fmt.Println("All tests passed")

}
