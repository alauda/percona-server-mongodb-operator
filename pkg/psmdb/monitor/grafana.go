package monitor

import (
	gdv1beta1 "gomod.alauda.cn/ait-apis/grafanadashboard/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const MongoGrafanaJson = `
{
  "annotations": {
    "list": [
      {
        "builtIn": 1,
        "datasource": "-- Grafana --",
        "enable": true,
        "hide": true,
        "iconColor": "rgba(0, 211, 255, 1)",
        "name": "Annotations & Alerts",
        "type": "dashboard"
      }
    ]
  },
  "editable": true,
  "gnetId": null,
  "graphTooltip": 1,
  "id": 58,
  "iteration": 1611230392235,
  "links": [],
  "panels": [
    {
      "aliasColors": {},
      "bars": false,
      "dashLength": 10,
      "dashes": false,
      "datasource": "prometheus",
      "decimals": 1,
      "description": "Shows how many times a command is executed per second on average during the selected interval.",
      "editable": true,
      "error": false,
      "fill": 2,
      "grid": {
        "leftLogBase": 1,
        "leftMax": null,
        "leftMin": 0,
        "rightLogBase": 1,
        "rightMax": null,
        "rightMin": null
      },
      "gridPos": {
        "h": 7,
        "w": 24,
        "x": 0,
        "y": 0
      },
      "id": 15,
      "legend": {
        "alignAsTable": true,
        "avg": true,
        "current": false,
        "hideZero": true,
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
          "expr": "rate(mongodb_ss_opcounters{instance=~\"$host\", legacy_op_type!=\"command\"}[$interval]) or \nirate(mongodb_ss_opcounters{instance=~\"$host\", legacy_op_type!=\"command\"}[5m])",
          "format": "time_series",
          "hide": false,
          "interval": "$interval",
          "intervalFactor": 1,
          "legendFormat": "{{legacy_op_type}}",
          "refId": "J",
          "step": 300
        },
        {
          "expr": "rate(mongodb_mongod_op_counters_repl_total{instance=~\"$host\", type!~\"(command|query|getmore)\"}[$interval]) or \nirate(mongodb_mongod_op_counters_repl_total{instance=~\"$host\", type!~\"(command|query|getmore)\"}[5m]) or\nrate(mongodb_mongos_op_counters_repl_total{instance=~\"$host\", type!~\"(command|query|getmore)\"}[$interval]) or \nirate(mongodb_mongos_op_counters_repl_total{instance=~\"$host\", type!~\"(command|query|getmore)\"}[5m])",
          "format": "time_series",
          "hide": false,
          "interval": "$interval",
          "intervalFactor": 1,
          "legendFormat": "repl_{{type}}",
          "refId": "A",
          "step": 300
        },
        {
          "expr": "rate(mongodb_mongod_metrics_ttl_deleted_documents_total{instance=~\"$host\"}[$interval]) or \nirate(mongodb_mongod_metrics_ttl_deleted_documents_total{instance=~\"$host\"}[5m]) or\nrate(mongodb_mongos_metrics_ttl_deleted_documents_total{instance=~\"$host\"}[$interval]) or \nirate(mongodb_mongos_metrics_ttl_deleted_documents_total{instance=~\"$host\"}[5m])",
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
        "sort": 0,
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
      "datasource": "prometheus",
      "decimals": null,
      "description": "Keep in mind the hard limit on the maximum number of connections set by your distribution.",
      "editable": true,
      "error": false,
      "fill": 2,
      "grid": {
        "leftLogBase": 1,
        "leftMax": null,
        "leftMin": 0,
        "rightLogBase": 1,
        "rightMax": null,
        "rightMin": null
      },
      "gridPos": {
        "h": 7,
        "w": 12,
        "x": 0,
        "y": 7
      },
      "height": "250px",
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
          "expr": "mongodb_ss_connections{instance=~\"$host\", conn_type=\"current\"}",
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
        "sort": 0,
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
      "datasource": "prometheus",
      "decimals": null,
      "description": "Helps identify why connections are increasing. Shows active cursors compared to cursors being automatically killed after 10 minutes due to an application not closing the connection.",
      "editable": true,
      "error": false,
      "fill": 2,
      "grid": {
        "leftLogBase": 1,
        "leftMax": null,
        "leftMin": 0,
        "rightLogBase": 1,
        "rightMax": null,
        "rightMin": null
      },
      "gridPos": {
        "h": 7,
        "w": 12,
        "x": 12,
        "y": 7
      },
      "height": "250px",
      "id": 25,
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
          "expr": "mongodb_ss_metrics_cursor_open{instance=~\"$host\"}",
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
        "sort": 0,
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
      "datasource": "prometheus",
      "decimals": null,
      "description": "When used in combination with 'Command Operations', this graph can help identify write amplification. For example, when one insert or update command actually inserts or updates hundreds, thousands, or even millions of documents.",
      "editable": true,
      "error": false,
      "fill": 2,
      "grid": {
        "leftLogBase": 1,
        "leftMax": null,
        "leftMin": 0,
        "rightLogBase": 1,
        "rightMax": null,
        "rightMin": null
      },
      "gridPos": {
        "h": 7,
        "w": 12,
        "x": 0,
        "y": 14
      },
      "height": "250px",
      "id": 36,
      "legend": {
        "alignAsTable": true,
        "avg": true,
        "current": false,
        "hideZero": true,
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
          "expr": "rate(mongodb_ss_metrics_document{instance=~\"$host\"}[$interval]) or \nirate(mongodb_ss_metrics_document{instance=~\"$host\"}[5m])",
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
        "sort": 0,
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
      "datasource": "prometheus",
      "decimals": null,
      "description": "The periods (in us) waited by application thread for global lock to further operate",
      "editable": true,
      "error": false,
      "fill": 2,
      "grid": {
        "leftLogBase": 1,
        "leftMax": null,
        "leftMin": 0,
        "rightLogBase": 1,
        "rightMax": null,
        "rightMin": null
      },
      "gridPos": {
        "h": 7,
        "w": 12,
        "x": 12,
        "y": 14
      },
      "height": "250px",
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
          "expr": "  mongodb_ss_wt_lock_txn_global_lock_application_thread_time_waiting_usecs{instance=~\"$host\"}",
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
      "title": "Global lock waiting time by Application thread",
      "tooltip": {
        "msResolution": false,
        "shared": true,
        "sort": 0,
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
          "format": "µs",
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
      "datasource": "prometheus",
      "decimals": null,
      "description": "Asserts are not important by themselves, but you can correlate spikes with other graphs.",
      "editable": true,
      "error": false,
      "fill": 2,
      "grid": {
        "leftLogBase": 1,
        "leftMax": null,
        "leftMin": 0,
        "rightLogBase": 1,
        "rightMax": null,
        "rightMin": null
      },
      "gridPos": {
        "h": 7,
        "w": 12,
        "x": 0,
        "y": 21
      },
      "height": "250px",
      "id": 37,
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
          "expr": "rate(mongodb_ss_asserts{instance=~\"$host\"}[$interval]) or \nirate(mongodb_ss_asserts{instance=~\"$host\"}[5m])",
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
        "sort": 0,
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
      "datasource": "prometheus",
      "decimals": null,
      "description": "Page faults indicate that requests are processed from disk either because an index is missing or there is not enough memory for the data set. Consider increasing memory or sharding out.",
      "editable": true,
      "error": false,
      "fill": 2,
      "grid": {
        "leftLogBase": 1,
        "leftMax": null,
        "leftMin": 0,
        "rightLogBase": 1,
        "rightMax": null,
        "rightMin": null
      },
      "gridPos": {
        "h": 7,
        "w": 12,
        "x": 12,
        "y": 21
      },
      "height": "250px",
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
          "expr": "rate(mongodb_ss_extra_info_page_faults{instance=~\"$host\"}[$interval]) or \nirate(mongodb_ss_extra_info_page_faults{instance=~\"$host\"}[5m])",
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
        "sort": 0,
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
    }
  ],
  "refresh": false,
  "schemaVersion": 16,
  "style": "dark",
  "tags": [
    "MongoDB",
    "Percona"
  ],
  "templating": {
    "list": [
      {
        "allFormat": "glob",
        "allValue": "",
        "auto": true,
        "auto_count": 200,
        "auto_min": "10s",
        "current": {
          "selected": true,
          "text": "auto",
          "value": "$__auto_interval_interval"
        },
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
        "regex": "",
        "skipUrlSync": false,
        "sort": 0,
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
        "includeAll": true,
        "label": "Cluster",
        "multi": false,
        "multiFormat": "glob",
        "name": "cluster",
        "options": [],
        "query": "label_values(mongodb_ss_connections, rs_nm)",
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
        "allValue": "",
        "current": {
          "selected": true,
          "text": "10.33.0.248:9104",
          "value": "10.33.0.248:9104"
        },
        "datasource": "prometheus",
        "definition": "",
        "hide": 0,
        "includeAll": false,
        "label": "Instance",
        "multi": false,
        "multiFormat": "glob",
        "name": "host",
        "options": [],
        "query": "label_values({__name__=~\"mongodb_ss_connections{cluster=~\\\"$cluster\\\"}|mongodb_up\"}, instance)",
        "refresh": 2,
        "regex": "",
        "skipUrlSync": false,
        "sort": 1,
        "tagValuesQuery": "",
        "tags": [],
        "tagsQuery": "",
        "type": "query",
        "useTags": false
      }
    ]
  },
  "time": {
    "from": "now-3h",
    "to": "now"
  },
  "timepicker": {
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
  "timezone": "browser",
  "title": "MongoDB Overview",
  "uid": "6Lk9wMHik",
  "version": 15
}
`

const GrafanaName = "mongodb-overview"

func GenerateGrafana(ns string) *gdv1beta1.GrafanaDashboard {
	return &gdv1beta1.GrafanaDashboard{
		ObjectMeta: metav1.ObjectMeta{
			Name:       GrafanaName,
			Namespace:  ns,
			Finalizers: []string{"grafana.dashboard.finalizers.ait.cpaas.io"},
		},
		Spec: gdv1beta1.GrafanaDashboardSpec{
			Folder: "MongoDB",
			Json:   MongoGrafanaJson,
		},
	}
}
