package codes

// Refer to: http://www.hl7.org/OID/

func LookupOidCodeSystem(oid string) CodeSystem {
	if oid == "" {
		return CODE_SYSTEM_UNKNOWN
	}

	// OID's seen from real data, but not identified here:
	// 2.16.840.1.113883.3.247.1.1: Intelligent Medical Objects ProblemIT
	// 2.16.840.1.113883.5.4: Act Class: https://www.hl7.org/fhir/v3/ActClass/cs.html
	// 2.16.840.1.113883.5.6: Act Class: https://www.hl7.org/fhir/v3/ActClass/cs.html
	// 2.16.840.1.113883.5.8: Act Reason
	// 2.16.840.1.113883.6.68: Medispan GPI
	// 2.16.840.1.113883.6.162: Master Drug Database
	// 2.16.840.1.113883.6.253: Medispan Drug

	switch oid {
	case "2.16.840.1.113883.6.96":
		return CODE_SYSTEM_SNOMED
	case "2.16.840.1.113883.6.1":
		return CODE_SYSTEM_LOINC
	case "2.16.840.1.113883.6.88":
		return CODE_SYSTEM_RXNORM
	case "2.16.840.1.113883.6.12":
		return CODE_SYSTEM_CPT
	case "2.16.840.1.113883.6.285", "2.16.840.1.113883.6.14":
		return CODE_SYSTEM_HCPCS
	case "2.16.840.1.113883.6.69":
		return CODE_SYSTEM_NDC
	case "2.16.840.1.113883.12.292":
		return CODE_SYSTEM_CVX
	case "2.16.840.1.113883.6.104":
		return CODE_SYSTEM_ICD9_PROC
	case "2.16.840.1.113883.6.103":
		return CODE_SYSTEM_ICD9_DIAG
	case "2.16.840.1.113883.6.3":
		return CODE_SYSTEM_ICD10_DIAG
	case "2.16.840.1.113883.6.4", "2.16.840.1.113883.6.90":
		return CODE_SYSTEM_ICD10_PROC
	case "2.16.840.1.113883.15.4":
		return CODE_SYSTEM_HIPPS // https: //oidref.com/2.16.840.1.113883.15.4
	default:
		return CODE_SYSTEM_UNKNOWN
	}
}
