🐛🐙🐜百度贴吧的爬虫,可以爬去网页的信息并保存.
## 过程分析

### 1、获取url
> golang主页信息:https://tieba.baidu.com/f?ie=utf-8&kw=golang  
> 第二页:https://tieba.baidu.com/f?kw=golang&ie=utf-8&pn=50  
> 第三页:https://tieba.baidu.com/f?kw=golang&ie=utf-8&pn=100  
> 第四页:https://tieba.baidu.com/f?kw=golang&ie=utf-8&pn=150  
> 最后一页:就是第五页,其他的都看不了.  

### 2、其中kw表示搜索的关键字,pn表示(页数)0,50,100,150,200

> 使用的是横向爬取,按照一页一页的爬取.  
> 使用的是net/http的包.模拟浏览器的请求,具体的步骤:  
>> 1)构建发送请求链接 使用的是http.get(url)  
>> 2)获取服务器相应的信息 resp.body.read([]buffer) result+=string  (buffer[:n]) 别忘了关闭 resp.body.close()  
>> 3)过滤保存使用得到的信息  
>> 4)关闭请求链接 每次保存文件之后就是用file.close  
> 文件爱你的保存需要使用os的Getwd() (dir string, err error)    //获取当前目录，类似linux中的pwd  

### 3、并发爬取信息  

> 使用的是go和chan int  
> for循环内使用go func 并且将chan进行传入(使用管道是因为防止主协程先结束,达到主次同步)  
> 子协程执行完之后要chan<-i 写入一个标识量到管道中  
> 主协成要使用同样的for循环,阻塞读取每一次的chan传入的数据 <-chan int  