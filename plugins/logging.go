package plugins

import (
	"github.com/go-kit/kit/log"
	"go_consul/service"
	"time"
)

// loggingMiddleware Make a new type
// that contains Service interface and logger instance

//装饰器模式，对service进行装饰 简单说，就是结构体内充填东西  service内最后一行代码定义
//该logging等价于service，

type StringServiceWithLogging struct {
	service.Service
	logger log.Logger
}

// LoggingMiddleware make logging middleware
func LoggingMiddleware(logger log.Logger) service.ServiceMiddleware {
	return func(next service.Service) service.Service {
		return StringServiceWithLogging{next, logger}
	}
}

func (mw StringServiceWithLogging) Concat(a, b string) (ret string, err error) {
	// 函数执行结束后打印日志
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "Concat",
			"a", a,
			"b", b,
			"result", ret,
			"took", time.Since(begin),
		)
	}(time.Now()) //调用匿名函数

	ret, err = mw.Service.Concat(a, b)
	return ret, err
}

func (mw StringServiceWithLogging) Diff(a, b string) (ret string, err error) {
	// 函数执行结束后打印日志
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "Diff",
			"a", a,
			"b", b,
			"result", ret,
			"took", time.Since(begin),
		)
	}(time.Now()) //传参数(begin time.Time)

	ret, err = mw.Service.Diff(a, b)
	return ret, err
}

func (mw StringServiceWithLogging) HealthCheck() (result bool) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "HealthChcek",
			"result", result,
			"took", time.Since(begin),
		)
	}(time.Now())
	result = mw.Service.HealthCheck()
	return
}
