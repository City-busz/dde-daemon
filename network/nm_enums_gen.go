/* Generated data (by glib-mkenums) */

package network

/* enumerations from "/tmp/nm-client.h" */

// NMClientPermission
const (
	NM_CLIENT_PERMISSION_NONE                     = 0
	NM_CLIENT_PERMISSION_ENABLE_DISABLE_NETWORK   = 1
	NM_CLIENT_PERMISSION_ENABLE_DISABLE_WIFI      = 2
	NM_CLIENT_PERMISSION_ENABLE_DISABLE_WWAN      = 3
	NM_CLIENT_PERMISSION_ENABLE_DISABLE_WIMAX     = 4
	NM_CLIENT_PERMISSION_SLEEP_WAKE               = 5
	NM_CLIENT_PERMISSION_NETWORK_CONTROL          = 6
	NM_CLIENT_PERMISSION_WIFI_SHARE_PROTECTED     = 7
	NM_CLIENT_PERMISSION_WIFI_SHARE_OPEN          = 8
	NM_CLIENT_PERMISSION_SETTINGS_MODIFY_SYSTEM   = 9
	NM_CLIENT_PERMISSION_SETTINGS_MODIFY_OWN      = 10
	NM_CLIENT_PERMISSION_SETTINGS_MODIFY_HOSTNAME = 11
)

// NMClientPermissionResult
const (
	NM_CLIENT_PERMISSION_RESULT_UNKNOWN = 0
	NM_CLIENT_PERMISSION_RESULT_YES     = 1
	NM_CLIENT_PERMISSION_RESULT_AUTH    = 2
	NM_CLIENT_PERMISSION_RESULT_NO      = 3
)

// NMClientError
const (
	NM_CLIENT_ERROR_FAILED                 = 0
	NM_CLIENT_ERROR_MANAGER_NOT_RUNNING    = 1
	NM_CLIENT_ERROR_OBJECT_CREATION_FAILED = 2
)

/* enumerations from "/usr/include/libnm/nm-connection.h" */

// NMConnectionSerializationFlags
const (
	NM_CONNECTION_SERIALIZE_ALL          = 0
	NM_CONNECTION_SERIALIZE_NO_SECRETS   = 1
	NM_CONNECTION_SERIALIZE_ONLY_SECRETS = 2
)

/* enumerations from "/tmp/nm-dbus-interface.h" */

// NMState
const (
	NM_STATE_UNKNOWN          = 0
	NM_STATE_ASLEEP           = 10
	NM_STATE_DISCONNECTED     = 20
	NM_STATE_DISCONNECTING    = 30
	NM_STATE_CONNECTING       = 40
	NM_STATE_CONNECTED_LOCAL  = 50
	NM_STATE_CONNECTED_SITE   = 60
	NM_STATE_CONNECTED_GLOBAL = 70
)

// NMConnectivityState
const (
	NM_CONNECTIVITY_UNKNOWN = 0
	NM_CONNECTIVITY_NONE    = 1
	NM_CONNECTIVITY_PORTAL  = 2
	NM_CONNECTIVITY_LIMITED = 3
	NM_CONNECTIVITY_FULL    = 4
)

// NMDeviceType
const (
	NM_DEVICE_TYPE_UNKNOWN    = 0
	NM_DEVICE_TYPE_ETHERNET   = 1
	NM_DEVICE_TYPE_WIFI       = 2
	NM_DEVICE_TYPE_UNUSED1    = 3
	NM_DEVICE_TYPE_UNUSED2    = 4
	NM_DEVICE_TYPE_BT         = 5
	NM_DEVICE_TYPE_OLPC_MESH  = 6
	NM_DEVICE_TYPE_WIMAX      = 7
	NM_DEVICE_TYPE_MODEM      = 8
	NM_DEVICE_TYPE_INFINIBAND = 9
	NM_DEVICE_TYPE_BOND       = 10
	NM_DEVICE_TYPE_VLAN       = 11
	NM_DEVICE_TYPE_ADSL       = 12
	NM_DEVICE_TYPE_BRIDGE     = 13
	NM_DEVICE_TYPE_GENERIC    = 14
	NM_DEVICE_TYPE_TEAM       = 15
)

