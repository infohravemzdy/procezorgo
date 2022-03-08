package types

const (
	versionZeroCode int32 = 0
)

type VersionCode struct {
	value int32
}

func (v VersionCode) Value() int32 {
	return v.value
}

func (v VersionCode) IsZero() bool {
	return v.value == versionZeroCode
}

func VersionCodeZero() VersionCode {
	return newVersionCode()
}

func newVersionCode() VersionCode {
	return VersionCode{	value: versionZeroCode}
}

func GetVersionCode(id int32) VersionCode {
	return VersionCode{	value: id }
}

