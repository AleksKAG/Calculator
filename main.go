package main
import ("fmt"
        "net/http"
        
)
type Calculation struct {
ID string 'json:"id"'
Expression string 'json:"expression"'
Result string 'json:"result"'
}

type CalculationRequest struct {
Expression string 'json:"expression"'
}

var calculations = []Calculation{}

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
  calc := Calculation {
    ID: uuid.NewString(),
    Expression: req.Expression,
    Result: result,
  }
  calculations = append(calculations, calc)
  return c.JSON(http.StatusCreated, calc)
}

func patchCalculations(c echo.Context) error {
id := c.Param(name: "id")
        
var req CalculationRequest
  if err := c.Bind(&req); err != nil {
    return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
  }
  result, err := calculatteExpression(req.Expression)
if err != nil {
  return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid expression"})
}
for i, calculation := range calculations {
if calculation.ID == id {
        calculations[i].Expression = req.Expression
        calculations[i].Result = req.Result
        return c.JSON(http.StatusOk, calculations[i])
}
}
}

  
}

func main() {
e := echo.New()
  e.Use(middleware.CORS())
  e.Use(middleware.Logger())
  e.GET(path: "/calculations", getCalculations)
  e.Start(address: "localhost:8080")
}
