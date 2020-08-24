---
title: "Shortcodes"
date: 2020-01-25T06:40:51+09:00
description: tabs, code-tabs, expand, alert, warning, notice, img, box
draft: false
hideToc: false
enableToc: true
enableTocContent: true
tocPosition: inner
tags:
- shortcode
series:
-
categories:
-
image: images/feature3/code-file.png
---

## Markdownify box

{{< boxmd >}}
This is **boxmd** shortcode
{{< /boxmd >}}

## Simple box

{{< box >}}
This is **box** shortcode
{{< /box >}}

## Code tabs

Make it easy to switch between different code

{{< codes java javascript >}}
  {{< code >}}

  ```java
  System.out.println('Hello World!');
  ```

  {{< /code >}}

  {{< code >}}

  ```javascript
  console.log('Hello World!');
  ```
  
  {{< /code >}}
{{< /codes >}}

## Tabs for general purpose

{{< tabs Windows MacOS Ubuntu >}}
  {{< tab >}}

  ### Windows section

  ```javascript
  console.log('Hello World!');
  ```

  ⚠️Becareful that the content in the tab should be different from each other. The tab makes unique id hashes depending on the tab contents. So, If you just copy-paste the tabs with multiple times, since it has the same contents, the tab will not work.

  {{< /tab >}}
  {{< tab >}}

  ### MacOS section

  Hello world!
  {{< /tab >}}
  {{< tab >}}

  ### Ubuntu section

  Great!
  {{< /tab >}}
{{< /tabs >}}

## Expand

{{< expand "Expand me" >}}

### Title

contents

{{< /expand >}}

{{< expand "Expand me2" >}}

### Title2

contents2

{{< /expand >}}

## Alert

Colored box

{{< alert theme="warning" >}}
**this** is a text
{{< /alert >}}

{{< alert theme="info" >}}
**this** is a text
{{< /alert >}}

{{< alert theme="success" >}}
**this** is a text
{{< /alert >}}

{{< alert theme="danger" >}}
**this** is a text
{{< /alert >}}

## Notice

{{< notice success >}}
success text
{{< /notice >}}

{{< notice info >}}
info text
{{< /notice >}}

{{< notice warning >}}
warning text
{{< /notice >}}

{{< notice error >}}
error text
{{< /notice >}}
