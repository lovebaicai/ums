package utils

var (
	RegisterVerify = Rules{"Username": {NotEmpty()}, "CNname": {NotEmpty()}}
	PageInfoVerify = Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}}
	IdVerify       = Rules{"ID": []string{NotEmpty()}}
	LoginVerify    = Rules{"Username": {NotEmpty()}, "Password": {NotEmpty()}}
)
