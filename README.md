# Drone wxpusher

[![Build Status](https://cloud.drone.io/api/badges/beanjs-framework/drone-wxpusher/status.svg)](https://cloud.drone.io/beanjs-framework/drone-wxpusher)

drone-wxpusher微信消息通知插件

## 简介

基于 [wxpusher](http://wxpusher.zjiecode.com/docs/#/) 封装的微信消息通知插件

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
      uids: uid1,uid2,uid3 
      content: {消息内容}
      content_type: {消息类型：1，2，3}
      app_token: {应用ID}
      url: {跳转链接} (可选:)
```
