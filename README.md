# TableKV
TableKV 基于leveldb，实现了分区的KV库，网络协议采用聚合方式，目前支持thrift，后续计划支持gRPC等，天然支持多语言开发。

来源：有个项目每天都有大量的数据录入和删除，同时要求较高的录入和删除性能。按天分区恰好可以实现所有需求，所以有了TableKV。

### 已支持功能：
* 支持分区的kv库，以分目录方式实现。
* 支持分区的关闭和快速删除。
* 支持分区的自动创建。
* 支持多协议接入，已支持thrift。
* 支持分区按小时和天进行过期。

### 未来支持功能:
* 支持分布式
* 支持同步和备份

### 基本要求
* 分区名仅支持英文字母和数字，不能包括其它特殊的字符。

# TableKV
TableKV implements partitioned KV library based on level db. the network protocol adopts aggregation method. At present, TableKV supports thrift, gRPC and so on, and naturally supports multilingual development.

Source: there is a project that has a lot of data input and deletion every day, and requires high input and delete performance.

### is supported:
* support partitioned KV library, implemented in a directory way.
* supports partition closure and quick deletion.
* supports automatic creation of partitions.
* support multi-protocol access, has been supported thrift.
* Supports partition expiration by hour and day.

### Future support:
* support for distributed
* support for synchronization and backup

### Requirements
* Partition names only support English letters and numbers , and cannot include other special characters .