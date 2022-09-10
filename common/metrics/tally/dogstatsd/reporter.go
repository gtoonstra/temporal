package metrics

import (
    "fmt"
    "sort"
    nflxTag "temporal/internal/tag"
    "time"

    "github.com/DataDog/datadog-go/statsd"
    "github.com/uber-go/tally"
    "go.temporal.io/server/common/log"
    "go.temporal.io/server/common/log/tag"
)

const defaultFlushBytes = 1432
const extensionName = "dogstatsd"

// DogstatsdReporterConfig contains the config items for the dogstatsd
// metrics reporter.
type DogstatsdReporterConfig struct {
    // The host and port of the statsd server
    HostPort string
    // FlusherInterval is the maximum interval for sending packets.
    // If it is not specified, it defaults to 1 second.
    FlushInterval time.Duration
    // FlushBytes specifies the maximum UDP packet size you wish to send.
    // If FlushBytes is unspecified, it defaults to 1432 bytes, which is
    // considered safe for local traffic.
    FlushBytes int
}

type dogstatsdReporter struct {
    dogstatsd *statsd.Client
    log       log.Logger
}

// NewReporter is a wrapper on top of "github.com/DataDog/datadog-go/statsd"
// The purpose is to support datadog-formatted statsd metric tagging.
func NewReporter(config *DogstatsdReporterConfig, logger log.Logger) tally.StatsReporter {
    hostPort := config.HostPort
    flushInterval := config.FlushInterval
    flushBytes := config.FlushBytes

    if hostPort == "" {
        hostPort = "127.0.0.1:8125"
    }
    if flushInterval == 0 {
        flushInterval = time.Second
    }
    if flushBytes == 0 {
        flushBytes = defaultFlushBytes
    }

    client, err := statsd.New(hostPort, statsd.WithBufferFlushInterval(flushInterval), statsd.WithMaxBytesPerPayload(flushBytes))
    if err != nil {
        logger.Fatal("error creating dogstatsd client", nflxTag.Extension(extensionName), tag.Error(err))
    }
    return dogstatsdReporter{
        dogstatsd: client,
        log:       logger,
    }
}

func (r dogstatsdReporter) Capabilities() tally.Capabilities {
    return r
}

func (r dogstatsdReporter) Reporting() bool {
    return true
}

func (r dogstatsdReporter) Tagging() bool {
    return true
}

func (r dogstatsdReporter) Flush() {
    if err := r.dogstatsd.Flush(); err != nil {
        r.log.Error("error while flushing", nflxTag.Extension(extensionName), tag.Error(err))
    }
}

func (r dogstatsdReporter) ReportCounter(name string, tags map[string]string, value int64) {
    name = r.sanitizeMetricName(name)
    if err := r.dogstatsd.Count(name, value, r.marshalTags(tags), 1); err != nil {
        r.log.Error("failed reporting counter", nflxTag.Extension(extensionName), tag.Error(err))
    }
}

func (r dogstatsdReporter) ReportGauge(name string, tags map[string]string, value float64) {
    name = r.sanitizeMetricName(name)
    if err := r.dogstatsd.Gauge(name, value, r.marshalTags(tags), 1); err != nil {
        r.log.Error("failed reporting gauge", nflxTag.Extension(extensionName), tag.Error(err))
    }
}

func (r dogstatsdReporter) ReportTimer(name string, tags map[string]string, interval time.Duration) {
    name = r.sanitizeMetricName(name)
    if err := r.dogstatsd.Timing(name, interval, r.marshalTags(tags), 1); err != nil {
        r.log.Error("failed reporting timer", nflxTag.Extension(extensionName), tag.Error(err))
    }
}

func (r dogstatsdReporter) ReportHistogramValueSamples(name string, tags map[string]string, buckets tally.Buckets, bucketLowerBound, bucketUpperBound float64, samples int64) {
    // TODO(rz): Temporal does not currently use histograms
    r.log.Warn("unexpected call to ReportHistogramValueSamples", nflxTag.Extension(extensionName))
}

func (r dogstatsdReporter) ReportHistogramDurationSamples(name string, tags map[string]string, buckets tally.Buckets, bucketLowerBound, bucketUpperBound time.Duration, samples int64) {
    // TODO(rz): Temporal does not currently use histograms
    r.log.Warn("unexpected call to ReportHistogramDurationSamples", nflxTag.Extension(extensionName))
}

func (r dogstatsdReporter) marshalTags(tags map[string]string) []string {
    var keys []string
    for k := range tags {
        keys = append(keys, k)
    }
    sort.Strings(keys)

    var dogTags []string
    for _, tk := range keys {
        dogTags = append(dogTags, fmt.Sprintf("%s:%s", tk, tags[tk]))
    }
    return dogTags
}

func (r dogstatsdReporter) sanitizeMetricName(name string) string {
    return fmt.Sprintf("temporal.%s", name)
}