// NMDeviceCapabilities
const (
	NM_DEVICE_CAP_NONE           = 0
	NM_DEVICE_CAP_NM_SUPPORTED   = 1
	NM_DEVICE_CAP_CARRIER_DETECT = 2
	NM_DEVICE_CAP_IS_SOFTWARE    = 4
)

// NMDeviceWifiCapabilities
const (
	NM_WIFI_DEVICE_CAP_NONE          = 0
	NM_WIFI_DEVICE_CAP_CIPHER_WEP40  = 1
	NM_WIFI_DEVICE_CAP_CIPHER_WEP104 = 2
	NM_WIFI_DEVICE_CAP_CIPHER_TKIP   = 4
	NM_WIFI_DEVICE_CAP_CIPHER_CCMP   = 8
	NM_WIFI_DEVICE_CAP_WPA           = 16
	NM_WIFI_DEVICE_CAP_RSN           = 32
	NM_WIFI_DEVICE_CAP_AP            = 64
	NM_WIFI_DEVICE_CAP_ADHOC         = 128
	NM_WIFI_DEVICE_CAP_FREQ_VALID    = 256
	NM_WIFI_DEVICE_CAP_FREQ_2GHZ     = 512
	NM_WIFI_DEVICE_CAP_FREQ_5GHZ     = 1024
)

// NM80211ApFlags
const (
	NM_802_11_AP_FLAGS_NONE    = 0
	NM_802_11_AP_FLAGS_PRIVACY = 1
)

// NM80211ApSecurityFlags
const (
	NM_802_11_AP_SEC_NONE            = 0
	NM_802_11_AP_SEC_PAIR_WEP40      = 1
	NM_802_11_AP_SEC_PAIR_WEP104     = 2
	NM_802_11_AP_SEC_PAIR_TKIP       = 4
	NM_802_11_AP_SEC_PAIR_CCMP       = 8
	NM_802_11_AP_SEC_GROUP_WEP40     = 16
	NM_802_11_AP_SEC_GROUP_WEP104    = 32
	NM_802_11_AP_SEC_GROUP_TKIP      = 64
	NM_802_11_AP_SEC_GROUP_CCMP      = 128
	NM_802_11_AP_SEC_KEY_MGMT_PSK    = 256
	NM_802_11_AP_SEC_KEY_MGMT_802_1X = 512
)

// NM80211Mode
const (
	NM_802_11_MODE_UNKNOWN = 0
	NM_802_11_MODE_ADHOC   = 1
	NM_802_11_MODE_INFRA   = 2
	NM_802_11_MODE_AP      = 3
)

// NMBluetoothCapabilities
const (
	NM_BT_CAPABILITY_NONE = 0
	NM_BT_CAPABILITY_DUN  = 1
	NM_BT_CAPABILITY_NAP  = 2
)

// NMDeviceModemCapabilities
const (
	NM_DEVICE_MODEM_CAPABILITY_NONE      = 0
	NM_DEVICE_MODEM_CAPABILITY_POTS      = 1
	NM_DEVICE_MODEM_CAPABILITY_CDMA_EVDO = 2
	NM_DEVICE_MODEM_CAPABILITY_GSM_UMTS  = 4
	NM_DEVICE_MODEM_CAPABILITY_LTE       = 8
)

// NMDeviceState
const (
	NM_DEVICE_STATE_UNKNOWN      = 0
	NM_DEVICE_STATE_UNMANAGED    = 10
	NM_DEVICE_STATE_UNAVAILABLE  = 20
	NM_DEVICE_STATE_DISCONNECTED = 30
	NM_DEVICE_STATE_PREPARE      = 40
	NM_DEVICE_STATE_CONFIG       = 50
	NM_DEVICE_STATE_NEED_AUTH    = 60
	NM_DEVICE_STATE_IP_CONFIG    = 70
	NM_DEVICE_STATE_IP_CHECK     = 80
	NM_DEVICE_STATE_SECONDARIES  = 90
	NM_DEVICE_STATE_ACTIVATED    = 100
	NM_DEVICE_STATE_DEACTIVATING = 110
	NM_DEVICE_STATE_FAILED       = 120
)

