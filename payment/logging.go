package payment

import (
	"github.com/go-kit/kit/log"
	"time"
//	"github.com/go-kit/kit/log/level"
	log1 "github.com/sirupsen/logrus"
	//"os"
)
// LoggingMiddleware logs method calls, parameters, results, and elapsed time.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next Service) Service {
		return loggingMiddleware{
			next:   next,
			logger: logger,
		}
	}
}

type loggingMiddleware struct {
	next   Service
	logger log.Logger
}

func (mw loggingMiddleware) Authorise(amount float32, traceID string,spanID string) (auth Authorisation, err error) {
	log1.SetFormatter(&log1.JSONFormatter{})
	defer func(begin time.Time) {
		if (err == nil) {
/*			level.Info(mw.logger).Log(
				"method", "Authorise",
				"result", auth.Authorised,
				"took", time.Since(begin),
				"traceID", traceID,
			)*/
			standardFields :=  log1.Fields{
				"method": "Authorise",
				"result": auth.Authorised,
				"took": time.Since(begin),
				"traceID": traceID,
				"spanID": spanID,

			}
		log1.WithFields(standardFields).Info("Authorise Payment")
		} else {
/*			level.Error(mw.logger).Log(
				"method", "Authorise",
				"result", auth.Authorised,
				"took", time.Since(begin),
				"traceID", traceID,
				"error", err,
			)*/
                        standardFields :=  log1.Fields{
				"method": "Authorise",
				"result": auth.Authorised,
				"took": time.Since(begin),
				"traceID": traceID,
				"error": err,
				"spanID": spanID,
			}
                log1.WithFields(standardFields).Error("Authorise Payment")

		}
	}(time.Now())
	return mw.next.Authorise(amount, traceID,spanID)
}

func (mw loggingMiddleware) Health() (health []Health) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "Health",
			"result", len(health),
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.next.Health()
}
