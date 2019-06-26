# Drone wxpusher

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
      ids: uid1,uid2,uid3 (必须)
      title: {消息标题} (可选:)
      subtitle_text: {工单类型-文字} (可选:简单提醒)
      subtitle_color: {工单类型-颜色} (可选:#C0C0C0)
      remark: {消息内容} (必须)
      url: {跳转链接} (可选:)
```
