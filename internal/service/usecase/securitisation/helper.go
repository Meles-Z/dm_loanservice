package securitisation

func computeEligibility(ddStatus string, flags []string) string {
	// ❌ If DD failed → not eligible
	if ddStatus == "fail" {
		return "not_eligible"
	}

	// ❌ If any blocking flag exists → not eligible
	for _, f := range flags {
		if f == "ineligible" {
			return "not_eligible"
		}
	}

	// ❌ If DD pending → review
	if ddStatus == "pending" {
		return "under_review"
	}

	// ✅ Passed everything
	return "eligible"
}

func calculateEligibility(ddStatus, flagStatus string) string {
	if ddStatus != "pass" {
		return "ineligible"
	}
	if flagStatus == "sec_excluded" || flagStatus == "manual_review" {
		return "ineligible"
	}
	return "eligible"
}
