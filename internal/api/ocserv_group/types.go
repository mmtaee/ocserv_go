package ocserv_group

type GroupConfig struct {
	RxDataPerSec         int64    `json:"rx_data_per_sec"`
	TxDataPerSec         int64    `json:"tx_data_per_sec"`
	MaxSameClients       int64    `json:"max_same_clients"`
	Ipv4Network          string   `json:"ipv4_network"`
	Dns                  []string `json:"dns"`
	NoUdp                bool     `json:"no_udp"`
	Keepalive            int64    `json:"keepalive"`
	Dpd                  int64    `json:"dpd"`
	MobileDpd            int64    `json:"mobile_dpd"`
	TunnelAllDns         bool     `json:"tunnel_all_dns"`
	RestrictUserToRoutes bool     `json:"restrict_user_to_routes"`
	Mtu                  int64    `json:"mtu"`
	IdleTimeout          int64    `json:"idle_timeout"`
	MobileIdleTimeout    int64    `json:"mobile_idle_timeout"`
	SessionTimeout       int64    `json:"session_timeout"`
	Routes               []string `json:"routes"`
	NoRoutes             []string `json:"no_routes"`
}

type CreateOcservGroupData struct {
	GroupName string      `json:"group_name" binding:"required"`
	Config    GroupConfig `json:"config" binding:"omitempty"`
}

type UpdateOcservGroupData struct {
	Config GroupConfig `json:"config" binding:"omitempty"`
}
