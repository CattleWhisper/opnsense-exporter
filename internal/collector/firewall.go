package collector

import (
	"log/slog"
	"time"

	"github.com/AthennaMind/opnsense-exporter/opnsense"
	"github.com/prometheus/client_golang/prometheus"
)

type firewallCollector struct {
	log                 *slog.Logger
	inIPv4PassPackets   *prometheus.Desc
	outIPv4PassPackets  *prometheus.Desc
	inIPv4BlockPackets  *prometheus.Desc
	outIPv4BlockPackets *prometheus.Desc

	inIPv6PassPackets   *prometheus.Desc
	outIPv6PassPackets  *prometheus.Desc
	inIPv6BlockPackets  *prometheus.Desc
	outIPv6BlockPackets *prometheus.Desc

	inIPv4PassBytes   *prometheus.Desc
	outIPv4PassBytes  *prometheus.Desc
	inIPv4BlockBytes  *prometheus.Desc
	outIPv4BlockBytes *prometheus.Desc

	inIPv6PassBytes   *prometheus.Desc
	outIPv6PassBytes  *prometheus.Desc
	inIPv6BlockBytes  *prometheus.Desc
	outIPv6BlockBytes *prometheus.Desc

	interfaceReferences    *prometheus.Desc
	interfaceCountersReset *prometheus.Desc

	stateTableEntries        *prometheus.Desc
	stateTableOperations     *prometheus.Desc
	sourceTrackingEntries    *prometheus.Desc
	sourceTrackingOperations *prometheus.Desc
	pfCounters               *prometheus.Desc
	pfLimitCounters          *prometheus.Desc

	memoryLimit    *prometheus.Desc
	timeoutSeconds *prometheus.Desc
	adaptiveStates *prometheus.Desc

	subsystem string
	instance  string
}

func init() {
	collectorInstances = append(collectorInstances, &firewallCollector{
		subsystem: FirewallSubsystem,
	})
}

func (c *firewallCollector) Name() string {
	return c.subsystem
}

