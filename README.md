# 分布式缓存项目

## 项目简介

这是一个功能强大的分布式缓存项目，旨在解决资源限制和性能瓶颈问题。通过多种优化技术和算法的实现，本项目能够高效地缓存数据并支持高并发访问。
## 学习来源

本项目的学习来源于 [GeekTutu](https://github.com/geektutu)，通过对该[项目](https://github.com/geektutu/7days-golang/tree/master/gee-cache)的分析与研究，实现了以下功能和特性。

## 特性

1. **LRU 缓存淘汰算法**  
   为了解决资源限制的问题，本项目实现了 LRU（Least Recently Used）缓存淘汰算法。该算法通过跟踪最近使用的缓存数据，确保在缓存空间不足时优先淘汰最久未使用的数据，提高了数据的访问效率。

2. **单机并发支持**  
   项目实现了单机并发访问，能够处理多个请求并发执行。同时，提供了用户自定义数据源的回调函数，使用户可以灵活地指定数据的获取方式，以适应不同的使用场景。

3. **HTTP 服务端**  
   本项目包含一个高性能的 HTTP 服务端，允许用户通过 RESTful API 进行缓存操作。用户可以方便地进行数据存储、查询和删除等操作，极大地简化了缓存管理的复杂性。

4. **一致性哈希算法**  
   为了解决分布式环境中远程节点的挑选问题，本项目实现了一致性哈希算法。这一算法能够在节点增减时有效地减少数据迁移，确保系统的高可用性和稳定性。

5. **HTTP 客户端**  
   项目中还实现了一个 HTTP 客户端，支持多节点间的通信。该客户端能够高效地与分布式缓存集群中的各个节点进行交互，确保数据的快速访问和一致性。

6. **Singleflight 机制**  
   为了解决缓存击穿的问题，本项目实现了 singleflight 机制。该机制能够在高并发情况下，避免重复的缓存请求，从而减少对后端数据源的压力，提高系统的整体性能。

7. **使用 Protobuf 优化通信性能**  
   项目使用 Protocol Buffers（protobuf）库来序列化和反序列化数据，优化了节点间通信的性能。protobuf 的高效数据编码和解码能力，显著提升了数据传输的速度和占用的带宽。

## 安装与使用

### 前提条件

- Go 1.16 或更高版本
- Protobuf 编译器（`protoc`）

### 安装步骤

1. **克隆项目**

   ```bash
   git clone https://github.com/KinoHui/distributed-cache-demo.git
   cd distributed-cache-demo
   ```

2. **运行服务**

   ```bash
   ./start_servers.sh
   ```

3. **访问 HTTP API**

   你可以使用 `curl` 或其他 HTTP 客户端工具来访问 API，例如：

   ```bash
   curl "http://localhost:9999/api?key=Tom" 
   ```

## 致谢

再次特别感谢 [GeekTutu](https://github.com/geektutu) 对开源社区的贡献。通过他的项目，我不仅获得了灵感，也提升了自己的技能。这个项目是我学习的一个重要里程碑，感谢他为我提供了这样一个学习的平台。


## 联系

如有任何问题或建议，请联系：

- [KinoHui](mailto:568000273@qq.com)
- GitHub: [KinoHui](https://github.com/KinoHui)


