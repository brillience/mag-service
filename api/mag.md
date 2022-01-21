
### 1. "登陆"

1. 路由定义

- Url: /user/login
- Method: POST
- Request: `LoginReq`
- Response: `LoginReply`

2. 请求定义


```golang
type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
```


3. 返回定义


```golang
type LoginReply struct {
	Username string `json:"username"`
	Nick string `json:"nick"`
	AccessToken string `json:"accessToken"`
	AccessExpire int64 `json:"accessExpire"`
	RefreshAfter int64 `json:"refreshAfter"`
}
```
  


### 2. "创建Mag摘要文档"

1. 路由定义

- Url: /mag
- Method: POST
- Request: `Abstract`
- Response: `commonResp`

2. 请求定义


```golang
type Abstract struct {
	Docid string `json:"docid"`
	Content string `json:"abstract"`
}
```


3. 返回定义


```golang
type CommonResp struct {
	Ok bool `json:"ok"`
	Error string `json:"error"`
}
```
  


### 3. "更新Mag摘要文档"

1. 路由定义

- Url: /mag
- Method: PUT
- Request: `Abstract`
- Response: `commonResp`

2. 请求定义


```golang
type Abstract struct {
	Docid string `json:"docid"`
	Content string `json:"abstract"`
}
```


3. 返回定义


```golang
type CommonResp struct {
	Ok bool `json:"ok"`
	Error string `json:"error"`
}
```
  


### 4. "通过docid获取Mag摘要文档"

1. 路由定义

- Url: /mag/id/:id
- Method: GET
- Request: `reqAbsId`
- Response: `Abstract`

2. 请求定义


```golang
type ReqAbsId struct {
	Docid string `from:"docid"`
}
```


3. 返回定义


```golang
type Abstract struct {
	Docid string `json:"docid"`
	Content string `json:"abstract"`
}
```
  


### 5. "支持通配符的模糊搜索"

1. 路由定义

- Url: /mag/search
- Method: POST
- Request: `reqKeyWord`
- Response: `Abstracts`

2. 请求定义


```golang
type ReqKeyWord struct {
	Key string `from:"key"`
}
```


3. 返回定义


```golang
type Abstracts struct {
	Data []Abstract `json:"data"`
}
```
  


### 6. "通过docid获取Mag摘要的NLP标签"

1. 路由定义

- Url: /mag/nlp
- Method: POST
- Request: `reqAbsId`
- Response: `NlpTags`

2. 请求定义


```golang
type ReqAbsId struct {
	Docid string `from:"docid"`
}
```


3. 返回定义


```golang
type NlpTags struct {
	DocId string `json:"docid"`
	Tags string `json:"tags"`
}
```
  

