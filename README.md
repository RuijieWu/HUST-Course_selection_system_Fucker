# HUST-Course_selection_system_Fucker
 抢课用简易脚本,经测试现已成功抢到
require httpx
你也可以把ctrl+f把httpx改成requests,效果是一样的

## 更简易的脚本
这是一个极其简单的发包程序，因此完全可以用更简单的shell程序完成这个任务
```
#! bin/bash
$resp = ''

:(){
    curl -s 'http://wsxk.hust.edu.cn/zxqstudentcourse/zxqcoursesresult.action' -H 'Connection: keep-alive' -H 'Cache-Control: max-age=1' -H 'Origin: http://wsxk.hust.edu.cn' -H 'Upgrade-Insecure-Requests: 1' -H 'DNT: 1' -H 'Content-Type: application/x-www-form-urlencoded' -H 'User-Agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.119 Safari/537.36' -H 'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8' -H 'Referer: http://wsxk.hust.edu.cn/zxqstudentcourse/coursesandclassroom.action?markZB=6&ggkdl=&GGKDLBH=0&skZC=&skJC=&kcmc=%E6%98%93%E7%BB%8F' -H 'Accept-Encoding: gzip, deflate' -H 'Accept-Language: en-US,en;q=0.9' -H 'Cookie: YOUR_COOKIE' --data 'kcbh=1438094&kczxf=2&ktbh=201821438094001&ktrl=200&ktrs=200&markZB=6&kcmc=%E6%98%93%E7%BB%8F%E4%B8%8E%E4%B8%AD%E5%9B%BD%E6%96%87%E5%8C%96' --compressed > $resp
    if grep '选课失败' $resp ; then
        echo 'Failed.'$resp
        :
    else
        echo 'Success.'
        exit 8
    fi
};
:
```
# 使用方式
<<<<<<< HEAD
在config的SESSION里填入你**成功登录后**按F12得到的session值，CLASS_NAME填入你打算抢的课程名字，CLASS_NUMBER填入课堂编号，COURSE_NUMBER填入课程编号，然后运行course_selection_system_fucker.py即可
=======
在config的session里填入你**成功登录后**按F12得到的session值，classname填入你打算抢的课程名字，classnumber然后运行course_selection_system_fucker.py即可
>>>>>>> parent of 4b20475 (0.1.1)

# Tips
学校禁止选课是通过前端实现的(把选课按钮给disbale了)，[选课页面的接口](http://wsxk.hust.edu.cn/zxqstudentcourse/zxqcourses.action)其实依然开放着，例如9.6选课但是9.5你也能通过这个接口访问到选课页面并查看选课情况。

因此，你可以直接前往该页面，进行一次抓包发现报文
```
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/117.0
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8
Accept-Language: zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2
Accept-Encoding: gzip, deflate
Content-Type: application/x-www-form-urlencoded
Content-Length: 111
Origin: http://wsxk.hust.edu.cn
Connection: close
Referer: http://wsxk.hust.edu.cn/zxqstudentcourse/zxqclassroom.action?kcbh=1437248&ggkdl=74&markZB=
Cookie: YOURSESSION0000(以JSESSIONID=起头)
Upgrade-Insecure-Requests: 1
DNT: 1
Sec-GPC: 1

kcbh=1437248&kczxf=2.0&ktbh=202311437248001&ktrl=100&ktrs=100&markZB=&kcmc=%E9%92%A2%E7%AC%94%E7%BB%98%E7%94%BB

//url解码后发现为
kcbh=1437248&kczxf=2.0&ktbh=202311437248001&ktrl=100&ktrs=100&markZB=&kcmc=钢笔绘画
```
参考学校其他系统,以及一点点猜测，不难发现学校的字段名都是汉语拼音的缩写。例如```kcmc```是课程名称，```ktrl```是课堂容量，```ktrs```是课堂人数。`kcbh`是课程编号，`ktbh`是课堂编号。构造POST报文的DATA字段值并发送即可实现选课目的

令人遗憾的是，猜测其他字段值含义的成本恐怕要大于我愿意为此付出的精力，因此这个脚本程序的进一步简易化，自动化要交给别人了，如果你有兴趣，欢迎pr
