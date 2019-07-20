from werkzeug.wrappers import Request, Response
from werkzeug.serving import run_simple

from jsonrpc import JSONRPCResponseManager, dispatcher

@Request.application
def application(request):
    print(dispatcher)
    print("*"*10)
    # Dispatcher is dictionary {<method_name>: callable}
    # dispatcher["Say.Hello"] = lambda s: "hello " + s["name"]
    

    # response = JSONRPCResponseManager.handle(
    #     request.data, dispatcher)
    # print(response.json)
    # return Response(response.json, mimetype='application/json')
    return {"code":2}


if __name__ == '__main__':
    run_simple('localhost', 4000, application)
