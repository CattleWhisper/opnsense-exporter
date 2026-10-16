# Grafana dashboards

| File | Dashboard | What it shows |
| --- | --- | --- |
| `opnsense-wan-ip-family.json` | OPNsense / WAN - IPv4 vs IPv6 | WAN download/upload split by IP family, IPv6 share, daily volume, packets, average packet size and blocked traffic. Pick the WAN by its pf name (e.g. `igb0`, `pppoe0`). |
| `opnsense-interfaces.json` | OPNsense / Interfaces | Link state, line rate, throughput, utilization, packets, multicast, errors, queue drops and send queue per interface. |
| `opnsense-firewall.json` | OPNsense / Firewall (pf) | State table usage, state inserts/removals, rule matches, pf anomaly and limit counters, source tracking, allowed/blocked traffic per pf interface, memory limits and timeouts. |
| `dashboard-v1.json` | OPNSense | The original overview: exporter status, services, gateways, interfaces and ARP. |

All dashboards share the `opnsense` tag and link to each other, keeping the
selected instance and time range.

## Import

Dashboards → New → Import → upload the JSON file. Each dashboard asks for a
Prometheus data source through its `Data source` variable, so no `__inputs`
mapping is needed.

## Requirements

The three new dashboards use metrics added after v0.0.17: the per-family
pf byte counters (`opnsense_firewall_*_bytes_total`), the extra interface
statistics and the global pf statistics. See [docs/metrics.md](../../docs/metrics.md).

Note that the `interface` label means different things in the two
collectors: in `opnsense_interfaces_*` it is the OPNsense description
(e.g. `WAN`) with the OS name in `device`, while in `opnsense_firewall_*`
it is the pf/OS name (e.g. `igb0`).
