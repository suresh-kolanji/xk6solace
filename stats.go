package ikeaxk6

import (
	"errors"

	"go.k6.io/k6/js/modules"
	"go.k6.io/k6/metrics"
)

type ikeaxk6Metrics struct {
	ReaderDials      *metrics.Metric
	ReaderFetches    *metrics.Metric
	ReaderMessages   *metrics.Metric
	ReaderBytes      *metrics.Metric
	ReaderRebalances *metrics.Metric
	ReaderTimeouts   *metrics.Metric
	ReaderErrors     *metrics.Metric

	ReaderDialTime   *metrics.Metric
	ReaderReadTime   *metrics.Metric
	ReaderWaitTime   *metrics.Metric
	ReaderFetchSize  *metrics.Metric
	ReaderFetchBytes *metrics.Metric

	ReaderOffset        *metrics.Metric
	ReaderLag           *metrics.Metric
	ReaderMinBytes      *metrics.Metric
	ReaderMaxBytes      *metrics.Metric
	ReaderMaxWait       *metrics.Metric
	ReaderQueueLength   *metrics.Metric
	ReaderQueueCapacity *metrics.Metric

	WriterWrites   *metrics.Metric
	WriterMessages *metrics.Metric
	WriterBytes    *metrics.Metric
	WriterErrors   *metrics.Metric

	WriterWriteTime  *metrics.Metric
	WriterWaitTime   *metrics.Metric
	WriterRetries    *metrics.Metric
	WriterBatchSize  *metrics.Metric
	WriterBatchBytes *metrics.Metric

	WriterMaxAttempts  *metrics.Metric
	WriterMaxBatchSize *metrics.Metric
	WriterBatchTimeout *metrics.Metric
	WriterReadTimeout  *metrics.Metric
	WriterWriteTimeout *metrics.Metric
	WriterRequiredAcks *metrics.Metric
	WriterAsync        *metrics.Metric
}

