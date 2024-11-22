package database

func InitializeMultitenantDomain() []MultitenantDomain {
	return []MultitenantDomain{
		{
			TenantID: "12345",
			Domain:   "localhost:8000",
		},
		{
			TenantID: "11223",
			Domain:   "multitenant.com",
		},
	}
}

func InitializeMultitenantInfomation() []TenantInformation {
	return []TenantInformation{
		{
			TenantID: "12345",
			Detail:   "I am localhost",
		},
		{
			TenantID: "11223",
			Detail:   "I am multitenant",
		},
	}
}
