// Code generated - EDITING IS FUTILE. DO NOT EDIT.
//
// Generated by:
//     public/app/plugins/gen.go
// Using jennies:
//     PluginGoTypesJenny
//
// Run 'make gen-cue' from repository root to regenerate.

package dataquery

// Defines values for LokiQueryDirection.
const (
	LokiQueryDirectionBackward LokiQueryDirection = "backward"
	LokiQueryDirectionForward  LokiQueryDirection = "forward"
	LokiQueryDirectionScan     LokiQueryDirection = "scan"
)

// Defines values for LokiQueryType.
const (
	LokiQueryTypeInstant LokiQueryType = "instant"
	LokiQueryTypeRange   LokiQueryType = "range"
	LokiQueryTypeStream  LokiQueryType = "stream"
)

// Defines values for QueryEditorMode.
const (
	QueryEditorModeBuilder QueryEditorMode = "builder"
	QueryEditorModeCode    QueryEditorMode = "code"
)

// Defines values for SupportingQueryType.
const (
	SupportingQueryTypeDataSample     SupportingQueryType = "dataSample"
	SupportingQueryTypeInfiniteScroll SupportingQueryType = "infiniteScroll"
	SupportingQueryTypeLogsSample     SupportingQueryType = "logsSample"
	SupportingQueryTypeLogsVolume     SupportingQueryType = "logsVolume"
)

// These are the common properties available to all queries in all datasources.
// Specific implementations will *extend* this interface, adding the required
// properties for the given context.
type DataQuery struct {
	// For mixed data sources the selected datasource is on the query level.
	// For non mixed scenarios this is undefined.
	// TODO find a better way to do this ^ that's friendly to schema
	// TODO this shouldn't be unknown but DataSourceRef | null
	Datasource *any `json:"datasource,omitempty"`

	// If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
	Hide *bool `json:"hide,omitempty"`

	// Specify the query flavor
	// TODO make this required and give it a default
	QueryType *string `json:"queryType,omitempty"`

	// A unique identifier for the query within the list of targets.
	// In server side expressions, the refId is used as a variable name to identify results.
	// By default, the UI will assign A->Z; however setting meaningful names may be useful.
	RefId string `json:"refId"`
}

// LokiDataQuery defines model for LokiDataQuery.
type LokiDataQuery struct {
	// For mixed data sources the selected datasource is on the query level.
	// For non mixed scenarios this is undefined.
	// TODO find a better way to do this ^ that's friendly to schema
	// TODO this shouldn't be unknown but DataSourceRef | null
	Datasource *any             `json:"datasource,omitempty"`
	EditorMode *QueryEditorMode `json:"editorMode,omitempty"`

	// The LogQL query.
	Expr *string `json:"expr,omitempty"`

	// If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
	Hide *bool `json:"hide,omitempty"`

	// @deprecated, now use queryType.
	Instant *bool `json:"instant,omitempty"`

	// Used to override the name of the series.
	LegendFormat *string `json:"legendFormat,omitempty"`

	// Used to limit the number of log rows returned.
	MaxLines *int64 `json:"maxLines,omitempty"`

	// Specify the query flavor
	// TODO make this required and give it a default
	QueryType *string `json:"queryType,omitempty"`

	// @deprecated, now use queryType.
	Range *bool `json:"range,omitempty"`

	// A unique identifier for the query within the list of targets.
	// In server side expressions, the refId is used as a variable name to identify results.
	// By default, the UI will assign A->Z; however setting meaningful names may be useful.
	RefId *string `json:"refId,omitempty"`

	// @deprecated, now use step.
	Resolution *int64 `json:"resolution,omitempty"`

	// Used to set step value for range queries.
	Step *string `json:"step,omitempty"`
}

// LokiQueryDirection defines model for LokiQueryDirection.
type LokiQueryDirection string

// LokiQueryType defines model for LokiQueryType.
type LokiQueryType string

// QueryEditorMode defines model for QueryEditorMode.
type QueryEditorMode string

// SupportingQueryType defines model for SupportingQueryType.
type SupportingQueryType string
