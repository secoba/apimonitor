import sys
sys.path.append('./gen-py')
  
# from helloworld import HelloWorld
# from helloworld.ttypes import *
 

from testone.rpc import RpcService
from testone.rpc.ttypes import *
from testone.rpc.constants import *


from thrift.transport import TSocket
from thrift.transport import TTransport
from thrift.protocol import TBinaryProtocol
from thrift.server import TServer
  
import socket
 
class HelloWorldHandler:
  def __init__(self):
    self.log = {}
 
  # def ping(self):
  #   print ("ping()")
 
  # def sayHello(self):
  #   # print ("sayHello()")
  #   return "say hello from ".format( socket.gethostbyname(socket.gethostname()))
 
  # def sayMsg(self, msg):
  #   # print( "sayMsg:{}".format (msg))
  #   return "say {} from {}".format(msg,socket.gethostbyname(socket.gethostname()))
  def testOne(self,msg):
    return "osk"+msg
 
handler = HelloWorldHandler()
processor = RpcService.Processor(handler)
transport = TSocket.TServerSocket('127.0.0.1',19090)
tfactory = TTransport.TBufferedTransportFactory()
pfactory = TBinaryProtocol.TBinaryProtocolFactory()
 
server = TServer.TSimpleServer(processor, transport, tfactory, pfactory)
 
print("Starting python server...")
server.serve()
print("done!")