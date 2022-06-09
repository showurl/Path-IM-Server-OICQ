package repository

import (
	"github.com/gocql/gocql"
	"github.com/showurl/Path-IM-Server-OICQ/app/msg/rpc/internal/svc"
	"github.com/showurl/Path-IM-Server-OICQ/app/msg/rpc/pb"
	"go.mongodb.org/mongo-driver/mongo"
)

type IPullHistoryMsg interface {
	GetMsgBySeqList(uid string, seqList []uint32) (seqMsg []*pb.MsgData, err error)
	GetMsgByGroupSeqList(groupId string, seqList []uint32) (seqMsg []*pb.MsgData, err error)
}
type MongoHistory struct {
	svcCtx      *svc.ServiceContext
	MongoClient *mongo.Client
}

type CassandraHistory struct {
	svcCtx          *svc.ServiceContext
	CassandraClient *gocql.Session
}
