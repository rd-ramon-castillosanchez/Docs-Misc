# Git Notes

## Conventions

- All commands to input in the CLI are shown in code blocks, e.g. `cd $HOME`, `cd ~`.
- Lines starting with `#` inside a code block are comments, not commands.
- Placeholders are written as `[name]` or `<name>`.

## Documentation

- Official documentation: https://git-scm.com/docs

---

## How Tos

### Setup and Configuration

#### Global config

```bash
git config --global user.name "Ramon Castillo"
git config --global user.email ramon@rcastillo.net

# Stores the repo credentials (only do it once when creating the local git).
# Note: credentials are saved in plain text in ~/.git-credentials
git config --global credential.helper store
```

#### Set the default branch name to `main` (otherwise the default is `master`)

```bash
git config --global init.defaultBranch main
```

#### Editor

```bash
git config --global core.editor "vi"
# or
export GIT_EDITOR=vim
# or (VS Code as editor)
git config --global core.editor "code --wait"

# test the configured editor
git config -e
```

#### Query the config

```bash
git config --list
git config -l
git config --list --global
git config --list --local
git config --edit --local
```

#### Aliases to shorten often-run long commands

```bash
git config --global alias.show3 'log --all --graph'
git show3

git config --global alias.showld 'log --graph --decorate --oneline'
git showld

git config --global alias.dog 'log --all --decorate --oneline --graph'
git dog
```

#### Getting help

```bash
git config -h
git help config     # Linux-like man pages, or through the browser on Windows
```

---

### Creating and Cloning Repositories

#### Create a repository

`cd` to the desired location (the directory to track, where the files are), then:

```bash
git init
```

#### How to clone

```bash
# a normal repo to a specific new directory
git clone <url> ../clonerepo

# a bare repo
git clone --bare <url> ../clonedbarerepo

# a specific branch from remote
git clone <url> --branch <branch> --single-branch <folder>

# a repo including its submodules
git clone --recurse-submodules https://github.com/[User/Repository]
```

#### How to push a local repo for the first time

```bash
git remote add origin git@github.com:<url>.git
git push -u origin master

# resetting the remote
git remote set-url origin git@<remoterepo>
```

#### Starting a new GitHub repo from local work in progress

Creating a repo in GitHub with a commit will generate an undesired issue when Git faces unrelated histories.
This is because the brand-new repo in GitHub has an initial commit, which makes it different from any local repo.
It was fixed by cloning the remote and manually rebasing the local.

Another possible solution is to run:

```bash
git pull origin main --allow-unrelated-histories
```

Better: create a new repo in GitHub with no README nor License, in order not to trigger any initial commit,
and then from the local repository (assuming it is already up to date and with the latest code committed):

```bash
git remote add origin https://github.com/rjcastillos/{newreponame}.git
git branch -M main
git push -u origin main
```

Note: to avoid any issues, the local repo's branch is normally `master` and GitHub's default one is `main`,
so before uploading the code rename the local branch:

```bash
git branch -m master main
# same as
git branch --move master main
```

