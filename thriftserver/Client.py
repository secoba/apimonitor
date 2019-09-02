import sys
sys.path.append('./gen-py')


# from Order.rpc import RpcService
# from Order.rpc.ttypes import *
# from Order.rpc.constants import *
from ThriftServer.rpc import RpcService
from ThriftServer.rpc.ttypes import *
from ThriftServer.rpc.constants import *

from thrift import Thrift
from thrift.transport import TSocket
from thrift.transport import TTransport
from thrift.protocol import TBinaryProtocol


def main():
  try:
    transport = TSocket.TSocket('192.168.248.188', 11000)
   
    # Buffering is critical. Raw sockets are very slow
    transport = TTransport.TBufferedTransport(transport)
   
    # Wrap in a protocol
    protocol = TBinaryProtocol.TBinaryProtocol(transport)
   
    # Create a client to use the protocol encoder
    client = RpcService.Client(protocol)
   
    # Connect!
    transport.open()
   
    # res = client.testOne("oppo","192.168.248.188")

    # res = client.devIce()
    # print(res)
    # order = client.Order("oppo")

    # atx = client.Atx()

    # adb = client.Adb()

    # print(adb)

    # devs = client.devIce()

    print(client.adbServer())

    # print(order,atx,adb,res,devs)
    transport.close()

  except Thrift.TException as tx:
    print(tx)

if __name__ == '__main__':
    main()