package inputdevices

import (
	"fmt"
	"gir/gio-2.0"
	"io/ioutil"
	"os"
	"os/exec"
	"pkg.deepin.io/dde/api/dxinput"
	dxutils "pkg.deepin.io/dde/api/dxinput/utils"
	"pkg.deepin.io/lib/dbus/property"
	dutils "pkg.deepin.io/lib/utils"
	"strconv"
	"strings"
)

const (
	tpadSchema = "com.deepin.dde.touchpad"

	tpadKeyEnabled       = "touchpad-enabled"
	tpadKeyLeftHanded    = "left-handed"
	tpadKeyWhileTyping   = "disable-while-typing"
	tpadKeyNaturalScroll = "natural-scroll"
	tpadKeyEdgeScroll    = "edge-scroll-enabled"
	tpadKeyHorizScroll   = "horiz-scroll-enabled"
	tpadKeyVertScroll    = "vert-scroll-enabled"
	tpadKeyAcceleration  = "motion-acceleration"
	tpadKeyThreshold     = "motion-threshold"
	tpadKeyTapClick      = "tap-to-click"
	tpadKeyScrollDelta   = "delta-scroll"
)

const (
	syndaemonPidFile = "/tmp/syndaemon.pid"
)

type Touchpad struct {
	TPadEnable      *property.GSettingsBoolProperty `access:"readwrite"`
	LeftHanded      *property.GSettingsBoolProperty `access:"readwrite"`
	DisableIfTyping *property.GSettingsBoolProperty `access:"readwrite"`
	NaturalScroll   *property.GSettingsBoolProperty `access:"readwrite"`
	EdgeScroll      *property.GSettingsBoolProperty `access:"readwrite"`
	HorizScroll     *property.GSettingsBoolProperty `access:"readwrite"`
	VertScroll      *property.GSettingsBoolProperty `access:"readwrite"`
	TapClick        *property.GSettingsBoolProperty `access:"readwrite"`

	MotionAcceleration *property.GSettingsFloatProperty `access:"readwrite"`
	MotionThreshold    *property.GSettingsFloatProperty `access:"readwrite"`

	DoubleClick   *property.GSettingsIntProperty `access:"readwrite"`
	DragThreshold *property.GSettingsIntProperty `access:"readwrite"`
	DeltaScroll   *property.GSettingsIntProperty `access:"readwrite"`

	Exist      bool
	DeviceList dxutils.DeviceInfos

	dxTPads      map[int32]*dxinput.Touchpad
	setting      *gio.Settings
	mouseSetting *gio.Settings
	synProcess   *os.Process
}

var _tpad *Touchpad

func getTouchpad() *Touchpad {
	if _tpad == nil {
		_tpad = NewTouchpad()

		_tpad.init()
		_tpad.handleGSettings()
	}

	return _tpad
}

