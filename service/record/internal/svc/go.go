package svc

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/j128919965/gopkg/errors"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/valyala/fasthttp"
	"queoj/service/problem/problemclient"
	"queoj/service/record/internal/model"
	record_status "queoj/service/record/internal/svc/record-status"
	"strconv"
)

const GoCompileReq = `{
    "cmd": [{
        "args": ["/usr/local/go/bin/go","build","main.go"],
        "env": ["GOPATH=/go","GOCACHE=/w/go-build-cache"],
        "files": [{
            "content": ""
        }, {
            "name": "stdout",
            "max": 10240
        }, {
            "name": "stderr",
            "max": 10240
        }],
        "cpuLimit": 10000000000,
        "memoryLimit": 104857600,
        "procLimit": 50,
        "strictMemoryLimit": false,
        "copyOut": ["stdout", "stderr"],
        "copyOutCached": ["main"],
        "copyIn": {
            "main.go": {
                "content": %s
            }
        }
    }]
}`

var GoRunReq = `
{
    "cmd": [{
        "args": ["main"],
        "env": ["PATH=/usr/bin:/bin"],
        "files": [{
            "content": "%s"
        }, {
            "name": "stdout",
            "max": 10240
        }, {
            "name": "stderr",
            "max": 10240
        }],
        "cpuLimit": %d,
        "memoryLimit": %d,
        "procLimit": 50,
        "strictMemoryLimit": false,
        "copyIn": {
            "main": {
                "fileId": "%s"
            }
        }
    }]
}
`

func (svc *ServiceContext) submitGo(record *model.Record) {
	// 编译
	classId, err := compileGo(&record.Code)
	if err != nil {
		logx.Error(err)
		svc.UpdateRecordStatus(record.Id, record_status.CompileError)
		return
	}
	io, err := svc.ProblemClient.GetProblemIO(context.Background(), &problemclient.Integer{Value: record.Pid})
	if err != nil {
		logx.Error(err)
		svc.UpdateRecordStatus(record.Id, record_status.InternalError)
		return
	}
	detail, err := svc.ProblemClient.GetProblemDetail(context.Background(), &problemclient.Integer{Value: record.Pid})
	if err != nil {
		return
	}
	status, result, err := runGo(classId,io.InTxt,detail.TimeLimit,detail.SpaceLimit)
	if err != nil {
		logx.Error(err)
		svc.UpdateRecordStatus(record.Id, record_status.InternalError)
		return
	}
	if status != 1 {
		svc.UpdateRecordStatus(record.Id, status)
		return
	}

	if result.output != io.OutTxt {
		svc.UpdateRecordStatus(record.Id,record_status.WrongAnswer)
		return
	}

	logx.Info(fmt.Sprintf("judge cpp {%d} success .",record.Id))
	//svc.UpdateRecordResult(record.Id,result.timeUsed,result.spaceUsed,record)
}

func compileGo(code *string) (string, error) {
	req := &fasthttp.Request{}
	compileStr := fmt.Sprintf(GoCompileReq, strconv.Quote(*code))
	fmt.Println(compileStr)
	req.SetBody([]byte(compileStr))
	// 默认是application/x-www-form-urlencoded
	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")
	req.SetRequestURI(JudgeUrl)

	resp := &fasthttp.Response{}
	client := &fasthttp.Client{}
	err := client.Do(req, resp)
	if err != nil {
		logx.Error(err)
		panic(err.Error())
	}

	var m []interface{}
	body := resp.Body()
	err = json.Unmarshal(body, &m)
	if err != nil {
		logx.Error(string(body))
		return "", err
	}

	respMap := m[0].(map[string]interface{})
	if respMap["status"] != "Accepted" {
		err := respMap["files"].(map[string]interface{})["stderr"].(string)
		return "", errors.New(err, 0)
	}

	classId := respMap["fileIds"].(map[string]interface{})["main"].(string)
	return classId, nil
}


func runGo(classId, input string, timeLimit, spaceLimit uint64) (uint32, *JudgeResult, error) {
	req := &fasthttp.Request{}
	compileStr := fmt.Sprintf(GoRunReq, input, timeLimit,spaceLimit,classId)
	req.SetBody([]byte(compileStr))
	// 默认是application/x-www-form-urlencoded
	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")
	req.SetRequestURI(JudgeUrl)

	resp := &fasthttp.Response{}
	client := &fasthttp.Client{}
	err := client.Do(req, resp)
	if err != nil {
		logx.Error(err)
		panic(err.Error())
	}

	var m []interface{}
	err = json.Unmarshal(resp.Body(), &m)
	if err != nil {
		return record_status.InternalError,nil, err
	}

	respMap := m[0].(map[string]interface{})
	if respMap["status"] != "Accepted" {
		switch respMap["status"] {
		case "Internal Error":
			return record_status.InternalError, nil, nil
		case "Memory Limit Exceeded":
			return record_status.MemoryLimitExceeded, nil, nil
		case "Time Limit Exceeded":
			return record_status.TimeLimitExceeded, nil, nil
		case "Output Limit Exceeded":
			return record_status.OutputLimitExceeded, nil, nil
		case "File Error":
			return record_status.FileError, nil, nil
		case "Nonzero Exit Status":
			return record_status.NonzeroExitStatus, nil, nil
		case "Signalled":
			return record_status.Signalled, nil, nil
		}
	}

	out := respMap["files"].(map[string]interface{})["stdout"].(string)
	tl := uint64(respMap["time"].(float64))
	sl := uint64(respMap["memory"].(float64))

	return record_status.Accept, &JudgeResult{
		output:    out,
		timeUsed:  tl,
		spaceUsed: sl,
	}, nil
}
