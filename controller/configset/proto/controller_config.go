package configset_proto

import (
	"context"
	"encoding/json"

	"github.com/aperturerobotics/controllerbus/bus"
	"github.com/aperturerobotics/controllerbus/config"
	"github.com/aperturerobotics/controllerbus/controller/configset"
	"github.com/aperturerobotics/controllerbus/controller/resolver"
	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
	"github.com/valyala/fastjson"
	jsonpb "google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// NewControllerConfig constructs a new controller config.
func NewControllerConfig(c configset.ControllerConfig) (*ControllerConfig, error) {
	conf := c.GetConfig()
	cID := conf.GetConfigID()
	confData, err := proto.Marshal(conf)
	if err != nil {
		return nil, err
	}
	return &ControllerConfig{
		Id:       cID,
		Config:   confData,
		Revision: c.GetRevision(),
	}, nil
}

// Validate performs cursory validation of the controller config.
func (c *ControllerConfig) Validate() error {
	if len(c.GetId()) == 0 {
		return ErrControllerConfigIdEmpty
	}
	if conf := c.GetConfig(); len(conf) != 0 {
		// json if first character is {
		if conf[0] == 123 {
			// validate json
			var p fastjson.Parser
			_, err := p.ParseBytes(conf)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Resolve resolves the config into a configset.ControllerConfig
func (c *ControllerConfig) Resolve(ctx context.Context, b bus.Bus) (configset.ControllerConfig, error) {
	if len(c.GetId()) == 0 {
		return nil, ErrControllerConfigIdEmpty
	}

	configCtorDir := resolver.NewLoadConfigConstructorByID(c.GetId())
	configCtorVal, _, configCtorRef, err := bus.ExecOneOff(ctx, b, configCtorDir, false, nil)
	if err != nil {
		if err == context.Canceled {
			return nil, err
		}
		return nil, errors.WithMessage(err, "resolve config object")
	}
	defer configCtorRef.Release()

	ctor, ctorOk := configCtorVal.GetValue().(config.Constructor)
	if !ctorOk {
		return nil, errors.New("load config constructor directive returned invalid object")
	}
	cf := ctor.ConstructConfig()

	// detect json or protobuf & parse
	if configData := c.GetConfig(); len(configData) != 0 {
		// if configData[0] == '{'
		if configData[0] == 123 {
			err = jsonpb.Unmarshal(configData, cf)
		} else {
			err = proto.Unmarshal(configData, cf)
		}
		if err != nil {
			return nil, err
		}
	}

	return configset.NewControllerConfig(c.GetRevision(), cf), nil
}

// UnmarshalJSON unmarshals json to the controller config.
// For the config field: supports JSON, YAML, or a string containing either.
func (c *ControllerConfig) UnmarshalJSON(data []byte) error {
	jdata, err := yaml.YAMLToJSON(data)
	if err != nil {
		return err
	}
	var p fastjson.Parser
	v, err := p.ParseBytes(jdata)
	if err != nil {
		return err
	}
	if v.Exists("id") {
		c.Id = string(v.GetStringBytes("id"))
	}
	if v.Exists("revision") {
		c.Revision = v.GetUint64("revision")
	}
	if v.Exists("config") {
		var configVal *fastjson.Value
		configStr := v.GetStringBytes("config")
		if len(configStr) != 0 {
			// parse json and/or yaml
			configJson, err := yaml.YAMLToJSON(configStr)
			if err != nil {
				return err
			}
			var cj fastjson.Parser
			configVal, err = cj.ParseBytes(configJson)
			if err != nil {
				return err
			}
		} else {
			// expect a object value
			configVal = v.Get("config")
			if t := configVal.Type(); t != fastjson.TypeObject {
				return errors.Errorf("config: expected json object but got %s", t.String())
			}
		}
		// re-marshal to json
		c.Config = configVal.MarshalTo(nil)
	}
	return nil
}

// _ is a type assertion
var (
	_ json.Unmarshaler = ((*ControllerConfig)(nil))
)
