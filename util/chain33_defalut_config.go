package util

// Chain33DefaultConfig Chain33 Default Config
var Chain33DefaultConfig = `
TestNet=false
CoinSymbol="bty"
EnableParaFork=true

[blockchain]
defCacheSize=128
maxFetchBlockNum=128
timeoutSeconds=5
batchBlockNum=128
driver="leveldb"
dbPath="paradatadir"
dbCache=64
isStrongConsistency=true
singleMode=true
batchsync=false
#平行链钱包通过平行链区块seqence索引高度，缺省是true
isRecordBlockSequence=true
isParaChain = true
enableTxQuickIndex=true
# 升级storedb是否重新执行localdb，bityuan主链升级不需要开启，平行链升级需要开启
enableReExecLocal=true
# 使能精简localdb
enableReduceLocaldb=false
enablePushSubscribe=false

[p2p]
types=["dht"]
enable=false
driver="leveldb"
dbPath="paradatadir/addrbook"
dbCache=4
grpcLogFile="grpc33.log"
waitPid=true

[p2p.sub.dht]
DHTDataPath="paradatadir/p2pstore"


[mempool]
name="para"
poolCacheSize=10240
#联盟链没有交易费，对应平行链minTxFeeRate需要设为0
minTxFeeRate=100000
maxTxNumPerAccount=10000

[consensus]
name="para"
genesisBlockTime=1514533390
genesis="14KEKbYtKKQm4wMthSK9J4La4nAiidGozt"
minerExecs=["paracross"]    #配置挖矿合约

[mver.consensus]
fundKeyAddr = "1BQXS6TxaYYG5mADaWij4AxhZZUTpw95a5"
powLimitBits = "0x1f00ffff"
maxTxNumber = 1600

[mver.consensus.paracross]
#超级节点挖矿奖励
coinReward=18
#发展基金奖励
coinDevFund=12
#如果超级节点上绑定了委托账户，则奖励超级节点coinBaseReward，其余部分(coinReward-coinBaseReward)按权重分给委托账户
coinBaseReward=3
#委托账户最少解绑定时间(按小时)
unBindTime=24
#支持挖矿奖励的1e8小数模式，比如18coin 需要配置成1800000000 以支持小数位后的配置,如果true，意味着已经打开即须配置coinReward=1800000000
decimalMode=false
#挖矿模式， normal：缺省挖矿，其他自定义，注册名字需要和配置名字保持一致
minerMode="normal"
#挖矿减半周期,按高度减半
halvePeriod=1000


[consensus.sub.para]
#主链节点的grpc服务器ip，当前可以支持多ip负载均衡，如“118.31.177.1:8802,39.97.2.127:8802”
#ParaRemoteGrpcClient="jiedian2.bityuan.com,cloud.bityuan.com"
ParaRemoteGrpcClient="localhost:8802"
#主链指定高度的区块开始同步
startHeight=345850
#主链指定高度后等待块数，防止主链回滚，联盟链最小为1，小于1则采用缺省高度100
#waitMainBlockNum=100
#等待打包主链区块时间间隔，单位毫秒
writeBlockMsec=2000
#共识节点账户，共识节点需要配置自己的账户，并且钱包导入对应种子，非共识节点留空
authAccount=""
#创世地址额度
genesisAmount=100000000
#主链支持平行链共识tx分叉高度，需要和主链保持严格一致,不可修改,2270000是bityuan主链对应高度， ycc或其他按实际修改
#不可为0，主链Local时候需特殊配置
mainForkParacrossCommitTx=2270000
#主链开启循环检查共识交易done的fork高度,需要和主链保持严格一致,不可修改,4320000是bityuan主链对应高度， ycc或其他按实际修改
#不可为0，主链Local时候需特殊配置
mainLoopCheckCommitTxDoneForkHeight=4320000
#无平行链交易的主链区块间隔，平行链产生一个空块，从高度0开始，配置[blockHeight:interval],比如["0:50","1000:100"]
emptyBlockInterval=["0:50"]

#平行链共识节点聚合签名配置
[consensus.sub.para.bls]
#是否开启聚合签名，缺省不开启
blsSign=false


[store]
name="kvmvccmavl"
driver="leveldb"
storedbVersion="2.0.0"
dbPath="paradatadir/mavltree"
dbCache=128



[wallet]
minFee=100000
driver="leveldb"
dbPath="parawallet"
dbCache=16
signType="secp256k1"
minerdisable=true

[exec]
enableStat=false
enableMVCC=false

[exec.sub.relay]
genesis="12qyocayNF7Lv6C9qW4avxs2E7U41fKSfv"

[exec.sub.manage]
superManager=["12qyocayNF7Lv6C9qW4avxs2E7U41fKSfv"]
#autonomy执行器名字,空则不开启,使用superManager list
autonomyExec=""

[exec.sub.token]
saveTokenTxList=true
tokenApprs=[]

[exec.sub.paracross]
#平行链自共识停止n个空块的对应主链高度后，超级账户可以直接参与投票,这个高度只在主链有效
paraConsensusStopBlocks=30000

[exec.sub.autonomy]
total="16htvcBNSEA7fZhAdLJphDwQRQJaHpyHTp"
useBalance=false

[exec.sub.evm]
#平行链evm合约ETH资产映射的合约和资产类型（symbol）
ethMapFromExecutor="paracross"
ethMapFromSymbol="coins.bty"

[exec.sub.mix]
#私对私的交易费,交易比较大，需要多的手续费
txFee=0
#私对私token转账，花费token(true)还是BTY(false),
tokenFee=false
#curve H point
pointHX="19172955941344617222923168298456110557655645809646772800021167670156933290312"
pointHY="21116962883761739586121793871108889864627195706475546685847911817475098399811"
#电路最大支持1024个叶子hash，10 level， 配置可以小于1024,但不能大于
maxTreeLeaves=1024
#管理员列表
mixApprs=[]

[exec.sub.zksync]
manager=[
    "14KEKbYtKKQm4wMthSK9J4La4nAiidGozt",
    "12qyocayNF7Lv6C9qW4avxs2E7U41fKSfv"
]

#系统中所有的fork,默认用chain33的测试网络的
#但是我们可以替换
[fork.system]
ForkChainParamV1= 0
ForkCheckTxDup=0
ForkBlockHash= 1
ForkMinerTime= 0
ForkTransferExec=0
ForkExecKey=0
ForkTxGroup=0
ForkResetTx0=0
ForkWithdraw=0
ForkExecRollback=0
ForkCheckBlockTime=0
ForkTxHeight=0
ForkTxGroupPara=0
ForkChainParamV2=0
ForkMultiSignAddress=0
ForkStateDBSet=0
ForkLocalDBAccess=0
ForkBlockCheck=0
ForkBase58AddressCheck=0
#平行链上使能平行链执行器如user.p.x.coins执行器的注册，缺省为0，对已有的平行链需要设置一个fork高度
ForkEnableParaRegExec=0
ForkCacheDriver=0
ForkTicketFundAddrV1=-1 #fork6.3
#主链和平行链都使用同一个fork高度
ForkRootHash=7200000
ForkFormatAddressKey=0

[fork.sub.coins]
Enable=0

[fork.sub.coinsx]
Enable=0

[fork.sub.pos33]
Enable=0
ForkReward15=725000 
ForkFixReward=5000000
UseEntrust=0

[fork.sub.ticket]
Enable=0
ForkTicketId =0
ForkTicketVrf =0

[fork.sub.retrieve]
Enable=0
ForkRetrive=0
ForkRetriveAsset=0

[fork.sub.hashlock]
Enable=0
ForkBadRepeatSecret=0

[fork.sub.manage]
Enable=0
ForkManageExec=0
#manage增加配置需要经过autonomy board成员审批，平行链不开启
ForkManageAutonomyEnable=-1

[fork.sub.token]
Enable=0
ForkTokenBlackList= 0
ForkBadTokenSymbol= 0
ForkTokenPrice=0
ForkTokenSymbolWithNumber=0
ForkTokenCheck= 0

[fork.sub.trade]
Enable=0
ForkTradeBuyLimit= 0
ForkTradeAsset= 0
ForkTradeID = 0
ForkTradeFixAssetDB = 0
ForkTradePrice = 0

[fork.sub.paracross]
Enable=0
ForkParacrossWithdrawFromParachain=0
ForkParacrossCommitTx=0
ForkLoopCheckCommitTxDone=0
#仅平行链适用，自共识分阶段开启，缺省是0，若对应主链高度7200000之前开启过自共识，需要重新配置此分叉，并为之前自共识设置selfConsensEnablePreContract配置项
ForkParaSelfConsStages=0
ForkParaAssetTransferRbk=0
ForkParaSupervision=0
#仅平行链适用，开启挖矿交易的高度，已有代码版本可能未在0高度开启挖矿，需要设置这个高度，新版本默认从0开启挖矿，通过交易配置分阶段奖励
ForkParaFullMinerHeight=0
#仅平行链适用，在旧的版本中计算blockTxHash输入高度为0，需要在此高度后统一采用新的主链高度值，旧的版本需要设置此分叉高度，新版本缺省为0即可
ForkParaRootHash=0
#nodegroup approve需要经过autonomy board成员审批,平行链不开启
ForkParaAutonomySuperGroup=-1
ForkParaFreeRegister=0

[fork.sub.evm]
Enable=0
ForkEVMState=0
ForkEVMABI=0
ForkEVMFrozen=0
ForkEVMKVHash=0
ForkEVMYoloV1=0
ForkEVMTxGroup=0

[fork.sub.blackwhite]
Enable=0
ForkBlackWhiteV2=0

[fork.sub.cert]
Enable=0

[fork.sub.guess]
Enable=0

[fork.sub.lottery]
Enable=0

[fork.sub.oracle]
Enable=0

[fork.sub.relay]
Enable=0

[fork.sub.norm]
Enable=0

[fork.sub.pokerbull]
Enable=0

[fork.sub.privacy]
Enable=0

[fork.sub.game]
Enable=0

[fork.sub.vote]
Enable=0

[fork.sub.accountmanager]
Enable=0

[fork.sub.wasm]
Enable=0

[fork.sub.valnode]
Enable=0
[fork.sub.dpos]
Enable=0
[fork.sub.echo]
Enable=0
[fork.sub.storage]
Enable=0
ForkStorageLocalDB=0

[fork.sub.qbftNode]
Enable=0

[fork.sub.multisig]
Enable=0

[fork.sub.mix]
Enable=0

[fork.sub.unfreeze]
Enable=0
ForkTerminatePart=0
ForkUnfreezeIDX= 0

[fork.sub.autonomy]
Enable=0
ForkAutonomyDelRule=0
ForkAutonomyEnableItem=0

[fork.sub.jsvm]
Enable=0

[fork.sub.evmxgo]
Enable=0

[fork.sub.issuance]
Enable=0
ForkIssuanceTableUpdate=0
ForkIssuancePrecision=0

[fork.sub.collateralize]
Enable=0
ForkCollateralizeTableUpdate=0
ForkCollateralizePrecision=0

#对已有的平行链如果不是从0开始同步数据，需要设置这个kvmvccmavl的对应平行链高度的fork，如果从0开始同步，statehash会跟以前mavl的不同
[fork.sub.store-kvmvccmavl]
ForkKvmvccmavl=0

[pprof]
listenAddr = "localhost:6061"

#exchange合约相关配置
[mver.exec.sub.exchange]
#银行帐户列表（现第一个地址用来收取手续费）
banks = ["1PTGVR7TUm1MJUH7M1UNcKBGMvfJ7nCrnN"]#Fee
#币种配置，
#rate每笔手续费率,配置时需*1e8(如：收取每笔交易千分之一的手续费，rate=100000)
#minFee最小手续费,配置时需*1e8(如：最小手续费收取1个，minFee=100000000)
coins = [
    { name = "bty", rate = 100000, minFee = 0 },
    { name = "coins.bty", rate = 100000, minFee = 0 },
    { name = "ETH", rate = 100000, minFee = 0 },
    { name = "USDT", rate = 100000, minFee = 0 },
]

[fork.sub.exchange]
Enable=0
ForkFix1=0
ForkParamV1 = 0
ForkParamV2 = 0
ForkParamV3 = 0
ForkParamV4 = 0
ForkParamV5 = 0
ForkParamV6 = 0
ForkParamV7 = 0
ForkParamV8 = 0
ForkParamV9 = 0

[fork.sub.zksync]
Enable=0
`

