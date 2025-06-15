package dbstore

import (
	"github.com/itemslabs/clubz-api/database/schema"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// GetCountryByID fetches a single Country record from the database based on its ID.
func (s *DBStore) GetCountryByID(countryID string) (*schema.Country, error) {
	return schema.FindCountry(s.db, countryID)
}

// GetAllCountries fetches all Country records from the database.
func (s *DBStore) GetAllCountries() (schema.CountrySlice, error) {
	return schema.Countries().All(s.db)
}

// AddCountry inserts a new Country record into the database.
func (s *DBStore) AddCountry(country *schema.Country) error {
	return country.Insert(s.db, boil.Infer())
}

// UpdateCountry updates an existing Country record in the database.
func (s *DBStore) UpdateCountry(country *schema.Country) (int64, error) {
	return country.Update(s.db, boil.Infer())
}

// DeleteCountry deletes a Country record from the database.
func (s *DBStore) DeleteCountry(countryID string) (int64, error) {
	country, err := s.GetCountryByID(countryID)
	if err != nil {
		return 0, err
	}
	return country.Delete(s.db)
}

// CountryExistsByID checks if a Country with the given ID exists in the database.
func (s *DBStore) CountryExistsByID(countryID string) (bool, error) {
	return schema.CountryExists(s.db, countryID)
}

func (s *DBStore) GetCountryByName(name string) (*schema.Country, error) {
	return schema.Countries(schema.CountryWhere.Name.EQ(name)).One(s.db)
}
