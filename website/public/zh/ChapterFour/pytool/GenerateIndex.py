import os
import glob

content = []
indexFile = 'index.txt'
current_working_dir = os.getcwd()
# print(f"current_working_dir: {current_working_dir}")

dir_names = glob.glob("*.md")
dir_names.sort()
# print(dir_names)
print(len(dir_names))
for file_name in dir_names:
        # - [0001. Two Sum]({{< relref "/ChapterFour/withouttoc/0001. Two Sum.md" >}})
        # content = '- [' + file_name[:-3] + ']' + '({{< relref "/ChapterFour/withouttoc/' + file_name + '" >}})'
        content.append('- [{}]({{{{< relref "/ChapterFour/{}" >}}}})'.format(file_name[:-3],file_name))

with open(indexFile, "w") as myfile:
        myfile.write('\n'.join(content))

print("Finished")