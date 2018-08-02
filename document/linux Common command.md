# linux 常见命令
## chmod
    属主的权限为：rw-
    属组的权限为：r--
    其他人的权限为：r--
    属主：4(读) + 2 (写) + 0(执行) = 6
    属组：4(读) + 0(写) + 0(执行) = 4
    其他人：0(读) + 0(写) + 0 (执行)= 0 
## netstat [-选项]
### netstat各选项参数说明： 
    -a : 列出所有连接，服务监听，Socket信息 
    -c : 持续列出网络状态 #每隔一秒输出网络信息 
    -t : 显示TCP端口 
    -u : 显示UDP端口 
    -l : 列出当前监听服务，只显示监听端口 
    -p : 显示PID（进程号）和进程名称 
    -n : 直接使用ip地址，而不通过域名服务器(加速输出，因为不用进行比对查询) 
    -s : 显示网络工作信息统计表 
    -r : 显示路由表信息
## netstat mac上的操作
### 列出所有端口 
    netstat 
### 列出所有tcp 端口 
    netstat -at
### 显示网络接口列表
    netstat -i
### 显示网络工作信息统计表
    netstat -s
### 显示伪装的网络连线
    netstat -m
### 显示核心路由信息
    netstat -r
### 显示合并的信息
    netstat -rs
##  查看域名对应的ip
    无论在windows还是在linux下都是用 nslookup
##  dig 
    查看详细的域名解析过程。
## 跟踪路由
  windows中使用 tracert，linux中使用 traceroute
  要查看dns,tcp connect, http rep/rsp中每个阶段的耗时呢？
## 如果判断某个ip是否连上？某个ip:port是否能连上？某个url:port是否能提供务？
### 方法一：ssh 
    SSH 是目前较可靠，专为远程登录会话和其他网络服务提供安全性的协议,在linux上可以通过ssh命令来测试端口的连通性，具体用法格式如下：
#### 用法: ssh -v -p port username@ip
    v说明：
        -v 调试模式(会打印日志).
        -p 指定端口
        username:远程主机的登录用户
        ip:远程主机
        如果远程主机开通了相应的端口，会有如下图所示的建立成功的提示。
### 方法二、：telnet法
    telnet为用户提供了在本地计算机上完成远程主机工作的能力，因此可以通过telnet来测试端口的连通性，具体用法格式：
    telnet ip port
    telnet 127.0.0.1  3306
    说明：
    ip：是测试主机的ip地址
    port：是端口，比如80
    如果telnet连接不存在的端口，那会如下图所示。
    如果telnet 连接存在端口会出现如下图所示的内容，下图中以80端口为例。
### 方法三、curl法
    curl是利用URL语法在命令行方式下工作的开源文件传输工具。也可以用来测试端口的连通性，具体用法:
    curl ip:port
    说明：
    ip：是测试主机的ip地址
    port：是端口，比如80
    如果远程主机开通了相应的端口，都会输出信息，如果没有开通相应的端口，则没有任何提示，需要CTRL+C断开。
### 方法四、wget方法
    wget是一个从网络上自动下载文件的自由工具，支持通过HTTP、HTTPS、FTP三个最常见的TCP/IP协议下载，并可以使用HTTP代理。wget名称的由来是“World Wide Web”与“get”的结合，它也可以用来测试端口的连通性具体用法:
    wget ip:port
    说明：
    ip：是测试主机的ip地址
    port：是端口，比如80
    如果远程主机不存在端口则会一直提示连接主机。
## telnet(选项)(参数)
    -8：允许使用8位字符资料，包括输入与输出；
    -a：尝试自动登入远端系统；
    -b<主机别名>：使用别名指定远端主机名称；
    -c：不读取用户专属目录里的.telnetrc文件；
    -d：启动排错模式；
    -e<脱离字符>：设置脱离字符；
    -E：滤除脱离字符；
    -f：此参数的效果和指定"-F"参数相同；
    -F：使用Kerberos V5认证时，加上此参数可把本地主机的认证数据上传到远端主机；
    -k<域名>：使用Kerberos认证时，加上此参数让远端主机采用指定的领域名，而非该主机的域名；
    -K：不自动登入远端主机；
    -l<用户名称>：指定要登入远端主机的用户名称；
    -L：允许输出8位字符资料；
    -n<记录文件>：指定文件记录相关信息；
    -r：使用类似rlogin指令的用户界面；
    -S<服务类型>：设置telnet连线所需的ip TOS信息；
    -x：假设主机有支持数据加密的功能，就使用它；
    -X<认证形态>：关闭指定的认证形态。
