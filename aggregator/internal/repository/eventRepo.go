package repository

import (
	"aggregator/cmd/config"
	"aggregator/internal/db"
	"aggregator/internal/logg"
	"aggregator/internal/model"
	"context"
	"strconv"
	"sync"

	"github.com/go-redis/redis/v8"
)

type EventRepo struct {
	DataBase db.DataBase[*redis.Client]
	ctx      context.Context
	store    map[string]*model.AgrTypeValue
	cfg      config.Config
	logger   logg.Logger
	mx       sync.RWMutex
}

func NewEventRepo(
	ctx context.Context,
	config config.Config,
	logger logg.Logger,
	database db.DataBase[*redis.Client]) *EventRepo {

	return &EventRepo{
		DataBase: database,
		ctx:      ctx,
		cfg:      config,
		logger:   logger,
		store:    make(map[string]*model.AgrTypeValue, 10),
		mx:       sync.RWMutex{},
	}
}

func (r *EventRepo) Set(dataArr []*model.AgrTypeValue) {
	for _, v := range dataArr {
		err := r.DataBase.GetDB().Set(r.ctx, v.Type, v.Value, 0).Err()

		if err != nil {
			r.logger.RaiseFatal(err.Error())
		}
	}
}

func (r *EventRepo) Read(key string) int64 {
	valStr, err := r.DataBase.GetDB().Get(r.ctx, key).Result()
	if err != nil {
		r.logger.RaiseFatal(err.Error())
	}

	valInt, err := strconv.Atoi(valStr)
	if err != nil {
		r.logger.RaiseFatal(err.Error())
	}
	return int64(valInt)
}

// func (r *EventRepo) Set(dataArr []*model.AgrTypeValue) {
// 	r.mx.Lock()
// 	defer r.mx.Unlock()

// 	for _, v := range dataArr {
// 		if modelArg, ok := r.store[v.Type]; ok {
// 			modelArg.Value += v.Value
// 			r.store[v.Type] = modelArg

// 		} else {
// 			r.store[v.Type] = v
// 		}
// 	}
// }

// func (r *EventRepo) Set(dataArr []*model.AgrTypeValue) {
// 	queryValues := make([]string, len(dataArr))
// 	params := make([]any, len(dataArr)*2)

// 	for i, upd := range dataArr {
// 		idx := i * 2
// 		queryValues[i] = fmt.Sprintf("(%d, %d)", idx+1, idx+2)
// 		params[idx], params[idx+1] = upd.Type, upd.Value
// 	}

// 	updateQuery := `
// 					BEGIN TRANSACTION;
// 					WITH tmpTb AS (
// 						SELECT *
// 						FROM (VALUES
// 							%s
// 						) AS t(type, value)
// 					)
// 					UPDATE args p
// 					SET value = p.value + u.value
// 					FROM tmpTb u
// 					WHERE p.type = u.type;
// 					COMMIT;
// 	`
// 	finalQuery := fmt.Sprintf(updateQuery, strings.Join(queryValues, ","))

// 	_, err := r.DataBase.GetDB().ExecContext(r.ctx, finalQuery, params...)
// 	if err != nil {
// 		r.logger.RaiseFatal(err.Error())
// 	}
// }
