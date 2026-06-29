import os
from os.path import join
from shutil import move

toc_string = """\n

---
bookToc: false
---

\n"""

current_working_dir = os.getcwd()
# print(f"current_working_dir: {current_working_dir}")

dir_names = os.listdir(current_working_dir)
# print(dir_names)
print(len(dir_names))
for file_name in dir_names:
        if(file_name.endswith('.md')):
                with open(file_name, "r") as myfile:
                        codeContent = toc_string + myfile.read()
                with open(file_name, "w") as myfile:
                        myfile.write(codeContent)
print("Finished")