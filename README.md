# 代码提交规则

## 比如共识的提交样例：(3段，每段用空格隔开)
```
第一段 标题
第二段 主要修改内容
第三段 主要修改区域，目录路径
```
eg:
```
Consensus: Create a witness list 

1.A total of 12 witnesses.
2.
3.

area: bbbOfChain/consensus
```


## 提交方式：
```
1. 保存修改的内容
git add .

2. 添加提交记录
git commit -a

3. 再次修改提交记录
git commit --amend

4. 提交到对应的远程分支，比如提到我的本地分支0-conDev到远程分支Consensus
git push origin 0-conDev:Consensus
```


