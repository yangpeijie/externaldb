// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: exdbconfig.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Chain33 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title             string  `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	GrpcHost          string  `protobuf:"bytes,2,opt,name=grpcHost,proto3" json:"grpcHost,omitempty"`
	Host              string  `protobuf:"bytes,3,opt,name=host,proto3" json:"host,omitempty"`
	Symbol            string  `protobuf:"bytes,4,opt,name=symbol,proto3" json:"symbol,omitempty"`
	OtherChainGenesis int64   `protobuf:"varint,5,opt,name=otherChainGenesis,proto3" json:"otherChainGenesis,omitempty"`
	PerBlockCoin      int64   `protobuf:"varint,6,opt,name=perBlockCoin,proto3" json:"perBlockCoin,omitempty"`
	RollbackSeq       []int64 `protobuf:"varint,7,rep,packed,name=rollbackSeq,proto3" json:"rollbackSeq,omitempty"`
	CoinPrecision     float64 `protobuf:"fixed64,8,opt,name=coinPrecision,proto3" json:"coinPrecision,omitempty"`
	PerTicketReward   float64 `protobuf:"fixed64,9,opt,name=perTicketReward,proto3" json:"perTicketReward,omitempty"`
}

func (x *Chain33) Reset() {
	*x = Chain33{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exdbconfig_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chain33) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chain33) ProtoMessage() {}

func (x *Chain33) ProtoReflect() protoreflect.Message {
	mi := &file_exdbconfig_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chain33.ProtoReflect.Descriptor instead.
func (*Chain33) Descriptor() ([]byte, []int) {
	return file_exdbconfig_proto_rawDescGZIP(), []int{0}
}

func (x *Chain33) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Chain33) GetGrpcHost() string {
	if x != nil {
		return x.GrpcHost
	}
	return ""
}

func (x *Chain33) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *Chain33) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *Chain33) GetOtherChainGenesis() int64 {
	if x != nil {
		return x.OtherChainGenesis
	}
	return 0
}

func (x *Chain33) GetPerBlockCoin() int64 {
	if x != nil {
		return x.PerBlockCoin
	}
	return 0
}

func (x *Chain33) GetRollbackSeq() []int64 {
	if x != nil {
		return x.RollbackSeq
	}
	return nil
}

func (x *Chain33) GetCoinPrecision() float64 {
	if x != nil {
		return x.CoinPrecision
	}
	return 0
}

func (x *Chain33) GetPerTicketReward() float64 {
	if x != nil {
		return x.PerTicketReward
	}
	return 0
}

type Kafka struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host  string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Group string `protobuf:"bytes,2,opt,name=group,proto3" json:"group,omitempty"`
	Topic string `protobuf:"bytes,3,opt,name=topic,proto3" json:"topic,omitempty"`
}

func (x *Kafka) Reset() {
	*x = Kafka{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exdbconfig_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Kafka) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Kafka) ProtoMessage() {}

func (x *Kafka) ProtoReflect() protoreflect.Message {
	mi := &file_exdbconfig_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Kafka.ProtoReflect.Descriptor instead.
func (*Kafka) Descriptor() ([]byte, []int) {
	return file_exdbconfig_proto_rawDescGZIP(), []int{1}
}

func (x *Kafka) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *Kafka) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *Kafka) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

