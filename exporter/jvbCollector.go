/*
 *  Copyright 2019 karriere tutor GmbH
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *  	http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */

package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

//statsSet describes the stats belonging to a jvb instance
//lastUpdated: time of last update
//stats: as the are unmarshalled by the PresExtension
//jvbIdentifier: will be attached as tag to metric to identify individual JVBs
type statsSet struct {
	lastUpdated   time.Time
	stats         Stats
	jvbIdentifier string
}

type metric struct {
	name       string
	desc       *prometheus.Desc
	metricType prometheus.ValueType
}

func newMetric(name string, metricType prometheus.ValueType, help string,
	varLabels []string, constLabels prometheus.Labels) metric {

	var metric = metric{
		name:       name,
		metricType: metricType,
	}
	metric.desc = prometheus.NewDesc(name, help, varLabels, constLabels)
	return metric
}

//JvbCollector collects metrics for jitsi JVBs
//NamePrefix for naming the metrics, see https://godoc.org/github.com/prometheus/client_golang/prometheus#Opts
//Retention defines how long the jvb collector will consider a set of stats valid, once retention has passed since the last update,
//	the stats set will not be included in the collect output anymore
type JvbCollector struct {
	NamePrefix string
	Retention  time.Duration
	statsSets  []statsSet
	metrics    []metric
}

