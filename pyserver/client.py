import sys
sys.path.append('./gen-py')
 
# from helloworld import HelloWorld
# from helloworld.ttypes import *
# from helloworld.constants import *

from testone.rpc import RpcService
from testone.rpc.ttypes import *
from testone.rpc.constants import *


from thrift import Thrift
from thrift.transport import TSocket
from thrift.transport import TTransport
from thrift.protocol import TBinaryProtocol
 
try:
  # Make socket
  transport = TSocket.TSocket('127.0.0.1', 19090)
 
  # Buffering is critical. Raw sockets are very slow
  transport = TTransport.TBufferedTransport(transport)
 
  # Wrap in a protocol
  protocol = TBinaryProtocol.TBinaryProtocol(transport)
 
  # Create a client to use the protocol encoder
  # client = HelloWorld.Client(protocol)
  client = RpcService.Client(protocol)
  # Connect!
  transport.open()
 
  # client.ping()
  # print("ping()")
 
  # msg = client.sayHello()
  # print(msg)
  # msg = client.sayMsg("ceshi")
  # print(HELLO_YK)
  # print(msg)
  msg = "ceshi"
  tmp = client.testOne(msg,"123")
  print(tmp)
 
  transport.close()
 
except Thrift.TException as tx:
  print( "%s" % (tx.message))