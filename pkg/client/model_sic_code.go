/*
 * Customers API
 *
 * Customers focuses on solving authentic identification of humans who are legally able to hold and transfer currency within the US. Primarily this project solves [Know Your Customer](https://en.wikipedia.org/wiki/Know_your_customer) (KYC), [Customer Identification Program](https://en.wikipedia.org/wiki/Customer_Identification_Program) (CIP), [Office of Foreign Asset Control](https://www.treasury.gov/about/organizational-structure/offices/Pages/Office-of-Foreign-Assets-Control.aspx) (OFAC) checks and verification workflows to comply with US federal law and ensure authentic transfers. Customers has an objective to be a service for detailed due diligence on individuals and companies for Financial Institutions and services in a modernized and extensible way.  Customer phone numbers and addresses are stored and partially used in KYC/OFAC validation. Arbitrary key/value pairs can be stored for a Customer. Documents and Disclaimers, and their acknowledgement are also stored under a Customer as they're accepted. Bank Accounts, which can be validated with micro-deposits currently, are stored under each Customer.  ![](https://raw.githubusercontent.com/adamdecaf/customers/create-accounts/docs/images/customer.png)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

// SicCode SIC Code (https://docs.google.com/spreadsheets/d/1erIdqoy60JwLAnpb91EfoJV5YrXDnbwSaA-aqcBlw48/edit#gid=1627409140)
type SicCode string

// List of SICCode
const (
	SICCODE__1   SicCode = "1"
	SICCODE__2   SicCode = "2"
	SICCODE__7   SicCode = "7"
	SICCODE__8_0 SicCode = "8.0"
	SICCODE__9_0 SicCode = "9.0"
	SICCODE__10  SicCode = "10"
	SICCODE__12  SicCode = "12"
	SICCODE__13  SicCode = "13"
	SICCODE__14  SicCode = "14"
	SICCODE__15  SicCode = "15"
	SICCODE__16  SicCode = "16"
	SICCODE__17  SicCode = "17"
	SICCODE__20  SicCode = "20"
	SICCODE__21  SicCode = "21"
	SICCODE__22  SicCode = "22"
	SICCODE__23  SicCode = "23"
	SICCODE__24  SicCode = "24"
	SICCODE__25  SicCode = "25"
	SICCODE__26  SicCode = "26"
	SICCODE__27  SicCode = "27"
	SICCODE__28  SicCode = "28"
	SICCODE__29  SicCode = "29"
	SICCODE__30  SicCode = "30"
	SICCODE__31  SicCode = "31"
	SICCODE__32  SicCode = "32"
	SICCODE__33  SicCode = "33"
	SICCODE__34  SicCode = "34"
	SICCODE__35  SicCode = "35"
	SICCODE__36  SicCode = "36"
	SICCODE__37  SicCode = "37"
	SICCODE__38  SicCode = "38"
	SICCODE__39  SicCode = "39"
	SICCODE__40  SicCode = "40"
	SICCODE__41  SicCode = "41"
	SICCODE__42  SicCode = "42"
	SICCODE__43  SicCode = "43"
	SICCODE__44  SicCode = "44"
	SICCODE__45  SicCode = "45"
	SICCODE__46  SicCode = "46"
	SICCODE__47  SicCode = "47"
	SICCODE__48  SicCode = "48"
	SICCODE__49  SicCode = "49"
	SICCODE__50  SicCode = "50"
	SICCODE__51  SicCode = "51"
	SICCODE__52  SicCode = "52"
	SICCODE__53  SicCode = "53"
	SICCODE__54  SicCode = "54"
	SICCODE__55  SicCode = "55"
	SICCODE__56  SicCode = "56"
	SICCODE__57  SicCode = "57"
	SICCODE__58  SicCode = "58"
	SICCODE__59  SicCode = "59"
	SICCODE__60  SicCode = "60"
	SICCODE__61  SicCode = "61"
	SICCODE__62  SicCode = "62"
	SICCODE__63  SicCode = "63"
	SICCODE__64  SicCode = "64"
	SICCODE__65  SicCode = "65"
	SICCODE__67  SicCode = "67"
	SICCODE__70  SicCode = "70"
	SICCODE__72  SicCode = "72"
	SICCODE__73  SicCode = "73"
	SICCODE__75  SicCode = "75"
	SICCODE__76  SicCode = "76"
	SICCODE__78  SicCode = "78"
	SICCODE__79  SicCode = "79"
	SICCODE__80  SicCode = "80"
	SICCODE__81  SicCode = "81"
	SICCODE__82  SicCode = "82"
	SICCODE__83  SicCode = "83"
	SICCODE__84  SicCode = "84"
	SICCODE__86  SicCode = "86"
	SICCODE__87  SicCode = "87"
	SICCODE__88  SicCode = "88"
	SICCODE__89  SicCode = "89"
	SICCODE__91  SicCode = "91"
	SICCODE__92  SicCode = "92"
	SICCODE__93  SicCode = "93"
	SICCODE__94  SicCode = "94"
	SICCODE__95  SicCode = "95"
	SICCODE__96  SicCode = "96"
	SICCODE__97  SicCode = "97"
	SICCODE__98  SicCode = "98"
	SICCODE__99  SicCode = "99"
)
