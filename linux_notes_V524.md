# Linux / Unix Notes

## Conventions

- All commands to input in the CLI are shown in code blocks, e.g. `cd $HOME`, `cd ~`.
- Lines starting with `#` inside a code block are comments, not commands.
- Lines starting with a prompt such as `[root@host ~]#` show an example session (command plus output).

---

## How Tos

### OS and System Information

#### How to query the OS release

```bash
lsb_release -i -r
cat /etc/*release
```

#### Kernel logs

```bash
sudo dmesg | tail -4
```

#### Checking how much memory

```bash
# Solaris
prtconf | grep Memory

# Linux
free -m
top
```

---

### Package Management

#### How to work with rpm (CentOS and RHEL)

| Task | Command |
|------|---------|
| Check the signature of a package | `rpm -K <package>.rpm` |
| Install a package | `rpm -ivh <package>.rpm` |
| List installed RPMs (latest first) | `rpm -qa --last \| tee` |
| Query dependencies | `rpm -qpR machradius-2.1.9-2.i386.rpm` |
| Query what a package provides | `rpm -qp --provides machradius-2.1.9-2.i386.rpm` |
| Query general info about an RPM (when installed) | `rpm -qi machradius` |
| List all the files of a package before it is installed | `rpm -qpl machradius-2.1.9-2.i386.rpm` |
| List all the files of a package already installed | `rpm -ql sudo-1.6.9p17-5.el5` |
| Determine which installed package a particular file belongs to | `rpm -qf /usr/bin/mysql` |
| Upgrade (freshen) a package | `rpm -Fvh mercurial-1.9.2-1.el5.rf.i386.rpm` |

---

### Users, Processes and System Control

#### Users and groups

```bash
groupadd -g 1001 radadmin
useradd -u 1001 -g radadmin radadmin
```

#### How to kill all the processes of oracle (or any user)

```bash
# using cut
ps -ef | grep oracle | cut -d" " -f4 | xargs -i kill -9 {}

# much better so
ps -ef | grep oracle | grep -v grep | awk '{print $2}' | xargs -i kill -9 {}
```

#### Reboot

```bash
init 6
```

---

### Date and Time

#### How to change the timezone (e.g. from CET to UTC)

```bash
dateconfig                  # (RH)
dpkg-reconfigure tzdata     # (Debian)
```

#### How to change the date

```bash
date 033121422003.55
Mon Mar 31 21:42:55 PST 2003
```

The above command does not use the `-s` option, and the fields are arranged like this: `MMDDhhmmCCYY.ss`
where MM = month, DD = day, hh = hour, mm = minute, CCYY = 4-digit year, and ss = seconds.

---

### Files and Archives

#### find

```bash
find ./ -name "19991015*" -print -exec cp {} ./totest \;

find ./ -name 'IMG_E1777*' -exec stat -L {} \;
```

- `-mtime n` : file's data was last modified n*24 hours ago (similar options: `-atime` for last access, `-ctime` for last status change).

##### Sneaking files

```bash
find . -name "*.pl" -exec grep "Comm.pl" {} \;       # lines matched
find . -name "*.pl" -exec grep -l "Comm.pl" {} \;    # (prefix)/files matched
```

#### How to find blank lines in a file

```bash
grep -E '^[[:space:]]*$' file
```

#### How to move files from (anyname) to b.(anyname)

```bash
ls -1 | xargs -I {} mv {} b.{}
```

#### Compressed files

- `zcat` to cat a `.gz` file
- `bzcat` for `bzip2` files

#### tar

```bash
tar -czvf kafka-acl-generator.tgz templates/ .gitignore *.yaml *.go go.* *.md
```

#### How to convert a Markdown document (.md) to a Word document (.docx)

```bash
pandoc inputdocument.md -f gfm -o outputdocument.docx
```

---

### Network Commands

#### netstat

```bash
netstat -nr     # Display routing table
netstat -ta     # Status of the current connections
netstat -i      # Display interface statistics

# Find out the process number of a netstat connection
netstat -anp | grep 93.128.92.137 | grep ESTABLISHED | awk '{print $7}' | awk -F '/' '{print $1}'
```

#### Interfaces and routes

Configuration files: `/etc/sysconfig/network-scripts`

```bash
service network restart

ifconfig eth0 up
ifconfig eth0 10.133.0.17 netmask 255.255.255.0 up

route add default gw 10.133.0.3
route add default netmask 255.255.255.0 gw 10.133.0.3

ipcalc -m -n 10.133.0.3 255.255.255.0
```

#### How to change the MTU

```bash
ifconfig eth1 mtu 1600
echo 'MTU="1600"' >> /etc/sysconfig/network-scripts/ifcfg-eth1

ping 79.170.49.233 -M do -s 1550 -I eth0   # (MTU > 1500)
```

#### tcpdump

```bash
tcpdump
tcpdump -lnXX -s 1500 -i eth1 -w <file.log>
tcpdump -lAn  -s 1500 host 80.187.100.162   # only ASCII
tcpdump -lnXX -s 1500 host 80.187.100.162   # Hex and ASCII
```

- `-s 1500` : snapshot length. The old default of 68 bytes captures only the header of each packet (1500 is the max ~ MTU).
- Some more: http://openmaniak.com/tcpdump.php

