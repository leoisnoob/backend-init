// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mod.proto

package api

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	regexp "regexp"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var _regex_PingReq_Msg = regexp.MustCompile(`^[a-z]{2,6}$`)

func (this *PingReq) Validate() error {
	if !_regex_PingReq_Msg.MatchString(this.Msg) {
		return github_com_mwitkow_go_proto_validators.FieldError("Msg", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-z]{2,6}$"`, this.Msg))
	}
	return nil
}
func (this *PingResp) Validate() error {
	return nil
}

var _regex_CreateUserReq_Username = regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9\-_@\.]{3,20}$`)

func (this *CreateUserReq) Validate() error {
	if !_regex_CreateUserReq_Username.MatchString(this.Username) {
		return github_com_mwitkow_go_proto_validators.FieldError("Username", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z0-9][a-zA-Z0-9\\-_@\\.]{3,20}$"`, this.Username))
	}
	if this.Password == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Password", fmt.Errorf(`value '%v' must not be an empty string`, this.Password))
	}
	return nil
}
func (this *CreateUserResp) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *UpdateUserReq) Validate() error {
	if this.Username == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Username", fmt.Errorf(`value '%v' must not be an empty string`, this.Username))
	}
	if this.User != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.User); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("User", err)
		}
	}
	return nil
}
func (this *UpdateUserBody) Validate() error {
	if this.NewPassword == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("NewPassword", fmt.Errorf(`value '%v' must not be an empty string`, this.NewPassword))
	}
	return nil
}
func (this *UpdateUserResp) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *DeleteUserReq) Validate() error {
	if this.Username == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Username", fmt.Errorf(`value '%v' must not be an empty string`, this.Username))
	}
	return nil
}

var _regex_LoginReq_Username = regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9\-_@\.]{3,20}$`)

func (this *LoginReq) Validate() error {
	if !_regex_LoginReq_Username.MatchString(this.Username) {
		return github_com_mwitkow_go_proto_validators.FieldError("Username", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z0-9][a-zA-Z0-9\\-_@\\.]{3,20}$"`, this.Username))
	}
	if this.Password == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Password", fmt.Errorf(`value '%v' must not be an empty string`, this.Password))
	}
	return nil
}
func (this *LoginResp) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *LogoutReq) Validate() error {
	return nil
}
func (this *ListUsersReq) Validate() error {
	return nil
}
func (this *ListUsersResp) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *User) Validate() error {
	return nil
}
func (this *SuccessResp) Validate() error {
	return nil
}
