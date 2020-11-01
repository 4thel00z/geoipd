package libgeoip

import (
	"errors"
	"plugin"
)

type Pipeline interface {
	Consume(objs ...interface{}) (interface{}, error)
	Init(objs ...interface{}) (interface{}, error)
}

func LoadPipeline(path string) (Pipeline, error) {

	plugin, err := plugin.Open(path)

	if err != nil {
		return nil, err
	}

	pipelineRaw, err := plugin.Lookup("Pipeline")

	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("could not cast symbol Pipeline to type Pipeline")
		}
	}()

	pipeline := pipelineRaw.(Pipeline)

	return pipeline, err
}