type ESDB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host   string `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	Prefix string `protobuf:"bytes,3,opt,name=prefix,proto3" json:"prefix,omitempty"`
	Bulk   bool   `protobuf:"varint,4,opt,name=bulk,proto3" json:"bulk,omitempty"`
	User   string `protobuf:"bytes,5,opt,name=user,proto3" json:"user,omitempty"`
	Pwd    string `protobuf:"bytes,6,opt,name=pwd,proto3" json:"pwd,omitempty"`
}

func (x *ESDB) Reset() {
	*x = ESDB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exdbconfig_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ESDB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ESDB) ProtoMessage() {}

func (x *ESDB) ProtoReflect() protoreflect.Message {
	mi := &file_exdbconfig_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ESDB.ProtoReflect.Descriptor instead.
func (*ESDB) Descriptor() ([]byte, []int) {
	return file_exdbconfig_proto_rawDescGZIP(), []int{2}
}

func (x *ESDB) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *ESDB) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

func (x *ESDB) GetBulk() bool {
	if x != nil {
		return x.Bulk
	}
	return false
}

func (x *ESDB) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *ESDB) GetPwd() string {
	if x != nil {
		return x.Pwd
	}
	return ""
}

type ESIndex struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NumberOfShards   int32 `protobuf:"varint,1,opt,name=numberOfShards,proto3" json:"numberOfShards,omitempty"`
	NumberOfReplicas int32 `protobuf:"varint,2,opt,name=numberOfReplicas,proto3" json:"numberOfReplicas,omitempty"`
}

func (x *ESIndex) Reset() {
	*x = ESIndex{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exdbconfig_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ESIndex) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ESIndex) ProtoMessage() {}

func (x *ESIndex) ProtoReflect() protoreflect.Message {
	mi := &file_exdbconfig_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ESIndex.ProtoReflect.Descriptor instead.
func (*ESIndex) Descriptor() ([]byte, []int) {
	return file_exdbconfig_proto_rawDescGZIP(), []int{3}
}

func (x *ESIndex) GetNumberOfShards() int32 {
	if x != nil {
		return x.NumberOfShards
	}
	return 0
}

func (x *ESIndex) GetNumberOfReplicas() int32 {
	if x != nil {
		return x.NumberOfReplicas
	}
	return 0
}

type RPC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host          string   `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	WhiteList     []string `protobuf:"bytes,2,rep,name=whiteList,proto3" json:"whiteList,omitempty"`
	SwaggerHost   string   `protobuf:"bytes,3,opt,name=swaggerHost,proto3" json:"swaggerHost,omitempty"`
	Name          string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	TriggerCount  float64  `protobuf:"fixed64,5,opt,name=triggerCount,proto3" json:"triggerCount,omitempty"`
	EnableSwagger bool     `protobuf:"varint,6,opt,name=enableSwagger,proto3" json:"enableSwagger,omitempty"`
	JrpcHost      string   `protobuf:"bytes,7,opt,name=jrpcHost,proto3" json:"jrpcHost,omitempty"`
	JrpcName      string   `protobuf:"bytes,8,opt,name=jrpcName,proto3" json:"jrpcName,omitempty"`
}

func (x *RPC) Reset() {
	*x = RPC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exdbconfig_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RPC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RPC) ProtoMessage() {}

func (x *RPC) ProtoReflect() protoreflect.Message {
	mi := &file_exdbconfig_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RPC.ProtoReflect.Descriptor instead.
func (*RPC) Descriptor() ([]byte, []int) {
	return file_exdbconfig_proto_rawDescGZIP(), []int{4}
}

func (x *RPC) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *RPC) GetWhiteList() []string {
	if x != nil {
		return x.WhiteList
	}
	return nil
}

func (x *RPC) GetSwaggerHost() string {
	if x != nil {
		return x.SwaggerHost
	}
	return ""
}

func (x *RPC) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RPC) GetTriggerCount() float64 {
	if x != nil {
		return x.TriggerCount
	}
	return 0
}

func (x *RPC) GetEnableSwagger() bool {
	if x != nil {
		return x.EnableSwagger
	}
	return false
}

func (x *RPC) GetJrpcHost() string {
	if x != nil {
		return x.JrpcHost
	}
	return ""
}

func (x *RPC) GetJrpcName() string {
	if x != nil {
		return x.JrpcName
	}
	return ""
}

type Sync struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,9,opt,name=name,proto3" json:"name,omitempty"`
	// pushBind 和 host 不一样， 在于云上， 外网ip 不可以bind
	PushHost string `protobuf:"bytes,1,opt,name=pushHost,proto3" json:"pushHost,omitempty"` // 外网ip
	PushName string `protobuf:"bytes,2,opt,name=pushName,proto3" json:"pushName,omitempty"`
	PushBind string `protobuf:"bytes,3,opt,name=pushBind,proto3" json:"pushBind,omitempty"` // 本地
	// 同步开始位置, 如果收到之前的, 忽略. (老版本的chain33, 会从0开始同步,
	// 不能指定开始)
	StartSeq       int64  `protobuf:"varint,4,opt,name=startSeq,proto3" json:"startSeq,omitempty"`
	StartHeight    int64  `protobuf:"varint,5,opt,name=startHeight,proto3" json:"startHeight,omitempty"`
	StartBlockHash string `protobuf:"bytes,6,opt,name=startBlockHash,proto3" json:"startBlockHash,omitempty"`
	// 指定推送格式,
	PushFormat string `protobuf:"bytes,11,opt,name=pushFormat,proto3" json:"pushFormat,omitempty"`
}

func (x *Sync) Reset() {
	*x = Sync{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exdbconfig_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sync) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sync) ProtoMessage() {}

