```
# HELP audiochannels The current number of audiochannels on the bridge.
# TYPE audiochannels gauge
audiochannels{app="jitsi",jvb_instance="jvb"} 0
# HELP bit_rate_download download rate bit/s
# TYPE bit_rate_download gauge
bit_rate_download{app="jitsi",jvb_instance="jvb"} 0
# HELP bit_rate_upload Current upload rate in bit/s.
# TYPE bit_rate_upload gauge
bit_rate_upload{app="jitsi",jvb_instance="jvb"} 0
# HELP conference_sizes total number of packets sent
# TYPE conference_sizes histogram
conference_sizes_bucket{app="jitsi",jvb_instance="jvb",le="0"} 0
conference_sizes_bucket{app="jitsi",jvb_instance="jvb",le="1"} 0
conference_sizes_bucket{app="jitsi",jvb_instance="jvb",le="2"} 0
conference_sizes_bucket{app="jitsi",jvb_instance="jvb",le="3"} 0
conference_sizes_bucket{app="jitsi",jvb_instance="jvb",le="4"} 0
conference_sizes_bucket{app="jitsi",jvb_instance="jvb",le="5"} 0
conference_sizes_bucket{app="jitsi",jvb_instance="jvb",le="6"} 0
conference_sizes_bucket{app="jitsi",jvb_instance="jvb",le="7"} 0
conference_sizes_bucket{app="jitsi",jvb_instance="jvb",le="8"} 0
conference_sizes_bucket{app="jitsi",jvb_instance="jvb",le="9"} 0
conference_sizes_bucket{app="jitsi",jvb_instance="jvb",le="10"} 0
conference_sizes_bucket{app="jitsi",jvb_instance="jvb",le="11"} 0
conference_sizes_bucket{app="jitsi",jvb_instance="jvb",le="12"} 0
conference_sizes_bucket{app="jitsi",jvb_instance="jvb",le="13"} 0
conference_sizes_bucket{app="jitsi",jvb_instance="jvb",le="14"} 0
conference_sizes_bucket{app="jitsi",jvb_instance="jvb",le="15"} 0
conference_sizes_bucket{app="jitsi",jvb_instance="jvb",le="16"} 0
conference_sizes_bucket{app="jitsi",jvb_instance="jvb",le="17"} 0
conference_sizes_bucket{app="jitsi",jvb_instance="jvb",le="18"} 0
conference_sizes_bucket{app="jitsi",jvb_instance="jvb",le="19"} 0
conference_sizes_bucket{app="jitsi",jvb_instance="jvb",le="20"} 0
conference_sizes_bucket{app="jitsi",jvb_instance="jvb",le="+Inf"} 0
conference_sizes_sum{app="jitsi",jvb_instance="jvb"} 0
conference_sizes_count{app="jitsi",jvb_instance="jvb"} 0
# HELP conferences The current number of conferences hosted by the bridge
# TYPE conferences gauge
conferences{app="jitsi",jvb_instance="jvb"} 0
# HELP cpu_usage CPU usage for the machine. The value is between 0 and 1 and is the fraction of the last interval that the CPU spent in either user, nice, system or iowait state (what would appear in the 'cpu' line in 'top').
# TYPE cpu_usage gauge
cpu_usage{app="jitsi",jvb_instance="jvb"} 0.19958634953464321
# HELP jitter_aggregate Experimental. An average value (in milliseconds) of the jitter calculated for incoming and outgoing streams. This hasn't been tested and it is currently not known whether the values are correct or not.
# TYPE jitter_aggregate gauge
jitter_aggregate{app="jitsi",jvb_instance="jvb"} 0
# HELP largest_conference The current number of participants in the largest conference
# TYPE largest_conference gauge
largest_conference{app="jitsi",jvb_instance="jvb"} 0
# HELP loss_rate_download The fraction of lost incoming RTP packets. This is based on RTP sequence numbers and is relatively accurate.
# TYPE loss_rate_download gauge
loss_rate_download{app="jitsi",jvb_instance="jvb"} 0
# HELP loss_rate_upload The fraction of lost outgoing RTP packets. This is based on incoming RTCP Receiver Reports, and an attempt to subtract the fraction of packets that were not sent (i.e. were lost before they reached the bridge). Further, this is averaged over all streams of all users as opposed to all packets, so it is not correctly weighted. This is not accurate, but may be a useful metric nonetheless.
# TYPE loss_rate_upload gauge
loss_rate_upload{app="jitsi",jvb_instance="jvb"} 0
# HELP packet_rate_download download packet rate
# TYPE packet_rate_download gauge
packet_rate_download{app="jitsi",jvb_instance="jvb"} 0
# HELP packet_rate_upload Upload packets/s
# TYPE packet_rate_upload gauge
packet_rate_upload{app="jitsi",jvb_instance="jvb"} 0
# HELP participants The current number of participants.
# TYPE participants gauge
participants{app="jitsi",jvb_instance="jvb"} 0
# HELP process_cpu_seconds_total Total user and system CPU time spent in seconds.
# TYPE process_cpu_seconds_total counter
process_cpu_seconds_total 79
# HELP process_max_fds Maximum number of open file descriptors.
# TYPE process_max_fds gauge
process_max_fds 1.048576e+06
# HELP process_open_fds Number of open file descriptors.
# TYPE process_open_fds gauge
process_open_fds 10
# HELP process_resident_memory_bytes Resident memory size in bytes.
# TYPE process_resident_memory_bytes gauge
process_resident_memory_bytes 1.6175104e+07
# HELP process_start_time_seconds Start time of the process since unix epoch in seconds.
# TYPE process_start_time_seconds gauge
process_start_time_seconds 1.56720128188e+09
# HELP process_virtual_memory_bytes Virtual memory size in bytes.
# TYPE process_virtual_memory_bytes gauge
process_virtual_memory_bytes 1.16805632e+08
# HELP process_virtual_memory_max_bytes Maximum amount of virtual memory available in bytes.
# TYPE process_virtual_memory_max_bytes gauge
process_virtual_memory_max_bytes -1
# HELP promhttp_metric_handler_requests_in_flight Current number of scrapes being served.
# TYPE promhttp_metric_handler_requests_in_flight gauge
promhttp_metric_handler_requests_in_flight 1
# HELP promhttp_metric_handler_requests_total Total number of scrapes by HTTP status code.
# TYPE promhttp_metric_handler_requests_total counter
promhttp_metric_handler_requests_total{code="200"} 15042
promhttp_metric_handler_requests_total{code="500"} 0
promhttp_metric_handler_requests_total{code="503"} 0
# HELP rtt_aggregate An average value (in milliseconds) of the RTT across all streams.
# TYPE rtt_aggregate gauge
rtt_aggregate{app="jitsi",jvb_instance="jvb"} 0
# HELP threads The current number of threads.
# TYPE threads gauge
threads{app="jitsi",jvb_instance="jvb"} 43
# HELP total_bytes_received Total bytes received.
# TYPE total_bytes_received counter
total_bytes_received{app="jitsi",jvb_instance="jvb"} 2.99091834e+08
# HELP total_bytes_received_octo The total number octo bytes sent.
# TYPE total_bytes_received_octo counter
total_bytes_received_octo{app="jitsi",jvb_instance="jvb"} 0
# HELP total_bytes_sent The number of total bytes sent.
# TYPE total_bytes_sent counter
total_bytes_sent{app="jitsi",jvb_instance="jvb"} 2.98784502e+08
# HELP total_bytes_sent_octo Total octo bytes sent.
# TYPE total_bytes_sent_octo counter
total_bytes_sent_octo{app="jitsi",jvb_instance="jvb"} 0
# HELP total_channels Current number of channels
# TYPE total_channels gauge
total_channels{app="jitsi",jvb_instance="jvb"} 44
# HELP total_colibri_web_socket_messages_received The total number messages received through COLIBRI web sockets.
# TYPE total_colibri_web_socket_messages_received counter
total_colibri_web_socket_messages_received{app="jitsi",jvb_instance="jvb"} 0
# HELP total_colibri_web_socket_messages_sent The total number messages sent through COLIBRI web sockets.
# TYPE total_colibri_web_socket_messages_sent counter
total_colibri_web_socket_messages_sent{app="jitsi",jvb_instance="jvb"} 0
# HELP total_conference_seconds The sum of the lengths of all completed conferences, in seconds.
# TYPE total_conference_seconds counter
total_conference_seconds{app="jitsi",jvb_instance="jvb"} 2909
# HELP total_conferences_completed Total conferences completed.
# TYPE total_conferences_completed counter
total_conferences_completed{app="jitsi",jvb_instance="jvb"} 10
# HELP total_data_channel_messages_received Total data channel messages received.
# TYPE total_data_channel_messages_received counter
total_data_channel_messages_received{app="jitsi",jvb_instance="jvb"} 1766
# HELP total_data_channel_messages_sent The total number of data channel messages sent.
# TYPE total_data_channel_messages_sent counter
total_data_channel_messages_sent{app="jitsi",jvb_instance="jvb"} 1799
# HELP total_failed_conferences The total number of failed conferences on the bridge. A conference is marked as failed when all of its channels have failed. A channel is marked as failed if it had no payload activity.
# TYPE total_failed_conferences counter
total_failed_conferences{app="jitsi",jvb_instance="jvb"} 0
# HELP total_loss_controlled_participant_seconds The total number of participant-seconds that are loss-controlled.
# TYPE total_loss_controlled_participant_seconds counter
total_loss_controlled_participant_seconds{app="jitsi",jvb_instance="jvb"} 2403
# HELP total_loss_degraded_participant_seconds The total number of participant-seconds that are loss-degraded.
# TYPE total_loss_degraded_participant_seconds counter
total_loss_degraded_participant_seconds{app="jitsi",jvb_instance="jvb"} 285
# HELP total_loss_limited_participant_seconds The total number of participant-seconds that are loss-limited.
# TYPE total_loss_limited_participant_seconds counter
total_loss_limited_participant_seconds{app="jitsi",jvb_instance="jvb"} 23
# HELP total_memory The total memory of the machine in megabytes.
# TYPE total_memory gauge
total_memory{app="jitsi",jvb_instance="jvb"} 7842
# HELP total_no_payload_channels The current number of payload channels.
# TYPE total_no_payload_channels gauge
total_no_payload_channels{app="jitsi",jvb_instance="jvb"} 2
# HELP total_no_transport_channels The current number of transport channels.
# TYPE total_no_transport_channels gauge
total_no_transport_channels{app="jitsi",jvb_instance="jvb"} 2
# HELP total_packets_received Total number of packets received
# TYPE total_packets_received counter
total_packets_received{app="jitsi",jvb_instance="jvb"} 326635
# HELP total_packets_received_octo Total octo packets received.
# TYPE total_packets_received_octo counter
total_packets_received_octo{app="jitsi",jvb_instance="jvb"} 0
# HELP total_packets_sent The total number of packets sent.
# TYPE total_packets_sent counter
total_packets_sent{app="jitsi",jvb_instance="jvb"} 322400
# HELP total_packets_sent_octo total number of octo packets sent
# TYPE total_packets_sent_octo counter
total_packets_sent_octo{app="jitsi",jvb_instance="jvb"} 0
# HELP total_partially_failed_conferences The total number of partially failed conferences on the bridge. A conference is marked as partially failed when some of its channels has failed. A channel is marked as failed if it had no payload activity.
# TYPE total_partially_failed_conferences counter
total_partially_failed_conferences{app="jitsi",jvb_instance="jvb"} 1
# HELP total_tcp_connections number of open tcp connections
# TYPE total_tcp_connections gauge
total_tcp_connections{app="jitsi",jvb_instance="jvb"} 0
# HELP total_udp_connections The current number of udp connections.
# TYPE total_udp_connections gauge
total_udp_connections{app="jitsi",jvb_instance="jvb"} 21
# HELP used_memory Total used memory on the machine (i.e. what 'free' would return) in megabytes (10^6 B).
# TYPE used_memory gauge
used_memory{app="jitsi",jvb_instance="jvb"} 7559
# HELP videochannels The current number of videochannels.
# TYPE videochannels gauge
videochannels{app="jitsi",jvb_instance="jvb"} 0
# HELP videostreams An estimation of the number of current video streams forwarded by the bridge.
# TYPE videostreams gauge
videostreams{app="jitsi",jvb_instance="jvb"} 0

```