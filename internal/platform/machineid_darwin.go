/*
 **********************************************************************
 * -------------------------------------------------------------------
 * Project Name : Abdal Machine Fingerprint
 * File Name    : machineid_darwin.go
 * Programmer   : Ebrahim Shafiei (EbraSha)
 * Email        : Prof.Shafiei@Gmail.com
 * Created On   : 2026-08-28 18:13:17
 * Description  : macOS IOPlatformUUID reader via native IOKit interfaces.
 * -------------------------------------------------------------------
 *
 * "Coding is an engaging and beloved hobby for me. I passionately and insatiably pursue knowledge in cybersecurity and programming."
 * – Ebrahim Shafiei
 *
 **********************************************************************
 */

//go:build darwin

package platform

/*
#cgo LDFLAGS: -framework IOKit -framework CoreFoundation
#include <CoreFoundation/CoreFoundation.h>
#include <IOKit/IOKitLib.h>
#include <stdlib.h>

static char* read_io_platform_uuid(void) {
    io_registry_entry_t entry = IORegistryEntryFromPath(kIOMainPortDefault, "IOService:/");
    if (entry == MACH_PORT_NULL) {
        return NULL;
    }

    CFStringRef uuidRef = IORegistryEntryCreateCFProperty(
        entry,
        CFSTR(kIOPlatformUUIDKey),
        kCFAllocatorDefault,
        0
    );
    IOObjectRelease(entry);
    if (uuidRef == NULL) {
        return NULL;
    }

    char buffer[64];
    if (!CFStringGetCString(uuidRef, buffer, sizeof(buffer), kCFStringEncodingUTF8)) {
        CFRelease(uuidRef);
        return NULL;
    }
    CFRelease(uuidRef);
    return strdup(buffer);
}
*/
import "C"

import "unsafe"

func readMachineID() (string, error) {
	cstr := C.read_io_platform_uuid()
	if cstr == nil {
		return "", ErrNotFound
	}
	defer C.free(unsafe.Pointer(cstr))

	value := C.GoString(cstr)
	if value == "" {
		return "", ErrNotFound
	}
	return value, nil
}