func (x *Sync) ProtoReflect() protoreflect.Message {
	mi := &file_exdbconfig_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sync.ProtoReflect.Descriptor instead.
func (*Sync) Descriptor() ([]byte, []int) {
	return file_exdbconfig_proto_rawDescGZIP(), []int{5}
}

func (x *Sync) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Sync) GetPushHost() string {
	if x != nil {
		return x.PushHost
	}
	return ""
}

func (x *Sync) GetPushName() string {
	if x != nil {
		return x.PushName
	}
	return ""
}

func (x *Sync) GetPushBind() string {
	if x != nil {
		return x.PushBind
	}
	return ""
}

func (x *Sync) GetStartSeq() int64 {
	if x != nil {
		return x.StartSeq
	}
	return 0
}

func (x *Sync) GetStartHeight() int64 {
	if x != nil {
		return x.StartHeight
	}
	return 0
}

func (x *Sync) GetStartBlockHash() string {
	if x != nil {
		return x.StartBlockHash
	}
	return ""
}

func (x *Sync) GetPushFormat() string {
	if x != nil {
		return x.PushFormat
	}
	return ""
}

type Convert struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppName       string        `protobuf:"bytes,1,opt,name=appName,proto3" json:"appName,omitempty"`
	ExecAddresses []string      `protobuf:"bytes,5,rep,name=execAddresses,proto3" json:"execAddresses,omitempty"`
	StartSeq      int64         `protobuf:"varint,9,opt,name=startSeq,proto3" json:"startSeq,omitempty"`
	Data          []*ExecConfig `protobuf:"bytes,10,rep,name=data,proto3" json:"data,omitempty"`
	Stat          []*StatConfig `protobuf:"bytes,11,rep,name=stat,proto3" json:"stat,omitempty"`
	// 默认的处理插件
	DefaultExec string `protobuf:"bytes,12,opt,name=defaultExec,proto3" json:"defaultExec,omitempty"`
	// 是否处理非本链数据
	DealOtherChain bool `protobuf:"varint,13,opt,name=dealOtherChain,proto3" json:"dealOtherChain,omitempty"`
	// 是否开启权限控制
	OpenAccessControl bool   `protobuf:"varint,14,opt,name=openAccessControl,proto3" json:"openAccessControl,omitempty"`
	SaveBlockInfo     bool   `protobuf:"varint,15,opt,name=saveBlockInfo,proto3" json:"saveBlockInfo,omitempty"`
	AddressDriver     string `protobuf:"bytes,16,opt,name=addressDriver,proto3" json:"addressDriver,omitempty"`
}

