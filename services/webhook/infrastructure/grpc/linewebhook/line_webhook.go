package linewebhook

import (
	"context"
	"os"
	"webhook/app/interactor"
	"webhook/infrastructure/grpc/linewebhook/pb"
	"webhook/infrastructure/spanner"
	"webhook/interface/controller"
	"webhook/interface/repository"
)

var (
	project  = os.Getenv("GCP_PROJECT_ID")
	instance = os.Getenv("DB_INSTANCE_NAME")
	dbName   = os.Getenv("DB_NAME")
)

type LineWebhook struct {
	ctlr    *controller.LineMessageController
	dbClose func()
}

func New() pb.LineWebhookServer {
	sql := spanner.NewSql(project, instance, dbName)
	return &LineWebhook{
		ctlr: controller.NewLineMessageController(
			interactor.NewNewsInteractor(
				repository.NewUserRepository(sql),
			),
		),
		dbClose: sql.Close,
	}
}

func (w *LineWebhook) Health(ctx context.Context, in *pb.HealthRequest) (*pb.HealthReply, error) {
	defer w.dbClose()
	// sql := spanner.NewSql(project, instance, dbName)
	// client, err := sql.NewClient(ctx)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// result := client.GetAll(ctx, "User", []string{"Id", "LineBotChannelId", "LineUID", "CreatedAt"})
	// var id string
	// var channelId string
	// var uid string
	// var createdat time.Time
	// items := make([]interface{})
	// err = result.Loop(func(row repository.Row) error {
	//   items = append(items, map[string]interface{}{"id"})
	// 	return row.Bind(&id, &channelId, &uid, &createdat)
	// })
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// log.Println(id, channelId, uid, createdat)
	return &pb.HealthReply{Data: "ok"}, nil
}

func (w *LineWebhook) Message(ctx context.Context, in *pb.MessageRequest) (*pb.MessageReply, error) {
	defer w.dbClose()
	return w.ctlr.Message(ctx, in)
}
