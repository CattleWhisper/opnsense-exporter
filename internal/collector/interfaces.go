package collector

import (
	"log/slog"

	"github.com/AthennaMind/opnsense-exporter/opnsense"
	"github.com/prometheus/client_golang/prometheus"
)

type interfacesCollector struct {
	log *slog.Logger

	mtu                   *prometheus.Desc
	bytesReceived         *prometheus.Desc
	bytesTransmited       *prometheus.Desc
	multicastsTransmitted *prometheus.Desc
	multicastsReceived    *prometheus.Desc
	inputErrors           *prometheus.Desc
	outputErrors          *prometheus.Desc
	collisions            *prometheus.Desc

	packetsReceived      *prometheus.Desc
	packetsTransmitted   *prometheus.Desc
	inputQueueDrops      *prometheus.Desc
	sendQueueDrops       *prometheus.Desc
	sendQueueLength      *prometheus.Desc
	sendQueueMaxLength   *prometheus.Desc
	unknownProtocol      *prometheus.Desc
	promiscuousListeners *prometheus.Desc
	linkState            *prometheus.Desc
	lineRate             *prometheus.Desc
	statsResetUptime     *prometheus.Desc

	subsystem string
	instance  string
}

func init() {
	collectorInstances = append(collectorInstances, &interfacesCollector{
		subsystem: InterfacesSubsystem,
	})
}

func (c *interfacesCollector) Name() string {
	return c.subsystem
}

func (c *interfacesCollector) Register(namespace, instanceLabel string, log *slog.Logger) {
	c.log = log
	c.instance = instanceLabel

	c.log.Debug("Registering collector", "collector", c.Name())

	c.mtu = buildPrometheusDesc(c.subsystem, "mtu_bytes",
		"The MTU value of the interface",
		[]string{"interface", "device", "type"},
	)
	c.bytesReceived = buildPrometheusDesc(c.subsystem, "received_bytes_total",
		"Bytes received on this interface by interface name and device",
		[]string{"interface", "device", "type"},
	)
	c.bytesTransmited = buildPrometheusDesc(c.subsystem, "transmitted_bytes_total",
		"Bytes transmitted on this interface by interface name and device",
		[]string{"interface", "device", "type"},
	)
	c.multicastsReceived = buildPrometheusDesc(c.subsystem, "received_multicasts_total",
		"Multicasts received on this interface by interface name and device",
		[]string{"interface", "device", "type"},
	)
	c.multicastsTransmitted = buildPrometheusDesc(c.subsystem, "transmitted_multicasts_total",
		"Multicasts transmitted on this interface by interface name and device",
		[]string{"interface", "device", "type"},
	)
	c.inputErrors = buildPrometheusDesc(c.subsystem, "input_errors_total",
		"Input errors on this interface by interface name and device",
		[]string{"interface", "device", "type"},
	)
	c.outputErrors = buildPrometheusDesc(c.subsystem, "output_errors_total",
		"Output errors on this interface by interface name and device",
		[]string{"interface", "device", "type"},
	)
	c.collisions = buildPrometheusDesc(c.subsystem, "collisions_total",
		"Collisions on this interface by interface name and device",
		[]string{"interface", "device", "type"},
	)
	c.packetsReceived = buildPrometheusDesc(c.subsystem, "received_packets_total",
		"Packets received on this interface by interface name and device",
		[]string{"interface", "device", "type"},
	)
	c.packetsTransmitted = buildPrometheusDesc(c.subsystem, "transmitted_packets_total",
		"Packets transmitted on this interface by interface name and device",
		[]string{"interface", "device", "type"},
	)
	c.inputQueueDrops = buildPrometheusDesc(c.subsystem, "input_queue_drops_total",
		"Packets dropped on input because the receive queue was full, by interface name and device",
		[]string{"interface", "device", "type"},
	)
	c.sendQueueDrops = buildPrometheusDesc(c.subsystem, "send_queue_drops_total",
		"Packets dropped on output because the send queue was full, by interface name and device",
		[]string{"interface", "device", "type"},
	)
	c.sendQueueLength = buildPrometheusDesc(c.subsystem, "send_queue_length",
		"Current length of the send queue by interface name and device",
		[]string{"interface", "device", "type"},
	)
	c.sendQueueMaxLength = buildPrometheusDesc(c.subsystem, "send_queue_max_length",
		"Maximum length of the send queue by interface name and device",
		[]string{"interface", "device", "type"},
	)
	c.unknownProtocol = buildPrometheusDesc(c.subsystem, "unknown_protocol_packets_total",
		"Packets received for an unknown or unsupported protocol by interface name and device",
		[]string{"interface", "device", "type"},
	)
	c.promiscuousListeners = buildPrometheusDesc(c.subsystem, "promiscuous_listeners",
		"Number of promiscuous listeners on this interface by interface name and device",
		[]string{"interface", "device", "type"},
	)
	c.linkState = buildPrometheusDesc(c.subsystem, "link_state",
		"Link state of the interface (0 = unknown, 1 = down, 2 = up) by interface name and device",
		[]string{"interface", "device", "type"},
	)
	c.lineRate = buildPrometheusDesc(c.subsystem, "line_rate_bits",
		"Line rate of the interface in bits per second by interface name and device",
		[]string{"interface", "device", "type"},
	)
	c.statsResetUptime = buildPrometheusDesc(c.subsystem, "stats_reset_uptime_seconds",
		"System uptime in seconds when the interface was attached or its counters were last reset, by interface name and device",
		[]string{"interface", "device", "type"},
	)
}

