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
func getCalculations(c echo.Context) error {
 return c.JSON(http.StatusOk, calculations) 
}

func postCalculations(c echo.Context) error {
var req CalculationRequest
  if err := c.Bind(&req); err != nil {
    return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
  }
  result, err := calculatteExpression(req.Expression)
if err != nil {
  return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid expression"})
}
  
}

func main() {
e := echo.New()
  e.Use(middleware.CORS())
  e.Use(middleware.Logger())
  e.GET(path: "/calculations", getCalculations)
  e.Start(address: "localhost:8080")
}
