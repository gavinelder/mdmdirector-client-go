package types

import (
	"github.com/google/uuid"
)

// DeviceProfile (s) are profiles that are individual to the device.
type DeviceProfile struct {
	// ID                uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	PayloadUUID       string `json:"PayloadUUID"`
	PayloadIdentifier string `json:"PayloadIdentifier" gorm:"primaryKey"`
	HashedPayloadUUID string `json:"HashedPayloadUUID"`
	MobileconfigData  []byte `json:"MobileconfigData"`
	MobileconfigHash  []byte `json:"MobileconfigHash"`
	DeviceUDID        string `json:"DeviceUDID" gorm:"primaryKey"`
	Installed         bool   `json:"Installed" gorm:"default:true"`
}

// SharedProfile (s) are profiles that go on every device.
type SharedProfile struct {
	ID                uuid.UUID `json:"ID" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	PayloadUUID       string    `json:"PayloadUUID"`
	HashedPayloadUUID string    `json:"HashedPayloadUUID"`
	PayloadIdentifier string    `json:"PayloadIdentifier"`
	MobileconfigData  []byte    `json:"MobileconfigData"`
	MobileconfigHash  []byte    `json:"MobileconfigHash"`
	Installed         bool      `json:"Installed" gorm:"default:true"`
}

// ProfilePayload - struct to unpack the payload sent to mdmdirector
/*
type ProfilePayload struct {
	SerialNumbers []string `json:"serial_numbers,omitempty"`
	DeviceUDIDs   []string `json:"udids,omitempty"`
	Mobileconfigs []string `json:"profiles"`
	PushNow       bool     `json:"push_now"`
	Metadata      bool     `json:"metadata"`
}

/*
type DeleteProfilePayload struct {
	SerialNumbers []string                     `json:"serial_numbers,omitempty"`
	DeviceUDIDs   []string                     `json:"udids,omitempty"`
	PushNow       bool                         `json:"push_now"`
	Mobileconfigs []DeletedMobileconfigPayload `json:"profiles"`
	Metadata      bool                         `json:"metadata"`
}

type DeletedMobileconfigPayload struct {
	PayloadIdentifier string `json:"payload_identifier"`
}

// ProfileList - returned data from the ProfileList MDM command
type ProfileListData struct {
	ProfileList []ProfileList
}

type ProfileList struct {
	ID                       uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	DeviceUDID               string
	HasRemovalPasscode       bool          `plist:"HasRemovalPasscode"`
	IsEncrypted              bool          `plist:"IsEncrypted"`
	IsManaged                bool          `plist:"IsManaged"`
	PayloadContent           []interface{} `plist:"PaylodContent" gorm:"-"`
	PayloadDescription       string        `plist:"PayloadDescription"`
	PayloadDisplayName       string        `plist:"PayloadDisplayName"`
	PayloadIdentifier        string        `plist:"PayloadIdentifier"`
	PayloadOrganization      string        `plist:"PayloadOrganization"`
	PayloadRemovalDisallowed bool          `plist:"PayloadRemovalDisallowed"`
	PayloadUUID              string        `plist:"PayloadUUID" gorm:"not null"`
	PayloadVersion           int           `plist:"PayloadVersion"`
	FullPayload              bool          `plist:"FullPayload"`
}

type MetadataItem struct {
	// Device          Device            `json:"device"`
	ProfileMetadata []ProfileMetadata `json:"profile_metadata"`
}

type ProfileMetadata struct {
	Status            string `json:"status"`
	PayloadIdentifier string `json:"payload_identifier"`
	PayloadUUID       string `json:"payload_uuid"`
	HashedPayloadUUID string `json:"hashed_payload_uuid"`
}

func (profile *DeviceProfile) AfterCreate(tx *gorm.DB) (err error) {
	BumpDeviceLastUpdated(profile.DeviceUDID)
	return nil
}

func (profile *DeviceProfile) AfterUpdate(tx *gorm.DB) (err error) {
	BumpDeviceLastUpdated(profile.DeviceUDID)
	return nil
}
*/
