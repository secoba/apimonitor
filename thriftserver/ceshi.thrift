namespace go Order.rpc
namespace java Order.rpc
namespace py Order.rpc

// 测试服务
service RpcService {
 
    // 发起远程调用
    list<string> funCall(1:i64 callTime, 2:string funCode, 3:map<string, string> paramMap),
    string testOne(1:string msg,2:string ip),
    string setLog(1:string msg),
    list<string> Adb(),
    string Atx(),
    string Order(1:string app),
    string devIce(),
    string dev_UP(),
}
