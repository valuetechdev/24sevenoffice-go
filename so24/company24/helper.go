package company24

func GetCompaniesByIds(
	c CompanyServiceSoap,
	ids []int32,
	searchParams *CompanySearchParameters,
	returnProps *ArrayOfString,
) (*GetCompaniesResponse, error) {
	searchParams.CompanyIds.int = ids
	return c.GetCompanies(&GetCompanies{SearchParams: searchParams, ReturnProperties: returnProps})
}