// NMDeviceStateReason
const (
	NM_DEVICE_STATE_REASON_NONE                           = 0
	NM_DEVICE_STATE_REASON_UNKNOWN                        = 1
	NM_DEVICE_STATE_REASON_NOW_MANAGED                    = 2
	NM_DEVICE_STATE_REASON_NOW_UNMANAGED                  = 3
	NM_DEVICE_STATE_REASON_CONFIG_FAILED                  = 4
	NM_DEVICE_STATE_REASON_IP_CONFIG_UNAVAILABLE          = 5
	NM_DEVICE_STATE_REASON_IP_CONFIG_EXPIRED              = 6
	NM_DEVICE_STATE_REASON_NO_SECRETS                     = 7
	NM_DEVICE_STATE_REASON_SUPPLICANT_DISCONNECT          = 8
	NM_DEVICE_STATE_REASON_SUPPLICANT_CONFIG_FAILED       = 9
	NM_DEVICE_STATE_REASON_SUPPLICANT_FAILED              = 10
	NM_DEVICE_STATE_REASON_SUPPLICANT_TIMEOUT             = 11
	NM_DEVICE_STATE_REASON_PPP_START_FAILED               = 12
	NM_DEVICE_STATE_REASON_PPP_DISCONNECT                 = 13
	NM_DEVICE_STATE_REASON_PPP_FAILED                     = 14
	NM_DEVICE_STATE_REASON_DHCP_START_FAILED              = 15
	NM_DEVICE_STATE_REASON_DHCP_ERROR                     = 16
	NM_DEVICE_STATE_REASON_DHCP_FAILED                    = 17
	NM_DEVICE_STATE_REASON_SHARED_START_FAILED            = 18
	NM_DEVICE_STATE_REASON_SHARED_FAILED                  = 19
	NM_DEVICE_STATE_REASON_AUTOIP_START_FAILED            = 20
	NM_DEVICE_STATE_REASON_AUTOIP_ERROR                   = 21
	NM_DEVICE_STATE_REASON_AUTOIP_FAILED                  = 22
	NM_DEVICE_STATE_REASON_MODEM_BUSY                     = 23
	NM_DEVICE_STATE_REASON_MODEM_NO_DIAL_TONE             = 24
	NM_DEVICE_STATE_REASON_MODEM_NO_CARRIER               = 25
	NM_DEVICE_STATE_REASON_MODEM_DIAL_TIMEOUT             = 26
	NM_DEVICE_STATE_REASON_MODEM_DIAL_FAILED              = 27
	NM_DEVICE_STATE_REASON_MODEM_INIT_FAILED              = 28
	NM_DEVICE_STATE_REASON_GSM_APN_FAILED                 = 29
	NM_DEVICE_STATE_REASON_GSM_REGISTRATION_NOT_SEARCHING = 30
	NM_DEVICE_STATE_REASON_GSM_REGISTRATION_DENIED        = 31
	NM_DEVICE_STATE_REASON_GSM_REGISTRATION_TIMEOUT       = 32
	NM_DEVICE_STATE_REASON_GSM_REGISTRATION_FAILED        = 33
	NM_DEVICE_STATE_REASON_GSM_PIN_CHECK_FAILED           = 34
	NM_DEVICE_STATE_REASON_FIRMWARE_MISSING               = 35
	NM_DEVICE_STATE_REASON_REMOVED                        = 36
	NM_DEVICE_STATE_REASON_SLEEPING                       = 37
	NM_DEVICE_STATE_REASON_CONNECTION_REMOVED             = 38
	NM_DEVICE_STATE_REASON_USER_REQUESTED                 = 39
	NM_DEVICE_STATE_REASON_CARRIER                        = 40
	NM_DEVICE_STATE_REASON_CONNECTION_ASSUMED             = 41
	NM_DEVICE_STATE_REASON_SUPPLICANT_AVAILABLE           = 42
	NM_DEVICE_STATE_REASON_MODEM_NOT_FOUND                = 43
	NM_DEVICE_STATE_REASON_BT_FAILED                      = 44
	NM_DEVICE_STATE_REASON_GSM_SIM_NOT_INSERTED           = 45
	NM_DEVICE_STATE_REASON_GSM_SIM_PIN_REQUIRED           = 46
	NM_DEVICE_STATE_REASON_GSM_SIM_PUK_REQUIRED           = 47
	NM_DEVICE_STATE_REASON_GSM_SIM_WRONG                  = 48
	NM_DEVICE_STATE_REASON_INFINIBAND_MODE                = 49
	NM_DEVICE_STATE_REASON_DEPENDENCY_FAILED              = 50
	NM_DEVICE_STATE_REASON_BR2684_FAILED                  = 51
	NM_DEVICE_STATE_REASON_MODEM_MANAGER_UNAVAILABLE      = 52
	NM_DEVICE_STATE_REASON_SSID_NOT_FOUND                 = 53
	NM_DEVICE_STATE_REASON_SECONDARY_CONNECTION_FAILED    = 54
	NM_DEVICE_STATE_REASON_DCB_FCOE_FAILED                = 55
	NM_DEVICE_STATE_REASON_TEAMD_CONTROL_FAILED           = 56
	NM_DEVICE_STATE_REASON_MODEM_FAILED                   = 57
	NM_DEVICE_STATE_REASON_MODEM_AVAILABLE                = 58
	NM_DEVICE_STATE_REASON_SIM_PIN_INCORRECT              = 59
	NM_DEVICE_STATE_REASON_NEW_ACTIVATION                 = 60
	NM_DEVICE_STATE_REASON_PARENT_CHANGED                 = 61
	NM_DEVICE_STATE_REASON_PARENT_MANAGED_CHANGED         = 62
)

