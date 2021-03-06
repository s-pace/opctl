package envvars

import (
	"fmt"

	"github.com/opctl/opctl/sdks/go/data/coerce"
	"github.com/opctl/opctl/sdks/go/model"
	"github.com/opctl/opctl/sdks/go/opspec/interpreter/object"
	"github.com/opctl/opctl/sdks/go/opspec/interpreter/str"
)

//counterfeiter:generate -o fakes/interpreter.go . Interpreter
type Interpreter interface {
	Interpret(
		scope map[string]*model.Value,
		scgContainerCallEnvVars interface{},
	) (map[string]string, error)
}

// NewInterpreter returns a new Interpreter instance
func NewInterpreter() Interpreter {
	return _interpreter{
		objectInterpreter: object.NewInterpreter(),
		stringInterpreter: str.NewInterpreter(),
		coerce:            coerce.New(),
	}
}

type _interpreter struct {
	coerce            coerce.Coerce
	objectInterpreter object.Interpreter
	stringInterpreter str.Interpreter
}

func (itp _interpreter) Interpret(
	scope map[string]*model.Value,
	scgContainerCallEnvVars interface{},
) (map[string]string, error) {
	if nil == scgContainerCallEnvVars {
		return nil, nil
	}

	envVarsMap, err := itp.objectInterpreter.Interpret(
		scope,
		scgContainerCallEnvVars,
	)
	if nil != err {
		return nil, fmt.Errorf(
			"unable to interpret '%v' as envVars; error was %v",
			scgContainerCallEnvVars,
			err,
		)
	}

	envVarsStringMap := map[string]string{}
	for envVarName, envVarValue := range *envVarsMap.Object {
		envVarValueString, err := itp.stringInterpreter.Interpret(scope, envVarValue)
		if nil != err {
			return nil, fmt.Errorf(
				"unable to interpret %+v as value of env var '%v'; error was %v",
				envVarValue,
				envVarName,
				err,
			)
		}

		envVarsStringMap[envVarName] = *envVarValueString.String
	}

	return envVarsStringMap, nil
}
