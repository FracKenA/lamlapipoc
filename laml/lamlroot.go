package laml

import (
	"strings"
)

type ConfigTree struct {
	Server      string
	Port        string
	Secure      bool
	User        string
	Pass        string
	ContentType string
	Accept      string
}

type LabelFieldTree struct {
	Field   string `json:"field"`
	MapFrom string `json:"map_from"`
}

type ESRuleTree struct {
	Timestamp            string         `json:"timestamp,omitempty"`
	AlgoType             string         `json:"algorithm_type"`
	PrepDate             string         `json:"preparation_date,omitempty"`
	MethodName           string         `json:"method_name"`
	ModelName            string         `json:"model_name"`
	ModelUID             string         `json:"model_uid,omitempty"`
	MachineStateUID      string         `json:"machine_state_uid,omitempty"`
	Search               string         `json:"search"`
	PathToLogs           string         `json:"path_to_logs,omitempty"`
	PathToMachineState   string         `json:"path_to_machine_state,omitempty"`
	LabelField           LabelFieldTree `json:"label_field"`
	Weights              []string       `json:"weights,omitempty"`
	Accuracy             string         `json:"accuracy,omitempty"`
	WeightedPrecision    string         `json:"weighted_precision,omitempty"`
	Layers               string         `json:"layers,omitempty"`
	MetricsData          []string       `json:"metrics_data,omitempty"`
	SearchQueryString    string         `json:"search_query_string,omitempty"`
	MaxIterations        int            `json:"max_iter,omitempty"`
	MaxClass             int            `json:"max_class,omitempty"`
	RegParam             int            `json:"reg_param,omitempty"`
	ElasticnetParam      int            `json:"elastic_net_param,omitempty"`
	MaxProbes            int            `json:"max_probes"`
	Timeframe            string         `json:"time_frame"`
	ValueType            string         `json:"value_type"`
	MaxPredictions       int            `json:"max_predictions"`
	SearchSourceJSON     string         `json:"searchSourceJSON,omitempty"`
	SkipItems            int            `json:"skip_items,omitempty"`
	Threshold            int            `json:"threshold,omitempty"`
	ProcessingTime       int            `json:"processing_time,omitempty"`
	PreditionCycle       string         `json:"prediction_cycle,omitempty"`
	StartDate            string         `json:"start_date"`
	MultiplyByValues     []string       `json:"multiply_by_values"`
	MultiplyByField      string         `json:"multiply_by_field"`
	SelectedRoles        []string       `json:"selectedroles,omitempty"`
	LastExecuteMili      int            `json:"last_execute_mili,omitempty"`
	LastExecuteTimestamp string         `json:"last_execute_timestamp,omitempty"`
	PID                  string         `json:"pid,omitempty"`
	ErrorMessage         string         `json:"error_message,omitempty"`
	ErrorDescription     string         `json:"error_description,omitempty"`
	ExitCode             int            `json:"exit_code,omitempty"`
}

type ESShardsTree struct {
	Total      int `json:"total"`
	Successful int `json:"successful"`
	Failed     int `json:"failed"`
}

type ESRuleCreatedTree struct {
	Index         string       `json:"_index"`
	Type          string       `json:"_type"`
	ID            string       `json:"_id"`
	Version       int          `json:"_version"`
	Result        string       `json:"result"`
	Shards        ESShardsTree `json:"_shards"`
	SeqenceNumber int          `json:"_seq_no"`
	PrimaryTerm   int          `json:"_primary_term"`
}

func BuildString(delim string, stringList ...string) string {
	return strings.Join(stringList, delim)
}
