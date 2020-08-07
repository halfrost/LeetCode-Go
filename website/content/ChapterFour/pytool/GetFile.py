import os
from os.path import join
from shutil import move

toc_string = """\n

---
bookToc: false
---

\n"""

pre_string = """

## 代码

```go

"""
end_string = "\n```"
codeContent = ""
current_working_dir = os.getcwd()
# print(f"current_working_dir: {current_working_dir}")

# dir_names = os.listdir(current_working_dir)

# 重命名目录。 os.rename(dir_name, new_dir_name)

# 遍历目录
# dir_path是当前遍历到的目录。dir_names是dir_path下的文件夹列表。file_names是是dir_path下的文件列表
# 如果想实现目录白名单，将白名单目录从dir_names中去除即可
for (dir_path, dir_names, file_names) in os.walk(current_working_dir):
    # print(f"当前遍历到的目录: {dir_path}")
    os.chdir(dir_path)
    files = dir_path.split("/")
    new_file_name = files[len(files) - 1] + '.md'
    
    print(f"当前所在的文件夹: {os.getcwd()}")
    for file_name in file_names:
        if(file_name.endswith('.go') and not file_name.endswith('_test.go')):
            print(f"当前所在文件: {file_name}")
            with open(file_name, "r") as myfile:
                codeContent = myfile.read()
            break
            # print(codeContent)
    for file_name in file_names:
        if(file_name.endswith('.md')):
            print(f"当前所在文件: {file_name}")
            with open(file_name, "a") as myfile:
                myfile.write(pre_string)
                myfile.write(codeContent)
                myfile.write(end_string)
            os.rename(file_name, new_file_name)
            move(join(dir_path, new_file_name), current_working_dir)
            break
    