package db_models

type Person struct {
  Name Name `json:"name" human:"Name"`
  BirthDate []string `json:"birth_date" human:"Date of Birth"`
  Country Country `json:"country" human:"Country"`
  Nationality Country `json:"nationality" human:"Nationality"`
  LegalForm LegalForm `json:"legal_form" human:"Legal Form"`
  Status Status `json:"status" human:"Status"`
  Address Address `json:"address" human:"Address"`
  Alias Alias `json:"alias" human:"Alias"`
  Appearance []string `json:"appearance" human:"Appearance"`
  Height Height `json:"height" human:"Height"`
  Gender Gender `json:"gender" human:"Gender"`
  BirthCountry Country  `json:"birth_country" human:"Country of Birth"`
  BirthPlace  Country `json:"birthPlace" human:"Birth Place"`
  Citizenship Country `json:"citizenship" human:"Citizenship"`
  Classification  string  `json:"classification" human:"Classification"`
  CreatedAt CreatedAt `json:"created_at" human:"Created at"`
  DeathDate Date  `json:"death_date" human:"Date of Death"` 
  Notes []string `json:"notes"`
  SourceUrl []string  `json:"SourceUrl" human:"Providence"`
  //DissolutionDate Date  `json:"dissolution_date"`
}