// registerMetrics registers the metrics for the ikeaxk6 module in the metrics registry
// nolint: funlen,maintidx
func registerMetrics(vu modules.VU) (ikeaxk6Metrics, error) {
	var err error
	registry := vu.InitEnv().Registry
	ikeaxk6Metrics := ikeaxk6Metrics{}

	if ikeaxk6Metrics.ReaderDials, err = registry.NewMetric(
		"ikeaxk6.reader.dial.count", metrics.Counter); err != nil {
		return ikeaxk6Metrics, errors.Unwrap(err)
	}

	if ikeaxk6Metrics.ReaderFetches, err = registry.NewMetric(
		"ikeaxk6.reader.fetches.count", metrics.Counter); err != nil {
		return ikeaxk6Metrics, errors.Unwrap(err)
	}

	if ikeaxk6Metrics.ReaderMessages, err = registry.NewMetric(
		"ikeaxk6.reader.message.count", metrics.Counter); err != nil {
		return ikeaxk6Metrics, errors.Unwrap(err)
	}

	if ikeaxk6Metrics.ReaderBytes, err = registry.NewMetric(
		"ikeaxk6.reader.message.bytes", metrics.Counter, metrics.Data); err != nil {
		return ikeaxk6Metrics, errors.Unwrap(err)
	}

	if ikeaxk6Metrics.ReaderRebalances, err = registry.NewMetric(
		"ikeaxk6.reader.rebalance.count", metrics.Counter); err != nil {
		return ikeaxk6Metrics, errors.Unwrap(err)
	}

	if ikeaxk6Metrics.ReaderTimeouts, err = registry.NewMetric(
		"ikeaxk6.reader.timeouts.count", metrics.Counter); err != nil {
		return ikeaxk6Metrics, errors.Unwrap(err)
	}

	if ikeaxk6Metrics.ReaderErrors, err = registry.NewMetric(
		"ikeaxk6.reader.error.count", metrics.Counter); err != nil {
		return ikeaxk6Metrics, errors.Unwrap(err)
	}

	if ikeaxk6Metrics.ReaderDialTime, err = registry.NewMetric(
		"ikeaxk6.reader.dial.seconds", metrics.Trend, metrics.Time); err != nil {
		return ikeaxk6Metrics, errors.Unwrap(err)
	}

	if ikeaxk6Metrics.ReaderReadTime, err = registry.NewMetric(
		"ikeaxk6.reader.read.seconds", metrics.Trend, metrics.Time); err != nil {
		return ikeaxk6Metrics, errors.Unwrap(err)
	}

	if ikeaxk6Metrics.ReaderWaitTime, err = registry.NewMetric(
		"ikeaxk6.reader.wait.seconds", metrics.Trend, metrics.Time); err != nil {
		return ikeaxk6Metrics, errors.Unwrap(err)
	}

	if ikeaxk6Metrics.ReaderFetchSize, err = registry.NewMetric(
		"ikeaxk6.reader.fetch.size", metrics.Counter); err != nil {
		return ikeaxk6Metrics, errors.Unwrap(err)
	}

	if ikeaxk6Metrics.ReaderFetchBytes, err = registry.NewMetric(
		"ikeaxk6.reader.fetch.bytes", metrics.Counter, metrics.Data); err != nil {
		return ikeaxk6Metrics, errors.Unwrap(err)
	}

	if ikeaxk6Metrics.ReaderOffset, err = registry.NewMetric(
		"ikeaxk6.reader.offset", metrics.Gauge); err != nil {
		return ikeaxk6Metrics, errors.Unwrap(err)
	}

	if ikeaxk6Metrics.ReaderLag, err = registry.NewMetric(
		"ikeaxk6.reader.lag", metrics.Gauge); err != nil {
		return ikeaxk6Metrics, errors.Unwrap(err)
	}

	if ikeaxk6Metrics.ReaderMinBytes, err = registry.NewMetric(
		"ikeaxk6.reader.fetch_bytes.min", metrics.Gauge); err != nil {
		return ikeaxk6Metrics, errors.Unwrap(err)
	}

	if ikeaxk6Metrics.ReaderMaxBytes, err = registry.NewMetric(
		"ikeaxk6.reader.fetch_bytes.max", metrics.Gauge); err != nil {
		return ikeaxk6Metrics, errors.Unwrap(err)
	}

	if ikeaxk6Metrics.ReaderMaxWait, err = registry.NewMetric(
		"ikeaxk6.reader.fetch_wait.max", metrics.Gauge, metrics.Time); err != nil {
		return ikeaxk6Metrics, errors.Unwrap(err)
	}

	if ikeaxk6Metrics.ReaderQueueLength, err = registry.NewMetric(
		"ikeaxk6.reader.queue.length", metrics.Gauge); err != nil {
		return ikeaxk6Metrics, errors.Unwrap(err)
	}

	if ikeaxk6Metrics.ReaderQueueCapacity, err = registry.NewMetric(
		"ikeaxk6.reader.queue.capacity", metrics.Gauge); err != nil {
		return ikeaxk6Metrics, errors.Unwrap(err)
	}

	if ikeaxk6Metrics.WriterWrites, err = registry.NewMetric(
		"ikeaxk6.writer.write.count", metrics.Counter); err != nil {
		return ikeaxk6Metrics, errors.Unwrap(err)
	}

	if ikeaxk6Metrics.WriterMessages, err = registry.NewMetric(
		"ikeaxk6.writer.message.count", metrics.Counter); err != nil {
		return ikeaxk6Metrics, errors.Unwrap(err)
	}

	if ikeaxk6Metrics.WriterBytes, err = registry.NewMetric(
		"ikeaxk6.writer.message.bytes", metrics.Counter, metrics.Data); err != nil {
		return ikeaxk6Metrics, errors.Unwrap(err)
	}

	if ikeaxk6Metrics.WriterErrors, err = registry.NewMetric(
		"ikeaxk6.writer.error.count", metrics.Counter); err != nil {
		return ikeaxk6Metrics, errors.Unwrap(err)
	}

	if ikeaxk6Metrics.WriterWriteTime, err = registry.NewMetric(
		"ikeaxk6.writer.write.seconds", metrics.Trend, metrics.Time); err != nil {
		return ikeaxk6Metrics, errors.Unwrap(err)
	}

	if ikeaxk6Metrics.WriterWaitTime, err = registry.NewMetric(
		"ikeaxk6.writer.wait.seconds", metrics.Trend, metrics.Time); err != nil {
		return ikeaxk6Metrics, errors.Unwrap(err)
	}

	if ikeaxk6Metrics.WriterRetries, err = registry.NewMetric(
		"ikeaxk6.writer.retries.count", metrics.Counter); err != nil {
		return ikeaxk6Metrics, errors.Unwrap(err)
	}

	if ikeaxk6Metrics.WriterBatchSize, err = registry.NewMetric(
		"ikeaxk6.writer.batch.size", metrics.Counter); err != nil {
		return ikeaxk6Metrics, errors.Unwrap(err)
	}

	if ikeaxk6Metrics.WriterBatchBytes, err = registry.NewMetric(
		"ikeaxk6.writer.batch.bytes", metrics.Counter, metrics.Data); err != nil {
		return ikeaxk6Metrics, errors.Unwrap(err)
	}

	if ikeaxk6Metrics.WriterMaxAttempts, err = registry.NewMetric(
		"ikeaxk6.writer.attempts.max", metrics.Gauge); err != nil {
		return ikeaxk6Metrics, errors.Unwrap(err)
	}

	if ikeaxk6Metrics.WriterMaxBatchSize, err = registry.NewMetric(
		"ikeaxk6.writer.batch.max", metrics.Gauge); err != nil {
		return ikeaxk6Metrics, errors.Unwrap(err)
	}

	if ikeaxk6Metrics.WriterBatchTimeout, err = registry.NewMetric(
		"ikeaxk6.writer.batch.timeout", metrics.Gauge, metrics.Time); err != nil {
		return ikeaxk6Metrics, errors.Unwrap(err)
	}

	if ikeaxk6Metrics.WriterReadTimeout, err = registry.NewMetric(
		"ikeaxk6.writer.read.timeout", metrics.Gauge, metrics.Time); err != nil {
		return ikeaxk6Metrics, errors.Unwrap(err)
	}

	if ikeaxk6Metrics.WriterWriteTimeout, err = registry.NewMetric(
		"ikeaxk6.writer.write.timeout", metrics.Gauge, metrics.Time); err != nil {
		return ikeaxk6Metrics, errors.Unwrap(err)
	}

	if ikeaxk6Metrics.WriterRequiredAcks, err = registry.NewMetric(
		"ikeaxk6.writer.acks.required", metrics.Gauge); err != nil {
		return ikeaxk6Metrics, errors.Unwrap(err)
	}

	if ikeaxk6Metrics.WriterAsync, err = registry.NewMetric(
		"ikeaxk6.writer.async", metrics.Rate); err != nil {
		return ikeaxk6Metrics, errors.Unwrap(err)
	}

	return ikeaxk6Metrics, nil
}
