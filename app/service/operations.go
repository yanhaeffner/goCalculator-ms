package service

type Operation struct {
	Signal       string  `json:"signal"`
	FirstNumber  float64 `json:"firstNumber"`
	SecondNumber float64 `json:"secondNumber"`
	Result       float64 `json:"result"`
	ErrorMsg	 string	 `json:"errorMsg"`
}

func Sum(a float64, b float64) (Operation){
	operationResult := Operation{
		Signal: "+",
		FirstNumber: a,
		SecondNumber: b,
		Result: a+b,
		ErrorMsg: "",
	}
	return operationResult
}

func Sub(a float64, b float64) (Operation){
	operationResult := Operation{
		Signal: "-",
		FirstNumber: a,
		SecondNumber: b,
		Result: a-b,
		ErrorMsg: "",
	}
	return operationResult
}

func Mul(a float64, b float64) (Operation){
	operationResult := Operation{
		Signal: "*",
		FirstNumber: a,
		SecondNumber: b,
		Result: a*b,
		ErrorMsg: "",
	}
	return operationResult
}

func Div(a float64, b float64) (Operation){
	result := 0.0
	errorMsg := ""
	if b == 0{
		result = 0
		errorMsg = "Denominator can not be zero"
	} else{
		result = a/b
		errorMsg = ""
	}
	operationResult := Operation{
		Signal: "/",
		FirstNumber: a,
		SecondNumber: b,
		Result: result,
		ErrorMsg: errorMsg,
	}
	return operationResult
}