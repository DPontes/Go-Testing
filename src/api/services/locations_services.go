package services

import (
  "api/domain/locations"
  "api/utils/errors"
  "api/providers/locations_provider"
)

func GetCountry(countryId string) (*locations.Country, *errors.ApiError) {
  return locations_provider.GetCountry(countryId)
}
