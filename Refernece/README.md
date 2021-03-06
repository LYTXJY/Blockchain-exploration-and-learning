 

 

 

 

 

 

 

**GO****语言核心编程**

 

**https://github.com/LYTXJY/Blockchain-exploration-and-learning**

 

![img](file:///C:/Users/xjy/AppData/Local/Temp/msohtmlclip1/01/clip_image002.gif)

 

 

 

海阔凭鱼跃，山高任鸟飞

 

 

**
**

目 录

​                  

[目 录. i](#_Toc32397483)

[第1章 Golang开山篇. 1](#_Toc32397484)

[1.1Golang的学习方向. 1](#_Toc32397485)

[1.2Golang的应用领域. 1](#_Toc32397486)

[1.2.1区块链应用. 1](#_Toc32397487)

[1.2.2后端服务器应用. 1](#_Toc32397488)

[1.2.3云计算/云服务后台应用. 2](#_Toc32397489)

[1.3Golang的学习方法. 2](#_Toc32397490)

[1.3.1主要研究内容. 2](#_Toc32397491)

[第2章 Golang的概述. 3](#_Toc32397492)

[2.1什么是程序. 3](#_Toc32397493)

[2.2GO语言的诞生. 3](#_Toc32397494)

[2.2.1为什么要创造Go语言. 3](#_Toc32397495)

[2.2.2GO语言发展简史. 4](#_Toc32397496)

[2.3GO语言的特点. 4](#_Toc32397497)

[2.4GO语言的开发工具. 4](#_Toc32397498)

[2.4VS code安装和使用. 5](#_Toc32397499)

[2.5Windows下GO语言开发环境搭建. 5](#_Toc32397500)

[2.5.1SDK介绍. 5](#_Toc32397501)

[2.5.2SDK下载. 6](#_Toc32397502)

[2.5.3SDK安装. 6](#_Toc32397503)

[2.5.4SDK环境变量配置. 7](#_Toc32397504)

[2.3GO语言的特点. 8](#_Toc32397505)

[2.3GO语言的特点. 9](#_Toc32397506)

[2.3GO语言的特点. 9](#_Toc32397507)

[第3章 基于OpenPose模型改进的2D平面多人姿态识别方法. 10](#_Toc32397508)

[3.1引言. 10](#_Toc32397509)

[3.2改进的OpenPose模型框架. 10](#_Toc32397510)

[3.2.1主干网络残差设计. 10](#_Toc32397511)

[3.2.1.1基于ResNet残差网络改进的OpenPose主干网络. 10](#_Toc32397512)

[3.2.1.2 VGG主干网络与ResNet残差主干网络的对比分析. 10](#_Toc32397513)

[3.2.2网络裁剪与卷积块堆叠. 11](#_Toc32397514)

[3.2.2.1 OpenPose上下分支网络与改进版的对比分析. 11](#_Toc32397515)

[3.3人体关节点的识别. 11](#_Toc32397516)

[3.3.1单人情形. 11](#_Toc32397517)

[3.3.2多人情形. 11](#_Toc32397518)

[3.4人体关节点的组合. 12](#_Toc32397519)

[3.4.1单人情形. 12](#_Toc32397520)

[3.4.2多人情形. 12](#_Toc32397521)

[3.5网络训练. 12](#_Toc32397522)

[3.5.1数据集准备. 12](#_Toc32397523)

[3.5.2超参数设置. 12](#_Toc32397524)

[3.6网络训练实验结果与分析. 12](#_Toc32397525)

[3.6.1模型评估. 12](#_Toc32397526)

[3.6.2模型测试. 13](#_Toc32397527)

[3.6.2.1图像测试. 13](#_Toc32397528)

[3.6.2.2视频测试. 13](#_Toc32397529)

[3.7本章小结. 13](#_Toc32397530)

[第4章 人体动作识别应用. 14](#_Toc32397531)

[4.1引言. 14](#_Toc32397532)

[4.2人体动作识别相关研究. 14](#_Toc32397533)

[4.3应用OpenPose模型的人体动作识别实现. 14](#_Toc32397534)

[4.4本章小结. 14](#_Toc32397535)

[第5章 总结与展望. 15](#_Toc32397536)

[5.1工作总结. 15](#_Toc32397537)

[5.2研究展望. 15](#_Toc32397538)

[参考文献. 16](#_Toc32397539)

[致谢. 17](#_Toc32397540)

[附录. 18](#_Toc32397541)



**
**

第1章 Golang开山篇

# 1.1Golang的学习方向

GO语言，又称Golang，是“GO Language”的简写。

![img](file:///C:/Users/xjy/AppData/Local/Temp/msohtmlclip1/01/clip_image004.jpg)

1.区块链本质是分布式账本、数据库，核心思想是去中心化。

2.服务器端：GO本身的开发者对C与C++进行了优化。数据处理与处理大并发是GO的优势，如美团后台。

3.盛大云cdn，京东消息推送均是使用GO语言。



# 1.2Golang的应用领域

## 1.2.1区块链应用

区块链应用：BT(block chain technology)，分布式账本技术，是一种互联网数据库技术，特点是去中心化、公开透明，每个人均可参与数据库记录。

## 1.2.2后端服务器应用

美团后台流量支撑程序：支撑主站后台流量（排序，推荐，搜索等），提供负载均衡，cache，容错，按条件分流，统计运行指标（qps，latency）等功能。----GO数据处理，大并发，计算是强项。

游戏：游戏服务端（通讯，逻辑，数据存储）

## 1.2.3云计算/云服务后台应用

盛大云CDN（内容分发网络）：CDN的调度系统，分发系统，监控系统，短域名服务，CDN内部开放平台，运营报表系统等。

京东消息推送云服务/京东分布式文件系统：GO实现后台所有服务。



# 1.3Golang的学习方法

## 1.3.1主要研究内容

1）高效而愉快的学习：一个月争取入门基本语法

2）先建立一个整体框架，然后细节

3）在实际工作中，要培养用到什么，能够快速学习什么的能力

4）先know how,再know why：重要的是先做出来，回来再思考其中原理

5）软件编程是一门“做中学”的学科，不是会了再做，而是做了才会！：课后练习

6）适当的囫囵吞枣：指针是难点，继续往下做

7）学习软件编程是在琢磨别人怎么做，而不是我认为应该怎么做的过程：注意作者规定的细节，“我认为”是错的。

![img](file:///C:/Users/xjy/AppData/Local/Temp/msohtmlclip1/01/clip_image006.jpg)



第2章 Golang的概述

# 2.1什么是程序

程序：完成某个功能的指令的集合。

![img](file:///C:/Users/xjy/AppData/Local/Temp/msohtmlclip1/01/clip_image008.jpg)

GO语言是Google公司创造的语言，也是Google主推的语言。

GO语言的发展还是很具有潜力的。

# 2.2GO语言的诞生

Ken Thompson（肯·汤普森）:Unix的原创者，C语言的主要发明者。

2.2.1为什么要创造Go语言

1）计算机硬件更新频繁，目前主流的编程语言发展明显落后于硬件，不能合理利用多核多CPU的优势提升软件系统性能。

2）软件系统复杂度越来越高，维护成本越来越高，目前缺乏一个足够简洁高效的编程语言。【现有编程语言：

1.风格不统一2.计算能力不够3.处理大并发不够好】

3）企业运行维护C/C++的项目，C/C++程序运行速度虽然很快，但是编译速度慢，同时还存在内存泄漏的一系列的困扰需要解决。【脚本语言：python运行慢，编译快。C运行快，编译慢。】

2.2.2GO语言发展简史

Ø 2007年，GO语言原型诞生。

Ø 2009年，全球开源

Ø 2015年，里程碑纪元日期，GO1.5本次更新中移除了“最后残余的C代码”

Ø 2017，GO1.8发布

Ø 2017，GO1.9发布【GO1.9.2】

Ø 2018，GO1.10发布

# 2.3GO语言的特点

**GO****语言保证了既能到达静态编译语言的安全和性能，又达到了动态语言开发维护的高效率，使用一个表达式来形容****GO****语言：****GO=C+python****，说明****GO****语言既有****C****静态语言程序的运行速度，又能达到****Python****动态语言的快速开发。**

**1****）从****C****语言中继承了很多理念，包括表达式语法，控制结构，基础数据类型，调用参数传值，指针等，也保留了和****C****语言一样的编译执行方式及弱化的指针。**

**2****）引入包的概念，用于组织程序结构，****GO****语言的一个文件都要归属一个包，而不能单独存在。**

**3****）垃圾回收机制，内存自动回收，不需要开发人员管理**

**4****）天然并发****(****重要特点****)**

   **a.****从语言层面支持并发，实现简单**

   **b.goroutine,****轻量级线程，可实现大并发处理，高效利用多核**

   **c.****基于****CPS****并发模型（****communicating sequential processes****）实现**

**5****）吸收了管道通信机制，形成****Go****语言特有的管道****channel**

**通过通道****channel****，可以实现不同的****goroute****之间的相互通信。**

**6****）函数返回多个值（示意代码）**

**7****）新的创新：比如切片****slice****、延时执行****defer****等**

 

# 2.4GO语言的开发工具

**工具介绍：**

**1）****visual studio code Microsoft:****支持语言高亮，支持智能提示，多平台（****win,mac,linux****）**

**2）****sublime text:****有购买广告**

**3）****vim:****专门讲****Linux****操作环境时用的**

**4）****emacs:**

**5）****eclipse IDE:**

**6）****LiteIDE:****国人开发**

**7）****JetBrains:ps,ws,pc****等****IDE****工具，安装****GO****插件即可**

**前期少用提示，要深刻理解****Go****语言技术，培养代码感，****->****写代码的感觉**

**利于手撕代码**

# 2.4VS code安装和使用

**1****）下载**

**2****）****Windows****下的安装与使用**

**步骤一：下载后，运行“****.exe****”文件，注意****64****与****32****位的区别，注意安装路径。**

**步骤二：建立专门的代码文件夹**

**3****）****Ubuntu****下的安装与使用**

**4****）****mac****下的安装与使用**

# 2.5Windows下GO语言开发环境搭建

## 2.5.1SDK介绍

**SDK****—****(Software development kit,****软件开发工具包****)****：****sdk****是提供给开发人员使用的，其中包含了对应开发语言的工具包。（源码需要被****SDK****编译）。**

**注意区分****32****位数，与****64****位数。**

## **2.5.2SDK****下载**

![img](file:///C:/Users/xjy/AppData/Local/Temp/msohtmlclip1/01/clip_image010.jpg)

## 2.5.3SDK安装

本机路径：D:\work_software\go1.9.2.windows-amd64\go

![img](file:///C:/Users/xjy/AppData/Local/Temp/msohtmlclip1/01/clip_image012.jpg)

1.使用go.exe可以编译和运行我们的GO源码

2.godoc说明文档

3.gofmt:代码格式化

 

！如何测试go这个SDK安装成功了？

**Windows****下的查看文件指令是：****dir**

**go version**

![img](file:///C:/Users/xjy/AppData/Local/Temp/msohtmlclip1/01/clip_image014.jpg)

**发现这个****go****命令只能在当前安装路径下进行使用，一旦更改路径则失效，提示“不是内部或外部命名。也不是可运行的程序或批处理文件”：当前执行的程序在当前目录下如果不存在，****Windows****系统会在系统中已有的一个名为****path****的环境变量指定的目录中查找。如果还找不到，则提示错误。**

**名词解释：**

**l** **内部命令：运行机制是先在当前路径下寻找****go.exe****文件是否存在，存在则运行。**

**l** **外部命令：内部命令失效时，通过环境变量中的定义，使用****go****命令。**

## **2.5.4SDK****环境变量配置**

配置环境变量介绍：

根据 windows系统在查找可执行程序的原理，可以将Go所在路径定义到环境变量中，让系统帮我们去找运行执行的程序，这样在任何目录下都可以执行go指令。

在Go开发中，需要配置的环境变量如下：

| **环境变量** | **说明**                                                     |
| ------------ | ------------------------------------------------------------ |
| **GOROOT**   | 指定SDK的安装路径  如本机：  D:\work_software\go1.9.2.windows-amd64\go |
| **Path**     | 添加SDK的/bin目录                                            |
| **GOPATH**   | 工作目录，将来go项目的工作路径                               |

配置截图：

步骤一：

![img](file:///C:/Users/xjy/AppData/Local/Temp/msohtmlclip1/01/clip_image016.jpg)

Xxx的用户变量：只对xxx生效

系统变量：对所有用户均生效

步骤二：

1)![img](file:///C:/Users/xjy/AppData/Local/Temp/msohtmlclip1/01/clip_image018.jpg)

2)![img](file:///C:/Users/xjy/AppData/Local/Temp/msohtmlclip1/01/clip_image020.jpg)

![img](file:///C:/Users/xjy/AppData/Local/Temp/msohtmlclip1/01/clip_image022.jpg) ;%GOROOT%\bin

3) ![img](file:///C:/Users/xjy/AppData/Local/Temp/msohtmlclip1/01/clip_image024.jpg)

！注意配置环境变量后点击“确定”才可以。

重启CMD即可。再次测试是否可以全局使用“go version”。

# 2.6Ubuntu搭建go开发环境

# 2.7Mac搭建go开发环境

# 2.8GO语言的快速入门

2.8.1需求

**要求开发一个****hello.go****程序，可以输出“****hello,world****”**

**2.8.2****开发的步骤**

开发这个程序时，电脑上的目录结构怎么有序的处理

**[1]**  **安装****Windows****版本的****VS Code**

**[2]**  **将****Go****代码编写到扩展名为****hello.go****的文件中**

**（目录结构为：****F:\xjy_goproject\src\go_code\chapter3\demo06****）**

![img](file:///C:/Users/xjy/AppData/Local/Temp/msohtmlclip1/01/clip_image026.jpg)

**[3]**   

# 2.8GO语言的快速入门

# 2.8GO语言的快速入门

# 2.8GO语言的快速入门

# 2.8GO语言的快速入门

# 2.8GO语言的快速入门

# 2.8GO语言的快速入门

 

**
**

 

第3章 基于OpenPose模型改进的2D平面多人姿态识别方法

# 3.1引言

出了本章的小结并讨论了未来的工作。

# 3.2改进的OpenPose模型框架

一定的进步。

![img](file:///C:/Users/xjy/AppData/Local/Temp/msohtmlclip1/01/clip_image028.jpg)

（a）部件模型   （b）增加数量的部件模型改进方法 （c）“铰链”实现方式

图1-4 铰链模型以及测试结果

 

## 3.2.1主干网络残差设计

）共18层结构，见图3-2 ResNet18网络结构，图中黑色实线表示为identity残差块，黑色虚线表示为convolutional残差块。

图3-2 ResNet18网络结构

3.2.1.1基于ResNet残差网络改进的OpenPose主干网络

主特征图通道数为128个。

图3-4 改进的主干网络特征图

3.2.1.2 VGG主干网络与ResNet残差主干网络的对比分析

8模型比VGG19耗时大幅衰减，到达1秒的差值。表明训练过程中，残差结构具有速度方面的优势。

## 3.2.2网络裁剪与卷积块堆叠

整体结构，ResNet结构作为主干网络，同时上下分支网络缩减到4个阶段。

的连接。

图3-10 改进的OpenPose上下分支网络结构

3.2.2.1 OpenPose上下分支网络与改进版的对比分析

参数量的对比。

表3-1 不同上下分支网络参数量的对比



|      |      |      |
| ---- | ---- | ---- |
|      |      |      |
|      |      |      |
|      |      |      |
|      |      |      |
|      |      |      |
|      |      |      |
|      |      |      |
|      |      |      |

# 3.3人体关节点的识别

## 3.3.1单人情形

并对每个关节点进行了位置标号的打印。

   

图3-14 输入原始图像与单人关节点检测结果

## 3.3.2多人情形

同样对于所有关节点的高斯响应的极大值会进行阈值筛选。图3-18多人测试结果，关节点位置预测结果展示到输入图像上。

 

图3-18 多人关节点检测结果

# 3.4人体关节点的组合

## 3.4.1单人情形

图3-19 单人情形人体骨架连接，并对各个关节点的位置信息打印在左上角。

 

图3-19 单人情形人体骨架连接

## 3.4.2多人情形

2．计算属于每个个体的关键点集合

最终遍历每个人，并在输入图像上绘制骨架，图3-26关节点组合成多人骨架模型。

​        

图3-26 关节点组合成多人骨架模型



# 3.5网络训练

## 3.5.1数据集准备

取值、超参数的变化趋势、模型的预测精度结果起决定性作用，关于COCO2017数据集标注信息的具体内容在下文的数据预处理阶段有详细说明。

三.机器环境配置

cuDNN版本是7.3.0。

四.数据集预处理：

将新生成的“.json”标注信息生成LMDB训练图像数据，供给caffe训练时使用。

## 3.5.2超参数设置

红色线条L2为关节点连接损失。

 

图3-37 训练过程损失值变化

# 3.6网络训练实验结果与分析

## 3.6.1模型评估

失值对比。

表3-2 不同模型100000迭代损失值对比

|      |      |      |      |      |
| ---- | ---- | ---- | ---- | ---- |
|      |      |      |      |      |
|      |      |      |      |      |
|      |      |      |      |      |

Pose模型。

表3-3 COCO2017 keypoint challenge(test-challenge)

|      |      |      |      |      |      |
| ---- | ---- | ---- | ---- | ---- | ---- |
|      |      |      |      |      |      |
|      |      |      |      |      |      |
|      |      |      |      |      |      |

表3-4 多测时间测试

|      |      |      |      |      |      |      |      |      |      |      |      |
| ---- | ---- | ---- | ---- | ---- | ---- | ---- | ---- | ---- | ---- | ---- | ---- |
|      |      |      |      |      |      |      |      |      |      |      |      |
|      |      |      |      |      |      |      |      |      |      |      |      |

## 3.6.2模型测试

情景。

3.6.2.1图像测试

首先使用改进的OpenPose框架及其训练好的模型，测试图像来源于COCO2017测试集。

3.6.2.2视频测试

，提升幅度约为（4.0-3.5）/3.5*100%≈14.3%。

 

# 3.7本章小结

本章。

 



 

第4章 人体动作识别应用

# 4.1引言

工作。

# 4.2人体动作识别相关研究

方法。

# 4.3应用OpenPose模型的人体动作识别实现

判别。

# 4.4本章小结

识别。

 



 



第5章 总结与展望

# 5.1工作总结

问题。

# 5.2研究展望

适用性。



 

参考文献

[1].  Kandlhofer M, Steinbauer G, Hirschmuglgaisch S, et al. Artificial intelligence and computer science in education: From kindergarten to university[C]. frontiers in education conference, 2016: 1-9.

[2].  Tang X， Wang B， Rong Y， et al. Artificial intelligence will reduce the need for clinical medical physicists[J]. Journal of Applied Clinical Medical Physics， 2018， 19(1): 6-9.



 

致谢

时光

 

 



 

附录

附录