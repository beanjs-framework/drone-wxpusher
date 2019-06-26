# Drone ServerChan(Server酱)

[![Build Status](https://cloud.drone.io/api/badges/beanjs-framework/drone-wxpusher/status.svg)](https://cloud.drone.io/beanjs-framework/drone-wxpusher)

drone-wxpusher微信消息通知插件

## 简介

基于 [wxpusher](http://wxpusher.dingliqc.com/) 封装的微信消息通知插件

## 例子

明文 key 配置

```yml
---
kind: pipeline
name: default

steps:
  - name: send-wechat
    image: beanjs/drone-wxpusher
    settings:
      msg: {消息内容}
      ids: {用户ID}
      title: {消息标题}
      url: {跳转链接}
```

密文 key 配置

```yml
---
kind: pipeline
name: default

steps:
  - name: send-wechat
    image: beanjs/drone-wxpusher
    settings:
      msg: {消息内容}
      ids:
        from_secret: {your ids}
      title: {消息标题}
      url: {跳转链接}

```
