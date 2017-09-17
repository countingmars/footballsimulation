package base

var Countries = struct{
	KR,
	UK, ES, IT, FR, DE Country
} {
	KR : Country{"KR", "Korea"},
	UK : Country{"UK", "United Kingdom"},
	ES : Country{"ES", "Spain"},
	IT : Country{"IT", "Italia"},
	FR : Country{"FR", "France"},
	DE : Country{"DE", "Germany"},
}
type Country struct {
	Code string `json:"code"`
	Name string `json:"name"`
}
