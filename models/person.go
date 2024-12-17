package db_models

type Person struct {
  Name Name `json:"name"`
  BirthDate []string `json:"birth_date"`
  Country Country `json:"country"`
  Nationality Country `json:"nationality"`
  LegalForm LegalForm `json:"legal_form"`
  Status Status `json:"status"`
  Address Address `json:"address"`
  Alias Alias `json:"alias"`
  Appearance string `json:"appearance"`
  //AddressEnty
  BirthCountry Country  `json:"birth_country"`
  BirthPlace  Country `json:"birthPlace"`
  Citizenship Country `json:"citizenship"`
  Classification  string  `json:"classification"`
  CreatedAt CreatedAt `json:"created_at"`
  DeathDate Date  `json:"death_date"` 
  Notes []string `json:"notes"`
  //DissolutionDate Date  `json:"dissolution_date"`
}