// NMMetered
const (
	NM_METERED_UNKNOWN   = 0
	NM_METERED_YES       = 1
	NM_METERED_NO        = 2
	NM_METERED_GUESS_YES = 3
	NM_METERED_GUESS_NO  = 4
)

// NMActiveConnectionState
const (
	NM_ACTIVE_CONNECTION_STATE_UNKNOWN      = 0
	NM_ACTIVE_CONNECTION_STATE_ACTIVATING   = 1
	NM_ACTIVE_CONNECTION_STATE_ACTIVATED    = 2
	NM_ACTIVE_CONNECTION_STATE_DEACTIVATING = 3
	NM_ACTIVE_CONNECTION_STATE_DEACTIVATED  = 4
)

// NMSecretAgentGetSecretsFlags
const (
	NM_SECRET_AGENT_GET_SECRETS_FLAG_NONE              = 0
	NM_SECRET_AGENT_GET_SECRETS_FLAG_ALLOW_INTERACTION = 1
	NM_SECRET_AGENT_GET_SECRETS_FLAG_REQUEST_NEW       = 2
	NM_SECRET_AGENT_GET_SECRETS_FLAG_USER_REQUESTED    = 4
	NM_SECRET_AGENT_GET_SECRETS_FLAG_ONLY_SYSTEM       = 2147483648
	NM_SECRET_AGENT_GET_SECRETS_FLAG_NO_ERRORS         = 1073741824
)

// NMSecretAgentCapabilities
const (
	NM_SECRET_AGENT_CAPABILITY_NONE      = 0
	NM_SECRET_AGENT_CAPABILITY_VPN_HINTS = 1
)

/* enumerations from "/usr/include/libnm/nm-errors.h" */

