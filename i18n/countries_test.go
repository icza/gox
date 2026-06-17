package i18n

import "testing"

func TestCountries(t *testing.T) {
	if len(Countries) != len(CountryCodeCountries) {
		t.Errorf("Mismatch in number of countries: %d vs %d", len(Countries), len(CountryCodeCountries))
	}

	for _, country := range Countries {
		if country.GetCode() == "" {
			t.Errorf("Country %v has empty code", country)
		}
		if country.GetFlag() == "" {
			t.Errorf("Country %v has empty flag", country)
		}
		if country.GetFlagName() == "" {
			t.Errorf("Country %v has empty flag name", country)
		}
	}
}