## 此进程建立的所有连接的对端的ipport

##  查看某一端口的占用情况
### linux 
    netstat -apn |grep 8080
### mac 
    lsof -i:8080
##  查看本机磁盘消耗
### mac  df:
    -a      Show all mount points, including those that were mounted with the MNT_IGNORE flag.
    -b      Use (the default) 512-byte blocks.  This is only useful as a way to override an BLOCKSIZE specification from the environment.
    -g      Use 1073741824-byte (1-Gbyte) blocks rather than the default.  Note that this overrides the BLOCKSIZE specification from the environment.
    -H      "Human-readable" output.  Use unit suffixes: Byte, Kilobyte, Megabyte, Gigabyte, Terabyte and Petabyte in order to reduce the number of digits to three or less using base 10 for sizes.
    -h      "Human-readable" output.  Use unit suffixes: Byte, Kilobyte, Megabyte, Gigabyte, Terabyte and Petabyte in order to reduce the number of digits to three or less using base 2 for sizes.
    -i      Include statistics on the number of free inodes. This option is now the default to conform to Version 3 of the Single UNIX Specification (``SUSv3'') Use -P to suppress this output.
    -k      Use 1024-byte (1-Kbyte) blocks, rather than the default.  Note that this overrides the BLOCKSIZE specification from the environment.
    -l      Only display information about locally-mounted filesystems.
    -m      Use 1048576-byte (1-Mbyte) blocks rather than the default.  Note that this overrides the BLOCKSIZE specification from the environment.
### linux   df：
    -a：显示所有文件系统的磁盘使用情况，包括0块（block）的文件系统，如/proc文件系统。
    -k：以k字节为单位显示。
    -i：显示i节点信息，而不是磁盘块。
    -t：显示各指定类型的文件系统的磁盘空间使用情况。
    -x：列出不是某一指定类型文件系统的磁盘空间使用情况（与t选项相反）。
    -T：显示文件系统类型。
