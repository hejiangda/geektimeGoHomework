# 第9周作业

1. 总结几种 socket 粘包的解包方式：fix length/delimiter based/length field based frame decoder。尝试举例其应用。

答：
   1. fix length 固定缓冲区大小
    发送方和接收方规定固定大小的缓冲区，当发送的数据长度不够时用空字符弥补，若发送的数据小于缓冲区长度时会有较大浪费，适合发送格式固定，长度固定的数据，如传感器的监测数据
   2. delimiter based 分隔符分隔
    在发送的数据结尾用特殊字符分隔。如用'\n'分隔，接收方每次接收1行数据，按行分隔，适合传输的数据中不含分隔字符的场景
   3. length field based frame decoder 数据头+数据正文
    在数据头中存储数据正文的大小，当读取的数据小于数据头中的大小时，继续读取数据，直到读取的数据长度等于数据头中的长度时才停止。适合知道要发送的数据长度的情况。
   
2. 实现一个从 socket connection 中解码出 goim 协议的解码器。
   1. 
