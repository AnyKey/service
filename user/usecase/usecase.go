package usecase

import (
	"context"
	"encoding/json"
	"github.com/AnyKey/service/user"
	log "github.com/sirupsen/logrus"
)

type userUseCase struct {
	userRepo         user.Repository
	userHttpDelivery user.HttpDelivery
}

func New(ur user.Repository, uh user.HttpDelivery) user.Usecase {
	return &userUseCase{
		userRepo:         ur,
		userHttpDelivery: uh,
	}
}

func (uuc *userUseCase) SetToken(token string) error {
	ctx := context.Background()
	bytes, err := json.Marshal(token)
	if err != nil {
		log.Errorln("[SetToken] Error: ", err)
		return err
	}
	uuc.userRepo.SetToken(ctx, bytes)
	return nil
}

func (uuc *userUseCase) GetList() ([]user.TitleList, error) {
	ctx := context.Background()
	token := uuc.userRepo.GetToken(ctx)
	res, err := uuc.userHttpDelivery.GetSubscriptions(*token)
	if err != nil {
		log.Errorln("[GetList] Error: ", err)
		return nil, err
	}
	list := make([]user.TitleList, 0)

	for i := range res.Items {
		list = append(list, res.Items[i].Snippet)
	}

	log.Infoln("sub list received")
	return list, nil
}
