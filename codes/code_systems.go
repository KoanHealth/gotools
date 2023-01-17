package codes

// HEDIS 2019 included these - not sure what data they come from:
// https://srs-health.com/wp-content/uploads/2016/04/SOP-Codes-3.21.14.pdf

type CodeSystem string

const (
	CODE_SYSTEM_SOURCE_OF_PAYMENT CodeSystem = "SOP"
	CODE_SYSTEM_LOINC             CodeSystem = "LOINC"

	// Medications

	CODE_SYSTEM_RXNORM CodeSystem = "RXNORM"
	CODE_SYSTEM_NDC    CodeSystem = "NDC"
	CODE_SYSTEM_MULTUM CodeSystem = "MULTUM" // Drugs
	CODE_SYSTEM_CVX    CodeSystem = "CVX"    // Vaccinations

	// Procedures

	CODE_SYSTEM_CPT        CodeSystem = "CPT"
	CODE_SYSTEM_CPT2       CodeSystem = "CPT2"
	CODE_SYSTEM_HCPCS      CodeSystem = "HCPCS"
	CODE_SYSTEM_ICD9_PROC  CodeSystem = "ICD9PCS"
	CODE_SYSTEM_ICD10_PROC CodeSystem = "ICD10PCS"
	CODE_SYSTEM_MODIFIER   CodeSystem = "MOD"   // CPT Modifier
	CODE_SYSTEM_HIPPS      CodeSystem = "HIPPS" // https://www.cms.gov/Medicare/Medicare-Fee-for-Service-Payment/ProspMedicareFeeSvcPmtGen/HIPPSCodes

	// Diagnoses

	CODE_SYSTEM_SNOMED     CodeSystem = "SNOMED"
	CODE_SYSTEM_ICD9_DIAG  CodeSystem = "ICD9CM"
	CODE_SYSTEM_ICD10_DIAG CodeSystem = "ICD10CM"

	// Claims

	CODE_SYSTEM_DRG              CodeSystem = "MSDRG"
	CODE_SYSTEM_REVENUE          CodeSystem = "REV"
	CODE_SYSTEM_TYPE_OF_BILL     CodeSystem = "TOB"
	CODE_SYSTEM_PLACE_OF_SERVICE CodeSystem = "POS"

	// Providers

	CODE_SYSTEM_TAAXONOMY CodeSystem = "TAXONOMY" // Specialty

	CODE_SYSTEM_UNKNOWN CodeSystem = "UNKNOWN"
)

var CodeSystemMap = map[CodeSystem]string{
	CODE_SYSTEM_SOURCE_OF_PAYMENT: "Source of Payment",
	CODE_SYSTEM_LOINC:             "LOINC",
	CODE_SYSTEM_RXNORM:            "RxNorm",
	CODE_SYSTEM_NDC:               "National Drug Code",
	CODE_SYSTEM_MULTUM:            "Multum",
	CODE_SYSTEM_CVX:               "Vaccines",
	CODE_SYSTEM_CPT:               "CPT",
	CODE_SYSTEM_CPT2:              "CPT2",
	CODE_SYSTEM_HCPCS:             "HCPCS",
	CODE_SYSTEM_HIPPS:             "HIPPS",
	CODE_SYSTEM_ICD9_PROC:         "ICD v9 Procedures",
	CODE_SYSTEM_ICD10_PROC:        "ICD v10 Procedures",
	CODE_SYSTEM_MODIFIER:          "CPT Modifier",
	CODE_SYSTEM_SNOMED:            "SNOMED",
	CODE_SYSTEM_ICD9_DIAG:         "ICD v9 Diagnoses",
	CODE_SYSTEM_ICD10_DIAG:        "ICD v10 Diagnoses",
	CODE_SYSTEM_DRG:               "MS DRG",
	CODE_SYSTEM_REVENUE:           "Revenue Code",
	CODE_SYSTEM_TYPE_OF_BILL:      "Ttype of Bill",
	CODE_SYSTEM_PLACE_OF_SERVICE:  "Place of Service",
	CODE_SYSTEM_TAAXONOMY:         "Provider Taxonomy",
	CODE_SYSTEM_UNKNOWN:           "Unknown",
}
