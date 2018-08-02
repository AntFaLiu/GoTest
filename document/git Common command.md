# git常用操作
## git init
    用 git init 在目录中创建新的 Git 仓库。 你可以在任何时候、任何目录中这么做，完全是本地化的。
    在目录中执行 git init，就可以创建一个 Git 仓库了。
# git clone
    git clone
    使用 git clone 拷贝一个 Git 仓库到本地，让自己能够查看该项目，或者进行修改。
    如果你需要与他人合作一个项目，或者想要复制一个项目，看看代码，你就可以克隆那个项目。
## git add
    git add 命令可将该文件添加到缓存
## git status 
    查看项目的当前状态
## git dff
    执行 git diff 来查看执行 git status 的结果的详细信息
## git log 
    查看提交历史
## git commit
    使用 git add 命令将想要快照的内容写入缓存区， 而执行 git commit 将缓存区内容添加到仓库中。
## git fetch origin
    同步远程服务器上的数据到本地。
## git branch --merge
    查看哪些分支已被并入当前分支
## git rebase git merge
### git merge
    先切换到要合并到的那个分支上，再使用git merge+要被合并的commit id 
    相当于新提交一个commit记录合并信息
### git rebase
    先找公共祖先，然后将被合并的那一个分支所有新的提交接到合并到的那个分支后面
    冲突处理策略的不同
    merge 遇见冲突后会直接停止，等待手动解决冲突并重新提交 commit 后，才能再次 merge
    rebase 遇见冲突后会暂停当前操作，开发者可以选择手动解决冲突，然后 git rebase --continue 继续，或者 --skip 跳过（注意此操作中当前分支的修改会直接覆盖目标分支的冲突部分），亦或者 --abort 直接停止该次 rebase 操作
### 合并多次提交
    在使用git rebase -i + 分支别名 合并
    git rebase -i HEAD~4
    git log --pretty=oneline
    git 将修复上一个commit的comment
## 当前有修改没有提交，但现在有个更重要的任务要做，要先保存本地修改，然后切到master，新建一个分支来做新任务
    git stash    暂存修改
    git stash pop  会把暂存的修改第一个pop出来
    git stash list 可以查看git栈的内的所有备份
## 将一个分支的commit移到另一个分支
    feature 分支上的commit 62ecb3 非常重要，它含有一个bug的修改，或其他人想访问的内容。无论什么原因，你现在只需要将62ecb3 合并到master，而不合并feature上的其他commits，所以我们用git cherry-pick命令来做
    git checkout master  
    git cherry-pick 62ecb3  
## 回退提交
    $ git reset --hard HEAD^         回退到上个版本
    $ git reset --hard HEAD~3        回退到前3次提交之前，以此类推，回退到n次提交之前
    $ git reset --hard commit_id     退到/进到 指定commit的sha码
## 创建并且切换分支：
    git checkout -b dev
    相当于：
    $ git branch dev
    $ git checkout dev
## 查看本地分支
    git branach
## 删除本地分支
    git branach -d
## 查看远程分支
    $ git branch -a
## 删除远程分支
    git push origin  :test
## git强推
    $ git push origin HEAD --force
    git push -f
## 删除远程分支：
    git push origin --delete test
    git push origin  :test
## 从当前分支切换到‘dev’分支：
    git checkout dev
## 建立并切换新分支：
    git checkout -b 'dev'
## 查看当前详细分支信息（可看到当前分支与对应的远程追踪分支）:
    git branch -vv
## 查看当前远程仓库信息
    git remote -vv
## Git本地分支与远程分支的追踪关系
    git push origin [本地分支名]:[远程分支名]

