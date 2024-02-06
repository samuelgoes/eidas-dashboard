package models

type TrustedListGPT struct {
	TSLType                     string `json:"tslType"`
	TSLVersion                  string `json:"tslVersion"`
	TSLSequenceNumber           string `json:"tslSequenceNumber"`
	SchemeOperatorName          string `json:"schemeOperatorName"`
	SchemeTypeCommunityRules    string `json:"schemeTypeCommunityRules"`
	StatusDeterminationApproach string `json:"statusDeterminationApproach"`
	PointersToOtherTSL          []struct {
		TSLPointer struct {
			CountryCode                 string   `json:"countryCode"`
			Territory                   string   `json:"territory"`
			TSLType                     string   `json:"tslType"`
			SchemeOperatorName          string   `json:"schemeOperatorName"`
			SchemeTypeCommunityRules    string   `json:"schemeTypeCommunityRules"`
			StatusDeterminationApproach string   `json:"statusDeterminationApproach"`
			TSLSequenceNumber           string   `json:"tslSequenceNumber"`
			HistoricalInformationPeriod []string `json:"historicalInformationPeriod"`
			NextUpdate                  string   `json:"nextUpdate"`
			DistributionPoints          []struct {
				URI string `json:"uri"`
			} `json:"distributionPoints"`
		} `json:"tslPointer"`
	} `json:"pointersToOtherTSL"`
}

type TrustedListDTO struct {
	TLId              int                     `json:"tlId"`
	TechnicalLocation string                  `json:"technicalLocation"`
	HRLocation        string                  `json:"hrLocation"`
	Name              string                  `json:"name"`
	ExternalId        string                  `json:"externalId"`
	CountryCode       string                  `json:"countryCode"`
	TSLId             string                  `json:"tslId"`
	TSLTag            string                  `json:"tslTag"`
	SchemeInformation TSLSchemeInformationDTO `json:"schemeInformation"`
	Pointers          []TLPointerDTO          `json:"pointers"`
	//ServiceProviders          string `json:"serviceProviders"`
	Status             string `json:"status"`             //PROD, DRAFT
	AvailabilityStatus string `json:"availabilityStatus"` // AVAILABLE, UNSUPPORTED, UNAVAILABLE
	//Signature string `json:"signature"` // AVAILABLE, UNSUPPORTED, UNAVAILABLE
	BrowsePreviousProd         bool   `json:"browsePreviousProd"`
	FirstScanDate              string `json:"firstScanDate"`
	FileName                   string `json:"fileName"`
	PerformCheckChange         bool   `json:"performCheckChange"`
	CurrentCheckSequenceNumber int    `json:"currentCheckSequenceNumber"`
	ConformanceStatus          string `json:"conformanceStatus"` // ERROR, WARNING, INFO, IGNORE, SUCCESS, COMPUTING
	//MraInfo                    string `json:"mraInfo"`
	//CurrentCountryPointer      string `json:"currentCountryPointer"`
}

type TSLSchemeInformationDTO struct {
	TlId                     int                     `json:"tlId"`
	TechnicalLocation        string                  `json:"technicalLocation"`
	HrLocation               string                  `json:"hrLocation"`
	TlIdentifier             int                     `json:"tlIdentifier"`
	SequenceNumber           int                     `json:"sequenceNumber"`
	Territory                string                  `json:"territory"`
	Type                     string                  `json:"type"`
	HistoricalPeriod         int                     `json:"historicalPeriod"`
	IssueDate                string                  `json:"issueDate"`
	NextUpdateDate           string                  `json:"nextUpdateDate"`
	StatusDetermination      string                  `json:"statusDetermination"`
	SchemeOperatorNames      []TLGenericLangEntryDTO `json:"schemeOperatorNames"`
	PostalAddresses          []TLPostalAddressDTO    `json:"postalAddresses"`
	ElectronicAddresses      []TLGenericLangEntryDTO `json:"electronicAddresses"`
	SchemeNames              []TLGenericLangEntryDTO `json:"schemeNames"`
	InformationURIs          []TLGenericLangEntryDTO `json:"informationURIs"`
	SchemeTypeCommunityRules []TLGenericLangEntryDTO `json:"schemeTypeCommunityRules"`
	SchemePolicies           []TLGenericLangEntryDTO `json:"schemePolicies"`
	DistributionPoints       []TLGenericEntryDTO     `json:"distributionPoints"`
}

