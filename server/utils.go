package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"tats-backend/api"
	"tats-backend/config"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var layout = "2006-01-02"
var osgLayout = "20060102"
var secondLayout = "2006-01-02 15:04:05"

var versionPat = regexp.MustCompile(`/v[1-9]/`)

func changeVersionToUnderscore(in string) string {
	return versionPat.ReplaceAllString(in, "/_/")
}
func mergeCoverUrl(in string) string {
	return fmt.Sprintf("%s/%s", changeVersionToUnderscore(config.Cfg.OsgUrl), in)
}

func pass2Vpsformat(t string) (time.Time, error) {
	if t == "" {
		return time.Now(), nil
	}
	return time.Parse(layout, t)
}

func timeFormat(t time.Time) string {
	return t.Format(secondLayout)
}
func timeFormatShort(t time.Time) string {
	return t.Format(layout)
}

func Einter(e error) error {
	if e == nil {
		return status.Error(codes.Internal, "unkown")
	}
	return status.Error(codes.Internal, e.Error())
}

func Einvalid(e error) error {
	if e == nil {
		return status.Error(codes.Internal, "unkown")
	}
	return status.Error(codes.InvalidArgument, e.Error())
}

func wrapCode(code int, err error) error {
	if err == nil {
		return status.New(codes.Code(code), "unkown").Err()
	}
	return status.New(codes.Code(code), err.Error()).Err()
}

func success() (*api.SuccessResp, error) {
	return &api.SuccessResp{
		Msg: "success",
	}, nil
}

// http

type Rsp struct {
	Err  string      `json:"error,omitempty"`
	Code int         `json:"code,omitempty"`
	Msg  string      `json:"message,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

func fail(w http.ResponseWriter, http_code int, err error) {
	r := Rsp{
		Err:  err.Error(),
		Code: -1,
		Msg:  "failed",
	}
	writeResp(w, r, http_code)
}
func writeResp(w http.ResponseWriter, r interface{}, code int) {
	cont, _ := json.Marshal(r)
	w.Header().Add("Content-Type", "application/json")
	if code != http.StatusOK {
		w.WriteHeader(code)
	}
	w.Write(cont)
}

func succ(w http.ResponseWriter, data interface{}) {
	r := Rsp{
		Msg:  "success",
		Data: data,
	}
	writeResp(w, r, 200)
}
