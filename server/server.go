package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"example-backend/api"
	"example-backend/store"

	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/protobuf/encoding/protojson"
	"gorm.io/gorm"
)

var logger *zap.Logger
var ServeAddr = "0.0.0.0:8080"
var db *gorm.DB

func injectDB() {
	db = store.GlobalStore.MysqlCli
	if db == nil {
		panic("inject db nil")
	}
}

type Ser struct{}

func (s *Ser) Ping(ctx context.Context, in *api.PingReq) (*api.PingResp, error) {
	return &api.PingResp{IMsg: "pong " + in.Msg}, nil
}

func Serve(cli *cli.Context) error {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGINT)
	httpSer := httpServer()
	<-ch
	httpSer.Close()
	log.Println("serve shutdonw gracefully")
	return nil
}

func addCORS(w http.ResponseWriter) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization")
	w.Header().Add("Access-Control-Allow-Credentials", "true")

}
func PathLogMiddleare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 兼容后缀的斜杠
		r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		addCORS(w)
		next.ServeHTTP(w, r)
	})
}

func httpServer() *http.Server {
	var err error
	logger.Info("服务启动,监听端口8080")
	rmux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames:   true,
			EmitUnpopulated: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		}}))

	err = api.RegisterHITLBackHandlerServer(context.Background(), rmux, &Ser{})
	if err != nil {
		log.Fatal(err)
	}
	// 用 mux 劫持，这样方便用自己的中间件
	router := mux.NewRouter()
	router.Use(PathLogMiddleare)

	// 单独的 http 接口
	router.Path("/api/v1/login").Methods(http.MethodPost).HandlerFunc(HttpLogin)
	router.Path("/api/v1/logout").Methods(http.MethodGet).HandlerFunc(HttpLogout)

	// 路由映射到转换后的 http 服务上
	v1Router := router.PathPrefix("/api/v1").Subrouter()
	v1Router.PathPrefix("/").Handler(rmux)

	server := &http.Server{Addr: ServeAddr, Handler: router}
	go func() {
		if err = server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
	return server
}

func InitLogger() {
	var err error
	cfg := zap.NewProductionConfig()
	cfg.InitialFields = map[string]interface{}{
		"service": "hitl-console-back",
	}
	cfg.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder

	logger, err = cfg.Build(zap.WithCaller(true))
	if err != nil {
		panic(err)
	}
	logger.Info("logger 初始化成功")
	store.InjectLog(logger)
}