type TLGenericLangEntryDTO struct {
	TlId              int    `json:"tlId"`
	TechnicalLocation string `json:"technicalLocation"`
	HrLocation        string `json:"hrLocation"`
	Language          string `json:"language"`
	Value             string `json:"value"`
}

type TLPostalAddressDTO struct {
	TlId              int    `json:"tlId"`
	TechnicalLocation string `json:"technicalLocation"`
	HrLocation        string `json:"hrLocation"`
	Street            string `json:"street"`
	Locality          string `json:"locality"`
	PostalCode        string `json:"postalCode"`
	CountryCode       string `json:"countryCode"`
	Language          string `json:"language"`
	StateOrProvince   string `json:"stateOrProvince"`
}

type TLGenericEntryDTO struct {
	TlId              int    `json:"tlId"`
	TechnicalLocation string `json:"technicalLocation"`
	HrLocation        string `json:"hrLocation"`
	Value             string `json:"value"`
}

type TLPointerDTO struct {
	TlId                int                     `json:"tlId"`
	TechnicalLocation   string                  `json:"technicalLocation"`
	HrLocation          string                  `json:"hrLocation"`
	TlLocation          string                  `json:"tlLocation"`
	SchemeTerritory     string                  `json:"schemeTerritory"`
	TlType              string                  `json:"tlType"`
	MimeType            string                  `json:"mimeType"` //XML, PDF
	SchemeOperatorNames []TLGenericLangEntryDTO `json:"schemeOperatorNames"`
	CommunityRules      []TLGenericLangEntryDTO `json:"communityRules"`
	DigitalIdentities   []TLDigitalIdentityDTO  `json:"digitalIdentities"`
	MraInfo             MRAInfoDTO              `json:"mraInfo"`
}

type TLDigitalIdentityDTO struct {
	TlId              int                `json:"tlId"`
	TechnicalLocation string             `json:"technicalLocation"`
	HrLocation        string             `json:"hrLocation"`
	SubjectName       string             `json:"subjectName"`
	X509ski           string             `json:"x509ski"`
	Other             []string           `json:"other"`
	Certificates      []TLCertificateDTO `json:"certificates"`
}

type TLCertificateDTO struct {
	TlId                 int                `json:"tlId"`
	TechnicalLocation    string             `json:"technicalLocation"`
	HrLocation           string             `json:"hrLocation"`
	Base64               string             `json:"base64"`
	SubjectShortName     string             `json:"subjectShortName"`
	Subject              string             `json:"subject"`
	SerialNumber         string             `json:"serialNumber"`
	Issuer               string             `json:"issuer"`
	DigestAlgorithm      string             `json:"digestAlgorithm"`
	EncryptionAlgorithm  string             `json:"encryptionAlgorithm"`
	NotBefore            string             `json:"notBefore"`
	NotAfter             string             `json:"notAfter"`
	Information          string             `json:"information"`
	KeyUsageList         []string           `json:"keyUsageList"`
	ExtendedKeyUsageList []string           `json:"extendedKeyUsageList"`
	SkiB64               string             `json:"skiB64"`
	IssuerX500           CertificateX500DTO `json:"issuerX500"`
	SubjectX500          CertificateX500DTO `json:"subjectX500"`
}

type CertificateX500DTO struct {
	CommonName       string `json:"commonName"`
	SerialNumber     string `json:"serialNumber"`
	Locality         string `json:"locality"`
	State            string `json:"state"`
	CountryName      string `json:"countryName"`
	OrganizationName string `json:"organizationName"`
	GivenName        string `json:"givenName"`
	Pseudonym        string `json:"pseudonym"`
	Email            string `json:"email"`
}

