package monitor

const MongoInmemoryDetail = `
{
    "editable": false,
    "gnetId": null,
    "graphTooltip": 1,
    "id": null,
    "iteration": 1573814423755,
    "links": [],
    "panels": [
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
            "description": "This panel shows the amount of data currently being stored InMemory. This data is in its uncompressed format and differs from how the data is stored on disk or in the file system cache. This value will always be lower than the counter shown in the *InMemory Max Data Size* panel.",
            "editable": true,
            "error": false,
            "format": "bytes",
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
                "y": 2
            },
            "hideTimeOverride": true,
            "id": 62,
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
            "options": {},
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
                    "expr": "avg by (service_name) (mongodb_mongod_wiredtiger_cache_bytes{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"total\"})",
                    "format": "time_series",
                    "interval": "5m",
                    "intervalFactor": 1,
                    "refId": "A",
                    "step": 300
                }
            ],
            "thresholds": "",
            "timeFrom": "1m",
            "title": "InMemory Data Size",
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
                "rgba(245, 54, 54, 0.9)",
                "rgba(237, 129, 40, 0.89)",
                "rgba(50, 172, 45, 0.97)"
            ],
            "datasource": null,
            "decimals": 2,
            "description": "This is the maximum size that the InMemory data can grow to and this value can be changed from the default value by setting the storage.wiredTiger.inMemory.cacheSizeGB value in the config file or passing in the --inMemoryCacheSizeGB parameter on the command line.",
            "editable": true,
            "error": false,
            "format": "bytes",
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
                "y": 2
            },
            "hideTimeOverride": true,
            "id": 63,
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
            "options": {},
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
                    "expr": "avg by (service_name) (mongodb_mongod_wiredtiger_cache_max_bytes{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"})",
                    "interval": "5m",
                    "intervalFactor": 1,
                    "refId": "A",
                    "step": 300
                }
            ],
            "thresholds": "",
            "timeFrom": "1m",
            "title": "InMemory Max Data Size",
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
                "rgba(50, 172, 45, 0.97)",
                "rgba(237, 129, 40, 0.89)",
                "rgba(245, 54, 54, 0.9)"
            ],
            "datasource": null,
            "decimals": 0,
            "description": "This panel shows the amount of memory left available for your InMemory databases.",
            "editable": true,
            "error": false,
            "format": "percentunit",
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
                "y": 2
            },
            "height": "125px",
            "id": 69,
            "interval": "$interval",
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
            "options": {},
            "postfix": "",
            "postfixFontSize": "50%",
            "prefix": "",
            "prefixFontSize": "80%",
            "rangeMaps": [
                {
                    "from": "null",
                    "text": "N/A",
                    "to": "null"
                }
            ],
            "sparkline": {
                "fillColor": "rgba(31, 118, 189, 0.18)",
                "full": true,
                "lineColor": "rgb(31, 120, 193)",
                "show": true,
                "ymax": 1,
                "ymin": 0
            },
            "tableColumn": "",
            "targets": [
                {
                    "calculatedInterval": "10m",
                    "datasourceErrors": {},
                    "errors": {},
                    "expr": "avg by (service_name) (1-(max(mongodb_mongod_wiredtiger_cache_bytes{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"total\"})/max(mongodb_mongod_wiredtiger_cache_max_bytes{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"})))",
                    "interval": "5m",
                    "intervalFactor": 1,
                    "legendFormat": "",
                    "metric": "node_mem",
                    "refId": "A",
                    "step": 300
                }
            ],
            "thresholds": "90,95",
            "title": "InMemory Available",
            "type": "singlestat",
            "valueFontSize": "50%",
            "valueMaps": [],
            "valueName": "current"
        },
        {
            "cacheTimeout": null,
            "colorBackground": false,
            "colorValue": true,
            "colors": [
                "rgba(42, 150, 37, 0.97)",
                "rgba(237, 129, 40, 0.89)",
                "rgba(245, 54, 54, 0.9)"
            ],
            "datasource": null,
            "decimals": 0,
            "description": "This panel shows the percentage of InMemory pages that have been changed and not yet had the modified data consolidated.",
            "editable": true,
            "error": false,
            "format": "percentunit",
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
                "y": 2
            },
            "height": "125px",
            "id": 70,
            "interval": "$interval",
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
            "options": {},
            "postfix": "",
            "postfixFontSize": "50%",
            "prefix": "",
            "prefixFontSize": "80%",
            "rangeMaps": [
                {
                    "from": "null",
                    "text": "N/A",
                    "to": "null"
                }
            ],
            "sparkline": {
                "fillColor": "rgba(189, 154, 31, 0.18)",
                "full": true,
                "lineColor": "rgb(193, 153, 31)",
                "show": true,
                "ymax": 1,
                "ymin": 0
            },
            "tableColumn": "",
            "targets": [
                {
                    "calculatedInterval": "10m",
                    "datasourceErrors": {},
                    "errors": {},
                    "expr": "avg by (service_name) (avg(mongodb_mongod_wiredtiger_cache_pages{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"dirty\"})/avg(mongodb_mongod_wiredtiger_cache_pages{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"total\"}))",
                    "interval": "5m",
                    "intervalFactor": 1,
                    "legendFormat": "",
                    "metric": "node_mem",
                    "refId": "A",
                    "step": 300
                }
            ],
            "thresholds": "30,50",
            "title": "InMemory Dirty Pages",
            "type": "singlestat",
            "valueFontSize": "50%",
            "valueMaps": [],
            "valueName": "current"
        },
        {
            "aliasColors": {},
            "bars": false,
            "dashLength": 10,
            "dashes": false,
            "datasource": null,
            "decimals": 2,
            "description": "WiredTiger internal transactions",
            "editable": true,
            "error": false,
            "fill": 2,
            "fillGradient": 0,
            "grid": {},
            "gridPos": {
                "h": 8,
                "w": 12,
                "x": 0,
                "y": 4
            },
            "hiddenSeries": false,
            "id": 52,
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
            "options": {
                "dataLinks": []
            },
            "paceLength": 10,
            "percentage": false,
            "pointradius": 5,
            "points": false,
            "renderer": "flot",
            "seriesOverrides": [],
            "spaceLength": 10,
            "stack": false,
            "steppedLine": false,
            "targets": [
                {
                    "expr": "avg by (service_name,type) (rate(mongodb_mongod_wiredtiger_transactions_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or irate(mongodb_mongod_wiredtiger_transactions_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]))",
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "{{type}}",
                    "refId": "A",
                    "step": 300
                }
            ],
            "thresholds": [],
            "timeFrom": null,
            "timeRegions": [],
            "timeShift": null,
            "title": "InMemory Transactions",
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
            "aliasColors": {},
            "bars": false,
            "dashLength": 10,
            "dashes": false,
            "datasource": null,
            "decimals": 2,
            "description": "Configured max and current size of the WiredTiger cache.",
            "editable": true,
            "error": false,
            "fill": 2,
            "fillGradient": 0,
            "grid": {},
            "gridPos": {
                "h": 8,
                "w": 12,
                "x": 12,
                "y": 4
            },
            "hiddenSeries": false,
            "id": 45,
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
                "dataLinks": []
            },
            "paceLength": 10,
            "percentage": false,
            "pointradius": 5,
            "points": false,
            "renderer": "flot",
            "seriesOverrides": [
                {
                    "alias": "Percent Overhead",
                    "yaxis": 2
                }
            ],
            "spaceLength": 10,
            "stack": false,
            "steppedLine": false,
            "targets": [
                {
                    "expr": "avg by (service_name) (mongodb_mongod_wiredtiger_cache_max_bytes{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"})",
                    "hide": false,
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "Maximum",
                    "refId": "C",
                    "step": 300
                },
                {
                    "expr": "avg by (service_name) (mongodb_mongod_wiredtiger_cache_bytes{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"total\"})",
                    "hide": false,
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "Used",
                    "refId": "A",
                    "step": 300
                }
            ],
            "thresholds": [],
            "timeFrom": null,
            "timeRegions": [],
            "timeShift": null,
            "title": "InMemory Capacity",
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
            "description": "Internal WiredTiger storage engine cursors and sessions currently open.",
            "editable": true,
            "error": false,
            "fill": 2,
            "fillGradient": 0,
            "grid": {},
            "gridPos": {
                "h": 8,
                "w": 12,
                "x": 0,
                "y": 12
            },
            "hiddenSeries": false,
            "id": 60,
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
                "dataLinks": []
            },
            "paceLength": 10,
            "percentage": false,
            "pointradius": 5,
            "points": false,
            "renderer": "flot",
            "seriesOverrides": [],
            "spaceLength": 10,
            "stack": false,
            "steppedLine": false,
            "targets": [
                {
                    "expr": "avg by (service_name) (mongodb_mongod_wiredtiger_session_open_cursors_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"})",
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "Cursors",
                    "metric": "",
                    "refId": "A",
                    "step": 300
                },
                {
                    "expr": "avg by (service_name) (mongodb_mongod_wiredtiger_session_open_sessions_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"})",
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "Sessions",
                    "metric": "",
                    "refId": "B",
                    "step": 300
                }
            ],
            "thresholds": [],
            "timeFrom": null,
            "timeRegions": [],
            "timeShift": null,
            "title": "InMemory Sessions",
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
            "aliasColors": {},
            "bars": false,
            "dashLength": 10,
            "dashes": false,
            "datasource": null,
            "decimals": 2,
            "description": "Pages in the WiredTiger cache",
            "editable": true,
            "error": false,
            "fill": 2,
            "fillGradient": 0,
            "grid": {},
            "gridPos": {
                "h": 8,
                "w": 12,
                "x": 12,
                "y": 12
            },
            "hiddenSeries": false,
            "id": 68,
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
                "dataLinks": []
            },
            "paceLength": 10,
            "percentage": false,
            "pointradius": 5,
            "points": false,
            "renderer": "flot",
            "seriesOverrides": [],
            "spaceLength": 10,
            "stack": false,
            "steppedLine": false,
            "targets": [
                {
                    "expr": "avg by (service_name,type) (mongodb_mongod_wiredtiger_cache_pages{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"})",
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "{{type}}",
                    "refId": "A",
                    "step": 300
                }
            ],
            "thresholds": [],
            "timeFrom": null,
            "timeRegions": [],
            "timeShift": null,
            "title": "InMemory Pages",
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
                    "format": "short",
                    "label": "",
                    "logBase": 1,
                    "max": null,
                    "min": 0,
                    "show": true
                },
                {
                    "decimals": 2,
                    "format": "short",
                    "label": "",
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
            "description": "A WT 'ticket' is assigned out for every operation running simultaneously in the WT storage engine. \"Tickets available\" = hardcoded high value - \"Tickets Out\".",
            "editable": true,
            "error": false,
            "fill": 2,
            "fillGradient": 0,
            "grid": {},
            "gridPos": {
                "h": 8,
                "w": 12,
                "x": 0,
                "y": 20
            },
            "hiddenSeries": false,
            "id": 55,
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
                "dataLinks": []
            },
            "paceLength": 10,
            "percentage": false,
            "pointradius": 5,
            "points": false,
            "renderer": "flot",
            "seriesOverrides": [
                {
                    "alias": "/^write_/",
                    "transform": "negative-Y"
                }
            ],
            "spaceLength": 10,
            "stack": false,
            "steppedLine": false,
            "targets": [
                {
                    "expr": "avg by (service_name,type) (mongodb_mongod_wiredtiger_concurrent_transactions_total_tickets{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}-mongodb_mongod_wiredtiger_concurrent_transactions_out_tickets{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"})",
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "{{type}}",
                    "metric": "",
                    "refId": "A",
                    "step": 300
                }
            ],
            "thresholds": [],
            "timeFrom": null,
            "timeRegions": [],
            "timeShift": null,
            "title": "InMemory Concurrency Tickets",
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
            "description": "Operations queued due to a lock",
            "editable": true,
            "error": false,
            "fill": 2,
            "fillGradient": 0,
            "grid": {},
            "gridPos": {
                "h": 8,
                "w": 12,
                "x": 12,
                "y": 20
            },
            "hiddenSeries": false,
            "id": 40,
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
                "dataLinks": []
            },
            "paceLength": 10,
            "percentage": false,
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
            "description": "Mixed metrics: Docs per second inserted, updated, deleted or returned on any type of node (primary or secondary); + replicated write Ops/sec; + TTL deletes per second.",
            "editable": true,
            "error": false,
            "fill": 2,
            "fillGradient": 0,
            "grid": {},
            "gridPos": {
                "h": 8,
                "w": 12,
                "x": 0,
                "y": 28
            },
            "hiddenSeries": false,
            "id": 36,
            "legend": {
                "alignAsTable": true,
                "avg": true,
                "current": false,
                "hideZero": true,
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
                "dataLinks": []
            },
            "paceLength": 10,
            "percentage": false,
            "pointradius": 5,
            "points": false,
            "renderer": "flot",
            "seriesOverrides": [],
            "spaceLength": 10,
            "stack": true,
            "steppedLine": false,
            "targets": [
                {
                    "expr": "avg by (service_name,state) (rate(mongodb_mongod_metrics_document_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or irate(mongodb_mongod_metrics_document_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]))",
                    "hide": false,
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "{{state}}",
                    "refId": "J",
                    "step": 300
                },
                {
                    "expr": "avg by (service_name) (rate(mongodb_mongod_op_counters_repl_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"delete\"}[$interval]) or irate(mongodb_mongod_op_counters_repl_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"delete\"}[5m]))",
                    "hide": false,
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "repl_deleted",
                    "refId": "A",
                    "step": 300
                },
                {
                    "expr": "avg by (service_name) (rate(mongodb_mongod_op_counters_repl_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"update\"}[$interval]) or irate(mongodb_mongod_op_counters_repl_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"update\"}[5m]))",
                    "hide": false,
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "repl_updated",
                    "refId": "B",
                    "step": 300
                },
                {
                    "expr": "avg by (service_name) (rate(mongodb_mongod_op_counters_repl_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"insert\"}[$interval]) or irate(mongodb_mongod_op_counters_repl_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", type=\"insert\"}[5m]))",
                    "hide": false,
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "repl_inserted",
                    "refId": "C",
                    "step": 300
                },
                {
                    "expr": "avg by (service_name) (rate(mongodb_mongod_metrics_ttl_deleted_documents_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or irate(mongodb_mongod_metrics_ttl_deleted_documents_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]))",
                    "hide": false,
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "ttl_deleted",
                    "refId": "D",
                    "step": 300
                }
            ],
            "thresholds": [],
            "timeFrom": null,
            "timeRegions": [],
            "timeShift": null,
            "title": "Document Changes",
            "tooltip": {
                "msResolution": false,
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
            "aliasColors": {},
            "bars": false,
            "dashLength": 10,
            "dashes": false,
            "datasource": null,
            "decimals": 2,
            "description": "This panel shows the number of pages that have been evicted from the WiredTiger cache for the given time period. The InMemory storage engine only evicts modified pages which signals a compaction of the data and removal of the dirty pages.",
            "editable": true,
            "error": false,
            "fill": 6,
            "fillGradient": 0,
            "grid": {},
            "gridPos": {
                "h": 8,
                "w": 12,
                "x": 12,
                "y": 28
            },
            "hiddenSeries": false,
            "id": 53,
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
                "dataLinks": []
            },
            "paceLength": 10,
            "percentage": false,
            "pointradius": 5,
            "points": false,
            "renderer": "flot",
            "seriesOverrides": [],
            "spaceLength": 10,
            "stack": true,
            "steppedLine": false,
            "targets": [
                {
                    "expr": "avg by (service_name) (rate(mongodb_mongod_wiredtiger_cache_evicted_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"modified\"}[$interval]) or irate(mongodb_mongod_wiredtiger_cache_evicted_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\",type=\"modified\"}[5m]))",
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "Evicted",
                    "refId": "A",
                    "step": 300
                }
            ],
            "thresholds": [],
            "timeFrom": null,
            "timeRegions": [],
            "timeShift": null,
            "title": "InMemory Cache Eviction",
            "tooltip": {
                "msResolution": false,
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
                    "format": "short",
                    "label": "Pages / sec",
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
            "aliasColors": {},
            "bars": false,
            "dashLength": 10,
            "dashes": false,
            "datasource": null,
            "decimals": 2,
            "description": "This panel shows the number of objects (both data (scanned_objects) and index (scanned)) as well as the number of documents that were moved to a new location due to the size of the document growing. Moved documents only apply to the MMAPv1 storage engine.",
            "editable": true,
            "error": false,
            "fill": 2,
            "fillGradient": 0,
            "grid": {},
            "gridPos": {
                "h": 8,
                "w": 12,
                "x": 0,
                "y": 36
            },
            "hiddenSeries": false,
            "id": 32,
            "legend": {
                "alignAsTable": true,
                "avg": true,
                "current": false,
                "hideZero": true,
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
                "dataLinks": []
            },
            "paceLength": 10,
            "percentage": false,
            "pointradius": 5,
            "points": false,
            "renderer": "flot",
            "seriesOverrides": [],
            "spaceLength": 10,
            "stack": false,
            "steppedLine": false,
            "targets": [
                {
                    "expr": "avg by (service_name,state) (rate(mongodb_mongod_metrics_query_executor_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or irate(mongodb_mongod_metrics_query_executor_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]))",
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "{{state}}",
                    "metric": "",
                    "refId": "A",
                    "step": 300
                },
                {
                    "expr": "avg by (service_name) (rate(mongodb_mongod_metrics_record_moves_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or irate(mongodb_mongod_metrics_record_moves_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]))",
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
            "description": "Unix or Window memory page faults. Not necessarily from mongodb.",
            "editable": true,
            "error": false,
            "fill": 2,
            "fillGradient": 0,
            "grid": {},
            "gridPos": {
                "h": 8,
                "w": 12,
                "x": 12,
                "y": 36
            },
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
            "options": {
                "dataLinks": []
            },
            "paceLength": 10,
            "percentage": false,
            "pointradius": 5,
            "points": false,
            "renderer": "flot",
            "seriesOverrides": [],
            "spaceLength": 10,
            "stack": false,
            "steppedLine": false,
            "targets": [
                {
                    "expr": "avg by (service_name) (rate(mongodb_mongod_extra_info_page_faults_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or irate(mongodb_mongod_extra_info_page_faults_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]) or rate(mongodb_extra_info_page_faults_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or irate(mongodb_extra_info_page_faults_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]))",
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
            "collapsed": true,
            "datasource": null,
            "gridPos": {
                "h": 1,
                "w": 24,
                "x": 0,
                "y": 44
            },
            "id": 1009,
            "panels": [
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
                        "x": 0,
                        "y": 45
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
                    "options": {},
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
                    "thresholds": "300,3600",
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
                        "w": 8,
                        "x": 8,
                        "y": 45
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
                    "options": {},
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
                        "w": 8,
                        "x": 16,
                        "y": 45
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
                    "options": {},
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
                    "aliasColors": {},
                    "bars": false,
                    "dashLength": 10,
                    "dashes": false,
                    "datasource": null,
                    "decimals": 2,
                    "description": "This shows the number of active connections on the server. Keep in mind the hard limit on the maximum number of connections set by your distribution.\n\nYou can read more about the [connection numbers](https://docs.mongodb.com/manual/administration/analyzing-mongodb-performance/#number-of-connections).",
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
                        "y": 47
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
                        "dataLinks": []
                    },
                    "paceLength": 10,
                    "percentage": false,
                    "pointradius": 5,
                    "points": false,
                    "renderer": "flot",
                    "seriesOverrides": [],
                    "spaceLength": 10,
                    "stack": false,
                    "steppedLine": false,
                    "targets": [
                        {
                            "expr": "avg by (service_name) (max_over_time(mongodb_mongod_connections{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"current\"}[$interval]) or max_over_time(mongodb_mongod_connections{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"current\"}[5m]) or\nmax_over_time(mongodb_mongos_connections{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"current\"}[$interval]) or\nmax_over_time(mongodb_mongos_connections{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"current\"}[5m]) or\nmax_over_time(mongodb_connections{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"current\"}[$interval]) or\nmax_over_time(mongodb_connections{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\", state=\"current\"}[5m]))",
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
                    "description": "This shows the number of open cursors for each shard in the cluster. A cursor in MongoDB is a pointer to the result of a given query that can be iterated over. By default a cursor times out after 10 minutes.",
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
                        "x": 12,
                        "y": 47
                    },
                    "height": "250px",
                    "hiddenSeries": false,
                    "id": 1013,
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
                    "options": {
                        "dataLinks": []
                    },
                    "paceLength": 10,
                    "percentage": false,
                    "pointradius": 5,
                    "points": false,
                    "renderer": "flot",
                    "seriesOverrides": [],
                    "spaceLength": 10,
                    "stack": false,
                    "steppedLine": false,
                    "targets": [
                        {
                            "expr": "avg by (service_name,state) (max_over_time(mongodb_mongod_metrics_cursor_open{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or\nmax_over_time(mongodb_mongod_metrics_cursor_open{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]) or\nmax_over_time(mongodb_mongod_cursors{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or\nmax_over_time(mongodb_mongod_cursors{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]) or\nmax_over_time(mongodb_mongos_metrics_cursor_open{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or \nmax_over_time(mongodb_mongos_metrics_cursor_open{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]) or\nmax_over_time(mongodb_mongos_cursors{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[$interval]) or\nmax_over_time(mongodb_mongos_cursors{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", service=~\"$service_name\"}[5m]))",
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
                }
            ],
            "title": "MongoDB Summary",
            "type": "row"
        }
    ],
    "refresh": "1m",
    "schemaVersion": 21,
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
                "includeAll": false,
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
                "definition": "",
                "hide": 0,
                "includeAll": true,
                "label": "Service Name",
                "multi": false,
                "multiFormat": "glob",
                "name": "service_name",
                "options": [],
                "query": "label_values(mongodb_mongod_storage_engine{vmcluster=\"$vmcluster\", namespace=~\"$namespace\",engine=\"wiredTiger\"}, service)",
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
    "title": "MongoDB InMemory Details",
    "uid": "mongodb-inmemory",
    "version": 1
}
`