- After the above you might need to change the default main/master branch (see [Set the default branch name](#set-the-default-branch-name-to-main-otherwise-the-default-is-master)).

---

### Files: Tracking, Staging and Committing

#### File states

- Working files
- Staging
- Commit

#### Start tracking a file or all

```bash
git add [filename]
git add .
git add --all
git add -A
```

#### Commit

```bash
# commit all previously added
git commit -m "Comment"

# add and commit all tracked, modified files in the current working directory
git commit -am "Comment"
```

#### Stop tracking / unstage

```bash
# stop tracking a file
git rm --cached [filename]

# file back from staged to working
git restore --staged [filename]
```

#### Remove, rename and restore files

```bash
# remove a file
git rm [filename]

# rename a file
git mv [oldfilename] [newfilename]

# restore a file that was deleted (it has to be committed to Git before being deleted)
git restore "[filename]"

# restore a file to the last commit point
git restore [filename]
```

#### See the changes of a file

```bash
git diff [filename]
```

#### Ignore a file

- Create the file `.gitignore`
- Templates: https://github.com/github/gitignore

---

### Branches

```bash
# query which branches exist locally
git branch

# create a branch
git branch [branchname]

# list remote (-r) or all (-a) branches, local and remote
git branch -r
git branch -a

# change (switch) to the working branch
git switch [branchtoswitchto]
git checkout [branchtoswitchto]

# create and switch to a new branch
git switch -c [newbranchname]

# rename (move) a branch
git branch -m [oldbranch] [newbranch]

# merge a branch into the one currently switched to
git merge -m "comment" [branchtomerge]

# delete a local branch
git branch -d [branchnametodelete]

# delete a remote branch
git push origin --delete [remotebranch]
```

#### Checking out a commit as a branch (code at a particular point in time)

```bash
git checkout -B [branchname] [commithash]
git checkout -b repo 4c34db222a9241d4438cef2fbedac0b1072c47aa
```

---

### Remotes

```bash
# query the remote origin
git remote -v
git remote show origin     # (tests the connection/security)

# bring commits from other contributors (fetch only; use git pull to also merge them)
git fetch origin

# update origin from my branch
git push origin main
```

---

### Status and History

```bash
# shows which branch you are on and the state of the files
git status

# info about the branch and remote
git branch -av
git status -sb

# see the commit history
git log
git log --oneline
git log --graph --decorate --oneline
git reflog          # all movements of HEAD, chronologically

# see info about a particular commit; shows *ALL* modifications contained in that commit
git show [commithash]
```

---

### Undoing Changes

#### Revert a commit

```bash
git revert [Commit ID shown on log / hash]
```

#### Reset a commit

Move HEAD back and discard all commits after it (and any uncommitted changes):

```bash
git reset --hard [Commit ID shown on log]
git push --force origin master     # rewrites the remote history
```

#### How to address the push error "Updates were rejected because the remote contains work that you do not have locally"

```bash
git reset --soft HEAD~1
```

Moves HEAD back 1 commit, keeping the changes staged, ready to be committed again.
This is useful when the origin contains changes that are not local; then you can `git pull`, update your base and continue with your work.

#### Putting away the current changes while pulling from origin, to later put them back

```bash
git stash
git pull
git stash pop
```

---

### Rewriting History

#### Amend the last commit

```bash
# change only the message of the last commit
git commit --amend -m "[new text]"

# add newly changed (staged) files to the last commit without a new message (forgot in the last commit)
git commit --amend --no-edit
```

- For older commits use interactive rebase.

#### Rewriting commits before pushing or merging code to an upstream branch

Use interactive rebase (`x` = number of commits back from HEAD to include in the range to manipulate):

```bash
git rebase -i HEAD~x
git rebase --interactive HEAD~4
```

- Never rebase a branch that is already pushed.
- More details (from min 4:10): https://youtu.be/qsTthZi23VE?si=9YRsCEz9VO266yi9

#### Cherry-pick

Takes a single commit (hash) from a branch (most likely it was an error to commit it there in the first place)
and applies it to the checked-out branch you are on:

```bash
git cherry-pick 26b1f84a
```

- Word of warning: in the branch where the commit was wrong you can reset (deleting the "last" wrong commit) or hard reset up to a hash (commit).

---

### Merge Conflicts

They appear when a file has been modified and committed in 2 branches, i.e. a file has been updated
(e.g. on master) while also working on it in a branch, so it is different from when the branch was created.

The error occurs when trying to merge a branch with its parent (master or main branch). Example:

1. Change the file in the new branch:

   ```bash
   git switch newbranch
   vi readme.txt          # change text
   git commit -a -m "updating readme.txt file"
   ```

2. Change the same file in the master branch:

   ```bash
   git switch masterbranch
   vi readme.txt          # change text
   git commit -a -m "updating readme.txt file"
   ```

3. On the master branch:

   ```bash
   git merge newbranch
   ```

4. Result: `Merge conflict in readme.txt`

   In the master branch the file(s) with a conflict will show both changes inside, in this way:

   ```text
   <<<<<<< HEAD
   Text on file
   =======
   Text changed on new branch
   >>>>>>> newbranch
   ```

5. Manually modify `readme.txt`, taking out the extra markers inserted by Git and choosing what to leave.
6. Once `readme.txt` is manually fixed:

   ```bash
   git commit -a -m "updating readme.txt file"
   ```

---

### Tags

```bash
# list tags
git tag -l

# create and show a tag
git tag v1.0.0
git show v1.0.0

# push a tag
git push origin v1.0.0

# delete a tag locally
git tag -d [tagname]

# delete a tag on the remote
git push origin -d [tagname]

# tag an existing commit (when you forgot): add the commit hash at the end
git tag -a 1.0.0 -m "Flat version pre packaging" 2cedf23c7ef58a57fa0a1de45000fc708762c4dd

# show the tag's commit
git show-ref -s mytag
```

- Version numbering follows https://semver.org/

---

### SSH

#### Generating an SSH key

```text
C:\Users\A233C1F>ssh-keygen -t ed25519 -C "ramon.castillosanchez@rewe-group.com"
Generating public/private ed25519 key pair.
Enter file in which to save the key (C:\Users\A233C1F/.ssh/id_ed25519):
Created directory 'C:\\Users\\A233C1F/.ssh'.
Enter passphrase (empty for no passphrase):
Enter same passphrase again:
Your identification has been saved in C:\Users\A233C1F/.ssh/id_ed25519
Your public key has been saved in C:\Users\A233C1F/.ssh/id_ed25519.pub
The key fingerprint is:
SHA256:xxxxxxx ramon.castillosanchez@rewe-group.com
The key's randomart image is:
+--[ED25519 256]--+
|                 |
|   .   . .       |
|    o . o        |
|   ..o+o.        |
|   .oX+=S .      |
|    =oO+B=..     |
|   o o &o+..     |
|    o O B..      |
|   ..+.=.*E      |
+----[SHA256]-----+
```

#### Checking SSH

```bash
ssh -T git@gitlab.com
```

#### Copy the public key to the clipboard (Windows)

```text
C:\Users\A233C1F\.ssh>clip < id_ed25519.pub
```

---

## Workflows

### Git with multiple devs

After checking out a local branch, the quicker you merge back to main the better;
otherwise keep it updated as often as possible, resolving any merge conflicts sooner rather than later.

Process, at least daily:

```bash
# from your local branch go to your local master branch
git checkout master

# pull all of the changes to your local master
git pull origin master

# go back to your local working branch
git checkout feature/workingon

# and then merge it with your local master
git merge master
```

Keep doing the above until ready to push:

```bash
git push origin feature/workingon
```

### Git Flow

- Main branches
  - `master`
  - `develop`
- Supporting branches
  - `hotfix/`
  - `feature/`
  - `bugfix/`
  - `release/`
  - `support/`
- Tags
  - Version tag prefixes

After a hotfix is finished, Git Flow will merge it into `master` and into `develop`.

```bash
# initialize a Git Flow repository
git flow init

# start a hotfix
git flow hotfix start '0.1.1'

# when finished
git flow hotfix finish '0.1.1'
```

### DevOps flow

#### DevOps rd-mine workflow

```bash
git clone [repo to work on]
git branch feature/<ticket>-<description>
# ...

# from main
git switch main
git pull

# from the feature branch
git switch feature/<ticket>-<description>
git merge main          # edit the commit message

# before pushing (never rebase a branch already pushed)
git rebase main
```

Alternatively:

```bash
# create a new remote branch from main containing fixes/new code/etc. lately made over main
git push origin main:feature/ticket12

# delete the remote branch if needed
git push origin --delete feature/ticket12

# delete the local tracking of a remote branch
git branch -d -r origin/feature/ticket12
```

#### DevOps suggested

```bash
git clone [repo]
git pull --rebase
```

- Undo if there are merge conflicts:

  ```bash
  git rebase --abort
  git pull              # then fix the merge conflicts
  ```

- Fix manually, merge and rebase:
  - Never rebase a branch already pushed. How to check:

    ```bash
    git branch -a
    ```

  - Interactive rebase (with VS Code as editor: `git config --global core.editor "code --wait"`):

    ```bash
    git rebase --interactive HEAD~4
    ```

    In the `git-rebase-todo` file, change the first commit line:
    - replace `pick` with `edit`
    - make the needed changes and stage them: `git add [filename]`
    - continue:

      ```bash
      git rebase --continue
      ```

---

## Misc

-

---

## Links, Readings & Reference Material

- Install Git: https://www.git-scm.com/ (on Windows use Git Bash)
- VS Code SFTP extension: https://github.com/Natizyskunk/vscode-sftp
- Semantic versioning: https://semver.org/
- Scott Chacon (GitHub co-founder) talk: https://blog.gitbutler.com/fosdem-git-talk/

---

## Notes

-

---

## Versioning History

- v. 1/30/2014
- 2:20 PM 6/29/2017
- 2:34 PM 6/30/2017
- 11:09 AM 4/10/2018
- 12:50 PM 4/24/2018
- 01:02 PM 4/25/2018
- 2:01 PM 5/18/2018
- 1.16.2025
- 2.13.2025
- 9:20 AM 1/29/2026
- 1:20 PM 2/9/2026
- 3.27.2026
- 10.27.2026 converted notes to Markdown and regrouped commands by topic
