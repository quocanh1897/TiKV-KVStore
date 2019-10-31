package main

import (
	configs "TiClientExample/config"
	"TiClientExample/dal"
	"context"
	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()

	testSetKeyValue(ctx)
	testGetKeyExist(ctx)
	testGetKeyNotExist(ctx)
	testDelKey(ctx)
}

func testSetKeyValue(ctx context.Context) {
	logrus.Info("=========[TEST SET KV]=========")
	cfg := loadConfig()
	dal, er := dal.NewDAO(ctx, cfg)
	if er != nil {
		logrus.Error("[TEST DAO FAIL]: ", er.Error())
	}

	key := []byte("test-key")
	val := []byte("val-corr")

	err := dal.Put(ctx, key, val)
	if err != nil {
		logrus.Fatal("[TEST SET KV]: ", err.Error())
	}

	logrus.Info("[TEST SET KV] ok!")

	err = dal.DisconnectStorage()
	if err != nil {
		logrus.Fatal("[TEST SET KV]: ", err.Error())
	}
}

func testGetKeyExist(ctx context.Context){
	logrus.Info("=========[TEST GET KV EXIST]=========")

	cfg := loadConfig()
	dal, er := dal.NewDAO(ctx, cfg)
	if er != nil {
		logrus.Error("[TEST DAO FAIL]: ", er.Error())
	}

	key := []byte("test-key")
	val := []byte("val-corr")

	err := dal.Put(ctx,key,val)
	if err != nil {
		logrus.Fatal("[TEST GET KEY EXIST]: ", err.Error())
	}

	val, err = dal.Get(ctx, key)
	if err != nil {
		logrus.Fatalf("[testGetKeyExist]" + err.Error())
	}

	logrus.Infof("[TEST GET KEY EXIST] key: %s for val: %s", key, val)
	logrus.Info("[TEST GET KEY EXIST] ok!")

	err = dal.DisconnectStorage()
	if err != nil {
		logrus.Fatalf("[TEST GET KEY EXIST]" + err.Error())
	}
}

func testGetKeyNotExist(ctx context.Context){
	logrus.Info("=========[TEST GET KV NOT EXIST]=========")

	key := []byte("test-key-not-exist")
	cfg := loadConfig()
	dal, er := dal.NewDAO(ctx, cfg)
	if er != nil {
		logrus.Error("[TEST DAO FAIL]: ", er.Error())
	}

	_, err := dal.Get(ctx, key)
	if err != nil {
		logrus.Fatal("[TEST GET KEY NOT EXIST]" + err.Error())
	}

	logrus.Info("[TEST GET KEY NOT EXIST] ok!")

	err = dal.DisconnectStorage()
	if err != nil {
		logrus.Fatal("[TEST GET KEY NOT EXIST]" + err.Error())
	}
}

func testDelKey(ctx context.Context){
	logrus.Info("=========[TEST DELETE KV]=========")

	key := []byte("test-del-key")
	val := []byte("val-of-test-del-key")
	cfg := loadConfig()
	dal, er := dal.NewDAO(ctx, cfg)
	if er != nil {
		logrus.Error("[TEST DAO FAIL]: ", er.Error())
	}

	//put
	err := dal.Put(ctx, key, val)
	if err != nil {
		logrus.Fatalf("[TEST DEL KEY]" + err.Error())
	}

	//get
	val, err = dal.Get(ctx, key)
	if err != nil {
		logrus.Fatalf("[TEST DEL KEY]" + err.Error())
	}

	logrus.Infof("[TEST DEL KEY] key: %s for val: %s", key, val)

	//delete
	err = dal.Delete(ctx, key)
	if err != nil {
		logrus.Fatalf("[TEST DEL KEY]" + err.Error())
	}

	//get
	val, err = dal.Get(ctx, key)
	if err != nil {
		logrus.Fatalf("[TEST DEL KEY]" + err.Error())
	}

	logrus.Infof("[TEST DEL KEY] key: %s for val: %s", key, val)

	logrus.Info("[TEST DEL KEY] ok!")

	err = dal.DisconnectStorage()
	if err != nil {
		logrus.Fatalf("[TEST DEL KEY]" + err.Error())
	}
}

// loadConfig use for setup test case
func loadConfig() *configs.Config{
	return configs.GetConfig()
}
