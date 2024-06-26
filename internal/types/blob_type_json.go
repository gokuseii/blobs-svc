package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

func init() {
	// stubs for imports
	_ = json.Delim('s')
	_ = driver.Int32

}

var ErrBlobTypeInvalid = errors.New("BlobType is invalid")

func init() {
	var v BlobType
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_BlobTypeNameToValue = map[string]BlobType{
			interface{}(BlobTypeAssetDescription).(fmt.Stringer).String():   BlobTypeAssetDescription,
			interface{}(BlobTypeFundOverview).(fmt.Stringer).String():       BlobTypeFundOverview,
			interface{}(BlobTypeFundUpdate).(fmt.Stringer).String():         BlobTypeFundUpdate,
			interface{}(BlobTypeNavUpdate).(fmt.Stringer).String():          BlobTypeNavUpdate,
			interface{}(BlobTypeFundDocument).(fmt.Stringer).String():       BlobTypeFundDocument,
			interface{}(BlobTypeAlpha).(fmt.Stringer).String():              BlobTypeAlpha,
			interface{}(BlobTypeBravo).(fmt.Stringer).String():              BlobTypeBravo,
			interface{}(BlobTypeCharlie).(fmt.Stringer).String():            BlobTypeCharlie,
			interface{}(BlobTypeDelta).(fmt.Stringer).String():              BlobTypeDelta,
			interface{}(BlobTypeTokenTerms).(fmt.Stringer).String():         BlobTypeTokenTerms,
			interface{}(BlobTypeTokenMetrics).(fmt.Stringer).String():       BlobTypeTokenMetrics,
			interface{}(BlobTypeKYCForm).(fmt.Stringer).String():            BlobTypeKYCForm,
			interface{}(BlobTypeKYCIdDocument).(fmt.Stringer).String():      BlobTypeKYCIdDocument,
			interface{}(BlobTypeKYCPoa).(fmt.Stringer).String():             BlobTypeKYCPoa,
			interface{}(BlobTypeIdentityMindReject).(fmt.Stringer).String(): BlobTypeIdentityMindReject,
		}
	}
}

var _BlobTypeNameToValue = map[string]BlobType{
	"asset_description":    BlobTypeAssetDescription,
	"fund_overview":        BlobTypeFundOverview,
	"fund_update":          BlobTypeFundUpdate,
	"nav_update":           BlobTypeNavUpdate,
	"fund_document":        BlobTypeFundDocument,
	"alpha":                BlobTypeAlpha,
	"bravo":                BlobTypeBravo,
	"charlie":              BlobTypeCharlie,
	"delta":                BlobTypeDelta,
	"token_terms":          BlobTypeTokenTerms,
	"token_metrics":        BlobTypeTokenMetrics,
	"kyc_form":             BlobTypeKYCForm,
	"kyc_id_document":      BlobTypeKYCIdDocument,
	"kyc_poa":              BlobTypeKYCPoa,
	"identity_mind_reject": BlobTypeIdentityMindReject,
}

var _BlobTypeValueToName = map[BlobType]string{
	BlobTypeAssetDescription:   "asset_description",
	BlobTypeFundOverview:       "fund_overview",
	BlobTypeFundUpdate:         "fund_update",
	BlobTypeNavUpdate:          "nav_update",
	BlobTypeFundDocument:       "fund_document",
	BlobTypeAlpha:              "alpha",
	BlobTypeBravo:              "bravo",
	BlobTypeCharlie:            "charlie",
	BlobTypeDelta:              "delta",
	BlobTypeTokenTerms:         "token_terms",
	BlobTypeTokenMetrics:       "token_metrics",
	BlobTypeKYCForm:            "kyc_form",
	BlobTypeKYCIdDocument:      "kyc_id_document",
	BlobTypeKYCPoa:             "kyc_poa",
	BlobTypeIdentityMindReject: "identity_mind_reject",
}

// String is generated so BlobType satisfies fmt.Stringer.
func (r BlobType) String() string {
	s, ok := _BlobTypeValueToName[r]
	if !ok {
		return fmt.Sprintf("BlobType(%d)", r)
	}
	return s
}

// Validate verifies that value is predefined for BlobType.
func (r BlobType) Validate() error {
	_, ok := _BlobTypeValueToName[r]
	if !ok {
		return ErrBlobTypeInvalid
	}
	return nil
}

// MarshalJSON is generated so BlobType satisfies json.Marshaler.
func (r BlobType) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _BlobTypeValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid BlobType: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so BlobType satisfies json.Unmarshaler.
func (r *BlobType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("BlobType should be a string, got %s", data)
	}
	v, ok := _BlobTypeNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid BlobType %q", s)
	}
	*r = v
	return nil
}

func (t *BlobType) Scan(src interface{}) error {
	i, ok := src.(int64)
	if !ok {
		return fmt.Errorf("can't scan from %T", src)
	}
	*t = BlobType(i)
	return nil
}

func (t BlobType) Value() (driver.Value, error) {
	return int64(t), nil
}
