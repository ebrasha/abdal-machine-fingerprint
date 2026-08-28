/*
 **********************************************************************
 * -------------------------------------------------------------------
 * Project Name : Abdal Machine Fingerprint
 * File Name    : machineid_linux.go
 * Programmer   : Ebrahim Shafiei (EbraSha)
 * Email        : Prof.Shafiei@Gmail.com
 * Created On   : 2026-08-28 18:13:17
 * Description  : Linux machine-id file reader for stable OS identity.
 * -------------------------------------------------------------------
 *
 * "Coding is an engaging and beloved hobby for me. I passionately and insatiably pursue knowledge in cybersecurity and programming."
 * – Ebrahim Shafiei
 *
 **********************************************************************
 */

//go:build linux

package platform

import "os"

const (
	dbusMachineIDPath = "/var/lib/dbus/machine-id"
	etcMachineIDPath  = "/etc/machine-id"
)

func readMachineID() (string, error) {
	id, err := readMachineIDFile(dbusMachineIDPath)
	if err != nil {
		id, err = readMachineIDFile(etcMachineIDPath)
	}
	if err != nil {
		return "", ErrNotFound
	}
	return id, nil
}

func readMachineIDFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	if len(data) == 0 {
		return "", ErrNotFound
	}
	return string(data), nil
}
