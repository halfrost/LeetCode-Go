---
title: "Flowchart support"
date: 2019-11-14T12:00:06+09:00
description: "flowchart.js is a flowchart DSL and SVG render that runs in the browser and terminal.
Nodes and connections are defined in separately so that nodes can be reused and connections can be quickly changed."
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
- flowchartjs
image: images/feature1/flowchart.png
---

```flowchart
st=>start: Start|past:>http://www.google.com[blank]
e=>end: End|future:>http://www.google.com
op1=>operation: My Operation|past
op2=>operation: Stuff|current
sub1=>subroutine: My Subroutine|invalid
cond=>condition: Yes
or No?|approved:>http://www.google.com
c2=>condition: Good idea|rejected
io=>inputoutput: catch something...|future

st->op1(right)->cond
cond(yes, right)->c2
cond(no)->sub1(left)->op1
c2(yes)->io->e
c2(no)->op2->e
```
