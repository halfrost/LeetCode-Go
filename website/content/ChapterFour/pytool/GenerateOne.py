import os
from os.path import join
from shutil import move
import glob

content = []
oneFile = 'ChapterTwo_OnePDF.md'
current_working_dir = os.getcwd()
# print(f"current_working_dir: {current_working_dir}")

dir_names = glob.glob("*.md")
dir_names.sort()
# print(dir_names)
print(len(dir_names))
for file_name in dir_names:
        # print(file_name)
        with open(file_name, "r") as myfile:
                content = myfile.read()
        with open(oneFile, "a") as myfile:
                myfile.write("\n")
                myfile.write(content)
                myfile.write("\n")

print("Finished")