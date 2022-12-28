package models

type Reference struct {
	ID                 int    `json:"id"`
	FullName           string `json:"full_name"`
	PhoneNumber        string `json:"phone_number"`
	Email              string `json:"email"`
	BirthDate          string `json:"birth_date"`
	ResidentialAddress string `json:"residential_address"`
	Inn                string `json:"inn"`
	PassportFront      string `json:"passport_front"`
	PassportBack       string `json:"passport_back"`
	PassportSelfie     string `json:"passport_selfie"`
	PaymentReceipt     string `json:"payment_receipt"`
	ReferenceLanguage  string `json:"reference_language"`
	ReferenceTariff    string `json:"reference_tariff"`
	Status             string `json:"status"`
	StatusType         string `json:"status_type"`
	Comment            string `json:"comment"`
	CreatedAt          string `json:"created_at"`
	ReceivingRegion    string `json:"receiving_region"`
}

type ReferenceStatus struct {
	Status  string `json:"status"`
	Comment string `json:"comment"`
}