func (x *Convert) Reset() {
	*x = Convert{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exdbconfig_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Convert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Convert) ProtoMessage() {}

func (x *Convert) ProtoReflect() protoreflect.Message {
	mi := &file_exdbconfig_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Convert.ProtoReflect.Descriptor instead.
func (*Convert) Descriptor() ([]byte, []int) {
	return file_exdbconfig_proto_rawDescGZIP(), []int{6}
}

func (x *Convert) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

func (x *Convert) GetExecAddresses() []string {
	if x != nil {
		return x.ExecAddresses
	}
	return nil
}

func (x *Convert) GetStartSeq() int64 {
	if x != nil {
		return x.StartSeq
	}
	return 0
}

func (x *Convert) GetData() []*ExecConfig {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Convert) GetStat() []*StatConfig {
	if x != nil {
		return x.Stat
	}
	return nil
}

func (x *Convert) GetDefaultExec() string {
	if x != nil {
		return x.DefaultExec
	}
	return ""
}

func (x *Convert) GetDealOtherChain() bool {
	if x != nil {
		return x.DealOtherChain
	}
	return false
}

func (x *Convert) GetOpenAccessControl() bool {
	if x != nil {
		return x.OpenAccessControl
	}
	return false
}

func (x *Convert) GetSaveBlockInfo() bool {
	if x != nil {
		return x.SaveBlockInfo
	}
	return false
}

func (x *Convert) GetAddressDriver() string {
	if x != nil {
		return x.AddressDriver
	}
	return ""
}

type ConfigNew struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 处理的是哪条链 bityuan 或 user.p.xxxx.
	// Chain 依赖的chain33节点jsonrpc地址
	Chain *Chain33 `protobuf:"bytes,1,opt,name=chain,proto3" json:"chain,omitempty"`
	// 同步和展开的数据库
	Dbtype         string      `protobuf:"bytes,2,opt,name=dbtype,proto3" json:"dbtype,omitempty"`
	EsVersion      int32       `protobuf:"varint,10,opt,name=es_version,json=esVersion,proto3" json:"es_version,omitempty"`
	SyncEs         *ESDB       `protobuf:"bytes,3,opt,name=sync_es,json=syncEs,proto3" json:"sync_es,omitempty"`
	ConvertEs      *ESDB       `protobuf:"bytes,4,opt,name=convert_es,json=convertEs,proto3" json:"convert_es,omitempty"`
	Sync           *Sync       `protobuf:"bytes,5,opt,name=sync,proto3" json:"sync,omitempty"`
	EsIndex        *ESIndex    `protobuf:"bytes,6,opt,name=es_index,json=esIndex,proto3" json:"es_index,omitempty"`
	Rpc            *RPC        `protobuf:"bytes,7,opt,name=rpc,proto3" json:"rpc,omitempty"`
	Convert        *Convert    `protobuf:"bytes,8,opt,name=convert,proto3" json:"convert,omitempty"`
	Kafka          *Kafka      `protobuf:"bytes,9,opt,name=kafka,proto3" json:"kafka,omitempty"`
	ManagerAddress []string    `protobuf:"bytes,11,rep,name=manager_address,json=managerAddress,proto3" json:"manager_address,omitempty"`
	Contracts      []*Contract `protobuf:"bytes,12,rep,name=Contracts,proto3" json:"Contracts,omitempty"`
}

func (x *ConfigNew) Reset() {
	*x = ConfigNew{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exdbconfig_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigNew) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigNew) ProtoMessage() {}

func (x *ConfigNew) ProtoReflect() protoreflect.Message {
	mi := &file_exdbconfig_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigNew.ProtoReflect.Descriptor instead.
func (*ConfigNew) Descriptor() ([]byte, []int) {
	return file_exdbconfig_proto_rawDescGZIP(), []int{7}
}

func (x *ConfigNew) GetChain() *Chain33 {
	if x != nil {
		return x.Chain
	}
	return nil
}

func (x *ConfigNew) GetDbtype() string {
	if x != nil {
		return x.Dbtype
	}
	return ""
}

func (x *ConfigNew) GetEsVersion() int32 {
	if x != nil {
		return x.EsVersion
	}
	return 0
}

func (x *ConfigNew) GetSyncEs() *ESDB {
	if x != nil {
		return x.SyncEs
	}
	return nil
}

func (x *ConfigNew) GetConvertEs() *ESDB {
	if x != nil {
		return x.ConvertEs
	}
	return nil
}

func (x *ConfigNew) GetSync() *Sync {
	if x != nil {
		return x.Sync
	}
	return nil
}

func (x *ConfigNew) GetEsIndex() *ESIndex {
	if x != nil {
		return x.EsIndex
	}
	return nil
}

func (x *ConfigNew) GetRpc() *RPC {
	if x != nil {
		return x.Rpc
	}
	return nil
}

func (x *ConfigNew) GetConvert() *Convert {
	if x != nil {
		return x.Convert
	}
	return nil
}

func (x *ConfigNew) GetKafka() *Kafka {
	if x != nil {
		return x.Kafka
	}
	return nil
}

func (x *ConfigNew) GetManagerAddress() []string {
	if x != nil {
		return x.ManagerAddress
	}
	return nil
}

func (x *ConfigNew) GetContracts() []*Contract {
	if x != nil {
		return x.Contracts
	}
	return nil
}

type Contract struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Abi  string `protobuf:"bytes,1,opt,name=Abi,proto3" json:"Abi,omitempty"`
	Bin  string `protobuf:"bytes,2,opt,name=Bin,proto3" json:"Bin,omitempty"`
	Type string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *Contract) Reset() {
	*x = Contract{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exdbconfig_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Contract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contract) ProtoMessage() {}

func (x *Contract) ProtoReflect() protoreflect.Message {
	mi := &file_exdbconfig_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contract.ProtoReflect.Descriptor instead.
func (*Contract) Descriptor() ([]byte, []int) {
	return file_exdbconfig_proto_rawDescGZIP(), []int{8}
}

func (x *Contract) GetAbi() string {
	if x != nil {
		return x.Abi
	}
	return ""
}

func (x *Contract) GetBin() string {
	if x != nil {
		return x.Bin
	}
	return ""
}

func (x *Contract) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

var File_exdbconfig_proto protoreflect.FileDescriptor

var file_exdbconfig_proto_rawDesc = []byte{
	0x0a, 0x10, 0x65, 0x78, 0x64, 0x62, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x63, 0x6f, 0x6e, 0x76, 0x65,
	0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xab, 0x02, 0x0a, 0x07, 0x43, 0x68, 0x61,
	0x69, 0x6e, 0x33, 0x33, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x72,
	0x70, 0x63, 0x48, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x72,
	0x70, 0x63, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79,
	0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62,
	0x6f, 0x6c, 0x12, 0x2c, 0x0a, 0x11, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x43, 0x68, 0x61, 0x69, 0x6e,
	0x47, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x6f,
	0x74, 0x68, 0x65, 0x72, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x47, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73,
	0x12, 0x22, 0x0a, 0x0c, 0x70, 0x65, 0x72, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x69, 0x6e,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x70, 0x65, 0x72, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x43, 0x6f, 0x69, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b,
	0x53, 0x65, 0x71, 0x18, 0x07, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0b, 0x72, 0x6f, 0x6c, 0x6c, 0x62,
	0x61, 0x63, 0x6b, 0x53, 0x65, 0x71, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x69, 0x6e, 0x50, 0x72,
	0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x63,
	0x6f, 0x69, 0x6e, 0x50, 0x72, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x0f,
	0x70, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x70, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x22, 0x47, 0x0a, 0x05, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x12,
	0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68,
	0x6f, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70,
	0x69, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x22,
	0x6c, 0x0a, 0x04, 0x45, 0x53, 0x44, 0x42, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x75, 0x6c, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x04, 0x62, 0x75, 0x6c, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x70,
	0x77, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x77, 0x64, 0x22, 0x5d, 0x0a,
	0x07, 0x45, 0x53, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x26, 0x0a, 0x0e, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x4f, 0x66, 0x53, 0x68, 0x61, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0e, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x53, 0x68, 0x61, 0x72, 0x64, 0x73,
	0x12, 0x2a, 0x0a, 0x10, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x52, 0x65, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x4f, 0x66, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x22, 0xef, 0x01, 0x0a,
	0x03, 0x52, 0x50, 0x43, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x77, 0x68, 0x69, 0x74,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x77, 0x68, 0x69,
	0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65,
	0x72, 0x48, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x77, 0x61,
	0x67, 0x67, 0x65, 0x72, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c,
	0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0c, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x24, 0x0a, 0x0d, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x77, 0x61, 0x67, 0x67, 0x65,
	0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x53,
	0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x6a, 0x72, 0x70, 0x63, 0x48, 0x6f,
	0x73, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6a, 0x72, 0x70, 0x63, 0x48, 0x6f,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6a, 0x72, 0x70, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6a, 0x72, 0x70, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xf4,
	0x01, 0x0a, 0x04, 0x53, 0x79, 0x6e, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x75, 0x73, 0x68, 0x48, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x75, 0x73, 0x68, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x75, 0x73, 0x68, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x75, 0x73, 0x68, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x75, 0x73, 0x68, 0x42, 0x69, 0x6e, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x75, 0x73, 0x68, 0x42, 0x69, 0x6e, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x53, 0x65, 0x71, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x53, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x0b, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x26, 0x0a,
	0x0e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x75, 0x73, 0x68, 0x46, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x75, 0x73, 0x68, 0x46,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x22, 0xf7, 0x02, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x65,
	0x78, 0x65, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0d, 0x65, 0x78, 0x65, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65,
	0x73, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x53, 0x65, 0x71, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x53, 0x65, 0x71, 0x12, 0x25, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x25, 0x0a, 0x04, 0x73, 0x74, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x04, 0x73, 0x74, 0x61, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x45, 0x78, 0x65, 0x63, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x45, 0x78, 0x65, 0x63, 0x12, 0x26, 0x0a,
	0x0e, 0x64, 0x65, 0x61, 0x6c, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x64, 0x65, 0x61, 0x6c, 0x4f, 0x74, 0x68, 0x65, 0x72,
	0x43, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x2c, 0x0a, 0x11, 0x6f, 0x70, 0x65, 0x6e, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x11, 0x6f, 0x70, 0x65, 0x6e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x61, 0x76, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x73, 0x61, 0x76, 0x65,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x22,
	0xca, 0x03, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4e, 0x65, 0x77, 0x12, 0x24, 0x0a,
	0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x33, 0x33, 0x52, 0x05, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x62, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x62, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x65,
	0x73, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x65, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x07, 0x73, 0x79,
	0x6e, 0x63, 0x5f, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x53, 0x44, 0x42, 0x52, 0x06, 0x73, 0x79, 0x6e, 0x63, 0x45, 0x73,
	0x12, 0x2a, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x5f, 0x65, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x53, 0x44,
	0x42, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x45, 0x73, 0x12, 0x1f, 0x0a, 0x04,
	0x73, 0x79, 0x6e, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x04, 0x73, 0x79, 0x6e, 0x63, 0x12, 0x29, 0x0a,
	0x08, 0x65, 0x73, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x53, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52,
	0x07, 0x65, 0x73, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1c, 0x0a, 0x03, 0x72, 0x70, 0x63, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x50,
	0x43, 0x52, 0x03, 0x72, 0x70, 0x63, 0x12, 0x28, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74,
	0x12, 0x22, 0x0a, 0x05, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x52, 0x05, 0x6b,
	0x61, 0x66, 0x6b, 0x61, 0x12, 0x27, 0x0a, 0x0f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2d, 0x0a,
	0x09, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x52, 0x09, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x22, 0x42, 0x0a, 0x08,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x41, 0x62, 0x69, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x41, 0x62, 0x69, 0x12, 0x10, 0x0a, 0x03, 0x42, 0x69,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x42, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_exdbconfig_proto_rawDescOnce sync.Once
	file_exdbconfig_proto_rawDescData = file_exdbconfig_proto_rawDesc
)

func file_exdbconfig_proto_rawDescGZIP() []byte {
	file_exdbconfig_proto_rawDescOnce.Do(func() {
		file_exdbconfig_proto_rawDescData = protoimpl.X.CompressGZIP(file_exdbconfig_proto_rawDescData)
	})
	return file_exdbconfig_proto_rawDescData
}

var file_exdbconfig_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_exdbconfig_proto_goTypes = []interface{}{
	(*Chain33)(nil),    // 0: proto.Chain33
	(*Kafka)(nil),      // 1: proto.Kafka
	(*ESDB)(nil),       // 2: proto.ESDB
	(*ESIndex)(nil),    // 3: proto.ESIndex
	(*RPC)(nil),        // 4: proto.RPC
	(*Sync)(nil),       // 5: proto.Sync
	(*Convert)(nil),    // 6: proto.Convert
	(*ConfigNew)(nil),  // 7: proto.ConfigNew
	(*Contract)(nil),   // 8: proto.Contract
	(*ExecConfig)(nil), // 9: proto.ExecConfig
	(*StatConfig)(nil), // 10: proto.StatConfig
}
var file_exdbconfig_proto_depIdxs = []int32{
	9,  // 0: proto.Convert.data:type_name -> proto.ExecConfig
	10, // 1: proto.Convert.stat:type_name -> proto.StatConfig
	0,  // 2: proto.ConfigNew.chain:type_name -> proto.Chain33
	2,  // 3: proto.ConfigNew.sync_es:type_name -> proto.ESDB
	2,  // 4: proto.ConfigNew.convert_es:type_name -> proto.ESDB
	5,  // 5: proto.ConfigNew.sync:type_name -> proto.Sync
	3,  // 6: proto.ConfigNew.es_index:type_name -> proto.ESIndex
	4,  // 7: proto.ConfigNew.rpc:type_name -> proto.RPC
	6,  // 8: proto.ConfigNew.convert:type_name -> proto.Convert
	1,  // 9: proto.ConfigNew.kafka:type_name -> proto.Kafka
	8,  // 10: proto.ConfigNew.Contracts:type_name -> proto.Contract
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_exdbconfig_proto_init() }
func file_exdbconfig_proto_init() {
	if File_exdbconfig_proto != nil {
		return
	}
	file_convert_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_exdbconfig_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chain33); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_exdbconfig_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Kafka); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_exdbconfig_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ESDB); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_exdbconfig_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ESIndex); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_exdbconfig_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RPC); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_exdbconfig_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sync); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_exdbconfig_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Convert); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_exdbconfig_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigNew); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_exdbconfig_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Contract); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_exdbconfig_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_exdbconfig_proto_goTypes,
		DependencyIndexes: file_exdbconfig_proto_depIdxs,
		MessageInfos:      file_exdbconfig_proto_msgTypes,
	}.Build()
	File_exdbconfig_proto = out.File
	file_exdbconfig_proto_rawDesc = nil
	file_exdbconfig_proto_goTypes = nil
	file_exdbconfig_proto_depIdxs = nil
}
