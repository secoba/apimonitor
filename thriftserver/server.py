import sys
import logging
sys.path.append('./gen-py')
# from helloworld import HelloWorld
# from helloworld.ttypes import *



from Order.rpc import RpcService
from Order.rpc.ttypes import *
from Order.rpc.constants import *


from thrift.transport import TSocket
from thrift.transport import TTransport
from thrift.protocol import TBinaryProtocol
from thrift.server import TServer
  
import socket



class ThriftHandler:
  def __init__(self):
    self.log = {}
    
  def Adb(self):
    data = {}
    data["mod"] = "1"
    data["port"] = "2"
    data["version"] = "3"
    data["memory"] = "4"
    data["cpu"] = "5"
    data["level"] = "6"
    data["sdk"] = "7"
    data["ip"] = "8"
    data["Online"] = "9"
    data["name"] = "10"
    data["img"] = "10"
    res = []
    res.append(data)
    return res
 

 

def ThriftServer():
  
  handler = ThriftHandler()
  processor = RpcService.Processor(handler)
  transport = TSocket.TServerSocket('192.168.248.188',10001)
  tfactory = TTransport.TBufferedTransportFactory()
  pfactory = TBinaryProtocol.TBinaryProtocolFactory()
   
  server = TServer.TSimpleServer(processor, transport, tfactory, pfactory)
  print("thrift server start")
  # print("Starting python server...")
  server.serve()
  # print("done!")

if __name__ == '__main__':
  ThriftServer()

