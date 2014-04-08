package main

import "dbus/org/freedesktop/upower"

const (
	UPOWER_BUS_NAME = "org.freedesktop.UPower"
)

const (
	//defined at http://upower.freedesktop.org/docs/Device.html#Device:Type
	DeviceTypeUnknow    = 0
	DeviceTypeLinePower = 1
	DeviceTypeBattery   = 2
	DeviceTypeUps       = 3
	DeviceTypeMonitor   = 4
	DeviceTypeMouse     = 5
	DeviceTypeKeyboard  = 6
	DeviceTypePda       = 7
	DeviceTypePhone     = 8
)

const (
	//defined at http://upower.freedesktop.org/docs/Device.html#Device:State
	BatteryStateUnknown          = 0
	BatteryStateCharging         = 1
	BatteryStateDischarging      = 2
	BatteryStateEmpty            = 3
	BatteryStateFullyCharged     = 4
	BatteryStatePendingCharge    = 5
	BatteryStatePendingDischarge = 6
)

func (p *Power) refreshUpower(up *upower.Upower) {
	p.setPropOnBattery(up.OnBattery.Get())

	p.setPropLidIsPresent(up.LidIsPresent.Get())

	if dev := getBattery(); dev != nil {
		p.setPropBatteryIsPresent(dev.IsPresent.Get())
		p.setPropBatteryState(dev.State.Get())
		p.setPropBatteryPercentage(dev.Percentage.Get())
	} else {
		p.setPropBatteryIsPresent(false)
	}
	//TODO: handle lowe battery
}

func (p *Power) initUpower() {
	up, err := upower.NewUpower(UPOWER_BUS_NAME, "/org/freedesktop/UPower")
	if err != nil {
		LOGGER.Error("Can't build org.freedesktop.UPower:", err)
	} else {
		up.ConnectChanged(func() {
			p.refreshUpower(up)
		})
	}
	p.refreshUpower(up)
}
func getBattery() *upower.Device {
	up, err := upower.NewUpower(UPOWER_BUS_NAME, "/org/freedesktop/UPower")
	if err != nil {
		LOGGER.Error("Can't build org.freedesktop.UPower:", err)
	}
	devs, err := up.EnumerateDevices()
	if err != nil {
		LOGGER.Error("Can't EnumerateDevices", err)
	}
	for _, path := range devs {
		dev, err := upower.NewDevice(UPOWER_BUS_NAME, path)
		if err == nil && dev.Type.Get() == DeviceTypeBattery {
			return dev
		}
	}
	return nil
}