func (c *firewallCollector) Register(namespace, instanceLabel string, log *slog.Logger) {
	c.log = log
	c.instance = instanceLabel
	c.log.Debug("Registering collector", "collector", c.Name())

	c.inIPv4PassPackets = buildPrometheusDesc(c.subsystem, "in_ipv4_pass_packets",
		"The number of IPv4 incoming packets that were allowed to pass through the firewall by interface",
		[]string{"interface"},
	)

	c.outIPv4PassPackets = buildPrometheusDesc(c.subsystem, "out_ipv4_pass_packets",
		"The number of IPv4 outgoing packets that were allowed to pass through the firewall by interface",
		[]string{"interface"},
	)

	c.inIPv4BlockPackets = buildPrometheusDesc(c.subsystem, "in_ipv4_block_packets",
		"The number of IPv4 incoming packets that were blocked by the firewall by interface",
		[]string{"interface"},
	)

	c.outIPv4BlockPackets = buildPrometheusDesc(c.subsystem, "out_ipv4_block_packets",
		"The number of IPv4 outgoing packets that were blocked by the firewall by interface",
		[]string{"interface"},
	)

	c.inIPv6PassPackets = buildPrometheusDesc(c.subsystem, "in_ipv6_pass_packets",
		"The number of IPv6 incoming packets that were allowed to pass through the firewall by interface",
		[]string{"interface"},
	)

	c.outIPv6PassPackets = buildPrometheusDesc(c.subsystem, "out_ipv6_pass_packets",
		"The number of IPv6 outgoing packets that were allowed to pass through the firewall by interface",
		[]string{"interface"},
	)

	c.inIPv6BlockPackets = buildPrometheusDesc(c.subsystem, "in_ipv6_block_packets",
		"The number of IPv6 incoming packets that were blocked by the firewall by interface",
		[]string{"interface"},
	)

	c.outIPv6BlockPackets = buildPrometheusDesc(c.subsystem, "out_ipv6_block_packets",
		"The number of IPv6 outgoing packets that were blocked by the firewall by interface",
		[]string{"interface"},
	)

	c.inIPv4PassBytes = buildPrometheusDesc(c.subsystem, "in_ipv4_pass_bytes_total",
		"The number of IPv4 incoming bytes that were allowed to pass through the firewall by interface",
		[]string{"interface"},
	)

	c.outIPv4PassBytes = buildPrometheusDesc(c.subsystem, "out_ipv4_pass_bytes_total",
		"The number of IPv4 outgoing bytes that were allowed to pass through the firewall by interface",
		[]string{"interface"},
	)

	c.inIPv4BlockBytes = buildPrometheusDesc(c.subsystem, "in_ipv4_block_bytes_total",
		"The number of IPv4 incoming bytes that were blocked by the firewall by interface",
		[]string{"interface"},
	)

	c.outIPv4BlockBytes = buildPrometheusDesc(c.subsystem, "out_ipv4_block_bytes_total",
		"The number of IPv4 outgoing bytes that were blocked by the firewall by interface",
		[]string{"interface"},
	)

	c.inIPv6PassBytes = buildPrometheusDesc(c.subsystem, "in_ipv6_pass_bytes_total",
		"The number of IPv6 incoming bytes that were allowed to pass through the firewall by interface",
		[]string{"interface"},
	)

	c.outIPv6PassBytes = buildPrometheusDesc(c.subsystem, "out_ipv6_pass_bytes_total",
		"The number of IPv6 outgoing bytes that were allowed to pass through the firewall by interface",
		[]string{"interface"},
	)

	c.inIPv6BlockBytes = buildPrometheusDesc(c.subsystem, "in_ipv6_block_bytes_total",
		"The number of IPv6 incoming bytes that were blocked by the firewall by interface",
		[]string{"interface"},
	)

	c.outIPv6BlockBytes = buildPrometheusDesc(c.subsystem, "out_ipv6_block_bytes_total",
		"The number of IPv6 outgoing bytes that were blocked by the firewall by interface",
		[]string{"interface"},
	)

	c.interfaceReferences = buildPrometheusDesc(c.subsystem, "interface_references",
		"The number of pf rules and tables referencing the interface",
		[]string{"interface"},
	)

	c.interfaceCountersReset = buildPrometheusDesc(c.subsystem, "interface_counters_cleared_timestamp_seconds",
		"Unix time the pf counters of the interface were last cleared. OPNsense reports it in the firewall's local time without a zone, so it is interpreted in the exporter's time zone",
		[]string{"interface"},
	)

	c.stateTableEntries = buildPrometheusDesc(c.subsystem, "state_table_entries",
		"The current number of entries in the pf state table",
		nil,
	)

	c.stateTableOperations = buildPrometheusDesc(c.subsystem, "state_table_operations_total",
		"The number of pf state table operations by operation (searches, inserts, removals)",
		[]string{"operation"},
	)

	c.sourceTrackingEntries = buildPrometheusDesc(c.subsystem, "source_tracking_entries",
		"The current number of entries in the pf source tracking table",
		nil,
	)

	c.sourceTrackingOperations = buildPrometheusDesc(c.subsystem, "source_tracking_operations_total",
		"The number of pf source tracking table operations by operation (searches, inserts, removals)",
		[]string{"operation"},
	)

	c.pfCounters = buildPrometheusDesc(c.subsystem, "pf_counters_total",
		"The global pf counters by counter (match, bad-offset, fragment, short, normalize, memory, ...)",
		[]string{"counter"},
	)

	c.pfLimitCounters = buildPrometheusDesc(c.subsystem, "pf_limit_counters_total",
		"The global pf limit counters by counter (max-states-per-rule, max-src-states, ...)",
		[]string{"counter"},
	)

	c.memoryLimit = buildPrometheusDesc(c.subsystem, "memory_limit",
		"The hard limit of the pf memory pools by pool (states, src-nodes, frags, tables, table-entries, ...)",
		[]string{"pool"},
	)

	c.timeoutSeconds = buildPrometheusDesc(c.subsystem, "timeout_seconds",
		"The configured pf timeouts in seconds by timeout (tcp.first, tcp.established, udp.single, interval, ...)",
		[]string{"timeout"},
	)

	c.adaptiveStates = buildPrometheusDesc(c.subsystem, "adaptive_states",
		"The pf adaptive timeout thresholds in number of states by threshold (adaptive.start, adaptive.end)",
		[]string{"threshold"},
	)
}

func (c *firewallCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.inIPv4PassPackets
	ch <- c.outIPv4PassPackets
	ch <- c.inIPv4BlockPackets
	ch <- c.outIPv4BlockPackets

	ch <- c.inIPv6PassPackets
	ch <- c.outIPv6PassPackets
	ch <- c.inIPv6BlockPackets
	ch <- c.outIPv6BlockPackets

	ch <- c.inIPv4PassBytes
	ch <- c.outIPv4PassBytes
	ch <- c.inIPv4BlockBytes
	ch <- c.outIPv4BlockBytes

	ch <- c.inIPv6PassBytes
	ch <- c.outIPv6PassBytes
	ch <- c.inIPv6BlockBytes
	ch <- c.outIPv6BlockBytes

	ch <- c.interfaceReferences
	ch <- c.interfaceCountersReset

	ch <- c.stateTableEntries
	ch <- c.stateTableOperations
	ch <- c.sourceTrackingEntries
	ch <- c.sourceTrackingOperations
	ch <- c.pfCounters
	ch <- c.pfLimitCounters

	ch <- c.memoryLimit
	ch <- c.timeoutSeconds
	ch <- c.adaptiveStates
}

