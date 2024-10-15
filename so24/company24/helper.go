package company24

func (c *companyServiceSoap) GetCompaniesByIds(
	ids []int32,
	searchParams *CompanySearchParameters,
	returnProps *ArrayOfString,
) (*GetCompaniesResponse, error) {
	searchParams.CompanyIds.int = ids
	return c.GetCompanies(&GetCompanies{SearchParams: searchParams, ReturnProperties: returnProps})
}
