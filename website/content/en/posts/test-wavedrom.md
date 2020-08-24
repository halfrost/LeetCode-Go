---
title: "Wavedrom support"
date: 2019-11-18T18:00:06+09:00
description: "WaveDrom is a Free and Open Source online digital timing diagram (waveform) rendering engine that uses javascript, HTML5 and SVG to convert a WaveJSON input text description into SVG vector graphics."
draft: false
enableToc: false
enableTocContent: false
tags:
-
series:
-
categories:
- diagram
libraries:
- wavedrom
image: images/feature1/wave.png
---

```wave
{ 
  "signal": [ {"name": "CLK", "wave": "p.....|..."},
            {"name":"DAT", "wave":"x.345x|=.x", "data":["A","B","C","D"]},
            {"name": "REQ", "wave": "0.1..0|1.0"},
            {},
            {"name": "ACK", "wave": "1.....|01."}
]}
```
