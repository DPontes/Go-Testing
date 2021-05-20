package controllers

import(
  "testing"
  "os"
  "net/http"
  "net/http/httptest"
  "encoding/json"
  "api/domain/locations"
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

  GetCountry(c)

  assert.EqualValues(t, http.StatusOK, response.Code)

  var country locations.Country
  err := json.Unmarshal(response.Body.Bytes(), &country)
  assert.Nil(t, err)

  assert.EqualValues(t, "AR", country.Id)
  assert.EqualValues(t, "Argentina", country.Name)
}
