#!/bin/bash
# 2018-12-22 (cc) <paul4hough@gmail.com>
#
# debug
set -x

trap 'kill $(jobs -p)' SIGINT SIGTERM EXIT

# start exporters
node_exporter \
  --collector.systemd \
  2> node_exporter.out &
ne_pid=$!

process-exporter \
  -config.path process-exporter-config.yml \
  2> process-exporter.out &
pe_pid=$!

grok_exporter \
  -config grok-config.yml \
  > grok_exporter.out &
ge_pid=$!

export DATA_SOURCE_NAME='postgresql://postgres_exporter:pg_metrics@localhost:5432/postgres?sslmode=disable'
postgres_exporter \
  2> postgres_exporter.out &
pge_pid=$!

cc-auto-tie \
  2> cc-auto-tie.out &
cat_pid=$!

# rm -rf am-data
alertmanager \
  --storage.path="am-data/" \
  2> alertmanager.out &
am_pid=$!

# rm -rf prom-data
prometheus \
  --storage.tsdb.path="prom-data/" \
  2> prometheus.out &
pm_pid=$!

jobs -p

rm mock-logger.log
mock-logger/mock-logger