// Chain33TitleM ...
var Chain33TitleM = `
Title="bityuan"
`

// Chain33TitleP ...
var Chain33TitleP = `
Title="user.p.xx."
`

// Chain33TitleFmt ...
var Chain33TitleFmt = `Title="%s"
`

// Chain33CoinSymbol ...
var Chain33CoinSymbol = `CoinSymbol="%s"
`

// Chain33V165 chain33 1.65的配置
// 支持不指定配置文件, 可以更方便的使用
var Chain33V165 = `

FixTime=false

[log]
# 日志级别，支持debug(dbug)/info/warn/error(eror)/crit
loglevel = "debug"
logConsoleLevel = "info"
# 日志文件名，可带目录，所有生成的日志文件都放到此目录下
logFile = "logs/chain33.log"
# 单个日志文件的最大值（单位：兆）
maxFileSize = 300
# 最多保存的历史日志文件个数
maxBackups = 100
# 最多保存的历史日志消息（单位：天）
maxAge = 28
# 日志文件名是否使用本地事件（否则使用UTC时间）
localTime = true
# 历史日志文件是否压缩（压缩格式为gz）
compress = true
# 是否打印调用源文件和行号
callerFile = false
# 是否打印调用方法
callerFunction = false

[blockchain]
dbPath="datadir"
dbCache=64
batchsync=false
isRecordBlockSequence=true
enableTxQuickIndex=true
# 升级storedb是否重新执行localdb，bityuan主链升级不需要开启，平行链升级需要开启
enableReExecLocal=false
# 使能精简localdb
enableReduceLocaldb=false
enablePushSubscribe=true

[p2p]
types=["gossip", "dht"]
dbPath="datadir/addrbook"
dbCache=4
grpcLogFile="grpc33.log"
#waitPid 等待seed导入
waitPid=false

[p2p.sub.gossip]
port=13802
seeds=[]
isSeed=false
innerSeedEnable=true
useGithub=true
innerBounds=300

[p2p.sub.dht]
port=13803


[rpc]
jrpcBindAddr="localhost:18801"
grpcBindAddr="localhost:18802"
whitelist=["127.0.0.1"]
jrpcFuncWhitelist=["*"]
grpcFuncWhitelist=["*"]
enableTLS=false
certFile="cert.pem"
keyFile="key.pem"

[mempool]
maxTxNumPerAccount=100

[store]
dbPath="datadir/mavltree"
dbCache=128

[store.sub.mavl]
enableMavlPrefix=true
enableMVCC=false
enableMavlPrune=true
pruneHeight=10000
enableMemTree=true
enableMemVal=true
# 缓存close ticket数目，该缓存越大同步速度越快，最大设置到1500000,默认200000
tkCloseCacheLen=200000

[store.sub.kvmvccmavl]
enableMVCCIter=true
enableMavlPrefix=false
enableMVCC=false
enableMavlPrune=false
pruneMavlHeight=10000
enableMVCCPrune=false
pruneMVCCHeight=10000
# 是否使能mavl数据载入内存
enableMemTree=true
# 是否使能mavl叶子节点数据载入内存
enableMemVal=true
# 缓存close ticket数目，该缓存越大同步速度越快，最大设置到1500000
tkCloseCacheLen=100000
# 该参数针对平行链，如果平行链的ForkKvmvccmavl高度不为0,需要开启此功能,开启此功能需要从0开始执行区块
enableEmptyBlockHandle=false

[wallet]
dbPath="wallet"
dbCache=16

[wallet.sub.ticket]
minerdisable=false
minerwhitelist=["*"]

[exec]
enableStat=false
enableMVCC=false

[exec.sub.token]
saveTokenTxList=false

[metrics]
#是否使能发送metrics数据的发送
enableMetrics=false
#数据保存模式
dataEmitMode="influxdb"

[metrics.sub.influxdb]
#以纳秒为单位的发送间隔
duration=1000000000
url="http://influxdb:8086"
database="chain33metrics"
username=""
password=""
namespace=""


`
