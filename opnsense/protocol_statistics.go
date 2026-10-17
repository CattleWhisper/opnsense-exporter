package opnsense

type protocolStatisticsResponse struct {
	Statistics struct {
		TCP struct {
			SentPackets                           flexCounter `json:"sent-packets"`
			SentDataPackets                       flexCounter `json:"sent-data-packets"`
			SentDataBytes                         flexCounter `json:"sent-data-bytes"`
			SentRetransmittedPackets              flexCounter `json:"sent-retransmitted-packets"`
			SentRetransmittedBytes                flexCounter `json:"sent-retransmitted-bytes"`
			SentUnnecessaryRetransmittedPackets   flexCounter `json:"sent-unnecessary-retransmitted-packets"`
			SentResendsByMtuDiscovery             flexCounter `json:"sent-resends-by-mtu-discovery"`
			SentAckOnlyPackets                    flexCounter `json:"sent-ack-only-packets"`
			SentPacketsDelayed                    flexCounter `json:"sent-packets-delayed"`
			SentUrgOnlyPackets                    flexCounter `json:"sent-urg-only-packets"`
			SentWindowProbePackets                flexCounter `json:"sent-window-probe-packets"`
			SentWindowUpdatePackets               flexCounter `json:"sent-window-update-packets"`
			SentControlPackets                    flexCounter `json:"sent-control-packets"`
			ReceivedPackets                       flexCounter `json:"received-packets"`
			ReceivedAckPackets                    flexCounter `json:"received-ack-packets"`
			ReceivedAckBytes                      flexCounter `json:"received-ack-bytes"`
			ReceivedDuplicateAcks                 flexCounter `json:"received-duplicate-acks"`
			ReceivedUDPTunneledPkts               flexCounter `json:"received-udp-tunneled-pkts"`
			ReceivedBadUDPTunneledPkts            flexCounter `json:"received-bad-udp-tunneled-pkts"`
			ReceivedAcksForUnsentData             flexCounter `json:"received-acks-for-unsent-data"`
			ReceivedInSequencePackets             flexCounter `json:"received-in-sequence-packets"`
			ReceivedInSequenceBytes               flexCounter `json:"received-in-sequence-bytes"`
			ReceivedCompletelyDuplicatePackets    flexCounter `json:"received-completely-duplicate-packets"`
			ReceivedCompletelyDuplicateBytes      flexCounter `json:"received-completely-duplicate-bytes"`
			ReceivedOldDuplicatePackets           flexCounter `json:"received-old-duplicate-packets"`
			ReceivedSomeDuplicatePackets          flexCounter `json:"received-some-duplicate-packets"`
			ReceivedSomeDuplicateBytes            flexCounter `json:"received-some-duplicate-bytes"`
			ReceivedOutOfOrder                    flexCounter `json:"received-out-of-order"`
			ReceivedOutOfOrderBytes               flexCounter `json:"received-out-of-order-bytes"`
			ReceivedAfterWindowPackets            flexCounter `json:"received-after-window-packets"`
			ReceivedAfterWindowBytes              flexCounter `json:"received-after-window-bytes"`
			ReceivedWindowProbes                  flexCounter `json:"received-window-probes"`
			ReceiveWindowUpdatePackets            flexCounter `json:"receive-window-update-packets"`
			ReceivedAfterClosePackets             flexCounter `json:"received-after-close-packets"`
			DiscardBadChecksum                    flexCounter `json:"discard-bad-checksum"`
			DiscardBadHeaderOffset                flexCounter `json:"discard-bad-header-offset"`
			DiscardTooShort                       flexCounter `json:"discard-too-short"`
			DiscardReassemblyQueueFull            flexCounter `json:"discard-reassembly-queue-full"`
			ConnectionRequests                    flexCounter `json:"connection-requests"`
			ConnectionsAccepts                    flexCounter `json:"connections-accepts"`
			BadConnectionAttempts                 flexCounter `json:"bad-connection-attempts"`
			ListenQueueOverflows                  flexCounter `json:"listen-queue-overflows"`
			IgnoredInWindowResets                 flexCounter `json:"ignored-in-window-resets"`
			ConnectionsEstablished                flexCounter `json:"connections-established"`
			ConnectionsHostcacheRtt               flexCounter `json:"connections-hostcache-rtt"`
			ConnectionsHostcacheRttvar            flexCounter `json:"connections-hostcache-rttvar"`
			ConnectionsHostcacheSsthresh          flexCounter `json:"connections-hostcache-ssthresh"`
			ConnectionsClosed                     flexCounter `json:"connections-closed"`
			ConnectionDrops                       flexCounter `json:"connection-drops"`
			ConnectionsUpdatedRttOnClose          flexCounter `json:"connections-updated-rtt-on-close"`
			ConnectionsUpdatedVarianceOnClose     flexCounter `json:"connections-updated-variance-on-close"`
			ConnectionsUpdatedSsthreshOnClose     flexCounter `json:"connections-updated-ssthresh-on-close"`
			EmbryonicConnectionsDropped           flexCounter `json:"embryonic-connections-dropped"`
			SegmentsUpdatedRtt                    flexCounter `json:"segments-updated-rtt"`
			SegmentUpdateAttempts                 flexCounter `json:"segment-update-attempts"`
			RetransmitTimeouts                    flexCounter `json:"retransmit-timeouts"`
			ConnectionsDroppedByRetransmitTimeout flexCounter `json:"connections-dropped-by-retransmit-timeout"`
			PersistTimeout                        flexCounter `json:"persist-timeout"`
			ConnectionsDroppedByPersistTimeout    flexCounter `json:"connections-dropped-by-persist-timeout"`
			ConnectionsDroppedByFinwait2Timeout   flexCounter `json:"connections-dropped-by-finwait2-timeout"`
			KeepaliveTimeout                      flexCounter `json:"keepalive-timeout"`
			KeepaliveProbes                       flexCounter `json:"keepalive-probes"`
			ConnectionsDroppedByKeepalives        flexCounter `json:"connections-dropped-by-keepalives"`
			AckHeaderPredictions                  flexCounter `json:"ack-header-predictions"`
			DataPacketHeaderPredictions           flexCounter `json:"data-packet-header-predictions"`
			Syncache                              struct {
				EntriesAdded   flexCounter `json:"entries-added"`
				Retransmitted  flexCounter `json:"retransmitted"`
				Duplicates     flexCounter `json:"duplicates"`
				Dropped        flexCounter `json:"dropped"`
				Completed      flexCounter `json:"completed"`
				BucketOverflow flexCounter `json:"bucket-overflow"`
				CacheOverflow  flexCounter `json:"cache-overflow"`
				Reset          flexCounter `json:"reset"`
				Stale          flexCounter `json:"stale"`
				Aborted        flexCounter `json:"aborted"`
				BadAck         flexCounter `json:"bad-ack"`
				Unreachable    flexCounter `json:"unreachable"`
				ZoneFailures   flexCounter `json:"zone-failures"`
				SentCookies    flexCounter `json:"sent-cookies"`
				ReceivdCookies flexCounter `json:"receivd-cookies"`
			} `json:"syncache"`
			Hostcache struct {
				EntriesAdded    flexCounter `json:"entries-added"`
				BufferOverflows flexCounter `json:"buffer-overflows"`
			} `json:"hostcache"`
			Sack struct {
				RecoveryEpisodes    flexCounter `json:"recovery-episodes"`
				SegmentRetransmits  flexCounter `json:"segment-retransmits"`
				ByteRetransmits     flexCounter `json:"byte-retransmits"`
				ReceivedBlocks      flexCounter `json:"received-blocks"`
				SentOptionBlocks    flexCounter `json:"sent-option-blocks"`
				ScoreboardOverflows flexCounter `json:"scoreboard-overflows"`
			} `json:"sack"`
			Ecn struct {
				CePackets            flexCounter `json:"ce-packets"`
				Ect0Packets          flexCounter `json:"ect0-packets"`
				Ect1Packets          flexCounter `json:"ect1-packets"`
				Handshakes           flexCounter `json:"handshakes"`
				CongestionReductions flexCounter `json:"congestion-reductions"`
			} `json:"ecn"`
			TCPSignature struct {
				ReceivedGoodSignature flexCounter `json:"received-good-signature"`
				ReceivedBadSignature  flexCounter `json:"received-bad-signature"`
				FailedMakeSignature   flexCounter `json:"failed-make-signature"`
				NoSignatureExpected   flexCounter `json:"no-signature-expected"`
				NoSignatureProvided   flexCounter `json:"no-signature-provided"`
			} `json:"tcp-signature"`
			Pmtud struct {
				PmtudActivated       flexCounter `json:"pmtud-activated"`
				PmtudActivatedMinMss flexCounter `json:"pmtud-activated-min-mss"`
				PmtudFailed          flexCounter `json:"pmtud-failed"`
			} `json:"pmtud"`
			Tw struct {
				TwResponds flexCounter `json:"tw_responds"`
				TwRecycles flexCounter `json:"tw_recycles"`
				TwResets   flexCounter `json:"tw_resets"`
			} `json:"tw"`
			TCPConnectionCountByState struct {
				Closed      flexCounter `json:"CLOSED"`
				Listen      flexCounter `json:"LISTEN"`
				SynSent     flexCounter `json:"SYN_SENT"`
				SynRcvd     flexCounter `json:"SYN_RCVD"`
				Established flexCounter `json:"ESTABLISHED"`
				CloseWait   flexCounter `json:"CLOSE_WAIT"`
				FinWait1    flexCounter `json:"FIN_WAIT_1"`
				Closing     flexCounter `json:"CLOSING"`
				LastAck     flexCounter `json:"LAST_ACK"`
				FinWait2    flexCounter `json:"FIN_WAIT_2"`
				TimeWait    flexCounter `json:"TIME_WAIT"`
			} `json:"TCP connection count by state"`
		} `json:"tcp"`
		UDP struct {
			ReceivedDatagrams            flexCounter `json:"received-datagrams"`
			DroppedIncompleteHeaders     flexCounter `json:"dropped-incomplete-headers"`
			DroppedBadDataLength         flexCounter `json:"dropped-bad-data-length"`
			DroppedBadChecksum           flexCounter `json:"dropped-bad-checksum"`
			DroppedNoChecksum            flexCounter `json:"dropped-no-checksum"`
			DroppedNoSocket              flexCounter `json:"dropped-no-socket"`
			DroppedBroadcastMulticast    flexCounter `json:"dropped-broadcast-multicast"`
			DroppedFullSocketBuffer      flexCounter `json:"dropped-full-socket-buffer"`
			NotForHashedPcb              flexCounter `json:"not-for-hashed-pcb"`
			DeliveredPackets             flexCounter `json:"delivered-packets"`
			OutputPackets                flexCounter `json:"output-packets"`
			MulticastSourceFilterMatches flexCounter `json:"multicast-source-filter-matches"`
		} `json:"udp"`
		IP struct {
			ReceivedPackets               flexCounter `json:"received-packets"`
			DroppedBadChecksum            flexCounter `json:"dropped-bad-checksum"`
			DroppedBelowMinimumSize       flexCounter `json:"dropped-below-minimum-size"`
			DroppedShortPackets           flexCounter `json:"dropped-short-packets"`
			DroppedTooLong                flexCounter `json:"dropped-too-long"`
			DroppedShortHeaderLength      flexCounter `json:"dropped-short-header-length"`
			DroppedShortData              flexCounter `json:"dropped-short-data"`
			DroppedBadOptions             flexCounter `json:"dropped-bad-options"`
			DroppedBadVersion             flexCounter `json:"dropped-bad-version"`
			ReceivedFragments             flexCounter `json:"received-fragments"`
			DroppedFragments              flexCounter `json:"dropped-fragments"`
			DroppedFragmentsAfterTimeout  flexCounter `json:"dropped-fragments-after-timeout"`
			ReassembledPackets            flexCounter `json:"reassembled-packets"`
			ReceivedLocalPackets          flexCounter `json:"received-local-packets"`
			DroppedUnknownProtocol        flexCounter `json:"dropped-unknown-protocol"`
			ForwardedPackets              flexCounter `json:"forwarded-packets"`
			FastForwardedPackets          flexCounter `json:"fast-forwarded-packets"`
			PacketsCannotForward          flexCounter `json:"packets-cannot-forward"`
			ReceivedUnknownMulticastGroup flexCounter `json:"received-unknown-multicast-group"`
			RedirectsSent                 flexCounter `json:"redirects-sent"`
			SentPackets                   flexCounter `json:"sent-packets"`
			SendPacketsFabricatedHeader   flexCounter `json:"send-packets-fabricated-header"`
			DiscardNoMbufs                flexCounter `json:"discard-no-mbufs"`
			DiscardNoRoute                flexCounter `json:"discard-no-route"`
			SentFragments                 flexCounter `json:"sent-fragments"`
			FragmentsCreated              flexCounter `json:"fragments-created"`
			DiscardCannotFragment         flexCounter `json:"discard-cannot-fragment"`
			DiscardTunnelNoGif            flexCounter `json:"discard-tunnel-no-gif"`
			DiscardBadAddress             flexCounter `json:"discard-bad-address"`
		} `json:"ip"`
		Icmp struct {
			IcmpCalls                   flexCounter `json:"icmp-calls"`
			ErrorsNotFromMessage        flexCounter `json:"errors-not-from-message"`
			DroppedBadCode              flexCounter `json:"dropped-bad-code"`
			DroppedTooShort             flexCounter `json:"dropped-too-short"`
			DroppedBadChecksum          flexCounter `json:"dropped-bad-checksum"`
			DroppedBadLength            flexCounter `json:"dropped-bad-length"`
			DroppedMulticastEcho        flexCounter `json:"dropped-multicast-echo"`
			DroppedMulticastTimestamp   flexCounter `json:"dropped-multicast-timestamp"`
			SentPackets                 flexCounter `json:"sent-packets"`
			DiscardInvalidReturnAddress flexCounter `json:"discard-invalid-return-address"`
			DiscardNoRoute              flexCounter `json:"discard-no-route"`
			IcmpAddressResponses        string      `json:"icmp-address-responses"`
		} `json:"icmp"`
		Carp struct {
			ReceivedInetPackets      flexCounter `json:"received-inet-packets"`
			ReceivedInet6Packets     flexCounter `json:"received-inet6-packets"`
			DroppedWrongTTL          flexCounter `json:"dropped-wrong-ttl"`
			DroppedShortHeader       flexCounter `json:"dropped-short-header"`
			DroppedBadChecksum       flexCounter `json:"dropped-bad-checksum"`
			DroppedBadVersion        flexCounter `json:"dropped-bad-version"`
			DroppedShortPacket       flexCounter `json:"dropped-short-packet"`
			DroppedBadAuthentication flexCounter `json:"dropped-bad-authentication"`
			DroppedBadVhid           flexCounter `json:"dropped-bad-vhid"`
			DroppedBadAddressList    flexCounter `json:"dropped-bad-address-list"`
			SentInetPackets          flexCounter `json:"sent-inet-packets"`
			SentInet6Packets         flexCounter `json:"sent-inet6-packets"`
			SendFailedMemoryError    flexCounter `json:"send-failed-memory-error"`
		} `json:"carp"`
		Pfsync struct {
			ReceivedInetPackets  flexCounter `json:"received-inet-packets"`
			ReceivedInet6Packets flexCounter `json:"received-inet6-packets"`
			InputHistogram       []struct {
				Name  string      `json:"name"`
				Count flexCounter `json:"count"`
			} `json:"input-histogram"`
			DroppedBadInterface flexCounter `json:"dropped-bad-interface"`
			DroppedBadTTL       flexCounter `json:"dropped-bad-ttl"`
			DroppedShortHeader  flexCounter `json:"dropped-short-header"`
			DroppedBadVersion   flexCounter `json:"dropped-bad-version"`
			DroppedBadAuth      flexCounter `json:"dropped-bad-auth"`
			DroppedBadAction    flexCounter `json:"dropped-bad-action"`
			DroppedShort        flexCounter `json:"dropped-short"`
			DroppedBadValues    flexCounter `json:"dropped-bad-values"`
			DroppedStaleState   flexCounter `json:"dropped-stale-state"`
			DroppedFailedLookup flexCounter `json:"dropped-failed-lookup"`
			SentInetPackets     flexCounter `json:"sent-inet-packets"`
			SendInet6Packets    flexCounter `json:"send-inet6-packets"`
			OutputHistogram     []struct {
				Name  string      `json:"name"`
				Count flexCounter `json:"count"`
			} `json:"output-histogram"`
			DiscardedNoMemory flexCounter `json:"discarded-no-memory"`
			SendErrors        flexCounter `json:"send-errors"`
		} `json:"pfsync"`
		Arp struct {
			SentRequests            flexCounter `json:"sent-requests"`
			SentFailures            flexCounter `json:"sent-failures"`
			SentReplies             flexCounter `json:"sent-replies"`
			ReceivedRequests        flexCounter `json:"received-requests"`
			ReceivedReplies         flexCounter `json:"received-replies"`
			ReceivedPackets         flexCounter `json:"received-packets"`
			DroppedNoEntry          flexCounter `json:"dropped-no-entry"`
			EntriesTimeout          flexCounter `json:"entries-timeout"`
			DroppedDuplicateAddress flexCounter `json:"dropped-duplicate-address"`
		} `json:"arp"`
	} `json:"statistics"`
}

