#!/usr/bin/env python3
"""Count the total word count of the LeetCode-Go book.

Metric: Chinese characters (CJK ideographs) + English words, counted over the
authored Markdown only (the book content, the per-problem pages and the
top-level docs). Fenced / inline code blocks are stripped and the third-party
Hugo theme and code templates are excluded, so the number reflects written
prose rather than source code or numeric data.

The repo root is resolved from this file's location, so it can be run from any
directory:

    python3 website/content/ChapterFour/pytool/WordCount.py
"""
import glob
import os
import re

HAN = re.compile(r"[一-鿿]")   # CJK unified ideographs (汉字)
EN_WORD = re.compile(r"[A-Za-z]+")     # English words
FENCED = re.compile(r"```.*?```", re.S)  # ```fenced code blocks```
INLINE = re.compile(r"`[^`\n]*`")        # `inline code`

# repo root = four levels up from website/content/ChapterFour/pytool/
ROOT = os.path.abspath(os.path.join(os.path.dirname(__file__), "..", "..", "..", ".."))

# Globs of authored Markdown, relative to the repo root.
GLOBS = [
    "website/content/**/*.md",  # the book《LeetCode Cookbook》
    "leetcode/**/*.md",         # per-problem pages
    "note/*.md",
    "*.md",                     # README and other top-level docs
]


def strip_code(text):
    """Remove fenced and inline code so only prose is counted."""
    return INLINE.sub(" ", FENCED.sub(" ", text))


def count_words(text):
    """汉字 + English words, after stripping code."""
    text = strip_code(text)
    return len(HAN.findall(text)) + len(EN_WORD.findall(text))


def main():
    files = sorted({f for g in GLOBS for f in glob.glob(os.path.join(ROOT, g), recursive=True)})
    total = 0
    for file_name in files:
        with open(file_name, encoding="utf-8", errors="ignore") as fh:
            total += count_words(fh.read())
    print("files counted: {}".format(len(files)))
    print("total word count (汉字 + English words, code excluded): {}".format(total))


if __name__ == "__main__":
    main()
