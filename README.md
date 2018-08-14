# simple_rest

  RESTFul API 範例框架

## 專案結構

| 路徑 | 說明 |
|---|---|
| api | RESTFul API 伺服器相關 |
| api/controller | Http API 流程控制層 |
| api/middleware | gin 框架中間層 |
| api/protocol | Http 通用協定 |
| config | 設定檔讀取套件 |
| service | 業務邏輯層 |
| env | 常數 |

## Usage

### 啟動Mysql服務

在本專案底下執行:

```bash
cd $GOPATH/simple_rest/docker && docker-compose up -d
```

若要關閉Mysql則執行以下指令:

```bash
cd $GOPATH/simple_rest/docker && docker-compose down
```


### 啟動 RESTful Server

在本專案底下執行:

```bash
cd $GOPATH/simple_rest && go run main.go
```

此專案預設有提供三組API，啟動服務後，可執行以下語法測試伺服器是否正常運行:

```bash
curl -G http://127.0.0.1:8000/v1/get  \
--data-urlencode "Colors[]=red"   \
--data-urlencode "Colors[]=green" \
--data-urlencode "Name=hank" \
--data-urlencode "Birthday=2018-08-14T12:00:00+08:00" 
```

```bash
curl -X POST http://127.0.0.1:8000/v1/post -d 'Colors[]=red&Colors[]=green&Name=hank&Birthday=2018-08-14T12:00:00%2d08:00'
```

```bash
curl -G http://127.0.0.1:8000/v1/get?UserID=0
```

#### 預設三組API

*GET範例*

    GET API 範例，會回傳傳入的資訊

URL: `http://127.0.0.1:8000/v1/get`

協定: `GET`

Input:

| 變數名稱 | 型態 | 是否必要 | 說明 |
|---|---|---|---|
| Colors[] | []string | |  |
| Name | string | * |  |
| Birthday | time | | 時間格式須為 RFC3339，Example: 2018-08-05T12:00:00+08:00 |
| Address | string | |  |

*POST範例*

    POST API 範例，會回傳傳入的資訊

URL: `http://127.0.0.1:8000/v1/post`

協定: `POST`

Input:

| 變數名稱 | 型態 | 是否必要 | 說明 |
|---|---|---|---|
| Colors[] | []string | |  |
| Name | string | * |  |
| Birthday | time | | 時間格式須為 RFC3339，Example: 2018-08-05T12:00:00+08:00 |
| Address | string | |  |

*取得使用者資訊*

    從資料庫中撈取使用者資訊

URL: `http://127.0.0.1:8000/v1/user`

協定: `GET`

Input:

| 變數名稱 | 型態 | 是否必要 | 說明 |
|---|---|---|---|
| UserID | string | * | 要撈取的使用者ID |

Output: 

| 變數名稱 | 型態 | 說明 |
|---|---|---|---|
| ID | int32 | 使用者ID |
| Account | string | 使用者帳號 |
| Password | string | 使用者密碼 |

## 啟動Mysql時的預設資料結構

此專案使用docker來建置資料庫，其預設資料放在 docker/init_file/init.sql 中