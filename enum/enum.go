package enum

// UserStatus represents the status of a user account
type VenderStatus string

const (
	Active    VenderStatus = "ACTIVE"
	Inactive  VenderStatus = "INACTIVE"
	Suspended VenderStatus = "SUSPENDED"
)

type VenderData string

const (
	UnVerified    VenderData = "UNVERIFIED"
	Verified      VenderData = "VERIFIED"
	ReField       VenderData = "REFIELD"
)

type VenderType string

const (
	Indivisual    VenderType = "INDIVISUAL"
	Clinic        VenderType = "CLINIC"
	Hospital      VenderType = "HOSPITAL"
)

type ProductType string

const (
	Quick    	ProductType = "QUICK"
    Schedule    ProductType = "SCHEDULE"
	Emergency   ProductType = "EMERGENCY"
)