//NewJvbCollector initializes a Jvb collector
//namespace and subsystem may be empty if you dont need them, see https://godoc.org/github.com/prometheus/client_golang/prometheus#Opts
func NewJvbCollector(namespace, subsystem, labels string, retention time.Duration) *JvbCollector {
	var collector = &JvbCollector{
		Retention: retention,
	}

	var namePrefix = ""
	if subsystem != "" {
		namePrefix += subsystem
		namePrefix += "_"
	}

	if namespace != "" {
		namePrefix += namespace
		namePrefix += "_"
	}

	collector.NamePrefix = namePrefix

	var constLabels = prometheus.Labels{
		"app": "jitsi",
	}

	if labels != "" {
		s := strings.Split(labels, ",")
		for _, v := range s {
			d := strings.Split(v, "=")
			constLabels[d[0]] = d[1]
		}
	}

	//add metrics
	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"inactive_endpoints", prometheus.GaugeValue,
		"inactive endpoints", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"inactive_conferences", prometheus.GaugeValue,
		"inactive conferences", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"muc_clients_connected", prometheus.GaugeValue,
		"muc clients connected", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"p2p_conferences", prometheus.GaugeValue,
		"p2p conferences", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"total_aimd_bwe_expirations", prometheus.GaugeValue,
		"total aimd bwe expirations", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"octo_send_bitrate", prometheus.GaugeValue,
		"current outgoing bitrate on the octo channel (combined for all conferences) in bits per second.", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"octo_send_packet_rate", prometheus.GaugeValue,
		"current outgoing packet rate on the octo channel (combined for all conferences) in packets per second.", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"octo_receive_bitrate", prometheus.GaugeValue,
		"current incoming bitrate on the octo channel (combined for all conferences) in bits per second.", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"octo_receive_packet_rate", prometheus.GaugeValue,
		"current incoming packet rate on the octo channel (combined for all conferences) in packets per second.", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"total_dominant_speaker_changes", prometheus.CounterValue,
		"total dominant speaker changes", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"total_ice_succeeded", prometheus.CounterValue,
		"total ice succeeded", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"total_ice_succeeded_tcp", prometheus.CounterValue,
		"total ice succeeded tcp", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"total_ice_succeeded_relayed", prometheus.CounterValue,
		"total ice succeeded relayed", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"total_ice_failed", prometheus.CounterValue,
		"total ice failed", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"num_eps_no_msg_transport_after_delay", prometheus.GaugeValue,
		"num endpoints no message transport after delay", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"mucs_configured", prometheus.GaugeValue,
		"mucs configured", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"mucs_joined", prometheus.GaugeValue,
		"mucs joined", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"muc_clients_configured", prometheus.GaugeValue,
		"muc clients configured", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"incoming_loss", prometheus.GaugeValue,
		"incoming loss", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"outgoing_loss", prometheus.GaugeValue,
		"outgoing loss", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"overall_loss", prometheus.GaugeValue,
		"overall loss", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"conferences_by_video_senders", prometheus.UntypedValue,
		"histogram of conference sizes by number of video senders", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"conferences_by_audio_senders", prometheus.UntypedValue,
		"histogram of conference sizes by number of audio senders", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"stress_level", prometheus.GaugeValue,
		"stress level", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"average_participant_stress", prometheus.GaugeValue,
		"average participant stress", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"octo_endpoints", prometheus.GaugeValue,
		"The current number of octo endpoints (connected to remove jitsi-videobridge instances).", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"total_packets_dropped_octo", prometheus.CounterValue,
		"total number of packets dropped on the octo channel", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"dtls_failed_endpoints", prometheus.CounterValue,
		"dtls failed endpoints", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"receive_only_endpoints", prometheus.GaugeValue,
		"receive only endpoints", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"octo_conferences", prometheus.GaugeValue,
		"The current number of conferences in which octo is enabled.", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"num_eps_oversending", prometheus.GaugeValue,
		"current number of endpoints to which we are oversending", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"endpoints_with_high_outgoing_loss", prometheus.GaugeValue,
		"number of endpoints with high outgoing loss (more than 10% loss in the bridge->endpoint direction)", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"endpoints_with_spurious_remb", prometheus.GaugeValue,
		"total number of endpoints which have sent an RTCP REMB packet when REMB was not signaled", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"endpoints", prometheus.GaugeValue,
		"The current number of endpoints, including `octo` endpoints", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"local_endpoints", prometheus.GaugeValue,
		"The current number of local (non-`octo`) endpoints", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"local_active_endpoints", prometheus.GaugeValue,
		"The current number of local endpoints (not `octo`) which are in an active conference. This includes endpoints which are not sending audio or video, but are in an active conference (i.e. they are receive-only).", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"packet_rate_download", prometheus.GaugeValue,
		"current RTP incoming packet rate in packets per second", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"conference_sizes", prometheus.UntypedValue,
		"histogram of conference sizes (ie. how many conferences have 5 participants and so on)", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"total_packets_sent_octo", prometheus.CounterValue,
		"total number of octo packets sent", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"total_loss_degraded_participant_seconds", prometheus.CounterValue,
		"The total number of participant-seconds that are loss-degraded.", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"bit_rate_download", prometheus.GaugeValue,
		"The current incoming bitrate (RTP) in kilobits per second", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"jitter_aggregate", prometheus.GaugeValue,
		"Experimental. An average value (in milliseconds) of the jitter calculated for incoming and outgoing streams. This hasn't been tested and it is currently not known whether the values are correct or not.", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"total_packets_received", prometheus.CounterValue,
		"Total number of packets received", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"rtt_aggregate", prometheus.GaugeValue,
		"An average value (in milliseconds) of the RTT across all streams.", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"packet_rate_upload", prometheus.GaugeValue,
		"current RTP outgoing packet rate in packets per second", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"conferences", prometheus.GaugeValue,
		"The current number of conferences hosted by the bridge", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"participants", prometheus.GaugeValue,
		"The current number of participants. Deprecated.", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"total_loss_limited_participant_seconds", prometheus.CounterValue,
		"The total number of participant-seconds that are loss-limited.", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"preemptive_kfr_sent", prometheus.GaugeValue,
		"The total number of preemptive keyframe requests sent.", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"preemptive_kfr_suppressed", prometheus.GaugeValue,
		"preemptive kfr suppressed", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"largest_conference", prometheus.GaugeValue,
		"The current number of participants in the largest conference", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"total_packets_sent", prometheus.CounterValue,
		"The total number of packets sent.", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"total_data_channel_messages_sent", prometheus.CounterValue,
		"The total number of data channel messages sent.", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"total_bytes_received_octo", prometheus.CounterValue,
		"The total number octo bytes sent.", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"threads", prometheus.GaugeValue,
		"The current number of JVM threads", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"total_colibri_web_socket_messages_received", prometheus.CounterValue,
		"The total number messages received through COLIBRI web sockets.", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"videochannels", prometheus.GaugeValue,
		"The current number of videochannels.", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"total_packets_received_octo", prometheus.CounterValue,
		"total number packets received on the octo channel", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"total_colibri_web_socket_messages_sent", prometheus.CounterValue,
		"The total number messages sent through COLIBRI web sockets.", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"total_bytes_sent_octo", prometheus.CounterValue,
		"Total octo bytes sent.", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"total_data_channel_messages_received", prometheus.CounterValue,
		"Total data channel messages received.", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"total_conference_seconds", prometheus.CounterValue,
		"The sum of the lengths of all completed conferences, in seconds.", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"total_bytes_received", prometheus.CounterValue,
		"Total bytes received.", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"total_loss_controlled_participant_seconds", prometheus.CounterValue,
		"The total number of participant-seconds that are loss-controlled.", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"total_partially_failed_conferences", prometheus.CounterValue,
		"The total number of partially failed conferences on the bridge. A conference is marked as partially failed when some of its channels has failed. A channel is marked as failed if it had no payload activity.", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"bit_rate_upload", prometheus.GaugeValue,
		"Current upload rate in kbit/s.", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"total_conferences_completed", prometheus.CounterValue,
		"Total conferences completed.", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"total_bytes_sent", prometheus.CounterValue,
		"The number of total bytes sent.", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"total_failed_conferences", prometheus.CounterValue,
		"The total number of failed conferences on the bridge. A conference is marked as failed when all of its channels have failed. A channel is marked as failed if it had no payload activity.", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"endpoints_sending_audio", prometheus.GaugeValue,
		"The current number of sending audioendpoint on the bridge.", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"endpoints_sending_video", prometheus.GaugeValue,
		"The current number of sending videoendpoint on the bridge.", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"total_conferences_created", prometheus.GaugeValue,
		"The total number of created conferences.", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"total_participants", prometheus.GaugeValue,
		"The total number of participants.", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"total_keyframes_received", prometheus.CounterValue,
		"The total number of keyframes that were received (updated on endpoint expiration).", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"total_layering_changes_received", prometheus.CounterValue,
		"The total number of times the layering of an incoming video stream changed (updated on endpoint expiration).", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"total_video_stream_milliseconds_received", prometheus.CounterValue,
		"The total duration, in milliseconds, of video streams (SSRCs) that were received (updated on endpoint expiration).", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"total_relay_bytes_received", prometheus.CounterValue,
		"The total number of bytes received in RTP packets in relays in this conference (updated on endpoint expiration).", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"total_relay_bytes_sent", prometheus.CounterValue,
		"The total number of bytes sent in RTP packets in relays in this conference (updated on endpoint expiration).", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"total_relay_packets_received", prometheus.CounterValue,
		"The total number of RTP packets received in relays in this conference (updated on endpoint expiration).", []string{"jvb_instance"}, constLabels))

	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"total_relay_packets_sent", prometheus.CounterValue,
		"The total number of RTP packets sent in relays in this conference (updated on endpoint expiration).", []string{"jvb_instance"}, constLabels))

	return collector
}

