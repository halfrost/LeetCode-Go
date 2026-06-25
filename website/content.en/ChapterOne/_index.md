---
title: Chapter 1 Prologue
type: docs
weight: 1
---

# Chapter 1 Prologue

{{< columns >}}
## About LeetCode

Speaking of LeetCode, as a programmer, you should be no stranger to it; in recent years it is always mentioned in interviews. Programmers at home and abroad mainly use it to practice problems for interviews. According to historical records, this website was founded in 2011 and is about to celebrate its 10th anniversary. Weekly contests, biweekly contests, and monthly contests are held, and coding within a limited time can indeed test one's algorithmic ability very well. In contests sponsored and named by some large companies, those who rank among the top not only receive prizes, but can also directly get opportunities for internal referrals.

<--->

## What Is a Cookbook

Translated literally, it is a cooking book, a book that teaches you how to make all kinds of recipes and delicacies. Students who often read O'Reilly technical books will be very familiar with this term. Generally, hands-on, practical books have this name.

{{< /columns >}}

<img src="https://books.halfrost.com/leetcode/logo.png" alt="logo" height="600" align="right" style="padding-left: 30px;"/>

## Why I Wrote This Open-Source Book

The author has been practicing problems for a year and wants to share some thoughts on doing problems and solution methods with everyone. I want to make friends with people who have the same interests, and communicate and learn together. For myself, writing solutions is also a form of improvement. Explaining a profound problem to someone who has no clue at all, and making them fully understand it, can really train one's ability to express ideas. During the explanation, you may also encounter some questions from listeners; these questions may be gaps in your own knowledge, forcing you to fill them. The author has given related sharing sessions at the company and felt this deeply; both sides benefited quite a lot.

> In addition, during college, when the author was doing problems, what I hated most was writing solutions. I felt it was a waste of time and that I should spend more time doing more problems. Now I don't know whether this counts as "what goes around comes around".


## About the Book Cover

Students who often read O'Reilly animal books will know at a glance that this cover is a tribute to them. That is indeed the purpose. The animals on O'Reilly covers are all rare animals, and the art style is black-and-white sketching. These animals are all copyrighted, so I could only look online for copyright-free black-and-white sketch-style images. Commonly, about 40 images in this style can be found. However, too many people use them, so the author went to great lengths to find several other such images, and this peacock displaying its tail is one of them. The meaning of the peacock displaying its tail is to hope that after everyone finishes practicing LeetCode, they improve their own algorithmic ability and unfold their own "tail" on the stage of life. The color scheme of the whole book is also green, because this is the color of AC.


## About the Author

The author is a gopher newcomer who has just been in the industry for a year and a half; I hope all the experts will give me more guidance. I participated in ACM-ICPC for 3 years in college, but because my aptitude was not high, I did not win a gold medal. So in terms of algorithms, my evaluation of myself is that I am a beginner. The greatest gain from participating in ACM-ICPC was training my thinking ability, which is also applied in life. Secondly, I got to know many very smart contestants in China and saw the gap between myself and them. Finally, there are those more than 200 pages of densely printed [algorithm templates](https://github.com/halfrost/leetcode-go/releases/tag/Special), some of which I still have not fully understood myself. Knowledge you have learned belongs to you for life; if you have not learned it, that knowledge is merely something external.

The author started practicing problems on March 25, 2019, and by March 25, 2020, it had been exactly one year. The original plan was one problem per day. In reality, sometimes there was more than one problem per day, and in the end I completed 600+:

![](https://img.halfrost.com/Blog/ArticleImage/2019_leetcode.png)

> A warm reminder: the author originally thought that doing one problem every day would make this submissions graph completely green, but I found I was wrong. If you also want to persist and make this graph completely green, you must pay attention to the following issue: LeetCode servers are in the +0 time zone, and this graph is also calculated according to this time zone. In other words, before 8 a.m. every day in China, it counts as the previous day! It is also because of the time zone issue that these 22 squares were left blank for me. For example, if there is a very difficult Hard problem, and there is also a lot of work that day, by the time I think of the solution after getting home from work at night, it is already the early morning of the next day. Then I do another problem as the next day's quota. As a result, you will find that both of these 2 problems are counted as the previous day. Sometimes the author gets up at 6 a.m. to practice problems, and after submitting, they are also counted as the previous day.
> 
> (Of course, all of this is in the past and no longer important; just take it as some little episodes on the road of hard work)

In 2020, the author will definitely continue practicing problems, because I have not yet achieved some of my own goals. I may forge ahead toward 1000 problems, or I may start a second and third round after reaching 800 problems. (I guess I won't stop until I reach my goal~)

## About the Code in the Book

The code is all placed in the [github repo](https://github.com/halfrost/leetcode-go/tree/master/leetcode), and problems can be found by searching by problem number.
The code for the problems in this book has already reached beats 100%. Solutions that have not reached beats 100% have not been included in this book. The author will continue optimizing those problems to 100% before putting them in.

Some readers may ask why pursue beats 100%. The author believes that only after optimizing to beats 100% can it be considered that you have truly gotten a feel for the problem. For several Hard problems, the author used brute-force solutions to AC them, but then only beats 5%. Such a problem is almost as if it had not been done. Moreover, in an interview, if you give such an answer, the interviewer will not be satisfied either: "Is there a better solution?" If you can give a better solution through your own thinking, the interviewer will be more satisfied.

LeetCode's statistics for code runtime fluctuate; submitting the same code 10 times may result in beats 100%. The author did not notice this problem at first, and for many problems submitted the correct code many times in a row, submitting 3400+ times in a year, which caused my acceptance rate to become extremely high as well. 😢

Of course, if there are other more elegant solutions that can also beats 100%, you are welcome to submit a PR, and the author will learn together with everyone.

## Target Readers

Programming enthusiasts who want to improve their algorithmic ability through LeetCode.


## Programming Language

All algorithms in this book are implemented in Go.


## Instructions for Use

- There is a search bar in the upper left corner of this e-book, which can quickly help you find the chapters and problem numbers you want to read.
- Every page of this e-book is connected to Gitalk, and there is a comment box at the bottom of every page where you can comment. If it is not displayed, please check your network.
- Regarding solutions, the author suggests using them this way: first read the problem yourself and think about how to solve it. If you still have no idea after 15 minutes, then first read the author's solution idea, but do not look at the code. After you have an idea, implement it yourself in code. If you completely do not know how to write it, then look at the code provided by the author, find out exactly where you do not know how to write it, identify the problem and write it down; this is the knowledge gap you need to fill. If you implement it yourself and there are errors after submitting, debug it yourself first. After AC, if it has not reached 100%, also first think by yourself about how to optimize it. If you can optimize every problem to 100% yourself, then after some time you will make great progress. So overall, if you really have no idea, read the solution idea; if you really cannot optimize it to 100%, take a look at the code.

## Interaction and Errata

If anything is omitted in the articles in the book, you are welcome to click the edit button at the bottom of the corresponding page to comment and interact. Thank you for your support and help.

## Finally

Let's start practicing problems together~

![](https://img.halfrost.com/Blog/ArticleImage/hello_leetcode.png)

This work is licensed under the [Attribution-NonCommercial-NoDerivatives (BY-NC-ND) 4.0 International License](https://creativecommons.org/licenses/by-nc-nd/4.0/legalcode.zh-Hans).


The copyrights of all problems in the solutions belong to [LeetCode](https://leetcode.com/) and [LeetCode China](https://leetcode-cn.com/)
