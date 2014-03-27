package main

import (
	"dlib/dbus"
	"fmt"
)

const (
	pageGeneral = "General"
	pageIPv4    = "IPv4"
	pageIPv6    = "IPv6"
	page8021x   = "802.1xSecurity"
)

var supportedConnectionTypes = []string{
	NM_SETTING_WIRED_SETTING_NAME,
	NM_SETTING_WIRELESS_SETTING_NAME,
}

type ConnectionSession struct {
	objPath string

	data _ConnectionData

	connType string // TODO

	currentUUID string // TODO hide property
	HasChanged  bool

	currentPage string

	//前端只显示此列表中的字段,会跟随当前正在编辑的值而改变
	// TODO more documentation
	CurrentFields []string
	//返回当前page下错误的字段和对应的错误原因
	CurrentErrors []string
}

//所有字段值都为string，后端自行转换为需要的值后提供给NM

func NewConnectionSessionByCreate(connType string) (session *ConnectionSession, err error) {
	if !isStringInArray(connType, supportedConnectionTypes) {
		err = fmt.Errorf("connection type is out of support: %s", connType)
		LOGGER.Error(err)
		return
	}

	session = &ConnectionSession{}
	session.currentUUID = newUUID()
	session.objPath = fmt.Sprintf("/com/deepin/daemon/ConnectionSession/%s", randString(8))

	// TODO
	session.data = make(_ConnectionData)
	session.connType = connType

	// session.updatePropCurrentUUID(uuid)
	// session.updatePropHasChanged(true)

	// TODO
	// session.currentPage = session.getDefaultPage(connType)
	// session.updatePropCurrentFields()

	// TODO current errors
	return
}

func NewConnectionSessionByOpen(uuid string) (session *ConnectionSession, err error) {
	data := _Manager.getConnectionData(uuid)
	if data == nil {
		err = fmt.Errorf("counld not find connection with uuid equal %s", uuid)
		return
	}
	session = &ConnectionSession{}
	session.currentUUID = uuid
	session.objPath = fmt.Sprintf("/com/deepin/daemon/ConnectionSession/%s", randString(8))
	session.data = data
	return
}

// Save save current connection session.
func (session *ConnectionSession) Save() {
	// TODO
	if !session.HasChanged {
		dbus.UnInstallObject(session)
		return
	}
	dbus.UnInstallObject(session)
}

// Cancel cancel current connection session.
func (session *ConnectionSession) Cancel() {
	dbus.UnInstallObject(session)
}

//根据CurrentUUID返回此Connection支持的设置页面
func (session *ConnectionSession) ListPages() (pages []string) {
	// TODO
	switch session.connType {
	case NM_SETTING_WIRED_SETTING_NAME:
		pages = []string{
			pageGeneral,
			pageIPv4,
			pageIPv6,
			page8021x,
		}
	case NM_SETTING_WIRELESS_SETTING_NAME:
	}
	return
}

// get valid fields for target page
func (session *ConnectionSession) listFields(page string) (fields []string) {
	switch session.connType {
	case NM_SETTING_WIRED_SETTING_NAME:
		switch session.currentPage {
		case pageGeneral:
			fields = []string{"General_field1", "General_field2"}
		case pageIPv4:
			fields = []string{
				NM_SETTING_IP4_CONFIG_METHOD,
				NM_SETTING_IP4_CONFIG_DNS,
				NM_SETTING_IP4_CONFIG_DNS_SEARCH,
				NM_SETTING_IP4_CONFIG_ADDRESSES,
				NM_SETTING_IP4_CONFIG_ROUTES,
				NM_SETTING_IP4_CONFIG_IGNORE_AUTO_ROUTES,
				NM_SETTING_IP4_CONFIG_IGNORE_AUTO_DNS,
				NM_SETTING_IP4_CONFIG_DHCP_CLIENT_ID,
				NM_SETTING_IP4_CONFIG_DHCP_SEND_HOSTNAME,
				NM_SETTING_IP4_CONFIG_DHCP_HOSTNAME,
				NM_SETTING_IP4_CONFIG_NEVER_DEFAULT,
				NM_SETTING_IP4_CONFIG_MAY_FAIL,
			}
		case pageIPv6:
			fields = []string{"IPv6_field1", "IPv6_field2"}
		case page8021x:
			fields = []string{"802.1xSecurity_field1", "802.1xSecurity_field2"}
		}
	case NM_SETTING_WIRELESS_SETTING_NAME:
	}
	return
}