// NMAgentManagerError
const (
	NM_AGENT_MANAGER_ERROR_FAILED             = 0
	NM_AGENT_MANAGER_ERROR_PERMISSION_DENIED  = 1
	NM_AGENT_MANAGER_ERROR_INVALID_IDENTIFIER = 2
	NM_AGENT_MANAGER_ERROR_NOT_REGISTERED     = 3
	NM_AGENT_MANAGER_ERROR_NO_SECRETS         = 4
	NM_AGENT_MANAGER_ERROR_USER_CANCELED      = 5
)

// NMConnectionError
const (
	NM_CONNECTION_ERROR_FAILED              = 0
	NM_CONNECTION_ERROR_SETTING_NOT_FOUND   = 1
	NM_CONNECTION_ERROR_PROPERTY_NOT_FOUND  = 2
	NM_CONNECTION_ERROR_PROPERTY_NOT_SECRET = 3
	NM_CONNECTION_ERROR_MISSING_SETTING     = 4
	NM_CONNECTION_ERROR_INVALID_SETTING     = 5
	NM_CONNECTION_ERROR_MISSING_PROPERTY    = 6
	NM_CONNECTION_ERROR_INVALID_PROPERTY    = 7
)

// NMCryptoError
const (
	NM_CRYPTO_ERROR_FAILED            = 0
	NM_CRYPTO_ERROR_INVALID_DATA      = 1
	NM_CRYPTO_ERROR_INVALID_PASSWORD  = 2
	NM_CRYPTO_ERROR_UNKNOWN_CIPHER    = 3
	NM_CRYPTO_ERROR_DECRYPTION_FAILED = 4
	NM_CRYPTO_ERROR_ENCRYPTION_FAILED = 5
)

// NMDeviceError
const (
	NM_DEVICE_ERROR_FAILED                    = 0
	NM_DEVICE_ERROR_CREATION_FAILED           = 1
	NM_DEVICE_ERROR_INVALID_CONNECTION        = 2
	NM_DEVICE_ERROR_INCOMPATIBLE_CONNECTION   = 3
	NM_DEVICE_ERROR_NOT_ACTIVE                = 4
	NM_DEVICE_ERROR_NOT_SOFTWARE              = 5
	NM_DEVICE_ERROR_NOT_ALLOWED               = 6
	NM_DEVICE_ERROR_SPECIFIC_OBJECT_NOT_FOUND = 7
)

// NMManagerError
const (
	NM_MANAGER_ERROR_FAILED                      = 0
	NM_MANAGER_ERROR_PERMISSION_DENIED           = 1
	NM_MANAGER_ERROR_UNKNOWN_CONNECTION          = 2
	NM_MANAGER_ERROR_UNKNOWN_DEVICE              = 3
	NM_MANAGER_ERROR_CONNECTION_NOT_AVAILABLE    = 4
	NM_MANAGER_ERROR_CONNECTION_NOT_ACTIVE       = 5
	NM_MANAGER_ERROR_CONNECTION_ALREADY_ACTIVE   = 6
	NM_MANAGER_ERROR_DEPENDENCY_FAILED           = 7
	NM_MANAGER_ERROR_ALREADY_ASLEEP_OR_AWAKE     = 8
	NM_MANAGER_ERROR_ALREADY_ENABLED_OR_DISABLED = 9
	NM_MANAGER_ERROR_UNKNOWN_LOG_LEVEL           = 10
	NM_MANAGER_ERROR_UNKNOWN_LOG_DOMAIN          = 11
)

// NMSecretAgentError
const (
	NM_SECRET_AGENT_ERROR_FAILED             = 0
	NM_SECRET_AGENT_ERROR_PERMISSION_DENIED  = 1
	NM_SECRET_AGENT_ERROR_INVALID_CONNECTION = 2
	NM_SECRET_AGENT_ERROR_USER_CANCELED      = 3
	NM_SECRET_AGENT_ERROR_AGENT_CANCELED     = 4
	NM_SECRET_AGENT_ERROR_NO_SECRETS         = 5
)

