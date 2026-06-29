import os
import re
import glob

# file_name = 'Array.md'
reg = "## 题目大意"

current_working_dir = os.getcwd()
# print(f"current_working_dir: {current_working_dir}")

dir_names = glob.glob("*.md")
dir_names.sort()
print(len(dir_names))

for file_name in dir_names:
    # print(file_name)
    with open(file_name, "r") as myfile:
        codeContent = myfile.read()
        findIndex = codeContent.find(reg)
        # print(findIndex)
        content = codeContent[findIndex:]
    with open(file_name, "w") as myfile:
        myfile.write(content)
print("Finished")