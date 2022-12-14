package monitor

const MongoClusterSummary = `
{
    "editable": false,
    "gnetId": null,
    "graphTooltip": 1,
    "id": null,
    "iteration": 1573487532506,
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
            "id": 1069,
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
            "description": "Sharding is a method for distributing data across multiple machines. MongoDB uses sharding to support deployments with very large data sets and high throughput operations.",
            "editable": true,
            "error": false,
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
            "hideTimeOverride": true,
            "id": 35,
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
                    "expr": "max(mongodb_mongos_sharding_databases_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", type=\"partitioned\"})",
                    "interval": "5m",
                    "intervalFactor": 1,
                    "legendFormat": "Shards",
                    "refId": "A",
                    "step": 300
                }
            ],
            "thresholds": "",
            "timeFrom": "1m",
            "title": "Sharded DBs",
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
            "decimals": 0,
            "description": "A MongoDB sharded cluster deployment can contain collections that are either unsharded (the default when created) or sharded.",
            "editable": true,
            "error": false,
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
                "x": 4,
                "y": 3
            },
            "hideTimeOverride": true,
            "id": 39,
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
                    "expr": "max(mongodb_mongos_sharding_databases_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", type=\"unpartitioned\"})",
                    "interval": "5m",
                    "intervalFactor": 1,
                    "legendFormat": "Shards",
                    "refId": "A",
                    "step": 300
                }
            ],
            "thresholds": "",
            "timeFrom": "1m",
            "title": "Unsharded DBs",
            "type": "singlestat",
            "valueFontSize": "50%",
            "valueMaps": [],
            "valueName": "current"
        },
        {
            "activePatternIndex": 1,
            "datasource": null,
            "debug_mode": false,
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
                "null_value": "-",
                "pattern": "*",
                "row_col_wrapper": "_",
                "row_name": "_series_",
                "textColors": "red|orange|green",
                "textColors_overrides": "0->red|2->green|1->yellow",
                "thresholds": "70,90",
                "time_based_thresholds": [],
                "tooltipTemplate": "Series : _series_ <br/>Row Name : _row_name_ <br/>Col Name : _col_name_ <br/>Value : _value_",
                "transform_values": "_value_|_value_|_value_",
                "transform_values_overrides": "0->down|1->up",
                "valueName": "current"
            },
            "default_title_for_rows": "DB Name",
            "description": "MongoDB stores documents in collections. Collections are analogous to tables in relational databases.",
            "fieldConfig": {
                "defaults": {},
                "overrides": []
            },
            "gridPos": {
                "h": 8,
                "w": 8,
                "x": 8,
                "y": 3
            },
            "hide_headers": false,
            "id": 1043,
            "patterns": [
                {
                    "bgColors": "blue|blue|blue",
                    "bgColors_overrides": "0->green|2->red|1->yellow",
                    "clickable_cells_link": "/graph/d/mongodb-replicaset-summary/mongodb-replset-summary?var-interval=$__auto_interval_interval&var-replset=_1_",
                    "col_name": "_1_",
                    "decimals": "0",
                    "defaultBGColor": "",
                    "defaultTextColor": "",
                    "delimiter": "|",
                    "displayTemplate": "_value_",
                    "enable_bgColor": true,
                    "enable_bgColor_overrides": false,
                    "enable_clickable_cells": true,
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
                    "name": "Collection",
                    "null_color": "blue",
                    "null_textcolor": "white",
                    "null_value": "0",
                    "pattern": ".*Collections",
                    "row_col_wrapper": "_",
                    "row_name": "_0_",
                    "textColors": "red|orange|green",
                    "textColors_overrides": "0->red|2->green|1->yellow",
                    "thresholds": "70,90",
                    "time_based_thresholds": [],
                    "tooltipTemplate": "-",
                    "transform_values": "_value_|_value_|_value_",
                    "transform_values_overrides": "0->down|1->up",
                    "valueName": "current"
                },
                {
                    "bgColors": "green|orange|red",
                    "bgColors_overrides": "0->green|2->red|1->yellow",
                    "clickable_cells_link": "",
                    "col_name": "_2_",
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
                    "name": "Size",
                    "null_color": "darkred",
                    "null_textcolor": "black",
                    "null_value": "No data",
                    "pattern": ".*Size",
                    "row_col_wrapper": "_",
                    "row_name": "_1_",
                    "textColors": "red|orange|green",
                    "textColors_overrides": "0->red|2->green|1->yellow",
                    "thresholds": "70,90",
                    "time_based_thresholds": [],
                    "tooltipTemplate": "Series : _series_ <br/>Row Name : _row_name_ <br/>Col Name : _col_name_ <br/>Value : _value_",
                    "transform_values": "_value_|_value_|_value_",
                    "transform_values_overrides": "0->down|1->up",
                    "valueName": "avg"
                }
            ],
            "row_col_wrapper": "_",
            "sorting_props": {
                "col_index": -1,
                "direction": "desc"
            },
            "targets": [
                {
                    "expr": "max by (db,shard) (mongodb_mongos_db_collections_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", db!~\"admin|config\"})",
                    "format": "time_series",
                    "hide": false,
                    "instant": true,
                    "interval": "$interval",
                    "legendFormat": "{{db}} | {{shard}} | Collections",
                    "refId": "A"
                },
                {
                    "expr": "mongodb_mongos_db_data_size_bytes{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", db!~\"admin|config\"}",
                    "hide": true,
                    "instant": true,
                    "interval": "$intgerval",
                    "legendFormat": "{{db}} | {{shard}} | Size",
                    "refId": "B"
                }
            ],
            "timeFrom": null,
            "timeShift": null,
            "title": "Amount of Collections in Shards",
            "type": "yesoreyeram-boomtable-panel"
        },
        {
            "activePatternIndex": 0,
            "datasource": null,
            "debug_mode": false,
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
                "row_name": "_series_",
                "textColors": "red|orange|green",
                "textColors_overrides": "0->red|2->green|1->yellow",
                "thresholds": "70,90",
                "time_based_thresholds": [],
                "tooltipTemplate": "Series : _series_ <br/>Row Name : _row_name_ <br/>Col Name : _col_name_ <br/>Value : _value_",
                "transform_values": "_value_|_value_|_value_",
                "transform_values_overrides": "0->down|1->up",
                "valueName": "current"
            },
            "default_title_for_rows": "DB Name",
            "description": "MongoDB stores documents in collections. Collections are analogous to tables in relational databases.",
            "fieldConfig": {
                "defaults": {},
                "overrides": []
            },
            "gridPos": {
                "h": 8,
                "w": 8,
                "x": 16,
                "y": 3
            },
            "hide_headers": false,
            "id": 1030,
            "patterns": [
                {
                    "bgColors": "blue|blue|blue",
                    "bgColors_overrides": "0->green|2->red|1->yellow",
                    "clickable_cells_link": "/graph/d/mongodb-replicaset-summary/mongodb-replset-summary?var-interval=$__auto_interval_interval&var-replset=_1_",
                    "col_name": "_1_",
                    "decimals": "0",
                    "defaultBGColor": "",
                    "defaultTextColor": "",
                    "delimiter": "|",
                    "displayTemplate": "_value_",
                    "enable_bgColor": true,
                    "enable_bgColor_overrides": false,
                    "enable_clickable_cells": true,
                    "enable_textColor": false,
                    "enable_textColor_overrides": false,
                    "enable_time_based_thresholds": false,
                    "enable_transform": false,
                    "enable_transform_overrides": false,
                    "filter": {
                        "value_above": "",
                        "value_below": ""
                    },
                    "format": "bytes",
                    "name": "Size",
                    "null_color": "blue",
                    "null_textcolor": "white",
                    "null_value": "0",
                    "pattern": ".*Collections",
                    "row_col_wrapper": "_",
                    "row_name": "_0_",
                    "textColors": "red|orange|green",
                    "textColors_overrides": "0->red|2->green|1->yellow",
                    "thresholds": "70,90",
                    "time_based_thresholds": [],
                    "tooltipTemplate": "-",
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
                    "expr": "max by (db,shard) (mongodb_mongos_db_data_size_bytes{vmcluster=\"$vmcluster\", namespace=~\"$namespace\",db!~\"admin|config\"})",
                    "format": "time_series",
                    "hide": false,
                    "instant": true,
                    "interval": "$interval",
                    "legendFormat": "{{db}} | {{shard}} | Collections",
                    "refId": "A"
                }
            ],
            "text_alignment_header": "center",
            "timeFrom": null,
            "timeShift": null,
            "title": "Size of Collections in Shards",
            "type": "yesoreyeram-boomtable-panel"
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
            "description": "A shard contains a subset of sharded data for a sharded cluster. Together, the cluster\u2019s shards hold the entire data set for the cluster.",
            "editable": true,
            "error": false,
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
                "y": 5
            },
            "hideTimeOverride": true,
            "id": 36,
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
                    "expr": "max(mongodb_mongos_sharding_shards_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\"})",
                    "interval": "5m",
                    "intervalFactor": 1,
                    "legendFormat": "Shards",
                    "refId": "A",
                    "step": 300
                }
            ],
            "thresholds": "",
            "timeFrom": "1m",
            "title": "Shards",
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
            "decimals": 0,
            "description": "A chunk consists of a subset of sharded data.",
            "editable": true,
            "error": false,
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
                "x": 4,
                "y": 5
            },
            "hideTimeOverride": true,
            "id": 11,
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
                    "expr": "sum(mongodb_mongos_sharding_chunks_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\"})",
                    "interval": "5m",
                    "intervalFactor": 1,
                    "legendFormat": "Chunks",
                    "refId": "A",
                    "step": 300
                }
            ],
            "thresholds": "",
            "timeFrom": "1m",
            "title": "Chunks",
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
            "decimals": 0,
            "description": "When you run removeShard, MongoDB drains the shard by using the balancer to move the shard\u2019s chunks to other shards in the cluster. Once the shard is drained, MongoDB removes the shard from the cluster.",
            "editable": true,
            "error": false,
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
                "y": 7
            },
            "hideTimeOverride": true,
            "id": 1028,
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
                    "expr": "max(mongodb_mongos_sharding_shards_draining_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\"})",
                    "interval": "5m",
                    "intervalFactor": 1,
                    "legendFormat": "Draining Total",
                    "refId": "A",
                    "step": 300
                }
            ],
            "thresholds": "",
            "timeFrom": "1m",
            "title": "Draining Shards",
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
            "decimals": 0,
            "description": "MongoDB stores documents in collections. Collections are analogous to tables in relational databases.",
            "editable": true,
            "error": false,
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
                "x": 4,
                "y": 7
            },
            "hideTimeOverride": true,
            "id": 10,
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
                    "expr": "max(mongodb_mongos_sharding_collections_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\"})",
                    "interval": "5m",
                    "intervalFactor": 1,
                    "legendFormat": "Shards",
                    "refId": "A",
                    "step": 300
                }
            ],
            "thresholds": "",
            "timeFrom": "1m",
            "title": "Sharded Collections",
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
                "rgba(245, 54, 54, 0.9)",
                "rgba(237, 129, 40, 0.89)",
                "rgba(50, 172, 45, 0.97)"
            ],
            "datasource": null,
            "decimals": 0,
            "description": "The MongoDB balancer is a background process that monitors the number of chunks on each shard. ",
            "editable": true,
            "error": false,
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
                "y": 9
            },
            "hideTimeOverride": true,
            "id": 5,
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
                    "expr": "max(mongodb_mongos_sharding_balancer_enabled{vmcluster=\"$vmcluster\", namespace=~\"$namespace\"})",
                    "format": "time_series",
                    "interval": "5m",
                    "intervalFactor": 1,
                    "legendFormat": "Cluster Balanced",
                    "refId": "A",
                    "step": 300
                }
            ],
            "thresholds": "0,1",
            "timeFrom": "1m",
            "title": "Balancer Enabled",
            "type": "singlestat",
            "valueFontSize": "50%",
            "valueMaps": [
                {
                    "op": "=",
                    "text": "YES",
                    "value": "1"
                },
                {
                    "op": "=",
                    "text": "NO",
                    "value": "0"
                }
            ],
            "valueName": "current"
        },
        {
            "cacheTimeout": null,
            "colorBackground": false,
            "colorValue": true,
            "colors": [
                "rgba(245, 54, 54, 0.9)",
                "rgba(237, 129, 40, 0.89)",
                "rgba(50, 172, 45, 0.97)"
            ],
            "datasource": null,
            "decimals": 0,
            "description": "When the number of chunks on a given shard reaches specific migration thresholds, the balancer attempts to automatically migrate chunks between shards and reach an equal number of chunks per shard.",
            "editable": true,
            "error": false,
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
                "x": 4,
                "y": 9
            },
            "hideTimeOverride": true,
            "id": 4,
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
                    "expr": "min(mongodb_mongos_sharding_chunks_is_balanced{vmcluster=\"$vmcluster\", namespace=~\"$namespace\"})",
                    "format": "time_series",
                    "interval": "5m",
                    "intervalFactor": 1,
                    "legendFormat": "Cluster Balanced",
                    "refId": "A",
                    "step": 300
                }
            ],
            "thresholds": "0,1",
            "timeFrom": "1m",
            "title": "Chunks Balanced",
            "type": "singlestat",
            "valueFontSize": "50%",
            "valueMaps": [
                {
                    "op": "=",
                    "text": "YES",
                    "value": "1"
                },
                {
                    "op": "=",
                    "text": "NO",
                    "value": "0"
                }
            ],
            "valueName": "current"
        },
        {
            "colors": [
                "#299c46",
                "rgba(237, 129, 40, 0.89)",
                "#d44a3a",
                "#4040a0"
            ],
            "datasource": null,
            "description": "Config services store the metadata for a sharded cluster. The metadata reflects state and organization for all data and components within the sharded cluster. The metadata includes the list of chunks on every shard and the ranges that define the chunks.",
            "fieldConfig": {
                "defaults": {},
                "overrides": []
            },
            "gridPos": {
                "h": 5,
                "w": 12,
                "x": 0,
                "y": 11
            },
            "id": 1229,
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
            "polystat": {
                "animationSpeed": 2500,
                "columnAutoSize": true,
                "columns": "",
                "displayLimit": 100,
                "fontAutoColor": true,
                "fontAutoScale": true,
                "fontSize": 12,
                "fontType": "Roboto",
                "globalDecimals": 0,
                "globalDisplayMode": "all",
                "globalDisplayTextTriggeredEmpty": "OK",
                "globalOperatorName": "current",
                "globalUnitFormat": "ops",
                "gradientEnabled": true,
                "hexagonSortByDirection": 1,
                "hexagonSortByField": "name",
                "maxMetrics": 0,
                "polygonBorderColor": "black",
                "polygonBorderSize": 2,
                "polygonGlobalFillColor": "#8F3BB8",
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
                "tooltipSecondarySortField": "name",
                "tooltipTimestampEnabled": true,
                "valueEnabled": true
            },
            "rangeMaps": [
                {
                    "from": "null",
                    "text": "N/A",
                    "to": "null"
                }
            ],
            "repeatDirection": "v",
            "savedComposites": [],
            "savedOverrides": [],
            "targets": [
                {
                    "expr": "sum by (service_name) (mongodb_connections{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", state=\"current\"} * on (service_name) group_right avg by (service_name,set) (mongodb_mongod_replset_my_state{vmcluster=\"$vmcluster\", namespace=~\"$namespace\",set!~\"$shard\"}/ mongodb_mongod_replset_my_state{vmcluster=\"$vmcluster\", namespace=~\"$namespace\",set!~\"$shard\"}))",
                    "format": "time_series",
                    "hide": true,
                    "instant": false,
                    "interval": "$interval",
                    "legendFormat": "Config - {{service_name}} ",
                    "refId": "A"
                },
                {
                    "expr": "(sum by (service_name) (rate(mongodb_op_counters_total{service_name=~\"$service_name\",type!=\"command\"}[$interval]) or irate(mongodb_op_counters_total{service_name=~\"$service_name\",type!=\"command\"}[5m])) * on (service_name) group_right avg by (service_name,set) (mongodb_mongod_replset_my_state{vmcluster=\"$vmcluster\", namespace=~\"$namespace\",set!~\"$shard\"}/ mongodb_mongod_replset_my_state{vmcluster=\"$vmcluster\", namespace=~\"$namespace\",set!~\"$shard\"}))",
                    "interval": "$interval",
                    "legendFormat": "{{service_name}}",
                    "refId": "B"
                }
            ],
            "timeFrom": null,
            "timeShift": null,
            "title": "QPS of Config Services",
            "type": "grafana-polystat-panel",
            "valueMaps": [
                {
                    "op": "=",
                    "text": "N/A",
                    "value": "null"
                }
            ]
        },
        {
            "colors": [
                "#299c46",
                "rgba(237, 129, 40, 0.89)",
                "#d44a3a",
                "#4040a0"
            ],
            "datasource": null,
            "description": "Mongos is a routing service for MongoDB shard configurations that processes queries from the application layer, and determines the location of this data in the sharded cluster, in order to complete these operations. From the perspective of the application, a mongos instance behaves identically to any other MongoDB instance.",
            "fieldConfig": {
                "defaults": {},
                "overrides": []
            },
            "gridPos": {
                "h": 5,
                "w": 12,
                "x": 12,
                "y": 11
            },
            "id": 1231,
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
            "polystat": {
                "animationSpeed": 2500,
                "columnAutoSize": true,
                "columns": "",
                "displayLimit": 100,
                "fontAutoColor": true,
                "fontAutoScale": true,
                "fontSize": 12,
                "fontType": "Roboto",
                "globalDecimals": 0,
                "globalDisplayMode": "all",
                "globalDisplayTextTriggeredEmpty": "OK",
                "globalOperatorName": "current",
                "globalUnitFormat": "ops",
                "gradientEnabled": true,
                "hexagonSortByDirection": 1,
                "hexagonSortByField": "name",
                "maxMetrics": 0,
                "polygonBorderColor": "black",
                "polygonBorderSize": 2,
                "polygonGlobalFillColor": "#8AB8FF",
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
                "tooltipSecondarySortField": "name",
                "tooltipTimestampEnabled": true,
                "valueEnabled": true
            },
            "rangeMaps": [
                {
                    "from": "null",
                    "text": "N/A",
                    "to": "null"
                }
            ],
            "repeatDirection": "v",
            "savedComposites": [],
            "savedOverrides": [
                {
                    "clickThrough": "",
                    "colors": [
                        "#299c46",
                        "#e5ac0e",
                        "#bf1b00",
                        "#ffffff"
                    ],
                    "decimals": "",
                    "enabled": true,
                    "label": "OVERRIDE 1",
                    "metricName": "",
                    "newTabEnabled": true,
                    "operatorName": "avg",
                    "prefix": "",
                    "sanitizeURLEnabled": false,
                    "scaledDecimals": null,
                    "suffix": "",
                    "thresholds": [],
                    "unitFormat": "short"
                }
            ],
            "targets": [
                {
                    "exemplar": true,
                    "expr": "(sum by (service_name) (rate(mongodb_op_counters_total{service_name=~\"$service_name\",type!=\"command\"}[$interval]) or irate(mongodb_op_counters_total{service_name=~\"$service_name\",type!=\"command\"}[5m])) * on (service_name) group_right avg by (service_name) (avg by (service_name) (mongodb_mongos_db_collections_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\"}) / avg by (service_name) (mongodb_mongos_db_collections_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\"})))",
                    "format": "time_series",
                    "instant": false,
                    "interval": "$interval",
                    "legendFormat": "{{service_name}} ",
                    "refId": "A"
                }
            ],
            "timeFrom": null,
            "timeShift": null,
            "title": "QPS of Mongos Service",
            "type": "grafana-polystat-panel",
            "valueMaps": [
                {
                    "op": "=",
                    "text": "N/A",
                    "value": "null"
                }
            ]
        },
        {
            "colors": [
                "#299c46",
                "rgba(237, 129, 40, 0.89)",
                "#d44a3a",
                "#4040a0"
            ],
            "datasource": null,
            "description": "A shard contains a subset of sharded data for a sharded cluster. Together, the cluster\u2019s shards hold the entire data set for the cluster.",
            "fieldConfig": {
                "defaults": {},
                "overrides": []
            },
            "gridPos": {
                "h": 5,
                "w": 24,
                "x": 0,
                "y": 16
            },
            "id": 1227,
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
            "polystat": {
                "animationSpeed": 2500,
                "columnAutoSize": true,
                "columns": "",
                "displayLimit": 100,
                "fontAutoColor": true,
                "fontAutoScale": true,
                "fontSize": 12,
                "fontType": "Roboto",
                "globalDecimals": 0,
                "globalDisplayMode": "all",
                "globalDisplayTextTriggeredEmpty": "OK",
                "globalOperatorName": "current",
                "globalUnitFormat": "ops",
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
                "tooltipSecondarySortField": "name",
                "tooltipTimestampEnabled": true,
                "valueEnabled": true
            },
            "rangeMaps": [
                {
                    "from": "null",
                    "text": "N/A",
                    "to": "null"
                }
            ],
            "repeat": "shard",
            "repeatDirection": "v",
            "savedComposites": [],
            "savedOverrides": [
                {
                    "clickThrough": "",
                    "colors": [
                        "#299c46",
                        "#e5ac0e",
                        "#bf1b00",
                        "#ffffff"
                    ],
                    "decimals": "",
                    "enabled": true,
                    "label": "OVERRIDE 1",
                    "metricName": "",
                    "newTabEnabled": true,
                    "operatorName": "avg",
                    "prefix": "",
                    "sanitizeURLEnabled": false,
                    "scaledDecimals": null,
                    "suffix": "",
                    "thresholds": [],
                    "unitFormat": "short"
                }
            ],
            "targets": [
                {
                    "expr": "(sum by (service_name) (rate(mongodb_op_counters_total{service_name=~\"$service_name\",type!=\"command\"}[$interval]) or irate(mongodb_op_counters_total{service_name=~\"$service_name\",type!=\"command\"}[5m])) * on (service_name) group_right avg by (service_name,set) (mongodb_mongod_replset_my_state{vmcluster=\"$vmcluster\", namespace=~\"$namespace\",set=~\"$shard\"}/ mongodb_mongod_replset_my_state{vmcluster=\"$vmcluster\", namespace=~\"$namespace\",set=~\"$shard\"}))",
                    "format": "time_series",
                    "instant": false,
                    "interval": "$interval",
                    "legendFormat": "{{service_name}} ",
                    "refId": "A"
                }
            ],
            "timeFrom": null,
            "timeShift": null,
            "title": "QPS of Services in Shard - $shard",
            "type": "grafana-polystat-panel",
            "valueMaps": [
                {
                    "op": "=",
                    "text": "N/A",
                    "value": "null"
                }
            ]
        },
        {
            "colors": [
                "#299c46",
                "rgba(237, 129, 40, 0.89)",
                "#d44a3a",
                "#4040a0"
            ],
            "datasource": null,
            "description": "A shard contains a subset of sharded data for a sharded cluster. Together, the cluster\u2019s shards hold the entire data set for the cluster.",
            "fieldConfig": {
                "defaults": {},
                "overrides": []
            },
            "gridPos": {
                "h": 5,
                "w": 24,
                "x": 0,
                "y": 21
            },
            "id": 1252,
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
            "polystat": {
                "animationSpeed": 2500,
                "columnAutoSize": true,
                "columns": "",
                "displayLimit": 100,
                "fontAutoColor": true,
                "fontAutoScale": true,
                "fontSize": 12,
                "fontType": "Roboto",
                "globalDecimals": 0,
                "globalDisplayMode": "all",
                "globalDisplayTextTriggeredEmpty": "OK",
                "globalOperatorName": "current",
                "globalUnitFormat": "ops",
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
                "tooltipSecondarySortField": "name",
                "tooltipTimestampEnabled": true,
                "valueEnabled": true
            },
            "rangeMaps": [
                {
                    "from": "null",
                    "text": "N/A",
                    "to": "null"
                }
            ],
            "repeatDirection": "v",
            "repeatIteration": 1607082131822,
            "repeatPanelId": 1227,
            "savedComposites": [],
            "savedOverrides": [
                {
                    "clickThrough": "",
                    "colors": [
                        "#299c46",
                        "#e5ac0e",
                        "#bf1b00",
                        "#ffffff"
                    ],
                    "decimals": "",
                    "enabled": true,
                    "label": "OVERRIDE 1",
                    "metricName": "",
                    "newTabEnabled": true,
                    "operatorName": "avg",
                    "prefix": "",
                    "sanitizeURLEnabled": false,
                    "scaledDecimals": null,
                    "suffix": "",
                    "thresholds": [],
                    "unitFormat": "short"
                }
            ],
            "targets": [
                {
                    "expr": "(sum by (service_name) (rate(mongodb_op_counters_total{service_name=~\"$service_name\",type!=\"command\"}[$interval]) or irate(mongodb_op_counters_total{service_name=~\"$service_name\",type!=\"command\"}[5m])) * on (service_name) group_right avg by (service_name,set) (mongodb_mongod_replset_my_state{vmcluster=\"$vmcluster\", namespace=~\"$namespace\",set=~\"$shard\"}/ mongodb_mongod_replset_my_state{vmcluster=\"$vmcluster\", namespace=~\"$namespace\",set=~\"$shard\"}))",
                    "format": "time_series",
                    "instant": false,
                    "interval": "$interval",
                    "legendFormat": "{{service_name}} ",
                    "refId": "A"
                }
            ],
            "timeFrom": null,
            "timeShift": null,
            "title": "QPS of Services in Shard - $shard",
            "type": "grafana-polystat-panel",
            "valueMaps": [
                {
                    "op": "=",
                    "text": "N/A",
                    "value": "null"
                }
            ]
        },
        {
            "collapsed": true,
            "datasource": null,
            "gridPos": {
                "h": 1,
                "w": 24,
                "x": 0,
                "y": 26
            },
            "id": 1199,
            "panels": [
                {
                    "backgroundColor": "rgba(128,128,128,0.1)",
                    "colorMaps": [
                        {
                            "color": "#CCC",
                            "text": "N/A"
                        }
                    ],
                    "crosshairColor": "#8F070C",
                    "datasource": null,
                    "description": "The MongoDB balancer is a background process that monitors the number of chunks on each shard. When the number of chunks on a given shard reaches specific migration thresholds, the balancer attempts to automatically migrate chunks between shards and reach an equal number of chunks per shard.",
                    "display": "timeline",
                    "expandFromQueryS": 0,
                    "extendLastValue": true,
                    "fieldConfig": {
                        "defaults": {},
                        "overrides": []
                    },
                    "gridPos": {
                        "h": 4,
                        "w": 24,
                        "x": 0,
                        "y": 27
                    },
                    "highlightOnMouseover": true,
                    "id": 1215,
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
                    "showLegend": true,
                    "showLegendNames": true,
                    "showLegendPercent": true,
                    "showLegendValues": true,
                    "showTimeAxis": true,
                    "targets": [
                        {
                            "expr": "min(mongodb_mongos_sharding_chunks_is_balanced{vmcluster=\"$vmcluster\", namespace=~\"$namespace\"})",
                            "interval": "$interval",
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
                    "title": "Chunks Balanced",
                    "type": "natel-discrete-panel",
                    "units": "short",
                    "useTimePrecision": false,
                    "valueMaps": [
                        {
                            "op": "=",
                            "text": "YES",
                            "value": "1"
                        },
                        {
                            "op": "=",
                            "text": "NO",
                            "value": "0"
                        }
                    ],
                    "valueTextColor": "#000000",
                    "writeAllValues": false,
                    "writeLastValue": true,
                    "writeMetricNames": false
                },
                {
                    "activePatternIndex": 0,
                    "datasource": null,
                    "debug_mode": false,
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
                        "row_name": "_series_",
                        "textColors": "red|orange|green",
                        "textColors_overrides": "0->red|2->green|1->yellow",
                        "thresholds": "70,90",
                        "time_based_thresholds": [],
                        "tooltipTemplate": "Series : _series_ <br/>Row Name : _row_name_ <br/>Col Name : _col_name_ <br/>Value : _value_",
                        "transform_values": "_value_|_value_|_value_",
                        "transform_values_overrides": "0->down|1->up",
                        "valueName": "current"
                    },
                    "default_title_for_rows": "Shard Name",
                    "description": "A chunk consists of a subset of sharded data.",
                    "fieldConfig": {
                        "defaults": {},
                        "overrides": []
                    },
                    "gridPos": {
                        "h": 8,
                        "w": 8,
                        "x": 0,
                        "y": 31
                    },
                    "hide_headers": false,
                    "id": 1200,
                    "patterns": [
                        {
                            "bgColors": "blue|blue|blue",
                            "bgColors_overrides": "0->green|2->red|1->yellow",
                            "clickable_cells_link": "/graph/d/mongodb-replicaset-summary/mongodb-replset-summary?var-interval=$__auto_interval_interval&var-replset=_0_",
                            "col_name": "_1_",
                            "decimals": "0",
                            "defaultBGColor": "",
                            "defaultTextColor": "",
                            "delimiter": "|",
                            "displayTemplate": "_value_",
                            "enable_bgColor": true,
                            "enable_bgColor_overrides": false,
                            "enable_clickable_cells": true,
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
                            "name": "Collection",
                            "null_color": "darkred",
                            "null_textcolor": "black",
                            "null_value": "No data",
                            "pattern": ".*Chunks",
                            "row_col_wrapper": "_",
                            "row_name": "_0_",
                            "textColors": "red|orange|green",
                            "textColors_overrides": "0->red|2->green|1->yellow",
                            "thresholds": "70,90",
                            "time_based_thresholds": [],
                            "tooltipTemplate": "-",
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
                            "exemplar": true,
                            "expr": "avg by (shard) (mongodb_mongos_sharding_shard_chunks_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\",db!~\"admin|config\"})",
                            "format": "time_series",
                            "hide": false,
                            "instant": true,
                            "interval": "$interval",
                            "legendFormat": "{{shard}} | Chunks",
                            "refId": "A"
                        }
                    ],
                    "text_alignment_header": "center",
                    "timeFrom": null,
                    "timeShift": null,
                    "title": "Amount of Chunks in Shards",
                    "type": "yesoreyeram-boomtable-panel"
                },
                {
                    "aliasColors": {},
                    "bars": false,
                    "dashLength": 10,
                    "dashes": false,
                    "datasource": null,
                    "decimals": 2,
                    "description": "The sharding operation creates the initial chunk(s) to cover the entire range of the shard key values. The number of chunks created depends on the configured chunk size. After the initial chunk creation, the balancer migrates these initial chunks across the shards as appropriate as well as manages the chunk distribution going forward.",
                    "fieldConfig": {
                        "defaults": {
                            "links": [
                                {
                                    "targetBlank": true,
                                    "title": "MongoDB ReplSet Summary",
                                    "url": "/graph/d/mongodb-replicaset-summary/mongodb-replset-summary?var-replset=${__series.name}&$__url_time_range"
                                }
                            ]
                        },
                        "overrides": []
                    },
                    "fill": 2,
                    "fillGradient": 0,
                    "gridPos": {
                        "h": 8,
                        "w": 16,
                        "x": 8,
                        "y": 31
                    },
                    "hiddenSeries": false,
                    "id": 1201,
                    "legend": {
                        "alignAsTable": true,
                        "avg": true,
                        "current": false,
                        "max": true,
                        "min": true,
                        "rightSide": true,
                        "show": true,
                        "sideWidth": null,
                        "sort": "avg",
                        "sortDesc": true,
                        "total": false,
                        "values": true
                    },
                    "lines": true,
                    "linewidth": 2,
                    "nullPointMode": "null",
                    "options": {
                        "alertThreshold": true
                    },
                    "percentage": false,
                    "pluginVersion": "7.5.6",
                    "pointradius": 2,
                    "points": false,
                    "renderer": "flot",
                    "seriesOverrides": [],
                    "spaceLength": 10,
                    "stack": false,
                    "steppedLine": false,
                    "targets": [
                        {
                            "exemplar": true,
                            "expr": "avg by (shard) (rate(mongodb_mongos_sharding_shard_chunks_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\"}[$interval]) or irate(mongodb_mongos_sharding_shard_chunks_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\"}[5m]))",
                            "interval": "$interval",
                            "legendFormat": "{{shard}}",
                            "refId": "A"
                        }
                    ],
                    "thresholds": [],
                    "timeFrom": null,
                    "timeRegions": [],
                    "timeShift": null,
                    "title": "Dynamic of Chunks",
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
                            "format": "cps",
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
                    "bars": true,
                    "dashLength": 10,
                    "dashes": false,
                    "datasource": null,
                    "decimals": 2,
                    "description": "MongoDB splits chunks when they grow beyond the configured chunk size. Both inserts and updates can trigger a chunk split.",
                    "fieldConfig": {
                        "defaults": {
                            "links": []
                        },
                        "overrides": []
                    },
                    "fill": 2,
                    "fillGradient": 0,
                    "gridPos": {
                        "h": 8,
                        "w": 12,
                        "x": 0,
                        "y": 39
                    },
                    "hiddenSeries": false,
                    "id": 1212,
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
                    "nullPointMode": "null",
                    "options": {
                        "alertThreshold": true
                    },
                    "percentage": false,
                    "pluginVersion": "7.5.6",
                    "pointradius": 2,
                    "points": false,
                    "renderer": "flot",
                    "seriesOverrides": [],
                    "spaceLength": 10,
                    "stack": false,
                    "steppedLine": false,
                    "targets": [
                        {
                            "exemplar": true,
                            "expr": "avg by (event) (rate(mongodb_mongos_sharding_changelog_10min_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", event=~\".*split.*\"}[$interval]) or\nirate(mongodb_mongos_sharding_changelog_10min_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", event=~\".*split.*\"}[5m]))",
                            "interval": "$interval",
                            "legendFormat": "{{event}}",
                            "refId": "A"
                        }
                    ],
                    "thresholds": [],
                    "timeFrom": null,
                    "timeRegions": [],
                    "timeShift": null,
                    "title": "Chunks Split Events",
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
                            "format": "short",
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
                    "bars": true,
                    "dashLength": 10,
                    "dashes": false,
                    "datasource": null,
                    "decimals": 2,
                    "description": "MongoDB migrates chunks in a sharded cluster to distribute the chunks of a sharded collection evenly among shards. ",
                    "fieldConfig": {
                        "defaults": {
                            "links": []
                        },
                        "overrides": []
                    },
                    "fill": 2,
                    "fillGradient": 0,
                    "gridPos": {
                        "h": 8,
                        "w": 12,
                        "x": 12,
                        "y": 39
                    },
                    "hiddenSeries": false,
                    "id": 1216,
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
                    "nullPointMode": "null",
                    "options": {
                        "alertThreshold": true
                    },
                    "percentage": false,
                    "pluginVersion": "7.5.6",
                    "pointradius": 2,
                    "points": false,
                    "renderer": "flot",
                    "seriesOverrides": [],
                    "spaceLength": 10,
                    "stack": false,
                    "steppedLine": false,
                    "targets": [
                        {
                            "exemplar": true,
                            "expr": "avg by (event) (rate(mongodb_mongos_sharding_changelog_10min_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", event=~\".*moveChunk.*\"}[$interval]) or\nirate(mongodb_mongos_sharding_changelog_10min_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", event=~\".*moveChunk.*\"}[5m]))",
                            "interval": "$interval",
                            "legendFormat": "{{event}}",
                            "refId": "A"
                        }
                    ],
                    "thresholds": [],
                    "timeFrom": null,
                    "timeRegions": [],
                    "timeShift": null,
                    "title": "Chunks Move Events",
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
                            "format": "short",
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
                }
            ],
            "title": "Chunks in Shards",
            "type": "row"
        },
        {
            "collapsed": true,
            "datasource": null,
            "gridPos": {
                "h": 1,
                "w": 24,
                "x": 0,
                "y": 27
            },
            "id": 1045,
            "panels": [
                {
                    "activePatternIndex": 0,
                    "datasource": null,
                    "debug_mode": false,
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
                        "row_name": "_series_",
                        "textColors": "red|orange|green",
                        "textColors_overrides": "0->red|2->green|1->yellow",
                        "thresholds": "70,90",
                        "time_based_thresholds": [],
                        "tooltipTemplate": "Series : _series_ <br/>Row Name : _row_name_ <br/>Col Name : _col_name_ <br/>Value : _value_",
                        "transform_values": "_value_|_value_|_value_",
                        "transform_values_overrides": "0->down|1->up",
                        "valueName": "current"
                    },
                    "default_title_for_rows": "DB Name",
                    "description": "Indexes are special data structures that store a small portion of the collection\u2019s data set in an easy to traverse form. ",
                    "fieldConfig": {
                        "defaults": {},
                        "overrides": []
                    },
                    "gridPos": {
                        "h": 8,
                        "w": 8,
                        "x": 0,
                        "y": 28
                    },
                    "hide_headers": false,
                    "id": 1040,
                    "patterns": [
                        {
                            "bgColors": "blue|blue|blue",
                            "bgColors_overrides": "0->green|2->red|1->yellow",
                            "clickable_cells_link": "/graph/d/mongodb-replicaset-summary/mongodb-replset-summary?var-interval=$__auto_interval_interval&var-replset=_1_",
                            "col_name": "_1_",
                            "decimals": "0",
                            "defaultBGColor": "",
                            "defaultTextColor": "",
                            "delimiter": "|",
                            "displayTemplate": "_value_",
                            "enable_bgColor": true,
                            "enable_bgColor_overrides": false,
                            "enable_clickable_cells": true,
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
                            "name": "Collection",
                            "null_color": "darkred",
                            "null_textcolor": "black",
                            "null_value": "No data",
                            "pattern": ".*Collections",
                            "row_col_wrapper": "_",
                            "row_name": "_0_",
                            "textColors": "red|orange|green",
                            "textColors_overrides": "0->red|2->green|1->yellow",
                            "thresholds": "70,90",
                            "time_based_thresholds": [],
                            "tooltipTemplate": "-",
                            "transform_values": "_value_|_value_|_value_",
                            "transform_values_overrides": "0->down|1->up",
                            "valueName": "current"
                        },
                        {
                            "bgColors": "green|orange|red",
                            "bgColors_overrides": "0->green|2->red|1->yellow",
                            "clickable_cells_link": "",
                            "col_name": "_2_",
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
                            "name": "Size",
                            "null_color": "darkred",
                            "null_textcolor": "black",
                            "null_value": "No data",
                            "pattern": ".*Size",
                            "row_col_wrapper": "_",
                            "row_name": "_1_",
                            "textColors": "red|orange|green",
                            "textColors_overrides": "0->red|2->green|1->yellow",
                            "thresholds": "70,90",
                            "time_based_thresholds": [],
                            "tooltipTemplate": "Series : _series_ <br/>Row Name : _row_name_ <br/>Col Name : _col_name_ <br/>Value : _value_",
                            "transform_values": "_value_|_value_|_value_",
                            "transform_values_overrides": "0->down|1->up",
                            "valueName": "avg"
                        }
                    ],
                    "row_col_wrapper": "_",
                    "sorting_props": {
                        "col_index": -1,
                        "direction": "desc"
                    },
                    "targets": [
                        {
                            "exemplar": true,
                            "expr": "avg by (db,shard) (mongodb_mongos_db_indexes_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\",db!~\"admin|config\"})",
                            "format": "time_series",
                            "hide": false,
                            "instant": true,
                            "interval": "$interval",
                            "legendFormat": "{{db}} | {{shard}} | Collections",
                            "refId": "A"
                        },
                        {
                            "expr": "mongodb_mongos_db_data_size_bytes{vmcluster=\"$vmcluster\", namespace=~\"$namespace\",db!~\"admin|config\"}",
                            "hide": true,
                            "instant": true,
                            "interval": "$intgerval",
                            "legendFormat": "{{db}} | {{shard}} | Size",
                            "refId": "B"
                        }
                    ],
                    "text_alignment_header": "center",
                    "timeFrom": null,
                    "timeShift": null,
                    "title": "Amount of Indexes in Shards",
                    "type": "yesoreyeram-boomtable-panel"
                },
                {
                    "aliasColors": {},
                    "bars": false,
                    "dashLength": 10,
                    "dashes": false,
                    "datasource": null,
                    "decimals": 2,
                    "description": "",
                    "fieldConfig": {
                        "defaults": {
                            "links": []
                        },
                        "overrides": []
                    },
                    "fill": 2,
                    "fillGradient": 0,
                    "gridPos": {
                        "h": 8,
                        "w": 16,
                        "x": 8,
                        "y": 28
                    },
                    "hiddenSeries": false,
                    "id": 1042,
                    "legend": {
                        "alignAsTable": true,
                        "avg": true,
                        "current": false,
                        "max": true,
                        "min": true,
                        "rightSide": true,
                        "show": true,
                        "sideWidth": null,
                        "sort": "avg",
                        "sortDesc": true,
                        "total": false,
                        "values": true
                    },
                    "lines": true,
                    "linewidth": 2,
                    "nullPointMode": "null",
                    "options": {
                        "alertThreshold": true
                    },
                    "percentage": false,
                    "pluginVersion": "7.5.6",
                    "pointradius": 2,
                    "points": false,
                    "renderer": "flot",
                    "seriesOverrides": [],
                    "spaceLength": 10,
                    "stack": false,
                    "steppedLine": false,
                    "targets": [
                        {
                            "exemplar": true,
                            "expr": "avg by (db,shard) (rate(mongodb_mongos_db_indexes_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\",db!~\"admin|config\"}[$interval]) or irate(mongodb_mongos_db_indexes_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\",db!~\"admin|config\"}[5m]))",
                            "interval": "$interval",
                            "legendFormat": "{{shard}}-{{db}}",
                            "refId": "A"
                        }
                    ],
                    "thresholds": [],
                    "timeFrom": null,
                    "timeRegions": [],
                    "timeShift": null,
                    "title": "Dynamic of Indexes",
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
                            "format": "cps",
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
                    "activePatternIndex": 0,
                    "datasource": null,
                    "debug_mode": false,
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
                        "row_name": "_series_",
                        "textColors": "red|orange|green",
                        "textColors_overrides": "0->red|2->green|1->yellow",
                        "thresholds": "70,90",
                        "time_based_thresholds": [],
                        "tooltipTemplate": "Series : _series_ <br/>Row Name : _row_name_ <br/>Col Name : _col_name_ <br/>Value : _value_",
                        "transform_values": "_value_|_value_|_value_",
                        "transform_values_overrides": "0->down|1->up",
                        "valueName": "current"
                    },
                    "default_title_for_rows": "DB Name",
                    "description": "The index stores the value of a specific field or set of fields, ordered by the value of the field. ",
                    "fieldConfig": {
                        "defaults": {},
                        "overrides": []
                    },
                    "gridPos": {
                        "h": 8,
                        "w": 8,
                        "x": 0,
                        "y": 36
                    },
                    "hide_headers": false,
                    "id": 1072,
                    "patterns": [
                        {
                            "bgColors": "blue|blue|blue",
                            "bgColors_overrides": "0->green|2->red|1->yellow",
                            "clickable_cells_link": "/graph/d/mongodb-replicaset-summary/mongodb-replset-summary?var-interval=$__auto_interval_interval&var-replset=_1_",
                            "col_name": "_1_",
                            "decimals": "0",
                            "defaultBGColor": "",
                            "defaultTextColor": "",
                            "delimiter": "|",
                            "displayTemplate": "_value_",
                            "enable_bgColor": true,
                            "enable_bgColor_overrides": false,
                            "enable_clickable_cells": true,
                            "enable_textColor": false,
                            "enable_textColor_overrides": false,
                            "enable_time_based_thresholds": false,
                            "enable_transform": false,
                            "enable_transform_overrides": false,
                            "filter": {
                                "value_above": "",
                                "value_below": ""
                            },
                            "format": "bytes",
                            "name": "Size",
                            "null_color": "darkred",
                            "null_textcolor": "black",
                            "null_value": "No data",
                            "pattern": ".*Collections",
                            "row_col_wrapper": "_",
                            "row_name": "_0_",
                            "textColors": "red|orange|green",
                            "textColors_overrides": "0->red|2->green|1->yellow",
                            "thresholds": "70,90",
                            "time_based_thresholds": [],
                            "tooltipTemplate": "-",
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
                            "exemplar": true,
                            "expr": "avg by (db,shard) (mongodb_mongos_db_index_size_bytes{vmcluster=\"$vmcluster\", namespace=~\"$namespace\",db!~\"admin|config\"})",
                            "format": "time_series",
                            "hide": false,
                            "instant": true,
                            "interval": "$interval",
                            "legendFormat": "{{db}} | {{shard}} | Collections",
                            "refId": "A"
                        }
                    ],
                    "text_alignment_header": "center",
                    "timeFrom": null,
                    "timeShift": null,
                    "title": "Size of Indexes in Shards",
                    "type": "yesoreyeram-boomtable-panel"
                },
                {
                    "aliasColors": {},
                    "bars": false,
                    "dashLength": 10,
                    "dashes": false,
                    "datasource": null,
                    "decimals": 2,
                    "description": "",
                    "fieldConfig": {
                        "defaults": {
                            "links": []
                        },
                        "overrides": []
                    },
                    "fill": 2,
                    "fillGradient": 0,
                    "gridPos": {
                        "h": 8,
                        "w": 16,
                        "x": 8,
                        "y": 36
                    },
                    "hiddenSeries": false,
                    "id": 1073,
                    "legend": {
                        "alignAsTable": true,
                        "avg": true,
                        "current": false,
                        "max": true,
                        "min": true,
                        "rightSide": true,
                        "show": true,
                        "sideWidth": null,
                        "sort": "avg",
                        "sortDesc": true,
                        "total": false,
                        "values": true
                    },
                    "lines": true,
                    "linewidth": 2,
                    "nullPointMode": "null",
                    "options": {
                        "alertThreshold": true
                    },
                    "percentage": false,
                    "pluginVersion": "7.5.6",
                    "pointradius": 2,
                    "points": false,
                    "renderer": "flot",
                    "seriesOverrides": [],
                    "spaceLength": 10,
                    "stack": false,
                    "steppedLine": false,
                    "targets": [
                        {
                            "exemplar": true,
                            "expr": "avg by (db,shard) (rate(mongodb_mongos_db_index_size_bytes{vmcluster=\"$vmcluster\", namespace=~\"$namespace\",db!~\"admin|config\"}[$interval]) or rate(mongodb_mongos_db_index_size_bytes{vmcluster=\"$vmcluster\", namespace=~\"$namespace\",db!~\"admin|config\"}[5m]))",
                            "interval": "$interval",
                            "legendFormat": "{{shard}}-{{db}}",
                            "refId": "A"
                        }
                    ],
                    "thresholds": [],
                    "timeFrom": null,
                    "timeRegions": [],
                    "timeShift": null,
                    "title": "Dynamic of Indexes Size",
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
                            "format": "Bps",
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
                }
            ],
            "title": "Indexes in Shards",
            "type": "row"
        },
        {
            "collapsed": false,
            "datasource": null,
            "gridPos": {
                "h": 1,
                "w": 24,
                "x": 0,
                "y": 29
            },
            "id": 81,
            "panels": [],
            "title": "Connections",
            "type": "row"
        },
        {
            "aliasColors": {},
            "bars": false,
            "dashLength": 10,
            "dashes": false,
            "datasource": null,
            "description": "TCP connections (Incoming) in mongod processes.",
            "editable": true,
            "error": false,
            "fieldConfig": {
                "defaults": {
                    "links": [
                        {
                            "targetBlank": true,
                            "title": "MongoDB ReplSet Summary - ${__series.name}",
                            "url": "/graph/d/mongodb-replicaset-summary/mongodb-replset-summary?var-replset=${__series.name}&$__url_time_range"
                        }
                    ]
                },
                "overrides": []
            },
            "fill": 6,
            "fillGradient": 0,
            "grid": {},
            "gridPos": {
                "h": 8,
                "w": 12,
                "x": 0,
                "y": 30
            },
            "hiddenSeries": false,
            "id": 37,
            "legend": {
                "alignAsTable": true,
                "avg": true,
                "current": false,
                "hideEmpty": false,
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
                "alertThreshold": true
            },
            "percentage": false,
            "pluginVersion": "7.5.6",
            "pointradius": 5,
            "points": false,
            "renderer": "flot",
            "seriesOverrides": [
                {
                    "alias": "Total",
                    "color": "#C4162A",
                    "fill": 0,
                    "legend": false,
                    "stack": false
                }
            ],
            "spaceLength": 10,
            "stack": true,
            "steppedLine": false,
            "targets": [
                {
                    "expr": "sum by (set) (avg by (service_name,set) (mongodb_connections{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", state=\"current\"}) * on (service_name) group_right avg by (service_name,set) (mongodb_mongod_replset_my_state{vmcluster=\"$vmcluster\", namespace=~\"$namespace\"}/ mongodb_mongod_replset_my_state{vmcluster=\"$vmcluster\", namespace=~\"$namespace\"}))",
                    "format": "time_series",
                    "hide": false,
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "{{set}}",
                    "refId": "B",
                    "step": 300
                },
                {
                    "expr": "sum by () (avg by (service_name,set) (mongodb_connections{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", state=\"current\"}) * on (service_name) group_right avg by (service_name,set) (mongodb_mongod_replset_my_state{vmcluster=\"$vmcluster\", namespace=~\"$namespace\"}/ mongodb_mongod_replset_my_state{vmcluster=\"$vmcluster\", namespace=~\"$namespace\"}))",
                    "hide": false,
                    "interval": "$interval",
                    "legendFormat": "Total",
                    "refId": "A"
                },
                {
                    "expr": "sum by (shard) (avg by (service_name,shard) (mongodb_connections{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", state=\"current\"}) * on (service_name) group_right avg by (service_name,shard) (mongodb_mongos_db_collections_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\"} / mongodb_mongos_db_collections_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\"}))",
                    "format": "time_series",
                    "hide": true,
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "{{shard}}",
                    "refId": "C",
                    "step": 300
                }
            ],
            "thresholds": [],
            "timeFrom": null,
            "timeRegions": [],
            "timeShift": null,
            "title": "Current Connections Per Shard",
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
            "description": "Incoming connections to mongos nodes.",
            "editable": true,
            "error": false,
            "fieldConfig": {
                "defaults": {
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
                "y": 30
            },
            "hiddenSeries": false,
            "id": 7,
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
            "percentage": false,
            "pluginVersion": "7.5.6",
            "pointradius": 5,
            "points": false,
            "renderer": "flot",
            "seriesOverrides": [],
            "spaceLength": 10,
            "stack": false,
            "steppedLine": false,
            "targets": [
                {
                    "expr": "sum by (state) (max_over_time(mongodb_connections{vmcluster=\"$vmcluster\", namespace=~\"$namespace\"}[$interval]) or max_over_time(mongodb_connections{vmcluster=\"$vmcluster\", namespace=~\"$namespace\"}[5m]))",
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
            "title": "Total Connections",
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
                    "format": "short",
                    "label": "",
                    "logBase": 2,
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
            "collapsed": true,
            "datasource": null,
            "gridPos": {
                "h": 1,
                "w": 24,
                "x": 0,
                "y": 38
            },
            "id": 50,
            "panels": [
                {
                    "aliasColors": {},
                    "bars": false,
                    "dashLength": 10,
                    "dashes": false,
                    "datasource": null,
                    "decimals": 2,
                    "editable": true,
                    "error": false,
                    "fill": 6,
                    "fillGradient": 0,
                    "grid": {},
                    "gridPos": {
                        "h": 7,
                        "w": 24,
                        "x": 0,
                        "y": 39
                    },
                    "hiddenSeries": false,
                    "id": 1178,
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
                    "maxPerRow": 2,
                    "nullPointMode": "null",
                    "options": {
                        "dataLinks": []
                    },
                    "percentage": false,
                    "pointradius": 5,
                    "points": false,
                    "renderer": "flot",
                    "repeatDirection": "h",
                    "seriesOverrides": [],
                    "spaceLength": 10,
                    "stack": true,
                    "steppedLine": false,
                    "targets": [
                        {
                            "expr": "sum by (service_name) (max_over_time(mongodb_connections{vmcluster=\"$vmcluster\", namespace=~\"$namespace\",service_name=~\"$service_name\",state=\"current\"}[$interval]) or \nmax_over_time(mongodb_connections{vmcluster=\"$vmcluster\", namespace=~\"$namespace\",service_name=~\"$service_name\",state=\"current\"}[5m]))",
                            "format": "time_series",
                            "hide": false,
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}}",
                            "refId": "J",
                            "step": 300
                        }
                    ],
                    "thresholds": [],
                    "timeFrom": null,
                    "timeRegions": [],
                    "timeShift": null,
                    "title": "Current Connections ",
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
                    "description": "",
                    "editable": true,
                    "error": false,
                    "fill": 6,
                    "fillGradient": 0,
                    "grid": {},
                    "gridPos": {
                        "h": 7,
                        "w": 24,
                        "x": 0,
                        "y": 46
                    },
                    "hiddenSeries": false,
                    "id": 1177,
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
                    "maxPerRow": 2,
                    "nullPointMode": "null",
                    "options": {
                        "dataLinks": []
                    },
                    "percentage": false,
                    "pointradius": 5,
                    "points": false,
                    "renderer": "flot",
                    "repeatDirection": "h",
                    "seriesOverrides": [],
                    "spaceLength": 10,
                    "stack": true,
                    "steppedLine": false,
                    "targets": [
                        {
                            "expr": "sum by (service_name) (max_over_time(mongodb_connections{vmcluster=\"$vmcluster\", namespace=~\"$namespace\",service_name=~\"$service_name\",state=\"available\"}[$interval]) or \nmax_over_time(mongodb_connections{vmcluster=\"$vmcluster\", namespace=~\"$namespace\",service_name=~\"$service_name\",state=\"available\"}[5m]))",
                            "format": "time_series",
                            "hide": false,
                            "interval": "$interval",
                            "intervalFactor": 1,
                            "legendFormat": "{{service_name}}",
                            "refId": "J",
                            "step": 300
                        }
                    ],
                    "thresholds": [],
                    "timeFrom": null,
                    "timeRegions": [],
                    "timeShift": null,
                    "title": "Available Connections",
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
            "title": "Connections Details",
            "type": "row"
        },
        {
            "collapsed": false,
            "datasource": null,
            "gridPos": {
                "h": 1,
                "w": 24,
                "x": 0,
                "y": 39
            },
            "id": 1149,
            "panels": [],
            "title": "Cursors",
            "type": "row"
        },
        {
            "aliasColors": {},
            "bars": false,
            "dashLength": 10,
            "dashes": false,
            "datasource": null,
            "decimals": 0,
            "description": "The Cursor is a MongoDB Collection of the document which is returned upon the find method execution. ",
            "editable": true,
            "error": false,
            "fieldConfig": {
                "defaults": {
                    "links": [
                        {
                            "targetBlank": true,
                            "title": "MongoDB ReplSet Summary - ${__series.name}",
                            "url": "/graph/d/mongodb-replicaset-summary/mongodb-replset-summary?var-replset=${__series.name}&$__url_time_range"
                        }
                    ]
                },
                "overrides": []
            },
            "fill": 6,
            "fillGradient": 0,
            "grid": {},
            "gridPos": {
                "h": 8,
                "w": 12,
                "x": 0,
                "y": 40
            },
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
            "options": {
                "alertThreshold": true
            },
            "percentage": false,
            "pluginVersion": "7.5.6",
            "pointradius": 5,
            "points": false,
            "renderer": "flot",
            "seriesOverrides": [
                {
                    "alias": "Total",
                    "color": "#C4162A",
                    "fill": 0,
                    "legend": false,
                    "stack": false
                }
            ],
            "spaceLength": 10,
            "stack": true,
            "steppedLine": false,
            "targets": [
                {
                    "expr": "sum(sum(mongodb_mongod_metrics_cursor_open{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", state=\"total\"} or mongodb_mongod_cursors{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", state=\"total_open\"}) by (service_name) * on (service_name) group_right mongodb_mongod_replset_my_state{vmcluster=\"$vmcluster\", namespace=~\"$namespace\"} / mongodb_mongod_replset_my_state{vmcluster=\"$vmcluster\", namespace=~\"$namespace\"}) by (set)",
                    "format": "time_series",
                    "hide": false,
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "{{set}}",
                    "refId": "A",
                    "step": 300
                },
                {
                    "expr": "sum(sum(mongodb_mongod_metrics_cursor_open{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", state=\"total\"} or mongodb_mongod_cursors{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", state=\"total_open\"}) by (service_name) * on (service_name) group_right mongodb_mongod_replset_my_state{vmcluster=\"$vmcluster\", namespace=~\"$namespace\"} / mongodb_mongod_replset_my_state{vmcluster=\"$vmcluster\", namespace=~\"$namespace\"})",
                    "interval": "$interval",
                    "legendFormat": "Total",
                    "refId": "B"
                }
            ],
            "thresholds": [],
            "timeFrom": null,
            "timeRegions": [],
            "timeShift": null,
            "title": "Cursors Per Shard",
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
            "description": "The Cursor is a MongoDB Collection of the document which is returned upon the find method execution.",
            "editable": true,
            "error": false,
            "fieldConfig": {
                "defaults": {
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
                "y": 40
            },
            "hiddenSeries": false,
            "id": 31,
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
            "percentage": false,
            "pluginVersion": "7.5.6",
            "pointradius": 5,
            "points": false,
            "renderer": "flot",
            "seriesOverrides": [],
            "spaceLength": 10,
            "stack": false,
            "steppedLine": false,
            "targets": [
                {
                    "exemplar": true,
                    "expr": "sum by () (\nmax_over_time(mongodb_mongod_metrics_cursor_open{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", state=\"total\"}[$interval]) or\nmax_over_time(mongodb_mongod_metrics_cursor_open{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", state=\"total\"}[5m]) or \nmax_over_time(mongodb_mongos_metrics_cursor_open{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", state=\"total\"}[$interval]) or\nmax_over_time(mongodb_mongos_metrics_cursor_open{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", state=\"total\"}[5m]) or\nmax_over_time(mongodb_mongod_cursors{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", state=\"total_open\"}[$interval]) or\nmax_over_time(mongodb_mongod_cursors{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", state=\"total_open\"}[5m]) or\nmax_over_time(mongodb_mongos_cursors{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", state=\"total_open\"}[$interval]) or\nmax_over_time(mongodb_mongos_cursors{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", state=\"total_open\"}[5m])\n)",
                    "hide": false,
                    "interval": "$interval",
                    "legendFormat": "Coursors",
                    "refId": "A"
                }
            ],
            "thresholds": [],
            "timeFrom": null,
            "timeRegions": [],
            "timeShift": null,
            "title": "Mongos Cursors",
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
            "collapsed": true,
            "datasource": null,
            "gridPos": {
                "h": 1,
                "w": 24,
                "x": 0,
                "y": 48
            },
            "id": 62,
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
                    "fill": 6,
                    "fillGradient": 0,
                    "grid": {},
                    "gridPos": {
                        "h": 8,
                        "w": 24,
                        "x": 0,
                        "y": 49
                    },
                    "hiddenSeries": false,
                    "id": 1179,
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
                    "nullPointMode": "null",
                    "options": {
                        "dataLinks": []
                    },
                    "percentage": false,
                    "pointradius": 5,
                    "points": false,
                    "renderer": "flot",
                    "seriesOverrides": [
                        {
                            "alias": "Total",
                            "color": "#C4162A",
                            "fill": 0,
                            "legend": false,
                            "stack": false
                        }
                    ],
                    "spaceLength": 10,
                    "stack": true,
                    "steppedLine": false,
                    "targets": [
                        {
                            "expr": "sum by (service_name) (\nmax_over_time(mongodb_mongod_metrics_cursor_open{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", state=\"total\"}[$interval]) or\nmax_over_time(mongodb_mongod_metrics_cursor_open{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", state=\"total\"}[5m]) or max_over_time(mongodb_mongos_metrics_cursor_open{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", state=\"total\"}[$interval]) or\nmax_over_time(mongodb_mongos_metrics_cursor_open{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", state=\"total\"}[5m]) or\nmax_over_time(mongodb_mongod_cursors{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", state=\"total_open\"}[$interval]) or\nmax_over_time(mongodb_mongod_cursors{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", state=\"total_open\"}[5m]) or\nmax_over_time(mongodb_mongos_cursors{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", state=\"total_open\"}[$interval]) or\nmax_over_time(mongodb_mongos_cursors{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", state=\"total_open\"}[5m])\n)",
                            "hide": false,
                            "interval": "$interval",
                            "legendFormat": "{{service_name}}",
                            "refId": "A"
                        },
                        {
                            "expr": "sum by () (\nmax_over_time(mongodb_mongod_metrics_cursor_open{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", state=\"total\"}[$interval]) or\nmax_over_time(mongodb_mongod_metrics_cursor_open{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", state=\"total\"}[5m]) or max_over_time(mongodb_mongos_metrics_cursor_open{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", state=\"total\"}[$interval]) or\nmax_over_time(mongodb_mongos_metrics_cursor_open{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", state=\"total\"}[5m]) or\nmax_over_time(mongodb_mongod_cursors{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", state=\"total_open\"}[$interval]) or\nmax_over_time(mongodb_mongod_cursors{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", state=\"total_open\"}[5m]) or\nmax_over_time(mongodb_mongos_cursors{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", state=\"total_open\"}[$interval]) or\nmax_over_time(mongodb_mongos_cursors{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", state=\"total_open\"}[5m])\n)",
                            "interval": "$interval",
                            "legendFormat": "Total",
                            "refId": "B"
                        }
                    ],
                    "thresholds": [],
                    "timeFrom": null,
                    "timeRegions": [],
                    "timeShift": null,
                    "title": "Mongos Cursors",
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
            "title": "Cursors Details",
            "type": "row"
        },
        {
            "datasource": null,
            "gridPos": {
                "h": 1,
                "w": 24,
                "x": 0,
                "y": 49
            },
            "id": 1056,
            "title": "Mongos Operations",
            "type": "row"
        },
        {
            "aliasColors": {},
            "bars": false,
            "dashLength": 10,
            "dashes": false,
            "datasource": null,
            "decimals": 0,
            "description": "Ops/sec, classified by legacy wire protocol type (query, insert, update, delete, getmore).",
            "editable": true,
            "error": false,
            "fieldConfig": {
                "defaults": {
                    "links": [
                        {
                            "targetBlank": true,
                            "title": "MongoDB ReplSet Summary - ${__series.name}",
                            "url": "/graph/d/mongodb-replicaset-summary/mongodb-replset-summary?var-replset=${__series.name}&$__url_time_range"
                        }
                    ]
                },
                "overrides": []
            },
            "fill": 6,
            "fillGradient": 0,
            "grid": {},
            "gridPos": {
                "h": 8,
                "w": 12,
                "x": 0,
                "y": 50
            },
            "hiddenSeries": false,
            "id": 30,
            "legend": {
                "alignAsTable": true,
                "avg": true,
                "current": false,
                "hideEmpty": false,
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
                "alertThreshold": true
            },
            "percentage": false,
            "pluginVersion": "7.5.6",
            "pointradius": 5,
            "points": false,
            "renderer": "flot",
            "seriesOverrides": [
                {}
            ],
            "spaceLength": 10,
            "stack": true,
            "steppedLine": false,
            "targets": [
                {
                    "expr": "sum(sum(rate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", type!=\"command\"}[$interval]) or \nirate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", type!=\"command\"}[5m])) by (instance) * on (instance) \ngroup_right mongodb_mongod_replset_my_state{vmcluster=\"$vmcluster\", namespace=~\"$namespace\"} / mongodb_mongod_replset_my_state{vmcluster=\"$vmcluster\", namespace=~\"$namespace\"}) by (set)",
                    "format": "time_series",
                    "hide": false,
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "{{set}}",
                    "refId": "A",
                    "step": 300
                }
            ],
            "thresholds": [],
            "timeFrom": null,
            "timeRegions": [],
            "timeShift": null,
            "title": "Operations Per Shard",
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
            "description": "Ops/sec, classified by legacy wire protocol type (query, insert, update, delete, getmore).",
            "editable": true,
            "error": false,
            "fieldConfig": {
                "defaults": {
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
                "y": 50
            },
            "hiddenSeries": false,
            "id": 46,
            "interval": "",
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
                "alertThreshold": true
            },
            "percentage": false,
            "pluginVersion": "7.5.6",
            "pointradius": 5,
            "points": false,
            "renderer": "flot",
            "seriesOverrides": [],
            "spaceLength": 10,
            "stack": false,
            "steppedLine": false,
            "targets": [
                {
                    "expr": "sum by (type) (rate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", type!=\"command\"}[$interval]) or \nirate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", type!=\"command\"}[5m]))",
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
            "title": "Total Mongos Operations",
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
            "id": 52,
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
                    "grid": {},
                    "gridPos": {
                        "h": 7,
                        "w": 12,
                        "x": 0,
                        "y": 59
                    },
                    "hiddenSeries": false,
                    "id": 53,
                    "interval": "",
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
                    "maxPerRow": 2,
                    "nullPointMode": "null",
                    "options": {
                        "dataLinks": []
                    },
                    "percentage": false,
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
                            "expr": "sum by (type) (rate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", type!=\"command\",service_name=~\"$service_name\"}[$interval]) or \nirate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", type!=\"command\",service_name=~\"$service_name\"}[5m]))",
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
                    "title": "Operations $service_name",
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
                    "description": "",
                    "editable": true,
                    "error": false,
                    "fill": 2,
                    "fillGradient": 0,
                    "grid": {},
                    "gridPos": {
                        "h": 7,
                        "w": 12,
                        "x": 12,
                        "y": 59
                    },
                    "hiddenSeries": false,
                    "id": 1243,
                    "interval": "",
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
                    "maxPerRow": 2,
                    "nullPointMode": "null",
                    "options": {
                        "dataLinks": []
                    },
                    "percentage": false,
                    "pointradius": 5,
                    "points": false,
                    "renderer": "flot",
                    "repeat": null,
                    "repeatDirection": "h",
                    "repeatIteration": 1592916617858,
                    "repeatPanelId": 53,
                    "seriesOverrides": [],
                    "spaceLength": 10,
                    "stack": false,
                    "steppedLine": false,
                    "targets": [
                        {
                            "expr": "sum by (type) (rate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", type!=\"command\",service_name=~\"$service_name\"}[$interval]) or \nirate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", type!=\"command\",service_name=~\"$service_name\"}[5m]))",
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
                    "title": "Operations $service_name",
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
                    "description": "",
                    "editable": true,
                    "error": false,
                    "fill": 2,
                    "fillGradient": 0,
                    "grid": {},
                    "gridPos": {
                        "h": 7,
                        "w": 12,
                        "x": 0,
                        "y": 66
                    },
                    "hiddenSeries": false,
                    "id": 1244,
                    "interval": "",
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
                    "maxPerRow": 2,
                    "nullPointMode": "null",
                    "options": {
                        "dataLinks": []
                    },
                    "percentage": false,
                    "pointradius": 5,
                    "points": false,
                    "renderer": "flot",
                    "repeat": null,
                    "repeatDirection": "h",
                    "repeatIteration": 1592916617858,
                    "repeatPanelId": 53,
                    "seriesOverrides": [],
                    "spaceLength": 10,
                    "stack": false,
                    "steppedLine": false,
                    "targets": [
                        {
                            "expr": "sum by (type) (rate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", type!=\"command\",service_name=~\"$service_name\"}[$interval]) or \nirate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", type!=\"command\",service_name=~\"$service_name\"}[5m]))",
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
                    "title": "Operations $service_name",
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
                    "description": "",
                    "editable": true,
                    "error": false,
                    "fill": 2,
                    "fillGradient": 0,
                    "grid": {},
                    "gridPos": {
                        "h": 7,
                        "w": 12,
                        "x": 12,
                        "y": 66
                    },
                    "hiddenSeries": false,
                    "id": 1245,
                    "interval": "",
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
                    "maxPerRow": 2,
                    "nullPointMode": "null",
                    "options": {
                        "dataLinks": []
                    },
                    "percentage": false,
                    "pointradius": 5,
                    "points": false,
                    "renderer": "flot",
                    "repeat": null,
                    "repeatDirection": "h",
                    "repeatIteration": 1592916617858,
                    "repeatPanelId": 53,
                    "seriesOverrides": [],
                    "spaceLength": 10,
                    "stack": false,
                    "steppedLine": false,
                    "targets": [
                        {
                            "expr": "sum by (type) (rate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", type!=\"command\",service_name=~\"$service_name\"}[$interval]) or \nirate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", type!=\"command\",service_name=~\"$service_name\"}[5m]))",
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
                    "title": "Operations $service_name",
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
                    "description": "",
                    "editable": true,
                    "error": false,
                    "fill": 2,
                    "fillGradient": 0,
                    "grid": {},
                    "gridPos": {
                        "h": 7,
                        "w": 12,
                        "x": 0,
                        "y": 73
                    },
                    "hiddenSeries": false,
                    "id": 1246,
                    "interval": "",
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
                    "maxPerRow": 2,
                    "nullPointMode": "null",
                    "options": {
                        "dataLinks": []
                    },
                    "percentage": false,
                    "pointradius": 5,
                    "points": false,
                    "renderer": "flot",
                    "repeat": null,
                    "repeatDirection": "h",
                    "repeatIteration": 1592916617858,
                    "repeatPanelId": 53,
                    "seriesOverrides": [],
                    "spaceLength": 10,
                    "stack": false,
                    "steppedLine": false,
                    "targets": [
                        {
                            "expr": "sum by (type) (rate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", type!=\"command\",service_name=~\"$service_name\"}[$interval]) or \nirate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", type!=\"command\",service_name=~\"$service_name\"}[5m]))",
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
                    "title": "Operations $service_name",
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
                    "description": "",
                    "editable": true,
                    "error": false,
                    "fill": 2,
                    "fillGradient": 0,
                    "grid": {},
                    "gridPos": {
                        "h": 7,
                        "w": 12,
                        "x": 12,
                        "y": 73
                    },
                    "hiddenSeries": false,
                    "id": 1247,
                    "interval": "",
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
                    "maxPerRow": 2,
                    "nullPointMode": "null",
                    "options": {
                        "dataLinks": []
                    },
                    "percentage": false,
                    "pointradius": 5,
                    "points": false,
                    "renderer": "flot",
                    "repeat": null,
                    "repeatDirection": "h",
                    "repeatIteration": 1592916617858,
                    "repeatPanelId": 53,
                    "seriesOverrides": [],
                    "spaceLength": 10,
                    "stack": false,
                    "steppedLine": false,
                    "targets": [
                        {
                            "expr": "sum by (type) (rate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", type!=\"command\",service_name=~\"$service_name\"}[$interval]) or \nirate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", type!=\"command\",service_name=~\"$service_name\"}[5m]))",
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
                    "title": "Operations $service_name",
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
                    "description": "",
                    "editable": true,
                    "error": false,
                    "fill": 2,
                    "fillGradient": 0,
                    "grid": {},
                    "gridPos": {
                        "h": 7,
                        "w": 12,
                        "x": 0,
                        "y": 80
                    },
                    "hiddenSeries": false,
                    "id": 1248,
                    "interval": "",
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
                    "maxPerRow": 2,
                    "nullPointMode": "null",
                    "options": {
                        "dataLinks": []
                    },
                    "percentage": false,
                    "pointradius": 5,
                    "points": false,
                    "renderer": "flot",
                    "repeat": null,
                    "repeatDirection": "h",
                    "repeatIteration": 1592916617858,
                    "repeatPanelId": 53,
                    "seriesOverrides": [],
                    "spaceLength": 10,
                    "stack": false,
                    "steppedLine": false,
                    "targets": [
                        {
                            "expr": "sum by (type) (rate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", type!=\"command\",service_name=~\"$service_name\"}[$interval]) or \nirate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", type!=\"command\",service_name=~\"$service_name\"}[5m]))",
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
                    "title": "Operations $service_name",
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
                    "description": "",
                    "editable": true,
                    "error": false,
                    "fill": 2,
                    "fillGradient": 0,
                    "grid": {},
                    "gridPos": {
                        "h": 7,
                        "w": 12,
                        "x": 12,
                        "y": 80
                    },
                    "hiddenSeries": false,
                    "id": 1249,
                    "interval": "",
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
                    "maxPerRow": 2,
                    "nullPointMode": "null",
                    "options": {
                        "dataLinks": []
                    },
                    "percentage": false,
                    "pointradius": 5,
                    "points": false,
                    "renderer": "flot",
                    "repeat": null,
                    "repeatDirection": "h",
                    "repeatIteration": 1592916617858,
                    "repeatPanelId": 53,
                    "seriesOverrides": [],
                    "spaceLength": 10,
                    "stack": false,
                    "steppedLine": false,
                    "targets": [
                        {
                            "expr": "sum by (type) (rate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", type!=\"command\",service_name=~\"$service_name\"}[$interval]) or \nirate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", type!=\"command\",service_name=~\"$service_name\"}[5m]))",
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
                    "title": "Operations $service_name",
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
                    "description": "",
                    "editable": true,
                    "error": false,
                    "fill": 2,
                    "fillGradient": 0,
                    "grid": {},
                    "gridPos": {
                        "h": 7,
                        "w": 12,
                        "x": 0,
                        "y": 87
                    },
                    "hiddenSeries": false,
                    "id": 1250,
                    "interval": "",
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
                    "maxPerRow": 2,
                    "nullPointMode": "null",
                    "options": {
                        "dataLinks": []
                    },
                    "percentage": false,
                    "pointradius": 5,
                    "points": false,
                    "renderer": "flot",
                    "repeat": null,
                    "repeatDirection": "h",
                    "repeatIteration": 1592916617858,
                    "repeatPanelId": 53,
                    "seriesOverrides": [],
                    "spaceLength": 10,
                    "stack": false,
                    "steppedLine": false,
                    "targets": [
                        {
                            "expr": "sum by (type) (rate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", type!=\"command\",service_name=~\"$service_name\"}[$interval]) or \nirate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", type!=\"command\",service_name=~\"$service_name\"}[5m]))",
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
                    "title": "Operations $service_name",
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
                    "description": "",
                    "editable": true,
                    "error": false,
                    "fill": 2,
                    "fillGradient": 0,
                    "grid": {},
                    "gridPos": {
                        "h": 7,
                        "w": 12,
                        "x": 12,
                        "y": 87
                    },
                    "hiddenSeries": false,
                    "id": 1251,
                    "interval": "",
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
                    "maxPerRow": 2,
                    "nullPointMode": "null",
                    "options": {
                        "dataLinks": []
                    },
                    "percentage": false,
                    "pointradius": 5,
                    "points": false,
                    "renderer": "flot",
                    "repeat": null,
                    "repeatDirection": "h",
                    "repeatIteration": 1592916617858,
                    "repeatPanelId": 53,
                    "seriesOverrides": [],
                    "spaceLength": 10,
                    "stack": false,
                    "steppedLine": false,
                    "targets": [
                        {
                            "expr": "sum by (type) (rate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", type!=\"command\",service_name=~\"$service_name\"}[$interval]) or \nirate(mongodb_op_counters_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", type!=\"command\",service_name=~\"$service_name\"}[5m]))",
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
                    "title": "Operations $service_name",
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
                            "show": true
                        }
                    ],
                    "yaxis": {
                        "align": false,
                        "alignLevel": null
                    }
                }
            ],
            "title": "Operations Details",
            "type": "row"
        },
        {
            "collapsed": false,
            "datasource": null,
            "gridPos": {
                "h": 1,
                "w": 24,
                "x": 0,
                "y": 59
            },
            "id": 97,
            "panels": [],
            "title": "Other",
            "type": "row"
        },
        {
            "aliasColors": {},
            "bars": false,
            "dashLength": 10,
            "dashes": false,
            "datasource": null,
            "editable": true,
            "error": false,
            "fieldConfig": {
                "defaults": {
                    "links": [
                        {
                            "targetBlank": true,
                            "title": "MongoDB ReplSet Summary - ${__series.name}",
                            "url": "/graph/d/mongodb-replicaset-summary/mongodb-replset-summary?var-replset=${__series.name}&$__url_time_range"
                        }
                    ]
                },
                "overrides": []
            },
            "fill": 2,
            "fillGradient": 0,
            "grid": {},
            "gridPos": {
                "h": 8,
                "w": 12,
                "x": 0,
                "y": 60
            },
            "hiddenSeries": false,
            "id": 14,
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
                "alertThreshold": true
            },
            "percentage": false,
            "pluginVersion": "7.5.6",
            "pointradius": 5,
            "points": false,
            "renderer": "flot",
            "seriesOverrides": [],
            "spaceLength": 10,
            "stack": false,
            "steppedLine": false,
            "targets": [
                {
                    "expr": "max by (set) (max(max_over_time(mongodb_mongod_replset_member_replication_lag{vmcluster=\"$vmcluster\", namespace=~\"$namespace\"}[$interval]) > 0) by (service_name,set) or max(max_over_time(mongodb_mongod_replset_member_replication_lag{vmcluster=\"$vmcluster\", namespace=~\"$namespace\"}[5m]) > 0) by (service_name,set))",
                    "format": "time_series",
                    "hide": false,
                    "interval": "$interval",
                    "intervalFactor": 1,
                    "legendFormat": "{{set}}",
                    "refId": "B",
                    "step": 300
                }
            ],
            "thresholds": [],
            "timeFrom": null,
            "timeRegions": [],
            "timeShift": null,
            "title": "Replication Lag by Set",
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
            "bars": true,
            "dashLength": 10,
            "dashes": false,
            "datasource": null,
            "decimals": 2,
            "description": "Count, over last 10 minutes, of all types of config db changelog events.",
            "fieldConfig": {
                "defaults": {
                    "links": []
                },
                "overrides": []
            },
            "fill": 2,
            "fillGradient": 0,
            "gridPos": {
                "h": 8,
                "w": 12,
                "x": 12,
                "y": 60
            },
            "hiddenSeries": false,
            "id": 1213,
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
            "nullPointMode": "null",
            "options": {
                "alertThreshold": true
            },
            "percentage": false,
            "pluginVersion": "7.5.6",
            "pointradius": 2,
            "points": false,
            "renderer": "flot",
            "seriesOverrides": [],
            "spaceLength": 10,
            "stack": false,
            "steppedLine": false,
            "targets": [
                {
                    "expr": "rate(mongodb_mongos_sharding_changelog_10min_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", event=~\".*(shard|Shard).*\"}[$interval]) or\nirate(mongodb_mongos_sharding_changelog_10min_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\", event=~\".*(shard|Shard).*\"}[5m])",
                    "interval": "$interval",
                    "legendFormat": "{{event}}",
                    "refId": "A"
                }
            ],
            "thresholds": [],
            "timeFrom": null,
            "timeRegions": [],
            "timeShift": null,
            "title": "Change Log Events",
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
                    "format": "short",
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
        }
    ],
    "refresh": "1m",
    "schemaVersion": 27,
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
                "allValue": "",
                "current": {
                    "selected": true,
                    "text": "All",
                    "value": "$__all"
                },
                "datasource": "prometheus",
                "definition": "",
                "hide": 0,
                "allValue": ".*",
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
                "allValue": null,
                "current": {},
                "datasource": "prometheus",
                "definition": "label_values(mongodb_up{vmcluster=\"$vmcluster\", namespace=~\"$namespace\"},service_name)",
                "description": null,
                "error": null,
                "hide": 2,
                "includeAll": true,
                "label": "Service Name",
                "multi": true,
                "name": "service_name",
                "options": [],
                "query": {
                    "query": "label_values(mongodb_up{vmcluster=\"$vmcluster\", namespace=~\"$namespace\"},service_name)",
                    "refId": "Metrics-service_name-Variable-Query"
                },
                "refresh": 2,
                "regex": "",
                "skipUrlSync": false,
                "sort": 5,
                "tagValuesQuery": "",
                "tags": [],
                "tagsQuery": "",
                "type": "query",
                "useTags": false
            },
            {
                "allValue": null,
                "current": {},
                "datasource": "prometheus",
                "definition": "label_values(mongodb_mongos_db_collections_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\"},shard)",
                "description": null,
                "error": null,
                "hide": 2,
                "includeAll": true,
                "label": "Shard Name",
                "multi": true,
                "name": "shard",
                "options": [],
                "query": {
                    "query": "label_values(mongodb_mongos_db_collections_total{vmcluster=\"$vmcluster\", namespace=~\"$namespace\"},shard)",
                    "refId": "Metrics-shard-Variable-Query"
                },
                "refresh": 2,
                "regex": "",
                "skipUrlSync": false,
                "sort": 5,
                "tagValuesQuery": "",
                "tags": [],
                "tagsQuery": "",
                "type": "query",
                "useTags": false
            },
            {
                "allValue": null,
                "current": {},
                "description": null,
                "error": null,
                "hide": 2,
                "includeAll": true,
                "label": "Schema",
                "multi": true,
                "name": "schema",
                "options": [],
                "query": "",
                "skipUrlSync": false,
                "type": "custom"
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
    "title": "MongoDB Cluster Summary",
    "uid": "mongodb-cluster-summary",
    "version": 1
}
`
