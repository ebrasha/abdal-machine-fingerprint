/*
 **********************************************************************
 * -------------------------------------------------------------------
 * Project Name : Abdal Machine Fingerprint
 * File Name    : machineid_windows.go
 * Programmer   : Ebrahim Shafiei (EbraSha)
 * Email        : Prof.Shafiei@Gmail.com
 * Created On   : 2026-08-28 18:13:17
 * Description  : Windows MachineGuid registry reader.
 * -------------------------------------------------------------------
 *
 * "Coding is an engaging and beloved hobby for me. I passionately and insatiably pursue knowledge in cybersecurity and programming."
 * – Ebrahim Shafiei
 *
 **********************************************************************
 */

//go:build windows

package platform

import (
	"golang.org/x/sys/windows/registry"
)

const (
	registryPath = `SOFTWARE\Microsoft\Cryptography`
	registryKey  = "MachineGuid"
)

func readMachineID() (string, error) {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, registryPath, registry.QUERY_VALUE|registry.WOW64_64KEY)
	if err != nil {
		return "", ErrNotFound
	}
	defer k.Close()

	value, _, err := k.GetStringValue(registryKey)
	if err != nil {
		return "", ErrNotFound
	}
	if value == "" {
		return "", ErrNotFound
	}
	return value, nil
}