//设置/获得字段的值都受这里设置page的影响。
func (session *ConnectionSession) SwitchPage(page string) {
	// TODO HasChanged
	session.currentPage = page
	session.updatePropCurrentFields()
}

//比如获得当前链接支持的加密方式 EAP字段: TLS、MD5、FAST、PEAP
//获得ip设置方式 : Manual、Link-Local Only、Automatic(DHCP)
//获得当前可用mac地址(这种字段是有几个可选值但用户也可用手动输入一个其他值)
func (session *ConnectionSession) GetAvailableValue(key string) (values []string) {
	// TODO
	switch key {
	case NM_SETTING_IP4_CONFIG_METHOD:
		values = []string{
			NM_SETTING_IP4_CONFIG_METHOD_AUTO,
			NM_SETTING_IP4_CONFIG_METHOD_LINK_LOCAL,
			NM_SETTING_IP4_CONFIG_METHOD_MANUAL,
			NM_SETTING_IP4_CONFIG_METHOD_SHARED,
		}
	case NM_SETTING_IP4_CONFIG_DNS:
		values = []string{}
	}
	return
}

//仅仅调试使用，返回某个页面支持的字段。 因为字段如何安排(位置、我们是否要提供这个字段)是由前端决定的。
//*****在设计前端显示内容的时候和这个返回值关联很大*****
// DebugListFields return all fields of current page, only for debugging.
func (session *ConnectionSession) DebugListFields() []string {
	// TODO
	return session.listFields(session.currentPage)
}

// DebugConnectionTypes return all supported connection types, only for debugging.
func (session *ConnectionSession) DebugConnectionTypes() []string {
	return supportedConnectionTypes
}

// TODO
func (session *ConnectionSession) DebugGetKeyType(page, key string) ktypeDescription {
	return ktypeDescriptions[0]
}

// TODO panic error
func (*Manager) DebugListKeyTypes() (uint32, string) {
	LOGGER.Debug(ktypeDescriptions)
	return ktypeDescriptions[0].t, ktypeDescriptions[0].desc
}

//设置某个字段， 会影响CurrentFields属性，某些值会导致其他属性进入不可用状态
// TODO SetField
// func (session *ConnectionSession) SetField(key, value string) (err error) {
// 	page, ok := session.data[session.currentPage]
// 	if !ok {
// 		err = fmt.Errorf("invalid page: data[%s]", session.currentPage)
// 		return
// 	}

// 	varient, ok := page[key]
// 	if !ok {
// 		// field is not yet exits
// 		page[key] = dbus.MakeVariant(value)
// 		session.HasChanged = true
// 		return
// 	}

// 	// TODO oldValue, err := varientToString(varient)
// 	// oldValue, err := varientToString(varient)
// 	// if err != nil {
// 	// 	return
// 	// }
// 	// if oldValue == value {
// 	// 	return
// 	// }
// 	// page[key] = dbus.MakeVariant(value)
// 	// session.HasChanged = true

// 	// // TODO processing logic
// 	// session.updatePropCurrentFields()

// 	return
// }

// func (session *ConnectionSession) GetField(key string) (value string, err error) {
// 	page, ok := session.data[session.currentPage]
// 	if !ok {
// 		err = fmt.Errorf("invalid page: data[%s]", session.currentPage)
// 		return
// 	}

// 	// TODO
// 	// varient, ok := page[key]
// 	// if !ok {
// 	// 	err = fmt.Errorf("invalid field: data[%s][%s]", session.currentPage, key)
// 	// 	return
// 	// }

// 	// TODO
// 	// value, err = varientToString(varient)
// 	return
// }