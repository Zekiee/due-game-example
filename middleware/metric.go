package middleware

import (
	"due-game-example/util/logger"
	"github.com/dobyte/due/v2/cluster/node"
	"time"
)

// RecordMetric1 指标
func RecordMetric1(middleware *node.Middleware, ctx node.Context) {
	logger.Info("RecordMetric1 in")
	start := time.Now()
	middleware.Next(ctx)
	logger.Info("RecordMetric1 out", time.Since(start))
}

// RecordMetric2 指标
func RecordMetric2(middleware *node.Middleware, ctx node.Context) {
	logger.Info("RecordMetric2 in")
	start := time.Now()
	middleware.Next(ctx)
	logger.Info("RecordMetric2 out", time.Since(start))
}
