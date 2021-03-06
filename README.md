### 字段 Fields:

| Tag | 中文描述 | Description |
| - | - | - |
| eqcsfield | 字段等于另一个字段 | Field Equals Another Field (relative)|
| eqfield | 字段等于另一个字段 |  Field Equals Another Field |
| fieldcontains | doc.go 中未记录 | NOT DOCUMENTED IN doc.go |
| fieldexcludes | doc.go 中未记录 | NOT DOCUMENTED IN doc.go |
| gtcsfield | 字段大于另一个字段|  Field Greater Than Another Relative Field |
| gtecsfield | 字段大于等于另一个字段 | Field Greater Than or Equal To Another Relative Field |
| gtefield | 字段大于或等于另一个字段 |  Field Greater Than or Equal To Another Field |
| gtfield | 场大于另一个场 |  Field Greater Than Another Field |
| ltcsfield | 小于另一个相对字段 |  Less Than Another Relative Field |
| ltecsfield |小于或等于另一个相对字段|  Less Than or Equal To Another Relative Field |
| ltefield | 小于或等于另一个字段 | Less Than or Equal To Another Field |
| ltfield | 少于另一个领域 | Less Than Another Field |
| necsfield | 字段不等于另一个字段（相对）|  Field Does Not Equal Another Field (relative) |
| nefield | 字段不等于另一个字段 | Field Does Not Equal Another Field |

### 网络 Network:

| Tag | 中文描述 | Description |
| - | - | - |
| cidr |  | Classless Inter-Domain Routing CIDR |
| cidrv4 |   | Classless Inter-Domain Routing CIDRv4 |
| cidrv6 |   | Classless Inter-Domain Routing CIDRv6 |
| datauri |   | Data URL |
| fqdn |   | Full Qualified Domain Name (FQDN) |
| hostname |   | Hostname RFC 952 |
| hostname_port |   | HostPort |
| hostname_rfc1123 |   | Hostname RFC 1123 |
| ip |   | Internet Protocol Address IP |
| ip4_addr |   | Internet Protocol Address IPv4 |
| ip6_addr |  |Internet Protocol Address IPv6 |
| ip_addr |   | Internet Protocol Address IP |
| ipv4  |  | Internet Protocol Address IPv4 |
| ipv6 |   | Internet Protocol Address IPv6 |
| mac |   | Media Access Control Address MAC |
| tcp4_addr |   | Transmission Control Protocol Address TCPv4 |
| tcp6_addr |   | Transmission Control Protocol Address TCPv6 |
| tcp_addr |   | Transmission Control Protocol Address TCP |
| udp4_addr |   | User Datagram Protocol Address UDPv4 |
| udp6_addr |   | User Datagram Protocol Address UDPv6 |
| udp_addr |  用户数据报协议地址 UDP | User Datagram Protocol Address UDP |
| unix_addr |  Unix域套接字端点地址 | Unix domain socket end point Address |
| uri |  URI 字符串 | URI String |
| url |  网址字符串 | URL String |
| url_encoded |  网址编码 | URL Encoded |
| urn_rfc2141 | 瓮 RFC 2141 字符串  | Urn RFC 2141 String |

### Strings:

| Tag | Description |
| - | - |
| alpha | Alpha Only |
| alphanum | Alphanumeric |
| alphanumunicode | Alphanumeric Unicode |
| alphaunicode | Alpha Unicode |
| ascii | ASCII |
| contains | Contains |
| containsany | Contains Any |
| containsrune | Contains Rune |
| endswith | Ends With |
| lowercase | Lowercase |
| multibyte | Multi-Byte Characters |
| number | NOT DOCUMENTED IN doc.go |
| numeric | Numeric |
| printascii | Printable ASCII |
| startswith | Starts With |
| uppercase | Uppercase |

### Format:
| Tag | Description |
| - | - |
| base64 | Base64 String |
| base64url | Base64URL String |
| btc_addr | Bitcoin Address |
| btc_addr_bech32 | Bitcoin Bech32 Address (segwit) |
| datetime | Datetime |
| e164 | e164 formatted phone number |
| email | E-mail String
| eth_addr | Ethereum Address |
| hexadecimal | Hexadecimal String |
| hexcolor | Hexcolor String |
| hsl | HSL String |
| hsla | HSLA String |
| html | HTML Tags |
| html_encoded | HTML Encoded |
| isbn | International Standard Book Number |
| isbn10 | International Standard Book Number 10 |
| isbn13 | International Standard Book Number 13 |
| json | JSON |
| latitude | Latitude |
| longitude | Longitude |
| rgb | RGB String |
| rgba | RGBA String |
| ssn | Social Security Number SSN |
| uuid | Universally Unique Identifier UUID |
| uuid3 | Universally Unique Identifier UUID v3 |
| uuid3_rfc4122 | Universally Unique Identifier UUID v3 RFC4122 |
| uuid4 | Universally Unique Identifier UUID v4 |
| uuid4_rfc4122 | Universally Unique Identifier UUID v4 RFC4122 |
| uuid5 | Universally Unique Identifier UUID v5 |
| uuid5_rfc4122 | Universally Unique Identifier UUID v5 RFC4122 |
| uuid_rfc4122 | Universally Unique Identifier UUID RFC4122 |

### Comparisons:
| Tag | Description |
| - | - |
| eq | Equals |
| gt | Greater than|
| gte |Greater than or equal |
| lt | Less Than |
| lte | Less Than or Equal |
| ne | Not Equal |

### Other:
| Tag | Description |
| - | - |
| dir | Directory |
| endswith | Ends With |
| excludes | Excludes |
| excludesall | Excludes All |
| excludesrune | Excludes Rune |
| file | File path |
| isdefault | Is Default |
| len | Length |
| max | Maximum |
| min | Minimum |
| oneof | One Of |
| required | Required |
| required_if | Required If |
| required_unless | Required Unless |
| required_with | Required With |
| required_with_all | Required With All |
| required_without | Required Without |
| required_without_all | Required Without All |
| excluded_with | Excluded With |
| excluded_with_all | Excluded With All |
| excluded_without | Excluded Without |
| excluded_without_all | Excluded Without All |
| unique | Unique |