// NMSettingsError
const (
	NM_SETTINGS_ERROR_FAILED               = 0
	NM_SETTINGS_ERROR_PERMISSION_DENIED    = 1
	NM_SETTINGS_ERROR_NOT_SUPPORTED        = 2
	NM_SETTINGS_ERROR_INVALID_CONNECTION   = 3
	NM_SETTINGS_ERROR_READ_ONLY_CONNECTION = 4
	NM_SETTINGS_ERROR_UUID_EXISTS          = 5
	NM_SETTINGS_ERROR_INVALID_HOSTNAME     = 6
)

// NMVpnPluginError
const (
	NM_VPN_PLUGIN_ERROR_FAILED                    = 0
	NM_VPN_PLUGIN_ERROR_STARTING_IN_PROGRESS      = 1
	NM_VPN_PLUGIN_ERROR_ALREADY_STARTED           = 2
	NM_VPN_PLUGIN_ERROR_STOPPING_IN_PROGRESS      = 3
	NM_VPN_PLUGIN_ERROR_ALREADY_STOPPED           = 4
	NM_VPN_PLUGIN_ERROR_WRONG_STATE               = 5
	NM_VPN_PLUGIN_ERROR_BAD_ARGUMENTS             = 6
	NM_VPN_PLUGIN_ERROR_LAUNCH_FAILED             = 7
	NM_VPN_PLUGIN_ERROR_INVALID_CONNECTION        = 8
	NM_VPN_PLUGIN_ERROR_INTERACTIVE_NOT_SUPPORTED = 9
)

/* enumerations from "/usr/include/libnm/nm-setting-8021x.h" */

// NMSetting8021xCKFormat
const (
	NM_SETTING_802_1X_CK_FORMAT_UNKNOWN = 0
	NM_SETTING_802_1X_CK_FORMAT_X509    = 1
	NM_SETTING_802_1X_CK_FORMAT_RAW_KEY = 2
	NM_SETTING_802_1X_CK_FORMAT_PKCS12  = 3
)

// NMSetting8021xCKScheme
const (
	NM_SETTING_802_1X_CK_SCHEME_UNKNOWN = 0
	NM_SETTING_802_1X_CK_SCHEME_BLOB    = 1
	NM_SETTING_802_1X_CK_SCHEME_PATH    = 2
)

/* enumerations from "/usr/include/libnm/nm-setting-connection.h" */

// NMSettingConnectionAutoconnectSlaves
const (
	NM_SETTING_CONNECTION_AUTOCONNECT_SLAVES_DEFAULT = -1
	NM_SETTING_CONNECTION_AUTOCONNECT_SLAVES_NO      = 0
	NM_SETTING_CONNECTION_AUTOCONNECT_SLAVES_YES     = 1
)

/* enumerations from "/usr/include/libnm/nm-setting-dcb.h" */

// NMSettingDcbFlags
const (
	NM_SETTING_DCB_FLAG_NONE      = 0
	NM_SETTING_DCB_FLAG_ENABLE    = 1
	NM_SETTING_DCB_FLAG_ADVERTISE = 2
	NM_SETTING_DCB_FLAG_WILLING   = 4
)

/* enumerations from "/usr/include/libnm/nm-setting.h" */

// NMSettingSecretFlags
const (
	NM_SETTING_SECRET_FLAG_NONE         = 0
	NM_SETTING_SECRET_FLAG_AGENT_OWNED  = 1
	NM_SETTING_SECRET_FLAG_NOT_SAVED    = 2
	NM_SETTING_SECRET_FLAG_NOT_REQUIRED = 4
)

