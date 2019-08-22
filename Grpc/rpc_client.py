# -*- coding: utf-8 -*-
import grpc
import os
import sys
from grpc._channel import _Rendezvous

current_fonder_path = os.path.split(os.path.realpath(__file__))[0]
print (current_fonder_path)
protocal_path = os.path.join(current_fonder_path,"..","example")
print (protocal_path)
sys.path.append(protocal_path)
import data_pb2, data_pb2_grpc

_HOST = 'localhost'
_PORT = '8082'


def run():
    conn = grpc.insecure_channel(_HOST + ':' + _PORT)  # 监听频道
    client = data_pb2_grpc.FormatDataStub(channel=conn)   # 客户端使用Stub类发送请求,参数为频道,为了绑定链接
    response = client.DoFormat(data_pb2.actionrequest(text='hello,world!',corpus="NEWS"))   # 返回的结果就是proto中定义的类
    for i in response.result:
        print(i)


if __name__ == '__main__':
    try:
        run()
    except _Rendezvous as e:
        print("connect error")
