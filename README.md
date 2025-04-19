# 🔍 Multi-DNS Resolver

This is a high-performance, concurrent DNS query tool written in Go. It supports customizable DNS record types (A, AAAA, MX, TXT, etc.) and performs parallel queries to multiple DNS servers, returning the **first successful response**.

## 🚀 Features

- ✅ Supports multiple DNS servers
- ✅ Supports various DNS record types (A, AAAA, MX, TXT, CNAME)
- ✅ Concurrent queries to all servers
- ✅ Uses the **fastest valid response**
- ✅ Built with [miekg/dns](https://github.com/miekg/dns) library

---