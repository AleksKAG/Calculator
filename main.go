package main
import "fmt"

type Calculation struct {
ID string 'json:"id"'
Expression string 'json:"expression"'
Result string 'json:"result"'
}

type CalculationRequest struct {
Expression string 'json:"expression"'
}

vav calculations = []Calculation{}

func calculateExpression(expression string) (string, error) {
  expr, err := govaluate.NewEvaluableExpression(expression)
  if err != nil {
   return "", err
  }
return fmt.Printf(format: "%v", result), err
}

func main() {
e := echo.New()
  e.Use(middleware.CORS())
  e.Use(middleware.Logger())
}
