package staticdata

import (
	"fmt"
	"strings"
)

// GetPricePointID returns the price point ID for a given territory and price.
// Returns an error if not found.
func GetPricePointID(territory string, price string) (string, error) {
	key := fmt.Sprintf("%s-%s", strings.ToUpper(territory), price)
	if id, ok := pricingMatrix[key]; ok {
		return id, nil
	}
	return "", fmt.Errorf("price point not found for territory %s and price %s in static data", territory, price)
}

// pricingMatrix maps "Territory-Price" to "PricePointID".
// IDs are standard App Store Connect price tier IDs (e.g. "1", "2", "1000").
var pricingMatrix = map[string]string{
	// --- USA (United States) ---
	"USA-0":    "1000", // Free
	"USA-0.00": "1000", // Free
	"USA-0.99": "1",    // Tier 1
	"USA-1.99": "2",    // Tier 2
	"USA-2.99": "3",    // Tier 3
	"USA-3.99": "4",    // Tier 4
	"USA-4.99": "5",    // Tier 5
	"USA-5.99": "6",    // Tier 6
	"USA-6.99": "7",    // Tier 7
	"USA-7.99": "8",    // Tier 8
	"USA-8.99": "9",    // Tier 9
	"USA-9.99": "10",   // Tier 10

	// --- CHN (China) ---
	"CHN-0":     "1000",
	"CHN-0.00":  "1000",
	"CHN-1.00":  "1", // Alternate Tier 1
	"CHN-6.00":  "1", // Standard Tier 1
	"CHN-8.00":  "1", // Alternate
	"CHN-12.00": "2",
	"CHN-18.00": "3",
	"CHN-25.00": "4",
	"CHN-30.00": "5",
	"CHN-40.00": "6",
	"CHN-45.00": "7",
	"CHN-50.00": "8",
	"CHN-60.00": "9",
	"CHN-68.00": "10",

	// --- JPN (Japan) ---
	"JPN-0":    "1000",
	"JPN-0.00": "1000",
	"JPN-160":  "1",
	"JPN-320":  "2",
	"JPN-480":  "3",
	"JPN-650":  "4",
	"JPN-800":  "5",
	"JPN-1000": "6",
	"JPN-1100": "7",
	"JPN-1200": "8",
	"JPN-1500": "9",
	"JPN-1600": "10",

	// --- GBR (United Kingdom) ---
	"GBR-0":    "1000",
	"GBR-0.00": "1000",
	"GBR-0.99": "1",
	"GBR-1.99": "2",
	"GBR-2.99": "3",
	"GBR-3.99": "4",
	"GBR-4.99": "5",
	"GBR-5.99": "6",
	"GBR-6.99": "7",
	"GBR-7.99": "8",
	"GBR-8.99": "9",
	"GBR-9.99": "10",

	// --- EUR (Eurozone) ---
	// Using DEU/FRA/ESP/ITA as keys for EUR pricing
	"DEU-0":    "1000",
	"DEU-0.00": "1000",
	"DEU-0.99": "1",
	"DEU-1.99": "2",
	"DEU-2.99": "3",
	"DEU-3.99": "4",
	"DEU-4.99": "5",
	"DEU-5.99": "6",
	"DEU-6.99": "7",
	"DEU-7.99": "8",
	"DEU-8.99": "9",
	"DEU-9.99": "10",

	"FRA-0":    "1000",
	"FRA-0.00": "1000",
	"FRA-0.99": "1",
	"FRA-1.99": "2",
	"FRA-2.99": "3",
	"FRA-3.99": "4",
	"FRA-4.99": "5",
	"FRA-5.99": "6",
	"FRA-6.99": "7",
	"FRA-7.99": "8",
	"FRA-8.99": "9",
	"FRA-9.99": "10",

	// --- CAN (Canada) ---
	"CAN-0":    "1000",
	"CAN-0.00": "1000",
	"CAN-1.29": "1",
	"CAN-2.79": "2",
	"CAN-3.99": "3",
	"CAN-5.49": "4",
	"CAN-6.99": "5",
	"CAN-8.49": "6",
	"CAN-9.99": "7",
	"CAN-11.49": "8",
	"CAN-12.99": "9",
	"CAN-14.49": "10",

	// --- AUS (Australia) ---
	"AUS-0":    "1000",
	"AUS-0.00": "1000",
	"AUS-1.49": "1",
	"AUS-2.99": "2",
	"AUS-4.49": "3",
	"AUS-5.99": "4",
	"AUS-7.99": "5",
	"AUS-8.99": "6",
	"AUS-10.99": "7",
	"AUS-12.99": "8",
	"AUS-13.99": "9",
	"AUS-14.99": "10",

	// --- IND (India) ---
	"IND-0":     "1000",
	"IND-0.00":  "1000",
	"IND-99.00": "1",
	"IND-199.00": "2",
	"IND-299.00": "3",
	"IND-399.00": "4",
	"IND-499.00": "5",
	"IND-599.00": "6",
	"IND-699.00": "7",
	"IND-799.00": "8",
	"IND-899.00": "9",
	"IND-999.00": "10",

	// --- KOR (South Korea) ---
	"KOR-0":      "1000",
	"KOR-0.00":   "1000",
	"KOR-1500":   "1",
	"KOR-3300":   "2",
	"KOR-4400":   "3",
	"KOR-6000":   "4",
	"KOR-7500":   "5",
	"KOR-9000":   "6",
	"KOR-11000":  "7",
	"KOR-12000":  "8",
	"KOR-14000":  "9",
	"KOR-15000":  "10",

	// --- HKG (Hong Kong) ---
	"HKG-0":     "1000",
	"HKG-0.00":  "1000",
	"HKG-8.00":  "1",
	"HKG-15.00": "2",
	"HKG-23.00": "3",
	"HKG-33.00": "4",
	"HKG-38.00": "5",
	"HKG-48.00": "6",
	"HKG-58.00": "7",
	"HKG-68.00": "8",
	"HKG-78.00": "9",
	"HKG-88.00": "10",

	// --- TWN (Taiwan) ---
	"TWN-0":      "1000",
	"TWN-0.00":   "1000",
	"TWN-33.00":  "1",
	"TWN-70.00":  "2",
	"TWN-100.00": "3",
	"TWN-130.00": "4",
	"TWN-170.00": "5",
	"TWN-200.00": "6",
	"TWN-230.00": "7",
	"TWN-270.00": "8",
	"TWN-290.00": "9",
	"TWN-330.00": "10",
}
