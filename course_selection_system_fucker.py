'''
和这牛魔选课系统爆了 :P
'''
import sys
import threading
import time
import httpx
import config
#! 构造请求报文
HEADER = {
    "User-Agent"                :   "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/117.0",
    "Accept"                    :   "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8",
    "Accept-Language"           :   "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2",
    "Accept-Encoding"           :   "gzip, deflate",
    "Content-Type"              :   "application/x-www-form-urlencoded",
    "Origin"                    :   "http://wsxk.hust.edu.cn",
    "Connection"                :   "close",
    "Referer"                   :   "http://wsxk.hust.edu.cn/zxqstudentcourse/zxqclassroom.action?kcbh=1437248&ggkdl=74&markZB=",
    "Cookie"                    :   config.SESSION,
    "Upgrade-Insecure-Requests" :   "1",
    "DNT"                       :   "1",
    "Sec-GPC"                   :   "1"
}

DATA = {
    "kcbh"  :config.CLASS_NUMBER,
    "kczxf" :"2.0",
    "ktbh"  :"202311437248001",
    "ktrl"  :"100",
    "ktrs"  :"99",
    "markZB":"",
    "kcmc"  :config.CLASS_NAME
}

API_URL = "http://wsxk.hust.edu.cn/zxqstudentcourse/zxqcoursesresult.action"

def sign()->None:
    '''send post request and judge if it is successful'''
    try:
        resp = httpx.post(url=API_URL,headers=HEADER,data=DATA,timeout=10)
    except httpx.RequestError as err:
        print(err)
    if resp:
        if resp.status_code == 200:
            html = resp.text
            #print(html)
            if "选课失败，课堂人数已满！" in html:
                print(time.ctime())
            else:
                sys.exit()

if __name__ == "__main__":
    try:
        while True:
            client_handler = threading.Thread(target=sign)
            client_handler.start()
            #! 请求过于频繁会被HUB系统当作DDOS
            time.sleep(10)
    except KeyboardInterrupt:
        print("Sign in stopped")