func (c *interfacesCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.mtu
	ch <- c.bytesReceived
	ch <- c.bytesTransmited
	ch <- c.multicastsReceived
	ch <- c.multicastsTransmitted
	ch <- c.inputErrors
	ch <- c.outputErrors
	ch <- c.collisions
	ch <- c.packetsReceived
	ch <- c.packetsTransmitted
	ch <- c.inputQueueDrops
	ch <- c.sendQueueDrops
	ch <- c.sendQueueLength
	ch <- c.sendQueueMaxLength
	ch <- c.unknownProtocol
	ch <- c.promiscuousListeners
	ch <- c.linkState
	ch <- c.lineRate
	ch <- c.statsResetUptime
}

func (c *interfacesCollector) update(ch chan<- prometheus.Metric, desc *prometheus.Desc, valueType prometheus.ValueType, value float64, labelValues ...string) {
	ch <- prometheus.MustNewConstMetric(
		desc, valueType, value, labelValues...,
	)
}

func (c *interfacesCollector) Update(client *opnsense.Client, ch chan<- prometheus.Metric) *opnsense.APICallError {
	data, err := client.FetchInterfaces()
	if err != nil {
		return err
	}

	for _, iface := range data.Interfaces {
		c.update(ch, c.mtu, prometheus.GaugeValue, float64(iface.MTU), iface.Name, iface.Device, iface.Type, c.instance)
		c.update(ch, c.bytesReceived, prometheus.CounterValue, float64(iface.BytesReceived), iface.Name, iface.Device, iface.Type, c.instance)
		c.update(ch, c.bytesTransmited, prometheus.CounterValue, float64(iface.BytesTransmitted), iface.Name, iface.Device, iface.Type, c.instance)
		c.update(ch, c.multicastsReceived, prometheus.CounterValue, float64(iface.MulticastsReceived), iface.Name, iface.Device, iface.Type, c.instance)
		c.update(ch, c.multicastsTransmitted, prometheus.CounterValue, float64(iface.MulticastsTransmitted), iface.Name, iface.Device, iface.Type, c.instance)
		c.update(ch, c.inputErrors, prometheus.CounterValue, float64(iface.InputErrors), iface.Name, iface.Device, iface.Type, c.instance)
		c.update(ch, c.outputErrors, prometheus.CounterValue, float64(iface.OutputErrors), iface.Name, iface.Device, iface.Type, c.instance)
		c.update(ch, c.collisions, prometheus.CounterValue, float64(iface.Collisions), iface.Name, iface.Device, iface.Type, c.instance)
		c.update(ch, c.packetsReceived, prometheus.CounterValue, float64(iface.PacketsReceived), iface.Name, iface.Device, iface.Type, c.instance)
		c.update(ch, c.packetsTransmitted, prometheus.CounterValue, float64(iface.PacketsTransmitted), iface.Name, iface.Device, iface.Type, c.instance)

		optional := []struct {
			desc      *prometheus.Desc
			valueType prometheus.ValueType
			value     *float64
		}{
			{c.inputQueueDrops, prometheus.CounterValue, iface.InputQueueDrops},
			{c.sendQueueDrops, prometheus.CounterValue, iface.SendQueueDrops},
			{c.sendQueueLength, prometheus.GaugeValue, iface.SendQueueLength},
			{c.sendQueueMaxLength, prometheus.GaugeValue, iface.SendQueueMaxLength},
			{c.unknownProtocol, prometheus.CounterValue, iface.PacketsForUnknownProtocol},
			{c.promiscuousListeners, prometheus.GaugeValue, iface.PromiscuousListeners},
			{c.linkState, prometheus.GaugeValue, iface.LinkState},
			{c.lineRate, prometheus.GaugeValue, iface.LineRateBits},
			{c.statsResetUptime, prometheus.GaugeValue, iface.UptimeAtAttachOrStatReset},
		}
		for _, m := range optional {
			if m.value == nil {
				continue
			}
			c.update(ch, m.desc, m.valueType, *m.value, iface.Name, iface.Device, iface.Type, c.instance)
		}
	}

	return nil
}
