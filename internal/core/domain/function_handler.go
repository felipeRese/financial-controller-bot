package domain


type FunctionHandler interface {
  GetAvailableFunctions() []Function
	ExecuteFunction(name string, args map[string]interface{}) (string, error)
}

