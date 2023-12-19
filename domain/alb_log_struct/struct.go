package alb_log_struct

type ALBLogStruct struct {
	Type                   string `json:"type"`
	Timestamp              string `json:"timestamp"`
	Elb                    string `json:"elb"`
	Client                 string `json:"client"`
	Target                 string `json:"target"`
	RequestProcessingTime  string `json:"request_processing_time"`
	TargetProcessingTime   string `json:"target_processing_time"`
	ResponseProcessingTime string `json:"response_processing_time"`
	ElbStatusCode          string `json:"elb_status_code"`
	TargetStatusCode       string `json:"target_status_code"`
	ReceivedBytes          string `json:"received_bytes"`
	SentBytes              string `json:"sent_bytes"`
	Method                 string `json:"method"`
	HttpVersion            string `json:"http_version"`
	Protocol               string `json:"protocol"`
	Host                   string `json:"host"`
	Port                   string `json:"port"`
	Uri                    string `json:"uri"`
	UserAgent              string `json:"user_agent"`
	SslCipher              string `json:"ssl_cipher"`
	SslProtocol            string `json:"ssl_protocol"`
	TargetGroupArn         string `json:"target_group_arn"`
	TraceId                string `json:"trace_id"`
	DomainName             string `json:"domain_name"`
	ChosenCertArn          string `json:"chosen_cert_arn"`
	MatchedRulePriority    string `json:"matched_rule_priority"`
	RequestCreationTime    string `json:"request_creation_time"`
	ActionsExecuted        string `json:"actions_executed"`
	RedirectUrl            string `json:"redirect_url"`
	ErrorReason            string `json:"error_reason"`
	TargetPortList         string `json:"target:port_list"`
	TargetStatusCodeList   string `json:"target_status_code_list"`
}