#### tshark

```bash
tshark -Y "gtp && (gtp.message != 0x01 && gtp.message != 0x02)" -i eth1
```

- In older versions the display filter option was `-R`.
- http://www.codealias.info/technotes/the_tshark_capture_and_filter_example_page

#### Using cURL

```bash
curl -k -v https://mach.onavo.com
```

#### Some iptables commands on RHEL 5.4

```bash
chkconfig --list iptables
/etc/init.d/iptables stop
chkconfig iptables off
chkconfig --list iptables

lsmod | grep ip_tables

iptables -L
iptables -L OUTPUT -v -n -x

/sbin/service iptables stop
```

Examples:

```text
[root@babyface radadmin]# chkconfig --list iptables
iptables        0:off   1:off   2:off   3:off   4:off   5:off   6:off

[root@CGRAMS1 ~]# cat /etc/sysconfig/iptables
# TO ALLOW SSH
-A RH-Firewall-1-INPUT -p tcp --dport 22 -j ACCEPT

[root@ThinkpadX41 ~]# service sshd status
openssh-daemon (pid  2695) is running...

[root@ThinkpadX41 ~]# iptables -L | grep ssh
ACCEPT     tcp  --  anywhere             anywhere            state NEW tcp dpt:ssh
```

---

### SSH

#### How to generate keys on Linux

```bash
ssh-keygen -t ed25519 -C "user@localserver"
```

Stored under `$HOME/.ssh`

#### Copy keys to server

```bash
cat ~/.ssh/<Key ID>.pub | ssh user@server "mkdir -p ~/.ssh && cat >> ~/.ssh/authorized_keys"
```

OR

```bash
ssh-copy-id <user@server>
```

---

### Bash

#### Source .bashrc from .bash_profile (especially on WSL)

```bash
## Loads .bashrc
if [[ -f ~/.bashrc ]] ; then
        . ~/.bashrc
fi
```

#### Shortcuts and redirection

- `Ctrl + r` : reverse search of the command history
- Redirect stderr to stdout:

```bash
command > output.log 2>&1
```

#### Script: loop over files

```bash
for f in *;
   do
     diff $f ../$f;
   done
```

---

### vi

#### Move cursor

| Key | Action |
|-----|--------|
| `h` | left |
| `l` | right |
| `j` | down |
| `k` | up |
| `w` | a word forward |
| `b` | a word backward |
| `0` | beginning of the line |
| `$` | end of the line |
| `Shift + g` | bottom of the file |
| `10 Shift + g` | go to line 10 of the file |
| `Ctrl + b` | one page back at a time |
| `Ctrl + f` | one page forward at a time |
| `Ctrl + n` | next line |

#### Editing

| Key | Action |
|-----|--------|
| `a` | append after the cursor |
| `o` | open a new line below the cursor |
| `r` | replace a character |
| `x` | delete 1 character where the cursor is on |
| `5x` | delete 5 characters from where the cursor is on |
| `dw` | delete a word |
| `dd` | delete a line |
| `4dd` | delete 4 lines |
| `u` | undo last command (e.g. search and replace) |
| `Shift + j` | join the current line and the next (normal mode) |

#### Copy and paste is yanking and putting

| Key | Action |
|-----|--------|
| `Shift + y` | yank (copy) the line |
| `Shift + p` | put (paste) before the cursor |

#### Search

| Key | Action |
|-----|--------|
| `/` | search forward (find a word) |
| `?` | search backward (find a word) |
| `n` | after a search, jump to the next occurrence |

#### Command mode

| Command | Action |
|---------|--------|
| `:` | enter command mode |
| `:1 [Enter]` | go to line 1 |
| `:set` | show all the vi environment variables and their settings |
| `:set nu` | set line numbering on (numbers at the left of lines in the vi view) |
| `:set ic` | set ignore case on (bash and BASH are the same) |
| `:set noic` | go back to case sensitive |
| `:r !date` | execute the Linux command `date` and append its output to the open file |

#### Search and replace

| Command | Action |
|---------|--------|
| `:%s/search_string/replacement_string/g` | replace in the whole file |
| `:3,5s/search_string/replacement_string/g` | only between lines 3 to 5 |
| `:.,$s/search_string/replacement_string/g` | only between the cursor and the end of the file |
| `:%s/\n/ /` | substitute *ALL* new lines in the file, creating only one long line |
| `:%s/^M//g` | remove CONTROL-M (fix Windows to Unix); type `^M` as `Ctrl + v` `Ctrl + m` |
| `:%s/.$//` | remove any character (`.`) exactly at the end of the line (`$`) |

---

## 911 Commands

- Convert a string to Base64:

```bash
echo -n password | base64
```

---

## Misc

-

---

## Links, Readings & Reference Material

- HTML docs: https://www.w3schools.com/html/html_forms_attributes.asp
- IP CIDR block calculator: https://mxtoolbox.com/subnetcalculator.aspx
- Terraform: https://developer.hashicorp.com/terraform/install

---

## Notes

-

---

## Versioning History

- 1.16.2025 notes template date
- 3.4.2026 start merging Unix/Linux notes
- 10.27.2026 converted notes to Markdown and regrouped commands by topic
