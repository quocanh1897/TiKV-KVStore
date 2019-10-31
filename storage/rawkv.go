package storage

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/tikv/client-go/config"
	"github.com/tikv/client-go/rawkv"
)

type Storage interface {
	ConnectTiKV(ctx context.Context) error
	DisConnectTiKV() error
	GetClient() interface{}
}

type RawTiKV struct {
	host []string
	client *rawkv.Client
}

func NewRawKV(ctx context.Context, host []string)(Storage, error){
	rawTiKV := &RawTiKV{
		host:   host,
	}

	if err := rawTiKV.ConnectTiKV(ctx); err != nil {
		return nil, err
	}

	return rawTiKV, nil
}

func (rawtikv *RawTiKV) ConnectTiKV(ctx context.Context) error {

	logrus.Info("[Storage] connecting with TiKV ... ", rawtikv.host)

	cli, err := rawkv.NewClient(ctx, rawtikv.host, config.Default())
	if err != nil {
		logrus.Error("[Storage] failed to connect with TiKV ", err)
		return err
	}

	logrus.Info("[Storage] connected with TiKV ", rawtikv.host)

	rawtikv.client = cli
	return nil
}

func (rawtikv *RawTiKV) DisConnectTiKV() error {
	if err := rawtikv.client.Close(); err != nil {
		logrus.Error("[Storage] DisConnect with TiKV error: ", err)
		return err
	}

	logrus.Info("[Storage] DisConnected with success")

	return nil
}

func (rawtikv *RawTiKV) GetClient() interface{} {
	return rawtikv.client
}
