package tools

import (
	"due-game-example/shared/pb/pb_common"
	"due-game-example/util/logger"
	"fmt"
	"github.com/dobyte/due/v2/cluster/node"
	"github.com/golang/protobuf/proto"
	"reflect"
	"strings"
)

// IsNil 判断接口是否为空
func IsNil(v interface{}) bool {
	if v == nil {
		return true
	}
	return reflect.ValueOf(v).IsNil()
}

func Wrap[T, S proto.Message](fn func(ctx node.Context, in T) (out S)) func(ctx node.Context) {
	return func(ctx node.Context) {
		reqTyp := reflect.TypeOf((*T)(nil)).Elem().Elem()
		ackTyp := reflect.TypeOf((*S)(nil)).Elem().Elem()

		req := reflect.New(reqTyp).Interface().(T)
		ack := reflect.New(ackTyp).Interface().(S)

		if err := ctx.Parse(req); err != nil {
			logger.Errorf("parse message failed: %v", err)
			_ = ctx.Disconnect(true)
			return
		}

		defer func() {
			name := proto.MessageReflect(ack).Descriptor().Name()
			pbName := strings.ToUpper(string(name))
			fmt.Println(name, pbName, pb_common.ROUTE_value[pbName])
			if !IsNil(ack) {
				if err := ctx.Response(ack); err != nil {
					logger.Errorf("response message failed: %v", err)
				}
			}
		}()

		ack = fn(ctx, req)
	}
}