type MRAInfoDTO struct {
	TechnicalType                       string                       `json:"technicalType"`
	PointingContractingPartyLegislation string                       `json:"pointingContractingPartyLegislation"`
	PointedContractingPartyLegislation  string                       `json:"pointedContractingPartyLegislation"`
	Version                             string                       `json:"version"`
	ServiceEquivalences                 []ServiceEquivalenceDTO      `json:"serviceEquivalences"`
	ParsedEquivalences                  MRAEquivalencesForServiceDTO `json:"parsedEquivalences"`
}

type ServiceEquivalenceDTO struct {
	LegalInfo                                string                                  `json:"legalInfo"`
	Status                                   string                                  `json:"status"` // ENACTED, REPEALED
	StartDate                                string                                  `json:"startDate"`
	ValidStatusEquivalences                  []StatusEquivalenceDTO                  `json:"validStatusEquivalences"`
	InvalidStatusEquivalences                []StatusEquivalenceDTO                  `json:"invalidStatusEquivalences"`
	CertificateContentEquivalenceContextDTOS CertificateContentEquivalanceContextDTO `json:"certificateContentEquivalenceContextDTOS"`
	QualifierEquivalence                     string                                  `json:"qualifierEquivalence"`
	TypeAsiEquivalence                       []TypeAsiEquivalenceDTO                 `json:"typeAsiEquivalence"`
}

type StatusEquivalenceDTO struct {
	PointedParties  []string `json:"pointedParties"`
	PointingParties []string `json:"pointingParties"`
}

type CertificateContentEquivalanceContextDTO struct {
	MraEquivalenceContext            string                           `json:"pointedParties"` // QC_COMPLIANCE, QC_TYPE, QC_QSCD
	CertificateContentEquivalenceDTO CertificateContentEquivalanceDTO `json:"certificateContentEquivalenceDTO"`
}

type CertificateContentEquivalanceDTO struct {
	PointedCriteria    string             `json:"pointedCriteria"`
	PointingCriteria   string             `json:"pointingCriteria"`
	ContentReplacement QCStatementOidsDTO `json:"contentReplacement"`
}

type QCStatementOidsDTO struct {
	QcStatementIds           []string `json:"qcStatementIds"`
	QcTypeIds                []string `json:"qcTypeIds"`
	QcCClegislations         []string `json:"qcCClegislations"`
	QcStatementIdsToRemove   []string `json:"qcStatementIdsToRemove"`
	QcTypeIdsToRemove        []string `json:"qcTypeIdsToRemove"`
	QcCClegislationsToRemove []string `json:"qcCClegislationsToRemove"`
}

type TypeAsiEquivalenceDTO struct {
	PointingParty TypeAsiDTO `json:"pointingParty"`
	PointedParty  TypeAsiDTO `json:"pointedParty"`
}

type TypeAsiDTO struct {
	Type string `json:"type"`
	Asi  string `json:"asi"`
}

type MRAEquivalencesForServiceDTO struct {
	ServiceTypeAsiEu             TypeAsiDTO              `json:"serviceTypeAsiEu"`
	AsiMraToAsiEuDtoMap          TypeAsiEquivalenceDTO   `json:"asiMraToAsiEuDtoMap"`
	EuStatusToAsiEquivalencesMap []TypeAsiEquivalenceDTO `json:"euStatusToAsiEquivalencesMap"`
}

