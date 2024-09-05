# 配置文件
config.json
```
{
  "moonshot": {
    "api_key": "<kimi api key>"
  }
}
```

moonshoot的api key从这里获取 https://platform.moonshot.cn/console/api-keys

# Docker运行

## Docker build
```
docker build -t shangxiehui-ai:latest .
```

## Docker run
```
docker run -v ./config.json:/config.json -p 8080:8080 -it shangxiehui-ai /apiserver --config.file=/config.json
```


# Chat API

URL: http://localhost:8080/v1/chat/completion/text/stream
Method: POST
Body:
```
{
  "messages": [
    {
      "role": "user",
      "message": "你好"
    },
    {
      "role": "assistant",
      "message": "有什么可以帮助你的吗？"
    },
    {
      "role": "user",
      "message": "你叫什么名字"
    }
  ]
}
```

messages是一个列表，可以带上历史聊天记录，role是user为用户的问题，role是assistant为机器人的回答。
由于Prompt长度有限，前端可以传有限的历史记录，比如最后10条对话记录。

Response:
```
data:{"role":"assistant","chunk":"你好"}

data:{"role":"assistant","chunk":"呀"}

data:{"role":"assistant","chunk":"！"}

data:{"role":"assistant","chunk":"我是"}

data:{"role":"assistant","chunk":"你的"}

...
```
Response是text/event格式的http response，每次返回一个chunk，前端可以逐步显示出来。

# 前端代码示例
```javascript
const processStream = async (reader: any, callback: (line: string) => void) => {
  let buf = "";
  while (true) {
    // .read() returns 2 properties
    const result = await reader?.read();

    // if done is true
    if (result?.done) {
      console.log("stream completed");
      break;
    }
    // value is a binary data in Uint8Array format, as Uint8Array is suitable data structure for binary data
    // we decode Uint8Array using TextDecoder
    let chunk = new TextDecoder("utf-8").decode(result?.value);

    buf += chunk;

    const lines = buf.split("\n");
    if (lines.length > 1) {
      for (let i = 0; i < lines.length - 1; i++) {
        const line = lines[i].trim();
        if (line) {
          callback(lines[i]);
        }
      }

      buf = lines[lines.length - 1];
    }
  }

  if (buf != "") {
    callback(buf);
  }
};

const res = await fetch(url, {
  method: method,
  body: data,
});

if (res.ok) {
  const reader = res.body?.getReader();
  processStream(reader, (line: string) => {
    if (line.startsWith("data:")) {
      line = line.replace(/^data:/, "");
      // 实际的数据，需要json解析后获取内容
      console.log(line);
    }
  }).catch(err => {
    console.error(err);
  });
}
```
