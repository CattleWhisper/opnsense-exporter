package opnsense

type protocolStatisticsResponse struct {
	Statistics struct {
		TCP struct {
			SentPackets                           int64 `json:"sent-packets"`
			SentDataPackets                       int64 `json:"sent-data-packets"`
			SentDataBytes                         int64 `json:"sent-data-bytes"`
			SentRetransmittedPackets              int64 `json:"sent-retransmitted-packets"`
			SentRetransmittedBytes                int64 `json:"sent-retransmitted-bytes"`
			SentUnnecessaryRetransmittedPackets   int64 `json:"sent-unnecessary-retransmitted-packets"`
			SentResendsByMtuDiscovery             int64 `json:"sent-resends-by-mtu-discovery"`
			SentAckOnlyPackets                    int64 `json:"sent-ack-only-packets"`
			SentPacketsDelayed                    int64 `json:"sent-packets-delayed"`
			SentUrgOnlyPackets                    int64 `json:"sent-urg-only-packets"`
			SentWindowProbePackets                int64 `json:"sent-window-probe-packets"`
			SentWindowUpdatePackets               int64 `json:"sent-window-update-packets"`
			SentControlPackets                    int64 `json:"sent-control-packets"`
			ReceivedPackets                       int64 `json:"received-packets"`
			ReceivedAckPackets                    int64 `json:"received-ack-packets"`
			ReceivedAckBytes                      int64 `json:"received-ack-bytes"`
			ReceivedDuplicateAcks                 int64 `json:"received-duplicate-acks"`
			ReceivedUDPTunneledPkts               int64 `json:"received-udp-tunneled-pkts"`
			ReceivedBadUDPTunneledPkts            int64 `json:"received-bad-udp-tunneled-pkts"`
			ReceivedAcksForUnsentData             int64 `json:"received-acks-for-unsent-data"`
			ReceivedInSequencePackets             int64 `json:"received-in-sequence-packets"`
			ReceivedInSequenceBytes               int64 `json:"received-in-sequence-bytes"`
			ReceivedCompletelyDuplicatePackets    int64 `json:"received-completely-duplicate-packets"`
			ReceivedCompletelyDuplicateBytes      int64 `json:"received-completely-duplicate-bytes"`
			ReceivedOldDuplicatePackets           int64 `json:"received-old-duplicate-packets"`
			ReceivedSomeDuplicatePackets          int64 `json:"received-some-duplicate-packets"`
			ReceivedSomeDuplicateBytes            int64 `json:"received-some-duplicate-bytes"`
			ReceivedOutOfOrder                    int64 `json:"received-out-of-order"`
			ReceivedOutOfOrderBytes               int64 `json:"received-out-of-order-bytes"`
			ReceivedAfterWindowPackets            int64 `json:"received-after-window-packets"`
			ReceivedAfterWindowBytes              int64 `json:"received-after-window-bytes"`
			ReceivedWindowProbes                  int64 `json:"received-window-probes"`
			ReceiveWindowUpdatePackets            int64 `json:"receive-window-update-packets"`
			ReceivedAfterClosePackets             int64 `json:"received-after-close-packets"`
			DiscardBadChecksum                    int64 `json:"discard-bad-checksum"`
			DiscardBadHeaderOffset                int64 `json:"discard-bad-header-offset"`
			DiscardTooShort                       int64 `json:"discard-too-short"`
			DiscardReassemblyQueueFull            int64 `json:"discard-reassembly-queue-full"`
			ConnectionRequests                    int64 `json:"connection-requests"`
			ConnectionsAccepts                    int64 `json:"connections-accepts"`
			BadConnectionAttempts                 int64 `json:"bad-connection-attempts"`
			ListenQueueOverflows                  int64 `json:"listen-queue-overflows"`
			IgnoredInWindowResets                 int64 `json:"ignored-in-window-resets"`
			ConnectionsEstablished                int64 `json:"connections-established"`
			ConnectionsHostcacheRtt               int64 `json:"connections-hostcache-rtt"`
			ConnectionsHostcacheRttvar            int64 `json:"connections-hostcache-rttvar"`
			ConnectionsHostcacheSsthresh          int64 `json:"connections-hostcache-ssthresh"`
			ConnectionsClosed                     int64 `json:"connections-closed"`
			ConnectionDrops                       int64 `json:"connection-drops"`
			ConnectionsUpdatedRttOnClose          int64 `json:"connections-updated-rtt-on-close"`
			ConnectionsUpdatedVarianceOnClose     int64 `json:"connections-updated-variance-on-close"`
			ConnectionsUpdatedSsthreshOnClose     int64 `json:"connections-updated-ssthresh-on-close"`
			EmbryonicConnectionsDropped           int64 `json:"embryonic-connections-dropped"`
			SegmentsUpdatedRtt                    int64 `json:"segments-updated-rtt"`
			SegmentUpdateAttempts                 int64 `json:"segment-update-attempts"`
			RetransmitTimeouts                    int64 `json:"retransmit-timeouts"`
			ConnectionsDroppedByRetransmitTimeout int64 `json:"connections-dropped-by-retransmit-timeout"`
			PersistTimeout                        int64 `json:"persist-timeout"`
			ConnectionsDroppedByPersistTimeout    int64 `json:"connections-dropped-by-persist-timeout"`
			ConnectionsDroppedByFinwait2Timeout   int64 `json:"connections-dropped-by-finwait2-timeout"`
			KeepaliveTimeout                      int64 `json:"keepalive-timeout"`
			KeepaliveProbes                       int64 `json:"keepalive-probes"`
			ConnectionsDroppedByKeepalives        int64 `json:"connections-dropped-by-keepalives"`
			AckHeaderPredictions                  int64 `json:"ack-header-predictions"`
			DataPacketHeaderPredictions           int64 `json:"data-packet-header-predictions"`
			Syncache                              struct {
				EntriesAdded   int64 `json:"entries-added"`
				Retransmitted  int64 `json:"retransmitted"`
				Duplicates     int64 `json:"duplicates"`
				Dropped        int64 `json:"dropped"`
				Completed      int64 `json:"completed"`
				BucketOverflow int64 `json:"bucket-overflow"`
				CacheOverflow  int64 `json:"cache-overflow"`
				Reset          int64 `json:"reset"`
				Stale          int64 `json:"stale"`
				Aborted        int64 `json:"aborted"`
				BadAck         int64 `json:"bad-ack"`
				Unreachable    int64 `json:"unreachable"`
				ZoneFailures   int64 `json:"zone-failures"`
				SentCookies    int64 `json:"sent-cookies"`
				ReceivdCookies int64 `json:"receivd-cookies"`
			} `json:"syncache"`
			Hostcache struct {
				EntriesAdded    int64 `json:"entries-added"`
				BufferOverflows int64 `json:"buffer-overflows"`
			} `json:"hostcache"`
			Sack struct {
				RecoveryEpisodes    int64 `json:"recovery-episodes"`
				SegmentRetransmits  int64 `json:"segment-retransmits"`
				ByteRetransmits     int64 `json:"byte-retransmits"`
				ReceivedBlocks      int64 `json:"received-blocks"`
				SentOptionBlocks    int64 `json:"sent-option-blocks"`
				ScoreboardOverflows int64 `json:"scoreboard-overflows"`
			} `json:"sack"`
			Ecn struct {
				CePackets            int64 `json:"ce-packets"`
				Ect0Packets          int64 `json:"ect0-packets"`
				Ect1Packets          int64 `json:"ect1-packets"`
				Handshakes           int64 `json:"handshakes"`
				CongestionReductions int64 `json:"congestion-reductions"`
			} `json:"ecn"`
			TCPSignature struct {
				ReceivedGoodSignature int64 `json:"received-good-signature"`
				ReceivedBadSignature  int64 `json:"received-bad-signature"`
				FailedMakeSignature   int64 `json:"failed-make-signature"`
				NoSignatureExpected   int64 `json:"no-signature-expected"`
				NoSignatureProvided   int64 `json:"no-signature-provided"`
			} `json:"tcp-signature"`
			Pmtud struct {
				PmtudActivated       int64 `json:"pmtud-activated"`
				PmtudActivatedMinMss int64 `json:"pmtud-activated-min-mss"`
				PmtudFailed          int64 `json:"pmtud-failed"`
			} `json:"pmtud"`
			Tw struct {
				TwResponds int64 `json:"tw_responds"`
				TwRecycles int64 `json:"tw_recycles"`
				TwResets   int64 `json:"tw_resets"`
			} `json:"tw"`
			TCPConnectionCountByState struct {
				Closed      int64 `json:"CLOSED"`
				Listen      int64 `json:"LISTEN"`
				SynSent     int64 `json:"SYN_SENT"`
				SynRcvd     int64 `json:"SYN_RCVD"`
				Established int64 `json:"ESTABLISHED"`
				CloseWait   int64 `json:"CLOSE_WAIT"`
				FinWait1    int64 `json:"FIN_WAIT_1"`
				Closing     int64 `json:"CLOSING"`
				LastAck     int64 `json:"LAST_ACK"`
				FinWait2    int64 `json:"FIN_WAIT_2"`
				TimeWait    int64 `json:"TIME_WAIT"`
			} `json:"TCP connection count by state"`
		} `json:"tcp"`
		UDP struct {
			ReceivedDatagrams            int64 `json:"received-datagrams"`
			DroppedIncompleteHeaders     int64 `json:"dropped-incomplete-headers"`
			DroppedBadDataLength         int64 `json:"dropped-bad-data-length"`
			DroppedBadChecksum           int64 `json:"dropped-bad-checksum"`
			DroppedNoChecksum            int64 `json:"dropped-no-checksum"`
			DroppedNoSocket              int64 `json:"dropped-no-socket"`
			DroppedBroadcastMulticast    int64 `json:"dropped-broadcast-multicast"`
			DroppedFullSocketBuffer      int64 `json:"dropped-full-socket-buffer"`
			NotForHashedPcb              int64 `json:"not-for-hashed-pcb"`
			DeliveredPackets             int64 `json:"delivered-packets"`
			OutputPackets                int64 `json:"output-packets"`
			MulticastSourceFilterMatches int64 `json:"multicast-source-filter-matches"`
		} `json:"udp"`
		IP struct {
			ReceivedPackets               int64 `json:"received-packets"`
			DroppedBadChecksum            int64 `json:"dropped-bad-checksum"`
			DroppedBelowMinimumSize       int64 `json:"dropped-below-minimum-size"`
			DroppedShortPackets           int64 `json:"dropped-short-packets"`
			DroppedTooLong                int64 `json:"dropped-too-long"`
			DroppedShortHeaderLength      int64 `json:"dropped-short-header-length"`
			DroppedShortData              int64 `json:"dropped-short-data"`
			DroppedBadOptions             int64 `json:"dropped-bad-options"`
			DroppedBadVersion             int64 `json:"dropped-bad-version"`
			ReceivedFragments             int64 `json:"received-fragments"`
			DroppedFragments              int64 `json:"dropped-fragments"`
			DroppedFragmentsAfterTimeout  int64 `json:"dropped-fragments-after-timeout"`
			ReassembledPackets            int64 `json:"reassembled-packets"`
			ReceivedLocalPackets          int64 `json:"received-local-packets"`
			DroppedUnknownProtocol        int64 `json:"dropped-unknown-protocol"`
			ForwardedPackets              int64 `json:"forwarded-packets"`
			FastForwardedPackets          int64 `json:"fast-forwarded-packets"`
			PacketsCannotForward          int64 `json:"packets-cannot-forward"`
			ReceivedUnknownMulticastGroup int64 `json:"received-unknown-multicast-group"`
			RedirectsSent                 int64 `json:"redirects-sent"`
			SentPackets                   int64 `json:"sent-packets"`
			SendPacketsFabricatedHeader   int64 `json:"send-packets-fabricated-header"`
			DiscardNoMbufs                int64 `json:"discard-no-mbufs"`
			DiscardNoRoute                int64 `json:"discard-no-route"`
			SentFragments                 int64 `json:"sent-fragments"`
			FragmentsCreated              int64 `json:"fragments-created"`
			DiscardCannotFragment         int64 `json:"discard-cannot-fragment"`
			DiscardTunnelNoGif            int64 `json:"discard-tunnel-no-gif"`
			DiscardBadAddress             int64 `json:"discard-bad-address"`
		} `json:"ip"`
		Icmp struct {
			IcmpCalls                   int64  `json:"icmp-calls"`
			ErrorsNotFromMessage        int64  `json:"errors-not-from-message"`
			DroppedBadCode              int64  `json:"dropped-bad-code"`
			DroppedTooShort             int64  `json:"dropped-too-short"`
			DroppedBadChecksum          int64  `json:"dropped-bad-checksum"`
			DroppedBadLength            int64  `json:"dropped-bad-length"`
			DroppedMulticastEcho        int64  `json:"dropped-multicast-echo"`
			DroppedMulticastTimestamp   int64  `json:"dropped-multicast-timestamp"`
			SentPackets                 int64  `json:"sent-packets"`
			DiscardInvalidReturnAddress int64  `json:"discard-invalid-return-address"`
			DiscardNoRoute              int64  `json:"discard-no-route"`
			IcmpAddressResponses        string `json:"icmp-address-responses"`
		} `json:"icmp"`
		Carp struct {
			ReceivedInetPackets      int64 `json:"received-inet-packets"`
			ReceivedInet6Packets     int64 `json:"received-inet6-packets"`
			DroppedWrongTTL          int64 `json:"dropped-wrong-ttl"`
			DroppedShortHeader       int64 `json:"dropped-short-header"`
			DroppedBadChecksum       int64 `json:"dropped-bad-checksum"`
			DroppedBadVersion        int64 `json:"dropped-bad-version"`
			DroppedShortPacket       int64 `json:"dropped-short-packet"`
			DroppedBadAuthentication int64 `json:"dropped-bad-authentication"`
			DroppedBadVhid           int64 `json:"dropped-bad-vhid"`
			DroppedBadAddressList    int64 `json:"dropped-bad-address-list"`
			SentInetPackets          int64 `json:"sent-inet-packets"`
			SentInet6Packets         int64 `json:"sent-inet6-packets"`
			SendFailedMemoryError    int64 `json:"send-failed-memory-error"`
		} `json:"carp"`
		Pfsync struct {
			ReceivedInetPackets  int64 `json:"received-inet-packets"`
			ReceivedInet6Packets int64 `json:"received-inet6-packets"`
			InputHistogram       []struct {
				Name  string `json:"name"`
				Count int64  `json:"count"`
			} `json:"input-histogram"`
			DroppedBadInterface int64 `json:"dropped-bad-interface"`
			DroppedBadTTL       int64 `json:"dropped-bad-ttl"`
			DroppedShortHeader  int64 `json:"dropped-short-header"`
			DroppedBadVersion   int64 `json:"dropped-bad-version"`
			DroppedBadAuth      int64 `json:"dropped-bad-auth"`
			DroppedBadAction    int64 `json:"dropped-bad-action"`
			DroppedShort        int64 `json:"dropped-short"`
			DroppedBadValues    int64 `json:"dropped-bad-values"`
			DroppedStaleState   int64 `json:"dropped-stale-state"`
			DroppedFailedLookup int64 `json:"dropped-failed-lookup"`
			SentInetPackets     int64 `json:"sent-inet-packets"`
			SendInet6Packets    int64 `json:"send-inet6-packets"`
			OutputHistogram     []struct {
				Name  string `json:"name"`
				Count int64  `json:"count"`
			} `json:"output-histogram"`
			DiscardedNoMemory int64 `json:"discarded-no-memory"`
			SendErrors        int64 `json:"send-errors"`
		} `json:"pfsync"`
		Arp struct {
			SentRequests            int64 `json:"sent-requests"`
			SentFailures            int64 `json:"sent-failures"`
			SentReplies             int64 `json:"sent-replies"`
			ReceivedRequests        int64 `json:"received-requests"`
			ReceivedReplies         int64 `json:"received-replies"`
			ReceivedPackets         int64 `json:"received-packets"`
			DroppedNoEntry          int64 `json:"dropped-no-entry"`
			EntriesTimeout          int64 `json:"entries-timeout"`
			DroppedDuplicateAddress int64 `json:"dropped-duplicate-address"`
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
		TCPSentPackets:      resp.Statistics.TCP.SentPackets,
		TCPReceivedPackets:  resp.Statistics.TCP.ReceivedPackets,
		ARPSentRequests:     resp.Statistics.Arp.SentRequests,
		ARPReceivedRequests: resp.Statistics.Arp.ReceivedRequests,
		TCPConnectionCountByState: map[string]int64{
			"CLOSED":      resp.Statistics.TCP.TCPConnectionCountByState.Closed,
			"LISTEN":      resp.Statistics.TCP.TCPConnectionCountByState.Listen,
			"SYN_SENT":    resp.Statistics.TCP.TCPConnectionCountByState.SynSent,
			"SYN_RCVD":    resp.Statistics.TCP.TCPConnectionCountByState.SynRcvd,
			"ESTABLISHED": resp.Statistics.TCP.TCPConnectionCountByState.Established,
			"CLOSE_WAIT":  resp.Statistics.TCP.TCPConnectionCountByState.CloseWait,
			"FIN_WAIT_1":  resp.Statistics.TCP.TCPConnectionCountByState.FinWait1,
			"CLOSING":     resp.Statistics.TCP.TCPConnectionCountByState.Closing,
			"LAST_ACK":    resp.Statistics.TCP.TCPConnectionCountByState.LastAck,
			"FIN_WAIT_2":  resp.Statistics.TCP.TCPConnectionCountByState.FinWait2,
			"TIME_WAIT":   resp.Statistics.TCP.TCPConnectionCountByState.TimeWait,
		},
		ICMPCalls:       resp.Statistics.Icmp.IcmpCalls,
		ICMPSentPackets: resp.Statistics.Icmp.SentPackets,
		ICMPDroppedByReason: map[string]int64{
			"BAD_CODE":            resp.Statistics.Icmp.DroppedBadCode,
			"TOO_SHORT":           resp.Statistics.Icmp.DroppedTooShort,
			"BAD_CHECKSUM":        resp.Statistics.Icmp.DroppedBadChecksum,
			"BAD_LENGTH":          resp.Statistics.Icmp.DroppedBadLength,
			"MULTICAST_ECHO":      resp.Statistics.Icmp.DroppedMulticastEcho,
			"MULTICAST_TIMESTAMP": resp.Statistics.Icmp.DroppedMulticastTimestamp,
		},
		UDPDeliveredPackets:  resp.Statistics.UDP.DeliveredPackets,
		UDPOutputPackets:     resp.Statistics.UDP.OutputPackets,
		UDPReceivedDatagrams: resp.Statistics.UDP.ReceivedDatagrams,
		UDPDroppedByReason: map[string]int64{
			"INCOMPLETE_HEADERS":  resp.Statistics.UDP.DroppedIncompleteHeaders,
			"BAD_DATA_LENGTH":     resp.Statistics.UDP.DroppedBadDataLength,
			"BAD_CHECKSUM":        resp.Statistics.UDP.DroppedBadChecksum,
			"NO_CHECKSUM":         resp.Statistics.UDP.DroppedNoChecksum,
			"NO_SOCKET":           resp.Statistics.UDP.DroppedNoSocket,
			"BROADCAST_MULTICAST": resp.Statistics.UDP.DroppedBroadcastMulticast,
			"FULL_SOCKET_BUFFER":  resp.Statistics.UDP.DroppedFullSocketBuffer,
		},
	}

	return out, nil
}