type TLServiceProviderDTO struct {
	TlId                   int                         `json:"tlId"`
	TechnicalLocation      string                      `json:"technicalLocation"`
	HrLocation             string                      `json:"hrLocation"`
	Names                  []TLGenericLangEntryDTO     `json:"names"`
	TradeNames             []TLGenericLangEntryDTO     `json:"tradeNames"`
	PostalAddresses        []TLPostalAddressDTO        `json:"postalAddresses"`
	ElectronicAddresses    []TLGenericLangEntryDTO     `json:"electronicAddresses"`
	InformationURIs        []TLGenericLangEntryDTO     `json:"informationURIs"`
	Extensions             []TLInformationExtensionDTO `json:"extensions"`
	Services               []TLServiceDTO              `json:"services"`
	GetqServiceTypes       []string                    `json:"getqServiceTypes"`
	TspId                  int                         `json:"tspId"`
	GlobalStatus           string                      `json:"globalStatus"` // ACTIVE, INACTIVE, TOB
	NbServicesActive       int                         `json:"nbServicesActive"`
	TakenOverByServicesMap []TLServiceDTO              `json:"takenOverByServicesMap"`
	TobservicesMap         []TLServiceDTO              `json:"tobservicesMap"`
}

type TLInformationExtensionDTO struct {
	TlId              int    `json:"tlId"`
	TechnicalLocation string `json:"technicalLocation"`
	HrLocation        string `json:"hrLocation"`
	Critical          bool   `json:"critical"`
	Value             string `json:"value"`
}

type TLServiceDTO struct {
	TlId                      int                     `json:"tlId"`
	TechnicalLocation         string                  `json:"technicalLocation"`
	HrLocation                string                  `json:"hrLocation"`
	TypeIdentifier            string                  `json:"typeIdentifier"`
	CurrentStatus             string                  `json:"currentStatus"`
	Active                    bool                    `json:"active"`
	TakenOverBy               bool                    `json:"takenOverBy"`
	CurrentStatusStartingDate string                  `json:"currentStatusStartingDate"`
	DigitalIdentity           TLDigitalIdentityDTO    `json:"digitalIdentity"`
	Names                     []TLGenericLangEntryDTO `json:"names"`
	SchemeDefinitionURIs      []TLGenericLangEntryDTO `json:"schemeDefinitionURIs"`
	TspDefinitionURIs         []TLGenericLangEntryDTO `json:"tspDefinitionURIs"`
	SupplyPoints              []TLGenericEntryDTO     `json:"supplyPoints"`
	Extensions                []TLServiceExtensionDTO `json:"extensions"`
	Histories                 []TLServiceHistoryDTO   `json:"histories"`
	GetqServiceTypes          []string                `json:"getqServiceTypes"`
}

type TLServiceExtensionDTO struct {
	TlId                             int                           `json:"tlId"`
	TechnicalLocation                string                        `json:"technicalLocation"`
	HrLocation                       string                        `json:"hrLocation"`
	Critical                         bool                          `json:"critical"`
	TakenOverBy                      TLTakenOverByDTO              `json:"takenOverBy"`
	AdditionnalServiceInformation    TLGenericLangEntryDTO         `json:"additionnalServiceInformation"`
	QualificationExtensions          []TLQualificationExtensionDTO `json:"qualificationExtensions"`
	ExpiredCertificateRevocationDate string                        `json:"expiredCertificateRevocationDate"`
	AnyType                          string                        `json:"anyType"`
}

type TLTakenOverByDTO struct {
	TlId              int                     `json:"tlId"`
	TechnicalLocation string                  `json:"technicalLocation"`
	HrLocation        string                  `json:"hrLocation"`
	Url               TLGenericLangEntryDTO   `json:"url"`
	Territory         string                  `json:"territory"`
	TspNames          []TLGenericLangEntryDTO `json:"tspNames"`
	OperatorNames     []TLGenericLangEntryDTO `json:"operatorNames"`
	OtherQualifiers   []string                `json:"otherQualifiers"`
}

type TLQualificationExtensionDTO struct {
	TlId               int           `json:"tlId"`
	TechnicalLocation  string        `json:"technicalLocation"`
	HrLocation         string        `json:"hrLocation"`
	QualificationTypes []string      `json:"qualificationTypes"`
	Criteria           TLCriteriaDTO `json:"criteria"`
}

