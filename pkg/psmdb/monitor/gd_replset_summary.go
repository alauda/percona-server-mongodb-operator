package monitor

const MongoReplsetSummary = `
{
    "editable": false,
    "gnetId": null,
    "graphTooltip": 1,
    "id": null,
    "iteration": 1573821287689,
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
            "id": 1022,
            "panels": [],
            "title": "Overview",
            "type": "row"
        },
        {
            "cacheTimeout": null,
            "colorBackground": false,
            "colorValue": false,
            "colors": [
                "rgba(245, 54, 54, 0.9)",
                "rgba(237, 129, 40, 0.89)",
                "rgba(50, 172, 45, 0.97)"
            ],
            "datasource": null,
            "decimals": 0,
            "description": "This shows how many members are configured in the replica set.",
            "editable": true,
            "error": false,
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
                "x": 0,
                "y": 3
            },
            "height": "125px",
            "hideTimeOverride": false,
            "id": 59,
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
            "nullPointMode": "null",
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
                "show": true
            },
            "tableColumn": "",
            "targets": [
                {
                    "expr": "count by (set) (mongodb_mongod_replset_number_of_members{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", set=~\"$replset\"} or mongodb_mongod_replset_my_state{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", set=~\"$replset\"})",
                    "interval": "5m",
                    "intervalFactor": 1,
                    "legendFormat": "",
                    "metric": "",
                    "refId": "A",
                    "step": 300
                }
            ],
            "thresholds": "",
            "timeFrom": null,
            "title": "ReplSet Members",
            "type": "singlestat",
            "valueFontSize": "50%",
            "valueMaps": [],
            "valueName": "current"
        },
        {
            "cacheTimeout": null,
            "colorBackground": false,
            "colorValue": false,
            "colors": [
                "rgba(245, 54, 54, 0.9)",
                "rgba(237, 129, 40, 0.89)",
                "rgba(50, 172, 45, 0.97)"
            ],
            "datasource": null,
            "decimals": 2,
            "description": "This shows the time since the last election.",
            "editable": true,
            "error": false,
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
                "w": 8,
                "x": 8,
                "y": 3
            },
            "height": "125px",
            "hideTimeOverride": false,
            "id": 1227,
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
            "nullPointMode": "null",
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
                "show": true
            },
            "tableColumn": "",
            "targets": [
                {
                    "expr": "time() - max(mongodb_mongod_replset_member_election_date{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"})",
                    "interval": "5m",
                    "intervalFactor": 1,
                    "legendFormat": "",
                    "metric": "",
                    "refId": "A",
                    "step": 300
                }
            ],
            "thresholds": "",
            "timeFrom": null,
            "title": "ReplSet Last Election",
            "type": "singlestat",
            "valueFontSize": "50%",
            "valueMaps": [],
            "valueName": "current"
        },
        {
            "cacheTimeout": null,
            "colorBackground": false,
            "colorValue": false,
            "colors": [
                "rgba(245, 54, 54, 0.9)",
                "rgba(237, 129, 40, 0.89)",
                "rgba(50, 172, 45, 0.97)"
            ],
            "datasource": null,
            "decimals": 2,
            "description": "This panel shows how far behind in replication this member is if it is a secondary. This number may be high it the instance is running as a delayed secondary member.",
            "editable": true,
            "error": false,
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
                "w": 8,
                "x": 16,
                "y": 3
            },
            "height": "125px",
            "hideTimeOverride": true,
            "id": 77,
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
            "nullPointMode": "null",
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
                "show": true
            },
            "tableColumn": "",
            "targets": [
                {
                    "expr": "avg by (set) (max_over_time(mongodb_mongod_replset_member_replication_lag{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", set=\"$replset\",service=~\"$service_name\"}[${__range}]))",
                    "format": "time_series",
                    "hide": false,
                    "instant": true,
                    "interval": "$interval",
                    "legendFormat": "",
                    "refId": "B"
                }
            ],
            "thresholds": "",
            "timeFrom": "1m",
            "title": "Avg ReplSet Lag",
            "type": "singlestat",
            "valueFontSize": "50%",
            "valueMaps": [],
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
                },
                {
                    "color": "#FF7383",
                    "text": "Exporter is not connected"
                }
            ],
            "crosshairColor": "#8F070C",
            "datasource": null,
            "description": "ReplSet statuses during the select time range.",
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
                "h": 8,
                "w": 24,
                "x": 0,
                "y": 5
            },
            "highlightOnMouseover": true,
            "id": 1015,
            "legendSortBy": "count",
            "lineColor": "rgba(0,0,0,0.1)",
            "metricNameColor": "#000000",
            "pluginVersion": "6.5.1",
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
                    "expr": "mongodb_mongod_replset_my_state{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", set=~\"$replset\",service=~\"$service_name\"}",
                    "interval": "$interval",
                    "legendFormat": "{{service_name}}",
                    "refId": "A"
                }
            ],
            "textSize": 18,
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
            "title": "ReplSet States",
            "type": "natel-discrete-panel",
            "units": "short",
            "useTimePrecision": false,
            "valueMaps": [
                {
                    "op": "=",
                    "text": "Exporter is not connected",
                    "value": "null"
                },
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
            "writeMetricNames": true
        },
        {
            "collapsed": false,
            "datasource": null,
            "gridPos": {
                "h": 1,
                "w": 24,
                "x": 0,
                "y": 13
            },
            "id": 1024,
            "panels": [],
            "title": "Replication Lag",
            "type": "row"
        },
        {
            "aliasColors": {},
            "bars": false,
            "dashLength": 10,
            "dashes": false,
            "datasource": null,
            "decimals": 2,
            "description": "MongoDB replication lag occurs when the secondary node cannot replicate data fast enough to keep up with the rate that data is being written to the primary node. It could be caused by something as simple as network latency, packet loss within your network, or a routing issue.",
            "editable": true,
            "error": false,
            "fieldConfig": {
                "defaults": {
                    "custom": {},
                    "links": [
                        {
                            "targetBlank": true,
                            "title": "MongoDB Instance Summary - ${__series.name}",
                            "url": "/graph/d/mongodb-instance-summary/mongodb-instance-summary?var-service_name=${__series.name}$&$__url_time_range"
                        }
                    ]
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
                "y": 14
            },
            "hiddenSeries": false,
            "id": 14,
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
            "links": [],
            "nullPointMode": "null as zero",
            "options": {
                "alertThreshold": true
            },
            "paceLength": 10,
            "percentage": false,
            "pluginVersion": "7.3.7",
            "pointradius": 5,
            "points": false,
            "renderer": "flot",
            "seriesOverrides": [
                {
                    "alias": "Avg",
                    "color": "#C4162A",
                    "fill": 0,
                    "legend": false,
                    "stack": false
                }
            ],
            "spaceLength": 10,
            "stack": false,
            "steppedLine": false,
            "targets": [
                {
                    "expr": "avg by (service_name) (max(max_over_time(mongodb_mongod_replset_member_replication_lag{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", set=~\"$replset\",service=~\"$secondary\"}[$interval]) > 0) by (service_name,set) or max(max_over_time(mongodb_mongod_replset_member_replication_lag{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", set=~\"$replset\",service=~\"$secondary\"}[5m]) > 0) by (service_name,set))",
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "{{service_name}}",
                    "refId": "A",
                    "step": 300
                },
                {
                    "expr": "avg by (set) (max(max_over_time(mongodb_mongod_replset_member_replication_lag{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", set=~\"$replset\",service=~\"$secondary\"}[$interval]) > 0) by (service_name,set) or max(max_over_time(mongodb_mongod_replset_member_replication_lag{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", set=~\"$replset\",service=~\"$secondary\"}[5m]) > 0) by (service_name,set))",
                    "hide": true,
                    "interval": "$interval",
                    "legendFormat": "Avg",
                    "refId": "B"
                }
            ],
            "thresholds": [],
            "timeFrom": null,
            "timeRegions": [],
            "timeShift": null,
            "title": "Replication Lag",
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
                "s",
                "short"
            ],
            "yaxes": [
                {
                    "decimals": 2,
                    "format": "s",
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
            "collapsed": false,
            "datasource": null,
            "gridPos": {
                "h": 1,
                "w": 24,
                "x": 0,
                "y": 22
            },
            "id": 1019,
            "panels": [],
            "title": "Operations",
            "type": "row"
        },
        {
            "aliasColors": {},
            "bars": false,
            "dashLength": 10,
            "dashes": false,
            "datasource": null,
            "decimals": 2,
            "description": "Operations are classified by legacy wire protocol type (insert, update, and delete only).",
            "editable": true,
            "error": false,
            "fieldConfig": {
                "defaults": {
                    "custom": {},
                    "links": [
                        {
                            "targetBlank": true,
                            "title": "MongoDB Instance Summary - ${__field.labels.service_name}",
                            "url": "/graph/d/mongodb-instance-summary/mongodb-instance-summary?var-service_name=${__field.labels.service_name}&$__url_time_range"
                        }
                    ]
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
                "w": 8,
                "x": 0,
                "y": 23
            },
            "hiddenSeries": false,
            "id": 1020,
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
            "links": [
                {
                    "targetBlank": true,
                    "title": "MongoDB Instance Summary - $service_name",
                    "url": "/graph/d/mongodb-instance-summary/mongodb-instance-summary?var-service_name=$service_name$__url_time_range"
                }
            ],
            "maxPerRow": null,
            "nullPointMode": "null",
            "options": {
                "alertThreshold": true
            },
            "paceLength": 10,
            "percentage": false,
            "pluginVersion": "7.3.7",
            "pointradius": 5,
            "points": false,
            "renderer": "flot",
            "repeat": "service_name",
            "repeatDirection": "h",
            "seriesOverrides": [],
            "spaceLength": 10,
            "stack": false,
            "steppedLine": false,
            "targets": [
                {
                    "expr": "avg by (service_name,type) (rate(mongodb_op_counters_repl_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or irate(mongodb_op_counters_repl_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]))",
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "repl - {{type}}",
                    "refId": "A",
                    "step": 300
                },
                {
                    "expr": "avg by (service_name,type) (rate(mongodb_mongod_op_counters_repl_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or irate(mongodb_mongod_op_counters_repl_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]))",
                    "interval": "$interval",
                    "legendFormat": "repl - {{type}}",
                    "refId": "B"
                },
                {
                    "expr": "avg by (service_name,type) (rate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or irate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]))",
                    "interval": "$interval",
                    "legendFormat": "{{type}}",
                    "refId": "C"
                }
            ],
            "thresholds": [],
            "timeFrom": null,
            "timeRegions": [],
            "timeShift": null,
            "title": "Operations - $service_name",
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
            "description": "Operations are classified by legacy wire protocol type (insert, update, and delete only).",
            "editable": true,
            "error": false,
            "fieldConfig": {
                "defaults": {
                    "custom": {},
                    "links": [
                        {
                            "targetBlank": true,
                            "title": "MongoDB Instance Summary - ${__field.labels.service_name}",
                            "url": "/graph/d/mongodb-instance-summary/mongodb-instance-summary?var-service_name=${__field.labels.service_name}&$__url_time_range"
                        }
                    ]
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
                "w": 8,
                "x": 8,
                "y": 23
            },
            "hiddenSeries": false,
            "id": 1837,
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
            "links": [
                {
                    "targetBlank": true,
                    "title": "MongoDB Instance Summary - $service_name",
                    "url": "/graph/d/mongodb-instance-summary/mongodb-instance-summary?var-service_name=$service_name$__url_time_range"
                }
            ],
            "maxPerRow": null,
            "nullPointMode": "null",
            "options": {
                "alertThreshold": true
            },
            "paceLength": 10,
            "percentage": false,
            "pluginVersion": "7.3.7",
            "pointradius": 5,
            "points": false,
            "renderer": "flot",
            "repeatDirection": "h",
            "repeatIteration": 1620988880578,
            "repeatPanelId": 1020,
            "seriesOverrides": [],
            "spaceLength": 10,
            "stack": false,
            "steppedLine": false,
            "targets": [
                {
                    "expr": "avg by (service_name,type) (rate(mongodb_op_counters_repl_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or irate(mongodb_op_counters_repl_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]))",
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "repl - {{type}}",
                    "refId": "A",
                    "step": 300
                },
                {
                    "expr": "avg by (service_name,type) (rate(mongodb_mongod_op_counters_repl_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or irate(mongodb_mongod_op_counters_repl_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]))",
                    "interval": "$interval",
                    "legendFormat": "repl - {{type}}",
                    "refId": "B"
                },
                {
                    "expr": "avg by (service_name,type) (rate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or irate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]))",
                    "interval": "$interval",
                    "legendFormat": "{{type}}",
                    "refId": "C"
                }
            ],
            "thresholds": [],
            "timeFrom": null,
            "timeRegions": [],
            "timeShift": null,
            "title": "Operations - $service_name",
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
            "description": "Operations are classified by legacy wire protocol type (insert, update, and delete only).",
            "editable": true,
            "error": false,
            "fieldConfig": {
                "defaults": {
                    "custom": {},
                    "links": [
                        {
                            "targetBlank": true,
                            "title": "MongoDB Instance Summary - ${__field.labels.service_name}",
                            "url": "/graph/d/mongodb-instance-summary/mongodb-instance-summary?var-service_name=${__field.labels.service_name}&$__url_time_range"
                        }
                    ]
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
                "w": 8,
                "x": 16,
                "y": 23
            },
            "hiddenSeries": false,
            "id": 1838,
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
            "links": [
                {
                    "targetBlank": true,
                    "title": "MongoDB Instance Summary - $service_name",
                    "url": "/graph/d/mongodb-instance-summary/mongodb-instance-summary?var-service_name=$service_name$__url_time_range"
                }
            ],
            "maxPerRow": null,
            "nullPointMode": "null",
            "options": {
                "alertThreshold": true
            },
            "paceLength": 10,
            "percentage": false,
            "pluginVersion": "7.3.7",
            "pointradius": 5,
            "points": false,
            "renderer": "flot",
            "repeatDirection": "h",
            "repeatIteration": 1620988880578,
            "repeatPanelId": 1020,
            "seriesOverrides": [],
            "spaceLength": 10,
            "stack": false,
            "steppedLine": false,
            "targets": [
                {
                    "expr": "avg by (service_name,type) (rate(mongodb_op_counters_repl_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or irate(mongodb_op_counters_repl_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]))",
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "repl - {{type}}",
                    "refId": "A",
                    "step": 300
                },
                {
                    "expr": "avg by (service_name,type) (rate(mongodb_mongod_op_counters_repl_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or irate(mongodb_mongod_op_counters_repl_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]))",
                    "interval": "$interval",
                    "legendFormat": "repl - {{type}}",
                    "refId": "B"
                },
                {
                    "expr": "avg by (service_name,type) (rate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or irate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]))",
                    "interval": "$interval",
                    "legendFormat": "{{type}}",
                    "refId": "C"
                }
            ],
            "thresholds": [],
            "timeFrom": null,
            "timeRegions": [],
            "timeShift": null,
            "title": "Operations - $service_name",
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
            "collapsed": false,
            "datasource": null,
            "gridPos": {
                "h": 1,
                "w": 24,
                "x": 0,
                "y": 40
            },
            "id": 1017,
            "panels": [],
            "title": "Max Heartbeat Time / Elections",
            "type": "row"
        },
        {
            "aliasColors": {},
            "bars": false,
            "dashLength": 10,
            "dashes": false,
            "datasource": null,
            "decimals": 2,
            "description": "Time span between now and last heartbeat from replicaset members.",
            "editable": true,
            "error": false,
            "fieldConfig": {
                "defaults": {
                    "custom": {},
                    "links": [
                        {
                            "targetBlank": true,
                            "title": "MongoDB Instance Summary - ${__series.name}",
                            "url": "/graph/d/mongodb-instance-summary/mongodb-instance-summary?var-service_name=${__series.name}$&$__url_time_range"
                        }
                    ]
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
                "y": 41
            },
            "hiddenSeries": false,
            "id": 75,
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
            "options": {
                "alertThreshold": true
            },
            "paceLength": 10,
            "percentage": false,
            "pluginVersion": "7.3.7",
            "pointradius": 5,
            "points": false,
            "renderer": "flot",
            "seriesOverrides": [],
            "spaceLength": 10,
            "stack": false,
            "steppedLine": false,
            "targets": [
                {
                    "expr": "time() - avg by (service_name) (max(mongodb_mongod_replset_member_last_heartbeat{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}) by (name)) * on (name) group_right avg by (service_name) (mongodb_mongod_replset_my_name{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"})",
                    "hide": false,
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "{{service_name}}",
                    "metric": "",
                    "refId": "J",
                    "step": 300
                },
                {
                    "expr": "avg by (service_name) (max(mongodb_rs_heartbeatIntervalMillis{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}) by (name) / 1000) * on (name) group_right avg by (service_name) (mongodb_mongod_replset_my_name{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"})",
                    "hide": true,
                    "interval": "$interval",
                    "legendFormat": "Interval - {service_name}}",
                    "refId": "B"
                }
            ],
            "thresholds": [],
            "timeFrom": null,
            "timeRegions": [],
            "timeShift": null,
            "title": "Max Heartbeat Time",
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
                "s",
                "short"
            ],
            "yaxes": [
                {
                    "decimals": 2,
                    "format": "s",
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
            "description": "Count of elections. Usually zero; 1 count by each healthy node will appear in each election. Happens when the primary role changes due to either normal maintenance or trouble events.",
            "editable": true,
            "error": false,
            "fieldConfig": {
                "defaults": {
                    "custom": {},
                    "links": [
                        {
                            "targetBlank": true,
                            "title": "MongoDB Instance Summary - ${__series.name}",
                            "url": "/graph/d/mongodb-instance-summary/mongodb-instance-summary?var-service_name=${__series.name}$&$__url_time_range"
                        }
                    ]
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
                "y": 41
            },
            "hiddenSeries": false,
            "id": 12,
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
            "options": {
                "alertThreshold": true
            },
            "paceLength": 10,
            "percentage": false,
            "pluginVersion": "7.3.7",
            "pointradius": 5,
            "points": false,
            "renderer": "flot",
            "seriesOverrides": [],
            "spaceLength": 10,
            "stack": false,
            "steppedLine": false,
            "targets": [
                {
                    "expr": "max by (service_name) (changes(mongodb_mongod_replset_member_election_date{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]))",
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "{{service_name}}",
                    "refId": "A",
                    "step": 300
                }
            ],
            "thresholds": [],
            "timeFrom": null,
            "timeRegions": [],
            "timeShift": null,
            "title": "Elections",
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
            "description": "Timespan 'window' between newest and the oldest op in the Oplog collection.",
            "editable": true,
            "error": false,
            "fieldConfig": {
                "defaults": {
                    "custom": {},
                    "links": [
                        {
                            "targetBlank": true,
                            "title": "MongoDB Instance Summary - ${__field.labels.service_name}",
                            "url": "/graph/d/mongodb-instance-summary/mongodb-instance-summary?var-service_name=${__field.labels.service_name}&$__url_time_range"
                        }
                    ]
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
                "w": 8,
                "x": 8,
                "y": 50
            },
            "hiddenSeries": false,
            "id": 1841,
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
            "options": {
                "alertThreshold": true
            },
            "paceLength": 10,
            "percentage": false,
            "pluginVersion": "7.3.7",
            "pointradius": 5,
            "points": false,
            "renderer": "flot",
            "repeatDirection": "h",
            "repeatIteration": 1620988880578,
            "repeatPanelId": 27,
            "seriesOverrides": [
                {
                    "alias": "Oplog Range",
                    "yaxis": 2
                }
            ],
            "spaceLength": 10,
            "stack": false,
            "steppedLine": false,
            "targets": [
                {
                    "expr": "time()-avg by (service_name) (mongodb_mongod_replset_oplog_tail_timestamp{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"})",
                    "hide": false,
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "Now to End",
                    "metric": "",
                    "refId": "J",
                    "step": 300
                },
                {
                    "expr": "(time() * 1000) - avg by (service_name) (mongodb_oplog_stats_end{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"})",
                    "interval": "$interval",
                    "legendFormat": "Now to End",
                    "refId": "B"
                },
                {
                    "expr": "avg by (service_name) (mongodb_mongod_replset_oplog_head_timestamp{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}-mongodb_mongod_replset_oplog_tail_timestamp{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"})",
                    "hide": false,
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "Oplog Range",
                    "metric": "",
                    "refId": "A",
                    "step": 300
                },
                {
                    "expr": "avg by (service_name) (mongodb_oplog_stats_end{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"} - mongodb_oplog_stats_start{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}) / 1000",
                    "interval": "$interval",
                    "legendFormat": "Oplog Range",
                    "refId": "C"
                }
            ],
            "thresholds": [],
            "timeFrom": null,
            "timeRegions": [],
            "timeShift": null,
            "title": "Oplog Recovery Window - $service_name",
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
                "s",
                "short"
            ],
            "yaxes": [
                {
                    "decimals": 2,
                    "format": "s",
                    "label": "",
                    "logBase": 1,
                    "max": null,
                    "min": 0,
                    "show": true
                },
                {
                    "format": "s",
                    "logBase": 1,
                    "max": null,
                    "min": null,
                    "show": true
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
            "description": "Timespan 'window' between newest and the oldest op in the Oplog collection.",
            "editable": true,
            "error": false,
            "fieldConfig": {
                "defaults": {
                    "custom": {},
                    "links": [
                        {
                            "targetBlank": true,
                            "title": "MongoDB Instance Summary - ${__field.labels.service_name}",
                            "url": "/graph/d/mongodb-instance-summary/mongodb-instance-summary?var-service_name=${__field.labels.service_name}&$__url_time_range"
                        }
                    ]
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
                "w": 8,
                "x": 16,
                "y": 50
            },
            "hiddenSeries": false,
            "id": 1842,
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
            "options": {
                "alertThreshold": true
            },
            "paceLength": 10,
            "percentage": false,
            "pluginVersion": "7.3.7",
            "pointradius": 5,
            "points": false,
            "renderer": "flot",
            "repeatDirection": "h",
            "repeatIteration": 1620988880578,
            "repeatPanelId": 27,
            "seriesOverrides": [
                {
                    "alias": "Oplog Range",
                    "yaxis": 2
                }
            ],
            "spaceLength": 10,
            "stack": false,
            "steppedLine": false,
            "targets": [
                {
                    "expr": "time()-avg by (service_name) (mongodb_mongod_replset_oplog_tail_timestamp{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"})",
                    "hide": false,
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "Now to End",
                    "metric": "",
                    "refId": "J",
                    "step": 300
                },
                {
                    "expr": "(time() * 1000) - avg by (service_name) (mongodb_oplog_stats_end{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"})",
                    "interval": "$interval",
                    "legendFormat": "Now to End",
                    "refId": "B"
                },
                {
                    "expr": "avg by (service_name) (mongodb_mongod_replset_oplog_head_timestamp{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}-mongodb_mongod_replset_oplog_tail_timestamp{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"})",
                    "hide": false,
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "Oplog Range",
                    "metric": "",
                    "refId": "A",
                    "step": 300
                },
                {
                    "expr": "avg by (service_name) (mongodb_oplog_stats_end{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"} - mongodb_oplog_stats_start{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}) / 1000",
                    "interval": "$interval",
                    "legendFormat": "Oplog Range",
                    "refId": "C"
                }
            ],
            "thresholds": [],
            "timeFrom": null,
            "timeRegions": [],
            "timeShift": null,
            "title": "Oplog Recovery Window - $service_name",
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
                "s",
                "short"
            ],
            "yaxes": [
                {
                    "decimals": 2,
                    "format": "s",
                    "label": "",
                    "logBase": 1,
                    "max": null,
                    "min": 0,
                    "show": true
                },
                {
                    "format": "s",
                    "logBase": 1,
                    "max": null,
                    "min": null,
                    "show": true
                }
            ],
            "yaxis": {
                "align": false,
                "alignLevel": null
            }
        },
        {
            "collapsed": true,
            "datasource": null,
            "gridPos": {
                "h": 1,
                "w": 24,
                "x": 0,
                "y": 58
            },
            "id": 1070,
            "panels": [
                {
                    "aliasColors": {},
                    "bars": false,
                    "dashLength": 10,
                    "dashes": false,
                    "datasource": null,
                    "decimals": 2,
                    "description": "Repl buffer ops applied per sec.",
                    "editable": true,
                    "error": false,
                    "fieldConfig": {
                        "defaults": {
                            "custom": {},
                            "links": [
                                {
                                    "targetBlank": true,
                                    "title": "MongoDB Instance Summary - ${__series.name}",
                                    "url": "/graph/d/mongodb-instance-summary/mongodb-instance-summary?var-service_name=${__series.name}$&$__url_time_range"
                                }
                            ]
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
                    "hiddenSeries": false,
                    "id": 85,
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
                    "options": {
                        "alertThreshold": true
                    },
                    "paceLength": 10,
                    "percentage": false,
                    "pluginVersion": "7.3.7",
                    "pointradius": 5,
                    "points": false,
                    "renderer": "flot",
                    "seriesOverrides": [],
                    "spaceLength": 10,
                    "stack": false,
                    "steppedLine": false,
                    "targets": [
                        {
                            "expr": "avg by (service_name) (mongodb_mongod_metrics_repl_buffer_count{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"})",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}}",
                            "refId": "A",
                            "step": 300
                        }
                    ],
                    "thresholds": [],
                    "timeFrom": null,
                    "timeRegions": [],
                    "timeShift": null,
                    "title": "Oplog Buffered Operations",
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
                        "ms",
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
                    "description": "Time spent per second waiting for or fetching oplog docs in replication.",
                    "editable": true,
                    "error": false,
                    "fieldConfig": {
                        "defaults": {
                            "custom": {},
                            "links": [
                                {
                                    "targetBlank": true,
                                    "title": "MongoDB Instance Summary - ${__series.name}",
                                    "url": "/graph/d/mongodb-instance-summary/mongodb-instance-summary?var-service_name=${__series.name}$&$__url_time_range"
                                }
                            ]
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
                    "hiddenSeries": false,
                    "id": 79,
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
                    "options": {
                        "alertThreshold": true
                    },
                    "paceLength": 10,
                    "percentage": false,
                    "pluginVersion": "7.3.7",
                    "pointradius": 5,
                    "points": false,
                    "renderer": "flot",
                    "seriesOverrides": [],
                    "spaceLength": 10,
                    "stack": false,
                    "steppedLine": false,
                    "targets": [
                        {
                            "expr": "avg by (service_name) (rate(mongodb_mongod_metrics_repl_network_getmores_total_milliseconds{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or irate(mongodb_mongod_metrics_repl_network_getmores_total_milliseconds{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]))",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}}",
                            "refId": "A",
                            "step": 300
                        }
                    ],
                    "thresholds": [],
                    "timeFrom": null,
                    "timeRegions": [],
                    "timeShift": null,
                    "title": "Oplog Getmore Time",
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
                        "ms",
                        "short"
                    ],
                    "yaxes": [
                        {
                            "decimals": 2,
                            "format": "ms",
                            "label": "",
                            "logBase": 1,
                            "max": null,
                            "min": "0",
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
                    "description": "Times spent per second A) pre-loading oplog ops into parallel-executable batches B) Times spent pre-loading index values and C) repl batch apply phase.",
                    "editable": true,
                    "error": false,
                    "fieldConfig": {
                        "defaults": {
                            "custom": {},
                            "links": [
                                {
                                    "targetBlank": true,
                                    "title": "MongoDB Instance Summary - ${__field.labels.service_name}",
                                    "url": "/graph/d/mongodb-instance-summary/mongodb-instance-summary?var-service_name=${__field.labels.service_name}&$__url_time_range"
                                }
                            ]
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
                        "w": 8,
                        "x": 0,
                        "y": 67
                    },
                    "hiddenSeries": false,
                    "id": 84,
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
                    "options": {
                        "alertThreshold": true
                    },
                    "paceLength": 10,
                    "percentage": false,
                    "pluginVersion": "7.3.7",
                    "pointradius": 5,
                    "points": false,
                    "renderer": "flot",
                    "repeat": "service_name",
                    "repeatDirection": "h",
                    "seriesOverrides": [],
                    "spaceLength": 10,
                    "stack": false,
                    "steppedLine": false,
                    "targets": [
                        {
                            "expr": "avg by (service_name) (rate(mongodb_mongod_metrics_repl_preload_docs_total_milliseconds{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or irate(mongodb_mongod_metrics_repl_preload_docs_total_milliseconds{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]))",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "Document Preload",
                            "refId": "A",
                            "step": 300
                        },
                        {
                            "expr": "avg by (service_name) (rate(mongodb_mongod_metrics_repl_preload_indexes_total_milliseconds{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or irate(mongodb_mongod_metrics_repl_preload_indexes_total_milliseconds{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]))",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "Index Preload",
                            "metric": "mongodb_mongod_metrics_repl_preload_indexes_total_milliseconds",
                            "refId": "B",
                            "step": 300
                        },
                        {
                            "expr": "avg by (service_name) (rate(mongodb_mongod_metrics_repl_apply_batches_total_milliseconds{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or irate(mongodb_mongod_metrics_repl_apply_batches_total_milliseconds{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]))",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "Batch Apply",
                            "metric": "mongodb_mongod_metrics_repl_preload_indexes_total_milliseconds",
                            "refId": "C",
                            "step": 300
                        }
                    ],
                    "thresholds": [],
                    "timeFrom": null,
                    "timeRegions": [],
                    "timeShift": null,
                    "title": "Oplog Processing Time - $service_name",
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
                        "ms",
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
                    "description": "Current hard-coded max and actual size of repl batch buffer.",
                    "editable": true,
                    "error": false,
                    "fieldConfig": {
                        "defaults": {
                            "custom": {},
                            "links": [
                                {
                                    "targetBlank": true,
                                    "title": "MongoDB Instance Summary - ${__field.labels.service_name}",
                                    "url": "/graph/d/mongodb-instance-summary/mongodb-instance-summary?var-service_name=${__field.labels.service_name}&$__url_time_range"
                                }
                            ]
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
                        "w": 8,
                        "x": 8,
                        "y": 67
                    },
                    "hiddenSeries": false,
                    "id": 80,
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
                    "options": {
                        "alertThreshold": true
                    },
                    "paceLength": 10,
                    "percentage": false,
                    "pluginVersion": "7.3.7",
                    "pointradius": 5,
                    "points": false,
                    "renderer": "flot",
                    "repeat": "service_name",
                    "repeatDirection": "h",
                    "seriesOverrides": [
                        {
                            "alias": "Max",
                            "color": "#C4162A",
                            "fill": 0
                        }
                    ],
                    "spaceLength": 10,
                    "stack": false,
                    "steppedLine": false,
                    "targets": [
                        {
                            "expr": "avg by (service_name) (mongodb_mongod_metrics_repl_buffer_size_bytes{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"})",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "Used",
                            "refId": "A",
                            "step": 300
                        },
                        {
                            "expr": "avg by (service_name) (mongodb_mongod_metrics_repl_buffer_max_size_bytes{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"})",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "Max",
                            "refId": "B",
                            "step": 300
                        }
                    ],
                    "thresholds": [],
                    "timeFrom": null,
                    "timeRegions": [],
                    "timeShift": null,
                    "title": "Oplog Buffer Capacity - $service_name",
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
                        "ms",
                        "short"
                    ],
                    "yaxes": [
                        {
                            "decimals": 2,
                            "format": "bytes",
                            "label": "",
                            "logBase": 1,
                            "max": null,
                            "min": "0",
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
                    "description": "Count of A) getmores executed B) index values pre-loaded C) oplog docs applied per second against oplog collection in replication.",
                    "editable": true,
                    "error": false,
                    "fieldConfig": {
                        "defaults": {
                            "custom": {},
                            "links": [
                                {
                                    "targetBlank": true,
                                    "title": "MongoDB Instance Summary - ${__field.labels.service_name}",
                                    "url": "/graph/d/mongodb-instance-summary/mongodb-instance-summary?var-service_name=${__field.labels.service_name}&$__url_time_range"
                                }
                            ]
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
                        "w": 8,
                        "x": 16,
                        "y": 67
                    },
                    "hiddenSeries": false,
                    "id": 81,
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
                    "options": {
                        "alertThreshold": true
                    },
                    "paceLength": 10,
                    "percentage": false,
                    "pluginVersion": "7.3.7",
                    "pointradius": 5,
                    "points": false,
                    "renderer": "flot",
                    "repeat": "service_name",
                    "repeatDirection": "h",
                    "seriesOverrides": [],
                    "spaceLength": 10,
                    "stack": false,
                    "steppedLine": false,
                    "targets": [
                        {
                            "expr": "avg by (service_name) (rate(mongodb_mongod_metrics_repl_preload_docs_num_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or irate(mongodb_mongod_metrics_repl_preload_docs_num_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]))",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "Document Preload",
                            "metric": "mongodb_mongod_metrics_repl_preload_indexes_num_total",
                            "refId": "A",
                            "step": 300
                        },
                        {
                            "expr": "avg by (service_name) (rate(mongodb_mongod_metrics_repl_preload_indexes_num_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or irate(mongodb_mongod_metrics_repl_preload_indexes_num_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]))",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "Index Preload",
                            "metric": "mongodb_mongod_metrics_repl_preload_indexes_num_total",
                            "refId": "B",
                            "step": 300
                        },
                        {
                            "expr": "avg by (service_name) (rate(mongodb_mongod_metrics_repl_apply_ops_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or irate(mongodb_mongod_metrics_repl_apply_ops_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]))",
                            "interval": "$interval",
                            "intervalFactor": 2,
                            "legendFormat": "Batch Apply",
                            "refId": "C",
                            "step": 120
                        }
                    ],
                    "thresholds": [],
                    "timeFrom": null,
                    "timeRegions": [],
                    "timeShift": null,
                    "title": "Oplog Operations - $service_name",
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
                        "ms",
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
                    "description": "Count of A) getmores executed B) index values pre-loaded C) oplog docs applied per second against oplog collection in replication.",
                    "editable": true,
                    "error": false,
                    "fieldConfig": {
                        "defaults": {
                            "custom": {},
                            "links": [
                                {
                                    "targetBlank": true,
                                    "title": "MongoDB Instance Summary - ${__field.labels.service_name}",
                                    "url": "/graph/d/mongodb-instance-summary/mongodb-instance-summary?var-service_name=${__field.labels.service_name}&$__url_time_range"
                                }
                            ]
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
                        "w": 8,
                        "x": 8,
                        "y": 83
                    },
                    "hiddenSeries": false,
                    "id": 1847,
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
                    "options": {
                        "alertThreshold": true
                    },
                    "paceLength": 10,
                    "percentage": false,
                    "pluginVersion": "7.3.7",
                    "pointradius": 5,
                    "points": false,
                    "renderer": "flot",
                    "repeatDirection": "h",
                    "repeatIteration": 1620988880578,
                    "repeatPanelId": 81,
                    "seriesOverrides": [],
                    "spaceLength": 10,
                    "stack": false,
                    "steppedLine": false,
                    "targets": [
                        {
                            "expr": "avg by (service_name) (rate(mongodb_mongod_metrics_repl_preload_docs_num_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or irate(mongodb_mongod_metrics_repl_preload_docs_num_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]))",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "Document Preload",
                            "metric": "mongodb_mongod_metrics_repl_preload_indexes_num_total",
                            "refId": "A",
                            "step": 300
                        },
                        {
                            "expr": "avg by (service_name) (rate(mongodb_mongod_metrics_repl_preload_indexes_num_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or irate(mongodb_mongod_metrics_repl_preload_indexes_num_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]))",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "Index Preload",
                            "metric": "mongodb_mongod_metrics_repl_preload_indexes_num_total",
                            "refId": "B",
                            "step": 300
                        },
                        {
                            "expr": "avg by (service_name) (rate(mongodb_mongod_metrics_repl_apply_ops_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or irate(mongodb_mongod_metrics_repl_apply_ops_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]))",
                            "interval": "$interval",
                            "intervalFactor": 2,
                            "legendFormat": "Batch Apply",
                            "refId": "C",
                            "step": 120
                        }
                    ],
                    "thresholds": [],
                    "timeFrom": null,
                    "timeRegions": [],
                    "timeShift": null,
                    "title": "Oplog Operations - $service_name",
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
                        "ms",
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
                    "description": "Count of A) getmores executed B) index values pre-loaded C) oplog docs applied per second against oplog collection in replication.",
                    "editable": true,
                    "error": false,
                    "fieldConfig": {
                        "defaults": {
                            "custom": {},
                            "links": [
                                {
                                    "targetBlank": true,
                                    "title": "MongoDB Instance Summary - ${__field.labels.service_name}",
                                    "url": "/graph/d/mongodb-instance-summary/mongodb-instance-summary?var-service_name=${__field.labels.service_name}&$__url_time_range"
                                }
                            ]
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
                        "w": 8,
                        "x": 16,
                        "y": 83
                    },
                    "hiddenSeries": false,
                    "id": 1848,
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
                    "options": {
                        "alertThreshold": true
                    },
                    "paceLength": 10,
                    "percentage": false,
                    "pluginVersion": "7.3.7",
                    "pointradius": 5,
                    "points": false,
                    "renderer": "flot",
                    "repeatDirection": "h",
                    "repeatIteration": 1620988880578,
                    "repeatPanelId": 81,
                    "seriesOverrides": [],
                    "spaceLength": 10,
                    "stack": false,
                    "steppedLine": false,
                    "targets": [
                        {
                            "expr": "avg by (service_name) (rate(mongodb_mongod_metrics_repl_preload_docs_num_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or irate(mongodb_mongod_metrics_repl_preload_docs_num_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]))",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "Document Preload",
                            "metric": "mongodb_mongod_metrics_repl_preload_indexes_num_total",
                            "refId": "A",
                            "step": 300
                        },
                        {
                            "expr": "avg by (service_name) (rate(mongodb_mongod_metrics_repl_preload_indexes_num_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or irate(mongodb_mongod_metrics_repl_preload_indexes_num_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]))",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "Index Preload",
                            "metric": "mongodb_mongod_metrics_repl_preload_indexes_num_total",
                            "refId": "B",
                            "step": 300
                        },
                        {
                            "expr": "avg by (service_name) (rate(mongodb_mongod_metrics_repl_apply_ops_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or irate(mongodb_mongod_metrics_repl_apply_ops_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]))",
                            "interval": "$interval",
                            "intervalFactor": 2,
                            "legendFormat": "Batch Apply",
                            "refId": "C",
                            "step": 120
                        }
                    ],
                    "thresholds": [],
                    "timeFrom": null,
                    "timeRegions": [],
                    "timeShift": null,
                    "title": "Oplog Operations - $service_name",
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
                        "ms",
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
                }
            ],
            "title": "Oplog Details",
            "type": "row"
        },
        {
            "collapsed": true,
            "datasource": null,
            "gridPos": {
                "h": 1,
                "w": 24,
                "x": 0,
                "y": 59
            },
            "id": 1306,
            "panels": [
                {
                    "activePatternIndex": 0,
                    "datasource": null,
                    "defaultPattern": {
                        "bgColors": "green|orange|red",
                        "bgColors_overrides": "0->green|2->red|1->yellow",
                        "clickable_cells_link": "",
                        "col_name": "Value",
                        "decimals": 2,
                        "defaultBGColor": "",
                        "defaultTextColor": "",
                        "delimiter": "|",
                        "displayTemplate": "_value_",
                        "enable_bgColor": false,
                        "enable_bgColor_overrides": false,
                        "enable_clickable_cells": false,
                        "enable_textColor": false,
                        "enable_textColor_overrides": false,
                        "enable_time_based_thresholds": false,
                        "enable_transform": false,
                        "enable_transform_overrides": false,
                        "filter": {
                            "value_above": "",
                            "value_below": ""
                        },
                        "format": "none",
                        "name": "Default Pattern",
                        "null_color": "darkred",
                        "null_textcolor": "black",
                        "null_value": "No data",
                        "pattern": "*",
                        "row_col_wrapper": "_",
                        "row_name": "Service Name",
                        "textColors": "red|orange|green",
                        "textColors_overrides": "0->red|2->green|1->yellow",
                        "thresholds": "70,90",
                        "time_based_thresholds": [],
                        "tooltipTemplate": "Series : _series_ <br/>Row Name : _row_name_ <br/>Col Name : _col_name_ <br/>Value : _value_",
                        "transform_values": "_value_|_value_|_value_",
                        "transform_values_overrides": "0->down|1->up",
                        "valueName": "current"
                    },
                    "default_title_for_rows": "Service Name",
                    "fieldConfig": {
                        "defaults": {
                            "custom": {}
                        },
                        "overrides": []
                    },
                    "first_column_link": "/graph/d/mongodb-instance-summary/mongodb-instance-summary",
                    "gridPos": {
                        "h": 6,
                        "w": 24,
                        "x": 0,
                        "y": 60
                    },
                    "hide_headers": false,
                    "id": 1352,
                    "patterns": [
                        {
                            "bgColors": "green|orange|red",
                            "bgColors_overrides": "0->green|2->red|1->yellow",
                            "clickable_cells_link": "",
                            "col_name": "QPS",
                            "decimals": 2,
                            "defaultBGColor": "",
                            "defaultTextColor": "",
                            "delimiter": "|",
                            "displayTemplate": "_value_",
                            "enable_bgColor": false,
                            "enable_bgColor_overrides": false,
                            "enable_clickable_cells": false,
                            "enable_textColor": false,
                            "enable_textColor_overrides": false,
                            "enable_time_based_thresholds": false,
                            "enable_transform": false,
                            "enable_transform_overrides": false,
                            "filter": {
                                "value_above": "",
                                "value_below": ""
                            },
                            "format": "ops",
                            "name": "QPS",
                            "null_color": "darkred",
                            "null_textcolor": "black",
                            "null_value": "No data",
                            "pattern": ".*QPS",
                            "row_col_wrapper": "_",
                            "row_name": "_0_",
                            "textColors": "red|orange|green",
                            "textColors_overrides": "0->red|2->green|1->yellow",
                            "thresholds": "70,90",
                            "time_based_thresholds": [],
                            "tooltipTemplate": "Queries per Second (QPS)",
                            "transform_values": "_value_|_value_|_value_",
                            "transform_values_overrides": "0->down|1->up",
                            "valueName": "current"
                        },
                        {
                            "bgColors": "green|orange|red",
                            "bgColors_overrides": "0->green|2->red|1->yellow",
                            "clickable_cells_link": "",
                            "col_name": "Average Latency",
                            "decimals": 2,
                            "defaultBGColor": "",
                            "defaultTextColor": "",
                            "delimiter": "|",
                            "displayTemplate": "_value_",
                            "enable_bgColor": false,
                            "enable_bgColor_overrides": false,
                            "enable_clickable_cells": false,
                            "enable_textColor": false,
                            "enable_textColor_overrides": false,
                            "enable_time_based_thresholds": false,
                            "enable_transform": false,
                            "enable_transform_overrides": false,
                            "filter": {
                                "value_above": "",
                                "value_below": ""
                            },
                            "format": "\u00b5s",
                            "name": "Latency",
                            "null_color": "darkred",
                            "null_textcolor": "black",
                            "null_value": "No data",
                            "pattern": ".*Latency",
                            "row_col_wrapper": "_",
                            "row_name": "_0_",
                            "textColors": "red|orange|green",
                            "textColors_overrides": "0->red|2->green|1->yellow",
                            "thresholds": "70,90",
                            "time_based_thresholds": [],
                            "tooltipTemplate": "Latency statistics for database commands.",
                            "transform_values": "_value_|_value_|_value_",
                            "transform_values_overrides": "0->down|1->up",
                            "valueName": "current"
                        },
                        {
                            "bgColors": "green|orange|red",
                            "bgColors_overrides": "0->green|2->red|1->yellow",
                            "clickable_cells_link": "",
                            "col_name": "_1_",
                            "decimals": 2,
                            "defaultBGColor": "",
                            "defaultTextColor": "",
                            "delimiter": "|",
                            "displayTemplate": "_value_",
                            "enable_bgColor": false,
                            "enable_bgColor_overrides": false,
                            "enable_clickable_cells": false,
                            "enable_textColor": false,
                            "enable_textColor_overrides": false,
                            "enable_time_based_thresholds": false,
                            "enable_transform": false,
                            "enable_transform_overrides": false,
                            "filter": {
                                "value_above": "",
                                "value_below": ""
                            },
                            "format": "s",
                            "name": "Uptime",
                            "null_color": "darkred",
                            "null_textcolor": "black",
                            "null_value": "No data",
                            "pattern": ".*Uptime",
                            "row_col_wrapper": "_",
                            "row_name": "_0_",
                            "textColors": "red|orange|green",
                            "textColors_overrides": "0->red|2->green|1->yellow",
                            "thresholds": "70,90",
                            "time_based_thresholds": [],
                            "tooltipTemplate": "The parameter shows how long a service has been \u201cup\u201d and running without a shut down or restart.",
                            "transform_values": "_value_|_value_|_value_",
                            "transform_values_overrides": "0->down|1->up",
                            "valueName": "current"
                        },
                        {
                            "bgColors": "green|orange|red",
                            "bgColors_overrides": "0->green|2->red|1->yellow",
                            "clickable_cells_link": "",
                            "col_name": "Open Connections",
                            "decimals": "0",
                            "defaultBGColor": "",
                            "defaultTextColor": "",
                            "delimiter": "|",
                            "displayTemplate": "_value_",
                            "enable_bgColor": false,
                            "enable_bgColor_overrides": false,
                            "enable_clickable_cells": false,
                            "enable_textColor": false,
                            "enable_textColor_overrides": false,
                            "enable_time_based_thresholds": false,
                            "enable_transform": false,
                            "enable_transform_overrides": false,
                            "filter": {
                                "value_above": "",
                                "value_below": ""
                            },
                            "format": "none",
                            "name": "Open Connections",
                            "null_color": "darkred",
                            "null_textcolor": "black",
                            "null_value": "No data",
                            "pattern": ".*Connections",
                            "row_col_wrapper": "_",
                            "row_name": "_0_",
                            "textColors": "red|orange|green",
                            "textColors_overrides": "0->red|2->green|1->yellow",
                            "thresholds": "70,90",
                            "time_based_thresholds": [],
                            "tooltipTemplate": "Connects to a MongoDB instance and to a specified database on that instance.",
                            "transform_values": "_value_|_value_|_value_",
                            "transform_values_overrides": "0->down|1->up",
                            "valueName": "current"
                        },
                        {
                            "bgColors": "green|orange|red",
                            "bgColors_overrides": "0->green|2->red|1->yellow",
                            "clickable_cells_link": "",
                            "col_name": "_1_",
                            "decimals": "0",
                            "defaultBGColor": "",
                            "defaultTextColor": "",
                            "delimiter": "|",
                            "displayTemplate": "_value_",
                            "enable_bgColor": false,
                            "enable_bgColor_overrides": false,
                            "enable_clickable_cells": false,
                            "enable_textColor": false,
                            "enable_textColor_overrides": false,
                            "enable_time_based_thresholds": false,
                            "enable_transform": false,
                            "enable_transform_overrides": false,
                            "filter": {
                                "value_above": "",
                                "value_below": ""
                            },
                            "format": "none",
                            "name": "Cursors",
                            "null_color": "darkred",
                            "null_textcolor": "black",
                            "null_value": "No data",
                            "pattern": ".* Cursors",
                            "row_col_wrapper": "_",
                            "row_name": "_0_",
                            "textColors": "red|orange|green",
                            "textColors_overrides": "0->red|2->green|1->yellow",
                            "thresholds": "70,90",
                            "time_based_thresholds": [],
                            "tooltipTemplate": "When the db. collection. find () function is used to search for documents in the collection, the result returns a pointer to the collection of documents returned which is called a cursor.",
                            "transform_values": "_value_|_value_|_value_",
                            "transform_values_overrides": "0->down|1->up",
                            "valueName": "current"
                        }
                    ],
                    "row_col_wrapper": "_",
                    "sorting_props": {
                        "col_index": -1,
                        "direction": "desc"
                    },
                    "targets": [
                        {
                            "expr": "avg by (service_name) (mongodb_instance_uptime_seconds{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"})",
                            "interval": "$interval",
                            "legendFormat": "{{service_name}} | Uptime",
                            "refId": "C"
                        },
                        {
                            "expr": "sum by (service_name) (rate(mongodb_mongod_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type!=\"command\"}[$interval]) or irate(mongodb_mongod_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type!=\"command\"}[5m]) or rate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type!=\"command\"}[$interval]) or irate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type!=\"command\"}[5m]))",
                            "instant": false,
                            "interval": "$interval",
                            "legendFormat": "{{service_name}} | QPS",
                            "refId": "A"
                        },
                        {
                            "expr": "avg by (service_name) (rate(mongodb_mongod_op_latencies_latency_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"command\"}[$interval]) / (rate(mongodb_mongod_op_latencies_ops_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"command\"}[$interval]) > 0) or\nirate(mongodb_mongod_op_latencies_latency_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"command\"}[5m]) / (irate(mongodb_mongod_op_latencies_ops_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"command\"}[5m]) > 0))",
                            "interval": "$interval",
                            "legendFormat": "{{service_name}} | Latency",
                            "refId": "B"
                        },
                        {
                            "expr": "avg by (service_name) (max_over_time(mongodb_mongod_connections{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"current\"}[$interval]) or \nmax_over_time(mongodb_mongod_connections{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"current\"}[5m]) or\nmax_over_time(mongodb_mongos_connections{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"current\"}[$interval]) or\nmax_over_time(mongodb_mongos_connections{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"current\"}[5m]) or\nmax_over_time(mongodb_connections{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"current\"}[$interval]) or\nmax_over_time(mongodb_connections{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"current\"}[5m]))",
                            "interval": "$interval",
                            "legendFormat": "{{service_name}} | Connections",
                            "refId": "D"
                        },
                        {
                            "expr": "avg by (service_name) (max_over_time(mongodb_mongod_metrics_cursor_open{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",state=\"total\"}[$interval]) or\nmax_over_time(mongodb_mongod_metrics_cursor_open{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",state=\"total\"}[5m]) or\nmax_over_time(mongodb_mongod_cursors{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",state=\"total\"}[$interval]) or\nmax_over_time(mongodb_mongod_cursors{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",state=\"total\"}[5m]) or\nmax_over_time(mongodb_mongos_metrics_cursor_open{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",state=\"total\"}[$interval]) or \nmax_over_time(mongodb_mongos_metrics_cursor_open{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",state=\"total\"}[5m]) or\nmax_over_time(mongodb_mongos_cursors{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",state=\"total\"}[$interval]) or\nmax_over_time(mongodb_mongos_cursors{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",state=\"total\"}[5m]))",
                            "instant": true,
                            "interval": "$interval",
                            "legendFormat": "{{service_name}} | Cursors ",
                            "refId": "E"
                        }
                    ],
                    "text_alignment_header": "center",
                    "text_alignment_values": "center",
                    "timeFrom": null,
                    "timeShift": null,
                    "title": "Services Details",
                    "type": "yesoreyeram-boomtable-panel"
                },
                {
                    "aliasColors": {},
                    "bars": false,
                    "dashLength": 10,
                    "dashes": false,
                    "datasource": null,
                    "decimals": 2,
                    "description": "Average latency of operations (classified by read, write, or (other) command).",
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
                        "w": 8,
                        "x": 0,
                        "y": 66
                    },
                    "height": "250px",
                    "hiddenSeries": false,
                    "id": 1558,
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
                    "options": {
                        "alertThreshold": true
                    },
                    "paceLength": 10,
                    "percentage": false,
                    "pluginVersion": "7.3.7",
                    "pointradius": 5,
                    "points": false,
                    "renderer": "flot",
                    "repeat": "service_name",
                    "repeatDirection": "h",
                    "seriesOverrides": [],
                    "spaceLength": 10,
                    "stack": false,
                    "steppedLine": false,
                    "targets": [
                        {
                            "expr": "avg by (service_name,type) (rate(mongodb_mongod_op_latencies_latency_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) / (rate(mongodb_mongod_op_latencies_ops_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) > 0) or irate(mongodb_mongod_op_latencies_latency_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]) / (irate(mongodb_mongod_op_latencies_ops_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]) > 0))",
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
                    "title": "Latency Detail - $service_name",
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
                            "format": "\u00b5s",
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
                        "w": 8,
                        "x": 8,
                        "y": 66
                    },
                    "height": "250px",
                    "hiddenSeries": false,
                    "id": 1397,
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
                    "options": {
                        "alertThreshold": true
                    },
                    "paceLength": 10,
                    "percentage": false,
                    "pluginVersion": "7.3.7",
                    "pointradius": 5,
                    "points": false,
                    "renderer": "flot",
                    "repeat": "service_name",
                    "repeatDirection": "h",
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
                    "title": "Document Operations - $service_name",
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
                    "description": "MongoDB keeps most recently used data in RAM. If you have created indexes for your queries and your working data set fits in RAM, MongoDB serves all queries from memory.",
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
                        "w": 8,
                        "x": 16,
                        "y": 66
                    },
                    "height": "250px",
                    "hiddenSeries": false,
                    "id": 1012,
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
                    "options": {
                        "alertThreshold": true
                    },
                    "paceLength": 10,
                    "percentage": false,
                    "pluginVersion": "7.3.7",
                    "pointradius": 5,
                    "points": false,
                    "renderer": "flot",
                    "repeat": "service_name",
                    "repeatDirection": "h",
                    "seriesOverrides": [],
                    "spaceLength": 10,
                    "stack": false,
                    "steppedLine": false,
                    "targets": [
                        {
                            "expr": "rate(mongodb_tcmalloc_cache_bytes{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or irate(mongodb_tcmalloc_cache_bytes{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m])",
                            "format": "time_series",
                            "hide": false,
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{cache}} - {{type}}",
                            "refId": "J",
                            "step": 300
                        },
                        {
                            "expr": "avg by (service_name) (mongodb_ss_tcmalloc_tcmalloc_thread_cache_free_bytes{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"})",
                            "interval": "$interval",
                            "legendFormat": "Free Thread Cache",
                            "refId": "A"
                        },
                        {
                            "expr": "avg by (service_name) (mongodb_ss_tcmalloc_tcmalloc_central_cache_free_bytes{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"})",
                            "interval": "$interval",
                            "legendFormat": "Free Central Cache",
                            "refId": "B"
                        },
                        {
                            "expr": "avg by (service_name) (mongodb_ss_tcmalloc_tcmalloc_transfer_cache_free_bytes{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"})",
                            "interval": "$interval",
                            "legendFormat": "Free Transfer Cache",
                            "refId": "C"
                        }
                    ],
                    "thresholds": [],
                    "timeFrom": null,
                    "timeRegions": [],
                    "timeShift": null,
                    "title": "Cache - $service_name",
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
                            "format": "bytes",
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
            "title": "MongoDB Services Summary",
            "type": "row"
        }
    ],
    "refresh": "1m",
    "schemaVersion": 26,
    "style": "dark",
    "tags": [
        "MongoDB_HA",
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
                "error": null,
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
                "definition": "label_values(mongodb_mongod_replset_my_state{vmcluster=\"$vmcluster\", namespace=~\"$namespace\"}, set)",
                "error": null,
                "hide": 0,
                "includeAll": true,
                "label": "Replica Set",
                "multi": false,
                "multiFormat": "glob",
                "name": "replset",
                "options": [],
                "query": "label_values(mongodb_mongod_replset_my_state{vmcluster=\"$vmcluster\", namespace=~\"$namespace\"}, set)",
                "refresh": 2,
                "regex": "",
                "skipUrlSync": false,
                "sort": 5,
                "tagValuesQuery": null,
                "tags": [],
                "tagsQuery": null,
                "type": "query",
                "useTags": false
            },
            {
                "allFormat": "glob",
                "allValue": ".*",
                "current": {},
                "datasource": "prometheus",
                "definition": "",
                "error": null,
                "hide": 0,
                "includeAll": true,
                "label": "Service Name",
                "multi": true,
                "multiFormat": "glob",
                "name": "service_name",
                "options": [],
                "query": "label_values(mongodb_mongod_replset_my_state{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", set=~\"$replset\"}, service)",
                "refresh": 2,
                "regex": "",
                "skipUrlSync": false,
                "sort": 5,
                "tagValuesQuery": null,
                "tags": [],
                "tagsQuery": "",
                "type": "query",
                "useTags": false
            },
            {
                "allValue": ".*",
                "current": {},
                "datasource": "prometheus",
                "definition": "query_result(mongodb_mongod_replset_my_state{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", set=\"$replset\"}==2)",
                "error": null,
                "hide": 2,
                "includeAll": true,
                "label": "Secondary",
                "multi": true,
                "name": "secondary",
                "options": [],
                "query": "query_result(mongodb_mongod_replset_my_state{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", set=\"$replset\"}==2)",
                "refresh": 2,
                "regex": "/.*service_name=\"(.*)\",service_type.*/",
                "skipUrlSync": false,
                "sort": 0,
                "tagValuesQuery": "",
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
    "title": "MongoDB ReplSet Summary",
    "uid": "mongodb-replicaset-summary",
    "version": 1
}
`
