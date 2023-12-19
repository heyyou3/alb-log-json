package fetch_alb_log

import (
	"alb-log-parser/adapter/output/filestorage"
	"net/url"
	"regexp"
	"strings"
)

var (
	regALB = regexp.MustCompile(`(.+) (.+) (.+) (.+) (.+) (.+) (.+) (.+) (.+) (.+) (.+) (.+) "(.+)" "(.+)" (.+) (.+) (.+) "(.+)" "(.+)" "(.+)" (.+) (.+) "(.+)" "(.+)" "(.+)" "(.+)" "(.+)"`)
	regReq = regexp.MustCompile(`(.+) (.+) (.+)`)
)

type FetchALBLogStruct struct {
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
type IFetchALBLogUsecase interface {
	Invoke() ([]*FetchALBLogStruct, error)
}
type FetchALBLogUsecase struct {
	Output           IFetchALBLogOutput
	FetchALBLogParam filestorage.FetchALBLogParam
}

// Invoke
// NOTE: fetch alb log return []*FetchALBLogStruct
//
//	Convert stored S3 logs from gz to txt
//	obligation
//	- fetch s3 object(directory)
//	- convert gz file to []*FetchALBLogStruct
func (u *FetchALBLogUsecase) Invoke() ([]*FetchALBLogStruct, error) {
	var res []*FetchALBLogStruct
	albLogs, err := u.Output.FetchALBLog(u.FetchALBLogParam)
	if err != nil {
		return nil, err
	}
	for _, log := range albLogs {
		data := regALB.FindStringSubmatch(log)
		if len(data) < 25 {
			// TODO: invalid Log Data output log
			continue
		}

		for i, _ := range data {
			data[i] = strings.Trim(data[i], `"`)
		}

		method, version, protocol, host, port, uri, err := parseRequest(data[13])

		if err != nil {
			// TODO: invalid Log Data output log
			continue
		}

		res = append(res, &FetchALBLogStruct{
			Type:                   data[1],
			Timestamp:              data[2],
			Elb:                    data[3],
			Client:                 data[4],
			Target:                 data[5],
			RequestProcessingTime:  data[6],
			TargetProcessingTime:   data[7],
			ResponseProcessingTime: data[8],
			ElbStatusCode:          data[9],
			TargetStatusCode:       data[10],
			ReceivedBytes:          data[11],
			SentBytes:              data[12],
			Method:                 method,
			HttpVersion:            version,
			Protocol:               protocol,
			Host:                   host,
			Port:                   port,
			Uri:                    uri,
			UserAgent:              data[14],
			SslCipher:              data[15],
			SslProtocol:            data[16],
			TargetGroupArn:         data[17],
			TraceId:                data[18],
			DomainName:             data[19],
			ChosenCertArn:          data[20],
			MatchedRulePriority:    data[21],
			RequestCreationTime:    data[22],
			ActionsExecuted:        data[23],
			RedirectUrl:            data[24],
			ErrorReason:            data[25],
			TargetPortList:         data[26],
			TargetStatusCodeList:   data[27],
		})
	}
	return res, nil
}

func parseRequest(data string) (string, string, string, string, string, string, error) {
	req := regReq.FindStringSubmatch(data)
	u, err := url.Parse(req[2])
	if err != nil {
		return "", "", "", "", "", "", err
	}

	hostPort := strings.Split(u.Host, `:`)

	return req[1], req[3], u.Scheme, hostPort[0], hostPort[1], u.Path, nil
}
