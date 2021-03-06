goim
==============
goim is a im server writen by golang.

## Features
 * Light weight
 * High performance
 * Pure Golang
 * Supports single push, multiple push and broadcasting
 * Supports one key to multiple subscribers (Configurable maximum subscribers count)
 * Supports heartbeats (Application heartbeats, TCP, KeepAlive, HTTP long pulling)
 * Supports authentication (Unauthenticated user can't subscribe)
 * Supports multiple protocols (WebSocket，TCP，HTTP）
 * Scalable architecture (Unlimited dynamic job and logic modules)
 * Asynchronous push notification based on Kafka

## Architecture
![arch](https://github.com/Terry-Mao/goim/blob/master/doc/arch.png)

Protocol:

[proto](https://github.com/Terry-Mao/goim/blob/master/doc/protocol.png)

## Document
[English](./README_en.md)

[中文](./README_cn.md)

## Examples
Websocket: [Websocket Client Demo](https://github.com/Terry-Mao/goim/tree/master/examples/javascript)

Android: [Android](https://github.com/roamdy/goim-sdk)

iOS: [iOS](https://github.com/roamdy/goim-oc-sdk)

## Benchmark

![benchmark](https://github.com/Terry-Mao/goim/blob/master/doc/benchmark.png)

### Benchmark Server

| CPU | Memory | Instance |
| :---- | :---- | :---- |
| Intel(R) Xeon(R) CPU E5-2630 v2 @ 2.60GHz  | DDR3 32GB | 5 |

### Benchmark Case

* Online: 500,000
* Duration: 10min
* Push Speed: 20/s (broadcast room)
* Push Message: {"test":1}

### Benchmark Resource

* CPU: 2340% (almost all busy)
* Memory: 2.65GB
* GC Pause: 41ms
* Network: Incoming(500MBit/s), Outgoing(780MBit/s)

### Benchmark Result

5.6 million/second message received with 5 24c server, 1.2 million/second per server.

[中文](./doc/benchmark_cn.md)

[English](./doc/benchmark_en.md)

## LICENSE
goim is is distributed under the terms of the GNU General Public License, version 3.0 [GPLv3](http://www.gnu.org/licenses/gpl.txt)
