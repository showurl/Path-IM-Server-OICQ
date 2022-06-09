package logic

import (
	"github.com/Path-IM/Path-IM-Server/common/xtrace"
	chatpb "github.com/showurl/Path-IM-Server-OICQ/app/msg/rpc/pb"
)

func (l *SendMsgLogic) sendMsgToKafka(m *chatpb.MsgDataToMQ, key string) error {
	m.TraceId = xtrace.TraceIdFromContext(l.ctx)
	pid, offset, err := l.svcCtx.Producer.SendMessage(l.ctx, m, key)
	if err != nil {
		l.Logger.Error(m.TraceId, " kafka send failed ", "send data ", m.String(), "pid ", pid, "offset ", offset, "err ", err.Error(), "key ", key)
	}
	return err
}