// NMSettingCompareFlags
const (
	NM_SETTING_COMPARE_FLAG_EXACT                      = 0
	NM_SETTING_COMPARE_FLAG_FUZZY                      = 1
	NM_SETTING_COMPARE_FLAG_IGNORE_ID                  = 2
	NM_SETTING_COMPARE_FLAG_IGNORE_SECRETS             = 4
	NM_SETTING_COMPARE_FLAG_IGNORE_AGENT_OWNED_SECRETS = 8
	NM_SETTING_COMPARE_FLAG_IGNORE_NOT_SAVED_SECRETS   = 16
	NM_SETTING_COMPARE_FLAG_DIFF_RESULT_WITH_DEFAULT   = 32
	NM_SETTING_COMPARE_FLAG_DIFF_RESULT_NO_DEFAULT     = 64
	NM_SETTING_COMPARE_FLAG_IGNORE_TIMESTAMP           = 128
)

// NMSettingDiffResult
const (
	NM_SETTING_DIFF_RESULT_UNKNOWN      = 0
	NM_SETTING_DIFF_RESULT_IN_A         = 1
	NM_SETTING_DIFF_RESULT_IN_B         = 2
	NM_SETTING_DIFF_RESULT_IN_A_DEFAULT = 4
	NM_SETTING_DIFF_RESULT_IN_B_DEFAULT = 4
)

/* enumerations from "/usr/include/libnm/nm-setting-ip6-config.h" */

// NMSettingIP6ConfigPrivacy
const (
	NM_SETTING_IP6_CONFIG_PRIVACY_UNKNOWN            = -1
	NM_SETTING_IP6_CONFIG_PRIVACY_DISABLED           = 0
	NM_SETTING_IP6_CONFIG_PRIVACY_PREFER_PUBLIC_ADDR = 1
	NM_SETTING_IP6_CONFIG_PRIVACY_PREFER_TEMP_ADDR   = 2
)

/* enumerations from "/usr/include/libnm/nm-setting-serial.h" */

// NMSettingSerialParity
const (
	NM_SETTING_SERIAL_PARITY_NONE = 0
	NM_SETTING_SERIAL_PARITY_EVEN = 1
	NM_SETTING_SERIAL_PARITY_ODD  = 2
)

/* enumerations from "/usr/include/libnm/nm-setting-vlan.h" */

// NMVlanPriorityMap
const (
	NM_VLAN_INGRESS_MAP = 0
	NM_VLAN_EGRESS_MAP  = 1
)

// NMVlanFlags
const (
	NM_VLAN_FLAG_REORDER_HEADERS = 1
	NM_VLAN_FLAG_GVRP            = 2
	NM_VLAN_FLAG_LOOSE_BINDING   = 4
)

/* enumerations from "/usr/include/libnm/nm-setting-wired.h" */

// NMSettingWiredWakeOnLan
const (
	NM_SETTING_WIRED_WAKE_ON_LAN_PHY       = 2
	NM_SETTING_WIRED_WAKE_ON_LAN_UNICAST   = 4
	NM_SETTING_WIRED_WAKE_ON_LAN_MULTICAST = 8
	NM_SETTING_WIRED_WAKE_ON_LAN_BROADCAST = 16
	NM_SETTING_WIRED_WAKE_ON_LAN_ARP       = 32
	NM_SETTING_WIRED_WAKE_ON_LAN_MAGIC     = 64
	NM_SETTING_WIRED_WAKE_ON_LAN_DEFAULT   = 1
	NM_SETTING_WIRED_WAKE_ON_LAN_IGNORE    = 32768
)

/* enumerations from "/tmp/nm-setting-wireless-security.h" */

// NMWepKeyType
const (
	NM_WEP_KEY_TYPE_UNKNOWN    = 0
	NM_WEP_KEY_TYPE_KEY        = 1
	NM_WEP_KEY_TYPE_PASSPHRASE = 2
)

/* enumerations from "/usr/include/libnm/nm-utils.h" */

// NMUtilsSecurityType
const (
	NMU_SEC_INVALID         = 0
	NMU_SEC_NONE            = 1
	NMU_SEC_STATIC_WEP      = 2
	NMU_SEC_LEAP            = 3
	NMU_SEC_DYNAMIC_WEP     = 4
	NMU_SEC_WPA_PSK         = 5
	NMU_SEC_WPA_ENTERPRISE  = 6
	NMU_SEC_WPA2_PSK        = 7
	NMU_SEC_WPA2_ENTERPRISE = 8
)

