package appearance

import (
	"gir/gio-2.0"
	"pkg.deepin.io/dde/daemon/appearance/background"
	"time"
)

func (m *Manager) listenGSettingChanged() {
	m.setting.Connect("changed::theme", func(s *gio.Settings, key string) {
		m.doSetDTheme(m.setting.GetString(key))
	})
	m.setting.GetString(gsKeyTheme)

	m.setting.Connect("changed::font-size", func(s *gio.Settings, key string) {
		m.doSetFontSize(m.setting.GetInt(key))
	})
	m.setting.GetInt(gsKeyFontSize)

	m.listenBgGsettings()
}

func (m *Manager) listenBgGsettings() {
	m.wrapBgSetting.Connect("changed::picture-uri", func(s *gio.Settings, key string) {
		uri := m.wrapBgSetting.GetString(gsKeyBackground)
		logger.Debug("[Wrap background] changed:", key, uri)
		err := m.doSetBackground(uri)
		if err != nil {
			logger.Debugf("[Wrap background] set '%s' failed: %s", uri, err)
			return
		}
		logger.Debug("[Wrap background] changed OVER ENDDDDDDDDDDD:", key, uri)
	})
	m.wrapBgSetting.GetString(gsKeyBackground)

	if m.gnomeBgSetting == nil {
		return
	}
	m.gnomeBgSetting.Connect("changed::picture-uri", func(s *gio.Settings, key string) {
		// Wait for file copy finished
		time.Sleep(time.Millisecond * 500)
		uri := m.gnomeBgSetting.GetString(gsKeyBackground)
		logger.Debug("[Gnome background] sync wrap bg:", uri, m.wrapBgSetting.GetString(gsKeyBackground))
		if uri == m.wrapBgSetting.GetString(gsKeyBackground) {
			return
		}
		if !background.IsBackgroundFile(uri) {
			logger.Debugf("[Gnome background] Invalid background file '%v'", uri)
			return
		}

		m.wrapBgSetting.SetString(gsKeyBackground, uri)
		logger.Debug("[Gnome background] sync wrap bg OVER ENDDDDDDDD:", uri, m.wrapBgSetting.GetString(gsKeyBackground))
	})
	m.gnomeBgSetting.GetString(gsKeyBackground)
}
