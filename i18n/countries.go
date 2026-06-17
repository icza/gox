package i18n

import "unicode/utf8"

// Country holds some info about a country.
type Country struct {
	// 2-letter ISO 3166-1 alpha-2 country code
	// https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2
	Code string

	// 3-letter ISO 3166-1 alpha-3 country code
	// https://en.wikipedia.org/wiki/ISO_3166-1_alpha-3
	Code3 string

	// English name of the country
	Name string

	// Emoji flag
	Flag string

	// Flag + name
	FlagName string
}

// GetCode returns the 2-letter ISO 3166-1 alpha-2 country code.
// Empty string is returned if Country is nil.
func (c *Country) GetCode() string {
	if c == nil {
		return ""
	}
	return c.Code
}

// GetName returns the name of the country.
// Empty string is returned if Country is nil.
func (c *Country) GetName() string {
	if c == nil {
		return ""
	}
	return c.Name
}

// GetFlag returns the flag of the country.
// Empty string is returned if Country is nil.
func (c *Country) GetFlag() string {
	if c == nil {
		return ""
	}
	return c.Flag
}

// GetFlagName returns the flag+name of the country.
// Empty string is returned if Country is nil.
func (c *Country) GetFlagName() string {
	if c == nil {
		return ""
	}
	return c.FlagName
}

