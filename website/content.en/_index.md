---
title: Preface
type: docs
---

# Preface

{{< columns >}}
## About LeetCode

Speaking of LeetCode, as a programmer, you should be no stranger to it; in recent years, it has come up in interviews. Programmers both in China and abroad mainly use it to practice problems for interviews. According to historical records, this website was founded in 2011 and is about to reach its 10th anniversary. It holds weekly contests, biweekly contests, and monthly contests. Coding within a limited time is indeed a very good test of one's algorithmic ability. In some competitions sponsored and named by large companies, those who place near the top not only receive prizes, but can also directly get opportunities for referrals.

<--->

## What Is a Cookbook

Literally translated, it is a cooking book, a book that teaches you how to make various recipes and dishes. Students who often read O'Reilly technical books will be very familiar with this term. Generally, hands-on and practice-oriented books have this name.

{{< /columns >}}

<img src="https://books.halfrost.com/leetcode/logo.png" alt="logo" height="600" align="right" style="padding-left: 30px;"/>

## Why Write This Open-Source Book

I have been solving problems for a year and want to share with everyone some insights and methods for solving problems. I want to make friends with people who have the same interests and communicate and learn together. For myself, writing solutions is also a way to improve. Explaining a profound problem to someone who has no clue at all, and being able to make them fully understand it, is a great exercise in expression. During the explanation, you may very likely encounter questions from the listener; these questions may be gaps in your own knowledge, forcing you to fill them. I have given related talks at the company and felt this deeply; both sides benefited quite a lot.

> In addition, during college, when I solved problems, I hated writing solutions the most. I felt it was a waste of time and that I should spend more time solving more problems. Now I don't know whether this counts as “what goes around comes around”.


## About the Book Cover

Students who often read O'Reilly animal books will know at a glance that this cover is a tribute to them. That is indeed the purpose. The animals on O'Reilly covers are all rare animals, and the drawing style is black-and-white sketching. These animals are all copyrighted, so I could only search online for copyright-free black-and-white sketch-style images. Commonly, one can find 40 images in this style. However, too many people use them, so I painstakingly found several other such images, and this peacock spreading its tail is one of them. The meaning of the peacock spreading its tail is the hope that after everyone finishes practicing LeetCode, they will improve their own algorithmic ability and unfold their own “tail” on the stage of life. The color scheme of the whole book is also green, because this is the color of AC.


## About the Author

I am a gopher newcomer who has only just entered the industry for a year and a half, so I ask all the experts to give this junior more guidance. In college, I participated in ACM-ICPC for 3 years, but because my aptitude was not high, I did not win a single gold medal. So in terms of algorithms, my evaluation of myself would be that I am a beginner. The greatest gain from participating in ACM-ICPC was training my thinking ability, and this ability is also applied in life. The second was getting to know many very smart contestants in China and seeing the gap between myself and them. Finally, there were those 200-plus pages of densely printed [algorithm templates](https://github.com/halfrost/leetcode-go/releases/tag/Special), some of which I did not even fully understand myself. Knowledge that you have learned is yours for life; if you have not learned it, that knowledge is just something external.

I started solving problems on March 25, 2019, and by March 25, 2020, it had been exactly one year. The original plan was one problem per day. In fact, sometimes there was more than one problem per day, and in the end I completed 600+:

![](https://img.halfrost.com/Blog/ArticleImage/2019_leetcode.png)

> A warm tip: I originally thought that doing one problem every day would make this submissions graph entirely green, but I found that I was wrong. If you also want to persist and make this graph entirely green, be sure to pay attention to the following issue: the LeetCode server is in the +0 time zone, and this graph is also calculated according to this time zone. In other words, before 8 a.m. every day in China, it counts as the previous day! It was also because of the time zone issue that I ended up with these 22 blank squares. For example, if there is a Hard problem that is very difficult, and there is also a lot of work that day, by the time I think of the solution after getting home from work at night, it may already be early the next morning. Then I do another problem as the next day's quota. As a result, I would find that both of these 2 problems counted as the previous day. Sometimes I get up at 6 a.m. to solve problems, and after submitting, they are also counted as the previous day.
> 
> (Of course, all of this is in the past and no longer important; just treat it as some small episodes on the road of struggle)

In 2020, I will definitely continue solving problems, because I have not yet reached some of my goals. I may strive toward 1000 problems, or I may start going back for a second pass and third pass when I reach 800 problems. (I probably will not give up until I reach my goal~)

## About the Code in the Book

The code is all placed in the [github repo](https://github.com/halfrost/leetcode-go/tree/master/leetcode), and problems can be searched by problem number.
The code for the problems in this book has all already beaten 100%. Solutions that have not beaten 100% have not been included in this book. I will continue optimizing those problems to 100% before adding them.

Some readers may ask why pursue beats 100%. I believe that only after optimizing to beats 100% can it be considered that I have gotten the feel for this problem. For several Hard problems, I used brute-force solutions to AC them, but they only beat 5%. That is just like not having done the problem at all. Moreover, if you give such an answer in an interview, the interviewer will not be satisfied either: “Is there a more optimal solution?” If you can provide a more optimal solution through your own thinking, the interviewer will be more satisfied.

LeetCode's statistics for code runtime fluctuate; submitting the same code 10 times may result in beats 100%. I did not notice this problem at first, and for many problems I submitted the correct code many times in a row. Over the year, I submitted 3400+ times, which caused my correctness rate to become extremely high as well. 😢

Of course, if there are other more elegant solutions that can also beat 100%, you are welcome to submit a PR, and I will learn together with everyone.

## Target Readers

Programming enthusiasts who want to improve their algorithmic ability through LeetCode.


## Programming Language

All algorithms in this book are implemented in Go.

## Instructions

- There is a search bar in the upper-left corner of this e-book, which can quickly help you find the chapter and problem number you want to read.
- Every page of this e-book is integrated with Gitalk, and there is a comment box at the bottom of each page for comments. If it does not appear, please check your network.
- Regarding solutions, I recommend using them this way: first read the problem yourself and think about how to solve it. If you still have no idea after 15 minutes, then first read my solution approach, but do not look at the code. Once you have an idea, implement it in code yourself. If you completely do not know how to write it, then look at the code I provide, find out exactly where you do not know how to write it, identify the problem and note it down; this is the knowledge gap you need to fill. If you implement it yourself and there are errors after submission, debug it yourself first. After AC, if it has not reached 100%, first think on your own about how to optimize it. If you can optimize every problem to 100% yourself, then after a period of time you will make great progress. So in general, if you really have no ideas, read the solution approach; if you really cannot optimize it to 100%, take a look at the code.

## Interaction and Errata

If there are omissions in the articles in the book, you are welcome to click the edit button at the bottom of the page to comment and interact. Thank you for your support and help.

## Finally

Let's start solving problems together~

![](https://img.halfrost.com/Blog/ArticleImage/hello_leetcode.png)

This work is licensed under the [Attribution-NonCommercial-NoDerivatives (BY-NC-ND) 4.0 International License](https://creativecommons.org/licenses/by-nc-nd/4.0/legalcode.zh-Hans).

The copyrights of all problems in the solutions belong to [LeetCode](https://leetcode.com/) and [LeetCode China](https://leetcode-cn.com/)
 
