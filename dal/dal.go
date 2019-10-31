package dal

import (
	configs "TiClientExample/config"
	"TiClientExample/storage"
	"context"
	"github.com/tikv/client-go/rawkv"
)

type DAO interface {
	Get(ctx context.Context, key []byte) ([]byte, error)
	Put(ctx context.Context, key []byte, value []byte) error
	Delete(ctx context.Context, key []byte) error
	DisconnectStorage() error
}

type Impl struct {
	rawKV storage.Storage
}

func NewDAO(ctx context.Context, cfg *configs.Config) (DAO, error) {
	newrawkv, err := storage.NewRawKV(ctx, cfg.PDConfig)
	if err != nil {
		return nil, err
	}

	return &Impl{
		rawKV: newrawkv,
	}, nil
}

func (dao Impl) Get(ctx context.Context, key []byte) ([]byte, error) {
	cli := dao.rawKV.GetClient().(*rawkv.Client)

	return cli.Get(ctx, key)
}
func (dao Impl) Put(ctx context.Context, key []byte, value []byte) error {
	cli := dao.rawKV.GetClient().(*rawkv.Client)

	return cli.Put(ctx, key, value)
}
func (dao Impl) Delete(ctx context.Context, key []byte) error {
	cli := dao.rawKV.GetClient().(*rawkv.Client)
	return cli.Delete(ctx, key)
}
func (dao Impl) DisconnectStorage() error {
	return dao.rawKV.DisConnectTiKV()
}
