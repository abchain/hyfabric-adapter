## 关于demo

### 1.0 fabric 1.4 网络启动

#### 1.1 通过项目 [deploy-fabric](https://dev.stringon.com/yinminghao/deployfabirc) 启动网络

#### 1.2 使用 fabric-sample 启动网络

1. 获取 fabric-sample 项目release-1.4版本
    ```shell
    $ git clone https://github.com/hyperledger/fabric-samples.git
    $ git checkout  v1.4.7
    ```
2.  启动 fabric 1.4 网络
    ```shell
    $ cd first-network 
    ## 清理之前环境
    $ ./byfn.sh -m down 
    ## 生成启动所需文件
    $ ./byfn.sh generate
    ## 启动网络, 1.4版本，包含ca，可通过调整参数定制网络
    $ ./byfn.sh -m up -i 1.4.7 -a
    ```

### 2.0 修改 demo/config.yaml

根据网络配置修改相关参数（没有修改默认网络参数可跳过此步）

### 3.0 运行 demo

执行 demo.go, 打印出连接成功与查询configblock成功日志   
```shell
$ go run demo/demo.go
```

### 4.0 可能遇到的错误(持续补充)

1. **Endorser Client Status Code...connection is in TRANSIENT_FAILURE** 连接节点出现问题可讲节点host添加进/etc/hosts文件中
    ```
    ### /etc/hosts
    127.0.0.1 peer0.org1.example.com
    127.0.0.1 peer1.org1.example.com
    ```
2. **Key with SKI ... not found in /tmp/state-store/msp/keystore**, 私钥文件未找到
    ```shell
    $ cp $GOPATH/src/github.com/hyperledger/fabric-samples/first-network/crypto-config/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp/keystore/priv_sk /tmp/state-store/msp/keystore
    ```
3. **user not found**，考虑路径问题，可将config.yaml中的${GOPATH}换为绝对路径
4. **x509: cannot validate certificate for 127.0.0.1 because it doesn't contain any IP SANs**, ca启动的是根据hostname而不是ip，可通过修改/etc/hosts解决
    ```
    ### /etc/hosts
    127.0.0.1 ca.org1.example.com
    ```
5. **access denied: channel [mychannel] creator org [Org1MSP]** /tmp/state-store/msp/keystore 中sk过多，删掉即可
6. **failed constructing descriptor for chaincodes:<name:"mycc" >** 链码安装失败，可尝试重新安装、实例化