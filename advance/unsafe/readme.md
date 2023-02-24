

# unsafe



对象内存布局
    
    内存对齐 + 填充


操作内存，本质上是对象的起始地址

起始地址 + 偏移量




对象如果挪动了，uintptr 那就瞎了


unsafe.Pointer - entityAddr起始地址

uintptr - offset 偏移量


性能瓶颈 
    
    unsafe 取代 reflect，性能显著提升。


    

读写 加速