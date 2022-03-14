package monitor

const MongoInstanceSummary = `
{
    "editable": false,
    "gnetId": null,
    "graphTooltip": 1,
    "id": null,
    "iteration": 1573818420013,
    "links": [],
    "panels": [
        {
            "collapsed": false,
            "datasource": null,
            "gridPos": {
                "h": 1,
                "w": 24,
                "x": 0,
                "y": 2
            },
            "id": 1009,
            "panels": [],
            "title": "",
            "type": "row"
        },
        {
            "cacheTimeout": null,
            "colorBackground": false,
            "colorPostfix": true,
            "colorValue": true,
            "colors": [
                "#d44a3a",
                "rgba(237, 129, 40, 0.89)",
                "#299c46"
            ],
            "datasource": null,
            "decimals": 2,
            "fieldConfig": {
                "defaults": {
                    "custom": {}
                },
                "overrides": []
            },
            "format": "s",
            "gauge": {
                "maxValue": 100,
                "minValue": 0,
                "show": false,
                "thresholdLabels": false,
                "thresholdMarkers": true
            },
            "gridPos": {
                "h": 2,
                "w": 4,
                "x": 0,
                "y": 3
            },
            "id": 1001,
            "interval": null,
            "links": [],
            "mappingType": 1,
            "mappingTypes": [
                {
                    "name": "value to text",
                    "value": 1
                },
                {
                    "name": "range to text",
                    "value": 2
                }
            ],
            "maxDataPoints": 100,
            "nullPointMode": "connected",
            "nullText": null,
            "postfix": "s",
            "postfixFontSize": "50%",
            "prefix": "",
            "prefixFontSize": "50%",
            "rangeMaps": [
                {
                    "from": "null",
                    "text": "N/A",
                    "to": "null"
                }
            ],
            "sparkline": {
                "fillColor": "rgba(31, 118, 189, 0.18)",
                "full": false,
                "lineColor": "rgb(31, 120, 193)",
                "show": true,
                "ymax": null,
                "ymin": null
            },
            "tableColumn": "",
            "targets": [
                {
                    "expr": "avg by (service_name) (mongodb_instance_uptime_seconds{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"})",
                    "interval": "$interval",
                    "refId": "A"
                }
            ],
            "thresholds": "3600,86400",
            "timeFrom": null,
            "timeShift": null,
            "title": "MongoDB Uptime",
            "type": "singlestat",
            "valueFontSize": "50%",
            "valueMaps": [
                {
                    "op": "=",
                    "text": "N/A",
                    "value": "null"
                }
            ],
            "valueName": "current"
        },
        {
            "cacheTimeout": null,
            "colorBackground": false,
            "colorValue": false,
            "colors": [
                "#299c46",
                "rgba(237, 129, 40, 0.89)",
                "#d44a3a"
            ],
            "datasource": null,
            "decimals": 2,
            "description": "",
            "fieldConfig": {
                "defaults": {
                    "custom": {}
                },
                "overrides": []
            },
            "format": "ops",
            "gauge": {
                "maxValue": 100,
                "minValue": 0,
                "show": false,
                "thresholdLabels": false,
                "thresholdMarkers": true
            },
            "gridPos": {
                "h": 2,
                "w": 4,
                "x": 4,
                "y": 3
            },
            "id": 1005,
            "interval": null,
            "links": [],
            "mappingType": 1,
            "mappingTypes": [
                {
                    "name": "value to text",
                    "value": 1
                },
                {
                    "name": "range to text",
                    "value": 2
                }
            ],
            "maxDataPoints": 100,
            "nullPointMode": "connected",
            "nullText": null,
            "postfix": "",
            "postfixFontSize": "50%",
            "prefix": "",
            "prefixFontSize": "50%",
            "rangeMaps": [
                {
                    "from": "null",
                    "text": "N/A",
                    "to": "null"
                }
            ],
            "sparkline": {
                "fillColor": "rgba(31, 118, 189, 0.18)",
                "full": false,
                "lineColor": "rgb(31, 120, 193)",
                "show": true,
                "ymax": null,
                "ymin": null
            },
            "tableColumn": "",
            "targets": [
                {
                    "expr": "sum(rate(mongodb_mongod_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type!=\"command\"}[$interval]) or irate(mongodb_mongod_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type!=\"command\"}[5m]) or rate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type!=\"command\"}[$interval]) or irate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type!=\"command\"}[5m]))",
                    "interval": "$interval",
                    "refId": "A"
                }
            ],
            "thresholds": "",
            "timeFrom": null,
            "timeShift": null,
            "title": "QPS",
            "type": "singlestat",
            "valueFontSize": "50%",
            "valueMaps": [
                {
                    "op": "=",
                    "text": "N/A",
                    "value": "null"
                }
            ],
            "valueName": "current"
        },
        {
            "cacheTimeout": null,
            "colorBackground": false,
            "colorValue": false,
            "colors": [
                "#299c46",
                "rgba(237, 129, 40, 0.89)",
                "#d44a3a"
            ],
            "datasource": null,
            "decimals": 2,
            "fieldConfig": {
                "defaults": {
                    "custom": {}
                },
                "overrides": []
            },
            "format": "\u00b5s",
            "gauge": {
                "maxValue": 100,
                "minValue": 0,
                "show": false,
                "thresholdLabels": false,
                "thresholdMarkers": true
            },
            "gridPos": {
                "h": 2,
                "w": 4,
                "x": 8,
                "y": 3
            },
            "id": 1007,
            "interval": null,
            "links": [],
            "mappingType": 1,
            "mappingTypes": [
                {
                    "name": "value to text",
                    "value": 1
                },
                {
                    "name": "range to text",
                    "value": 2
                }
            ],
            "maxDataPoints": 100,
            "nullPointMode": "connected",
            "nullText": null,
            "postfix": "",
            "postfixFontSize": "50%",
            "prefix": "",
            "prefixFontSize": "50%",
            "rangeMaps": [
                {
                    "from": "null",
                    "text": "N/A",
                    "to": "null"
                }
            ],
            "sparkline": {
                "fillColor": "rgba(31, 118, 189, 0.18)",
                "full": false,
                "lineColor": "rgb(31, 120, 193)",
                "show": true,
                "ymax": null,
                "ymin": null
            },
            "tableColumn": "",
            "targets": [
                {
                    "expr": "avg by (service_name) (rate(mongodb_mongod_op_latencies_latency_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"command\"}[$interval]) / (rate(mongodb_mongod_op_latencies_ops_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"command\"}[$interval]) > 0) or\nirate(mongodb_mongod_op_latencies_latency_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"command\"}[5m]) / (irate(mongodb_mongod_op_latencies_ops_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"command\"}[5m]) > 0))",
                    "interval": "$interval",
                    "refId": "A"
                }
            ],
            "thresholds": "",
            "timeFrom": null,
            "timeShift": null,
            "title": "Latency",
            "type": "singlestat",
            "valueFontSize": "50%",
            "valueMaps": [
                {
                    "op": "=",
                    "text": "N/A",
                    "value": "null"
                }
            ],
            "valueName": "current"
        },
        {
            "cacheTimeout": null,
            "colorBackground": false,
            "colorValue": false,
            "colors": [
                "#299c46",
                "rgba(237, 129, 40, 0.89)",
                "#d44a3a"
            ],
            "datasource": null,
            "fieldConfig": {
                "defaults": {
                    "custom": {}
                },
                "overrides": []
            },
            "format": "none",
            "gauge": {
                "maxValue": 100,
                "minValue": 0,
                "show": false,
                "thresholdLabels": false,
                "thresholdMarkers": true
            },
            "gridPos": {
                "h": 2,
                "w": 4,
                "x": 12,
                "y": 3
            },
            "id": 1020,
            "interval": null,
            "links": [
                {
                    "targetBlank": true,
                    "title": "MongoDB ReplSet Summary - $service_name",
                    "url": "/graph/d/mongodb-replicaset-summary/mongodb-replset-summary?var-service_name=$service_name"
                }
            ],
            "mappingType": 1,
            "mappingTypes": [
                {
                    "name": "value to text",
                    "value": 1
                },
                {
                    "name": "range to text",
                    "value": 2
                }
            ],
            "maxDataPoints": 100,
            "nullPointMode": "connected",
            "nullText": null,
            "postfix": "",
            "postfixFontSize": "50%",
            "prefix": "",
            "prefixFontSize": "50%",
            "rangeMaps": [
                {
                    "from": "null",
                    "text": "N/A",
                    "to": "null"
                }
            ],
            "sparkline": {
                "fillColor": "rgba(31, 118, 189, 0.18)",
                "full": false,
                "lineColor": "rgb(31, 120, 193)",
                "show": false,
                "ymax": null,
                "ymin": null
            },
            "tableColumn": "set",
            "targets": [
                {
                    "expr": "mongodb_mongod_replset_my_state{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}",
                    "format": "table",
                    "interval": "$interval",
                    "legendFormat": "{{set}}",
                    "refId": "A"
                }
            ],
            "thresholds": "",
            "timeFrom": null,
            "timeShift": null,
            "title": "ReplSet",
            "type": "singlestat",
            "valueFontSize": "50%",
            "valueMaps": [
                {
                    "op": "=",
                    "text": "N/A",
                    "value": "null"
                }
            ],
            "valueName": "current"
        },
        {
            "cacheTimeout": null,
            "colorBackground": false,
            "colorValue": false,
            "colors": [
                "#299c46",
                "rgba(237, 129, 40, 0.89)",
                "#d44a3a"
            ],
            "datasource": null,
            "description": "This shows the role of the selected service. Normally this should be one of PRIMARY, SECONDARY and ARBITER, but if the system is newly added it could show STARTUP2 during its initial sync.",
            "fieldConfig": {
                "defaults": {
                    "custom": {}
                },
                "overrides": []
            },
            "format": "none",
            "gauge": {
                "maxValue": 100,
                "minValue": 0,
                "show": false,
                "thresholdLabels": false,
                "thresholdMarkers": true
            },
            "gridPos": {
                "h": 2,
                "w": 8,
                "x": 16,
                "y": 3
            },
            "id": 1016,
            "interval": null,
            "links": [
                {
                    "targetBlank": true,
                    "title": "MongoDB ReplSet Summary - $service_name",
                    "url": "/graph/d/mongodb-replicaset-summary/mongodb-replset-summary?var-service_name=$service_name"
                }
            ],
            "mappingType": 1,
            "mappingTypes": [
                {
                    "name": "value to text",
                    "value": 1
                },
                {
                    "name": "range to text",
                    "value": 2
                }
            ],
            "maxDataPoints": 100,
            "nullPointMode": "connected",
            "nullText": null,
            "postfix": "",
            "postfixFontSize": "50%",
            "prefix": "",
            "prefixFontSize": "50%",
            "rangeMaps": [
                {
                    "from": "null",
                    "text": "N/A",
                    "to": "null"
                }
            ],
            "sparkline": {
                "fillColor": "rgba(31, 118, 189, 0.18)",
                "full": false,
                "lineColor": "rgb(31, 120, 193)",
                "show": false,
                "ymax": null,
                "ymin": null
            },
            "tableColumn": "",
            "targets": [
                {
                    "expr": "mongodb_mongod_replset_my_state{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}",
                    "interval": "$interval",
                    "refId": "A"
                }
            ],
            "thresholds": "",
            "timeFrom": null,
            "timeShift": null,
            "title": "Current ReplSet State",
            "type": "singlestat",
            "valueFontSize": "50%",
            "valueMaps": [
                {
                    "op": "=",
                    "text": "PRIMARY",
                    "value": "1"
                },
                {
                    "op": "=",
                    "text": "SECONDARY",
                    "value": "2"
                },
                {
                    "op": "=",
                    "text": "STARTUP",
                    "value": "0"
                },
                {
                    "op": "=",
                    "text": "RECOVERING",
                    "value": "3"
                },
                {
                    "op": "=",
                    "text": "STARTUP2",
                    "value": "5"
                },
                {
                    "op": "=",
                    "text": "UNKNOWN",
                    "value": "6"
                },
                {
                    "op": "=",
                    "text": "ARBITER",
                    "value": "7"
                },
                {
                    "op": "=",
                    "text": "DOWN",
                    "value": "8"
                },
                {
                    "op": "=",
                    "text": "ROLLBACK",
                    "value": "9"
                },
                {
                    "op": "=",
                    "text": "REMOVED",
                    "value": "10"
                }
            ],
            "valueName": "current"
        },
        {
            "backgroundColor": "rgba(128,128,128,0.1)",
            "colorMaps": [
                {
                    "color": "rgb(107, 152, 102)",
                    "text": "PRIMARY"
                },
                {
                    "color": "rgb(193, 159, 20)",
                    "text": "SECONDARY"
                }
            ],
            "crosshairColor": "#8F070C",
            "datasource": null,
            "description": "Read more about [Replica Set Member States](https://docs.mongodb.com/manual/reference/replica-states/).",
            "display": "timeline",
            "expandFromQueryS": 0,
            "extendLastValue": true,
            "fieldConfig": {
                "defaults": {
                    "custom": {}
                },
                "overrides": []
            },
            "gridPos": {
                "h": 4,
                "w": 24,
                "x": 0,
                "y": 5
            },
            "highlightOnMouseover": true,
            "id": 1018,
            "legendSortBy": "-ms",
            "lineColor": "rgba(0,0,0,0.1)",
            "metricNameColor": "#000000",
            "rangeMaps": [
                {
                    "from": "null",
                    "text": "N/A",
                    "to": "null"
                }
            ],
            "rowHeight": 50,
            "showDistinctCount": false,
            "showLegend": true,
            "showLegendCounts": false,
            "showLegendNames": true,
            "showLegendPercent": true,
            "showLegendTime": true,
            "showLegendValues": true,
            "showTimeAxis": true,
            "showTransitionCount": false,
            "targets": [
                {
                    "expr": "sum by (set) (mongodb_mongod_replset_my_state{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"})",
                    "interval": "$interval",
                    "legendFormat": "{{set}}",
                    "refId": "A"
                }
            ],
            "textSize": 24,
            "textSizeTime": 12,
            "timeFrom": null,
            "timeOptions": [
                {
                    "name": "Years",
                    "value": "years"
                },
                {
                    "name": "Months",
                    "value": "months"
                },
                {
                    "name": "Weeks",
                    "value": "weeks"
                },
                {
                    "name": "Days",
                    "value": "days"
                },
                {
                    "name": "Hours",
                    "value": "hours"
                },
                {
                    "name": "Minutes",
                    "value": "minutes"
                },
                {
                    "name": "Seconds",
                    "value": "seconds"
                },
                {
                    "name": "Milliseconds",
                    "value": "milliseconds"
                }
            ],
            "timePrecision": {
                "name": "Minutes",
                "value": "minutes"
            },
            "timeShift": null,
            "timeTextColor": "#d8d9da",
            "title": "ReplSet $service_name states",
            "type": "natel-discrete-panel",
            "units": "short",
            "useTimePrecision": false,
            "valueMaps": [
                {
                    "op": "=",
                    "text": "PRIMARY",
                    "value": "1"
                },
                {
                    "op": "=",
                    "text": "SECONDARY",
                    "value": "2"
                },
                {
                    "op": "=",
                    "text": "STARTUP",
                    "value": "0"
                },
                {
                    "op": "=",
                    "text": "RECOVERING",
                    "value": "3"
                },
                {
                    "op": "=",
                    "text": "STARTUP2",
                    "value": "5"
                },
                {
                    "op": "=",
                    "text": "UNKNOWN",
                    "value": "6"
                },
                {
                    "op": "=",
                    "text": "ARBITER",
                    "value": "7"
                },
                {
                    "op": "=",
                    "text": "DOWN",
                    "value": "8"
                },
                {
                    "op": "=",
                    "text": "ROLLBACK",
                    "value": "9"
                },
                {
                    "op": "=",
                    "text": "REMOVED",
                    "value": "10"
                }
            ],
            "valueTextColor": "#000000",
            "writeAllValues": false,
            "writeLastValue": true,
            "writeMetricNames": false
        },
        {
            "collapsed": false,
            "datasource": null,
            "gridPos": {
                "h": 1,
                "w": 24,
                "x": 0,
                "y": 10
            },
            "id": 1011,
            "panels": [],
            "title": "",
            "type": "row"
        },
        {
            "aliasColors": {},
            "bars": false,
            "dashLength": 10,
            "dashes": false,
            "datasource": null,
            "decimals": 2,
            "description": "Ops or Replicated Ops/sec classified by legacy wire protocol type (query, insert, update, delete, getmore). And (from the internal TTL threads) the docs deletes/sec by TTL indexes.",
            "editable": true,
            "error": false,
            "fieldConfig": {
                "defaults": {
                    "custom": {},
                    "links": []
                },
                "overrides": []
            },
            "fill": 2,
            "fillGradient": 0,
            "grid": {
                "leftLogBase": 1,
                "leftMax": null,
                "leftMin": 0,
                "rightLogBase": 1,
                "rightMax": null,
                "rightMin": null
            },
            "gridPos": {
                "h": 8,
                "w": 24,
                "x": 0,
                "y": 11
            },
            "hiddenSeries": false,
            "id": 15,
            "legend": {
                "alignAsTable": true,
                "avg": true,
                "current": false,
                "hideZero": false,
                "max": true,
                "min": true,
                "rightSide": true,
                "show": true,
                "sort": "avg",
                "sortDesc": true,
                "total": false,
                "values": true
            },
            "lines": true,
            "linewidth": 2,
            "links": [],
            "nullPointMode": "null",
            "paceLength": 10,
            "percentage": false,
            "pluginVersion": "7.1.3",
            "pointradius": 5,
            "points": false,
            "renderer": "flot",
            "seriesOverrides": [],
            "spaceLength": 10,
            "stack": false,
            "steppedLine": false,
            "targets": [
                {
                    "expr": "avg by (service_name) (rate(mongodb_mongod_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type!=\"command\"}[$interval]) or \nirate(mongodb_mongod_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type!=\"command\"}[5m]) or \nrate(mongodb_mongos_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type!=\"command\"}[$interval]) or \nirate(mongodb_mongos_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type!=\"command\"}[5m]))",
                    "format": "time_series",
                    "hide": false,
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "{{type}}",
                    "refId": "J",
                    "step": 300
                },
                {
                    "expr": "avg by (service_name) (rate(mongodb_mongod_op_counters_repl_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type!~\"(command|query|getmore)\"}[$interval]) or \nirate(mongodb_mongod_op_counters_repl_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type!~\"(command|query|getmore)\"}[5m]) or\nrate(mongodb_mongos_op_counters_repl_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type!~\"(command|query|getmore)\"}[$interval]) or \nirate(mongodb_mongos_op_counters_repl_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type!~\"(command|query|getmore)\"}[5m]))",
                    "format": "time_series",
                    "hide": false,
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "repl_{{type}}",
                    "refId": "A",
                    "step": 300
                },
                {
                    "expr": "avg by (service_name) (rate(mongodb_mongod_metrics_ttl_deleted_documents_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or \nirate(mongodb_mongod_metrics_ttl_deleted_documents_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]) or\nrate(mongodb_mongos_metrics_ttl_deleted_documents_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or \nirate(mongodb_mongos_metrics_ttl_deleted_documents_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]))",
                    "format": "time_series",
                    "hide": false,
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "ttl_delete",
                    "refId": "B",
                    "step": 300
                }
            ],
            "thresholds": [],
            "timeFrom": null,
            "timeRegions": [],
            "timeShift": null,
            "title": "Command Operations",
            "tooltip": {
                "msResolution": false,
                "shared": true,
                "sort": 5,
                "value_type": "individual"
            },
            "type": "graph",
            "x-axis": true,
            "xaxis": {
                "buckets": null,
                "mode": "time",
                "name": null,
                "show": true,
                "values": []
            },
            "y-axis": true,
            "y_formats": [
                "short",
                "short"
            ],
            "yaxes": [
                {
                    "decimals": 2,
                    "format": "ops",
                    "label": "",
                    "logBase": 1,
                    "max": null,
                    "min": 0,
                    "show": true
                },
                {
                    "format": "short",
                    "logBase": 1,
                    "max": null,
                    "min": null,
                    "show": false
                }
            ],
            "yaxis": {
                "align": false,
                "alignLevel": null
            }
        },
        {
            "aliasColors": {},
            "bars": false,
            "dashLength": 10,
            "dashes": false,
            "datasource": null,
            "decimals": 2,
            "description": "Average latency of operations (classified by read, write, or (other) command)",
            "fieldConfig": {
                "defaults": {
                    "custom": {},
                    "links": []
                },
                "overrides": []
            },
            "fill": 2,
            "fillGradient": 0,
            "gridPos": {
                "h": 8,
                "w": 24,
                "x": 0,
                "y": 19
            },
            "hiddenSeries": false,
            "id": 1014,
            "legend": {
                "alignAsTable": true,
                "avg": true,
                "current": false,
                "max": true,
                "min": true,
                "rightSide": true,
                "show": true,
                "sort": "avg",
                "sortDesc": true,
                "total": false,
                "values": true
            },
            "lines": true,
            "linewidth": 2,
            "nullPointMode": "null",
            "percentage": false,
            "pluginVersion": "7.1.3",
            "pointradius": 2,
            "points": false,
            "renderer": "flot",
            "seriesOverrides": [],
            "spaceLength": 10,
            "stack": false,
            "steppedLine": false,
            "targets": [
                {
                    "expr": "avg by (service_name,type) (rate(mongodb_mongod_op_latencies_latency_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) / (rate(mongodb_mongod_op_latencies_ops_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) > 0) or irate(mongodb_mongod_op_latencies_latency_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]) / (irate(mongodb_mongod_op_latencies_ops_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]) > 0))",
                    "interval": "$interval",
                    "legendFormat": "{{type}}",
                    "refId": "A"
                }
            ],
            "thresholds": [],
            "timeFrom": null,
            "timeRegions": [],
            "timeShift": null,
            "title": "Latency Detail",
            "tooltip": {
                "shared": true,
                "sort": 5,
                "value_type": "individual"
            },
            "type": "graph",
            "xaxis": {
                "buckets": null,
                "mode": "time",
                "name": null,
                "show": true,
                "values": []
            },
            "yaxes": [
                {
                    "decimals": 2,
                    "format": "\u00b5s",
                    "label": null,
                    "logBase": 1,
                    "max": null,
                    "min": "0",
                    "show": true
                },
                {
                    "format": "short",
                    "label": null,
                    "logBase": 1,
                    "max": null,
                    "min": null,
                    "show": false
                }
            ],
            "yaxis": {
                "align": false,
                "alignLevel": null
            }
        },
        {
            "aliasColors": {},
            "bars": false,
            "dashLength": 10,
            "dashes": false,
            "datasource": null,
            "decimals": 2,
            "description": "TCP connections (Incoming)",
            "editable": true,
            "error": false,
            "fieldConfig": {
                "defaults": {
                    "custom": {},
                    "links": []
                },
                "overrides": []
            },
            "fill": 2,
            "fillGradient": 0,
            "grid": {
                "leftLogBase": 1,
                "leftMax": null,
                "leftMin": 0,
                "rightLogBase": 1,
                "rightMax": null,
                "rightMin": null
            },
            "gridPos": {
                "h": 8,
                "w": 12,
                "x": 0,
                "y": 27
            },
            "height": "250px",
            "hiddenSeries": false,
            "id": 38,
            "legend": {
                "alignAsTable": true,
                "avg": true,
                "current": false,
                "max": true,
                "min": true,
                "rightSide": false,
                "show": true,
                "sort": "avg",
                "sortDesc": true,
                "total": false,
                "values": true
            },
            "lines": true,
            "linewidth": 2,
            "links": [],
            "nullPointMode": "null",
            "paceLength": 10,
            "percentage": false,
            "pluginVersion": "7.1.3",
            "pointradius": 5,
            "points": false,
            "renderer": "flot",
            "seriesOverrides": [],
            "spaceLength": 10,
            "stack": false,
            "steppedLine": false,
            "targets": [
                {
                    "expr": "avg by (service_name) (mongodb_mongod_connections{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"current\"} or\nmongodb_mongos_connections{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"current\"} or\nmongodb_connections{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"current\"})",
                    "format": "time_series",
                    "hide": false,
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "Connections",
                    "refId": "J",
                    "step": 300
                }
            ],
            "thresholds": [],
            "timeFrom": null,
            "timeRegions": [],
            "timeShift": null,
            "title": "Connections",
            "tooltip": {
                "msResolution": false,
                "shared": true,
                "sort": 5,
                "value_type": "cumulative"
            },
            "type": "graph",
            "x-axis": true,
            "xaxis": {
                "buckets": null,
                "mode": "time",
                "name": null,
                "show": true,
                "values": []
            },
            "y-axis": true,
            "y_formats": [
                "short",
                "short"
            ],
            "yaxes": [
                {
                    "decimals": 2,
                    "format": "short",
                    "label": "",
                    "logBase": 1,
                    "max": null,
                    "min": 0,
                    "show": true
                },
                {
                    "format": "short",
                    "logBase": 1,
                    "max": null,
                    "min": null,
                    "show": false
                }
            ],
            "yaxis": {
                "align": false,
                "alignLevel": null
            }
        },
        {
            "aliasColors": {},
            "bars": false,
            "dashLength": 10,
            "dashes": false,
            "datasource": null,
            "decimals": 2,
            "description": "Open cursors. Includes idle cursors.",
            "editable": true,
            "error": false,
            "fieldConfig": {
                "defaults": {
                    "custom": {},
                    "links": []
                },
                "overrides": []
            },
            "fill": 2,
            "fillGradient": 0,
            "grid": {
                "leftLogBase": 1,
                "leftMax": null,
                "leftMin": 0,
                "rightLogBase": 1,
                "rightMax": null,
                "rightMin": null
            },
            "gridPos": {
                "h": 8,
                "w": 12,
                "x": 12,
                "y": 27
            },
            "height": "250px",
            "hiddenSeries": false,
            "id": 25,
            "legend": {
                "alignAsTable": true,
                "avg": true,
                "current": false,
                "hideZero": false,
                "max": true,
                "min": true,
                "rightSide": false,
                "show": true,
                "sort": "avg",
                "sortDesc": true,
                "total": false,
                "values": true
            },
            "lines": true,
            "linewidth": 2,
            "links": [],
            "nullPointMode": "null",
            "paceLength": 10,
            "percentage": false,
            "pluginVersion": "7.1.3",
            "pointradius": 5,
            "points": false,
            "renderer": "flot",
            "seriesOverrides": [],
            "spaceLength": 10,
            "stack": false,
            "steppedLine": false,
            "targets": [
                {
                    "expr": "avg by (service_name,state) (mongodb_mongod_metrics_cursor_open{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"} or\nmongodb_mongod_cursors{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"} or\nmongodb_mongos_metrics_cursor_open{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"} or \nmongodb_mongos_cursors{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"})",
                    "format": "time_series",
                    "hide": false,
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "{{state}}",
                    "refId": "J",
                    "step": 300
                }
            ],
            "thresholds": [],
            "timeFrom": null,
            "timeRegions": [],
            "timeShift": null,
            "title": "Cursors",
            "tooltip": {
                "msResolution": false,
                "shared": true,
                "sort": 5,
                "value_type": "cumulative"
            },
            "type": "graph",
            "x-axis": true,
            "xaxis": {
                "buckets": null,
                "mode": "time",
                "name": null,
                "show": true,
                "values": []
            },
            "y-axis": true,
            "y_formats": [
                "short",
                "short"
            ],
            "yaxes": [
                {
                    "decimals": 2,
                    "format": "short",
                    "label": "",
                    "logBase": 1,
                    "max": null,
                    "min": 0,
                    "show": true
                },
                {
                    "format": "short",
                    "logBase": 1,
                    "max": null,
                    "min": null,
                    "show": false
                }
            ],
            "yaxis": {
                "align": false,
                "alignLevel": null
            }
        },
        {
            "aliasColors": {},
            "bars": false,
            "dashLength": 10,
            "dashes": false,
            "datasource": null,
            "decimals": 2,
            "description": "Docs per second inserted, updated, deleted or returned. (N.b. not 1-to-1 with operation counts.)",
            "editable": true,
            "error": false,
            "fieldConfig": {
                "defaults": {
                    "custom": {},
                    "links": []
                },
                "overrides": []
            },
            "fill": 2,
            "fillGradient": 0,
            "grid": {
                "leftLogBase": 1,
                "leftMax": null,
                "leftMin": 0,
                "rightLogBase": 1,
                "rightMax": null,
                "rightMin": null
            },
            "gridPos": {
                "h": 8,
                "w": 12,
                "x": 0,
                "y": 35
            },
            "height": "250px",
            "hiddenSeries": false,
            "id": 36,
            "legend": {
                "alignAsTable": true,
                "avg": true,
                "current": false,
                "hideZero": false,
                "max": true,
                "min": true,
                "show": true,
                "sort": "avg",
                "sortDesc": true,
                "total": false,
                "values": true
            },
            "lines": true,
            "linewidth": 2,
            "links": [],
            "nullPointMode": "null",
            "paceLength": 10,
            "percentage": false,
            "pluginVersion": "7.1.3",
            "pointradius": 5,
            "points": false,
            "renderer": "flot",
            "seriesOverrides": [],
            "spaceLength": 10,
            "stack": false,
            "steppedLine": false,
            "targets": [
                {
                    "expr": "avg by (service_name,state) (rate(mongodb_mongod_metrics_document_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or \nirate(mongodb_mongod_metrics_document_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]))",
                    "format": "time_series",
                    "hide": false,
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "{{state}}",
                    "refId": "J",
                    "step": 300
                }
            ],
            "thresholds": [],
            "timeFrom": null,
            "timeRegions": [],
            "timeShift": null,
            "title": "Document Operations",
            "tooltip": {
                "msResolution": false,
                "shared": true,
                "sort": 5,
                "value_type": "cumulative"
            },
            "type": "graph",
            "x-axis": true,
            "xaxis": {
                "buckets": null,
                "mode": "time",
                "name": null,
                "show": true,
                "values": []
            },
            "y-axis": true,
            "y_formats": [
                "short",
                "short"
            ],
            "yaxes": [
                {
                    "decimals": 2,
                    "format": "ops",
                    "label": "",
                    "logBase": 1,
                    "max": null,
                    "min": 0,
                    "show": true
                },
                {
                    "format": "short",
                    "logBase": 1,
                    "max": null,
                    "min": null,
                    "show": false
                }
            ],
            "yaxis": {
                "align": false,
                "alignLevel": null
            }
        },
        {
            "aliasColors": {},
            "bars": false,
            "dashLength": 10,
            "dashes": false,
            "datasource": null,
            "decimals": 2,
            "description": "Operations queued due to a lock.",
            "editable": true,
            "error": false,
            "fieldConfig": {
                "defaults": {
                    "custom": {},
                    "links": []
                },
                "overrides": []
            },
            "fill": 2,
            "fillGradient": 0,
            "grid": {
                "leftLogBase": 1,
                "leftMax": null,
                "leftMin": 0,
                "rightLogBase": 1,
                "rightMax": null,
                "rightMin": null
            },
            "gridPos": {
                "h": 8,
                "w": 12,
                "x": 12,
                "y": 35
            },
            "height": "250px",
            "hiddenSeries": false,
            "id": 40,
            "legend": {
                "alignAsTable": true,
                "avg": true,
                "current": false,
                "hideZero": false,
                "max": true,
                "min": true,
                "show": true,
                "sort": "avg",
                "sortDesc": true,
                "total": false,
                "values": true
            },
            "lines": true,
            "linewidth": 2,
            "links": [],
            "nullPointMode": "null",
            "paceLength": 10,
            "percentage": false,
            "pluginVersion": "7.1.3",
            "pointradius": 5,
            "points": false,
            "renderer": "flot",
            "seriesOverrides": [],
            "spaceLength": 10,
            "stack": false,
            "steppedLine": false,
            "targets": [
                {
                    "expr": "avg by (service_name,type) (mongodb_mongod_global_lock_current_queue{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"})",
                    "format": "time_series",
                    "hide": false,
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "{{type}}",
                    "refId": "J",
                    "step": 300
                }
            ],
            "thresholds": [],
            "timeFrom": null,
            "timeRegions": [],
            "timeShift": null,
            "title": "Queued Operations",
            "tooltip": {
                "msResolution": false,
                "shared": true,
                "sort": 5,
                "value_type": "cumulative"
            },
            "type": "graph",
            "x-axis": true,
            "xaxis": {
                "buckets": null,
                "mode": "time",
                "name": null,
                "show": true,
                "values": []
            },
            "y-axis": true,
            "y_formats": [
                "short",
                "short"
            ],
            "yaxes": [
                {
                    "decimals": 2,
                    "format": "ops",
                    "label": "",
                    "logBase": 1,
                    "max": null,
                    "min": 0,
                    "show": true
                },
                {
                    "format": "short",
                    "logBase": 1,
                    "max": null,
                    "min": null,
                    "show": false
                }
            ],
            "yaxis": {
                "align": false,
                "alignLevel": null
            }
        },
        {
            "aliasColors": {},
            "bars": false,
            "dashLength": 10,
            "dashes": false,
            "datasource": null,
            "decimals": 2,
            "description": "Ratio of Documents returned or Index entries scanned / full documents scanned",
            "editable": true,
            "error": false,
            "fieldConfig": {
                "defaults": {
                    "custom": {},
                    "links": []
                },
                "overrides": []
            },
            "fill": 2,
            "fillGradient": 0,
            "grid": {
                "leftLogBase": 1,
                "leftMax": null,
                "leftMin": 0,
                "rightLogBase": 1,
                "rightMax": null,
                "rightMin": null
            },
            "gridPos": {
                "h": 8,
                "w": 12,
                "x": 0,
                "y": 43
            },
            "height": "250px",
            "hiddenSeries": false,
            "id": 63,
            "legend": {
                "alignAsTable": true,
                "avg": true,
                "current": false,
                "max": true,
                "min": true,
                "show": true,
                "sort": "avg",
                "sortDesc": true,
                "total": false,
                "values": true
            },
            "lines": true,
            "linewidth": 2,
            "links": [],
            "nullPointMode": "null as zero",
            "paceLength": 10,
            "percentage": true,
            "pluginVersion": "7.1.3",
            "pointradius": 5,
            "points": false,
            "renderer": "flot",
            "seriesOverrides": [],
            "spaceLength": 10,
            "stack": false,
            "steppedLine": false,
            "targets": [
                {
                    "expr": "sum(rate(mongodb_mongod_metrics_document_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"returned\"}[$interval]))/\nsum(rate(mongodb_mongod_metrics_query_executor_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"scanned_objects\"}[$interval])) \nor\nsum(irate(mongodb_mongod_metrics_document_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"returned\"}[5m]))/\nsum(irate(mongodb_mongod_metrics_query_executor_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"scanned_objects\"}[5m]))",
                    "format": "time_series",
                    "hide": false,
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "Document",
                    "refId": "J",
                    "step": 300
                },
                {
                    "expr": "(sum(rate(mongodb_mongod_metrics_query_executor_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"scanned\"}[$interval]))/\nsum(rate(mongodb_mongod_metrics_query_executor_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"scanned_objects\"}[$interval])) \nor\nsum(irate(mongodb_mongod_metrics_query_executor_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"scanned\"}[5m]))/\nsum(irate(mongodb_mongod_metrics_query_executor_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"scanned_objects\"}[5m])))",
                    "format": "time_series",
                    "hide": false,
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "Index",
                    "refId": "A",
                    "step": 300
                }
            ],
            "thresholds": [],
            "timeFrom": null,
            "timeRegions": [],
            "timeShift": null,
            "title": "Query Efficiency",
            "tooltip": {
                "msResolution": false,
                "shared": true,
                "sort": 5,
                "value_type": "cumulative"
            },
            "type": "graph",
            "x-axis": true,
            "xaxis": {
                "buckets": null,
                "mode": "time",
                "name": null,
                "show": true,
                "values": []
            },
            "y-axis": true,
            "y_formats": [
                "short",
                "short"
            ],
            "yaxes": [
                {
                    "decimals": null,
                    "format": "percentunit",
                    "label": "",
                    "logBase": 1,
                    "max": null,
                    "min": 0,
                    "show": true
                },
                {
                    "format": "none",
                    "logBase": 1,
                    "max": null,
                    "min": null,
                    "show": false
                }
            ],
            "yaxis": {
                "align": false,
                "alignLevel": null
            }
        },
        {
            "aliasColors": {},
            "bars": false,
            "dashLength": 10,
            "dashes": false,
            "datasource": null,
            "decimals": 2,
            "description": "This panel shows the number of objects (both data (scanned_objects) and index (scanned)) as well as the number of documents that were moved to a new location due to the size of the document growing. Moved documents only apply to the MMAPv1 storage engine.",
            "editable": true,
            "error": false,
            "fieldConfig": {
                "defaults": {
                    "custom": {},
                    "links": []
                },
                "overrides": []
            },
            "fill": 2,
            "fillGradient": 0,
            "grid": {},
            "gridPos": {
                "h": 8,
                "w": 12,
                "x": 12,
                "y": 43
            },
            "hiddenSeries": false,
            "id": 64,
            "legend": {
                "alignAsTable": true,
                "avg": true,
                "current": false,
                "hideZero": false,
                "max": true,
                "min": true,
                "rightSide": false,
                "show": true,
                "sort": "avg",
                "sortDesc": true,
                "total": false,
                "values": true
            },
            "lines": true,
            "linewidth": 2,
            "links": [],
            "nullPointMode": "null as zero",
            "paceLength": 10,
            "percentage": false,
            "pluginVersion": "7.1.3",
            "pointradius": 5,
            "points": false,
            "renderer": "flot",
            "seriesOverrides": [],
            "spaceLength": 10,
            "stack": false,
            "steppedLine": false,
            "targets": [
                {
                    "expr": "avg by (service_name) (rate(mongodb_mongod_metrics_query_executor_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or \nirate(mongodb_mongod_metrics_query_executor_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]))",
                    "format": "time_series",
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "{{state}}",
                    "metric": "",
                    "refId": "A",
                    "step": 300
                },
                {
                    "expr": "avg by (service_name) (rate(mongodb_mongod_metrics_record_moves_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or irate(mongodb_mongod_metrics_record_moves_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]))",
                    "format": "time_series",
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "moved",
                    "refId": "B",
                    "step": 300
                }
            ],
            "thresholds": [],
            "timeFrom": null,
            "timeRegions": [],
            "timeShift": null,
            "title": "Scanned and Moved Objects",
            "tooltip": {
                "msResolution": false,
                "shared": true,
                "sort": 5,
                "value_type": "cumulative"
            },
            "type": "graph",
            "xaxis": {
                "buckets": null,
                "mode": "time",
                "name": null,
                "show": true,
                "values": []
            },
            "yaxes": [
                {
                    "decimals": 2,
                    "format": "ops",
                    "label": "",
                    "logBase": 1,
                    "max": null,
                    "min": 0,
                    "show": true
                },
                {
                    "format": "short",
                    "logBase": 1,
                    "max": null,
                    "min": null,
                    "show": false
                }
            ],
            "yaxis": {
                "align": false,
                "alignLevel": null
            }
        },
        {
            "aliasColors": {},
            "bars": false,
            "dashLength": 10,
            "dashes": false,
            "datasource": null,
            "decimals": 2,
            "description": "Legacy driver operation: Number of, and Sum of time spent, per second executing getLastError commands to confirm write concern.",
            "editable": true,
            "error": false,
            "fieldConfig": {
                "defaults": {
                    "custom": {},
                    "links": []
                },
                "overrides": []
            },
            "fill": 2,
            "fillGradient": 0,
            "grid": {
                "leftLogBase": 1,
                "leftMax": null,
                "leftMin": 0,
                "rightLogBase": 1,
                "rightMax": null,
                "rightMin": null
            },
            "gridPos": {
                "h": 8,
                "w": 12,
                "x": 0,
                "y": 51
            },
            "hiddenSeries": false,
            "id": 41,
            "legend": {
                "alignAsTable": true,
                "avg": true,
                "current": false,
                "max": true,
                "min": true,
                "show": true,
                "sort": "avg",
                "sortDesc": true,
                "total": false,
                "values": true
            },
            "lines": true,
            "linewidth": 2,
            "links": [],
            "nullPointMode": "null",
            "paceLength": 10,
            "percentage": false,
            "pluginVersion": "7.1.3",
            "pointradius": 5,
            "points": false,
            "renderer": "flot",
            "seriesOverrides": [],
            "spaceLength": 10,
            "stack": false,
            "steppedLine": false,
            "targets": [
                {
                    "expr": "avg by (service_name) (rate(mongodb_mongod_metrics_get_last_error_wtime_total_milliseconds{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or \nirate(mongodb_mongod_metrics_get_last_error_wtime_total_milliseconds{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]) or\nrate(mongodb_mongos_metrics_get_last_error_wtime_total_milliseconds{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or \nirate(mongodb_mongos_metrics_get_last_error_wtime_total_milliseconds{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]))",
                    "format": "time_series",
                    "hide": false,
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "Write Wait Time",
                    "refId": "J",
                    "step": 300
                }
            ],
            "thresholds": [],
            "timeFrom": null,
            "timeRegions": [],
            "timeShift": null,
            "title": "getLastError Write Time",
            "tooltip": {
                "msResolution": false,
                "shared": true,
                "sort": 5,
                "value_type": "cumulative"
            },
            "type": "graph",
            "x-axis": true,
            "xaxis": {
                "buckets": null,
                "mode": "time",
                "name": null,
                "show": true,
                "values": []
            },
            "y-axis": true,
            "y_formats": [
                "short",
                "short"
            ],
            "yaxes": [
                {
                    "decimals": 2,
                    "format": "ms",
                    "label": "",
                    "logBase": 1,
                    "max": null,
                    "min": 0,
                    "show": true
                },
                {
                    "format": "short",
                    "logBase": 1,
                    "max": null,
                    "min": null,
                    "show": false
                }
            ],
            "yaxis": {
                "align": false,
                "alignLevel": null
            }
        },
        {
            "aliasColors": {},
            "bars": false,
            "dashLength": 10,
            "dashes": false,
            "datasource": null,
            "decimals": 2,
            "description": "Legacy driver operation: Number of getLastError commands that timed out trying to confirm write concern.",
            "editable": true,
            "error": false,
            "fieldConfig": {
                "defaults": {
                    "custom": {},
                    "links": []
                },
                "overrides": []
            },
            "fill": 2,
            "fillGradient": 0,
            "grid": {
                "leftLogBase": 1,
                "leftMax": null,
                "leftMin": 0,
                "rightLogBase": 1,
                "rightMax": null,
                "rightMin": null
            },
            "gridPos": {
                "h": 8,
                "w": 12,
                "x": 12,
                "y": 51
            },
            "height": "250px",
            "hiddenSeries": false,
            "id": 62,
            "legend": {
                "alignAsTable": true,
                "avg": true,
                "current": false,
                "hideZero": false,
                "max": true,
                "min": true,
                "show": true,
                "sort": "avg",
                "sortDesc": true,
                "total": false,
                "values": true
            },
            "lines": true,
            "linewidth": 2,
            "links": [],
            "nullPointMode": "null",
            "paceLength": 10,
            "percentage": false,
            "pluginVersion": "7.1.3",
            "pointradius": 5,
            "points": false,
            "renderer": "flot",
            "seriesOverrides": [
                {
                    "alias": "Total",
                    "color": "#C4162A",
                    "fill": 0
                }
            ],
            "spaceLength": 10,
            "stack": false,
            "steppedLine": false,
            "targets": [
                {
                    "expr": "avg by (service_name) (rate(mongodb_mongod_metrics_get_last_error_wtime_num_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or \nirate(mongodb_mongod_metrics_get_last_error_wtime_num_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]) or\nrate(mongodb_mongos_metrics_get_last_error_wtime_num_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or \nirate(mongodb_mongos_metrics_get_last_error_wtime_num_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]))",
                    "format": "time_series",
                    "hide": false,
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "Total",
                    "refId": "J",
                    "step": 300
                },
                {
                    "expr": "avg by (service_name) (rate(mongodb_mongod_metrics_get_last_error_wtimeouts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or\nirate(mongodb_mongod_metrics_get_last_error_wtimeouts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]) or\nrate(mongodb_mongos_metrics_get_last_error_wtimeouts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or\nirate(mongodb_mongos_metrics_get_last_error_wtimeouts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]))",
                    "format": "time_series",
                    "hide": false,
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "Timeouts",
                    "refId": "A",
                    "step": 300
                }
            ],
            "thresholds": [],
            "timeFrom": null,
            "timeRegions": [],
            "timeShift": null,
            "title": "getLastError Write Operations",
            "tooltip": {
                "msResolution": false,
                "shared": true,
                "sort": 5,
                "value_type": "cumulative"
            },
            "type": "graph",
            "x-axis": true,
            "xaxis": {
                "buckets": null,
                "mode": "time",
                "name": null,
                "show": true,
                "values": []
            },
            "y-axis": true,
            "y_formats": [
                "short",
                "short"
            ],
            "yaxes": [
                {
                    "decimals": 2,
                    "format": "ops",
                    "label": "",
                    "logBase": 1,
                    "max": null,
                    "min": 0,
                    "show": true
                },
                {
                    "format": "short",
                    "logBase": 1,
                    "max": null,
                    "min": null,
                    "show": false
                }
            ],
            "yaxis": {
                "align": false,
                "alignLevel": null
            }
        },
        {
            "aliasColors": {},
            "bars": false,
            "dashLength": 10,
            "dashes": false,
            "datasource": null,
            "decimals": 2,
            "description": "This panel shows the number of assert events per second on average over the given time period. In most cases assertions are trivial, but you would want to check your log files if this counter spikes or is consistently high.",
            "editable": true,
            "error": false,
            "fieldConfig": {
                "defaults": {
                    "custom": {},
                    "links": []
                },
                "overrides": []
            },
            "fill": 2,
            "fillGradient": 0,
            "grid": {
                "leftLogBase": 1,
                "leftMax": null,
                "leftMin": 0,
                "rightLogBase": 1,
                "rightMax": null,
                "rightMin": null
            },
            "gridPos": {
                "h": 8,
                "w": 12,
                "x": 0,
                "y": 59
            },
            "height": "250px",
            "hiddenSeries": false,
            "id": 37,
            "legend": {
                "alignAsTable": true,
                "avg": true,
                "current": false,
                "hideZero": false,
                "max": true,
                "min": true,
                "rightSide": false,
                "show": true,
                "sort": "avg",
                "sortDesc": true,
                "total": false,
                "values": true
            },
            "lines": true,
            "linewidth": 2,
            "links": [],
            "nullPointMode": "null",
            "paceLength": 10,
            "percentage": false,
            "pluginVersion": "7.1.3",
            "pointradius": 5,
            "points": false,
            "renderer": "flot",
            "seriesOverrides": [],
            "spaceLength": 10,
            "stack": false,
            "steppedLine": false,
            "targets": [
                {
                    "expr": "avg by (service_name,type) (rate(mongodb_mongod_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or \nirate(mongodb_mongod_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]) or\nrate(mongodb_mongos_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or \nirate(mongodb_mongos_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]) or\nrate(mongodb_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or \nirate(mongodb_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]))",
                    "format": "time_series",
                    "hide": false,
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "{{type}}",
                    "refId": "J",
                    "step": 300
                }
            ],
            "thresholds": [],
            "timeFrom": null,
            "timeRegions": [],
            "timeShift": null,
            "title": "Assert Events",
            "tooltip": {
                "msResolution": false,
                "shared": true,
                "sort": 5,
                "value_type": "cumulative"
            },
            "type": "graph",
            "x-axis": true,
            "xaxis": {
                "buckets": null,
                "mode": "time",
                "name": null,
                "show": true,
                "values": []
            },
            "y-axis": true,
            "y_formats": [
                "short",
                "short"
            ],
            "yaxes": [
                {
                    "decimals": null,
                    "format": "short",
                    "label": "",
                    "logBase": 1,
                    "max": null,
                    "min": 0,
                    "show": true
                },
                {
                    "format": "short",
                    "logBase": 1,
                    "max": null,
                    "min": null,
                    "show": false
                }
            ],
            "yaxis": {
                "align": false,
                "alignLevel": null
            }
        },
        {
            "aliasColors": {},
            "bars": false,
            "dashLength": 10,
            "dashes": false,
            "datasource": null,
            "decimals": 2,
            "description": "Unix or Window memory page faults. Not necessarily from mongodb.",
            "editable": true,
            "error": false,
            "fieldConfig": {
                "defaults": {
                    "custom": {},
                    "links": []
                },
                "overrides": []
            },
            "fill": 2,
            "fillGradient": 0,
            "grid": {
                "leftLogBase": 1,
                "leftMax": null,
                "leftMin": 0,
                "rightLogBase": 1,
                "rightMax": null,
                "rightMin": null
            },
            "gridPos": {
                "h": 8,
                "w": 12,
                "x": 12,
                "y": 59
            },
            "height": "250px",
            "hiddenSeries": false,
            "id": 39,
            "legend": {
                "alignAsTable": true,
                "avg": true,
                "current": false,
                "max": true,
                "min": true,
                "show": true,
                "sort": "avg",
                "sortDesc": true,
                "total": false,
                "values": true
            },
            "lines": true,
            "linewidth": 2,
            "links": [],
            "nullPointMode": "null",
            "paceLength": 10,
            "percentage": false,
            "pluginVersion": "7.1.3",
            "pointradius": 5,
            "points": false,
            "renderer": "flot",
            "seriesOverrides": [],
            "spaceLength": 10,
            "stack": false,
            "steppedLine": false,
            "targets": [
                {
                    "expr": "avg by (service_name) (rate(mongodb_mongod_extra_info_page_faults_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or \nirate(mongodb_mongod_extra_info_page_faults_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]) or\nrate(mongodb_mongos_extra_info_page_faults_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or \nirate(mongodb_mongos_extra_info_page_faults_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]) or\nrate(mongodb_extra_info_page_faults_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or \nirate(mongodb_extra_info_page_faults_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]))",
                    "format": "time_series",
                    "hide": false,
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "Faults",
                    "refId": "J",
                    "step": 300
                }
            ],
            "thresholds": [],
            "timeFrom": null,
            "timeRegions": [],
            "timeShift": null,
            "title": "Page Faults",
            "tooltip": {
                "msResolution": false,
                "shared": true,
                "sort": 5,
                "value_type": "cumulative"
            },
            "type": "graph",
            "x-axis": true,
            "xaxis": {
                "buckets": null,
                "mode": "time",
                "name": null,
                "show": true,
                "values": []
            },
            "y-axis": true,
            "y_formats": [
                "short",
                "short"
            ],
            "yaxes": [
                {
                    "decimals": 2,
                    "format": "short",
                    "label": "",
                    "logBase": 1,
                    "max": null,
                    "min": 0,
                    "show": true
                },
                {
                    "format": "short",
                    "logBase": 1,
                    "max": null,
                    "min": null,
                    "show": false
                }
            ],
            "yaxis": {
                "align": false,
                "alignLevel": null
            }
        }
    ],
    "refresh": "1m",
    "schemaVersion": 26,
    "style": "dark",
    "tags": [
        "MongoDB",
        "Percona"
    ],
    "templating": {
        "list": [
            {
                "allFormat": "glob",
                "auto": true,
                "auto_count": 200,
                "auto_min": "1s",
                "current": {},
                "datasource": "prometheus",
                "hide": 0,
                "includeAll": false,
                "label": "Interval",
                "multi": false,
                "multiFormat": "glob",
                "name": "interval",
                "options": [
                    {
                        "selected": true,
                        "text": "auto",
                        "value": "$__auto_interval_interval"
                    },
                    {
                        "selected": false,
                        "text": "1s",
                        "value": "1s"
                    },
                    {
                        "selected": false,
                        "text": "5s",
                        "value": "5s"
                    },
                    {
                        "selected": false,
                        "text": "1m",
                        "value": "1m"
                    },
                    {
                        "selected": false,
                        "text": "5m",
                        "value": "5m"
                    },
                    {
                        "selected": false,
                        "text": "1h",
                        "value": "1h"
                    },
                    {
                        "selected": false,
                        "text": "6h",
                        "value": "6h"
                    },
                    {
                        "selected": false,
                        "text": "1d",
                        "value": "1d"
                    }
                ],
                "query": "1s,5s,1m,5m,1h,6h,1d",
                "refresh": 2,
                "skipUrlSync": false,
                "type": "interval"
            },
            {
                "allFormat": "blob",
                "allValue": "",
                "current": {
                    "selected": true,
                    "text": "All",
                    "value": "$__all"
                },
                "datasource": "prometheus",
                "definition": "",
                "hide": 0,
                "includeAll": false,
                "label": "Kubernetes Cluster",
                "multi": false,
                "multiFormat": "glob",
                "name": "vmcluster",
                "options": [],
                "query": "label_values(up, vmcluster)",
                "refresh": 2,
                "regex": "",
                "skipUrlSync": false,
                "sort": 1,
                "tagValuesQuery": "",
                "tags": [],
                "tagsQuery": "",
                "type": "query",
                "useTags": false
            },
            {
                "allFormat": "blob",
                "allValue": ".*",
                "current": {
                    "selected": true,
                    "text": "All",
                    "value": "$__all"
                },
                "datasource": "prometheus",
                "definition": "",
                "hide": 0,
                "includeAll": true,
                "label": "Namespace",
                "multi": false,
                "multiFormat": "glob",
                "name": "namespace",
                "options": [],
                "query": "label_values(mongodb_ss_connections{vmcluster=\"$vmcluster\"}, namespace)",
                "refresh": 2,
                "regex": "",
                "skipUrlSync": false,
                "sort": 1,
                "tagValuesQuery": "",
                "tags": [],
                "tagsQuery": "",
                "type": "query",
                "useTags": false
            },
            {
                "allFormat": "glob",
                "allValue": ".*",
                "current": {},
                "datasource": "prometheus",
                "definition": "label_values(mongodb_up{vmcluster=\"$vmcluster\", namespace=~\"$namespace\"}, service)",
                "hide": 0,
                "includeAll": true,
                "label": "Service Name",
                "multi": false,
                "multiFormat": "glob",
                "name": "service_name",
                "options": [],
                "query": "label_values(mongodb_up{vmcluster=\"$vmcluster\", namespace=~\"$namespace\"}, service)",
                "refresh": 2,
                "regex": "",
                "skipUrlSync": false,
                "sort": 5,
                "tagValuesQuery": null,
                "tags": [],
                "tagsQuery": "",
                "type": "query",
                "useTags": false
            }
        ]
    },
    "time": {
        "from": "now-12h",
        "to": "now"
    },
    "timepicker": {
        "hidden": false,
        "now": true,
        "refresh_intervals": [
            "5s",
            "10s",
            "30s",
            "1m",
            "5m",
            "15m",
            "30m",
            "1h",
            "2h",
            "1d"
        ],
        "time_options": [
            "5m",
            "15m",
            "1h",
            "6h",
            "12h",
            "24h",
            "2d",
            "7d",
            "30d"
        ]
    },
    "timezone": "",
    "title": "MongoDB Instance Summary",
    "uid": "mongodb-instance-summary",
    "version": 1
}
`
