package monitor

const MongoInstanceOverview = `
{
    "editable": false,
    "gnetId": null,
    "graphTooltip": 1,
    "id": null,
    "iteration": 1574863062208,
    "links": [],
    "panels": [
        {
            "datasource": null,
            "gridPos": {
                "h": 1,
                "w": 24,
                "x": 0,
                "y": 2
            },
            "id": 68,
            "title": "Overview",
            "type": "row"
        },
        {
            "cacheTimeout": null,
            "colorBackground": false,
            "colorValue": false,
            "colors": [
                "#299c46",
                "#1F60C4",
                "#d44a3a"
            ],
            "datasource": null,
            "decimals": null,
            "fieldConfig": {
                "defaults": {},
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
                "x": 0,
                "y": 3
            },
            "id": 70,
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
                "show": true
            },
            "tableColumn": "",
            "targets": [
                {
                    "expr": "count(mongodb_up{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"})",
                    "format": "time_series",
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "",
                    "refId": "A"
                }
            ],
            "thresholds": "",
            "timeFrom": null,
            "timeShift": null,
            "title": "Services",
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
            "colorValue": true,
            "colors": [
                "#d44a3a",
                "#FF780A",
                "#299c46"
            ],
            "datasource": null,
            "decimals": 2,
            "fieldConfig": {
                "defaults": {},
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
                "w": 5,
                "x": 4,
                "y": 3
            },
            "id": 72,
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
                "show": true
            },
            "tableColumn": "",
            "targets": [
                {
                    "expr": "min(mongodb_instance_uptime_seconds{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"})",
                    "format": "time_series",
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "",
                    "refId": "A"
                }
            ],
            "thresholds": "300,3600",
            "timeFrom": null,
            "timeShift": null,
            "title": "Min MongoDB Uptime",
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
                "#1F60C4",
                "#d44a3a"
            ],
            "datasource": null,
            "decimals": 2,
            "fieldConfig": {
                "defaults": {},
                "overrides": []
            },
            "format": "decmbytes",
            "gauge": {
                "maxValue": 100,
                "minValue": 0,
                "show": false,
                "thresholdLabels": false,
                "thresholdMarkers": true
            },
            "gridPos": {
                "h": 2,
                "w": 5,
                "x": 9,
                "y": 3
            },
            "id": 98,
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
                "show": true
            },
            "tableColumn": "",
            "targets": [
                {
                    "expr": "sum(mongodb_memory{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=~\"resident\"})",
                    "format": "time_series",
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "",
                    "refId": "A"
                }
            ],
            "thresholds": "",
            "timeFrom": null,
            "timeShift": null,
            "title": "Total Used Resident Memory",
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
                "#1F60C4",
                "#d44a3a"
            ],
            "datasource": null,
            "decimals": 2,
            "fieldConfig": {
                "defaults": {},
                "overrides": []
            },
            "format": "decmbytes",
            "gauge": {
                "maxValue": 100,
                "minValue": 0,
                "show": false,
                "thresholdLabels": false,
                "thresholdMarkers": true
            },
            "gridPos": {
                "h": 2,
                "w": 5,
                "x": 14,
                "y": 3
            },
            "id": 1000,
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
                "show": true
            },
            "tableColumn": "",
            "targets": [
                {
                    "expr": "sum(mongodb_memory{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=~\"virtual\"})",
                    "format": "time_series",
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "",
                    "refId": "A"
                }
            ],
            "thresholds": "",
            "timeFrom": null,
            "timeShift": null,
            "title": "Total Used Virtual Memory",
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
                "#1F60C4",
                "#d44a3a"
            ],
            "datasource": null,
            "decimals": 2,
            "fieldConfig": {
                "defaults": {},
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
                "w": 5,
                "x": 19,
                "y": 3
            },
            "id": 75,
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
                "show": true
            },
            "tableColumn": "",
            "targets": [
                {
                    "expr": "sum(rate(mongodb_mongod_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type!=\"command\"}[$interval]) or \nirate(mongodb_mongod_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type!=\"command\"}[5m]) or \nrate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type!=\"command\"}[$interval]) or \nirate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type!=\"command\"}[5m]))",
                    "format": "time_series",
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "",
                    "refId": "A"
                }
            ],
            "thresholds": "",
            "timeFrom": null,
            "timeShift": null,
            "title": "Total Current QPS",
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
            "collapsed": false,
            "datasource": null,
            "gridPos": {
                "h": 1,
                "w": 24,
                "x": 0,
                "y": 5
            },
            "id": 77,
            "panels": [],
            "type": "row"
        },
        {
            "cacheTimeout": null,
            "colorBackground": false,
            "colorValue": false,
            "colors": [
                "#299c46",
                "#FF780A",
                "#d44a3a"
            ],
            "datasource": null,
            "decimals": 0,
            "fieldConfig": {
                "defaults": {},
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
                "w": 6,
                "x": 0,
                "y": 6
            },
            "id": 135,
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
                "show": true
            },
            "tableColumn": "",
            "targets": [
                {
                    "expr": "sum(mongodb_mongod_connections{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"current\"} or\nmongodb_mongos_connections{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"current\"} or\nmongodb_connections{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"current\"})",
                    "format": "time_series",
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "",
                    "refId": "A"
                }
            ],
            "thresholds": "10,50",
            "timeFrom": null,
            "timeShift": null,
            "title": "Top Connections",
            "type": "singlestat",
            "valueFontSize": "50%",
            "valueMaps": [
                {
                    "op": "=",
                    "text": "N/A",
                    "value": "null"
                }
            ],
            "valueName": "max"
        },
        {
            "cacheTimeout": null,
            "colorBackground": false,
            "colorValue": false,
            "colors": [
                "#299c46",
                "#FF780A",
                "#d44a3a"
            ],
            "datasource": null,
            "decimals": 0,
            "fieldConfig": {
                "defaults": {},
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
                "w": 6,
                "x": 6,
                "y": 6
            },
            "id": 136,
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
                "show": true
            },
            "tableColumn": "",
            "targets": [
                {
                    "expr": "sum(mongodb_mongod_metrics_cursor_open{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"} or\nmongodb_mongod_cursors{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"} or\nmongodb_mongos_metrics_cursor_open{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"} or \nmongodb_mongos_cursors{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"})",
                    "format": "time_series",
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "",
                    "refId": "A"
                }
            ],
            "thresholds": "10,50",
            "timeFrom": null,
            "timeShift": null,
            "title": "Top Opened Cursors",
            "type": "singlestat",
            "valueFontSize": "50%",
            "valueMaps": [
                {
                    "op": "=",
                    "text": "N/A",
                    "value": "null"
                }
            ],
            "valueName": "max"
        },
        {
            "cacheTimeout": null,
            "colorBackground": false,
            "colorValue": false,
            "colors": [
                "#299c46",
                "#1F60C4",
                "#d44a3a"
            ],
            "datasource": null,
            "decimals": 2,
            "fieldConfig": {
                "defaults": {},
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
                "w": 6,
                "x": 12,
                "y": 6
            },
            "id": 140,
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
                "show": true
            },
            "tableColumn": "",
            "targets": [
                {
                    "expr": "sum(rate(mongodb_mongod_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type!=\"command\"}[$interval]) or \nirate(mongodb_mongod_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type!=\"command\"}[5m]) or \nrate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type!=\"command\"}[$interval]) or\nirate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type!=\"command\"}[5m]))",
                    "format": "time_series",
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "",
                    "refId": "A"
                }
            ],
            "thresholds": "",
            "timeFrom": null,
            "timeShift": null,
            "title": "Min QPS",
            "type": "singlestat",
            "valueFontSize": "50%",
            "valueMaps": [
                {
                    "op": "=",
                    "text": "N/A",
                    "value": "null"
                }
            ],
            "valueName": "min"
        },
        {
            "cacheTimeout": null,
            "colorBackground": false,
            "colorValue": false,
            "colors": [
                "#299c46",
                "#1F60C4",
                "#d44a3a"
            ],
            "datasource": null,
            "decimals": 2,
            "fieldConfig": {
                "defaults": {},
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
                "w": 6,
                "x": 18,
                "y": 6
            },
            "id": 139,
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
                "show": true
            },
            "tableColumn": "",
            "targets": [
                {
                    "expr": "max(rate(mongodb_mongod_op_latencies_latency_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"command\"}[$interval]) / \n(rate(mongodb_mongod_op_latencies_ops_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"command\"}[$interval]) > 0) or\nirate(mongodb_mongod_op_latencies_latency_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"command\"}[5m]) / \n(irate(mongodb_mongod_op_latencies_ops_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"command\"}[5m]) > 0)\n)",
                    "format": "time_series",
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "",
                    "refId": "A"
                }
            ],
            "thresholds": "",
            "timeFrom": null,
            "timeShift": null,
            "title": "Max Latency",
            "type": "singlestat",
            "valueFontSize": "50%",
            "valueMaps": [
                {
                    "op": "=",
                    "text": "N/A",
                    "value": "null"
                }
            ],
            "valueName": "max"
        },
        {
            "collapsed": true,
            "datasource": null,
            "gridPos": {
                "h": 1,
                "w": 24,
                "x": 0,
                "y": 8
            },
            "id": 114,
            "panels": [
                {
                    "aliasColors": {},
                    "bars": false,
                    "dashLength": 10,
                    "dashes": false,
                    "datasource": null,
                    "decimals": 2,
                    "description": "TCP connections (Incoming) in mongod processes",
                    "editable": true,
                    "error": false,
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
                        "y": 9
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
                    "lines": false,
                    "linewidth": 2,
                    "links": [],
                    "nullPointMode": "null",
                    "options": {
                        "dataLinks": [
                            {
                                "targetBlank": true,
                                "title": "MongoDB Instance Summary - ${__series.name}",
                                "url": "/graph/d/mongodb-instance-summary/mongodb-instance-summary?var-service_name=${__series.name}&$__url_time_range"
                            }
                        ]
                    },
                    "paceLength": 10,
                    "percentage": false,
                    "pointradius": 1,
                    "points": true,
                    "renderer": "flot",
                    "seriesOverrides": [
                        {
                            "alias": "Avg",
                            "color": "#C4162A",
                            "fill": 0,
                            "legend": false,
                            "lines": true,
                            "linewidth": 2,
                            "points": false
                        }
                    ],
                    "spaceLength": 10,
                    "stack": false,
                    "steppedLine": false,
                    "targets": [
                        {
                            "expr": "topk(5,avg by (service_name) ((max_over_time(mongodb_mongod_connections{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"current\"}[$interval]) or max_over_time(mongodb_mongod_connections{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"current\"}[5m]) or\nmax_over_time(mongodb_mongos_connections{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"current\"}[$interval]) or max_over_time(mongodb_mongos_connections{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"current\"}[5m]) or\nmax_over_time(mongodb_connections{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"current\"}[$interval]) or max_over_time(mongodb_connections{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"current\"}[5m]))))",
                            "format": "time_series",
                            "hide": false,
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}}",
                            "refId": "J",
                            "step": 300
                        },
                        {
                            "expr": "avg(max_over_time(mongodb_mongod_connections{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"current\"}[$interval]) or max_over_time(mongodb_mongod_connections{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"current\"}[5m]) or\nmax_over_time(mongodb_mongos_connections{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"current\"}[$interval]) or max_over_time(mongodb_mongos_connections{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"current\"}[5m]) or\nmax_over_time(mongodb_connections{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"current\"}[$interval]) or max_over_time(mongodb_connections{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"current\"}[5m]))",
                            "format": "time_series",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "Avg",
                            "refId": "A"
                        }
                    ],
                    "thresholds": [],
                    "timeFrom": null,
                    "timeRegions": [],
                    "timeShift": null,
                    "title": "Top 5 Connections",
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
                    "colors": [
                        "#299c46",
                        "rgba(237, 129, 40, 0.89)",
                        "#d44a3a"
                    ],
                    "datasource": null,
                    "gridPos": {
                        "h": 8,
                        "w": 12,
                        "x": 12,
                        "y": 9
                    },
                    "id": 138,
                    "interval": "",
                    "links": [],
                    "options": {},
                    "polystat": {
                        "animationSpeed": 2500,
                        "columnAutoSize": true,
                        "columns": "",
                        "displayLimit": 100,
                        "fontAutoScale": true,
                        "fontSize": 12,
                        "fontType": "Roboto",
                        "globalDecimals": 2,
                        "globalDisplayMode": "all",
                        "globalDisplayTextTriggeredEmpty": "OK",
                        "globalOperatorName": "current",
                        "globalUnitFormat": "short",
                        "gradientEnabled": true,
                        "hexagonSortByDirection": 1,
                        "hexagonSortByField": "name",
                        "maxMetrics": 0,
                        "polygonBorderColor": "black",
                        "polygonBorderSize": 2,
                        "polygonGlobalFillColor": "#0a50a1",
                        "radius": "",
                        "radiusAutoSize": true,
                        "rowAutoSize": true,
                        "rows": "",
                        "shape": "hexagon_pointed_top",
                        "tooltipDisplayMode": "all",
                        "tooltipDisplayTextTriggeredEmpty": "OK",
                        "tooltipFontSize": 12,
                        "tooltipFontType": "Roboto",
                        "tooltipPrimarySortDirection": 2,
                        "tooltipPrimarySortField": "thresholdLevel",
                        "tooltipSecondarySortDirection": 2,
                        "tooltipSecondarySortField": "value",
                        "tooltipTimestampEnabled": true
                    },
                    "savedComposites": [],
                    "savedOverrides": [
                        {
                            "colors": [
                                "#299c46",
                                "#e5ac0e",
                                "#bf1b00",
                                "#ffffff"
                            ],
                            "decimals": "",
                            "enabled": true,
                            "label": "OVERRIDE 1",
                            "metricName": "/.*/",
                            "operatorName": "current",
                            "prefix": "",
                            "sanitizeURLEnabled": false,
                            "scaledDecimals": null,
                            "suffix": "",
                            "thresholds": [
                                {
                                    "color": "#299c46",
                                    "state": 0,
                                    "value": 10
                                },
                                {
                                    "color": "#e5ac0e",
                                    "state": 1,
                                    "value": 30
                                },
                                {
                                    "color": "#bf1b00",
                                    "state": 2,
                                    "value": 50
                                }
                            ],
                            "unitFormat": "short"
                        }
                    ],
                    "targets": [
                        {
                            "expr": "avg by (service_name) (max_over_time(mongodb_mongod_connections{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"current\"}[$interval]) or max_over_time(mongodb_mongod_connections{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"current\"}[5m]) or\nmax_over_time(mongodb_mongos_connections{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"current\"}[$interval]) or max_over_time(mongodb_mongos_connections{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"current\"}[5m]) or\nmax_over_time(mongodb_connections{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"current\"}[$interval]) or max_over_time(mongodb_connections{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"current\"}[5m]))",
                            "format": "time_series",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}}",
                            "refId": "A"
                        }
                    ],
                    "timeFrom": null,
                    "timeShift": null,
                    "title": "Current Connections",
                    "type": "grafana-polystat-panel"
                }
            ],
            "title": "Connections Detail",
            "type": "row"
        },
        {
            "collapsed": true,
            "datasource": null,
            "gridPos": {
                "h": 1,
                "w": 24,
                "x": 0,
                "y": 9
            },
            "id": 118,
            "panels": [
                {
                    "aliasColors": {},
                    "bars": false,
                    "dashLength": 10,
                    "dashes": false,
                    "datasource": null,
                    "decimals": 2,
                    "description": "Open cursors count",
                    "editable": true,
                    "error": false,
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
                        "y": 10
                    },
                    "height": "250px",
                    "hiddenSeries": false,
                    "id": 142,
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
                    "lines": false,
                    "linewidth": 2,
                    "links": [],
                    "nullPointMode": "null",
                    "options": {
                        "dataLinks": [
                            {
                                "targetBlank": true,
                                "title": "MongoDB Instance Summary - ${__series.name}",
                                "url": "/graph/d/mongodb-instance-summary/mongodb-instance-summary?var-service_name=${__series.name}&$__url_time_range"
                            }
                        ]
                    },
                    "paceLength": 10,
                    "percentage": false,
                    "pointradius": 1,
                    "points": true,
                    "renderer": "flot",
                    "seriesOverrides": [
                        {
                            "alias": "Avg",
                            "color": "#C4162A",
                            "fill": 0,
                            "legend": false,
                            "lines": true,
                            "points": false
                        }
                    ],
                    "spaceLength": 10,
                    "stack": false,
                    "steppedLine": false,
                    "targets": [
                        {
                            "expr": "topk(5,avg by (service_name) ((mongodb_mongod_metrics_cursor_open{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"total\"} or\nmongodb_mongod_cursors{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"total\"} or\nmongodb_mongos_metrics_cursor_open{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"total\"} or \nmongodb_mongos_cursors{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"total\"})))",
                            "format": "time_series",
                            "hide": false,
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}} ",
                            "refId": "J",
                            "step": 300
                        },
                        {
                            "expr": "avg(mongodb_mongod_metrics_cursor_open{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"total\"} or\nmongodb_mongod_cursors{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"total\"} or\nmongodb_mongos_metrics_cursor_open{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"total\"} or \nmongodb_mongos_cursors{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"total\"})",
                            "format": "time_series",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "Avg",
                            "refId": "A"
                        }
                    ],
                    "thresholds": [],
                    "timeFrom": null,
                    "timeRegions": [],
                    "timeShift": null,
                    "title": "Top 5 Total Cursors",
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
                            "show": true
                        }
                    ],
                    "yaxis": {
                        "align": false,
                        "alignLevel": null
                    }
                },
                {
                    "colors": [
                        "#299c46",
                        "rgba(237, 129, 40, 0.89)",
                        "#d44a3a"
                    ],
                    "datasource": null,
                    "description": "",
                    "gridPos": {
                        "h": 8,
                        "w": 12,
                        "x": 12,
                        "y": 10
                    },
                    "id": 143,
                    "links": [],
                    "options": {},
                    "polystat": {
                        "animationSpeed": 2500,
                        "columnAutoSize": true,
                        "columns": "",
                        "displayLimit": 100,
                        "fontAutoScale": true,
                        "fontSize": 12,
                        "fontType": "Roboto",
                        "globalDecimals": 2,
                        "globalDisplayMode": "all",
                        "globalDisplayTextTriggeredEmpty": "OK",
                        "globalOperatorName": "avg",
                        "globalUnitFormat": "short",
                        "gradientEnabled": true,
                        "hexagonSortByDirection": "asc",
                        "hexagonSortByField": "name",
                        "maxMetrics": 0,
                        "polygonBorderColor": "black",
                        "polygonBorderSize": 2,
                        "polygonGlobalFillColor": "#0a50a1",
                        "radius": "",
                        "radiusAutoSize": true,
                        "rowAutoSize": true,
                        "rows": "",
                        "shape": "hexagon_pointed_top",
                        "tooltipDisplayMode": "all",
                        "tooltipDisplayTextTriggeredEmpty": "OK",
                        "tooltipFontSize": 12,
                        "tooltipFontType": "Roboto",
                        "tooltipPrimarySortDirection": "desc",
                        "tooltipPrimarySortField": "thresholdLevel",
                        "tooltipSecondarySortDirection": "desc",
                        "tooltipSecondarySortField": "value",
                        "tooltipTimestampEnabled": true
                    },
                    "savedComposites": [],
                    "savedOverrides": [
                        {
                            "colors": [
                                "#299c46",
                                "#e5ac0e",
                                "#bf1b00",
                                "#ffffff"
                            ],
                            "decimals": 0,
                            "enabled": true,
                            "label": "OVERRIDE 1",
                            "metricName": "/.*/",
                            "operatorName": "current",
                            "prefix": "",
                            "sanitizeURLEnabled": false,
                            "scaledDecimals": null,
                            "suffix": "",
                            "thresholds": [
                                {
                                    "color": "#299c46",
                                    "state": 0,
                                    "value": 10
                                },
                                {
                                    "color": "#e5ac0e",
                                    "state": 1,
                                    "value": 30
                                },
                                {
                                    "color": "#bf1b00",
                                    "state": 2,
                                    "value": 50
                                }
                            ],
                            "unitFormat": "short"
                        }
                    ],
                    "targets": [
                        {
                            "expr": "avg by (service_name) (mongodb_mongod_metrics_cursor_open{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"total\"} or\nmongodb_mongod_cursors{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"total\"} or\nmongodb_mongos_metrics_cursor_open{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"total\"} or \nmongodb_mongos_cursors{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"total\"})",
                            "format": "time_series",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}}",
                            "refId": "A"
                        }
                    ],
                    "timeFrom": null,
                    "timeShift": null,
                    "title": "Total Cursors",
                    "type": "grafana-polystat-panel"
                },
                {
                    "aliasColors": {},
                    "bars": false,
                    "dashLength": 10,
                    "dashes": false,
                    "datasource": null,
                    "decimals": 2,
                    "description": "Open cursors that are in pinned state.",
                    "editable": true,
                    "error": false,
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
                        "y": 18
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
                    "lines": false,
                    "linewidth": 2,
                    "links": [],
                    "nullPointMode": "null",
                    "options": {
                        "dataLinks": [
                            {
                                "targetBlank": true,
                                "title": "MongoDB Instance Summary - ${__series.name}",
                                "url": "/graph/d/mongodb-instance-summary/mongodb-instance-summary?var-service_name=${__series.name}&$__url_time_range"
                            }
                        ]
                    },
                    "paceLength": 10,
                    "percentage": false,
                    "pointradius": 1,
                    "points": true,
                    "renderer": "flot",
                    "seriesOverrides": [
                        {
                            "alias": "Avg",
                            "color": "#C4162A",
                            "fill": 0,
                            "legend": false,
                            "lines": true,
                            "points": false
                        }
                    ],
                    "spaceLength": 10,
                    "stack": false,
                    "steppedLine": false,
                    "targets": [
                        {
                            "expr": "topk(5,avg by (service_name) ((mongodb_mongod_metrics_cursor_open{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"pinned\"} or\nmongodb_mongod_cursors{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"pinned\"} or\nmongodb_mongos_metrics_cursor_open{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"pinned\"} or \nmongodb_mongos_cursors{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"pinned\"})))",
                            "format": "time_series",
                            "hide": false,
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}} ",
                            "refId": "J",
                            "step": 300
                        },
                        {
                            "expr": "avg(mongodb_mongod_metrics_cursor_open{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"pinned\"} or\nmongodb_mongod_cursors{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"pinned\"} or\nmongodb_mongos_metrics_cursor_open{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"pinned\"} or \nmongodb_mongos_cursors{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"pinned\"})",
                            "format": "time_series",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "Avg",
                            "refId": "A"
                        }
                    ],
                    "thresholds": [],
                    "timeFrom": null,
                    "timeRegions": [],
                    "timeShift": null,
                    "title": "Pinned Cursors ",
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
                    "colors": [
                        "#299c46",
                        "rgba(237, 129, 40, 0.89)",
                        "#d44a3a"
                    ],
                    "datasource": null,
                    "description": "",
                    "gridPos": {
                        "h": 8,
                        "w": 12,
                        "x": 12,
                        "y": 18
                    },
                    "id": 144,
                    "links": [],
                    "options": {},
                    "polystat": {
                        "animationSpeed": 2500,
                        "columnAutoSize": true,
                        "columns": "",
                        "displayLimit": 100,
                        "fontAutoScale": true,
                        "fontSize": 12,
                        "fontType": "Roboto",
                        "globalDecimals": 2,
                        "globalDisplayMode": "all",
                        "globalDisplayTextTriggeredEmpty": "OK",
                        "globalOperatorName": "avg",
                        "globalUnitFormat": "short",
                        "gradientEnabled": true,
                        "hexagonSortByDirection": "asc",
                        "hexagonSortByField": "name",
                        "maxMetrics": 0,
                        "polygonBorderColor": "black",
                        "polygonBorderSize": 2,
                        "polygonGlobalFillColor": "#0a50a1",
                        "radius": "",
                        "radiusAutoSize": true,
                        "rowAutoSize": true,
                        "rows": "",
                        "shape": "hexagon_pointed_top",
                        "tooltipDisplayMode": "all",
                        "tooltipDisplayTextTriggeredEmpty": "OK",
                        "tooltipFontSize": 12,
                        "tooltipFontType": "Roboto",
                        "tooltipPrimarySortDirection": "desc",
                        "tooltipPrimarySortField": "thresholdLevel",
                        "tooltipSecondarySortDirection": "desc",
                        "tooltipSecondarySortField": "value",
                        "tooltipTimestampEnabled": true
                    },
                    "savedComposites": [],
                    "savedOverrides": [
                        {
                            "colors": [
                                "#299c46",
                                "#e5ac0e",
                                "#bf1b00",
                                "#ffffff"
                            ],
                            "decimals": 0,
                            "enabled": true,
                            "label": "OVERRIDE 1",
                            "metricName": "/.*/",
                            "operatorName": "current",
                            "prefix": "",
                            "sanitizeURLEnabled": false,
                            "scaledDecimals": null,
                            "suffix": "",
                            "thresholds": [
                                {
                                    "color": "#299c46",
                                    "state": 0,
                                    "value": 10
                                },
                                {
                                    "color": "#e5ac0e",
                                    "state": 1,
                                    "value": 30
                                },
                                {
                                    "color": "#bf1b00",
                                    "state": 2,
                                    "value": 50
                                }
                            ],
                            "unitFormat": "short"
                        }
                    ],
                    "targets": [
                        {
                            "expr": "avg by (service_name) (mongodb_mongod_metrics_cursor_open{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"pinned\"} or\nmongodb_mongod_cursors{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"pinned\"} or\nmongodb_mongos_metrics_cursor_open{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"pinned\"} or \nmongodb_mongos_cursors{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"pinned\"})",
                            "format": "time_series",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}}",
                            "refId": "A"
                        }
                    ],
                    "timeFrom": null,
                    "timeShift": null,
                    "title": "Pinned Cursors",
                    "type": "grafana-polystat-panel"
                },
                {
                    "aliasColors": {},
                    "bars": false,
                    "dashLength": 10,
                    "dashes": false,
                    "datasource": null,
                    "decimals": 2,
                    "description": "noTimeout cursors currently open",
                    "editable": true,
                    "error": false,
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
                        "y": 26
                    },
                    "height": "250px",
                    "hiddenSeries": false,
                    "id": 141,
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
                    "lines": false,
                    "linewidth": 2,
                    "links": [],
                    "nullPointMode": "null",
                    "options": {
                        "dataLinks": [
                            {
                                "targetBlank": true,
                                "title": "MongoDB Instance Summary - ${__series.name}",
                                "url": "/graph/d/mongodb-instance-summary/mongodb-instance-summary?var-service_name=${__series.name}&$__url_time_range"
                            }
                        ]
                    },
                    "paceLength": 10,
                    "percentage": false,
                    "pointradius": 1,
                    "points": true,
                    "renderer": "flot",
                    "seriesOverrides": [
                        {
                            "alias": "Avg",
                            "color": "#C4162A",
                            "fill": 0,
                            "legend": false,
                            "lines": true,
                            "points": false
                        }
                    ],
                    "spaceLength": 10,
                    "stack": false,
                    "steppedLine": false,
                    "targets": [
                        {
                            "expr": "topk(5,avg by (service_name) ((mongodb_mongod_metrics_cursor_open{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"noTimeout\"} or\nmongodb_mongod_cursors{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"noTimeout\"} or\nmongodb_mongos_metrics_cursor_open{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"noTimeout\"} or \nmongodb_mongos_cursors{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"noTimeout\"})))",
                            "format": "time_series",
                            "hide": false,
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}} ",
                            "refId": "J",
                            "step": 300
                        },
                        {
                            "expr": "avg(mongodb_mongod_metrics_cursor_open{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"noTimeout\"} or\nmongodb_mongod_cursors{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"noTimeout\"} or\nmongodb_mongos_metrics_cursor_open{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"noTimeout\"} or \nmongodb_mongos_cursors{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"noTimeout\"})",
                            "format": "time_series",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "Avg",
                            "refId": "A"
                        }
                    ],
                    "thresholds": [],
                    "timeFrom": null,
                    "timeRegions": [],
                    "timeShift": null,
                    "title": "Top 5 noTimeout Cursors",
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
                            "show": true
                        }
                    ],
                    "yaxis": {
                        "align": false,
                        "alignLevel": null
                    }
                },
                {
                    "colors": [
                        "#299c46",
                        "rgba(237, 129, 40, 0.89)",
                        "#d44a3a"
                    ],
                    "datasource": null,
                    "description": "",
                    "gridPos": {
                        "h": 8,
                        "w": 12,
                        "x": 12,
                        "y": 26
                    },
                    "id": 145,
                    "links": [],
                    "options": {},
                    "polystat": {
                        "animationSpeed": 2500,
                        "columnAutoSize": true,
                        "columns": "",
                        "displayLimit": 100,
                        "fontAutoScale": true,
                        "fontSize": 12,
                        "fontType": "Roboto",
                        "globalDecimals": 2,
                        "globalDisplayMode": "all",
                        "globalDisplayTextTriggeredEmpty": "OK",
                        "globalOperatorName": "avg",
                        "globalUnitFormat": "short",
                        "gradientEnabled": true,
                        "hexagonSortByDirection": "asc",
                        "hexagonSortByField": "name",
                        "maxMetrics": 0,
                        "polygonBorderColor": "black",
                        "polygonBorderSize": 2,
                        "polygonGlobalFillColor": "#0a50a1",
                        "radius": "",
                        "radiusAutoSize": true,
                        "rowAutoSize": true,
                        "rows": "",
                        "shape": "hexagon_pointed_top",
                        "tooltipDisplayMode": "all",
                        "tooltipDisplayTextTriggeredEmpty": "OK",
                        "tooltipFontSize": 12,
                        "tooltipFontType": "Roboto",
                        "tooltipPrimarySortDirection": "desc",
                        "tooltipPrimarySortField": "thresholdLevel",
                        "tooltipSecondarySortDirection": "desc",
                        "tooltipSecondarySortField": "value",
                        "tooltipTimestampEnabled": true
                    },
                    "savedComposites": [],
                    "savedOverrides": [
                        {
                            "colors": [
                                "#299c46",
                                "#e5ac0e",
                                "#bf1b00",
                                "#ffffff"
                            ],
                            "decimals": 0,
                            "enabled": true,
                            "label": "OVERRIDE 1",
                            "metricName": "/.*/",
                            "operatorName": "current",
                            "prefix": "",
                            "sanitizeURLEnabled": false,
                            "scaledDecimals": null,
                            "suffix": "",
                            "thresholds": [
                                {
                                    "color": "#299c46",
                                    "state": 0,
                                    "value": 10
                                },
                                {
                                    "color": "#e5ac0e",
                                    "state": 1,
                                    "value": 30
                                },
                                {
                                    "color": "#bf1b00",
                                    "state": 2,
                                    "value": 50
                                }
                            ],
                            "unitFormat": "short"
                        }
                    ],
                    "targets": [
                        {
                            "expr": "avg by (service_name) (mongodb_mongod_metrics_cursor_open{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"noTimeout\"} or\nmongodb_mongod_cursors{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"noTimeout\"} or\nmongodb_mongos_metrics_cursor_open{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"noTimeout\"} or \nmongodb_mongos_cursors{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"noTimeout\"})",
                            "format": "time_series",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}}",
                            "refId": "A"
                        }
                    ],
                    "timeFrom": null,
                    "timeShift": null,
                    "title": "noTimeout Cursors",
                    "type": "grafana-polystat-panel"
                }
            ],
            "title": "Cursors Detail",
            "type": "row"
        },
        {
            "collapsed": true,
            "datasource": null,
            "gridPos": {
                "h": 1,
                "w": 24,
                "x": 0,
                "y": 10
            },
            "id": 209,
            "panels": [
                {
                    "aliasColors": {},
                    "bars": false,
                    "dashLength": 10,
                    "dashes": false,
                    "datasource": null,
                    "decimals": 2,
                    "description": "Average latency of (other) command operations",
                    "editable": true,
                    "error": false,
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
                        "y": 11
                    },
                    "height": "250px",
                    "hiddenSeries": false,
                    "id": 213,
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
                    "lines": false,
                    "linewidth": 2,
                    "links": [],
                    "nullPointMode": "null",
                    "options": {
                        "dataLinks": [
                            {
                                "targetBlank": true,
                                "title": "MongoDB Instance Summary - ${__series.name}",
                                "url": "/graph/d/mongodb-instance-summary/mongodb-instance-summary?var-service_name=${__series.name}&$__url_time_range"
                            }
                        ]
                    },
                    "paceLength": 10,
                    "percentage": false,
                    "pointradius": 1,
                    "points": true,
                    "renderer": "flot",
                    "seriesOverrides": [
                        {
                            "alias": "Avg",
                            "color": "#C4162A",
                            "fill": 0,
                            "legend": false,
                            "lines": true,
                            "points": false
                        }
                    ],
                    "spaceLength": 10,
                    "stack": false,
                    "steppedLine": false,
                    "targets": [
                        {
                            "expr": "topk(5,avg by (service_name) ((rate(mongodb_mongod_op_latencies_latency_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"command\"}[$interval]) / \n(rate(mongodb_mongod_op_latencies_ops_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"command\"}[$interval]) > 0) or\nirate(mongodb_mongod_op_latencies_latency_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"command\"}[5m]) / \n(irate(mongodb_mongod_op_latencies_ops_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"command\"}[5m]) > 0)\n)))",
                            "format": "time_series",
                            "hide": false,
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}} ",
                            "refId": "J",
                            "step": 300
                        },
                        {
                            "expr": "avg(rate(mongodb_mongod_op_latencies_latency_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"command\"}[$interval]) / \n(rate(mongodb_mongod_op_latencies_ops_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"command\"}[$interval]) > 0) or\nirate(mongodb_mongod_op_latencies_latency_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"command\"}[5m]) / \n(irate(mongodb_mongod_op_latencies_ops_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"command\"}[5m]) > 0)\n)",
                            "format": "time_series",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "Avg",
                            "refId": "A"
                        }
                    ],
                    "thresholds": [],
                    "timeFrom": null,
                    "timeRegions": [],
                    "timeShift": null,
                    "title": "Top 5 Command Latency",
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
                    "colors": [
                        "#299c46",
                        "rgba(237, 129, 40, 0.89)",
                        "#d44a3a"
                    ],
                    "datasource": null,
                    "description": "",
                    "gridPos": {
                        "h": 8,
                        "w": 12,
                        "x": 12,
                        "y": 11
                    },
                    "id": 211,
                    "links": [],
                    "options": {},
                    "polystat": {
                        "animationSpeed": 2500,
                        "columnAutoSize": true,
                        "columns": "",
                        "displayLimit": 100,
                        "fontAutoScale": true,
                        "fontSize": 12,
                        "fontType": "Roboto",
                        "globalDecimals": 2,
                        "globalDisplayMode": "all",
                        "globalDisplayTextTriggeredEmpty": "OK",
                        "globalOperatorName": "avg",
                        "globalUnitFormat": "short",
                        "gradientEnabled": true,
                        "hexagonSortByDirection": "asc",
                        "hexagonSortByField": "name",
                        "maxMetrics": 0,
                        "polygonBorderColor": "black",
                        "polygonBorderSize": 2,
                        "polygonGlobalFillColor": "#0a50a1",
                        "radius": "",
                        "radiusAutoSize": true,
                        "rowAutoSize": true,
                        "rows": "",
                        "shape": "hexagon_pointed_top",
                        "tooltipDisplayMode": "all",
                        "tooltipDisplayTextTriggeredEmpty": "OK",
                        "tooltipFontSize": 12,
                        "tooltipFontType": "Roboto",
                        "tooltipPrimarySortDirection": "desc",
                        "tooltipPrimarySortField": "thresholdLevel",
                        "tooltipSecondarySortDirection": "desc",
                        "tooltipSecondarySortField": "value",
                        "tooltipTimestampEnabled": true
                    },
                    "savedComposites": [],
                    "savedOverrides": [
                        {
                            "colors": [
                                "#299c46",
                                "#e5ac0e",
                                "#bf1b00",
                                "#ffffff"
                            ],
                            "decimals": 2,
                            "enabled": true,
                            "label": "OVERRIDE 1",
                            "metricName": "/.*/",
                            "operatorName": "current",
                            "prefix": "",
                            "sanitizeURLEnabled": false,
                            "scaledDecimals": null,
                            "suffix": "",
                            "thresholds": [
                                {
                                    "color": "#1F60C4",
                                    "state": 0,
                                    "value": 10
                                }
                            ],
                            "unitFormat": "\u00b5s"
                        }
                    ],
                    "targets": [
                        {
                            "expr": "avg by (service_name) (rate(mongodb_mongod_op_latencies_latency_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"command\"}[$interval]) / \n(rate(mongodb_mongod_op_latencies_ops_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"command\"}[$interval]) > 0) or\nirate(mongodb_mongod_op_latencies_latency_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"command\"}[5m]) / \n(irate(mongodb_mongod_op_latencies_ops_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"command\"}[5m]) > 0))\n",
                            "format": "time_series",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}}",
                            "refId": "A"
                        }
                    ],
                    "timeFrom": null,
                    "timeShift": null,
                    "title": "Command Latency",
                    "type": "grafana-polystat-panel"
                },
                {
                    "aliasColors": {},
                    "bars": false,
                    "dashLength": 10,
                    "dashes": false,
                    "datasource": null,
                    "decimals": 2,
                    "description": "Average latency of read operations",
                    "editable": true,
                    "error": false,
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
                        "y": 19
                    },
                    "height": "250px",
                    "hiddenSeries": false,
                    "id": 210,
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
                    "lines": false,
                    "linewidth": 2,
                    "links": [],
                    "nullPointMode": "null",
                    "options": {
                        "dataLinks": [
                            {
                                "targetBlank": true,
                                "title": "MongoDB Instance Summary - ${__series.name}",
                                "url": "/graph/d/mongodb-instance-summary/mongodb-instance-summary?var-service_name=${__series.name}&$__url_time_range"
                            }
                        ]
                    },
                    "paceLength": 10,
                    "percentage": false,
                    "pointradius": 1,
                    "points": true,
                    "renderer": "flot",
                    "seriesOverrides": [
                        {
                            "alias": "Avg",
                            "color": "#C4162A",
                            "fill": 0,
                            "legend": false,
                            "lines": true,
                            "points": false
                        }
                    ],
                    "spaceLength": 10,
                    "stack": false,
                    "steppedLine": false,
                    "targets": [
                        {
                            "expr": "topk(5,avg by (service_name) ((rate(mongodb_mongod_op_latencies_latency_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"read\"}[$interval]) / \n(rate(mongodb_mongod_op_latencies_ops_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"read\"}[$interval]) > 0) or\nirate(mongodb_mongod_op_latencies_latency_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"read\"}[5m]) / \n(irate(mongodb_mongod_op_latencies_ops_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"read\"}[5m]) > 0)\n)))",
                            "format": "time_series",
                            "hide": false,
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}} ",
                            "refId": "J",
                            "step": 300
                        },
                        {
                            "expr": "avg(rate(mongodb_mongod_op_latencies_latency_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"read\"}[$interval]) / \n(rate(mongodb_mongod_op_latencies_ops_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"read\"}[$interval]) > 0) or\nirate(mongodb_mongod_op_latencies_latency_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"read\"}[5m]) / \n(irate(mongodb_mongod_op_latencies_ops_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"read\"}[5m]) > 0)\n)",
                            "format": "time_series",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "Avg",
                            "refId": "A"
                        }
                    ],
                    "thresholds": [],
                    "timeFrom": null,
                    "timeRegions": [],
                    "timeShift": null,
                    "title": "Top 5 Read Latency",
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
                    "colors": [
                        "#299c46",
                        "rgba(237, 129, 40, 0.89)",
                        "#d44a3a"
                    ],
                    "datasource": null,
                    "description": "",
                    "gridPos": {
                        "h": 8,
                        "w": 12,
                        "x": 12,
                        "y": 19
                    },
                    "id": 215,
                    "links": [],
                    "options": {},
                    "polystat": {
                        "animationSpeed": 2500,
                        "columnAutoSize": true,
                        "columns": "",
                        "displayLimit": 100,
                        "fontAutoScale": true,
                        "fontSize": 12,
                        "fontType": "Roboto",
                        "globalDecimals": 2,
                        "globalDisplayMode": "all",
                        "globalDisplayTextTriggeredEmpty": "OK",
                        "globalOperatorName": "avg",
                        "globalUnitFormat": "short",
                        "gradientEnabled": true,
                        "hexagonSortByDirection": "asc",
                        "hexagonSortByField": "name",
                        "maxMetrics": 0,
                        "polygonBorderColor": "black",
                        "polygonBorderSize": 2,
                        "polygonGlobalFillColor": "#0a50a1",
                        "radius": "",
                        "radiusAutoSize": true,
                        "rowAutoSize": true,
                        "rows": "",
                        "shape": "hexagon_pointed_top",
                        "tooltipDisplayMode": "all",
                        "tooltipDisplayTextTriggeredEmpty": "OK",
                        "tooltipFontSize": 12,
                        "tooltipFontType": "Roboto",
                        "tooltipPrimarySortDirection": "desc",
                        "tooltipPrimarySortField": "thresholdLevel",
                        "tooltipSecondarySortDirection": "desc",
                        "tooltipSecondarySortField": "value",
                        "tooltipTimestampEnabled": true
                    },
                    "savedComposites": [],
                    "savedOverrides": [
                        {
                            "colors": [
                                "#299c46",
                                "#e5ac0e",
                                "#bf1b00",
                                "#ffffff"
                            ],
                            "decimals": 2,
                            "enabled": true,
                            "label": "OVERRIDE 1",
                            "metricName": "/.*/",
                            "operatorName": "current",
                            "prefix": "",
                            "sanitizeURLEnabled": false,
                            "scaledDecimals": null,
                            "suffix": "",
                            "thresholds": [
                                {
                                    "color": "#1F60C4",
                                    "state": 0,
                                    "value": 10
                                }
                            ],
                            "unitFormat": "\u00b5s"
                        }
                    ],
                    "targets": [
                        {
                            "expr": "avg by (service_name) (rate(mongodb_mongod_op_latencies_latency_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"read\"}[$interval]) / \n(rate(mongodb_mongod_op_latencies_ops_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"read\"}[$interval]) > 0) or\nirate(mongodb_mongod_op_latencies_latency_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"read\"}[5m]) / \n(irate(mongodb_mongod_op_latencies_ops_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"read\"}[5m]) > 0))\n",
                            "format": "time_series",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}}",
                            "refId": "A"
                        }
                    ],
                    "timeFrom": null,
                    "timeShift": null,
                    "title": "Read Latency",
                    "type": "grafana-polystat-panel"
                },
                {
                    "aliasColors": {},
                    "bars": false,
                    "dashLength": 10,
                    "dashes": false,
                    "datasource": null,
                    "decimals": 2,
                    "description": "Average latency of write operations",
                    "editable": true,
                    "error": false,
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
                    "id": 212,
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
                    "lines": false,
                    "linewidth": 2,
                    "links": [],
                    "nullPointMode": "null as zero",
                    "options": {
                        "dataLinks": [
                            {
                                "targetBlank": true,
                                "title": "MongoDB Instance Summary - ${__series.name}",
                                "url": "/graph/d/mongodb-instance-summary/mongodb-instance-summary?var-service_name=${__series.name}&$__url_time_range"
                            }
                        ]
                    },
                    "paceLength": 10,
                    "percentage": false,
                    "pointradius": 1,
                    "points": true,
                    "renderer": "flot",
                    "seriesOverrides": [
                        {
                            "alias": "Avg",
                            "color": "#C4162A",
                            "fill": 0,
                            "legend": false,
                            "lines": true,
                            "points": false
                        }
                    ],
                    "spaceLength": 10,
                    "stack": false,
                    "steppedLine": false,
                    "targets": [
                        {
                            "expr": "topk(5,avg by (service_name) ((rate(mongodb_mongod_op_latencies_latency_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"write\"}[$interval]) / \n(rate(mongodb_mongod_op_latencies_ops_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"write\"}[$interval]) > 0) or\nirate(mongodb_mongod_op_latencies_latency_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"write\"}[5m]) / \n(irate(mongodb_mongod_op_latencies_ops_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"write\"}[5m]) > 0)\n)))",
                            "format": "time_series",
                            "hide": false,
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}} ",
                            "refId": "J",
                            "step": 300
                        },
                        {
                            "expr": "avg(rate(mongodb_mongod_op_latencies_latency_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"write\"}[$interval]) / \n(rate(mongodb_mongod_op_latencies_ops_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"write\"}[$interval]) > 0) or\nirate(mongodb_mongod_op_latencies_latency_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"write\"}[5m]) / \n(irate(mongodb_mongod_op_latencies_ops_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"write\"}[5m]) > 0)\n)",
                            "format": "time_series",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "Avg",
                            "refId": "A"
                        }
                    ],
                    "thresholds": [],
                    "timeFrom": null,
                    "timeRegions": [],
                    "timeShift": null,
                    "title": "Top 5 Write Latency",
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
                    "colors": [
                        "#299c46",
                        "rgba(237, 129, 40, 0.89)",
                        "#d44a3a"
                    ],
                    "datasource": null,
                    "description": "",
                    "gridPos": {
                        "h": 8,
                        "w": 12,
                        "x": 12,
                        "y": 27
                    },
                    "id": 214,
                    "links": [],
                    "options": {},
                    "polystat": {
                        "animationSpeed": 2500,
                        "columnAutoSize": true,
                        "columns": "",
                        "displayLimit": 100,
                        "fontAutoScale": true,
                        "fontSize": 12,
                        "fontType": "Roboto",
                        "globalDecimals": 2,
                        "globalDisplayMode": "all",
                        "globalDisplayTextTriggeredEmpty": "OK",
                        "globalOperatorName": "avg",
                        "globalUnitFormat": "short",
                        "gradientEnabled": true,
                        "hexagonSortByDirection": "asc",
                        "hexagonSortByField": "name",
                        "maxMetrics": 0,
                        "polygonBorderColor": "black",
                        "polygonBorderSize": 2,
                        "polygonGlobalFillColor": "#0a50a1",
                        "radius": "",
                        "radiusAutoSize": true,
                        "rowAutoSize": true,
                        "rows": "",
                        "shape": "hexagon_pointed_top",
                        "tooltipDisplayMode": "all",
                        "tooltipDisplayTextTriggeredEmpty": "OK",
                        "tooltipFontSize": 12,
                        "tooltipFontType": "Roboto",
                        "tooltipPrimarySortDirection": "desc",
                        "tooltipPrimarySortField": "thresholdLevel",
                        "tooltipSecondarySortDirection": "desc",
                        "tooltipSecondarySortField": "value",
                        "tooltipTimestampEnabled": true
                    },
                    "savedComposites": [],
                    "savedOverrides": [
                        {
                            "colors": [
                                "#299c46",
                                "#e5ac0e",
                                "#bf1b00",
                                "#ffffff"
                            ],
                            "decimals": 2,
                            "enabled": true,
                            "label": "OVERRIDE 1",
                            "metricName": "/.*/",
                            "operatorName": "current",
                            "prefix": "",
                            "sanitizeURLEnabled": false,
                            "scaledDecimals": null,
                            "suffix": "",
                            "thresholds": [
                                {
                                    "color": "#1F60C4",
                                    "state": 0,
                                    "value": 10
                                }
                            ],
                            "unitFormat": "\u00b5s"
                        }
                    ],
                    "targets": [
                        {
                            "expr": "avg by (service_name) (rate(mongodb_mongod_op_latencies_latency_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"write\"}[$interval]) / \n(rate(mongodb_mongod_op_latencies_ops_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"write\"}[$interval]) > 0) or\nirate(mongodb_mongod_op_latencies_latency_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"write\"}[5m]) / \n(irate(mongodb_mongod_op_latencies_ops_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"write\"}[5m]) > 0))\n",
                            "format": "time_series",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}}",
                            "refId": "A"
                        }
                    ],
                    "timeFrom": null,
                    "timeShift": null,
                    "title": "Write Latency",
                    "type": "grafana-polystat-panel"
                }
            ],
            "title": "Latency Detail",
            "type": "row"
        },
        {
            "collapsed": false,
            "datasource": null,
            "gridPos": {
                "h": 1,
                "w": 24,
                "x": 0,
                "y": 11
            },
            "id": 94,
            "panels": [],
            "type": "row"
        },
        {
            "cacheTimeout": null,
            "colorBackground": false,
            "colorValue": false,
            "colors": [
                "#d44a3a",
                "#3274D9",
                "#299c46"
            ],
            "datasource": null,
            "decimals": 2,
            "description": "Ratio of index entries scanned / number of documents returned",
            "fieldConfig": {
                "defaults": {},
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
                "w": 6,
                "x": 0,
                "y": 12
            },
            "id": 97,
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
                "show": true
            },
            "tableColumn": "",
            "targets": [
                {
                    "expr": "min(rate(mongodb_mongod_metrics_query_executor_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"scanned\"}[$interval]) / ignoring(state) \nrate(mongodb_mongod_metrics_document_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"returned\"}[$interval])\nor\n(irate(mongodb_mongod_metrics_query_executor_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"scanned\"}[5m]) / ignoring(state) \nirate(mongodb_mongod_metrics_document_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"returned\"}[5m]))\n)",
                    "format": "time_series",
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "",
                    "refId": "A"
                }
            ],
            "thresholds": "",
            "timeFrom": null,
            "timeShift": null,
            "title": "Min Index Scanned Ratio",
            "type": "singlestat",
            "valueFontSize": "50%",
            "valueMaps": [
                {
                    "op": "=",
                    "text": "0",
                    "value": "null"
                },
                {
                    "op": "=",
                    "text": "0",
                    "value": "N/A"
                }
            ],
            "valueName": "min"
        },
        {
            "cacheTimeout": null,
            "colorBackground": false,
            "colorValue": false,
            "colors": [
                "#d44a3a",
                "#3274D9",
                "#299c46"
            ],
            "datasource": null,
            "decimals": 2,
            "description": "Ratio of number of documents returned / index entries scanned",
            "fieldConfig": {
                "defaults": {},
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
                "w": 6,
                "x": 6,
                "y": 12
            },
            "id": 112,
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
                "show": true
            },
            "tableColumn": "",
            "targets": [
                {
                    "expr": "max(rate(mongodb_mongod_metrics_query_executor_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"scanned\"}[$interval]) / ignoring(state) \nrate(mongodb_mongod_metrics_document_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"returned\"}[$interval])\nor\n(irate(mongodb_mongod_metrics_query_executor_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"scanned\"}[5m]) / ignoring(state) \nirate(mongodb_mongod_metrics_document_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"returned\"}[5m]))\n)",
                    "format": "time_series",
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "",
                    "refId": "A"
                }
            ],
            "thresholds": "",
            "timeFrom": null,
            "timeShift": null,
            "title": "Max Index Scanned Ratio",
            "type": "singlestat",
            "valueFontSize": "50%",
            "valueMaps": [
                {
                    "op": "=",
                    "text": "0",
                    "value": "null"
                },
                {
                    "op": "=",
                    "text": "0",
                    "value": "N/A"
                }
            ],
            "valueName": "max"
        },
        {
            "cacheTimeout": null,
            "colorBackground": false,
            "colorPrefix": false,
            "colorValue": false,
            "colors": [
                "#299c46",
                "#3274D9",
                "#d44a3a"
            ],
            "datasource": null,
            "decimals": 2,
            "description": "Ratio of index entries scanned compared divided by sum of index entries scanned plus whole documents scanned.",
            "fieldConfig": {
                "defaults": {},
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
                "w": 6,
                "x": 12,
                "y": 12
            },
            "id": 86,
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
                "show": true
            },
            "tableColumn": "",
            "targets": [
                {
                    "expr": "min(rate(mongodb_mongod_metrics_document_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"returned\"}[$interval])/ ignoring(state)\nrate(mongodb_mongod_metrics_query_executor_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"scanned_objects\"}[$interval])\nor\nirate(mongodb_mongod_metrics_document_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"returned\"}[5m])/ ignoring(state)\nirate(mongodb_mongod_metrics_query_executor_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"scanned_objects\"}[5m]))",
                    "format": "time_series",
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "",
                    "refId": "A"
                }
            ],
            "thresholds": "",
            "timeFrom": null,
            "timeShift": null,
            "title": "Min Document Scanned Ratio",
            "type": "singlestat",
            "valueFontSize": "50%",
            "valueMaps": [
                {
                    "op": "=",
                    "text": "0",
                    "value": "null"
                }
            ],
            "valueName": "min"
        },
        {
            "cacheTimeout": null,
            "colorBackground": false,
            "colorPrefix": false,
            "colorValue": false,
            "colors": [
                "#299c46",
                "#3274D9",
                "#d44a3a"
            ],
            "datasource": null,
            "decimals": 2,
            "fieldConfig": {
                "defaults": {},
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
                "w": 6,
                "x": 18,
                "y": 12
            },
            "id": 84,
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
                "show": true
            },
            "tableColumn": "",
            "targets": [
                {
                    "expr": "max(rate(mongodb_mongod_metrics_document_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"returned\"}[$interval])/ ignoring(state)\nrate(mongodb_mongod_metrics_query_executor_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"scanned_objects\"}[$interval])\nor\nirate(mongodb_mongod_metrics_document_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"returned\"}[5m])/ ignoring(state)\nirate(mongodb_mongod_metrics_query_executor_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"scanned_objects\"}[5m]))",
                    "format": "time_series",
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "",
                    "refId": "A"
                }
            ],
            "thresholds": "",
            "timeFrom": null,
            "timeShift": null,
            "title": "Max Document Scanned Ratio",
            "type": "singlestat",
            "valueFontSize": "50%",
            "valueMaps": [
                {
                    "op": "=",
                    "text": "0",
                    "value": "null"
                }
            ],
            "valueName": "max"
        },
        {
            "collapsed": true,
            "datasource": null,
            "gridPos": {
                "h": 1,
                "w": 24,
                "x": 0,
                "y": 14
            },
            "id": 126,
            "panels": [
                {
                    "aliasColors": {},
                    "bars": false,
                    "dashLength": 10,
                    "dashes": false,
                    "datasource": null,
                    "decimals": 2,
                    "description": "",
                    "editable": true,
                    "error": false,
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
                        "y": 15
                    },
                    "height": "250px",
                    "hiddenSeries": false,
                    "id": 63,
                    "interval": "",
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
                    "lines": false,
                    "linewidth": 2,
                    "links": [],
                    "nullPointMode": "null as zero",
                    "options": {
                        "dataLinks": [
                            {
                                "targetBlank": true,
                                "title": "MongoDB Instance Summary - ${__series.name}",
                                "url": "/graph/d/mongodb-instance-summary/mongodb-instance-summary?var-service_name=${__series.name}&$__url_time_range"
                            }
                        ]
                    },
                    "paceLength": 10,
                    "percentage": true,
                    "pointradius": 1,
                    "points": true,
                    "renderer": "flot",
                    "seriesOverrides": [
                        {
                            "alias": "Avg",
                            "color": "#C4162A",
                            "fill": 0,
                            "legend": false,
                            "lines": true,
                            "points": false
                        }
                    ],
                    "spaceLength": 10,
                    "stack": false,
                    "steppedLine": false,
                    "targets": [
                        {
                            "expr": "topk(5,avg by (service_name) ((rate(mongodb_mongod_metrics_query_executor_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"scanned\"}[$interval]) / ignoring(state) \nrate(mongodb_mongod_metrics_document_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"returned\"}[$interval])\nor\n(irate(mongodb_mongod_metrics_query_executor_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"scanned\"}[5m]) / ignoring(state) \nirate(mongodb_mongod_metrics_document_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"returned\"}[5m]))\n)))",
                            "format": "time_series",
                            "hide": false,
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}}",
                            "refId": "J",
                            "step": 300
                        },
                        {
                            "expr": "avg(rate(mongodb_mongod_metrics_query_executor_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"scanned\"}[$interval]) / ignoring(state) \nrate(mongodb_mongod_metrics_document_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"returned\"}[$interval])\nor\n(irate(mongodb_mongod_metrics_query_executor_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"scanned\"}[5m]) / ignoring(state) \nirate(mongodb_mongod_metrics_document_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"returned\"}[5m]))\n)",
                            "format": "time_series",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "Avg",
                            "refId": "A"
                        }
                    ],
                    "thresholds": [],
                    "timeFrom": null,
                    "timeRegions": [],
                    "timeShift": null,
                    "title": "Top 5 Index Scan Ratios",
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
                    "colors": [
                        "#299c46",
                        "rgba(237, 129, 40, 0.89)",
                        "#d44a3a"
                    ],
                    "datasource": null,
                    "gridPos": {
                        "h": 8,
                        "w": 12,
                        "x": 12,
                        "y": 15
                    },
                    "id": 191,
                    "links": [],
                    "options": {},
                    "polystat": {
                        "animationSpeed": 2500,
                        "columnAutoSize": true,
                        "columns": "",
                        "displayLimit": 100,
                        "fontAutoScale": true,
                        "fontSize": 12,
                        "fontType": "Roboto",
                        "globalDecimals": 2,
                        "globalDisplayMode": "all",
                        "globalDisplayTextTriggeredEmpty": "OK",
                        "globalOperatorName": "current",
                        "globalUnitFormat": "short",
                        "gradientEnabled": true,
                        "hexagonSortByDirection": "asc",
                        "hexagonSortByField": "name",
                        "maxMetrics": 0,
                        "polygonBorderColor": "black",
                        "polygonBorderSize": 2,
                        "polygonGlobalFillColor": "#0a50a1",
                        "radius": "",
                        "radiusAutoSize": true,
                        "rowAutoSize": true,
                        "rows": "",
                        "shape": "hexagon_pointed_top",
                        "tooltipDisplayMode": "all",
                        "tooltipDisplayTextTriggeredEmpty": "OK",
                        "tooltipFontSize": 12,
                        "tooltipFontType": "Roboto",
                        "tooltipPrimarySortDirection": "desc",
                        "tooltipPrimarySortField": "thresholdLevel",
                        "tooltipSecondarySortDirection": "desc",
                        "tooltipSecondarySortField": "value",
                        "tooltipTimestampEnabled": true
                    },
                    "savedComposites": [],
                    "savedOverrides": [
                        {
                            "colors": [
                                "#299c46",
                                "#e5ac0e",
                                "#bf1b00",
                                "#ffffff"
                            ],
                            "decimals": 2,
                            "enabled": true,
                            "label": "OVERRIDE 1",
                            "metricName": "/.*/",
                            "operatorName": "current",
                            "prefix": "",
                            "sanitizeURLEnabled": false,
                            "scaledDecimals": null,
                            "suffix": "",
                            "thresholds": [],
                            "unitFormat": "ops"
                        }
                    ],
                    "targets": [
                        {
                            "expr": "avg by (service_name) (rate(mongodb_mongod_metrics_query_executor_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"scanned\"}[$interval]) / ignoring(state) \nrate(mongodb_mongod_metrics_document_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"returned\"}[$interval])\nor\n(irate(mongodb_mongod_metrics_query_executor_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"scanned\"}[5m]) / ignoring(state) \nirate(mongodb_mongod_metrics_document_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"returned\"}[5m])))",
                            "format": "time_series",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}}",
                            "refId": "A"
                        }
                    ],
                    "timeFrom": null,
                    "timeShift": null,
                    "title": "Index Scan Ratios",
                    "type": "grafana-polystat-panel"
                },
                {
                    "aliasColors": {},
                    "bars": false,
                    "dashLength": 10,
                    "dashes": false,
                    "datasource": null,
                    "decimals": 2,
                    "description": "",
                    "editable": true,
                    "error": false,
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
                        "y": 23
                    },
                    "height": "250px",
                    "hiddenSeries": false,
                    "id": 189,
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
                    "lines": false,
                    "linewidth": 2,
                    "links": [],
                    "nullPointMode": "null as zero",
                    "options": {
                        "dataLinks": [
                            {
                                "targetBlank": true,
                                "title": "MongoDB Instance Summary - ${__series.name}",
                                "url": "/graph/d/mongodb-instance-summary/mongodb-instance-summary?var-service_name=${__series.name}&$__url_time_range"
                            }
                        ]
                    },
                    "paceLength": 10,
                    "percentage": true,
                    "pointradius": 1,
                    "points": true,
                    "renderer": "flot",
                    "seriesOverrides": [
                        {
                            "alias": "Avg",
                            "color": "#C4162A",
                            "fill": 0,
                            "legend": false,
                            "lines": true,
                            "points": false
                        }
                    ],
                    "spaceLength": 10,
                    "stack": false,
                    "steppedLine": false,
                    "targets": [
                        {
                            "expr": "topk(5,avg by (service_name) ((rate(mongodb_mongod_metrics_document_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"returned\"}[$interval])/ ignoring(state)\nrate(mongodb_mongod_metrics_query_executor_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"scanned_objects\"}[$interval])\nor\nirate(mongodb_mongod_metrics_document_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"returned\"}[5m])/ ignoring(state)\nirate(mongodb_mongod_metrics_query_executor_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"scanned_objects\"}[5m]))))",
                            "format": "time_series",
                            "hide": false,
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}}",
                            "refId": "J",
                            "step": 300
                        },
                        {
                            "expr": "avg(rate(mongodb_mongod_metrics_document_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"returned\"}[$interval])/ ignoring(state)\nrate(mongodb_mongod_metrics_query_executor_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"scanned_objects\"}[$interval])\nor\nirate(mongodb_mongod_metrics_document_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"returned\"}[5m])/ ignoring(state)\nirate(mongodb_mongod_metrics_query_executor_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"scanned_objects\"}[5m]))",
                            "format": "time_series",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "Avg",
                            "refId": "A"
                        }
                    ],
                    "thresholds": [],
                    "timeFrom": null,
                    "timeRegions": [],
                    "timeShift": null,
                    "title": "Top 5 Document Scan Ratios",
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
                    "colors": [
                        "#299c46",
                        "rgba(237, 129, 40, 0.89)",
                        "#d44a3a"
                    ],
                    "datasource": null,
                    "gridPos": {
                        "h": 8,
                        "w": 12,
                        "x": 12,
                        "y": 23
                    },
                    "id": 190,
                    "links": [],
                    "options": {},
                    "polystat": {
                        "animationSpeed": 2500,
                        "columnAutoSize": true,
                        "columns": "",
                        "displayLimit": 100,
                        "fontAutoScale": true,
                        "fontSize": 12,
                        "fontType": "Roboto",
                        "globalDecimals": 2,
                        "globalDisplayMode": "all",
                        "globalDisplayTextTriggeredEmpty": "OK",
                        "globalOperatorName": "avg",
                        "globalUnitFormat": "short",
                        "gradientEnabled": true,
                        "hexagonSortByDirection": "asc",
                        "hexagonSortByField": "name",
                        "maxMetrics": 0,
                        "polygonBorderColor": "black",
                        "polygonBorderSize": 2,
                        "polygonGlobalFillColor": "#0a50a1",
                        "radius": "",
                        "radiusAutoSize": true,
                        "rowAutoSize": true,
                        "rows": "",
                        "shape": "hexagon_pointed_top",
                        "tooltipDisplayMode": "all",
                        "tooltipDisplayTextTriggeredEmpty": "OK",
                        "tooltipFontSize": 12,
                        "tooltipFontType": "Roboto",
                        "tooltipPrimarySortDirection": "desc",
                        "tooltipPrimarySortField": "thresholdLevel",
                        "tooltipSecondarySortDirection": "desc",
                        "tooltipSecondarySortField": "value",
                        "tooltipTimestampEnabled": true
                    },
                    "savedComposites": [],
                    "savedOverrides": [
                        {
                            "colors": [
                                "#299c46",
                                "#e5ac0e",
                                "#bf1b00",
                                "#ffffff"
                            ],
                            "decimals": 2,
                            "enabled": true,
                            "label": "OVERRIDE 1",
                            "metricName": "/.*/",
                            "operatorName": "current",
                            "prefix": "",
                            "sanitizeURLEnabled": false,
                            "scaledDecimals": null,
                            "suffix": "",
                            "thresholds": [],
                            "unitFormat": "ops"
                        }
                    ],
                    "targets": [
                        {
                            "expr": "avg by (service_name) (rate(mongodb_mongod_metrics_document_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"returned\"}[$interval])/ ignoring(state)\nrate(mongodb_mongod_metrics_query_executor_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"scanned_objects\"}[$interval])\nor\nirate(mongodb_mongod_metrics_document_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"returned\"}[5m])/ ignoring(state)\nirate(mongodb_mongod_metrics_query_executor_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"scanned_objects\"}[5m]))",
                            "format": "time_series",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}}",
                            "refId": "A"
                        }
                    ],
                    "timeFrom": null,
                    "timeShift": null,
                    "title": "Document Scan Ratios",
                    "type": "grafana-polystat-panel"
                },
                {
                    "aliasColors": {},
                    "bars": false,
                    "dashLength": 10,
                    "dashes": false,
                    "datasource": null,
                    "decimals": 2,
                    "description": "Ratio of index entries scanned compared divided by sum of index entries scanned plus whole documents scanned.",
                    "editable": true,
                    "error": false,
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
                        "y": 31
                    },
                    "height": "250px",
                    "hiddenSeries": false,
                    "id": 217,
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
                    "lines": false,
                    "linewidth": 2,
                    "links": [],
                    "nullPointMode": "null as zero",
                    "options": {
                        "dataLinks": [
                            {
                                "targetBlank": true,
                                "title": "MongoDB Instance Summary - ${__series.name}",
                                "url": "/graph/d/mongodb-instance-summary/mongodb-instance-summary?var-service_name=${__series.name}&$__url_time_range"
                            }
                        ]
                    },
                    "paceLength": 10,
                    "percentage": true,
                    "pointradius": 1,
                    "points": true,
                    "renderer": "flot",
                    "seriesOverrides": [
                        {
                            "alias": "Avg",
                            "color": "#C4162A",
                            "fill": 0,
                            "legend": false,
                            "lines": true,
                            "points": false
                        }
                    ],
                    "spaceLength": 10,
                    "stack": false,
                    "steppedLine": false,
                    "targets": [
                        {
                            "expr": "topk(5,(sum by (service_name) (mongodb_mongod_metrics_query_executor_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=~\"scanned_objects\"}) / ignoring(state)\nsum by (service_name) (mongodb_mongod_metrics_query_executor_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=~\"scanned_objects|scanned\"})))",
                            "format": "time_series",
                            "hide": false,
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}}",
                            "refId": "B"
                        },
                        {
                            "expr": "avg(sum by (service_name) (mongodb_mongod_metrics_query_executor_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=~\"scanned_objects\"}) / ignoring(state)\nsum by (service_name) (mongodb_mongod_metrics_query_executor_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=~\"scanned_objects|scanned\"}))",
                            "format": "time_series",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "Avg",
                            "refId": "C"
                        }
                    ],
                    "thresholds": [],
                    "timeFrom": null,
                    "timeRegions": [],
                    "timeShift": null,
                    "title": "Top 5 Index Filtering Effectiveness",
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
                    "colors": [
                        "#299c46",
                        "rgba(237, 129, 40, 0.89)",
                        "#d44a3a"
                    ],
                    "datasource": null,
                    "gridPos": {
                        "h": 8,
                        "w": 12,
                        "x": 12,
                        "y": 31
                    },
                    "id": 219,
                    "links": [],
                    "options": {},
                    "polystat": {
                        "animationSpeed": 2500,
                        "columnAutoSize": true,
                        "columns": "",
                        "displayLimit": 100,
                        "fontAutoScale": true,
                        "fontSize": 12,
                        "fontType": "Roboto",
                        "globalDecimals": 2,
                        "globalDisplayMode": "all",
                        "globalDisplayTextTriggeredEmpty": "OK",
                        "globalOperatorName": "avg",
                        "globalUnitFormat": "short",
                        "gradientEnabled": true,
                        "hexagonSortByDirection": "asc",
                        "hexagonSortByField": "name",
                        "maxMetrics": 0,
                        "polygonBorderColor": "black",
                        "polygonBorderSize": 2,
                        "polygonGlobalFillColor": "#0a50a1",
                        "radius": "",
                        "radiusAutoSize": true,
                        "rowAutoSize": true,
                        "rows": "",
                        "shape": "hexagon_pointed_top",
                        "tooltipDisplayMode": "all",
                        "tooltipDisplayTextTriggeredEmpty": "OK",
                        "tooltipFontSize": 12,
                        "tooltipFontType": "Roboto",
                        "tooltipPrimarySortDirection": "desc",
                        "tooltipPrimarySortField": "thresholdLevel",
                        "tooltipSecondarySortDirection": "desc",
                        "tooltipSecondarySortField": "value",
                        "tooltipTimestampEnabled": true
                    },
                    "savedComposites": [],
                    "savedOverrides": [
                        {
                            "colors": [
                                "#299c46",
                                "#e5ac0e",
                                "#bf1b00",
                                "#ffffff"
                            ],
                            "decimals": 2,
                            "enabled": true,
                            "label": "OVERRIDE 1",
                            "metricName": "/.*/",
                            "operatorName": "current",
                            "prefix": "",
                            "sanitizeURLEnabled": false,
                            "scaledDecimals": null,
                            "suffix": "",
                            "thresholds": [
                                {
                                    "color": "#299c46",
                                    "state": 0,
                                    "value": 0.3
                                },
                                {
                                    "color": "#e5ac0e",
                                    "state": 1,
                                    "value": 0.5
                                }
                            ],
                            "unitFormat": "percentunit"
                        }
                    ],
                    "targets": [
                        {
                            "expr": "sum by (service_name) (mongodb_mongod_metrics_query_executor_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=~\"scanned_objects\"}) / ignoring(state)\nsum by (service_name) (mongodb_mongod_metrics_query_executor_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=~\"scanned_objects|scanned\"})",
                            "format": "time_series",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}}",
                            "refId": "A"
                        }
                    ],
                    "timeFrom": null,
                    "timeShift": null,
                    "title": "Index Filtering Effectiveness",
                    "type": "grafana-polystat-panel"
                }
            ],
            "title": "Query Efficiency Detail",
            "type": "row"
        },
        {
            "collapsed": false,
            "datasource": null,
            "gridPos": {
                "h": 1,
                "w": 24,
                "x": 0,
                "y": 15
            },
            "id": 105,
            "panels": [],
            "type": "row"
        },
        {
            "cacheTimeout": null,
            "colorBackground": false,
            "colorValue": true,
            "colors": [
                "#299c46",
                "#FF780A",
                "#d44a3a"
            ],
            "datasource": null,
            "decimals": 2,
            "description": "",
            "fieldConfig": {
                "defaults": {},
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
                "w": 6,
                "x": 0,
                "y": 16
            },
            "id": 78,
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
                "show": true
            },
            "tableColumn": "",
            "targets": [
                {
                    "expr": "sum(rate(mongodb_mongod_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type!=\"command\"}[$interval]) or \nirate(mongodb_mongod_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type!=\"command\"}[5m]) or \nrate(mongodb_mongos_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type!=\"command\"}[$interval]) or \nirate(mongodb_mongos_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type!=\"command\"}[5m]) or \nrate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type!=\"command\"}[$interval]) or \nirate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type!=\"command\"}[5m])) + \nsum(rate(mongodb_mongod_op_counters_repl_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type!~\"(command|query|getmore)\"}[$interval]) or \nirate(mongodb_mongod_op_counters_repl_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type!~\"(command|query|getmore)\"}[5m]) or\nrate(mongodb_mongos_op_counters_repl_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type!~\"(command|query|getmore)\"}[$interval]) or \nirate(mongodb_mongos_op_counters_repl_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type!~\"(command|query|getmore)\"}[5m]) or \nrate(mongodb_op_counters_repl_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type!~\"(command|query|getmore)\"}[$interval]) or \nirate(mongodb_op_counters_repl_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type!~\"(command|query|getmore)\"}[5m])) +\nsum(rate(mongodb_mongod_metrics_ttl_deleted_documents_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or \nirate(mongodb_mongod_metrics_ttl_deleted_documents_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]) or\nrate(mongodb_mongos_metrics_ttl_deleted_documents_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or \nirate(mongodb_mongos_metrics_ttl_deleted_documents_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]))",
                    "format": "time_series",
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "",
                    "refId": "A"
                }
            ],
            "thresholds": "1000,10000",
            "timeFrom": null,
            "timeShift": null,
            "title": "Top Opcounters",
            "type": "singlestat",
            "valueFontSize": "50%",
            "valueMaps": [
                {
                    "op": "=",
                    "text": "N/A",
                    "value": "null"
                }
            ],
            "valueName": "max"
        },
        {
            "cacheTimeout": null,
            "colorBackground": false,
            "colorValue": true,
            "colors": [
                "#299c46",
                "#FF780A",
                "#d44a3a"
            ],
            "datasource": null,
            "decimals": 2,
            "description": "Documents inserted, updated, deleted or returned per second.",
            "fieldConfig": {
                "defaults": {},
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
                "w": 6,
                "x": 6,
                "y": 16
            },
            "id": 81,
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
                "show": true
            },
            "tableColumn": "",
            "targets": [
                {
                    "expr": "sum(rate(mongodb_mongod_metrics_document_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or \nirate(mongodb_mongod_metrics_document_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]))",
                    "format": "time_series",
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "",
                    "refId": "A"
                }
            ],
            "thresholds": "1000,10000",
            "timeFrom": null,
            "timeShift": null,
            "title": "Top Document Operations",
            "type": "singlestat",
            "valueFontSize": "50%",
            "valueMaps": [
                {
                    "op": "=",
                    "text": "N/A",
                    "value": "null"
                }
            ],
            "valueName": "max"
        },
        {
            "cacheTimeout": null,
            "colorBackground": false,
            "colorValue": true,
            "colors": [
                "#299c46",
                "#FF780A",
                "#d44a3a"
            ],
            "datasource": null,
            "decimals": 2,
            "description": "Operations that are currently queued and waiting for read or write locks",
            "fieldConfig": {
                "defaults": {},
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
                "w": 6,
                "x": 12,
                "y": 16
            },
            "id": 82,
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
                "show": true
            },
            "tableColumn": "",
            "targets": [
                {
                    "expr": "sum(max_over_time(mongodb_mongod_global_lock_current_queue{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or \nmax_over_time(mongodb_mongod_global_lock_current_queue{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]))",
                    "format": "time_series",
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "",
                    "refId": "A"
                }
            ],
            "thresholds": "10,100",
            "timeFrom": null,
            "timeShift": null,
            "title": "Top Queued Operations",
            "type": "singlestat",
            "valueFontSize": "50%",
            "valueMaps": [
                {
                    "op": "=",
                    "text": "N/A",
                    "value": "null"
                }
            ],
            "valueName": "max"
        },
        {
            "cacheTimeout": null,
            "colorBackground": false,
            "colorValue": true,
            "colors": [
                "#299c46",
                "#FF9830",
                "#d44a3a"
            ],
            "datasource": null,
            "decimals": 2,
            "description": "Asserts. (Not a primary indicator of anything.)",
            "fieldConfig": {
                "defaults": {},
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
                "w": 6,
                "x": 18,
                "y": 16
            },
            "id": 107,
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
                "show": true
            },
            "tableColumn": "",
            "targets": [
                {
                    "expr": "sum(rate(mongodb_mongod_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or \nirate(mongodb_mongod_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]) or\nrate(mongodb_mongos_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or \nirate(mongodb_mongos_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]) or\nrate(mongodb_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or \nirate(mongodb_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]))",
                    "format": "time_series",
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "",
                    "refId": "A"
                }
            ],
            "thresholds": "50,100",
            "timeFrom": null,
            "timeShift": null,
            "title": "Total Assert Events",
            "type": "singlestat",
            "valueFontSize": "50%",
            "valueMaps": [
                {
                    "op": "=",
                    "text": "N/A",
                    "value": "null"
                }
            ],
            "valueName": "total"
        },
        {
            "collapsed": true,
            "datasource": null,
            "gridPos": {
                "h": 1,
                "w": 24,
                "x": 0,
                "y": 18
            },
            "id": 147,
            "panels": [
                {
                    "aliasColors": {},
                    "bars": false,
                    "dashLength": 10,
                    "dashes": false,
                    "datasource": null,
                    "decimals": 2,
                    "description": "The average rate of commands performed per second over the selected sample period",
                    "editable": true,
                    "error": false,
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
                        "y": 19
                    },
                    "height": "250px",
                    "hiddenSeries": false,
                    "id": 148,
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
                    "lines": false,
                    "linewidth": 2,
                    "links": [],
                    "nullPointMode": "null",
                    "options": {
                        "dataLinks": [
                            {
                                "targetBlank": true,
                                "title": "MongoDB Instance Summary - ${__series.name}",
                                "url": "/graph/d/mongodb-instance-summary/mongodb-instance-summary?var-service_name=${__series.name}&$__url_time_range"
                            }
                        ]
                    },
                    "paceLength": 10,
                    "percentage": false,
                    "pointradius": 1,
                    "points": true,
                    "renderer": "flot",
                    "seriesOverrides": [
                        {
                            "alias": "Avg",
                            "color": "#C4162A",
                            "fill": 0,
                            "legend": false,
                            "lines": true,
                            "linewidth": 2,
                            "points": false
                        }
                    ],
                    "spaceLength": 10,
                    "stack": false,
                    "steppedLine": false,
                    "targets": [
                        {
                            "expr": "topk(5,sum by (service_name) (rate(mongodb_mongod_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"command\"}[$interval]) or \nirate(mongodb_mongod_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"command\"}[5m]) or \nrate(mongodb_mongos_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"command\"}[$interval]) or \nirate(mongodb_mongos_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"command\"}[5m]) or \nrate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"command\"}[$interval]) or \nirate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"command\"}[5m])))",
                            "format": "time_series",
                            "hide": false,
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}}",
                            "refId": "J",
                            "step": 300
                        },
                        {
                            "expr": "avg(sum by (service_name) (rate(mongodb_mongod_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"command\"}[$interval]) or \nirate(mongodb_mongod_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"command\"}[5m]) or \nrate(mongodb_mongos_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"command\"}[$interval]) or \nirate(mongodb_mongos_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"command\"}[5m]) or \nrate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"command\"}[$interval]) or \nirate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"command\"}[5m])))",
                            "format": "time_series",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "Avg",
                            "refId": "A"
                        }
                    ],
                    "thresholds": [],
                    "timeFrom": null,
                    "timeRegions": [],
                    "timeShift": null,
                    "title": "Top 5 Command Operations",
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
                    "colors": [
                        "#299c46",
                        "rgba(237, 129, 40, 0.89)",
                        "#d44a3a"
                    ],
                    "datasource": null,
                    "gridPos": {
                        "h": 8,
                        "w": 12,
                        "x": 12,
                        "y": 19
                    },
                    "id": 158,
                    "links": [],
                    "options": {},
                    "polystat": {
                        "animationSpeed": 2500,
                        "columnAutoSize": true,
                        "columns": "",
                        "displayLimit": 100,
                        "fontAutoScale": true,
                        "fontSize": 12,
                        "fontType": "Roboto",
                        "globalDecimals": 2,
                        "globalDisplayMode": "all",
                        "globalDisplayTextTriggeredEmpty": "OK",
                        "globalOperatorName": "avg",
                        "globalUnitFormat": "short",
                        "gradientEnabled": true,
                        "hexagonSortByDirection": 1,
                        "hexagonSortByField": "name",
                        "maxMetrics": 0,
                        "polygonBorderColor": "black",
                        "polygonBorderSize": 2,
                        "polygonGlobalFillColor": "#0a50a1",
                        "radius": "",
                        "radiusAutoSize": true,
                        "rowAutoSize": true,
                        "rows": "",
                        "shape": "hexagon_pointed_top",
                        "tooltipDisplayMode": "all",
                        "tooltipDisplayTextTriggeredEmpty": "OK",
                        "tooltipFontSize": 12,
                        "tooltipFontType": "Roboto",
                        "tooltipPrimarySortDirection": 2,
                        "tooltipPrimarySortField": "thresholdLevel",
                        "tooltipSecondarySortDirection": 2,
                        "tooltipSecondarySortField": "value",
                        "tooltipTimestampEnabled": true
                    },
                    "savedComposites": [],
                    "savedOverrides": [
                        {
                            "colors": [
                                "#299c46",
                                "#e5ac0e",
                                "#bf1b00",
                                "#ffffff"
                            ],
                            "decimals": 2,
                            "enabled": true,
                            "label": "OVERRIDE 1",
                            "metricName": "/.*/",
                            "operatorName": "current",
                            "prefix": "",
                            "sanitizeURLEnabled": false,
                            "scaledDecimals": null,
                            "suffix": "",
                            "thresholds": [
                                {
                                    "color": "#299c46",
                                    "state": 0,
                                    "value": 20
                                },
                                {
                                    "color": "#e5ac0e",
                                    "state": 1,
                                    "value": 50
                                },
                                {
                                    "color": "#bf1b00",
                                    "state": 2,
                                    "value": 100
                                }
                            ],
                            "unitFormat": "ops"
                        }
                    ],
                    "targets": [
                        {
                            "expr": "sum by (service_name) (rate(mongodb_mongod_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"command\"}[$interval]) or \nirate(mongodb_mongod_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"command\"}[5m]) or \nrate(mongodb_mongos_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"command\"}[$interval]) or \nirate(mongodb_mongos_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"command\"}[5m]) or \nrate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"command\"}[$interval]) or \nirate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"command\"}[5m]))",
                            "format": "time_series",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}}",
                            "refId": "A"
                        }
                    ],
                    "timeFrom": null,
                    "timeShift": null,
                    "title": "Command Operations",
                    "type": "grafana-polystat-panel"
                },
                {
                    "aliasColors": {},
                    "bars": false,
                    "dashLength": 10,
                    "dashes": false,
                    "datasource": null,
                    "decimals": 2,
                    "description": "",
                    "editable": true,
                    "error": false,
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
                    "id": 152,
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
                    "lines": false,
                    "linewidth": 2,
                    "links": [],
                    "nullPointMode": "null",
                    "options": {
                        "dataLinks": [
                            {
                                "targetBlank": true,
                                "title": "MongoDB Instance Summary - ${__series.name}",
                                "url": "/graph/d/mongodb-instance-summary/mongodb-instance-summary?var-service_name=${__series.name}&$__url_time_range"
                            }
                        ]
                    },
                    "paceLength": 10,
                    "percentage": false,
                    "pointradius": 1,
                    "points": true,
                    "renderer": "flot",
                    "seriesOverrides": [
                        {
                            "alias": "Avg",
                            "color": "#C4162A",
                            "fill": 0,
                            "legend": false,
                            "lines": true,
                            "linewidth": 2,
                            "points": false
                        }
                    ],
                    "spaceLength": 10,
                    "stack": false,
                    "steppedLine": false,
                    "targets": [
                        {
                            "expr": "topk(5,(sum by (service_name) (rate(mongodb_mongod_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"getmore\"}[$interval]) or \nirate(mongodb_mongod_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"getmore\"}[5m]) or \nrate(mongodb_mongos_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"getmore\"}[$interval]) or \nirate(mongodb_mongos_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"getmore\"}[5m]) or \nrate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"getmore\"}[$interval]) or \nirate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"getmore\"}[5m]))))",
                            "format": "time_series",
                            "hide": false,
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}}",
                            "refId": "J",
                            "step": 300
                        },
                        {
                            "expr": "avg(sum by (service_name) (rate(mongodb_mongod_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"getmore\"}[$interval]) or \nirate(mongodb_mongod_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"getmore\"}[5m]) or \nrate(mongodb_mongos_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"getmore\"}[$interval]) or \nirate(mongodb_mongos_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"getmore\"}[5m]) or \nrate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"getmore\"}[$interval]) or \nirate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"getmore\"}[5m])))",
                            "format": "time_series",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "Avg",
                            "refId": "A"
                        }
                    ],
                    "thresholds": [],
                    "timeFrom": null,
                    "timeRegions": [],
                    "timeShift": null,
                    "title": "Top 5 Getmore Operations",
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
                    "colors": [
                        "#299c46",
                        "rgba(237, 129, 40, 0.89)",
                        "#d44a3a"
                    ],
                    "datasource": null,
                    "gridPos": {
                        "h": 8,
                        "w": 12,
                        "x": 12,
                        "y": 27
                    },
                    "id": 156,
                    "links": [],
                    "options": {},
                    "polystat": {
                        "animationSpeed": 2500,
                        "columnAutoSize": true,
                        "columns": "",
                        "displayLimit": 100,
                        "fontAutoScale": true,
                        "fontSize": 12,
                        "fontType": "Roboto",
                        "globalDecimals": 2,
                        "globalDisplayMode": "all",
                        "globalDisplayTextTriggeredEmpty": "OK",
                        "globalOperatorName": "avg",
                        "globalUnitFormat": "short",
                        "gradientEnabled": true,
                        "hexagonSortByDirection": 1,
                        "hexagonSortByField": "name",
                        "maxMetrics": 0,
                        "polygonBorderColor": "black",
                        "polygonBorderSize": 2,
                        "polygonGlobalFillColor": "#0a50a1",
                        "radius": "",
                        "radiusAutoSize": true,
                        "rowAutoSize": true,
                        "rows": "",
                        "shape": "hexagon_pointed_top",
                        "tooltipDisplayMode": "all",
                        "tooltipDisplayTextTriggeredEmpty": "OK",
                        "tooltipFontSize": 12,
                        "tooltipFontType": "Roboto",
                        "tooltipPrimarySortDirection": 2,
                        "tooltipPrimarySortField": "thresholdLevel",
                        "tooltipSecondarySortDirection": 2,
                        "tooltipSecondarySortField": "value",
                        "tooltipTimestampEnabled": true
                    },
                    "savedComposites": [],
                    "savedOverrides": [
                        {
                            "colors": [
                                "#299c46",
                                "#e5ac0e",
                                "#bf1b00",
                                "#ffffff"
                            ],
                            "decimals": 2,
                            "enabled": true,
                            "label": "OVERRIDE 1",
                            "metricName": "/.*/",
                            "operatorName": "current",
                            "prefix": "",
                            "sanitizeURLEnabled": false,
                            "scaledDecimals": null,
                            "suffix": "",
                            "thresholds": [
                                {
                                    "color": "#299c46",
                                    "state": 0,
                                    "value": 20
                                },
                                {
                                    "color": "#e5ac0e",
                                    "state": 1,
                                    "value": 50
                                },
                                {
                                    "color": "#bf1b00",
                                    "state": 2,
                                    "value": 100
                                }
                            ],
                            "unitFormat": "ops"
                        }
                    ],
                    "targets": [
                        {
                            "expr": "sum by (service_name) (rate(mongodb_mongod_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"getmore\"}[$interval]) or \nirate(mongodb_mongod_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"getmore\"}[5m]) or \nrate(mongodb_mongos_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"getmore\"}[$interval]) or \nirate(mongodb_mongos_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"getmore\"}[5m]) or \nrate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"getmore\"}[$interval]) or \nirate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"getmore\"}[5m]))",
                            "format": "time_series",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}}",
                            "refId": "A"
                        }
                    ],
                    "timeFrom": null,
                    "timeShift": null,
                    "title": "Getmore Operations",
                    "type": "grafana-polystat-panel"
                },
                {
                    "aliasColors": {},
                    "bars": false,
                    "dashLength": 10,
                    "dashes": false,
                    "datasource": null,
                    "decimals": 2,
                    "description": "",
                    "editable": true,
                    "error": false,
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
                    "id": 150,
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
                    "lines": false,
                    "linewidth": 2,
                    "links": [],
                    "nullPointMode": "null",
                    "options": {
                        "dataLinks": [
                            {
                                "targetBlank": true,
                                "title": "MongoDB Instance Summary - ${__series.name}",
                                "url": "/graph/d/mongodb-instance-summary/mongodb-instance-summary?var-service_name=${__series.name}&$__url_time_range"
                            }
                        ]
                    },
                    "paceLength": 10,
                    "percentage": false,
                    "pointradius": 1,
                    "points": true,
                    "renderer": "flot",
                    "seriesOverrides": [
                        {
                            "alias": "Avg",
                            "color": "#C4162A",
                            "fill": 0,
                            "legend": false,
                            "lines": true,
                            "linewidth": 2,
                            "points": false
                        }
                    ],
                    "spaceLength": 10,
                    "stack": false,
                    "steppedLine": false,
                    "targets": [
                        {
                            "expr": "topk(5,(sum by (service_name) (rate(mongodb_mongod_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"delete\"}[$interval]) or \nirate(mongodb_mongod_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"delete\"}[5m]) or \nrate(mongodb_mongos_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"delete\"}[$interval]) or \nirate(mongodb_mongos_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"delete\"}[5m]) or \nrate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"delete\"}[$interval]) or \nirate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"delete\"}[5m]))))",
                            "format": "time_series",
                            "hide": false,
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}}",
                            "refId": "J",
                            "step": 300
                        },
                        {
                            "expr": "avg(sum by (service_name) (rate(mongodb_mongod_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"delete\"}[$interval]) or \nirate(mongodb_mongod_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"delete\"}[5m]) or \nrate(mongodb_mongos_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"delete\"}[$interval]) or \nirate(mongodb_mongos_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"delete\"}[5m]) or \nrate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"delete\"}[$interval]) or \nirate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"delete\"}[5m])))",
                            "format": "time_series",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "Avg",
                            "refId": "A"
                        }
                    ],
                    "thresholds": [],
                    "timeFrom": null,
                    "timeRegions": [],
                    "timeShift": null,
                    "title": "Top 5 Delete Operations",
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
                    "colors": [
                        "#299c46",
                        "rgba(237, 129, 40, 0.89)",
                        "#d44a3a"
                    ],
                    "datasource": null,
                    "gridPos": {
                        "h": 8,
                        "w": 12,
                        "x": 12,
                        "y": 35
                    },
                    "id": 157,
                    "links": [],
                    "options": {},
                    "polystat": {
                        "animationSpeed": 2500,
                        "columnAutoSize": true,
                        "columns": "",
                        "displayLimit": 100,
                        "fontAutoScale": true,
                        "fontSize": 12,
                        "fontType": "Roboto",
                        "globalDecimals": 2,
                        "globalDisplayMode": "all",
                        "globalDisplayTextTriggeredEmpty": "OK",
                        "globalOperatorName": "avg",
                        "globalUnitFormat": "short",
                        "gradientEnabled": true,
                        "hexagonSortByDirection": "asc",
                        "hexagonSortByField": "name",
                        "maxMetrics": 0,
                        "polygonBorderColor": "black",
                        "polygonBorderSize": 2,
                        "polygonGlobalFillColor": "#0a50a1",
                        "radius": "",
                        "radiusAutoSize": true,
                        "rowAutoSize": true,
                        "rows": "",
                        "shape": "hexagon_pointed_top",
                        "tooltipDisplayMode": "all",
                        "tooltipDisplayTextTriggeredEmpty": "OK",
                        "tooltipFontSize": 12,
                        "tooltipFontType": "Roboto",
                        "tooltipPrimarySortDirection": "desc",
                        "tooltipPrimarySortField": "thresholdLevel",
                        "tooltipSecondarySortDirection": "desc",
                        "tooltipSecondarySortField": "value",
                        "tooltipTimestampEnabled": true
                    },
                    "savedComposites": [],
                    "savedOverrides": [
                        {
                            "colors": [
                                "#299c46",
                                "#e5ac0e",
                                "#bf1b00",
                                "#ffffff"
                            ],
                            "decimals": 2,
                            "enabled": true,
                            "label": "OVERRIDE 1",
                            "metricName": "/.*/",
                            "operatorName": "current",
                            "prefix": "",
                            "sanitizeURLEnabled": false,
                            "scaledDecimals": null,
                            "suffix": "",
                            "thresholds": [
                                {
                                    "color": "#299c46",
                                    "state": 0,
                                    "value": 20
                                },
                                {
                                    "color": "#e5ac0e",
                                    "state": 1,
                                    "value": 50
                                },
                                {
                                    "color": "#bf1b00",
                                    "state": 2,
                                    "value": 100
                                }
                            ],
                            "unitFormat": "ops"
                        }
                    ],
                    "targets": [
                        {
                            "expr": "sum by (service_name) (rate(mongodb_mongod_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"delete\"}[$interval]) or \nirate(mongodb_mongod_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"delete\"}[5m]) or \nrate(mongodb_mongos_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"delete\"}[$interval]) or \nirate(mongodb_mongos_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"delete\"}[5m]) or \nrate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"delete\"}[$interval]) or \nirate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"delete\"}[5m]))",
                            "format": "time_series",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}}",
                            "refId": "A"
                        }
                    ],
                    "timeFrom": null,
                    "timeShift": null,
                    "title": "Delete Operations",
                    "type": "grafana-polystat-panel"
                },
                {
                    "aliasColors": {},
                    "bars": false,
                    "dashLength": 10,
                    "dashes": false,
                    "datasource": null,
                    "decimals": 2,
                    "description": "",
                    "editable": true,
                    "error": false,
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
                    "id": 151,
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
                    "lines": false,
                    "linewidth": 2,
                    "links": [],
                    "nullPointMode": "null",
                    "options": {
                        "dataLinks": [
                            {
                                "targetBlank": true,
                                "title": "MongoDB Instance Summary - ${__series.name}",
                                "url": "/graph/d/mongodb-instance-summary/mongodb-instance-summary?var-service_name=${__series.name}&$__url_time_range"
                            }
                        ]
                    },
                    "paceLength": 10,
                    "percentage": false,
                    "pointradius": 1,
                    "points": true,
                    "renderer": "flot",
                    "seriesOverrides": [
                        {
                            "alias": "Avg",
                            "color": "#C4162A",
                            "fill": 0,
                            "legend": false,
                            "lines": true,
                            "linewidth": 2,
                            "points": false
                        }
                    ],
                    "spaceLength": 10,
                    "stack": false,
                    "steppedLine": false,
                    "targets": [
                        {
                            "expr": "topk(5,(sum by (service_name) (rate(mongodb_mongod_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"insert\"}[$interval]) or \nirate(mongodb_mongod_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"insert\"}[5m]) or \nrate(mongodb_mongos_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"insert\"}[$interval]) or \nirate(mongodb_mongos_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"insert\"}[5m]) or \nrate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"insert\"}[$interval]) or \nirate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"insert\"}[5m]))))",
                            "format": "time_series",
                            "hide": false,
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}}",
                            "refId": "J",
                            "step": 300
                        },
                        {
                            "expr": "avg(sum by (service_name) (rate(mongodb_mongod_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"insert\"}[$interval]) or \nirate(mongodb_mongod_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"insert\"}[5m]) or \nrate(mongodb_mongos_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"insert\"}[$interval]) or \nirate(mongodb_mongos_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"insert\"}[5m]) or \nrate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"insert\"}[$interval]) or \nirate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"insert\"}[5m])))",
                            "format": "time_series",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "Avg",
                            "refId": "A"
                        }
                    ],
                    "thresholds": [],
                    "timeFrom": null,
                    "timeRegions": [],
                    "timeShift": null,
                    "title": "Top 5 Insert Operations",
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
                    "colors": [
                        "#299c46",
                        "rgba(237, 129, 40, 0.89)",
                        "#d44a3a"
                    ],
                    "datasource": null,
                    "gridPos": {
                        "h": 8,
                        "w": 12,
                        "x": 12,
                        "y": 43
                    },
                    "id": 155,
                    "links": [],
                    "options": {},
                    "polystat": {
                        "animationSpeed": 2500,
                        "columnAutoSize": true,
                        "columns": "",
                        "displayLimit": 100,
                        "fontAutoScale": true,
                        "fontSize": 12,
                        "fontType": "Roboto",
                        "globalDecimals": 2,
                        "globalDisplayMode": "all",
                        "globalDisplayTextTriggeredEmpty": "OK",
                        "globalOperatorName": "avg",
                        "globalUnitFormat": "short",
                        "gradientEnabled": true,
                        "hexagonSortByDirection": "asc",
                        "hexagonSortByField": "name",
                        "maxMetrics": 0,
                        "polygonBorderColor": "black",
                        "polygonBorderSize": 2,
                        "polygonGlobalFillColor": "#0a50a1",
                        "radius": "",
                        "radiusAutoSize": true,
                        "rowAutoSize": true,
                        "rows": "",
                        "shape": "hexagon_pointed_top",
                        "tooltipDisplayMode": "all",
                        "tooltipDisplayTextTriggeredEmpty": "OK",
                        "tooltipFontSize": 12,
                        "tooltipFontType": "Roboto",
                        "tooltipPrimarySortDirection": "desc",
                        "tooltipPrimarySortField": "thresholdLevel",
                        "tooltipSecondarySortDirection": "desc",
                        "tooltipSecondarySortField": "value",
                        "tooltipTimestampEnabled": true
                    },
                    "savedComposites": [],
                    "savedOverrides": [
                        {
                            "colors": [
                                "#299c46",
                                "#e5ac0e",
                                "#bf1b00",
                                "#ffffff"
                            ],
                            "decimals": 2,
                            "enabled": true,
                            "label": "OVERRIDE 1",
                            "metricName": "/.*/",
                            "operatorName": "current",
                            "prefix": "",
                            "sanitizeURLEnabled": false,
                            "scaledDecimals": null,
                            "suffix": "",
                            "thresholds": [
                                {
                                    "color": "#299c46",
                                    "state": 0,
                                    "value": 20
                                },
                                {
                                    "color": "#e5ac0e",
                                    "state": 1,
                                    "value": 50
                                },
                                {
                                    "color": "#bf1b00",
                                    "state": 2,
                                    "value": 100
                                }
                            ],
                            "unitFormat": "ops"
                        }
                    ],
                    "targets": [
                        {
                            "expr": "sum by (service_name) (rate(mongodb_mongod_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"insert\"}[$interval]) or \nirate(mongodb_mongod_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"insert\"}[5m]) or \nrate(mongodb_mongos_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"insert\"}[$interval]) or \nirate(mongodb_mongos_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"insert\"}[5m]) or \nrate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"insert\"}[$interval]) or \nirate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"insert\"}[5m]))",
                            "format": "time_series",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}}",
                            "refId": "A"
                        }
                    ],
                    "timeFrom": null,
                    "timeShift": null,
                    "title": "Insert Operations",
                    "type": "grafana-polystat-panel"
                },
                {
                    "aliasColors": {},
                    "bars": false,
                    "dashLength": 10,
                    "dashes": false,
                    "datasource": null,
                    "decimals": 2,
                    "description": "",
                    "editable": true,
                    "error": false,
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
                    "height": "250px",
                    "hiddenSeries": false,
                    "id": 154,
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
                    "lines": false,
                    "linewidth": 2,
                    "links": [],
                    "nullPointMode": "null",
                    "options": {
                        "dataLinks": [
                            {
                                "targetBlank": true,
                                "title": "MongoDB Instance Summary - ${__series.name}",
                                "url": "/graph/d/mongodb-instance-summary/mongodb-instance-summary?var-service_name=${__series.name}&$__url_time_range"
                            }
                        ]
                    },
                    "paceLength": 10,
                    "percentage": false,
                    "pointradius": 1,
                    "points": true,
                    "renderer": "flot",
                    "seriesOverrides": [
                        {
                            "alias": "Avg",
                            "color": "#C4162A",
                            "fill": 0,
                            "legend": false,
                            "lines": true,
                            "linewidth": 2,
                            "points": false
                        }
                    ],
                    "spaceLength": 10,
                    "stack": false,
                    "steppedLine": false,
                    "targets": [
                        {
                            "expr": "topk(5,(sum by (service_name) (rate(mongodb_mongod_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"update\"}[$interval]) or \nirate(mongodb_mongod_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"update\"}[5m]) or \nrate(mongodb_mongos_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"update\"}[$interval]) or \nirate(mongodb_mongos_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"update\"}[5m]) or \nrate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"update\"}[$interval]) or \nirate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"update\"}[5m]))))",
                            "format": "time_series",
                            "hide": false,
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}}",
                            "refId": "J",
                            "step": 300
                        },
                        {
                            "expr": "avg(sum by (service_name) (rate(mongodb_mongod_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"update\"}[$interval]) or \nirate(mongodb_mongod_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"update\"}[5m]) or \nrate(mongodb_mongos_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"update\"}[$interval]) or \nirate(mongodb_mongos_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"update\"}[5m]) or \nrate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"update\"}[$interval]) or \nirate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"update\"}[5m])))",
                            "format": "time_series",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "Avg",
                            "refId": "A"
                        }
                    ],
                    "thresholds": [],
                    "timeFrom": null,
                    "timeRegions": [],
                    "timeShift": null,
                    "title": "Top 5 Update Operations",
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
                    "colors": [
                        "#299c46",
                        "rgba(237, 129, 40, 0.89)",
                        "#d44a3a"
                    ],
                    "datasource": null,
                    "gridPos": {
                        "h": 8,
                        "w": 12,
                        "x": 12,
                        "y": 51
                    },
                    "id": 173,
                    "links": [],
                    "options": {},
                    "polystat": {
                        "animationSpeed": 2500,
                        "columnAutoSize": true,
                        "columns": "",
                        "displayLimit": 100,
                        "fontAutoScale": true,
                        "fontSize": 12,
                        "fontType": "Roboto",
                        "globalDecimals": 2,
                        "globalDisplayMode": "all",
                        "globalDisplayTextTriggeredEmpty": "OK",
                        "globalOperatorName": "avg",
                        "globalUnitFormat": "short",
                        "gradientEnabled": true,
                        "hexagonSortByDirection": "asc",
                        "hexagonSortByField": "name",
                        "maxMetrics": 0,
                        "polygonBorderColor": "black",
                        "polygonBorderSize": 2,
                        "polygonGlobalFillColor": "#0a50a1",
                        "radius": "",
                        "radiusAutoSize": true,
                        "rowAutoSize": true,
                        "rows": "",
                        "shape": "hexagon_pointed_top",
                        "tooltipDisplayMode": "all",
                        "tooltipDisplayTextTriggeredEmpty": "OK",
                        "tooltipFontSize": 12,
                        "tooltipFontType": "Roboto",
                        "tooltipPrimarySortDirection": "desc",
                        "tooltipPrimarySortField": "thresholdLevel",
                        "tooltipSecondarySortDirection": "desc",
                        "tooltipSecondarySortField": "value",
                        "tooltipTimestampEnabled": true
                    },
                    "savedComposites": [],
                    "savedOverrides": [
                        {
                            "colors": [
                                "#299c46",
                                "#e5ac0e",
                                "#bf1b00",
                                "#ffffff"
                            ],
                            "decimals": 2,
                            "enabled": true,
                            "label": "OVERRIDE 1",
                            "metricName": "/.*/",
                            "operatorName": "current",
                            "prefix": "",
                            "sanitizeURLEnabled": false,
                            "scaledDecimals": null,
                            "suffix": "",
                            "thresholds": [
                                {
                                    "color": "#299c46",
                                    "state": 0,
                                    "value": 20
                                },
                                {
                                    "color": "#e5ac0e",
                                    "state": 1,
                                    "value": 50
                                },
                                {
                                    "color": "#bf1b00",
                                    "state": 2,
                                    "value": 100
                                }
                            ],
                            "unitFormat": "ops"
                        }
                    ],
                    "targets": [
                        {
                            "expr": "sum by (service_name) (rate(mongodb_mongod_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"update\"}[$interval]) or \nirate(mongodb_mongod_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"update\"}[5m]) or \nrate(mongodb_mongos_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"update\"}[$interval]) or \nirate(mongodb_mongos_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"update\"}[5m]) or \nrate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"update\"}[$interval]) or \nirate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"update\"}[5m]))",
                            "format": "time_series",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}}",
                            "refId": "A"
                        }
                    ],
                    "timeFrom": null,
                    "timeShift": null,
                    "title": "Update Operations",
                    "type": "grafana-polystat-panel"
                },
                {
                    "aliasColors": {},
                    "bars": false,
                    "dashLength": 10,
                    "dashes": false,
                    "datasource": null,
                    "decimals": 2,
                    "description": "",
                    "editable": true,
                    "error": false,
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
                    "id": 153,
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
                    "lines": false,
                    "linewidth": 2,
                    "links": [],
                    "nullPointMode": "null",
                    "options": {
                        "dataLinks": [
                            {
                                "targetBlank": true,
                                "title": "MongoDB Instance Summary - ${__series.name}",
                                "url": "/graph/d/mongodb-instance-summary/mongodb-instance-summary?var-service_name=${__series.name}&$__url_time_range"
                            }
                        ]
                    },
                    "paceLength": 10,
                    "percentage": false,
                    "pointradius": 1,
                    "points": true,
                    "renderer": "flot",
                    "seriesOverrides": [
                        {
                            "alias": "Avg",
                            "color": "#C4162A",
                            "fill": 0,
                            "legend": false,
                            "lines": true,
                            "linewidth": 2,
                            "points": false
                        }
                    ],
                    "spaceLength": 10,
                    "stack": false,
                    "steppedLine": false,
                    "targets": [
                        {
                            "expr": "topk(5,(sum by (service_name) (rate(mongodb_mongod_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"query\"}[$interval]) or \nirate(mongodb_mongod_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"query\"}[5m]) or \nrate(mongodb_mongos_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"query\"}[$interval]) or \nirate(mongodb_mongos_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"query\"}[5m]) or \nrate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"query\"}[$interval]) or \nirate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"query\"}[5m]))))",
                            "format": "time_series",
                            "hide": false,
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}}",
                            "refId": "J",
                            "step": 300
                        },
                        {
                            "expr": "avg(sum by (service_name) (rate(mongodb_mongod_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"query\"}[$interval]) or \nirate(mongodb_mongod_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"query\"}[5m]) or \nrate(mongodb_mongos_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"query\"}[$interval]) or \nirate(mongodb_mongos_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"query\"}[5m]) or \nrate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"query\"}[$interval]) or \nirate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"query\"}[5m])))",
                            "format": "time_series",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "Avg",
                            "refId": "A"
                        }
                    ],
                    "thresholds": [],
                    "timeFrom": null,
                    "timeRegions": [],
                    "timeShift": null,
                    "title": "Top 5 Query Operations",
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
                    "colors": [
                        "#299c46",
                        "rgba(237, 129, 40, 0.89)",
                        "#d44a3a"
                    ],
                    "datasource": null,
                    "gridPos": {
                        "h": 8,
                        "w": 12,
                        "x": 12,
                        "y": 59
                    },
                    "id": 149,
                    "links": [],
                    "options": {},
                    "polystat": {
                        "animationSpeed": 2500,
                        "columnAutoSize": true,
                        "columns": "",
                        "displayLimit": 100,
                        "fontAutoScale": true,
                        "fontSize": 12,
                        "fontType": "Roboto",
                        "globalDecimals": 2,
                        "globalDisplayMode": "all",
                        "globalDisplayTextTriggeredEmpty": "OK",
                        "globalOperatorName": "avg",
                        "globalUnitFormat": "short",
                        "gradientEnabled": true,
                        "hexagonSortByDirection": "asc",
                        "hexagonSortByField": "name",
                        "maxMetrics": 0,
                        "polygonBorderColor": "black",
                        "polygonBorderSize": 2,
                        "polygonGlobalFillColor": "#0a50a1",
                        "radius": "",
                        "radiusAutoSize": true,
                        "rowAutoSize": true,
                        "rows": "",
                        "shape": "hexagon_pointed_top",
                        "tooltipDisplayMode": "all",
                        "tooltipDisplayTextTriggeredEmpty": "OK",
                        "tooltipFontSize": 12,
                        "tooltipFontType": "Roboto",
                        "tooltipPrimarySortDirection": "desc",
                        "tooltipPrimarySortField": "thresholdLevel",
                        "tooltipSecondarySortDirection": "desc",
                        "tooltipSecondarySortField": "value",
                        "tooltipTimestampEnabled": true
                    },
                    "savedComposites": [],
                    "savedOverrides": [
                        {
                            "colors": [
                                "#299c46",
                                "#e5ac0e",
                                "#bf1b00",
                                "#ffffff"
                            ],
                            "decimals": 2,
                            "enabled": true,
                            "label": "OVERRIDE 1",
                            "metricName": "/.*/",
                            "operatorName": "current",
                            "prefix": "",
                            "sanitizeURLEnabled": false,
                            "scaledDecimals": null,
                            "suffix": "",
                            "thresholds": [
                                {
                                    "color": "#299c46",
                                    "state": 0,
                                    "value": 20
                                },
                                {
                                    "color": "#e5ac0e",
                                    "state": 1,
                                    "value": 50
                                },
                                {
                                    "color": "#bf1b00",
                                    "state": 2,
                                    "value": 100
                                }
                            ],
                            "unitFormat": "ops"
                        }
                    ],
                    "targets": [
                        {
                            "expr": "sum by (service_name) (rate(mongodb_mongod_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"query\"}[$interval]) or \nirate(mongodb_mongod_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"query\"}[5m]) or \nrate(mongodb_mongos_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"query\"}[$interval]) or \nirate(mongodb_mongos_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"query\"}[5m]) or \nrate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"query\"}[$interval]) or \nirate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"query\"}[5m]))",
                            "format": "time_series",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}}",
                            "refId": "A"
                        }
                    ],
                    "timeFrom": null,
                    "timeShift": null,
                    "title": "Query Operations",
                    "type": "grafana-polystat-panel"
                }
            ],
            "title": "Opcounters Detail",
            "type": "row"
        },
        {
            "collapsed": true,
            "datasource": null,
            "gridPos": {
                "h": 1,
                "w": 24,
                "x": 0,
                "y": 19
            },
            "id": 132,
            "panels": [
                {
                    "aliasColors": {},
                    "bars": false,
                    "dashLength": 10,
                    "dashes": false,
                    "datasource": null,
                    "decimals": 2,
                    "description": "This panel shows the actual number of documents that were affected on average during the given time period. A single *update* command for example could perform updates on hundreds or thousands of documents. If these counters seem high, you might want to check your statements to make sure they are properly operating against the documents that they should.",
                    "editable": true,
                    "error": false,
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
                        "y": 20
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
                    "lines": false,
                    "linewidth": 2,
                    "links": [],
                    "nullPointMode": "null",
                    "options": {
                        "dataLinks": [
                            {
                                "targetBlank": true,
                                "title": "MongoDB Instance Summary - ${__series.name}",
                                "url": "/graph/d/mongodb-instance-summary/mongodb-instance-summary?var-service_name=${__series.name}&$__url_time_range"
                            }
                        ]
                    },
                    "paceLength": 10,
                    "percentage": false,
                    "pointradius": 1,
                    "points": true,
                    "renderer": "flot",
                    "seriesOverrides": [
                        {
                            "alias": "Avg",
                            "color": "#C4162A",
                            "fill": 0,
                            "legend": false,
                            "lines": true,
                            "points": false
                        }
                    ],
                    "spaceLength": 10,
                    "stack": false,
                    "steppedLine": false,
                    "targets": [
                        {
                            "expr": "topk(5,avg by (service_name) ((rate(mongodb_mongod_metrics_document_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",state=\"deleted\"}[$interval]) or \nirate(mongodb_mongod_metrics_document_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",state=\"deleted\"}[5m]))))",
                            "format": "time_series",
                            "hide": false,
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}}",
                            "refId": "J",
                            "step": 300
                        },
                        {
                            "expr": "avg(rate(mongodb_mongod_metrics_document_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",state=\"deleted\"}[$interval]) or \nirate(mongodb_mongod_metrics_document_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",state=\"deleted\"}[5m]))",
                            "format": "time_series",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "Avg",
                            "refId": "A"
                        }
                    ],
                    "thresholds": [],
                    "timeFrom": null,
                    "timeRegions": [],
                    "timeShift": null,
                    "title": "Top 5 Document Delete Operations",
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
                    "colors": [
                        "#299c46",
                        "rgba(237, 129, 40, 0.89)",
                        "#d44a3a"
                    ],
                    "datasource": null,
                    "gridPos": {
                        "h": 8,
                        "w": 12,
                        "x": 12,
                        "y": 20
                    },
                    "id": 174,
                    "links": [],
                    "options": {},
                    "polystat": {
                        "animationSpeed": 2500,
                        "columnAutoSize": true,
                        "columns": "",
                        "displayLimit": 100,
                        "fontAutoScale": true,
                        "fontSize": 12,
                        "fontType": "Roboto",
                        "globalDecimals": 2,
                        "globalDisplayMode": "all",
                        "globalDisplayTextTriggeredEmpty": "OK",
                        "globalOperatorName": "avg",
                        "globalUnitFormat": "short",
                        "gradientEnabled": true,
                        "hexagonSortByDirection": "asc",
                        "hexagonSortByField": "name",
                        "maxMetrics": 0,
                        "polygonBorderColor": "black",
                        "polygonBorderSize": 2,
                        "polygonGlobalFillColor": "#0a50a1",
                        "radius": "",
                        "radiusAutoSize": true,
                        "rowAutoSize": true,
                        "rows": "",
                        "shape": "hexagon_pointed_top",
                        "tooltipDisplayMode": "all",
                        "tooltipDisplayTextTriggeredEmpty": "OK",
                        "tooltipFontSize": 12,
                        "tooltipFontType": "Roboto",
                        "tooltipPrimarySortDirection": "desc",
                        "tooltipPrimarySortField": "thresholdLevel",
                        "tooltipSecondarySortDirection": "desc",
                        "tooltipSecondarySortField": "value",
                        "tooltipTimestampEnabled": true
                    },
                    "savedComposites": [],
                    "savedOverrides": [
                        {
                            "colors": [
                                "#299c46",
                                "#e5ac0e",
                                "#bf1b00",
                                "#ffffff"
                            ],
                            "decimals": 2,
                            "enabled": true,
                            "label": "OVERRIDE 1",
                            "metricName": "/.*/",
                            "operatorName": "current",
                            "prefix": "",
                            "sanitizeURLEnabled": false,
                            "scaledDecimals": null,
                            "suffix": "",
                            "thresholds": [
                                {
                                    "color": "#299c46",
                                    "state": 0,
                                    "value": 20
                                },
                                {
                                    "color": "#e5ac0e",
                                    "state": 1,
                                    "value": 50
                                },
                                {
                                    "color": "#bf1b00",
                                    "state": 2,
                                    "value": 100
                                }
                            ],
                            "unitFormat": "ops"
                        }
                    ],
                    "targets": [
                        {
                            "expr": "avg by (service_name) (rate(mongodb_mongod_metrics_document_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",state=\"deleted\"}[$interval]) or \nirate(mongodb_mongod_metrics_document_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",state=\"deleted\"}[5m]))",
                            "format": "time_series",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}}",
                            "refId": "A"
                        }
                    ],
                    "timeFrom": null,
                    "timeShift": null,
                    "title": "Document Delete Operations",
                    "type": "grafana-polystat-panel"
                },
                {
                    "aliasColors": {},
                    "bars": false,
                    "dashLength": 10,
                    "dashes": false,
                    "datasource": null,
                    "decimals": 2,
                    "description": "This panel shows the actual number of documents that were affected on average during the given time period. A single *update* command for example could perform updates on hundreds or thousands of documents. If these counters seem high, you might want to check your statements to make sure they are properly operating against the documents that they should.",
                    "editable": true,
                    "error": false,
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
                        "y": 28
                    },
                    "height": "250px",
                    "hiddenSeries": false,
                    "id": 175,
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
                    "lines": false,
                    "linewidth": 2,
                    "links": [],
                    "nullPointMode": "null",
                    "options": {
                        "dataLinks": [
                            {
                                "targetBlank": true,
                                "title": "MongoDB Instance Summary - ${__series.name}",
                                "url": "/graph/d/mongodb-instance-summary/mongodb-instance-summary?var-service_name=${__series.name}&$__url_time_range"
                            }
                        ]
                    },
                    "paceLength": 10,
                    "percentage": false,
                    "pointradius": 1,
                    "points": true,
                    "renderer": "flot",
                    "seriesOverrides": [
                        {
                            "alias": "Avg",
                            "color": "#C4162A",
                            "fill": 0,
                            "legend": false,
                            "lines": true,
                            "points": false
                        }
                    ],
                    "spaceLength": 10,
                    "stack": false,
                    "steppedLine": false,
                    "targets": [
                        {
                            "expr": "topk(5,avg by (service_name) ((rate(mongodb_mongod_metrics_document_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",state=\"inserted\"}[$interval]) or \nirate(mongodb_mongod_metrics_document_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",state=\"inserted\"}[5m]))))",
                            "format": "time_series",
                            "hide": false,
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}}",
                            "refId": "J",
                            "step": 300
                        },
                        {
                            "expr": "avg(rate(mongodb_mongod_metrics_document_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",state=\"inserted\"}[$interval]) or \nirate(mongodb_mongod_metrics_document_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",state=\"inserted\"}[5m]))",
                            "format": "time_series",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "Avg",
                            "refId": "A"
                        }
                    ],
                    "thresholds": [],
                    "timeFrom": null,
                    "timeRegions": [],
                    "timeShift": null,
                    "title": "Top 5 Document Insert Operations",
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
                    "colors": [
                        "#299c46",
                        "rgba(237, 129, 40, 0.89)",
                        "#d44a3a"
                    ],
                    "datasource": null,
                    "gridPos": {
                        "h": 8,
                        "w": 12,
                        "x": 12,
                        "y": 28
                    },
                    "id": 180,
                    "links": [],
                    "options": {},
                    "polystat": {
                        "animationSpeed": 2500,
                        "columnAutoSize": true,
                        "columns": "",
                        "displayLimit": 100,
                        "fontAutoScale": true,
                        "fontSize": 12,
                        "fontType": "Roboto",
                        "globalDecimals": 2,
                        "globalDisplayMode": "all",
                        "globalDisplayTextTriggeredEmpty": "OK",
                        "globalOperatorName": "avg",
                        "globalUnitFormat": "short",
                        "gradientEnabled": true,
                        "hexagonSortByDirection": "asc",
                        "hexagonSortByField": "name",
                        "maxMetrics": 0,
                        "polygonBorderColor": "black",
                        "polygonBorderSize": 2,
                        "polygonGlobalFillColor": "#0a50a1",
                        "radius": "",
                        "radiusAutoSize": true,
                        "rowAutoSize": true,
                        "rows": "",
                        "shape": "hexagon_pointed_top",
                        "tooltipDisplayMode": "all",
                        "tooltipDisplayTextTriggeredEmpty": "OK",
                        "tooltipFontSize": 12,
                        "tooltipFontType": "Roboto",
                        "tooltipPrimarySortDirection": "desc",
                        "tooltipPrimarySortField": "thresholdLevel",
                        "tooltipSecondarySortDirection": "desc",
                        "tooltipSecondarySortField": "value",
                        "tooltipTimestampEnabled": true
                    },
                    "savedComposites": [],
                    "savedOverrides": [
                        {
                            "colors": [
                                "#299c46",
                                "#e5ac0e",
                                "#bf1b00",
                                "#ffffff"
                            ],
                            "decimals": 2,
                            "enabled": true,
                            "label": "OVERRIDE 1",
                            "metricName": "/.*/",
                            "operatorName": "current",
                            "prefix": "",
                            "sanitizeURLEnabled": false,
                            "scaledDecimals": null,
                            "suffix": "",
                            "thresholds": [
                                {
                                    "color": "#299c46",
                                    "state": 0,
                                    "value": 20
                                },
                                {
                                    "color": "#e5ac0e",
                                    "state": 1,
                                    "value": 50
                                },
                                {
                                    "color": "#bf1b00",
                                    "state": 2,
                                    "value": 100
                                }
                            ],
                            "unitFormat": "ops"
                        }
                    ],
                    "targets": [
                        {
                            "expr": "avg by (service_name) (rate(mongodb_mongod_metrics_document_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",state=\"inserted\"}[$interval]) or \nirate(mongodb_mongod_metrics_document_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",state=\"inserted\"}[5m]))",
                            "format": "time_series",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}}",
                            "refId": "A"
                        }
                    ],
                    "timeFrom": null,
                    "timeShift": null,
                    "title": "Document Insert Operations",
                    "type": "grafana-polystat-panel"
                },
                {
                    "aliasColors": {},
                    "bars": false,
                    "dashLength": 10,
                    "dashes": false,
                    "datasource": null,
                    "decimals": 2,
                    "description": "This panel shows the actual number of documents that were affected on average during the given time period. A single *update* command for example could perform updates on hundreds or thousands of documents. If these counters seem high, you might want to check your statements to make sure they are properly operating against the documents that they should.",
                    "editable": true,
                    "error": false,
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
                        "y": 36
                    },
                    "height": "250px",
                    "hiddenSeries": false,
                    "id": 176,
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
                    "lines": false,
                    "linewidth": 2,
                    "links": [],
                    "nullPointMode": "null",
                    "options": {
                        "dataLinks": [
                            {
                                "targetBlank": true,
                                "title": "MongoDB Instance Summary - ${__series.name}",
                                "url": "/graph/d/mongodb-instance-summary/mongodb-instance-summary?var-service_name=${__series.name}&$__url_time_range"
                            }
                        ]
                    },
                    "paceLength": 10,
                    "percentage": false,
                    "pointradius": 1,
                    "points": true,
                    "renderer": "flot",
                    "seriesOverrides": [
                        {
                            "alias": "Avg",
                            "color": "#C4162A",
                            "fill": 0,
                            "legend": false,
                            "lines": true,
                            "points": false
                        }
                    ],
                    "spaceLength": 10,
                    "stack": false,
                    "steppedLine": false,
                    "targets": [
                        {
                            "expr": "topk(5,avg by (service_name) ((rate(mongodb_mongod_metrics_document_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",state=\"returned\"}[$interval]) or \nirate(mongodb_mongod_metrics_document_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",state=\"returned\"}[5m]))))",
                            "format": "time_series",
                            "hide": false,
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}}",
                            "refId": "J",
                            "step": 300
                        },
                        {
                            "expr": "avg(rate(mongodb_mongod_metrics_document_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",state=\"returned\"}[$interval]) or \nirate(mongodb_mongod_metrics_document_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",state=\"returned\"}[5m]))",
                            "format": "time_series",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "Avg",
                            "refId": "A"
                        }
                    ],
                    "thresholds": [],
                    "timeFrom": null,
                    "timeRegions": [],
                    "timeShift": null,
                    "title": "Top 5 Document Return Operations",
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
                    "colors": [
                        "#299c46",
                        "rgba(237, 129, 40, 0.89)",
                        "#d44a3a"
                    ],
                    "datasource": null,
                    "gridPos": {
                        "h": 8,
                        "w": 12,
                        "x": 12,
                        "y": 36
                    },
                    "id": 179,
                    "links": [],
                    "options": {},
                    "polystat": {
                        "animationSpeed": 2500,
                        "columnAutoSize": true,
                        "columns": "",
                        "displayLimit": 100,
                        "fontAutoScale": true,
                        "fontSize": 12,
                        "fontType": "Roboto",
                        "globalDecimals": 2,
                        "globalDisplayMode": "all",
                        "globalDisplayTextTriggeredEmpty": "OK",
                        "globalOperatorName": "avg",
                        "globalUnitFormat": "short",
                        "gradientEnabled": true,
                        "hexagonSortByDirection": "asc",
                        "hexagonSortByField": "name",
                        "maxMetrics": 0,
                        "polygonBorderColor": "black",
                        "polygonBorderSize": 2,
                        "polygonGlobalFillColor": "#0a50a1",
                        "radius": "",
                        "radiusAutoSize": true,
                        "rowAutoSize": true,
                        "rows": "",
                        "shape": "hexagon_pointed_top",
                        "tooltipDisplayMode": "all",
                        "tooltipDisplayTextTriggeredEmpty": "OK",
                        "tooltipFontSize": 12,
                        "tooltipFontType": "Roboto",
                        "tooltipPrimarySortDirection": "desc",
                        "tooltipPrimarySortField": "thresholdLevel",
                        "tooltipSecondarySortDirection": "desc",
                        "tooltipSecondarySortField": "value",
                        "tooltipTimestampEnabled": true
                    },
                    "savedComposites": [],
                    "savedOverrides": [
                        {
                            "colors": [
                                "#299c46",
                                "#e5ac0e",
                                "#bf1b00",
                                "#ffffff"
                            ],
                            "decimals": 2,
                            "enabled": true,
                            "label": "OVERRIDE 1",
                            "metricName": "/.*/",
                            "operatorName": "current",
                            "prefix": "",
                            "sanitizeURLEnabled": false,
                            "scaledDecimals": null,
                            "suffix": "",
                            "thresholds": [
                                {
                                    "color": "#299c46",
                                    "state": 0,
                                    "value": 20
                                },
                                {
                                    "color": "#e5ac0e",
                                    "state": 1,
                                    "value": 50
                                },
                                {
                                    "color": "#bf1b00",
                                    "state": 2,
                                    "value": 100
                                }
                            ],
                            "unitFormat": "ops"
                        }
                    ],
                    "targets": [
                        {
                            "expr": "avg by (service_name) (rate(mongodb_mongod_metrics_document_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",state=\"returned\"}[$interval]) or \nirate(mongodb_mongod_metrics_document_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",state=\"returned\"}[5m]))",
                            "format": "time_series",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}}",
                            "refId": "A"
                        }
                    ],
                    "timeFrom": null,
                    "timeShift": null,
                    "title": "Document Return Operations",
                    "type": "grafana-polystat-panel"
                },
                {
                    "aliasColors": {},
                    "bars": false,
                    "dashLength": 10,
                    "dashes": false,
                    "datasource": null,
                    "decimals": 2,
                    "description": "This panel shows the actual number of documents that were affected on average during the given time period. A single *update* command for example could perform updates on hundreds or thousands of documents. If these counters seem high, you might want to check your statements to make sure they are properly operating against the documents that they should.",
                    "editable": true,
                    "error": false,
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
                        "y": 44
                    },
                    "height": "250px",
                    "hiddenSeries": false,
                    "id": 177,
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
                    "lines": false,
                    "linewidth": 2,
                    "links": [],
                    "nullPointMode": "null",
                    "options": {
                        "dataLinks": [
                            {
                                "targetBlank": true,
                                "title": "MongoDB Instance Summary - ${__series.name}",
                                "url": "/graph/d/mongodb-instance-summary/mongodb-instance-summary?var-service_name=${__series.name}&$__url_time_range"
                            }
                        ]
                    },
                    "paceLength": 10,
                    "percentage": false,
                    "pointradius": 1,
                    "points": true,
                    "renderer": "flot",
                    "seriesOverrides": [
                        {
                            "alias": "Avg",
                            "color": "#C4162A",
                            "fill": 0,
                            "legend": false,
                            "lines": true,
                            "points": false
                        }
                    ],
                    "spaceLength": 10,
                    "stack": false,
                    "steppedLine": false,
                    "targets": [
                        {
                            "expr": "topk(5,avg by (service_name) ((rate(mongodb_mongod_metrics_document_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",state=\"updated\"}[$interval]) or \nirate(mongodb_mongod_metrics_document_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",state=\"updated\"}[5m]))))",
                            "format": "time_series",
                            "hide": false,
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}}",
                            "refId": "J",
                            "step": 300
                        },
                        {
                            "expr": "avg(rate(mongodb_mongod_metrics_document_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",state=\"updated\"}[$interval]) or \nirate(mongodb_mongod_metrics_document_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",state=\"updated\"}[5m]))",
                            "format": "time_series",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "Avg",
                            "refId": "A"
                        }
                    ],
                    "thresholds": [],
                    "timeFrom": null,
                    "timeRegions": [],
                    "timeShift": null,
                    "title": "Top 5 Document Update Operations",
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
                    "colors": [
                        "#299c46",
                        "rgba(237, 129, 40, 0.89)",
                        "#d44a3a"
                    ],
                    "datasource": null,
                    "gridPos": {
                        "h": 8,
                        "w": 12,
                        "x": 12,
                        "y": 44
                    },
                    "id": 178,
                    "links": [],
                    "options": {},
                    "polystat": {
                        "animationSpeed": 2500,
                        "columnAutoSize": true,
                        "columns": "",
                        "displayLimit": 100,
                        "fontAutoScale": true,
                        "fontSize": 12,
                        "fontType": "Roboto",
                        "globalDecimals": 2,
                        "globalDisplayMode": "all",
                        "globalDisplayTextTriggeredEmpty": "OK",
                        "globalOperatorName": "avg",
                        "globalUnitFormat": "short",
                        "gradientEnabled": true,
                        "hexagonSortByDirection": "asc",
                        "hexagonSortByField": "name",
                        "maxMetrics": 0,
                        "polygonBorderColor": "black",
                        "polygonBorderSize": 2,
                        "polygonGlobalFillColor": "#0a50a1",
                        "radius": "",
                        "radiusAutoSize": true,
                        "rowAutoSize": true,
                        "rows": "",
                        "shape": "hexagon_pointed_top",
                        "tooltipDisplayMode": "all",
                        "tooltipDisplayTextTriggeredEmpty": "OK",
                        "tooltipFontSize": 12,
                        "tooltipFontType": "Roboto",
                        "tooltipPrimarySortDirection": "desc",
                        "tooltipPrimarySortField": "thresholdLevel",
                        "tooltipSecondarySortDirection": "desc",
                        "tooltipSecondarySortField": "value",
                        "tooltipTimestampEnabled": true
                    },
                    "savedComposites": [],
                    "savedOverrides": [
                        {
                            "colors": [
                                "#299c46",
                                "#e5ac0e",
                                "#bf1b00",
                                "#ffffff"
                            ],
                            "decimals": 2,
                            "enabled": true,
                            "label": "OVERRIDE 1",
                            "metricName": "/.*/",
                            "operatorName": "current",
                            "prefix": "",
                            "sanitizeURLEnabled": false,
                            "scaledDecimals": null,
                            "suffix": "",
                            "thresholds": [
                                {
                                    "color": "#299c46",
                                    "state": 0,
                                    "value": 20
                                },
                                {
                                    "color": "#e5ac0e",
                                    "state": 1,
                                    "value": 50
                                },
                                {
                                    "color": "#bf1b00",
                                    "state": 2,
                                    "value": 100
                                }
                            ],
                            "unitFormat": "ops"
                        }
                    ],
                    "targets": [
                        {
                            "expr": "avg by (service_name) (rate(mongodb_mongod_metrics_document_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",state=\"updated\"}[$interval]) or \nirate(mongodb_mongod_metrics_document_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",state=\"updated\"}[5m]))",
                            "format": "time_series",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}}",
                            "refId": "A"
                        }
                    ],
                    "timeFrom": null,
                    "timeShift": null,
                    "title": "Document Update Operations",
                    "type": "grafana-polystat-panel"
                }
            ],
            "title": "Document Operations Detail",
            "type": "row"
        },
        {
            "collapsed": true,
            "datasource": null,
            "gridPos": {
                "h": 1,
                "w": 24,
                "x": 0,
                "y": 20
            },
            "id": 134,
            "panels": [
                {
                    "aliasColors": {},
                    "bars": false,
                    "dashLength": 10,
                    "dashes": false,
                    "datasource": null,
                    "decimals": 2,
                    "description": "This shows the number of read and write operations that are waiting due to for a lock. Consistently small values here should not be of concern, but if the values are consistently high the queries causing long lock times should be tracked down and fixed.",
                    "editable": true,
                    "error": false,
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
                        "y": 21
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
                    "lines": false,
                    "linewidth": 2,
                    "links": [],
                    "nullPointMode": "null",
                    "options": {
                        "dataLinks": [
                            {
                                "targetBlank": true,
                                "title": "MongoDB Instance Summary - ${__series.name}",
                                "url": "/graph/d/mongodb-instance-summary/mongodb-instance-summary?var-service_name=${__series.name}&$__url_time_range"
                            }
                        ]
                    },
                    "paceLength": 10,
                    "percentage": false,
                    "pointradius": 1,
                    "points": true,
                    "renderer": "flot",
                    "seriesOverrides": [
                        {
                            "alias": "Avg",
                            "color": "#C4162A",
                            "fill": 0,
                            "legend": false,
                            "lines": true,
                            "points": false
                        }
                    ],
                    "spaceLength": 10,
                    "stack": false,
                    "steppedLine": false,
                    "targets": [
                        {
                            "expr": "topk(5,avg by (service_name) ((max_over_time(mongodb_mongod_global_lock_current_queue{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"reader\"}[$interval]) or \nmax_over_time(mongodb_mongod_global_lock_current_queue{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"reader\"}[5m]))))",
                            "format": "time_series",
                            "hide": false,
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}}",
                            "refId": "J",
                            "step": 300
                        },
                        {
                            "expr": "avg(max_over_time(mongodb_mongod_global_lock_current_queue{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"reader\"}[$interval]) or \nmax_over_time(mongodb_mongod_global_lock_current_queue{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"reader\"}[5m]))",
                            "format": "time_series",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "Avg",
                            "refId": "A"
                        }
                    ],
                    "thresholds": [],
                    "timeFrom": null,
                    "timeRegions": [],
                    "timeShift": null,
                    "title": "Top 5 Queued Read Operations",
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
                    "colors": [
                        "#299c46",
                        "rgba(237, 129, 40, 0.89)",
                        "#d44a3a"
                    ],
                    "datasource": null,
                    "gridPos": {
                        "h": 8,
                        "w": 12,
                        "x": 12,
                        "y": 21
                    },
                    "id": 181,
                    "links": [],
                    "options": {},
                    "polystat": {
                        "animationSpeed": 2500,
                        "columnAutoSize": true,
                        "columns": "",
                        "displayLimit": 100,
                        "fontAutoScale": true,
                        "fontSize": 12,
                        "fontType": "Roboto",
                        "globalDecimals": 2,
                        "globalDisplayMode": "all",
                        "globalDisplayTextTriggeredEmpty": "OK",
                        "globalOperatorName": "avg",
                        "globalUnitFormat": "short",
                        "gradientEnabled": true,
                        "hexagonSortByDirection": "asc",
                        "hexagonSortByField": "name",
                        "maxMetrics": 0,
                        "polygonBorderColor": "black",
                        "polygonBorderSize": 2,
                        "polygonGlobalFillColor": "#0a50a1",
                        "radius": "",
                        "radiusAutoSize": true,
                        "rowAutoSize": true,
                        "rows": "",
                        "shape": "hexagon_pointed_top",
                        "tooltipDisplayMode": "all",
                        "tooltipDisplayTextTriggeredEmpty": "OK",
                        "tooltipFontSize": 12,
                        "tooltipFontType": "Roboto",
                        "tooltipPrimarySortDirection": "desc",
                        "tooltipPrimarySortField": "thresholdLevel",
                        "tooltipSecondarySortDirection": "desc",
                        "tooltipSecondarySortField": "value",
                        "tooltipTimestampEnabled": true
                    },
                    "savedComposites": [],
                    "savedOverrides": [
                        {
                            "colors": [
                                "#299c46",
                                "#e5ac0e",
                                "#bf1b00",
                                "#ffffff"
                            ],
                            "decimals": 2,
                            "enabled": true,
                            "label": "OVERRIDE 1",
                            "metricName": "/.*/",
                            "operatorName": "current",
                            "prefix": "",
                            "sanitizeURLEnabled": false,
                            "scaledDecimals": null,
                            "suffix": "",
                            "thresholds": [
                                {
                                    "color": "#299c46",
                                    "state": 0,
                                    "value": 20
                                },
                                {
                                    "color": "#e5ac0e",
                                    "state": 1,
                                    "value": 50
                                },
                                {
                                    "color": "#bf1b00",
                                    "state": 2,
                                    "value": 100
                                }
                            ],
                            "unitFormat": "ops"
                        }
                    ],
                    "targets": [
                        {
                            "expr": "avg by (service_name) (max_over_time(mongodb_mongod_global_lock_current_queue{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"reader\"}[$interval]) or \nmax_over_time(mongodb_mongod_global_lock_current_queue{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"reader\"}[5m]))",
                            "format": "time_series",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}}",
                            "refId": "A"
                        }
                    ],
                    "timeFrom": null,
                    "timeShift": null,
                    "title": "Queued Read Operations",
                    "type": "grafana-polystat-panel"
                },
                {
                    "aliasColors": {},
                    "bars": false,
                    "dashLength": 10,
                    "dashes": false,
                    "datasource": null,
                    "decimals": 2,
                    "description": "This shows the number of read and write operations that are waiting due to for a lock. Consistently small values here should not be of concern, but if the values are consistently high the queries causing long lock times should be tracked down and fixed.",
                    "editable": true,
                    "error": false,
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
                        "y": 29
                    },
                    "height": "250px",
                    "hiddenSeries": false,
                    "id": 182,
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
                    "lines": false,
                    "linewidth": 2,
                    "links": [],
                    "nullPointMode": "null",
                    "options": {
                        "dataLinks": [
                            {
                                "targetBlank": true,
                                "title": "MongoDB Instance Summary - ${__series.name}",
                                "url": "/graph/d/mongodb-instance-summary/mongodb-instance-summary?var-service_name=${__series.name}&$__url_time_range"
                            }
                        ]
                    },
                    "paceLength": 10,
                    "percentage": false,
                    "pointradius": 1,
                    "points": true,
                    "renderer": "flot",
                    "seriesOverrides": [
                        {
                            "alias": "Avg",
                            "color": "#C4162A",
                            "fill": 0,
                            "legend": false,
                            "lines": true,
                            "points": false
                        }
                    ],
                    "spaceLength": 10,
                    "stack": false,
                    "steppedLine": false,
                    "targets": [
                        {
                            "expr": "topk(5,avg by (service_name) ((max_over_time(mongodb_mongod_global_lock_current_queue{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"writer\"}[$interval]) or \nmax_over_time(mongodb_mongod_global_lock_current_queue{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"writer\"}[5m]))))",
                            "format": "time_series",
                            "hide": false,
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}}",
                            "refId": "J",
                            "step": 300
                        },
                        {
                            "expr": "avg(max_over_time(mongodb_mongod_global_lock_current_queue{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"writer\"}[$interval]) or \nmax_over_time(mongodb_mongod_global_lock_current_queue{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"writer\"}[5m]))",
                            "format": "time_series",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "Avg",
                            "refId": "A"
                        }
                    ],
                    "thresholds": [],
                    "timeFrom": null,
                    "timeRegions": [],
                    "timeShift": null,
                    "title": "Top 5 Queued Write Operations",
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
                    "colors": [
                        "#299c46",
                        "rgba(237, 129, 40, 0.89)",
                        "#d44a3a"
                    ],
                    "datasource": null,
                    "gridPos": {
                        "h": 8,
                        "w": 12,
                        "x": 12,
                        "y": 29
                    },
                    "id": 183,
                    "links": [],
                    "options": {},
                    "polystat": {
                        "animationSpeed": 2500,
                        "columnAutoSize": true,
                        "columns": "",
                        "displayLimit": 100,
                        "fontAutoScale": true,
                        "fontSize": 12,
                        "fontType": "Roboto",
                        "globalDecimals": 2,
                        "globalDisplayMode": "all",
                        "globalDisplayTextTriggeredEmpty": "OK",
                        "globalOperatorName": "avg",
                        "globalUnitFormat": "short",
                        "gradientEnabled": true,
                        "hexagonSortByDirection": "asc",
                        "hexagonSortByField": "name",
                        "maxMetrics": 0,
                        "polygonBorderColor": "black",
                        "polygonBorderSize": 2,
                        "polygonGlobalFillColor": "#0a50a1",
                        "radius": "",
                        "radiusAutoSize": true,
                        "rowAutoSize": true,
                        "rows": "",
                        "shape": "hexagon_pointed_top",
                        "tooltipDisplayMode": "all",
                        "tooltipDisplayTextTriggeredEmpty": "OK",
                        "tooltipFontSize": 12,
                        "tooltipFontType": "Roboto",
                        "tooltipPrimarySortDirection": "desc",
                        "tooltipPrimarySortField": "thresholdLevel",
                        "tooltipSecondarySortDirection": "desc",
                        "tooltipSecondarySortField": "value",
                        "tooltipTimestampEnabled": true
                    },
                    "savedComposites": [],
                    "savedOverrides": [
                        {
                            "colors": [
                                "#299c46",
                                "#e5ac0e",
                                "#bf1b00",
                                "#ffffff"
                            ],
                            "decimals": 2,
                            "enabled": true,
                            "label": "OVERRIDE 1",
                            "metricName": "/.*/",
                            "operatorName": "current",
                            "prefix": "",
                            "sanitizeURLEnabled": false,
                            "scaledDecimals": null,
                            "suffix": "",
                            "thresholds": [
                                {
                                    "color": "#299c46",
                                    "state": 0,
                                    "value": 20
                                },
                                {
                                    "color": "#e5ac0e",
                                    "state": 1,
                                    "value": 50
                                },
                                {
                                    "color": "#bf1b00",
                                    "state": 2,
                                    "value": 100
                                }
                            ],
                            "unitFormat": "ops"
                        }
                    ],
                    "targets": [
                        {
                            "expr": "avg by (service_name) (max_over_time(mongodb_mongod_global_lock_current_queue{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"writer\"}[$interval]) or \nmax_over_time(mongodb_mongod_global_lock_current_queue{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"writer\"}[5m]))",
                            "format": "time_series",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}}",
                            "refId": "A"
                        }
                    ],
                    "timeFrom": null,
                    "timeShift": null,
                    "title": "Queued Write Operations",
                    "type": "grafana-polystat-panel"
                }
            ],
            "title": "Queued Operations Detail",
            "type": "row"
        },
        {
            "collapsed": true,
            "datasource": null,
            "gridPos": {
                "h": 1,
                "w": 24,
                "x": 0,
                "y": 21
            },
            "id": 122,
            "panels": [
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
                        "y": 22
                    },
                    "height": "250px",
                    "hiddenSeries": false,
                    "id": 202,
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
                    "lines": false,
                    "linewidth": 2,
                    "links": [],
                    "nullPointMode": "null",
                    "options": {
                        "dataLinks": [
                            {
                                "targetBlank": true,
                                "title": "MongoDB Instance Summary - ${__series.name}",
                                "url": "/graph/d/mongodb-instance-summary/mongodb-instance-summary?var-service_name=${__series.name}&$__url_time_range"
                            }
                        ]
                    },
                    "paceLength": 10,
                    "percentage": false,
                    "pointradius": 1,
                    "points": true,
                    "renderer": "flot",
                    "seriesOverrides": [
                        {
                            "alias": "Avg",
                            "color": "#C4162A",
                            "fill": 0,
                            "legend": false,
                            "lines": true,
                            "points": false
                        }
                    ],
                    "spaceLength": 10,
                    "stack": false,
                    "steppedLine": false,
                    "targets": [
                        {
                            "expr": "topk(5,avg by (service_name) ((rate(mongodb_mongod_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"msg\"}[$interval]) or \nirate(mongodb_mongod_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"msg\"}[5m]) or\nrate(mongodb_mongos_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"msg\"}[$interval]) or \nirate(mongodb_mongos_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"msg\"}[5m]) or\nrate(mongodb_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"msg\"}[$interval]) or \nirate(mongodb_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"msg\"}[5m]))))",
                            "format": "time_series",
                            "hide": false,
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}}",
                            "refId": "J",
                            "step": 300
                        },
                        {
                            "expr": "avg(rate(mongodb_mongod_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"msg\"}[$interval]) or \nirate(mongodb_mongod_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"msg\"}[5m]) or\nrate(mongodb_mongos_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"msg\"}[$interval]) or \nirate(mongodb_mongos_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"msg\"}[5m]) or\nrate(mongodb_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"msg\"}[$interval]) or \nirate(mongodb_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"msg\"}[5m]))",
                            "format": "time_series",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "Avg",
                            "refId": "A"
                        }
                    ],
                    "thresholds": [],
                    "timeFrom": null,
                    "timeRegions": [],
                    "timeShift": null,
                    "title": "Top 5 Assert Msg Events",
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
                            "show": true
                        }
                    ],
                    "yaxis": {
                        "align": false,
                        "alignLevel": null
                    }
                },
                {
                    "colors": [
                        "#299c46",
                        "rgba(237, 129, 40, 0.89)",
                        "#d44a3a"
                    ],
                    "datasource": null,
                    "description": "",
                    "gridPos": {
                        "h": 8,
                        "w": 12,
                        "x": 12,
                        "y": 22
                    },
                    "id": 203,
                    "links": [],
                    "options": {},
                    "polystat": {
                        "animationSpeed": 2500,
                        "columnAutoSize": true,
                        "columns": "",
                        "displayLimit": 100,
                        "fontAutoScale": true,
                        "fontSize": 12,
                        "fontType": "Roboto",
                        "globalDecimals": 2,
                        "globalDisplayMode": "all",
                        "globalDisplayTextTriggeredEmpty": "OK",
                        "globalOperatorName": "avg",
                        "globalUnitFormat": "short",
                        "gradientEnabled": true,
                        "hexagonSortByDirection": "asc",
                        "hexagonSortByField": "name",
                        "maxMetrics": 0,
                        "polygonBorderColor": "black",
                        "polygonBorderSize": 2,
                        "polygonGlobalFillColor": "#0a50a1",
                        "radius": "",
                        "radiusAutoSize": true,
                        "rowAutoSize": true,
                        "rows": "",
                        "shape": "hexagon_pointed_top",
                        "tooltipDisplayMode": "all",
                        "tooltipDisplayTextTriggeredEmpty": "OK",
                        "tooltipFontSize": 12,
                        "tooltipFontType": "Roboto",
                        "tooltipPrimarySortDirection": "desc",
                        "tooltipPrimarySortField": "thresholdLevel",
                        "tooltipSecondarySortDirection": "desc",
                        "tooltipSecondarySortField": "value",
                        "tooltipTimestampEnabled": true
                    },
                    "savedComposites": [],
                    "savedOverrides": [
                        {
                            "colors": [
                                "#299c46",
                                "#e5ac0e",
                                "#bf1b00",
                                "#ffffff"
                            ],
                            "decimals": 2,
                            "enabled": true,
                            "label": "OVERRIDE 1",
                            "metricName": "/.*/",
                            "operatorName": "current",
                            "prefix": "",
                            "sanitizeURLEnabled": false,
                            "scaledDecimals": null,
                            "suffix": "",
                            "thresholds": [
                                {
                                    "color": "#299c46",
                                    "state": 0,
                                    "value": 20
                                },
                                {
                                    "color": "#e5ac0e",
                                    "state": 1,
                                    "value": 50
                                },
                                {
                                    "color": "#bf1b00",
                                    "state": 2,
                                    "value": 100
                                }
                            ],
                            "unitFormat": "ops"
                        }
                    ],
                    "targets": [
                        {
                            "expr": "avg by (service_name) (rate(mongodb_mongod_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"msg\"}[$interval]) or \nirate(mongodb_mongod_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"msg\"}[5m]) or\nrate(mongodb_mongos_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"msg\"}[$interval]) or \nirate(mongodb_mongos_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"msg\"}[5m]) or\nrate(mongodb_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"msg\"}[$interval]) or \nirate(mongodb_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"msg\"}[5m]))",
                            "format": "time_series",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}}",
                            "refId": "A"
                        }
                    ],
                    "timeFrom": null,
                    "timeShift": null,
                    "title": "Assert Msg Events",
                    "type": "grafana-polystat-panel"
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
                        "y": 30
                    },
                    "height": "250px",
                    "hiddenSeries": false,
                    "id": 198,
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
                    "lines": false,
                    "linewidth": 2,
                    "links": [],
                    "nullPointMode": "null",
                    "options": {
                        "dataLinks": [
                            {
                                "targetBlank": true,
                                "title": "MongoDB Instance Summary - ${__series.name}",
                                "url": "/graph/d/mongodb-instance-summary/mongodb-instance-summary?var-service_name=${__series.name}&$__url_time_range"
                            }
                        ]
                    },
                    "paceLength": 10,
                    "percentage": false,
                    "pointradius": 1,
                    "points": true,
                    "renderer": "flot",
                    "seriesOverrides": [
                        {
                            "alias": "Avg",
                            "color": "#C4162A",
                            "fill": 0,
                            "legend": false,
                            "lines": true,
                            "points": false
                        }
                    ],
                    "spaceLength": 10,
                    "stack": false,
                    "steppedLine": false,
                    "targets": [
                        {
                            "expr": "topk(5,avg by (service_name) ((rate(mongodb_mongod_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"regular\"}[$interval]) or \nirate(mongodb_mongod_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"regular\"}[5m]) or\nrate(mongodb_mongos_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"regular\"}[$interval]) or \nirate(mongodb_mongos_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"regular\"}[5m]) or\nrate(mongodb_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"regular\"}[$interval]) or \nirate(mongodb_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"regular\"}[5m]))))",
                            "format": "time_series",
                            "hide": false,
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}}",
                            "refId": "J",
                            "step": 300
                        },
                        {
                            "expr": "avg(rate(mongodb_mongod_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"regular\"}[$interval]) or \nirate(mongodb_mongod_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"regular\"}[5m]) or\nrate(mongodb_mongos_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"regular\"}[$interval]) or \nirate(mongodb_mongos_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"regular\"}[5m]) or\nrate(mongodb_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"regular\"}[$interval]) or \nirate(mongodb_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"regular\"}[5m]))",
                            "format": "time_series",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "Avg",
                            "refId": "A"
                        }
                    ],
                    "thresholds": [],
                    "timeFrom": null,
                    "timeRegions": [],
                    "timeShift": null,
                    "title": "Top 5 Assert Regular Events",
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
                            "show": true
                        }
                    ],
                    "yaxis": {
                        "align": false,
                        "alignLevel": null
                    }
                },
                {
                    "colors": [
                        "#299c46",
                        "rgba(237, 129, 40, 0.89)",
                        "#d44a3a"
                    ],
                    "datasource": null,
                    "description": "",
                    "gridPos": {
                        "h": 8,
                        "w": 12,
                        "x": 12,
                        "y": 30
                    },
                    "id": 204,
                    "links": [],
                    "options": {},
                    "polystat": {
                        "animationSpeed": 2500,
                        "columnAutoSize": true,
                        "columns": "",
                        "displayLimit": 100,
                        "fontAutoScale": true,
                        "fontSize": 12,
                        "fontType": "Roboto",
                        "globalDecimals": 2,
                        "globalDisplayMode": "all",
                        "globalDisplayTextTriggeredEmpty": "OK",
                        "globalOperatorName": "avg",
                        "globalUnitFormat": "short",
                        "gradientEnabled": true,
                        "hexagonSortByDirection": "asc",
                        "hexagonSortByField": "name",
                        "maxMetrics": 0,
                        "polygonBorderColor": "black",
                        "polygonBorderSize": 2,
                        "polygonGlobalFillColor": "#0a50a1",
                        "radius": "",
                        "radiusAutoSize": true,
                        "rowAutoSize": true,
                        "rows": "",
                        "shape": "hexagon_pointed_top",
                        "tooltipDisplayMode": "all",
                        "tooltipDisplayTextTriggeredEmpty": "OK",
                        "tooltipFontSize": 12,
                        "tooltipFontType": "Roboto",
                        "tooltipPrimarySortDirection": "desc",
                        "tooltipPrimarySortField": "thresholdLevel",
                        "tooltipSecondarySortDirection": "desc",
                        "tooltipSecondarySortField": "value",
                        "tooltipTimestampEnabled": true
                    },
                    "savedComposites": [],
                    "savedOverrides": [
                        {
                            "colors": [
                                "#299c46",
                                "#e5ac0e",
                                "#bf1b00",
                                "#ffffff"
                            ],
                            "decimals": 2,
                            "enabled": true,
                            "label": "OVERRIDE 1",
                            "metricName": "/.*/",
                            "operatorName": "current",
                            "prefix": "",
                            "sanitizeURLEnabled": false,
                            "scaledDecimals": null,
                            "suffix": "",
                            "thresholds": [
                                {
                                    "color": "#299c46",
                                    "state": 0,
                                    "value": 20
                                },
                                {
                                    "color": "#e5ac0e",
                                    "state": 1,
                                    "value": 50
                                },
                                {
                                    "color": "#bf1b00",
                                    "state": 2,
                                    "value": 100
                                }
                            ],
                            "unitFormat": "ops"
                        }
                    ],
                    "targets": [
                        {
                            "expr": "avg by (service_name) (rate(mongodb_mongod_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"regular\"}[$interval]) or \nirate(mongodb_mongod_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"regular\"}[5m]) or\nrate(mongodb_mongos_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"regular\"}[$interval]) or \nirate(mongodb_mongos_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"regular\"}[5m]) or\nrate(mongodb_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"regular\"}[$interval]) or \nirate(mongodb_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"regular\"}[5m]))",
                            "format": "time_series",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}}",
                            "refId": "A"
                        }
                    ],
                    "timeFrom": null,
                    "timeShift": null,
                    "title": " Assert Regular Events",
                    "type": "grafana-polystat-panel"
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
                        "y": 38
                    },
                    "height": "250px",
                    "hiddenSeries": false,
                    "id": 200,
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
                    "lines": false,
                    "linewidth": 2,
                    "links": [],
                    "nullPointMode": "null",
                    "options": {
                        "dataLinks": [
                            {
                                "targetBlank": true,
                                "title": "MongoDB Instance Summary - ${__series.name}",
                                "url": "/graph/d/mongodb-instance-summary/mongodb-instance-summary?var-service_name=${__series.name}&$__url_time_range"
                            }
                        ]
                    },
                    "paceLength": 10,
                    "percentage": false,
                    "pointradius": 1,
                    "points": true,
                    "renderer": "flot",
                    "seriesOverrides": [
                        {
                            "alias": "Avg",
                            "color": "#C4162A",
                            "fill": 0,
                            "legend": false,
                            "lines": true,
                            "points": false
                        }
                    ],
                    "spaceLength": 10,
                    "stack": false,
                    "steppedLine": false,
                    "targets": [
                        {
                            "expr": "topk(5,avg by (service_name) ((rate(mongodb_mongod_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"rollovers\"}[$interval]) or \nirate(mongodb_mongod_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"rollovers\"}[5m]) or\nrate(mongodb_mongos_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"rollovers\"}[$interval]) or \nirate(mongodb_mongos_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"rollovers\"}[5m]) or\nrate(mongodb_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"rollovers\"}[$interval]) or \nirate(mongodb_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"rollovers\"}[5m]))))",
                            "format": "time_series",
                            "hide": false,
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}}",
                            "refId": "J",
                            "step": 300
                        },
                        {
                            "expr": "avg(rate(mongodb_mongod_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"rollovers\"}[$interval]) or \nirate(mongodb_mongod_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"rollovers\"}[5m]) or\nrate(mongodb_mongos_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"rollovers\"}[$interval]) or \nirate(mongodb_mongos_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"rollovers\"}[5m]) or\nrate(mongodb_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"rollovers\"}[$interval]) or \nirate(mongodb_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"rollovers\"}[5m]))",
                            "format": "time_series",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "Avg",
                            "refId": "A"
                        }
                    ],
                    "thresholds": [],
                    "timeFrom": null,
                    "timeRegions": [],
                    "timeShift": null,
                    "title": "Top 5 Assert Rollovers Events",
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
                            "show": true
                        }
                    ],
                    "yaxis": {
                        "align": false,
                        "alignLevel": null
                    }
                },
                {
                    "colors": [
                        "#299c46",
                        "rgba(237, 129, 40, 0.89)",
                        "#d44a3a"
                    ],
                    "datasource": null,
                    "description": "",
                    "gridPos": {
                        "h": 8,
                        "w": 12,
                        "x": 12,
                        "y": 38
                    },
                    "id": 205,
                    "links": [],
                    "options": {},
                    "polystat": {
                        "animationSpeed": 2500,
                        "columnAutoSize": true,
                        "columns": "",
                        "displayLimit": 100,
                        "fontAutoScale": true,
                        "fontSize": 12,
                        "fontType": "Roboto",
                        "globalDecimals": 2,
                        "globalDisplayMode": "all",
                        "globalDisplayTextTriggeredEmpty": "OK",
                        "globalOperatorName": "avg",
                        "globalUnitFormat": "short",
                        "gradientEnabled": true,
                        "hexagonSortByDirection": "asc",
                        "hexagonSortByField": "name",
                        "maxMetrics": 0,
                        "polygonBorderColor": "black",
                        "polygonBorderSize": 2,
                        "polygonGlobalFillColor": "#0a50a1",
                        "radius": "",
                        "radiusAutoSize": true,
                        "rowAutoSize": true,
                        "rows": "",
                        "shape": "hexagon_pointed_top",
                        "tooltipDisplayMode": "all",
                        "tooltipDisplayTextTriggeredEmpty": "OK",
                        "tooltipFontSize": 12,
                        "tooltipFontType": "Roboto",
                        "tooltipPrimarySortDirection": "desc",
                        "tooltipPrimarySortField": "thresholdLevel",
                        "tooltipSecondarySortDirection": "desc",
                        "tooltipSecondarySortField": "value",
                        "tooltipTimestampEnabled": true
                    },
                    "savedComposites": [],
                    "savedOverrides": [
                        {
                            "colors": [
                                "#299c46",
                                "#e5ac0e",
                                "#bf1b00",
                                "#ffffff"
                            ],
                            "decimals": 2,
                            "enabled": true,
                            "label": "OVERRIDE 1",
                            "metricName": "/.*/",
                            "operatorName": "current",
                            "prefix": "",
                            "sanitizeURLEnabled": false,
                            "scaledDecimals": null,
                            "suffix": "",
                            "thresholds": [
                                {
                                    "color": "#299c46",
                                    "state": 0,
                                    "value": 20
                                },
                                {
                                    "color": "#e5ac0e",
                                    "state": 1,
                                    "value": 50
                                },
                                {
                                    "color": "#bf1b00",
                                    "state": 2,
                                    "value": 100
                                }
                            ],
                            "unitFormat": "ops"
                        }
                    ],
                    "targets": [
                        {
                            "expr": "avg by (service_name) (rate(mongodb_mongod_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"rollovers\"}[$interval]) or \nirate(mongodb_mongod_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"rollovers\"}[5m]) or\nrate(mongodb_mongos_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"rollovers\"}[$interval]) or \nirate(mongodb_mongos_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"rollovers\"}[5m]) or\nrate(mongodb_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"rollovers\"}[$interval]) or \nirate(mongodb_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"rollovers\"}[5m]))",
                            "format": "time_series",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}}",
                            "refId": "A"
                        }
                    ],
                    "timeFrom": null,
                    "timeShift": null,
                    "title": "Assert Rollovers Events",
                    "type": "grafana-polystat-panel"
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
                        "y": 46
                    },
                    "height": "250px",
                    "hiddenSeries": false,
                    "id": 201,
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
                    "lines": false,
                    "linewidth": 2,
                    "links": [],
                    "nullPointMode": "null",
                    "options": {
                        "dataLinks": [
                            {
                                "targetBlank": true,
                                "title": "MongoDB Instance Summary - ${__series.name}",
                                "url": "/graph/d/mongodb-instance-summary/mongodb-instance-summary?var-service_name=${__series.name}&$__url_time_range"
                            }
                        ]
                    },
                    "paceLength": 10,
                    "percentage": false,
                    "pointradius": 1,
                    "points": true,
                    "renderer": "flot",
                    "seriesOverrides": [
                        {
                            "alias": "Avg",
                            "color": "#C4162A",
                            "fill": 0,
                            "legend": false,
                            "lines": true,
                            "points": false
                        }
                    ],
                    "spaceLength": 10,
                    "stack": false,
                    "steppedLine": false,
                    "targets": [
                        {
                            "expr": "topk(5,avg by (service_name) ((rate(mongodb_mongod_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"user\"}[$interval]) or \nirate(mongodb_mongod_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"user\"}[5m]) or\nrate(mongodb_mongos_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"user\"}[$interval]) or \nirate(mongodb_mongos_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"user\"}[5m]) or\nrate(mongodb_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"user\"}[$interval]) or \nirate(mongodb_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"user\"}[5m]))))",
                            "format": "time_series",
                            "hide": false,
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}}",
                            "refId": "J",
                            "step": 300
                        },
                        {
                            "expr": "avg(rate(mongodb_mongod_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"user\"}[$interval]) or \nirate(mongodb_mongod_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"user\"}[5m]) or\nrate(mongodb_mongos_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"user\"}[$interval]) or \nirate(mongodb_mongos_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"user\"}[5m]) or\nrate(mongodb_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"user\"}[$interval]) or \nirate(mongodb_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"user\"}[5m]))",
                            "format": "time_series",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "Avg",
                            "refId": "A"
                        }
                    ],
                    "thresholds": [],
                    "timeFrom": null,
                    "timeRegions": [],
                    "timeShift": null,
                    "title": "Top 5 Assert User Events",
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
                    "colors": [
                        "#299c46",
                        "rgba(237, 129, 40, 0.89)",
                        "#d44a3a"
                    ],
                    "datasource": null,
                    "description": "",
                    "gridPos": {
                        "h": 8,
                        "w": 12,
                        "x": 12,
                        "y": 46
                    },
                    "id": 206,
                    "links": [],
                    "options": {},
                    "polystat": {
                        "animationSpeed": 2500,
                        "columnAutoSize": true,
                        "columns": "",
                        "displayLimit": 100,
                        "fontAutoScale": true,
                        "fontSize": 12,
                        "fontType": "Roboto",
                        "globalDecimals": 2,
                        "globalDisplayMode": "all",
                        "globalDisplayTextTriggeredEmpty": "OK",
                        "globalOperatorName": "avg",
                        "globalUnitFormat": "short",
                        "gradientEnabled": true,
                        "hexagonSortByDirection": "asc",
                        "hexagonSortByField": "name",
                        "maxMetrics": 0,
                        "polygonBorderColor": "black",
                        "polygonBorderSize": 2,
                        "polygonGlobalFillColor": "#0a50a1",
                        "radius": "",
                        "radiusAutoSize": true,
                        "rowAutoSize": true,
                        "rows": "",
                        "shape": "hexagon_pointed_top",
                        "tooltipDisplayMode": "all",
                        "tooltipDisplayTextTriggeredEmpty": "OK",
                        "tooltipFontSize": 12,
                        "tooltipFontType": "Roboto",
                        "tooltipPrimarySortDirection": "desc",
                        "tooltipPrimarySortField": "thresholdLevel",
                        "tooltipSecondarySortDirection": "desc",
                        "tooltipSecondarySortField": "value",
                        "tooltipTimestampEnabled": true
                    },
                    "savedComposites": [],
                    "savedOverrides": [
                        {
                            "colors": [
                                "#299c46",
                                "#e5ac0e",
                                "#bf1b00",
                                "#ffffff"
                            ],
                            "decimals": 2,
                            "enabled": true,
                            "label": "OVERRIDE 1",
                            "metricName": "/.*/",
                            "operatorName": "current",
                            "prefix": "",
                            "sanitizeURLEnabled": false,
                            "scaledDecimals": null,
                            "suffix": "",
                            "thresholds": [
                                {
                                    "color": "#299c46",
                                    "state": 0,
                                    "value": 20
                                },
                                {
                                    "color": "#e5ac0e",
                                    "state": 1,
                                    "value": 50
                                },
                                {
                                    "color": "#bf1b00",
                                    "state": 2,
                                    "value": 100
                                }
                            ],
                            "unitFormat": "ops"
                        }
                    ],
                    "targets": [
                        {
                            "expr": "avg by (service_name) (rate(mongodb_mongod_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"user\"}[$interval]) or \nirate(mongodb_mongod_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"user\"}[5m]) or\nrate(mongodb_mongos_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"user\"}[$interval]) or \nirate(mongodb_mongos_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"user\"}[5m]) or\nrate(mongodb_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"user\"}[$interval]) or \nirate(mongodb_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"user\"}[5m]))",
                            "format": "time_series",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}}",
                            "refId": "A"
                        }
                    ],
                    "timeFrom": null,
                    "timeShift": null,
                    "title": "Assert User Events",
                    "type": "grafana-polystat-panel"
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
                        "y": 54
                    },
                    "height": "250px",
                    "hiddenSeries": false,
                    "id": 199,
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
                    "lines": false,
                    "linewidth": 2,
                    "links": [],
                    "nullPointMode": "null",
                    "options": {
                        "dataLinks": [
                            {
                                "targetBlank": true,
                                "title": "MongoDB Instance Summary - ${__series.name}",
                                "url": "/graph/d/mongodb-instance-summary/mongodb-instance-summary?var-service_name=${__series.name}&$__url_time_range"
                            }
                        ]
                    },
                    "paceLength": 10,
                    "percentage": false,
                    "pointradius": 1,
                    "points": true,
                    "renderer": "flot",
                    "seriesOverrides": [
                        {
                            "alias": "Avg",
                            "color": "#C4162A",
                            "fill": 0,
                            "legend": false,
                            "lines": true,
                            "points": false
                        }
                    ],
                    "spaceLength": 10,
                    "stack": false,
                    "steppedLine": false,
                    "targets": [
                        {
                            "expr": "topk(5,avg by (service_name) ((rate(mongodb_mongod_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"warning\"}[$interval]) or \nirate(mongodb_mongod_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"warning\"}[5m]) or\nrate(mongodb_mongos_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"warning\"}[$interval]) or \nirate(mongodb_mongos_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"warning\"}[5m]) or\nrate(mongodb_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"warning\"}[$interval]) or \nirate(mongodb_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"warning\"}[5m]))))",
                            "format": "time_series",
                            "hide": false,
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}}",
                            "refId": "J",
                            "step": 300
                        },
                        {
                            "expr": "avg(rate(mongodb_mongod_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"warning\"}[$interval]) or \nirate(mongodb_mongod_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"warning\"}[5m]) or\nrate(mongodb_mongos_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"warning\"}[$interval]) or \nirate(mongodb_mongos_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"warning\"}[5m]) or\nrate(mongodb_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"warning\"}[$interval]) or \nirate(mongodb_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"warning\"}[5m]))",
                            "format": "time_series",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "Avg",
                            "refId": "A"
                        }
                    ],
                    "thresholds": [],
                    "timeFrom": null,
                    "timeRegions": [],
                    "timeShift": null,
                    "title": "Top 5 Assert Msg Events",
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
                            "show": true
                        }
                    ],
                    "yaxis": {
                        "align": false,
                        "alignLevel": null
                    }
                },
                {
                    "colors": [
                        "#299c46",
                        "rgba(237, 129, 40, 0.89)",
                        "#d44a3a"
                    ],
                    "datasource": null,
                    "description": "",
                    "gridPos": {
                        "h": 8,
                        "w": 12,
                        "x": 12,
                        "y": 54
                    },
                    "id": 207,
                    "links": [],
                    "options": {},
                    "polystat": {
                        "animationSpeed": 2500,
                        "columnAutoSize": true,
                        "columns": "",
                        "displayLimit": 100,
                        "fontAutoScale": true,
                        "fontSize": 12,
                        "fontType": "Roboto",
                        "globalDecimals": 2,
                        "globalDisplayMode": "all",
                        "globalDisplayTextTriggeredEmpty": "OK",
                        "globalOperatorName": "avg",
                        "globalUnitFormat": "short",
                        "gradientEnabled": true,
                        "hexagonSortByDirection": "asc",
                        "hexagonSortByField": "name",
                        "maxMetrics": 0,
                        "polygonBorderColor": "black",
                        "polygonBorderSize": 2,
                        "polygonGlobalFillColor": "#0a50a1",
                        "radius": "",
                        "radiusAutoSize": true,
                        "rowAutoSize": true,
                        "rows": "",
                        "shape": "hexagon_pointed_top",
                        "tooltipDisplayMode": "all",
                        "tooltipDisplayTextTriggeredEmpty": "OK",
                        "tooltipFontSize": 12,
                        "tooltipFontType": "Roboto",
                        "tooltipPrimarySortDirection": "desc",
                        "tooltipPrimarySortField": "thresholdLevel",
                        "tooltipSecondarySortDirection": "desc",
                        "tooltipSecondarySortField": "value",
                        "tooltipTimestampEnabled": true
                    },
                    "savedComposites": [],
                    "savedOverrides": [
                        {
                            "colors": [
                                "#299c46",
                                "#e5ac0e",
                                "#bf1b00",
                                "#ffffff"
                            ],
                            "decimals": 2,
                            "enabled": true,
                            "label": "OVERRIDE 1",
                            "metricName": "/.*/",
                            "operatorName": "current",
                            "prefix": "",
                            "sanitizeURLEnabled": false,
                            "scaledDecimals": null,
                            "suffix": "",
                            "thresholds": [
                                {
                                    "color": "#299c46",
                                    "state": 0,
                                    "value": 20
                                },
                                {
                                    "color": "#e5ac0e",
                                    "state": 1,
                                    "value": 50
                                },
                                {
                                    "color": "#bf1b00",
                                    "state": 2,
                                    "value": 100
                                }
                            ],
                            "unitFormat": "ops"
                        }
                    ],
                    "targets": [
                        {
                            "expr": "avg by (service_name) (rate(mongodb_mongod_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"warning\"}[$interval]) or \nirate(mongodb_mongod_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"warning\"}[5m]) or\nrate(mongodb_mongos_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"warning\"}[$interval]) or \nirate(mongodb_mongos_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"warning\"}[5m]) or\nrate(mongodb_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"warning\"}[$interval]) or \nirate(mongodb_asserts_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"warning\"}[5m]))",
                            "format": "time_series",
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}}",
                            "refId": "A"
                        }
                    ],
                    "timeFrom": null,
                    "timeShift": null,
                    "title": "Assert Warning Events",
                    "type": "grafana-polystat-panel"
                }
            ],
            "title": "Assert Events Detail",
            "type": "row"
        }
    ],
    "refresh": "1m",
    "schemaVersion": 27,
    "style": "dark",
    "tags": [
        "MongoDB",
        "Percona",
        "Services"
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
                "description": null,
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
                "definition": "label_values(mongodb_up{vmcluster=\"$vmcluster\", namespace=~\"$namespace\"}, service)",
                "description": null,
                "error": null,
                "hide": 0,
                "includeAll": true,
                "label": "Service Name",
                "multi": true,
                "multiFormat": "glob",
                "name": "service_name",
                "options": [],
                "query": {
                    "query": "label_values(mongodb_up{vmcluster=\"$vmcluster\", namespace=~\"$namespace\"}, service)",
                    "refId": "Metrics-service_name-Variable-Query"
                },
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
            "1s",
            "5s",
            "1m",
            "5m",
            "1h",
            "6h",
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
    "title": "MongoDB Instances Overview",
    "uid": "mongodb-instance-overview",
    "version": 1
}
`