//Describe implements prometheus.Collector interface
func (c *JvbCollector) Describe(desc chan<- *prometheus.Desc) {
	for _, m := range c.metrics {
		desc <- m.desc
	}
}

//Collect implements prometheus.Collector interface
func (c *JvbCollector) Collect(metrics chan<- prometheus.Metric) {
	for _, set := range c.statsSets {
		if time.Since(set.lastUpdated) <= c.Retention {

			//match metric names with stats
			for _, stat := range set.stats.Stats {
				for _, metric := range c.metrics {
					if metric.name == c.NamePrefix+stat.Name {

						//special case for conference_sizes
						if (metric.name == c.NamePrefix+"conference_sizes") ||
							(metric.name == c.NamePrefix+"conferences_by_video_senders") ||
							(metric.name == c.NamePrefix+"conferences_by_audio_senders") {
							conSizes, sum := conferenceSizesHelper(stat.Value)
							m, err := prometheus.NewConstHistogram(metric.desc, sum, float64(sum), conSizes, set.jvbIdentifier)

							if err != nil {
								fmt.Printf("Unable to publish metric %s: %s\n", metric.name, err.Error())
								continue
							}

							metrics <- m
							continue
						}

						//simple metrics
						value, err := strconv.ParseFloat(stat.Value, 64)
						if err != nil {
							fmt.Printf("unable to convert value %s to numeric: %s\n", stat.Value, err.Error())
							continue
						}
						m, err := prometheus.NewConstMetric(metric.desc, metric.metricType, float64(value), set.jvbIdentifier)
						if err != nil {
							fmt.Printf("Unable to create metric %s: %s\n", metric.name, err.Error())
							continue
						}
						metrics <- m
					}
				}
			}
		}
	}

	//conference_sizes_combined, a new metric where we sum up all valid conference sizes histograms
	var combinedConferenceSizes = make(map[float64]uint64)
	var combinedSum uint64
	for _, s := range c.statsSets {
		if time.Since(s.lastUpdated) <= c.Retention {
			for _, stat := range s.stats.Stats {
				if stat.Name == "conference_sizes" {
					conSizes, sum := conferenceSizesHelper(stat.Value)
					for bucket, numConferences := range conSizes {
						combinedConferenceSizes[bucket] += numConferences
					}
					combinedSum += sum
				}
			}
		}
	}

	metric := newMetric("conference_sizes_combined", prometheus.UntypedValue,
		"All active conference_sizes summed up into this histogram, see conference_sizes", []string{}, prometheus.Labels{})
	m, err := prometheus.NewConstHistogram(metric.desc, combinedSum, float64(combinedSum), combinedConferenceSizes)

	if err != nil {
		fmt.Printf("Unable to create %s metric: %s\n", metric.name, err.Error())
	} else {
		metrics <- m
	}

}

