/**
 * Copyright (c) 2014 Deepin, Inc.
 *               2014 Xu FaSheng
 *
 * Author:      Xu FaSheng <fasheng.xu@gmail.com>
 * Maintainer:  Xu FaSheng <fasheng.xu@gmail.com>
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, see <http://www.gnu.org/licenses/>.
 **/

package network

// #cgo pkg-config: libudev
// #include <stdlib.h>
// #include "utils_udev.h"
import "C"

import "unsafe"

func udevGetDeviceVendor(syspath string) (vendor string) {
	cSyspath := C.CString(syspath)
	defer C.free(unsafe.Pointer(cSyspath))

	cVendor := C.get_device_vendor(cSyspath)
	defer C.free(unsafe.Pointer(cVendor))
	vendor = C.GoString(cVendor)
	return
}

func udevIsUsbDevice(syspath string) bool {
	cSyspath := C.CString(syspath)
	defer C.free(unsafe.Pointer(cSyspath))

	ret := C.is_usb_device(cSyspath)
	if ret == 0 {
		return true
	}
	return false
}
