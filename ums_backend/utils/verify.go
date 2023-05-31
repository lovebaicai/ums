package utils

var (
	RegisterVerify = Rules{"Username": {NotEmpty()}, "CNname": {NotEmpty()}}
	PageInfoVerify = Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}}
	IdVerify       = Rules{"ID": []string{NotEmpty()}}
	LoginVerify    = Rules{"Username": {NotEmpty()}, "Password": {NotEmpty()}}
	//ApiVerify              = Rules{"Path": {NotEmpty()}, "Description": {NotEmpty()}, "ApiGroup": {NotEmpty()}, "Method": {NotEmpty()}}
	//MenuVerify             = Rules{"Path": {NotEmpty()}, "ParentId": {NotEmpty()}, "Name": {NotEmpty()}, "Component": {NotEmpty()}, "Sort": {Ge("0")}}
	//MenuMetaVerify         = Rules{"Title": {NotEmpty()}}
	//LoginVerify            = Rules{"CaptchaId": {NotEmpty()}, "Captcha": {NotEmpty()}, "Username": {NotEmpty()}, "Password": {NotEmpty()}}
	//CustomerVerify         = Rules{"CustomerName": {NotEmpty()}, "CustomerPhoneData": {NotEmpty()}}
	//AutoCodeVerify         = Rules{"Abbreviation": {NotEmpty()}, "StructName": {NotEmpty()}, "PackageName": {NotEmpty()}, "Fields": {NotEmpty()}}
	//AutoPackageVerify      = Rules{"PackageName": {NotEmpty()}}
	//AuthorityVerify        = Rules{"AuthorityId": {NotEmpty()}, "AuthorityName": {NotEmpty()}}
	//AuthorityIdVerify      = Rules{"AuthorityId": {NotEmpty()}}
	//OldAuthorityVerify     = Rules{"OldAuthorityId": {NotEmpty()}}
	//ChangePasswordVerify   = Rules{"Password": {NotEmpty()}, "NewPassword": {NotEmpty()}}
	//SetUserAuthorityVerify = Rules{"AuthorityId": {NotEmpty()}}
)