//Update updates the cached stats for the JVB identified by identifier, inserts a new stats set if none present yet.
//identifier: any string that identifies the specific JVB, you might want to consider using the node part of the JVB jid (<node>@<domain>/<resource>)
//	instead of the whole jid. This helps to keep track of JVBs being autoscaled
//stats: as they are unmarshalled by the PresExtension
func (c *JvbCollector) Update(identifier string, stats *Stats) {
	for i, s := range c.statsSets {
		if s.jvbIdentifier == identifier {
			c.statsSets[i].lastUpdated = time.Now()
			c.statsSets[i].stats = *stats
			return
		}
	}

	c.statsSets = append(c.statsSets, statsSet{
		lastUpdated:   time.Now(),
		stats:         *stats,
		jvbIdentifier: identifier,
	})
}

func conferenceSizesHelper(conferenceSizes string) (conferenceSizesHistogram map[float64]uint64, sum uint64) {
	var sizes = make(map[float64]uint64)
	value := strings.Trim(conferenceSizes, "[]")
	var values []uint64
	for _, v := range strings.Split(value, ",") {
		vuint, _ := strconv.ParseUint(v, 10, 64)
		values = append(values, vuint)
	}

	//calculate sum (makes this metric independent from conferences metric)
	sum = 0
	for _, v := range values {
		sum += v
	}

	//for the histgram buckets we need to omit the last field b/c the +inf bucket is added automatically
	values = values[:len(values)-1]

	//the bucket values have to be cumulative
	var i int
	for i = len(values) - 1; i >= 0; i-- {
		var cumulative uint64
		var j int
		for j = i; j >= 0; j-- {
			cumulative += values[j]
		}
		values[i] = cumulative
	}

	for i, v := range values {
		sizes[float64(i)] = v
	}

	return sizes, sum
}
