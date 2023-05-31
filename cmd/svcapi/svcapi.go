package main

import (
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	ratelimit "github.com/asim/go-micro/plugins/wrapper/ratelimiter/uber/v3"
	"github.com/asim/go-micro/plugins/wrapper/select/roundrobin/v3"
	opentracing2 "github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/server"
	"github.com/opentracing/opentracing-go"
	"net"
	"net/http"
	"strconv"
	"tini-paas/api/svcapi/handler"
	"tini-paas/api/svcapi/proto/svcApi"
	microSvcService "tini-paas/internal/svc/proto/svc"
	"tini-paas/pkg/common"
	hystrix2 "tini-paas/plugin/hystrix"
)

var (
	hostIP               = "127.0.0.1"
	serviceHost          = hostIP // 服务地址
	servicePort          = "8084" // 服务端口
	consulHost           = hostIP // 注册中心地址
	consulPort     int64 = 8500   // 注册中心端口
	tracerHost           = hostIP // 链路追踪地址
	tracerPort           = 6831   // 链路追踪端口
	hystrixPort          = 9094   // 熔断端口（每个服务不能重复）
	prometheusPort       = 9194   // 监控端口（每个服务不能重复）
)

func main() {
	// 1、注册中心
	newRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			consulHost + ":" + strconv.FormatInt(consulPort, 10),
		}
	})

	// 2、添加链路追踪
	tracer, closer, err := common.NewTracer("go.micro.api.svcApi", tracerHost+":"+strconv.Itoa(tracerPort))
	if err != nil {
		common.Error(err)
	}
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	// 3、添加熔断器
	streamHandler := hystrix.NewStreamHandler()
	streamHandler.Start()

	// 4、添加日志，将日志采集到日志中心
	// 1) 需要程序日志打入到日志文件中
	// 2) 在程序中添加filebeat.yml 文件
	// 3) 启动filebeat, 启动命令 ./filebeat -e -c filebeat.yml
	fmt.Println("日志统一记录在根目录 micro.log 文件中，请点击查看日志")

	// 5、启动熔断监听程序、
	go func() {
		//http://192.168.0.112:9092/turbine/turbine.stream
		//看板访问地址 http://127.0.0.1:9002/hystrix，url后面一定要带 /hystrix
		err = http.ListenAndServe(net.JoinHostPort("0.0.0.0", strconv.Itoa(hystrixPort)), streamHandler)
		if err != nil {
			common.Error(err)
		}
	}()

	// 6、添加监控
	common.PrometheusBoot(prometheusPort)

	// 7、创建服务
	service := micro.NewService(
		// 自定义服务地址，且必须写在其它参数前面
		micro.Server(server.NewServer(func(options *server.Options) {
			options.Advertise = serviceHost + ":" + servicePort
		})),

		micro.Name("go.micro.api.svcApi"),
		micro.Version("latest"),
		// 指定服务端口
		micro.Address(":"+servicePort),
		// 添加注册中心
		micro.Registry(newRegistry),
		//添加链路追踪
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
		micro.WrapClient(opentracing2.NewClientWrapper(opentracing.GlobalTracer())),
		// 作为客户端范围启动熔断
		micro.WrapClient(hystrix2.NewClientHystrixWrapper()),
		// 添加限流
		micro.WrapHandler(ratelimit.NewHandlerWrapper(1000)),
		// 添加负载均衡
		micro.WrapClient(roundrobin.NewClientWrapper()),
	)

	service.Init()

	// 指定需要访问的服务，可以快速操作已开发的服务，
	// 默认API服务名称带有"Api"，程序会自动替换
	// 如果不带有特定字符会使用默认"XXX" 请自行替换
	svcService := microSvcService.NewSvcService("go.micro.service.svc", service.Client())
	err = svcApi.RegisterSvcApiHandler(service.Server(), &handler.SvcApi{
		SvcService: svcService,
	})
	if err != nil {
		common.Error(err)
	}

	// 启动服务
	err = service.Run()
	if err != nil {
		common.Fatal(err)
	}
}
