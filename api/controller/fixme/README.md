# fixme

GetDeposit 此方法在大量Request同時存取的情況下，資料庫中的值會有出現負數的情況。

嘗試修改程式邏輯避免資料庫出現負數情況。

壓力測試範例語法 :

```bash
 ab -c 20 -n 10000  'http://127.0.0.1:8000/v1/deposit?UserID=1&Amount=1'
```