type TLCriteriaDTO struct {
	TlId                   int                     `json:"tlId"`
	TechnicalLocation      string                  `json:"technicalLocation"`
	HrLocation             string                  `json:"hrLocation"`
	Asserts                string                  `json:"asserts"`
	Description            string                  `json:"description"`
	KeyUsageList           []TLKeyUsageDTO         `json:"keyUsageList"`
	PolicyList             []TLPolicyDTO           `json:"policyList"`
	CriteriaList           []string                `json:"criteriaList"`
	ExtendedKeyUsageList   []string                `json:"extendedKeyUsageList"`
	CertDnaList            []string                `json:"certDnaList"`
	CompositeConditionList []CompositeConditionDTO `json:"compositeConditionList"`
}

type TLKeyUsageDTO struct {
	TlId              int              `json:"tlId"`
	TechnicalLocation string           `json:"technicalLocation"`
	HrLocation        string           `json:"hrLocation"`
	KeyUsageBits      TLKeyUsageBitDTO `json:"keyUsageBits"`
}

type TLKeyUsageBitDTO struct {
	TlId              int    `json:"tlId"`
	TechnicalLocation string `json:"technicalLocation"`
	HrLocation        string `json:"hrLocation"`
	Value             string `json:"value"`
	IsValue           bool   `json:"isValue"`
}

type TLPolicyDTO struct {
	TlId              int                `json:"tlId"`
	TechnicalLocation string             `json:"technicalLocation"`
	HrLocation        string             `json:"hrLocation"`
	PolicyBit         []TLPoliciesBitDTO `json:"policyBit"`
}

type TLPoliciesBitDTO struct {
	TlId                    int      `json:"tlId"`
	TechnicalLocation       string   `json:"technicalLocation"`
	HrLocation              string   `json:"hrLocation"`
	Description             string   `json:"description"`
	IdentifierValue         string   `json:"identifierValue"`
	IdentifierType          string   `json:"identifierType"`
	DocumentationReferences []string `json:"documentationReferences"`
}

type CompositeConditionDTO struct {
	MatchingCriteriaIndicator string   `json:"matchingCriteriaIndicator"` // ALL, AT_LEAST_ONE, NONE
	Children                  []string `json:"children"`
}

type TLServiceHistoryDTO struct {
	TypeIdentifier            string                  `json:"typeIdentifier"` // ALL, AT_LEAST_ONE, NONE
	CurrentStatus             string                  `json:"currentStatus"`
	CurrentStatusStartingDate string                  `json:"currentStatusStartingDate"`
	Names                     []TLGenericLangEntryDTO `json:"names"`
	Extensions                []TLServiceExtensionDTO `json:"extensions"`
	DigitalIdentity           TLDigitalIdentityDTO    `json:"digitalIdentity"`
	HistoryChanges            []string                `json:"historyChanges"` // TYPE_IDENTIFIER, CURRENT_STATUS, NAMES, ADDITIONAL_SERVICE_INFORMATION, DIGITAL_IDENTITY, EXPIRED_CERTIFICATE_REVOCATION_DATE, TAKEN_OVER_BY, QUALIFICATION_EXTENSIONS
}

type TspDTO struct {
	TspId         int             `json:"tspId"`
	Name          string          `json:"name"`
	CountryCode   string          `json:"countryCode"`
	Trustmark     string          `json:"trustmark"`
	QServiceTypes []string        `json:"qServiceTypes"`
	Services      []TspServiceDTO `json:"services"`
}

type TspServiceDTO struct {
	TspId            int      `json:"tspId"`
	ServiceId        int      `json:"serviceId"`
	CountryCode      string   `json:"countryCode"`
	ServiceName      string   `json:"serviceName"`
	Type             string   `json:"type"`
	CurrentStatus    string   `json:"currentStatus"`
	Tob              string   `json:"tob"`
	GetqServiceTypes []string `json:"getqServiceTypes"`
}
