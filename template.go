package vmessconfig

var DefaultBalancerTemplate = `
{
  "log": {
    "loglevel": "warning"
  },
  "inbounds": [
    {
      "port": 1080,
      "listen": "0.0.0.0",
      "protocol": "socks",
      "sniffing": {
        "enabled": true,
        "destOverride": [
          "http",
          "tls"
        ]
      },
      "settings": {
        "auth": "noauth",
        "udp": true
      }
    },
    {
      "port": 80,
      "listen": "0.0.0.0",
      "protocol": "http",
      "settings": {
        "udp": true
      }
    }
  ],
  "outbounds": [
    {
      "tag": "vmessconfig-outbound-insert",
      "protocol": "freedom",
      "settings": {}
    },
    {
      "tag": "direct",
      "protocol": "freedom",
      "settings": {}
    }
  ],
  "routing": {
    "domainStrategy": "IPIfNonMatch",
    "rules": [
      {
        "type": "field",
        "balancerTag": "vmessconfig-autogenerated-balancer",
        "domainStrategy": "IPOnDemand",
        "ip": [
          "0.0.0.0/0"
        ]
      },
      {
        "type": "field",
        "outboundTag": "direct",
        "domainStrategy": "IPOnDemand",
        "ip": [
          "geoip:private",
          "geoip:cn"
        ]
      },
      {
        "type": "field",
        "outboundTag": "direct",
        "domain": [
          "geosite:private",
          "geosite:cn"
        ]
      }
    ],
    "balancers": [
      {
        "tag": "vmessconfig-autogenerated-balancer",
        "selector": [],
        "strategy": {
          "type": "random"
        }
      }
    ]
  }
}
`

var DefaultSingleNodeTemplate = `
{
  "log": {
    "loglevel": "warning"
  },
  "inbounds": [
    {
      "tag": "proxy",
      "port": 1080,
      "listen": "0.0.0.0",
      "protocol": "socks",
      "sniffing": {
        "enabled": true,
        "destOverride": [
          "http",
          "tls"
        ]
      },
      "settings": {
        "auth": "noauth",
        "udp": true
      }
    },
    {
      "port": 80,
      "listen": "0.0.0.0",
      "protocol": "http",
      "settings": {
        "udp": true
      }
    }
  ],
  "outbounds": [
    {
      "tag": "vmessconfig-outbound-insert",
      "protocol": "freedom",
      "settings": {}
    }
  ]
}
`