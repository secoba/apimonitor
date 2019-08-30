namespace go ThriftServer.rpc
namespace java ThriftServer.rpc
namespace py ThriftServer.rpc

// 测试服务
service RpcService {
 
    // 发起远程调用
    list<string> funCall(1:i64 callTime, 2:string funCode, 3:map<string, string> paramMap),
    string atxServer(),
    string adbServer(),
    string actServer(),
}
