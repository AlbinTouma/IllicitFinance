package db_models
import "html/template"


type Person struct {
  Id  string  `json:"id" human:"Id"`
  Name string `json:"name" human:"Name"`
  FirstName string  `json:"first_name" human:"First Name"`
  MiddleName string `json:"middle_name" human:"Middle Name"`
  LastName  string  `json:"last_name" human:"Last Name"`
  NameSuffix  string  `json:"name_suffix" human:"Suffix"`
  FatherName string `json:"father_name" human:"Father's name"`
  MotherName string `json:"mother_name" human:"Mother's name"`
  BirthDate string  `json:"birth_date"  human:"Date of Birth"`
  DeathDate  string `json:"death_date"  human:"Date of Death"`
  Gender  string  `json:"gender"  human:"Gender"`
  Ethnicity string  `json:"ethnicity" human:"Ethnicity"`
  Height  string  `json:"height"  human:"Height"`
  Weight  string  `json:"weight"  human:"Weight"`
  EyeColor string `json:"eye_color" human:"Eye Color"`
  HairColor  string `json:"hair_color"  human:"Hair Color"`
  Appearance  string  `json:"appearance"  human:"Appearance"`
  Religion  string  `json:"religion"  human:"Religion"`
  SpokenLanguage  string `json:"spoken_language"  human:"Spoken Languages"`
  Education string  `json:"education" human:"Eduction"`
  Address *Address
  Nationality string  `json:"nationality" human:"Nationality"`
  Citizenship string  `json:"citizenship" human:"Citizenship"`
  BirthPlace string `json:"birth_place" human:"Place of Birth"`
  BirthCountry string `json:"birth_country" human:"Country of Birth"`
  Position  string  `json:"position"  human:"Position"`
  Email string  `json:"email" human:"Email"`
  Phone string  `json:"phone" human:"Phone"`
  PassportNumber string `json:"passport_number" human:"Passport Number"`
  SocialSecurityNumber  string `json:"social_security_number"  human:"Social Security Number"`
  SourceName string `json:"source_name" human:"Source"`
  SourceUrl string    `json:"source_url"  human:"Source Link"`
  Note  template.HTML  `json:"note" human:"Description"`
  Image string  `json:"image" human:"Image"`
  Created string `json:"created" human:"Created"`
  LastUpdated string  `json:"updated" human:"Updated"`
}

type Address struct {
  FullAddress string  `json:"address_full"  human:"FullAddress"`
  AddressStreet string `json:"address_street" human:"street"`
  AddressStreet2  string  `json:"address_street2" human:street2"`
  City string `json:"address_city" human:"city"`
  PostalCode string  `json:"address_postal_code" human:Postal Code`
  Region string `json:"address_region"  human:"Region"`
  Country string `json:"address_country" human:"Country"`
  PostBox string `json:"address_post_office_box" human:"Post Box"`

}


