package controllers

import(
  "testing"
  "os"
  "api/services"
  "api/utils/errors"
  "api/domain/locations"
  "net/http"
  "net/http/httptest"
  "encoding/json"
  "github.com/mercadolibre/golang-restclient/rest"
  "github.com/stretchr/testify/assert"
  "github.com/gin-gonic/gin"
)

var (
  getCountryFunc func(countryId string) (*locations.Country, *errors.ApiError)
)

func TestMain(m *testing.M) {
  rest.StartMockupServer()
  os.Exit(m.Run())
}

type locationsServiceMock struct {}

func (* locationsServiceMock) GetCountry(countryId string) (*locations.Country, *errors.ApiError) {
  return getCountryFunc(countryId)
}

func TestGetCountryNotFound(t *testing.T) {
  // Mock LocationsService
  getCountryFunc = func(countryId string) (*locations.Country, *errors.ApiError) {
    return nil, &errors.ApiError{Status: http.StatusNotFound, Message: "Country not found"}
  }

  services.LocationsService = &locationsServiceMock{}

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
