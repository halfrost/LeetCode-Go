---
title: "Viz support"
date: 2019-11-18T21:00:06+09:00
description: "A hack to put Graphviz on the web."
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
- viz
image: images/feature2/graph.png
---

```viz-dot
  digraph G {

	subgraph cluster_0 {
		style=filled;
		color=lightgrey;
		node [style=filled,color=white];
		a0 -> a1 -> a2 -> a3;
		label = "process #1";
	}

	subgraph cluster_1 {
		node [style=filled];
		b0 -> b1 -> b2 -> b3;
		label = "process #2";
		color=blue
	}
	start -> a0;
	start -> b0;
	a1 -> b3;
	b2 -> a3;
	a3 -> a0;
	a3 -> end;
	b3 -> end;

	start [shape=Mdiamond];
	end [shape=Msquare];
}

```