func (c *firewallCollector) Update(client *opnsense.Client, ch chan<- prometheus.Metric) *opnsense.APICallError {
	data, err := client.FetchPFStatsByInterface()
	if err != nil {
		return err
	}

	for _, v := range data.Interfaces {
		metricsValueMapping := map[*prometheus.Desc]int64{
			c.inIPv4PassPackets:   v.In4PassPackets,
			c.outIPv4PassPackets:  v.Out4PassPackets,
			c.inIPv4BlockPackets:  v.In4BlockPackets,
			c.outIPv4BlockPackets: v.Out4BlockPackets,
			c.inIPv6PassPackets:   v.In6PassPackets,
			c.outIPv6PassPackets:  v.Out6PassPackets,
			c.inIPv6BlockPackets:  v.In6BlockPackets,
			c.outIPv6BlockPackets: v.Out6BlockPackets,
		}
		for metric, value := range metricsValueMapping {
			ch <- prometheus.MustNewConstMetric(
				metric,
				prometheus.GaugeValue,
				float64(value),
				v.InterfaceName,
				c.instance,
			)
		}

		bytesValueMapping := map[*prometheus.Desc]int64{
			c.inIPv4PassBytes:   v.In4PassBytes,
			c.outIPv4PassBytes:  v.Out4PassBytes,
			c.inIPv4BlockBytes:  v.In4BlockBytes,
			c.outIPv4BlockBytes: v.Out4BlockBytes,
			c.inIPv6PassBytes:   v.In6PassBytes,
			c.outIPv6PassBytes:  v.Out6PassBytes,
			c.inIPv6BlockBytes:  v.In6BlockBytes,
			c.outIPv6BlockBytes: v.Out6BlockBytes,
		}
		for metric, value := range bytesValueMapping {
			ch <- prometheus.MustNewConstMetric(
				metric,
				prometheus.CounterValue,
				float64(value),
				v.InterfaceName,
				c.instance,
			)
		}

		ch <- prometheus.MustNewConstMetric(
			c.interfaceReferences,
			prometheus.GaugeValue,
			float64(v.References),
			v.InterfaceName,
			c.instance,
		)

		if v.Cleared != "" {
			cleared, parseErr := time.ParseInLocation("2006-01-02T15:04:05", v.Cleared, time.Local)
			if parseErr != nil {
				c.log.Debug("failed to parse pf interface cleared time",
					"collector", c.Name(), "interface", v.InterfaceName, "value", v.Cleared, "err", parseErr)
			} else {
				ch <- prometheus.MustNewConstMetric(
					c.interfaceCountersReset,
					prometheus.GaugeValue,
					float64(cleared.Unix()),
					v.InterfaceName,
					c.instance,
				)
			}
		}
	}

	// The global pf statistics come from separate endpoints. A failure in
	// one of them must not hide the others, so collect them all and
	// report the first error.
	var firstErr *opnsense.APICallError
	keep := func(err *opnsense.APICallError) {
		if firstErr == nil {
			firstErr = err
		}
	}

	if info, err := client.FetchPFInfo(); err != nil {
		keep(err)
	} else {
		c.updateTable(ch, info.StateTable, c.stateTableEntries, c.stateTableOperations)
		c.updateTable(ch, info.SourceTracking, c.sourceTrackingEntries, c.sourceTrackingOperations)
		for name, value := range info.Counters {
			ch <- prometheus.MustNewConstMetric(c.pfCounters, prometheus.CounterValue, value, name, c.instance)
		}
		for name, value := range info.LimitCounters {
			ch <- prometheus.MustNewConstMetric(c.pfLimitCounters, prometheus.CounterValue, value, name, c.instance)
		}
	}

	if limits, err := client.FetchPFMemoryLimits(); err != nil {
		keep(err)
	} else {
		for pool, value := range limits.Limits {
			ch <- prometheus.MustNewConstMetric(c.memoryLimit, prometheus.GaugeValue, float64(value), pool, c.instance)
		}
	}

	if timeouts, err := client.FetchPFTimeouts(); err != nil {
		keep(err)
	} else {
		for name, value := range timeouts.Seconds {
			ch <- prometheus.MustNewConstMetric(c.timeoutSeconds, prometheus.GaugeValue, value, name, c.instance)
		}
		for name, value := range timeouts.AdaptiveStates {
			ch <- prometheus.MustNewConstMetric(c.adaptiveStates, prometheus.GaugeValue, value, name, c.instance)
		}
	}

	return firstErr
}

// updateTable emits a pf table section from pfctl -si: "current-entries"
// becomes the entries gauge, every other row is an operation counter.
func (c *firewallCollector) updateTable(ch chan<- prometheus.Metric, table map[string]float64, entries, operations *prometheus.Desc) {
	for name, value := range table {
		if name == "current-entries" {
			ch <- prometheus.MustNewConstMetric(entries, prometheus.GaugeValue, value, c.instance)
			continue
		}
		ch <- prometheus.MustNewConstMetric(operations, prometheus.CounterValue, value, name, c.instance)
	}
}
