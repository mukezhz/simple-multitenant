package database

type MultitenantDomain struct {
	TenantID string `json:"tenant_id"`
	Domain   string `json:"domain"`
}

type TenantInformation struct {
	TenantID string `json:"tenant_id"`
	Detail   string `json:"detail"`
}

type Database struct {
	Multitenants       []MultitenantDomain
	TenantInformations []TenantInformation
}

func NewDatabase() *Database {
	return &Database{
		InitializeMultitenantDomain(),
		InitializeMultitenantInfomation(),
	}
}

func (s *Database) FindTenantIDByDomain(domain string) string {
	for _, s := range s.Multitenants {
		if s.Domain == domain {
			return s.TenantID
		}
	}
	return ""
}

func (s *Database) FindDetailByTenantID(tenantID string) string {
	for _, s := range s.TenantInformations {
		if s.TenantID == tenantID {
			return s.Detail
		}
	}
	return ""
}
