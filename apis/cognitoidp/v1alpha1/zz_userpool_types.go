/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AccountRecoverySettingObservation struct {
}

type AccountRecoverySettingParameters struct {

	// +kubebuilder:validation:Required
	RecoveryMechanism []RecoveryMechanismParameters `json:"recoveryMechanism" tf:"recovery_mechanism,omitempty"`
}

type AdminCreateUserConfigObservation struct {
}

type AdminCreateUserConfigParameters struct {

	// +kubebuilder:validation:Optional
	AllowAdminCreateUserOnly *bool `json:"allowAdminCreateUserOnly,omitempty" tf:"allow_admin_create_user_only,omitempty"`

	// +kubebuilder:validation:Optional
	InviteMessageTemplate []InviteMessageTemplateParameters `json:"inviteMessageTemplate,omitempty" tf:"invite_message_template,omitempty"`
}

type CustomEmailSenderObservation struct {
}

type CustomEmailSenderParameters struct {

	// +kubebuilder:validation:Required
	LambdaArn *string `json:"lambdaArn" tf:"lambda_arn,omitempty"`

	// +kubebuilder:validation:Required
	LambdaVersion *string `json:"lambdaVersion" tf:"lambda_version,omitempty"`
}

type CustomSMSSenderObservation struct {
}

type CustomSMSSenderParameters struct {

	// +kubebuilder:validation:Required
	LambdaArn *string `json:"lambdaArn" tf:"lambda_arn,omitempty"`

	// +kubebuilder:validation:Required
	LambdaVersion *string `json:"lambdaVersion" tf:"lambda_version,omitempty"`
}

type DeviceConfigurationObservation struct {
}

type DeviceConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	ChallengeRequiredOnNewDevice *bool `json:"challengeRequiredOnNewDevice,omitempty" tf:"challenge_required_on_new_device,omitempty"`

	// +kubebuilder:validation:Optional
	DeviceOnlyRememberedOnUserPrompt *bool `json:"deviceOnlyRememberedOnUserPrompt,omitempty" tf:"device_only_remembered_on_user_prompt,omitempty"`
}

type EmailConfigurationObservation struct {
}

type EmailConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	ConfigurationSet *string `json:"configurationSet,omitempty" tf:"configuration_set,omitempty"`

	// +kubebuilder:validation:Optional
	EmailSendingAccount *string `json:"emailSendingAccount,omitempty" tf:"email_sending_account,omitempty"`

	// +kubebuilder:validation:Optional
	FromEmailAddress *string `json:"fromEmailAddress,omitempty" tf:"from_email_address,omitempty"`

	// +kubebuilder:validation:Optional
	ReplyToEmailAddress *string `json:"replyToEmailAddress,omitempty" tf:"reply_to_email_address,omitempty"`

	// +kubebuilder:validation:Optional
	SourceArn *string `json:"sourceArn,omitempty" tf:"source_arn,omitempty"`
}

type InviteMessageTemplateObservation struct {
}

type InviteMessageTemplateParameters struct {

	// +kubebuilder:validation:Optional
	EmailMessage *string `json:"emailMessage,omitempty" tf:"email_message,omitempty"`

	// +kubebuilder:validation:Optional
	EmailSubject *string `json:"emailSubject,omitempty" tf:"email_subject,omitempty"`

	// +kubebuilder:validation:Optional
	SMSMessage *string `json:"smsMessage,omitempty" tf:"sms_message,omitempty"`
}

type LambdaConfigObservation struct {
}

type LambdaConfigParameters struct {

	// +kubebuilder:validation:Optional
	CreateAuthChallenge *string `json:"createAuthChallenge,omitempty" tf:"create_auth_challenge,omitempty"`

	// +kubebuilder:validation:Optional
	CustomEmailSender []CustomEmailSenderParameters `json:"customEmailSender,omitempty" tf:"custom_email_sender,omitempty"`

	// +kubebuilder:validation:Optional
	CustomMessage *string `json:"customMessage,omitempty" tf:"custom_message,omitempty"`

	// +kubebuilder:validation:Optional
	CustomSMSSender []CustomSMSSenderParameters `json:"customSmsSender,omitempty" tf:"custom_sms_sender,omitempty"`

	// +kubebuilder:validation:Optional
	DefineAuthChallenge *string `json:"defineAuthChallenge,omitempty" tf:"define_auth_challenge,omitempty"`

	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// +kubebuilder:validation:Optional
	PostAuthentication *string `json:"postAuthentication,omitempty" tf:"post_authentication,omitempty"`

	// +kubebuilder:validation:Optional
	PostConfirmation *string `json:"postConfirmation,omitempty" tf:"post_confirmation,omitempty"`

	// +kubebuilder:validation:Optional
	PreAuthentication *string `json:"preAuthentication,omitempty" tf:"pre_authentication,omitempty"`

	// +kubebuilder:validation:Optional
	PreSignUp *string `json:"preSignUp,omitempty" tf:"pre_sign_up,omitempty"`

	// +kubebuilder:validation:Optional
	PreTokenGeneration *string `json:"preTokenGeneration,omitempty" tf:"pre_token_generation,omitempty"`

	// +kubebuilder:validation:Optional
	UserMigration *string `json:"userMigration,omitempty" tf:"user_migration,omitempty"`

	// +kubebuilder:validation:Optional
	VerifyAuthChallengeResponse *string `json:"verifyAuthChallengeResponse,omitempty" tf:"verify_auth_challenge_response,omitempty"`
}