/* enumerations from "/usr/include/libnm/nm-vpn-dbus-interface.h" */

// NMVpnServiceState
const (
	NM_VPN_SERVICE_STATE_UNKNOWN  = 0
	NM_VPN_SERVICE_STATE_INIT     = 1
	NM_VPN_SERVICE_STATE_SHUTDOWN = 2
	NM_VPN_SERVICE_STATE_STARTING = 3
	NM_VPN_SERVICE_STATE_STARTED  = 4
	NM_VPN_SERVICE_STATE_STOPPING = 5
	NM_VPN_SERVICE_STATE_STOPPED  = 6
)

// NMVpnConnectionState
const (
	NM_VPN_CONNECTION_STATE_UNKNOWN       = 0
	NM_VPN_CONNECTION_STATE_PREPARE       = 1
	NM_VPN_CONNECTION_STATE_NEED_AUTH     = 2
	NM_VPN_CONNECTION_STATE_CONNECT       = 3
	NM_VPN_CONNECTION_STATE_IP_CONFIG_GET = 4
	NM_VPN_CONNECTION_STATE_ACTIVATED     = 5
	NM_VPN_CONNECTION_STATE_FAILED        = 6
	NM_VPN_CONNECTION_STATE_DISCONNECTED  = 7
)

// NMVpnConnectionStateReason
const (
	NM_VPN_CONNECTION_STATE_REASON_UNKNOWN               = 0
	NM_VPN_CONNECTION_STATE_REASON_NONE                  = 1
	NM_VPN_CONNECTION_STATE_REASON_USER_DISCONNECTED     = 2
	NM_VPN_CONNECTION_STATE_REASON_DEVICE_DISCONNECTED   = 3
	NM_VPN_CONNECTION_STATE_REASON_SERVICE_STOPPED       = 4
	NM_VPN_CONNECTION_STATE_REASON_IP_CONFIG_INVALID     = 5
	NM_VPN_CONNECTION_STATE_REASON_CONNECT_TIMEOUT       = 6
	NM_VPN_CONNECTION_STATE_REASON_SERVICE_START_TIMEOUT = 7
	NM_VPN_CONNECTION_STATE_REASON_SERVICE_START_FAILED  = 8
	NM_VPN_CONNECTION_STATE_REASON_NO_SECRETS            = 9
	NM_VPN_CONNECTION_STATE_REASON_LOGIN_FAILED          = 10
	NM_VPN_CONNECTION_STATE_REASON_CONNECTION_REMOVED    = 11
)

// NMVpnPluginFailure
const (
	NM_VPN_PLUGIN_FAILURE_LOGIN_FAILED   = 0
	NM_VPN_PLUGIN_FAILURE_CONNECT_FAILED = 1
	NM_VPN_PLUGIN_FAILURE_BAD_IP_CONFIG  = 2
)

/* enumerations from "/usr/include/libnm/nm-vpn-editor-plugin.h" */

// NMVpnEditorPluginCapability
const (
	NM_VPN_EDITOR_PLUGIN_CAPABILITY_NONE   = 0
	NM_VPN_EDITOR_PLUGIN_CAPABILITY_IMPORT = 1
	NM_VPN_EDITOR_PLUGIN_CAPABILITY_EXPORT = 2
	NM_VPN_EDITOR_PLUGIN_CAPABILITY_IPV6   = 4
)

/* enumerations from "/usr/include/libnm/nm-wimax-nsp.h" */

// NMWimaxNspNetworkType
const (
	NM_WIMAX_NSP_NETWORK_TYPE_UNKNOWN         = 0
	NM_WIMAX_NSP_NETWORK_TYPE_HOME            = 1
	NM_WIMAX_NSP_NETWORK_TYPE_PARTNER         = 2
	NM_WIMAX_NSP_NETWORK_TYPE_ROAMING_PARTNER = 3
)

/* Generated data ends here */
