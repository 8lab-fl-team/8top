
# 8top

从rtop魔改而来的工具，用于记录服务器资源信息。

- 只支持Key验证，并且Key不能有密码，**仅在非常可信的环境运行该程序**
- 输出结果自行分析

```shell
Usage:
  8top [flags]

Flags:
  -c, --config string   config file (default "config.yml")
  -d, --dist string     distribution name
  -h, --help            help for 8top
  -i, --interval int    refresh interval in seconds (default 5)
  -k, --key string      path to private key
  -t, --toggle          Help message for toggle
```  

## 输出结果

形如以下格式的数据

注意：文件虽然叫.json，但一个文件中实际是N个json的组合，每行一个，追加写入
  
```json
{"uptime":2882355840000000,"hostname":"slave2","load_1":"0.26","load_5":"0.28","load_10":"0.50","running_procs":"1","total_procs":"891","mem_total":67416649728,"mem_free":9621766144,"mem_buffers":19817644032,"mem_cached":34371653632,"swap_total":6361702400,"swap_free":5778862080,"fs_infos":[{"MountPoint":"/","Used":40873766912,"Free":156598140928},{"MountPoint":"/boot","Used":168603648,"Free":784236544},{"MountPoint":"/data","Used":29629337600,"Free":1044111433728}],"net_intf":{"br-12d27dd74afe":{"IPv4":"192.168.49.1/24","IPv6":"fe80::42:d3ff:fe2d:618/64","Rx":31430986,"Tx":1159976},"br-51a5a3ebcd93":{"IPv4":"172.18.0.1/16","IPv6":"fe80::42:3cff:fe3b:805f/64","Rx":15053566891,"Tx":12854130219},"docker0":{"IPv4":"172.17.0.1/16","IPv6":"fe80::42:5dff:fe2a:3ff6/64","Rx":5962502,"Tx":272164},"eth0":{"IPv4":"192.168.1.78/24","IPv6":"fe80::250:56ff:feab:ea2b/64","Rx":57622454233,"Tx":123027864},"eth1":{"IPv4":"103.244.233.123/28","IPv6":"fe80::250:56ff:feab:16fc/64","Rx":11801983649,"Tx":10934999556},"ifb0":{"IPv4":"","IPv6":"fe80::ec57:fdff:fe03:7ff8/64","Rx":12710557,"Tx":12709717},"lo":{"IPv4":"127.0.0.1/8","IPv6":"::1/128","Rx":384060608,"Tx":384060608},"veth38a1a3c":{"IPv4":"","IPv6":"fe80::1cff:cdff:fe6c:c1e9/64","Rx":1175371509,"Tx":3367060418},"veth4bf835d":{"IPv4":"","IPv6":"fe80::6836:24ff:fea4:f514/64","Rx":553338030,"Tx":929070486},"veth7b7c3ba":{"IPv4":"","IPv6":"fe80::4cd9:44ff:fe89:38f3/64","Rx":102786154397,"Tx":76439267697},"veth8a28c39":{"IPv4":"","IPv6":"fe80::d06e:78ff:fe2e:ad9d/64","Rx":9652694921,"Tx":450129129661},"veth8ff6604":{"IPv4":"","IPv6":"fe80::b426:72ff:fe0e:8016/64","Rx":998485100,"Tx":304498552},"vetha2105f0":{"IPv4":"","IPv6":"fe80::5868:7ff:fe01:84ce/64","Rx":592152824104,"Tx":11598860626},"vethb31b2d3":{"IPv4":"","IPv6":"fe80::8ca3:1cff:fedd:fa4/64","Rx":38889738,"Tx":1178132},"vethfeefb3f":{"IPv4":"","IPv6":"fe80::e8ed:2cff:fe5c:93ff/64","Rx":59620622632,"Tx":132072886468}},"cpu":{"User":0,"Nice":0,"System":0,"Idle":0,"Iowait":0,"Irq":0,"SoftIrq":0,"Steal":0,"Guest":0},"time":1658146298,"host_ip":"103.244.233.123"}
```