func NewTouchpad() *Touchpad {
	var tpad = new(Touchpad)

	tpad.setting = gio.NewSettings(tpadSchema)

	tpad.TPadEnable = property.NewGSettingsBoolProperty(
		tpad, "TPadEnable",
		tpad.setting, tpadKeyEnabled)
	tpad.LeftHanded = property.NewGSettingsBoolProperty(
		tpad, "LeftHanded",
		tpad.setting, tpadKeyLeftHanded)
	tpad.DisableIfTyping = property.NewGSettingsBoolProperty(
		tpad, "DisableIfTyping",
		tpad.setting, tpadKeyWhileTyping)
	tpad.NaturalScroll = property.NewGSettingsBoolProperty(
		tpad, "NaturalScroll",
		tpad.setting, tpadKeyNaturalScroll)
	tpad.EdgeScroll = property.NewGSettingsBoolProperty(
		tpad, "EdgeScroll",
		tpad.setting, tpadKeyEdgeScroll)
	tpad.VertScroll = property.NewGSettingsBoolProperty(
		tpad, "VertScroll",
		tpad.setting, tpadKeyVertScroll)
	tpad.HorizScroll = property.NewGSettingsBoolProperty(
		tpad, "HorizScroll",
		tpad.setting, tpadKeyHorizScroll)
	tpad.TapClick = property.NewGSettingsBoolProperty(
		tpad, "TapClick",
		tpad.setting, tpadKeyTapClick)

	tpad.MotionAcceleration = property.NewGSettingsFloatProperty(
		tpad, "MotionAcceleration",
		tpad.setting, tpadKeyAcceleration)
	tpad.MotionThreshold = property.NewGSettingsFloatProperty(
		tpad, "MotionThreshold",
		tpad.setting, tpadKeyThreshold)

	tpad.DeltaScroll = property.NewGSettingsIntProperty(
		tpad, "DeltaScroll",
		tpad.setting, tpadKeyScrollDelta)

	tpad.mouseSetting = gio.NewSettings(mouseSchema)
	tpad.DoubleClick = property.NewGSettingsIntProperty(
		tpad, "DoubleClick",
		tpad.mouseSetting, mouseKeyDoubleClick)
	tpad.DragThreshold = property.NewGSettingsIntProperty(
		tpad, "DragThreshold",
		tpad.mouseSetting, mouseKeyDragThreshold)

	tpad.updateDeviceList()
	tpad.dxTPads = make(map[int32]*dxinput.Touchpad)
	tpad.updateDXTpads()

	return tpad
}

func (tpad *Touchpad) init() {
	if !tpad.Exist {
		return
	}

	tpad.enable(tpad.TPadEnable.Get())
	tpad.enableLeftHanded()
	tpad.enableNaturalScroll()
	tpad.enableEdgeScroll()
	tpad.enableTapToClick()
	tpad.enableTwoFingerScroll()
	tpad.motionAcceleration()
	tpad.motionThreshold()
	tpad.disableWhileTyping()
}

func (tpad *Touchpad) handleDeviceChanged() {
	tpad.updateDeviceList()
	tpad.updateDXTpads()
	tpad.init()
}

func (tpad *Touchpad) updateDeviceList() {
	tpad.DeviceList = getTPadInfos(false)
	if len(tpad.DeviceList) == 0 {
		tpad.setPropExist(false)
	} else {
		tpad.setPropExist(true)
	}
}

func (tpad *Touchpad) updateDXTpads() {
	for _, info := range tpad.DeviceList {
		_, ok := tpad.dxTPads[info.Id]
		if ok {
			continue
		}

		dxt, err := dxinput.NewTouchpad(info.Id)
		if err != nil {
			logger.Warning(err)
			continue
		}
		tpad.dxTPads[info.Id] = dxt
	}
}

func (tpad *Touchpad) enable(enabled bool) {
	for _, v := range tpad.dxTPads {
		err := v.Enable(enabled)
		if err != nil {
			logger.Warningf("Enable '%v - %v' failed: %v",
				v.Id, v.Name, err)
		}
	}
}

func (tpad *Touchpad) enableLeftHanded() {
	for _, v := range tpad.dxTPads {
		err := v.EnableLeftHanded(tpad.LeftHanded.Get())
		if err != nil {
			logger.Debugf("Enable left handed '%v - %v' failed: %v",
				v.Id, v.Name, err)
		}
	}
}

func (tpad *Touchpad) enableNaturalScroll() {
	for _, v := range tpad.dxTPads {
		err := v.EnableNaturalScroll(tpad.NaturalScroll.Get())
		if err != nil {
			logger.Debugf("Enable natural scroll '%v - %v' failed: %v",
				v.Id, v.Name, err)
		}
	}
}

func (tpad *Touchpad) setScrollDistance() {
	for _, v := range tpad.dxTPads {
		err := v.SetScrollDistance(tpad.DeltaScroll.Get(), tpad.DeltaScroll.Get())
		if err != nil {
			logger.Debugf("Set natural scroll distance '%v - %v' failed: %v",
				v.Id, v.Name, err)
		}
	}
}

