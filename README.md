# ZPanel
**An Advanced And Powerfull Web Panel • Built on Sing-Box**



> **Disclaimer:** This project is only for personal learning and communication, please do not use it for illegal purposes, please do not use it in a production environment  
 

## Quick Overview
| Features                               |      Enable?       |
| -------------------------------------- | :----------------: |
| Multi-Protocol                         | :heavy_check_mark: |
| Multi-Language                         | :heavy_check_mark: |
| Multi-Client/Inbound                   | :heavy_check_mark: |
| Advanced Traffic Routing Interface     | :heavy_check_mark: |
| Client & Traffic & System Status       | :heavy_check_mark: |
| Subscription Service (link/json + info)| :heavy_check_mark: |
| Dark/Light Theme                       | :heavy_check_mark: |
| API Interface                          | :heavy_check_mark: |

## Supported Platforms
| Platform | Architecture | Status |
|----------|--------------|---------|
| Linux    | amd64, arm64 | ✅ Supported |
## Screenshots

!["Main"](https://github.com/zamibd/ZPanel/blob/main/frontend/media/main.png)
 
 
## Default Installation Information
- Panel Port: 2080
- Panel Path: /admin/
- Subscription Port: 2081
- Subscription Path: /customer/
- User :    admin
- Password: admin

## Install & Upgrade to Latest Version

### Linux/macOS
```sh
bash <(curl -Ls https://raw.githubusercontent.com/zamibd/ZPanel/main/install.sh)
```

Follow the installation wizard

## Uninstall ZPanel

```sh
sudo -i

systemctl disable ZPanel  --now

rm -f /etc/systemd/system/sing-box.service
systemctl daemon-reload

rm -fr /usr/local/ZPanel
rm /usr/bin/ZPanel
```

 

## Languages

- English
- Farsi
- Vietnamese
- Chinese (Simplified)
- Chinese (Traditional)
- Russian
 
- Supported protocols:
  - General:  Mixed, SOCKS, HTTP, HTTPS, Direct, Redirect, TProxy
  - V2Ray based: VLESS, VMess, Trojan, Shadowsocks
  - Other protocols: ShadowTLS, Hysteria, Hysteria2, Naive, TUIC
- Supports XTLS protocols
- An advanced interface for routing traffic, incorporating PROXY Protocol, External, and Transparent Proxy, SSL Certificate, and Port
- An advanced interface for inbound and outbound configuration
- Clients’ traffic cap and expiration date
- Displays online clients, inbounds and outbounds with traffic statistics, and system status monitoring
- Subscription service with ability to add external links and subscription
- HTTPS for secure access to the web panel and subscription service (self-provided domain + SSL certificate)
- Dark/Light theme

## Environment Variables

<details>
  <summary>Click for details</summary>

### Usage

| Variable       |                      Type                      | Default       |
| -------------- | :--------------------------------------------: | :------------ |
| ZPANEL_LOG_LEVEL  | `"debug"` \| `"info"` \| `"warn"` \| `"error"` | `"info"`      |
| ZPANEL_DEBUG      |                   `boolean`                    | `false`       |
| ZPANEL_BIN_FOLDER |                    `string`                    | `"bin"`       |
| ZPANEL_DB_FOLDER  |                    `string`                    | `"db"`        |
| SINGBOX_API    |                    `string`                    | -             |

</details>