# XOR加密解密与破解

# 一、概述

基于异或位运算对数据进行加密或者解密，一般用在流式密码加密中。

# 二、安装

```bash
go get -u github.com/cryptography-research-lab/go-xor
```

# 三、代码示例

## 3.1 XOR加密和解密

```

```



## 3.2 XOR秘钥重用攻击

关于XOR秘钥重用攻击的介绍请参见[五、XOR秘钥重用攻击](#五、XOR秘钥重用攻击)部分，此处只是代码使用的API。

### 3.2.1 已知一对明文密文，解密任意其它密文

```go
```

## 3.2.2 已知多对密文，猜测秘钥 



# 四、XOR加密解密介绍

TODO 



# 五、XOR秘钥重用攻击

## 5.1 XOR秘钥重用攻击介绍

当XOR使用同一个秘钥多次进行加密的时候存在几种被攻击利用的方式： 

- 能够进行加密，拿到已知的明文和密文，就可以对其他任意密文进行解密

- 能够拿到一批密文，通过这批密文推测出明文，不需要已知任何明文 

## 5.2 XOR秘钥重用攻击：已知明文密文，解密任意密文

假设A、B为明文，cipherA和cipherB尾加密后的密文，K为秘钥

A的密文cipherA：

```text
cipherA = XOR(A, K) = A XOR K 
```

B的密文cipherB：

```text
cipherB = XOR(B, K) = B XOR K
```

则XOR(cipherA, cipherB)：

```text
XOR(cipherA, cipherB) = (A ^ K) ^ (B ^ K) = A ^ B ^ (K ^ K) = A ^ B ;
```

```text
E(A) XOR E(B) XOR B = A XOR B XOR B = A;
```

# 六、TODO 



























