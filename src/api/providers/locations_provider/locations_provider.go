package locations_provider

import (
  "fmt"
  "net/http"
  "encoding/json"
  "api/domain/locations"
  "api/utils/errors"
  "github.com/mercadolibre/golang-restclient/rest"
)

const (
  urlGetCountry = "https://api.mercadolibre.com/countries/%s"
)

func GetCountry(countryId string) (*locations.Country, *errors.ApiError) {
  response := rest.Get(fmt.Sprintf(urlGetCountry, countryId))

  if response == nil || response.Response == nil {
    return nil, &errors.ApiError{
      Status: http.StatusInternalServerError,
      Message: fmt.Sprintf("invalid restclient response when trying to get country %s", countryId),
    }
  }

  if response.StatusCode > 299 {    // Code 299 is last number before error return values
    var apiErr errors.ApiError
    if err := json.Unmarshal(response.Bytes(), &apiErr); err != nil {
      return nil, &errors.ApiError{
        Status: http.StatusInternalServerError,
        Message: fmt.Sprintf("invalid error interface when getting country %s", countryId),
      }
    }
    return nil, &apiErr
  }

  var result locations.Country
  if err := json.Unmarshal(response.Bytes(), &result); err != nil {
    return nil, &errors.ApiError{
      Status: http.StatusInternalServerError,
      Message: fmt.Sprintf("error when trying to unmarshal country data for %s", countryId),
    }
  }

  return &result, nil
}