type ProtocolStatistics struct {
	TCPSentPackets            int64
	TCPReceivedPackets        int64
	ARPSentRequests           int64
	ARPReceivedRequests       int64
	TCPConnectionCountByState map[string]int64
	ICMPCalls                 int64
	ICMPSentPackets           int64
	ICMPDroppedByReason       map[string]int64
	UDPDeliveredPackets       int64
	UDPOutputPackets          int64
	UDPReceivedDatagrams      int64
	UDPDroppedByReason        map[string]int64
}

func (c *Client) FetchProtocolStatistics() (ProtocolStatistics, *APICallError) {
	var resp protocolStatisticsResponse
	url, ok := c.endpoints["protocolStatistics"]
	if !ok {
		return ProtocolStatistics{}, &APICallError{
			Endpoint:   "protocolStatistics",
			StatusCode: 404,
			Message:    "endpoint not found in client endpoints",
		}
	}
	if err := c.do("GET", url, nil, &resp); err != nil {
		return ProtocolStatistics{}, err
	}

	out := ProtocolStatistics{
		TCPSentPackets:      int64(resp.Statistics.TCP.SentPackets),
		TCPReceivedPackets:  int64(resp.Statistics.TCP.ReceivedPackets),
		ARPSentRequests:     int64(resp.Statistics.Arp.SentRequests),
		ARPReceivedRequests: int64(resp.Statistics.Arp.ReceivedRequests),
		TCPConnectionCountByState: map[string]int64{
			"CLOSED":      int64(resp.Statistics.TCP.TCPConnectionCountByState.Closed),
			"LISTEN":      int64(resp.Statistics.TCP.TCPConnectionCountByState.Listen),
			"SYN_SENT":    int64(resp.Statistics.TCP.TCPConnectionCountByState.SynSent),
			"SYN_RCVD":    int64(resp.Statistics.TCP.TCPConnectionCountByState.SynRcvd),
			"ESTABLISHED": int64(resp.Statistics.TCP.TCPConnectionCountByState.Established),
			"CLOSE_WAIT":  int64(resp.Statistics.TCP.TCPConnectionCountByState.CloseWait),
			"FIN_WAIT_1":  int64(resp.Statistics.TCP.TCPConnectionCountByState.FinWait1),
			"CLOSING":     int64(resp.Statistics.TCP.TCPConnectionCountByState.Closing),
			"LAST_ACK":    int64(resp.Statistics.TCP.TCPConnectionCountByState.LastAck),
			"FIN_WAIT_2":  int64(resp.Statistics.TCP.TCPConnectionCountByState.FinWait2),
			"TIME_WAIT":   int64(resp.Statistics.TCP.TCPConnectionCountByState.TimeWait),
		},
		ICMPCalls:       int64(resp.Statistics.Icmp.IcmpCalls),
		ICMPSentPackets: int64(resp.Statistics.Icmp.SentPackets),
		ICMPDroppedByReason: map[string]int64{
			"BAD_CODE":            int64(resp.Statistics.Icmp.DroppedBadCode),
			"TOO_SHORT":           int64(resp.Statistics.Icmp.DroppedTooShort),
			"BAD_CHECKSUM":        int64(resp.Statistics.Icmp.DroppedBadChecksum),
			"BAD_LENGTH":          int64(resp.Statistics.Icmp.DroppedBadLength),
			"MULTICAST_ECHO":      int64(resp.Statistics.Icmp.DroppedMulticastEcho),
			"MULTICAST_TIMESTAMP": int64(resp.Statistics.Icmp.DroppedMulticastTimestamp),
		},
		UDPDeliveredPackets:  int64(resp.Statistics.UDP.DeliveredPackets),
		UDPOutputPackets:     int64(resp.Statistics.UDP.OutputPackets),
		UDPReceivedDatagrams: int64(resp.Statistics.UDP.ReceivedDatagrams),
		UDPDroppedByReason: map[string]int64{
			"INCOMPLETE_HEADERS":  int64(resp.Statistics.UDP.DroppedIncompleteHeaders),
			"BAD_DATA_LENGTH":     int64(resp.Statistics.UDP.DroppedBadDataLength),
			"BAD_CHECKSUM":        int64(resp.Statistics.UDP.DroppedBadChecksum),
			"NO_CHECKSUM":         int64(resp.Statistics.UDP.DroppedNoChecksum),
			"NO_SOCKET":           int64(resp.Statistics.UDP.DroppedNoSocket),
			"BROADCAST_MULTICAST": int64(resp.Statistics.UDP.DroppedBroadcastMulticast),
			"FULL_SOCKET_BUFFER":  int64(resp.Statistics.UDP.DroppedFullSocketBuffer),
		},
	}

	return out, nil
}