var Countries = []*Country{
	{Code: "AF", Code3: "AFG", FlagName: "🇦🇫Afghanistan"},
	{Code: "AX", Code3: "ALA", FlagName: "🇦🇽Åland Islands"},
	{Code: "AL", Code3: "ALB", FlagName: "🇦🇱Albania"},
	{Code: "DZ", Code3: "DZA", FlagName: "🇩🇿Algeria"},
	{Code: "AS", Code3: "ASM", FlagName: "🇦🇸American Samoa"},
	{Code: "AD", Code3: "AND", FlagName: "🇦🇩Andorra"},
	{Code: "AO", Code3: "AGO", FlagName: "🇦🇴Angola"},
	{Code: "AI", Code3: "AIA", FlagName: "🇦🇮Anguilla"},
	{Code: "AQ", Code3: "ATA", FlagName: "🇦🇶Antarctica"},
	{Code: "AG", Code3: "ATG", FlagName: "🇦🇬Antigua and Barbuda"},
	{Code: "AR", Code3: "ARG", FlagName: "🇦🇷Argentina"},
	{Code: "AM", Code3: "ARM", FlagName: "🇦🇲Armenia"},
	{Code: "AW", Code3: "ABW", FlagName: "🇦🇼Aruba"},
	{Code: "AU", Code3: "AUS", FlagName: "🇦🇺Australia"},
	{Code: "AT", Code3: "AUT", FlagName: "🇦🇹Austria"},
	{Code: "AZ", Code3: "AZE", FlagName: "🇦🇿Azerbaijan"},
	{Code: "BS", Code3: "BHS", FlagName: "🇧🇸Bahamas"},
	{Code: "BH", Code3: "BHR", FlagName: "🇧🇭Bahrain"},
	{Code: "BD", Code3: "BGD", FlagName: "🇧🇩Bangladesh"},
	{Code: "BB", Code3: "BRB", FlagName: "🇧🇧Barbados"},
	{Code: "BY", Code3: "BLR", FlagName: "🇧🇾Belarus"},
	{Code: "BE", Code3: "BEL", FlagName: "🇧🇪Belgium"},
	{Code: "BZ", Code3: "BLZ", FlagName: "🇧🇿Belize"},
	{Code: "BJ", Code3: "BEN", FlagName: "🇧🇯Benin"},
	{Code: "BM", Code3: "BMU", FlagName: "🇧🇲Bermuda"},
	{Code: "BT", Code3: "BTN", FlagName: "🇧🇹Bhutan"},
	{Code: "BO", Code3: "BOL", FlagName: "🇧🇴Bolivia (Plurinational State of)"}, // Previous ISO country name: Bolivia
	{Code: "BQ", Code3: "BES", FlagName: "🇧🇶Bonaire, Sint Eustatius and Saba"},
	{Code: "BA", Code3: "BIH", FlagName: "🇧🇦Bosnia and Herzegovina"},
	{Code: "BW", Code3: "BWA", FlagName: "🇧🇼Botswana"},
	{Code: "BV", Code3: "BVT", FlagName: "🇧🇻Bouvet Island"},
	{Code: "BR", Code3: "BRA", FlagName: "🇧🇷Brazil"},
	{Code: "IO", Code3: "IOT", FlagName: "🇮🇴British Indian Ocean Territory"},
	{Code: "BN", Code3: "BRN", FlagName: "🇧🇳Brunei Darussalam"},
	{Code: "BG", Code3: "BGR", FlagName: "🇧🇬Bulgaria"},
	{Code: "BF", Code3: "BFA", FlagName: "🇧🇫Burkina Faso"},
	{Code: "BI", Code3: "BDI", FlagName: "🇧🇮Burundi"},
	{Code: "KH", Code3: "KHM", FlagName: "🇰🇭Cambodia"},
	{Code: "CM", Code3: "CMR", FlagName: "🇨🇲Cameroon"},
	{Code: "CA", Code3: "CAN", FlagName: "🇨🇦Canada"},
	{Code: "CV", Code3: "CPV", FlagName: "🇨🇻Cabo Verde"}, // Common name and previous ISO country name: Cape Verde
	{Code: "KY", Code3: "CYM", FlagName: "🇰🇾Cayman Islands"},
	{Code: "CF", Code3: "CAF", FlagName: "🇨🇫Central African Republic"},
	{Code: "TD", Code3: "TCD", FlagName: "🇹🇩Chad"},
	{Code: "CL", Code3: "CHL", FlagName: "🇨🇱Chile"},
	{Code: "CN", Code3: "CHN", FlagName: "🇨🇳China"},
	{Code: "CX", Code3: "CXR", FlagName: "🇨🇽Christmas Island"},
	{Code: "CC", Code3: "CCK", FlagName: "🇨🇨Cocos (Keeling) Islands"},
	{Code: "CO", Code3: "COL", FlagName: "🇨🇴Colombia"},
	{Code: "KM", Code3: "COM", FlagName: "🇰🇲Comoros"},
	{Code: "CG", Code3: "COG", FlagName: "🇨🇬Congo"},
	{Code: "CD", Code3: "COD", FlagName: "🇨🇩Congo, Democratic Republic of the"},
	{Code: "CK", Code3: "COK", FlagName: "🇨🇰Cook Islands"},
	{Code: "CR", Code3: "CRI", FlagName: "🇨🇷Costa Rica"},
	{Code: "CI", Code3: "CIV", FlagName: "🇨🇮Côte d'Ivoire"},
	{Code: "HR", Code3: "HRV", FlagName: "🇭🇷Croatia"},
	{Code: "CU", Code3: "CUB", FlagName: "🇨🇺Cuba"},
	{Code: "CW", Code3: "CUW", FlagName: "🇨🇼Curaçao"},
	{Code: "CY", Code3: "CYP", FlagName: "🇨🇾Cyprus"},
	{Code: "CZ", Code3: "CZE", FlagName: "🇨🇿Czechia"}, // Previous ISO country name: Czech Republic
	{Code: "DK", Code3: "DNK", FlagName: "🇩🇰Denmark"},
	{Code: "DJ", Code3: "DJI", FlagName: "🇩🇯Djibouti"},
	{Code: "DM", Code3: "DMA", FlagName: "🇩🇲Dominica"},
	{Code: "DO", Code3: "DOM", FlagName: "🇩🇴Dominican Republic"},
	{Code: "EC", Code3: "ECU", FlagName: "🇪🇨Ecuador"},
	{Code: "EG", Code3: "EGY", FlagName: "🇪🇬Egypt"},
	{Code: "SV", Code3: "SLV", FlagName: "🇸🇻El Salvador"},
	{Code: "GQ", Code3: "GNQ", FlagName: "🇬🇶Equatorial Guinea"},
	{Code: "ER", Code3: "ERI", FlagName: "🇪🇷Eritrea"},
	{Code: "EE", Code3: "EST", FlagName: "🇪🇪Estonia"},
	{Code: "SZ", Code3: "SWZ", FlagName: "🇸🇿Eswatini"}, // Previous ISO country name: Swaziland
	{Code: "ET", Code3: "ETH", FlagName: "🇪🇹Ethiopia"},
	{Code: "FK", Code3: "FLK", FlagName: "🇫🇰Falkland Islands (Malvinas)"},
	{Code: "FO", Code3: "FRO", FlagName: "🇫🇴Faroe Islands"},
	{Code: "FJ", Code3: "FJI", FlagName: "🇫🇯Fiji"},
	{Code: "FI", Code3: "FIN", FlagName: "🇫🇮Finland"},
	{Code: "FR", Code3: "FRA", FlagName: "🇫🇷France"},
	{Code: "GF", Code3: "GUF", FlagName: "🇬🇫French Guiana"},
	{Code: "PF", Code3: "PYF", FlagName: "🇵🇫French Polynesia"},
	{Code: "TF", Code3: "ATF", FlagName: "🇹🇫French Southern Territories"},
	{Code: "GA", Code3: "GAB", FlagName: "🇬🇦Gabon"},
	{Code: "GM", Code3: "GMB", FlagName: "🇬🇲Gambia"},
	{Code: "GE", Code3: "GEO", FlagName: "🇬🇪Georgia"},
	{Code: "DE", Code3: "DEU", FlagName: "🇩🇪Germany"},
	{Code: "GH", Code3: "GHA", FlagName: "🇬🇭Ghana"},
	{Code: "GI", Code3: "GIB", FlagName: "🇬🇮Gibraltar"},
	{Code: "GR", Code3: "GRC", FlagName: "🇬🇷Greece"},
	{Code: "GL", Code3: "GRL", FlagName: "🇬🇱Greenland"},
	{Code: "GD", Code3: "GRD", FlagName: "🇬🇩Grenada"},
	{Code: "GP", Code3: "GLP", FlagName: "🇬🇵Guadeloupe"},
	{Code: "GU", Code3: "GUM", FlagName: "🇬🇺Guam"},
	{Code: "GT", Code3: "GTM", FlagName: "🇬🇹Guatemala"},
	{Code: "GG", Code3: "GGY", FlagName: "🇬🇬Guernsey"},
	{Code: "GN", Code3: "GIN", FlagName: "🇬🇳Guinea"},
	{Code: "GW", Code3: "GNB", FlagName: "🇬🇼Guinea-Bissau"},
	{Code: "GY", Code3: "GUY", FlagName: "🇬🇾Guyana"},
	{Code: "HT", Code3: "HTI", FlagName: "🇭🇹Haiti"},
	{Code: "HM", Code3: "HMD", FlagName: "🇭🇲Heard Island and McDonald Islands"},
	{Code: "VA", Code3: "VAT", FlagName: "🇻🇦Holy See"}, // Previous ISO country names: Vatican City State (Holy See) and Holy See (Vatican City State)
	{Code: "HN", Code3: "HND", FlagName: "🇭🇳Honduras"},
	{Code: "HK", Code3: "HKG", FlagName: "🇭🇰Hong Kong"},
	{Code: "HU", Code3: "HUN", FlagName: "🇭🇺Hungary"},
	{Code: "IS", Code3: "ISL", FlagName: "🇮🇸Iceland"},
	{Code: "IN", Code3: "IND", FlagName: "🇮🇳India"},
	{Code: "ID", Code3: "IDN", FlagName: "🇮🇩Indonesia"},
	{Code: "IR", Code3: "IRN", FlagName: "🇮🇷Iran (Islamic Republic of)"}, // Previous ISO country name: Iran
	{Code: "IQ", Code3: "IRQ", FlagName: "🇮🇶Iraq"},
	{Code: "IE", Code3: "IRL", FlagName: "🇮🇪Ireland"},
	{Code: "IM", Code3: "IMN", FlagName: "🇮🇲Isle of Man"},
	{Code: "IL", Code3: "ISR", FlagName: "🇮🇱Israel"},
	{Code: "IT", Code3: "ITA", FlagName: "🇮🇹Italy"},
	{Code: "JM", Code3: "JAM", FlagName: "🇯🇲Jamaica"},
	{Code: "JP", Code3: "JPN", FlagName: "🇯🇵Japan"},
	{Code: "JE", Code3: "JEY", FlagName: "🇯🇪Jersey"},
	{Code: "JO", Code3: "JOR", FlagName: "🇯🇴Jordan"},
	{Code: "KZ", Code3: "KAZ", FlagName: "🇰🇿Kazakhstan"},
	{Code: "KE", Code3: "KEN", FlagName: "🇰🇪Kenya"},
	{Code: "KI", Code3: "KIR", FlagName: "🇰🇮Kiribati"},
	{Code: "KP", Code3: "PRK", FlagName: "🇰🇵Korea (Democratic People's Republic of)"},
	{Code: "KR", Code3: "KOR", FlagName: "🇰🇷Korea, Republic of"},
	{Code: "XK", Code3: "XKX", FlagName: "🇽🇰Kosovo"},
	{Code: "KW", Code3: "KWT", FlagName: "🇰🇼Kuwait"},
	{Code: "KG", Code3: "KGZ", FlagName: "🇰🇬Kyrgyzstan"},
	{Code: "LA", Code3: "LAO", FlagName: "🇱🇦Lao People's Democratic Republic"},
	{Code: "LV", Code3: "LVA", FlagName: "🇱🇻Latvia"},
	{Code: "LB", Code3: "LBN", FlagName: "🇱🇧Lebanon"},
	{Code: "LS", Code3: "LSO", FlagName: "🇱🇸Lesotho"},
	{Code: "LR", Code3: "LBR", FlagName: "🇱🇷Liberia"},
	{Code: "LY", Code3: "LBY", FlagName: "🇱🇾Libya"},
	{Code: "LI", Code3: "LIE", FlagName: "🇱🇮Liechtenstein"},
	{Code: "LT", Code3: "LTU", FlagName: "🇱🇹Lithuania"},
	{Code: "LU", Code3: "LUX", FlagName: "🇱🇺Luxembourg"},
	{Code: "MO", Code3: "MAC", FlagName: "🇲🇴Macao"},
	{Code: "MG", Code3: "MDG", FlagName: "🇲🇬Madagascar"},
	{Code: "MW", Code3: "MWI", FlagName: "🇲🇼Malawi"},
	{Code: "MY", Code3: "MYS", FlagName: "🇲🇾Malaysia"},
	{Code: "MV", Code3: "MDV", FlagName: "🇲🇻Maldives"},
	{Code: "ML", Code3: "MLI", FlagName: "🇲🇱Mali"},
	{Code: "MT", Code3: "MLT", FlagName: "🇲🇹Malta"},
	{Code: "MH", Code3: "MHL", FlagName: "🇲🇭Marshall Islands"},
	{Code: "MQ", Code3: "MTQ", FlagName: "🇲🇶Martinique"},
	{Code: "MR", Code3: "MRT", FlagName: "🇲🇷Mauritania"},
	{Code: "MU", Code3: "MUS", FlagName: "🇲🇺Mauritius"},
	{Code: "YT", Code3: "MYT", FlagName: "🇾🇹Mayotte"},
	{Code: "MX", Code3: "MEX", FlagName: "🇲🇽Mexico"},
	{Code: "FM", Code3: "FSM", FlagName: "🇫🇲Micronesia (Federated States of)"},
	{Code: "MD", Code3: "MDA", FlagName: "🇲🇩Moldova, Republic of"},
	{Code: "MC", Code3: "MCO", FlagName: "🇲🇨Monaco"},
	{Code: "MN", Code3: "MNG", FlagName: "🇲🇳Mongolia"},
	{Code: "ME", Code3: "MNE", FlagName: "🇲🇪Montenegro"},
	{Code: "MS", Code3: "MSR", FlagName: "🇲🇸Montserrat"},
	{Code: "MA", Code3: "MAR", FlagName: "🇲🇦Morocco"},
	{Code: "MZ", Code3: "MOZ", FlagName: "🇲🇿Mozambique"},
	{Code: "MM", Code3: "MMR", FlagName: "🇲🇲Myanmar"},
	{Code: "NA", Code3: "NAM", FlagName: "🇳🇦Namibia"},
	{Code: "NR", Code3: "NRU", FlagName: "🇳🇷Nauru"},
	{Code: "NP", Code3: "NPL", FlagName: "🇳🇵Nepal"},
	{Code: "NL", Code3: "NLD", FlagName: "🇳🇱Netherlands, Kingdom of the"},
	{Code: "NC", Code3: "NCL", FlagName: "🇳🇨New Caledonia"},
	{Code: "NZ", Code3: "NZL", FlagName: "🇳🇿New Zealand"},
	{Code: "NI", Code3: "NIC", FlagName: "🇳🇮Nicaragua"},
	{Code: "NE", Code3: "NER", FlagName: "🇳🇪Niger"},
	{Code: "NG", Code3: "NGA", FlagName: "🇳🇬Nigeria"},
	{Code: "NU", Code3: "NIU", FlagName: "🇳🇺Niue"},
	{Code: "NF", Code3: "NFK", FlagName: "🇳🇫Norfolk Island"},
	{Code: "MK", Code3: "MKD", FlagName: "🇲🇰North Macedonia"}, // Previous ISO country name: Macedonia, the former Yugoslav Republic of
	{Code: "MP", Code3: "MNP", FlagName: "🇲🇵Northern Mariana Islands"},
	{Code: "NO", Code3: "NOR", FlagName: "🇳🇴Norway"},
	{Code: "OM", Code3: "OMN", FlagName: "🇴🇲Oman"},
	{Code: "PK", Code3: "PAK", FlagName: "🇵🇰Pakistan"},
	{Code: "PW", Code3: "PLW", FlagName: "🇵🇼Palau"},
	{Code: "PS", Code3: "PSE", FlagName: "🇵🇸Palestine, State of"},
	{Code: "PA", Code3: "PAN", FlagName: "🇵🇦Panama"},
	{Code: "PG", Code3: "PNG", FlagName: "🇵🇬Papua New Guinea"},
	{Code: "PY", Code3: "PRY", FlagName: "🇵🇾Paraguay"},
	{Code: "PE", Code3: "PER", FlagName: "🇵🇪Peru"},
	{Code: "PH", Code3: "PHL", FlagName: "🇵🇭Philippines"},
	{Code: "PN", Code3: "PCN", FlagName: "🇵🇳Pitcairn"},
	{Code: "PL", Code3: "POL", FlagName: "🇵🇱Poland"},
	{Code: "PT", Code3: "PRT", FlagName: "🇵🇹Portugal"},
	{Code: "PR", Code3: "PRI", FlagName: "🇵🇷Puerto Rico"},
	{Code: "QA", Code3: "QAT", FlagName: "🇶🇦Qatar"},
	{Code: "RE", Code3: "REU", FlagName: "🇷🇪Réunion"},
	{Code: "RO", Code3: "ROU", FlagName: "🇷🇴Romania"},
	{Code: "RU", Code3: "RUS", FlagName: "🇷🇺Russian Federation"},
	{Code: "RW", Code3: "RWA", FlagName: "🇷🇼Rwanda"},
	{Code: "BL", Code3: "BLM", FlagName: "🇧🇱Saint Barthélemy"},
	{Code: "SH", Code3: "SHN", FlagName: "🇸🇭Saint Helena, Ascension and Tristan da Cunha"},
	{Code: "KN", Code3: "KNA", FlagName: "🇰🇳Saint Kitts and Nevis"},
	{Code: "LC", Code3: "LCA", FlagName: "🇱🇨Saint Lucia"},
	{Code: "MF", Code3: "MAF", FlagName: "🇲🇫Saint Martin (French part)"},
	{Code: "PM", Code3: "SPM", FlagName: "🇵🇲Saint Pierre and Miquelon"},
	{Code: "VC", Code3: "VCT", FlagName: "🇻🇨Saint Vincent and the Grenadines"},
	{Code: "WS", Code3: "WSM", FlagName: "🇼🇸Samoa"},
	{Code: "SM", Code3: "SMR", FlagName: "🇸🇲San Marino"},
	{Code: "ST", Code3: "STP", FlagName: "🇸🇹Sao Tome and Principe"},
	{Code: "SA", Code3: "SAU", FlagName: "🇸🇦Saudi Arabia"},
	{Code: "SN", Code3: "SEN", FlagName: "🇸🇳Senegal"},
	{Code: "RS", Code3: "SRB", FlagName: "🇷🇸Serbia"},
	{Code: "SC", Code3: "SYC", FlagName: "🇸🇨Seychelles"},
	{Code: "SL", Code3: "SLE", FlagName: "🇸🇱Sierra Leone"},
	{Code: "SG", Code3: "SGP", FlagName: "🇸🇬Singapore"},
	{Code: "SX", Code3: "SXM", FlagName: "🇸🇽Sint Maarten (Dutch part)"},
	{Code: "SK", Code3: "SVK", FlagName: "🇸🇰Slovakia"},
	{Code: "SI", Code3: "SVN", FlagName: "🇸🇮Slovenia"},
	{Code: "SB", Code3: "SLB", FlagName: "🇸🇧Solomon Islands"},
	{Code: "SO", Code3: "SOM", FlagName: "🇸🇴Somalia"},
	{Code: "ZA", Code3: "ZAF", FlagName: "🇿🇦South Africa"},
	{Code: "GS", Code3: "SGS", FlagName: "🇬🇸South Georgia and the South Sandwich Islands"},
	{Code: "SS", Code3: "SSD", FlagName: "🇸🇸South Sudan"},
	{Code: "ES", Code3: "ESP", FlagName: "🇪🇸Spain"},
	{Code: "LK", Code3: "LKA", FlagName: "🇱🇰Sri Lanka"},
	{Code: "SD", Code3: "SDN", FlagName: "🇸🇩Sudan"},
	{Code: "SR", Code3: "SUR", FlagName: "🇸🇷Suriname"},
	{Code: "SJ", Code3: "SJM", FlagName: "🇸🇯Svalbard and Jan Mayen"},
	{Code: "SE", Code3: "SWE", FlagName: "🇸🇪Sweden"},
	{Code: "CH", Code3: "CHE", FlagName: "🇨🇭Switzerland"},
	{Code: "SY", Code3: "SYR", FlagName: "🇸🇾Syrian Arab Republic"},
	{Code: "TW", Code3: "TWN", FlagName: "🇹🇼Taiwan, Province of China"},
	{Code: "TJ", Code3: "TJK", FlagName: "🇹🇯Tajikistan"},
	{Code: "TZ", Code3: "TZA", FlagName: "🇹🇿Tanzania, United Republic of"},
	{Code: "TH", Code3: "THA", FlagName: "🇹🇭Thailand"},
	{Code: "TL", Code3: "TLS", FlagName: "🇹🇱Timor-Leste"},
	{Code: "TG", Code3: "TGO", FlagName: "🇹🇬Togo"},
	{Code: "TK", Code3: "TKL", FlagName: "🇹🇰Tokelau"},
	{Code: "TO", Code3: "TON", FlagName: "🇹🇴Tonga"},
	{Code: "TT", Code3: "TTO", FlagName: "🇹🇹Trinidad and Tobago"},
	{Code: "TN", Code3: "TUN", FlagName: "🇹🇳Tunisia"},
	{Code: "TR", Code3: "TUR", FlagName: "🇹🇷Türkiye"}, // Previous ISO country name: Turkey
	{Code: "TM", Code3: "TKM", FlagName: "🇹🇲Turkmenistan"},
	{Code: "TC", Code3: "TCA", FlagName: "🇹🇨Turks and Caicos Islands"},
	{Code: "TV", Code3: "TUV", FlagName: "🇹🇻Tuvalu"},
	{Code: "UG", Code3: "UGA", FlagName: "🇺🇬Uganda"},
	{Code: "UA", Code3: "UKR", FlagName: "🇺🇦Ukraine"},
	{Code: "AE", Code3: "ARE", FlagName: "🇦🇪United Arab Emirates"},
	{Code: "GB", Code3: "GBR", FlagName: "🇬🇧United Kingdom of Great Britain and Northern Ireland"}, // Previous ISO country name: United Kingdom
	{Code: "UM", Code3: "UMI", FlagName: "🇺🇲United States Minor Outlying Islands"},
	{Code: "US", Code3: "USA", FlagName: "🇺🇸United States of America"}, // Previous ISO country name: United States
	{Code: "UY", Code3: "URY", FlagName: "🇺🇾Uruguay"},
	{Code: "UZ", Code3: "UZB", FlagName: "🇺🇿Uzbekistan"},
	{Code: "VU", Code3: "VUT", FlagName: "🇻🇺Vanuatu"},
	{Code: "VE", Code3: "VEN", FlagName: "🇻🇪Venezuela (Bolivarian Republic of)"},
	{Code: "VN", Code3: "VNM", FlagName: "🇻🇳Viet Nam"},
	{Code: "VG", Code3: "VGB", FlagName: "🇻🇬Virgin Islands (British)"},
	{Code: "VI", Code3: "VIR", FlagName: "🇻🇮Virgin Islands (U.S.)"},
	{Code: "WF", Code3: "WLF", FlagName: "🇼🇫Wallis and Futuna"},
	{Code: "EH", Code3: "ESH", FlagName: "🇪🇭Western Sahara"},
	{Code: "YE", Code3: "YEM", FlagName: "🇾🇪Yemen"},
	{Code: "ZM", Code3: "ZMB", FlagName: "🇿🇲Zambia"},
	{Code: "ZW", Code3: "ZWE", FlagName: "🇿🇼Zimbabwe"},
}

// CountryCodeCountries maps from the 2-letter ISO 3166-1 alpha-2 country code to the Country descriptor.
var CountryCodeCountries = make(map[string]*Country, len(Countries))

// Initialize Countries and the CountryCodeCountries map.
func init() {
	// Country literals only contain FlagName as it's the concatenation of Flag and Name.
	// Slice FlagName and store the Flag and Name for each country.
	// Slicing reuses the memory of the FlagName string.
	for _, country := range Countries {
		// Flag emojies are 2 runes, so slice the first 2 runes from FlagName to get the Flag.
		_, size1 := utf8.DecodeRuneInString(country.FlagName)
		_, size2 := utf8.DecodeRuneInString(country.FlagName[size1:])
		flagLen := size1 + size2
		country.Flag = country.FlagName[:flagLen]
		country.Name = country.FlagName[flagLen:]

		CountryCodeCountries[country.Code] = country
	}
}
