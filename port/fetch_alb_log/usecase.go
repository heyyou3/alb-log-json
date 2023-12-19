package fetch_alb_log

import (
	"alb-log-json/adapter/output/filestorage"
	"alb-log-json/domain/alb_log_struct"
	"context"
	"errors"
	"golang.org/x/sync/errgroup"
	"net/url"
	"regexp"
	"strings"
)

var (
	regALB = regexp.MustCompile(`(.+) (.+) (.+) (.+) (.+) (.+) (.+) (.+) (.+) (.+) (.+) (.+) "(.+)" "(.+)" (.+) (.+) (.+) "(.+)" "(.+)" "(.+)" (.+) (.+) "(.+)" "(.+)" "(.+)" "(.+)" "(.+)"`)
	regReq = regexp.MustCompile(`(.+) (.+) (.+)`)
)

type IFetchALBLogUsecase interface {
	Invoke() ([]*alb_log_struct.ALBLogStruct, error)
}
type FetchALBLogUsecase struct {
	Output           IFetchALBLogOutput
	FetchALBLogParam filestorage.FetchALBLogParam
}

// Invoke
// NOTE: fetch alb log return []*alb_log_struct.ALBLogStruct
//
//	Convert stored S3 logs from gz to txt
//	obligation
//	- fetch s3 object(directory)
//	- convert gz file to []*alb_log_struct.ALBLogStruct
func (u *FetchALBLogUsecase) Invoke() ([]*alb_log_struct.ALBLogStruct, error) {
	var res []*alb_log_struct.ALBLogStruct
	albLogs, err := u.Output.FetchALBLog(u.FetchALBLogParam)
	if err != nil {
		return nil, err
	}
	eg, _ := errgroup.WithContext(context.Background())
	// NOTE: for local processing, Parallel Limit 1000
	eg.SetLimit(1000)
	for _, log := range albLogs {
		log := log
		eg.Go(func() error {
			data := regALB.FindStringSubmatch(log)
			if len(data) < 25 {
				// TODO: invalid Log Data think error handling...
				return nil
			}

			for i, _ := range data {
				data[i] = strings.Trim(data[i], `"`)
			}

			method, version, protocol, host, port, uri, err := parseRequest(data[13])

			if err != nil {
				// TODO: invalid Log Data think error handling...
				return nil
			}

			res = append(res, &alb_log_struct.ALBLogStruct{
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
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return nil, err
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

	if len(req) < 4 || len(hostPort) < 2 {
		return "", "", "", "", "", "", errors.New("invalid data")
	}

	return req[1], req[3], u.Scheme, hostPort[0], hostPort[1], u.Path, nil
}