func (tpad *Touchpad) enableEdgeScroll() {
	for _, v := range tpad.dxTPads {
		err := v.EnableEdgeScroll(tpad.EdgeScroll.Get())
		if err != nil {
			logger.Debugf("Enable edge scroll '%v - %v' failed: %v",
				v.Id, v.Name, err)
		}
	}
}

func (tpad *Touchpad) enableTwoFingerScroll() {
	for _, v := range tpad.dxTPads {
		err := v.EnableTwoFingerScroll(tpad.VertScroll.Get(),
			tpad.HorizScroll.Get())
		if err != nil {
			logger.Debugf("Enable two-finger scroll '%v - %v' failed: %v",
				v.Id, v.Name, err)
		}
	}
}

func (tpad *Touchpad) enableTapToClick() {
	for _, v := range tpad.dxTPads {
		err := v.EnableTapToClick(tpad.TapClick.Get())
		if err != nil {
			logger.Debugf("Enable tap to click '%v - %v' failed: %v",
				v.Id, v.Name, err)
		}
	}
}

func (tpad *Touchpad) motionAcceleration() {
	for _, v := range tpad.dxTPads {
		err := v.SetMotionAcceleration(
			float32(tpad.MotionAcceleration.Get()))
		if err != nil {
			logger.Debugf("Set acceleration for '%d - %v' failed: %v",
				v.Id, v.Name, err)
		}
	}
}

func (tpad *Touchpad) motionThreshold() {
	for _, v := range tpad.dxTPads {
		err := v.SetMotionThreshold(float32(tpad.MotionThreshold.Get()))
		if err != nil {
			logger.Debugf("Set threshold for '%d - %v' failed: %v",
				v.Id, v.Name, err)
		}
	}
}

func (tpad *Touchpad) doubleClick() {
	xsSetInt32(xsPropDoubleClick, tpad.DoubleClick.Get())
}

func (tpad *Touchpad) dragThreshold() {
	xsSetInt32(xsPropDragThres, tpad.DragThreshold.Get())
}

func (tpad *Touchpad) disableWhileTyping() {
	if !tpad.Exist {
		return
	}

	if tpad.DisableIfTyping.Get() {
		tpad.startSyndaemon()
	} else {
		tpad.stopSyndaemon()
	}
}

func (tpad *Touchpad) startSyndaemon() {
	if isSyndaemonExist(syndaemonPidFile) {
		logger.Debug("Syndaemon has running")
		return
	}

	var cmd = exec.Command("/bin/sh", "-c", "syndaemon -i 1 -K -t")
	err := cmd.Start()
	if err != nil {
		os.Remove(syndaemonPidFile)
		logger.Debug("[disableWhileTyping] start syndaemon failed:", err)
		return
	}
	tpad.synProcess = cmd.Process
	content := fmt.Sprintf("%v", tpad.synProcess.Pid)
	ioutil.WriteFile(syndaemonPidFile, []byte(content), 0777)
}

func (tpad *Touchpad) stopSyndaemon() {
	if tpad.synProcess == nil {
		return
	}

	tpad.synProcess.Kill()
	tpad.synProcess = nil
	os.Remove(syndaemonPidFile)
}

func isSyndaemonExist(pidFile string) bool {
	if !dutils.IsFileExist(pidFile) {
		return false
	}

	context, err := ioutil.ReadFile(pidFile)
	if err != nil {
		return false
	}

	pid, err := strconv.ParseInt(strings.TrimSpace(string(context)), 10, 64)
	if err != nil {
		return false
	}
	var file = fmt.Sprintf("/proc/%v/cmdline", pid)
	return isProcessExist(file, "syndaemon")
}

func isProcessExist(file, name string) bool {
	context, err := ioutil.ReadFile(file)
	if err != nil {
		return false
	}

	return strings.Contains(string(context), name)
}
