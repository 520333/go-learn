import json

import requests
if __name__ == '__main__':
    request = {
    "id": 0,
    "params": ["dawn"],
    "method": "HelloService.Hello"
    }
    resp=requests.post("http://localhost:1234/jsonrpc",json=request)
    print(json.loads(resp.text)['result'])
