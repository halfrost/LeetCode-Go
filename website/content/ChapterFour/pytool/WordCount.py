from collections import defaultdict
import glob
import os


def str_count2(str):
    count_zh = count_dg = 0
    for s in str:
        # 中文字符范围
        if '\u4e00' <= s <= '\u9fff':
            count_zh += 1
        if s.isdigit():
            count_dg += 1
    # print(count_zh + count_dg)
    return count_zh + count_dg

current_working_dir = os.getcwd()
# print(f"current_working_dir: {current_working_dir}")

dir_names = glob.glob("*.md")
dir_names.sort()

word_count = 0
for file_name in dir_names:
    with open(file_name, "r") as myfile:
        codeContent = myfile.read()
        print("当前读取文件: {}, 字数统计: {}".format(file_name, str_count2(codeContent)))
        word_count += str_count2(codeContent)
print(word_count)