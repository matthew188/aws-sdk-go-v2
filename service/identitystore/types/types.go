// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"github.com/matthew188/aws-sdk-go-v2/service/identitystore/document"
	smithydocument "github.com/aws/smithy-go/document"
)

// The address associated with the specified user.
type Address struct {

	// The country of the address.
	Country *string

	// A string containing a formatted version of the address for display.
	Formatted *string

	// A string of the address locality.
	Locality *string

	// The postal code of the address.
	PostalCode *string

	// A Boolean value representing whether this is the primary address for the
	// associated resource.
	Primary bool

	// The region of the address.
	Region *string

	// The street of the address.
	StreetAddress *string

	// A string representing the type of address. For example, "Home."
	Type *string

	noSmithyDocumentSerde
}

// A unique identifier for a user or group that is not the primary identifier. This
// value can be an identifier from an external identity provider (IdP) that is
// associated with the user, the group, or a unique attribute. For example, a
// unique GroupDisplayName.
//
// The following types satisfy this interface:
//
//	AlternateIdentifierMemberExternalId
//	AlternateIdentifierMemberUniqueAttribute
type AlternateIdentifier interface {
	isAlternateIdentifier()
}

// The identifier issued to this resource by an external identity provider.
type AlternateIdentifierMemberExternalId struct {
	Value ExternalId

	noSmithyDocumentSerde
}

func (*AlternateIdentifierMemberExternalId) isAlternateIdentifier() {}

// An entity attribute that's unique to a specific entity.
type AlternateIdentifierMemberUniqueAttribute struct {
	Value UniqueAttribute

	noSmithyDocumentSerde
}

func (*AlternateIdentifierMemberUniqueAttribute) isAlternateIdentifier() {}

// An operation that applies to the requested group. This operation might add,
// replace, or remove an attribute.
type AttributeOperation struct {

	// A string representation of the path to a given attribute or sub-attribute.
	// Supports JMESPath.
	//
	// This member is required.
	AttributePath *string

	// The value of the attribute. This is a Document type. This type is not supported
	// by Java V1, Go V1, and older versions of the AWS CLI.
	AttributeValue document.Interface

	noSmithyDocumentSerde
}

// The email address associated with the user.
type Email struct {

	// A Boolean value representing whether this is the primary email address for the
	// associated resource.
	Primary bool

	// A string representing the type of address. For example, "Work."
	Type *string

	// A string containing an email address. For example, "johndoe@amazon.com."
	Value *string

	noSmithyDocumentSerde
}

// The identifier issued to this resource by an external identity provider.
type ExternalId struct {

	// The identifier issued to this resource by an external identity provider.
	//
	// This member is required.
	Id *string

	// The issuer for an external identifier.
	//
	// This member is required.
	Issuer *string

	noSmithyDocumentSerde
}

// A query filter used by ListUsers and ListGroups. This filter object provides the
// attribute name and attribute value to search users or groups.
type Filter struct {

	// The attribute path that is used to specify which attribute name to search.
	// Length limit is 255 characters. For example, UserName is a valid attribute path
	// for the ListUsers API, and DisplayName is a valid attribute path for the
	// ListGroups API.
	//
	// This member is required.
	AttributePath *string

	// Represents the data for an attribute. Each attribute value is described as a
	// name-value pair.
	//
	// This member is required.
	AttributeValue *string

	noSmithyDocumentSerde
}

// A group object that contains a specified group’s metadata and attributes.
type Group struct {

	// The identifier for a group in the identity store.
	//
	// This member is required.
	GroupId *string

	// The globally unique identifier for the identity store.
	//
	// This member is required.
	IdentityStoreId *string

	// A string containing a description of the specified group.
	Description *string

	// The group’s display name value. The length limit is 1,024 characters. This value
	// can consist of letters, accented characters, symbols, numbers, punctuation, tab,
	// new line, carriage return, space, and nonbreaking space in this attribute. This
	// value is specified at the time the group is created and stored as an attribute
	// of the group object in the identity store.
	DisplayName *string

	// A list of ExternalId objects that contains the identifiers issued to this
	// resource by an external identity provider.
	ExternalIds []ExternalId

	noSmithyDocumentSerde
}

