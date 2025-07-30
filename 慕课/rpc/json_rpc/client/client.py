import json
import socket

if __name__ == '__main__':
    request = {
        "id": 0,
        "params": ["dawn"],
        "method": "HelloService.Hello"
    }

    client = socket.create_connection(("localhost", 1234))
    client.sendall(json.dumps(request).encode())
    rsp = client.recv(1024)  # 获取服务器返回的数据
    rsp = json.loads(rsp.decode())
    print(rsp["result"])