## 查看本机的磁盘空间消耗，当前目录每个文件夹的空间是多少
    df        du + 具体文件名
    1. df -lh 
    2. du -s /usr/* | sort -rn 
    这是按字节排序 
    3. du -sh /usr/* | sort -rn 
    这是按兆（M）来排序 
    4.选出排在前面的10个 
    du -s /usr/* | sort -rn | head 
    5.选出排在后面的10个 
    du -s /usr/* | sort -rn | tail 
## 怎么看dns解析
    使用dig + 网址
## 查看进程信息
### linux top  显示所有进程
    -b 批处理
    -c 显示完整的治命令
    -I 忽略失效过程
    -s 保密模式
    -S 累积模式
    -i<时间> 设置间隔时间
    -u<用户名> 指定用户名
    -p<进程号> 指定进程
    -n<次数> 循环显示的次数
### mac htop 显示所有
    -d --delay = DELAY
      更新之间的延迟，以十分之一秒为单位
    -C - 无色 - 无色
      以单色模式启动htop
    -h --help
      显示帮助消息并退出
    -p --pid = PID，PID ......
      仅显示给定的PID
    -s --sort-key COLUMN
     按此列排序（使用列表列表的--sort-key帮助）
### ps: 查看具体pid的进程信息
### ps -ef |grep
    -A显示有关其他用户进程的信息，包括那些没有控制终端的进程。
    -a显示有关其他用户进程以及您自己的进程的信息。 除非还指定了-x选项，否则这将跳过任何没有控制终端的进程。
    -C通过使用忽略“常驻”时间的“原始”CPU计算来改变计算CPU百分比的方式（这通常没有效果）。
    -c将``command''列输出更改为只包含可执行文件名，而不是完整的命令行。
    -d与-A类似，但不包括会话领导者。
    -E也显示环境。 这并不反映流程启动后环境的变化。
    某个ip:port是否能连上
## 查看从本机访问到指定IP/网址经历的路由器
### mac 
    traceroute
### windoows 
    tracert
## 一次访问网址的过程
(1)浏览器获取输入的域名www.baidu.com
(2) 浏览器向DNS请求解析www.baidu.com的IP地址
(3) 域名系统DNS解析出百度服务器的IP地址
(4) 浏览器与该服务器建立TCP连接(默认端口号80)
(5) 浏览器发出HTTP请求，请求百度首页
(6) 服务器通过HTTP响应把首页文件发送给浏览器
(7) TCP连接释放
(8) 浏览器将首页文件进行解析，并将Web页显示给用户。
涉及到的协议
(1) 应用层：HTTP(WWW访问协议)，DNS(域名解析服务)
(2) 传输层：TCP(为HTTP提供可靠的数据传输)，UDP(DNS使用UDP传输)
(3) 网络层：IP(IP数据数据包传输和路由选择)，ICMP(提供网络传输过程中的差错检测)，ARP(将本机的默认网关IP地址映射成物理MAC地址)
git stash  //暂存修改工作区
git pull  //拉取远端文件
git stash pop stash@{0}    //将工作区还原

## 抓包 tcpdump
## cat
### 查看日志中的特定字符串
cat filename | grep "dasdasdas"
## curl
###  使用curl查看HTTP请求各个时间阶段的耗时
    1。新建一个文件vim curl
    将下面的内容复制进去然后保存：
    \n
            time_namelookup:  %{time_namelookup}\n
               time_connect:  %{time_connect}\n
            time_appconnect:  %{time_appconnect}\n
           time_pretransfer:  %{time_pretransfer}\n
              time_redirect:  %{time_redirect}\n
         time_starttransfer:  %{time_starttransfer}\n
                            ----------\n
                 time_total:  %{time_total}\n
    \n
    语句：curl -w "@curl" -o /dev/null -s http://www.baidu.com
    2.如果curl这个文件在其他位置，需要将path写在curl前面：
    语句：curl -w "@path/curl" -o /dev/null -s http://www.baidu.com

## sed 命令
    sed命令在mac环境下，与linux有点不一样：
    1，需要在sed命令后面加上''引号
    2，添加的文本需要换行，必须在文本后面加上\且按回车才有换行的效果
    mac:
        sed  -i '' 's/n/en/g' t.txt
    linux: 
        sed  -i 's/n/en/g' t.txt

## awk打印指定行
    awk -F"," '{if($1=="505888"){print $0}}' mdl_13_14_0.csv
    $0           表示整个当前行
    $1           每行第一个字段
    NF          字段数量变量
    NR          每行的记录号，多文件记录递增
    FNR        与NR类似，不过多文件记录不递增，每个文件都从1开始
    \t            制表符
    \n           换行符
    FS          BEGIN时定义分隔符
    RS       输入的记录分隔符， 默认为换行符(即文本是按一行一行输入)
    ~            匹配，与==相比不是精确比较
    !~           不匹配，不精确比较
    ==         等于，必须全部相等，精确比较
    !=           不等于，精确比较
    &&　     逻辑与
    ||             逻辑或
    +            匹配时表示1个或1个以上
    /[0-9][0-9]+/   两个或两个以上数字
    /[0-9][0-9]*/    一个或一个以上数字
    FILENAME 文件名
## linux 批量重名名 Linux Shell 批量重命名的方法总览
    1、删除所有的 .bak 后缀：
    rename 's/\.bak$//' *.bak
    2、把 .jpe 文件后缀修改为 .jpg：
    rename 's/\.jpe$/\.jpg/' *.jpe
    3、把所有文件的文件名改为小写：
    rename 'y/A-Z/a-z/' *
    4、将 abcd.jpg 重命名为 abcd_efg.jpg：
    for var in *.jpg; do mv "$var" "${var%.jpg}_efg.jpg"; done
    5、将 abcd_efg.jpg 重命名为 abcd_lmn.jpg：
    for var in *.jpg; do mv "$var" "${var%_efg.jpg}_lmn.jpg"; done
    6、把文件名中所有小写字母改为大写字母：
    for var in `ls`; do mv -f "$var" `echo "$var" |tr a-z A-Z`; done
    7、把格式 *_?.jpg 的文件改为 *_0?.jpg：
    for var in `ls *_?.jpg`; do mv "$var" `echo "$var" |awk -F '_' '{print $1 "_0" $2}'`; done
    8、把文件名的前三个字母变为 vzomik：
    for var in `ls`; do mv -f "$var" `echo "$var" |sed 's/^.../vzomik/'`; done
    9、把文件名的后四个字母变为 vzomik：
    for var in `ls`; do mv -f "$var" `echo "$var" |sed 's/....$/vzomik/'`; done
    for var in `ls *_?.jpg`; do mv "$var" `echo "$var" |awk -F '_' '{print $1 "_0" $2}'`; done


    update daily_push_overall_statistics set udated_at = DATE_SUB(updated_at,interval,1 day)  where date < '2018-07-30';