// Contains the identifiers for a group, a group member, and a GroupMembership
// object in the identity store.
type GroupMembership struct {

	// The globally unique identifier for the identity store.
	//
	// This member is required.
	IdentityStoreId *string

	// The identifier for a group in the identity store.
	GroupId *string

	// An object that contains the identifier of a group member. Setting the UserID
	// field to the specific identifier for a user indicates that the user is a member
	// of the group.
	MemberId MemberId

	// The identifier for a GroupMembership object in an identity store.
	MembershipId *string

	noSmithyDocumentSerde
}

// Indicates whether a resource is a member of a group in the identity store.
type GroupMembershipExistenceResult struct {

	// The identifier for a group in the identity store.
	GroupId *string

	// An object that contains the identifier of a group member. Setting the UserID
	// field to the specific identifier for a user indicates that the user is a member
	// of the group.
	MemberId MemberId

	// Indicates whether a membership relation exists or not.
	MembershipExists bool

	noSmithyDocumentSerde
}

// An object containing the identifier of a group member.
//
// The following types satisfy this interface:
//
//	MemberIdMemberUserId
type MemberId interface {
	isMemberId()
}

// An object containing the identifiers of resources that can be members.
type MemberIdMemberUserId struct {
	Value string

	noSmithyDocumentSerde
}

func (*MemberIdMemberUserId) isMemberId() {}

// The full name of the user.
type Name struct {

	// The family name of the user.
	FamilyName *string

	// A string containing a formatted version of the name for display.
	Formatted *string

	// The given name of the user.
	GivenName *string

	// The honorific prefix of the user. For example, "Dr."
	HonorificPrefix *string

	// The honorific suffix of the user. For example, "M.D."
	HonorificSuffix *string

	// The middle name of the user.
	MiddleName *string

	noSmithyDocumentSerde
}

// The phone number associated with the user.
type PhoneNumber struct {

	// A Boolean value representing whether this is the primary phone number for the
	// associated resource.
	Primary bool

	// A string representing the type of a phone number. For example, "Mobile."
	Type *string

	// A string containing a phone number. For example, "8675309" or "+1 (800)
	// 123-4567".
	Value *string

	noSmithyDocumentSerde
}

// An entity attribute that's unique to a specific entity.
type UniqueAttribute struct {

	// A string representation of the path to a given attribute or sub-attribute.
	// Supports JMESPath.
	//
	// This member is required.
	AttributePath *string

	// The value of the attribute. This is a Document type. This type is not supported
	// by Java V1, Go V1, and older versions of the AWS CLI.
	//
	// This member is required.
	AttributeValue document.Interface

	noSmithyDocumentSerde
}

// A user object that contains a specified user’s metadata and attributes.
type User struct {

	// The globally unique identifier for the identity store.
	//
	// This member is required.
	IdentityStoreId *string

	// The identifier for a user in the identity store.
	//
	// This member is required.
	UserId *string

	// A list of Address objects containing addresses associated with the user.
	Addresses []Address

	// A string containing the user's name that's formatted for display when the user
	// is referenced. For example, "John Doe."
	DisplayName *string

	// A list of Email objects containing email addresses associated with the user.
	Emails []Email

	// A list of ExternalId objects that contains the identifiers issued to this
	// resource by an external identity provider.
	ExternalIds []ExternalId

	// A string containing the user's geographical region or location.
	Locale *string

	// An object containing the user's name.
	Name *Name

	// A string containing an alternate name for the user.
	NickName *string

	// A list of PhoneNumber objects containing phone numbers associated with the user.
	PhoneNumbers []PhoneNumber

	// A string containing the preferred language of the user. For example, "American
	// English" or "en-us."
	PreferredLanguage *string

	// A string containing a URL that may be associated with the user.
	ProfileUrl *string

	// A string containing the user's time zone.
	Timezone *string

	// A string containing the user's title. Possible values depend on each customer's
	// specific needs, so they are left unspecified.
	Title *string

	// A unique string used to identify the user. The length limit is 128 characters.
	// This value can consist of letters, accented characters, symbols, numbers, and
	// punctuation. This value is specified at the time the user is created and stored
	// as an attribute of the user object in the identity store.
	UserName *string

	// A string indicating the user's type. Possible values depend on each customer's
	// specific needs, so they are left unspecified.
	UserType *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde

// UnknownUnionMember is returned when a union member is returned over the wire,
// but has an unknown tag.
type UnknownUnionMember struct {
	Tag   string
	Value []byte

	noSmithyDocumentSerde
}

func (*UnknownUnionMember) isAlternateIdentifier() {}
func (*UnknownUnionMember) isMemberId()            {}
