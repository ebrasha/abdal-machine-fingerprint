/*
 **********************************************************************
 * -------------------------------------------------------------------
 * Project Name : Abdal Machine Fingerprint
 * File Name    : constants.go
 * Programmer   : Ebrahim Shafiei (EbraSha)
 * Email        : Prof.Shafiei@Gmail.com
 * Created On   : 2026-08-28 19:52:22
 * Description  : Platform-specific constants for machine identity retrieval.
 * -------------------------------------------------------------------
 *
 * "Coding is an engaging and beloved hobby for me. I passionately and insatiably pursue knowledge in cybersecurity and programming."
 * – Ebrahim Shafiei
 *
 **********************************************************************
 */

package platform

// Linux machine identity file paths.
const (
	linuxDBusMachineIDPath = "/var/lib/dbus/machine-id"
	linuxEtcMachineIDPath  = "/etc/machine-id"
)

// Windows registry paths and keys for machine identity.
const (
	windowsRegistryCryptographyPath = `SOFTWARE\Microsoft\Cryptography`
	windowsMachineGuidKey           = "MachineGuid"
)
