package controllers

import(
  "testing"
  "os"
  "api/utils/errors"
  "net/http"
  "net/http/httptest"
  "encoding/json"
  "github.com/mercadolibre/golang-restclient/rest"
  "github.com/stretchr/testify/assert"
  "github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
  rest.StartMockupServer()
  os.Exit(m.Run())
}

func TestGetCountryNotFound(t *testing.T) {
  rest.FlushMockups()
  rest.AddMockups(&rest.Mock{
    URL:            "https://api.mercadolibre.com/countries/AR",
    HTTPMethod:     http.MethodGet,
    RespHTTPCode:   http.StatusNotFound,
    RespBody:       `{"message": "Country not found","error": "not_found","status": 404,"cause": []}`,
  })

  response := httptest.NewRecorder()
  c, _ := gin.CreateTestContext(response)
  c.Request, _ = http.NewRequest(http.MethodGet, "", nil)
  c.Params = gin.Params {
    {Key: "country_id", Value: "AR"},
  }

  GetCountry(c)

  assert.EqualValues(t, http.StatusNotFound, response.Code)

  var apiErr errors.ApiError
  err := json.Unmarshal(response.Body.Bytes(), &apiErr)
  assert.Nil(t, err)

  assert.EqualValues(t, http.StatusNotFound, apiErr.Status)
  assert.EqualValues(t, "Country not found", apiErr.Message)
}
