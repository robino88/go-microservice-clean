package commercetools

type Customer struct {
	ID             string     `json:"id"`
	CustomerNumber string     `json:"customerNumber"`
	Key            string     `json:"key"`
	ExternalId     string     `json:"externalId"`
	Version        int        `json:"version"`
	CompanyName    string     `json:"companyName"`
	Email          string     `json:"email"`
	FirstName      string     `json:"firstName"`
	LastName       string     `json:"lastName"`
	Password       string     `json:"password"`
	Addresses      []*Address `json:"addresses,omitempty"`
	CustomerGroup  string     `json:"customerGroup"`
	Custom         string     `json:"custom"`
	Locale         string     `json:"locale"`
}