type NumberAttributeConstraintsObservation struct {
}

type NumberAttributeConstraintsParameters struct {

	// +kubebuilder:validation:Optional
	MaxValue *string `json:"maxValue,omitempty" tf:"max_value,omitempty"`

	// +kubebuilder:validation:Optional
	MinValue *string `json:"minValue,omitempty" tf:"min_value,omitempty"`
}

type PasswordPolicyObservation struct {
}

type PasswordPolicyParameters struct {

	// +kubebuilder:validation:Optional
	MinimumLength *float64 `json:"minimumLength,omitempty" tf:"minimum_length,omitempty"`

	// +kubebuilder:validation:Optional
	RequireLowercase *bool `json:"requireLowercase,omitempty" tf:"require_lowercase,omitempty"`

	// +kubebuilder:validation:Optional
	RequireNumbers *bool `json:"requireNumbers,omitempty" tf:"require_numbers,omitempty"`

	// +kubebuilder:validation:Optional
	RequireSymbols *bool `json:"requireSymbols,omitempty" tf:"require_symbols,omitempty"`

	// +kubebuilder:validation:Optional
	RequireUppercase *bool `json:"requireUppercase,omitempty" tf:"require_uppercase,omitempty"`

	// +kubebuilder:validation:Optional
	TemporaryPasswordValidityDays *float64 `json:"temporaryPasswordValidityDays,omitempty" tf:"temporary_password_validity_days,omitempty"`
}

type RecoveryMechanismObservation struct {
}

type RecoveryMechanismParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Priority *float64 `json:"priority" tf:"priority,omitempty"`
}

type SMSConfigurationObservation struct {
}

type SMSConfigurationParameters struct {

	// +kubebuilder:validation:Required
	ExternalID *string `json:"externalId" tf:"external_id,omitempty"`

	// +kubebuilder:validation:Required
	SnsCallerArn *string `json:"snsCallerArn" tf:"sns_caller_arn,omitempty"`
}

type SchemaObservation struct {
}

type SchemaParameters struct {

	// +kubebuilder:validation:Required
	AttributeDataType *string `json:"attributeDataType" tf:"attribute_data_type,omitempty"`

	// +kubebuilder:validation:Optional
	DeveloperOnlyAttribute *bool `json:"developerOnlyAttribute,omitempty" tf:"developer_only_attribute,omitempty"`

	// +kubebuilder:validation:Optional
	Mutable *bool `json:"mutable,omitempty" tf:"mutable,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	NumberAttributeConstraints []NumberAttributeConstraintsParameters `json:"numberAttributeConstraints,omitempty" tf:"number_attribute_constraints,omitempty"`

	// +kubebuilder:validation:Optional
	Required *bool `json:"required,omitempty" tf:"required,omitempty"`

	// +kubebuilder:validation:Optional
	StringAttributeConstraints []StringAttributeConstraintsParameters `json:"stringAttributeConstraints,omitempty" tf:"string_attribute_constraints,omitempty"`
}

type SoftwareTokenMfaConfigurationObservation struct {
}

type SoftwareTokenMfaConfigurationParameters struct {

	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`
}

type StringAttributeConstraintsObservation struct {
}

type StringAttributeConstraintsParameters struct {

	// +kubebuilder:validation:Optional
	MaxLength *string `json:"maxLength,omitempty" tf:"max_length,omitempty"`

	// +kubebuilder:validation:Optional
	MinLength *string `json:"minLength,omitempty" tf:"min_length,omitempty"`
}

type UserPoolAddOnsObservation struct {
}

type UserPoolAddOnsParameters struct {

	// +kubebuilder:validation:Required
	AdvancedSecurityMode *string `json:"advancedSecurityMode" tf:"advanced_security_mode,omitempty"`
}

type UserPoolObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	CreationDate *string `json:"creationDate,omitempty" tf:"creation_date,omitempty"`

	CustomDomain *string `json:"customDomain,omitempty" tf:"custom_domain,omitempty"`

	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	EstimatedNumberOfUsers *float64 `json:"estimatedNumberOfUsers,omitempty" tf:"estimated_number_of_users,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	LastModifiedDate *string `json:"lastModifiedDate,omitempty" tf:"last_modified_date,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type UserPoolParameters struct {

	// +kubebuilder:validation:Optional
	AccountRecoverySetting []AccountRecoverySettingParameters `json:"accountRecoverySetting,omitempty" tf:"account_recovery_setting,omitempty"`

	// +kubebuilder:validation:Optional
	AdminCreateUserConfig []AdminCreateUserConfigParameters `json:"adminCreateUserConfig,omitempty" tf:"admin_create_user_config,omitempty"`

	// +kubebuilder:validation:Optional
	AliasAttributes []*string `json:"aliasAttributes,omitempty" tf:"alias_attributes,omitempty"`

	// +kubebuilder:validation:Optional
	AutoVerifiedAttributes []*string `json:"autoVerifiedAttributes,omitempty" tf:"auto_verified_attributes,omitempty"`

	// +kubebuilder:validation:Optional
	DeviceConfiguration []DeviceConfigurationParameters `json:"deviceConfiguration,omitempty" tf:"device_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	EmailConfiguration []EmailConfigurationParameters `json:"emailConfiguration,omitempty" tf:"email_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	EmailVerificationMessage *string `json:"emailVerificationMessage,omitempty" tf:"email_verification_message,omitempty"`

	// +kubebuilder:validation:Optional
	EmailVerificationSubject *string `json:"emailVerificationSubject,omitempty" tf:"email_verification_subject,omitempty"`

	// +kubebuilder:validation:Optional
	LambdaConfig []LambdaConfigParameters `json:"lambdaConfig,omitempty" tf:"lambda_config,omitempty"`

	// +kubebuilder:validation:Optional
	MfaConfiguration *string `json:"mfaConfiguration,omitempty" tf:"mfa_configuration,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	PasswordPolicy []PasswordPolicyParameters `json:"passwordPolicy,omitempty" tf:"password_policy,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	SMSAuthenticationMessage *string `json:"smsAuthenticationMessage,omitempty" tf:"sms_authentication_message,omitempty"`

	// +kubebuilder:validation:Optional
	SMSConfiguration []SMSConfigurationParameters `json:"smsConfiguration,omitempty" tf:"sms_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	SMSVerificationMessage *string `json:"smsVerificationMessage,omitempty" tf:"sms_verification_message,omitempty"`

	// +kubebuilder:validation:Optional
	Schema []SchemaParameters `json:"schema,omitempty" tf:"schema,omitempty"`

	// +kubebuilder:validation:Optional
	SoftwareTokenMfaConfiguration []SoftwareTokenMfaConfigurationParameters `json:"softwareTokenMfaConfiguration,omitempty" tf:"software_token_mfa_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	UserPoolAddOns []UserPoolAddOnsParameters `json:"userPoolAddOns,omitempty" tf:"user_pool_add_ons,omitempty"`

	// +kubebuilder:validation:Optional
	UsernameAttributes []*string `json:"usernameAttributes,omitempty" tf:"username_attributes,omitempty"`

	// +kubebuilder:validation:Optional
	UsernameConfiguration []UsernameConfigurationParameters `json:"usernameConfiguration,omitempty" tf:"username_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	VerificationMessageTemplate []VerificationMessageTemplateParameters `json:"verificationMessageTemplate,omitempty" tf:"verification_message_template,omitempty"`
}

type UsernameConfigurationObservation struct {
}

type UsernameConfigurationParameters struct {

	// +kubebuilder:validation:Required
	CaseSensitive *bool `json:"caseSensitive" tf:"case_sensitive,omitempty"`
}

type VerificationMessageTemplateObservation struct {
}

type VerificationMessageTemplateParameters struct {

	// +kubebuilder:validation:Optional
	DefaultEmailOption *string `json:"defaultEmailOption,omitempty" tf:"default_email_option,omitempty"`

	// +kubebuilder:validation:Optional
	EmailMessage *string `json:"emailMessage,omitempty" tf:"email_message,omitempty"`

	// +kubebuilder:validation:Optional
	EmailMessageByLink *string `json:"emailMessageByLink,omitempty" tf:"email_message_by_link,omitempty"`

	// +kubebuilder:validation:Optional
	EmailSubject *string `json:"emailSubject,omitempty" tf:"email_subject,omitempty"`

	// +kubebuilder:validation:Optional
	EmailSubjectByLink *string `json:"emailSubjectByLink,omitempty" tf:"email_subject_by_link,omitempty"`

	// +kubebuilder:validation:Optional
	SMSMessage *string `json:"smsMessage,omitempty" tf:"sms_message,omitempty"`
}

// UserPoolSpec defines the desired state of UserPool
type UserPoolSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserPoolParameters `json:"forProvider"`
}

// UserPoolStatus defines the observed state of UserPool.
type UserPoolStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserPoolObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// UserPool is the Schema for the UserPools API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type UserPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UserPoolSpec   `json:"spec"`
	Status            UserPoolStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserPoolList contains a list of UserPools
type UserPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserPool `json:"items"`
}

// Repository type metadata.
var (
	UserPool_Kind             = "UserPool"
	UserPool_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: UserPool_Kind}.String()
	UserPool_KindAPIVersion   = UserPool_Kind + "." + CRDGroupVersion.String()
	UserPool_GroupVersionKind = CRDGroupVersion.WithKind(UserPool_Kind)
)

func init() {
	SchemeBuilder.Register(&UserPool{}, &UserPoolList{})
}
