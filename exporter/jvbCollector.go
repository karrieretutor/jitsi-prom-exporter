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
func NewJvbCollector(namespace, subsystem string, retention time.Duration) *JvbCollector {
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

	//add metrics
	collector.metrics = append(collector.metrics, newMetric(collector.NamePrefix+"total_packets_sent", prometheus.CounterValue,
		"total number of packets sent", []string{"jvb_instance"}, constLabels))

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
						value, err := strconv.Atoi(stat.Value)
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
