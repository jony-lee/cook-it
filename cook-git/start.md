# 开始学习

工具的学习和应用场景密切相关，脱离应用场景谈工具都毫无意义。因此，我们在学习一项工具的时候就必须了解它是针对现实世界存在的什么样的问题而诞生的，又在哪些场景的应用过程中得到了进一步的发展。

## git的来源
git是一种分布式版本控制系统，其中两个重要的关键词，版本控制和分布式。我们为什么需要版本控制，以写文章为例，一篇好的文章往往不是一蹴而就的，而是在写完之后总是需要修修改改，修改的过程就需要定义大量的修改语言来告知我们某处被修改的前后是什么样的以及为什么要这样修改，而一旦改动量大之后，想要频繁查找原来的修改就变得困难了起来，因此需要版本控制来告诉我们这篇文章的历史发生了什么，又是如何变化至今的。

至于分布式，其对应的是如SVN这样的集中式版本控制系统。
TODO:这部分讲不太清楚，捂脸

这里有一个linus开发git的历史过程，这里暂时先不涉及了。

## git
git add
git branch
git commit
git config
git diff 
git grep  //git grep \<name> 可以对当前项目内所有名为name的内容进行搜索
git log //--stat 添加对文章的统计
git rebase 
git reset //默认--mix即删除暂存区的改动
git stash //暂存内容，将工作区和暂存区内容都保存在一个特殊的地方
git status
git tag
git reflog
git merge
git checkout
git clean //  

![git-diff](./asserts/git-diff.png)
