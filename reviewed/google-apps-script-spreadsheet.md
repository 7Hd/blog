# Ref

* [Google developers apps-script](https://developers.google.com/apps-script/)

# About This

讓 Google Spreadsheets 依輸入資料自動更新。

# Code

* `totalSheet`: 原始資料 Sheet 所在位置
* `filterSheet`: 過濾後資料 Sheet 所在位置
* `diff`: 30 天毫秒數

## Filter Contents

```js
function ouFilter() {
  var verifyDoneDate = 9, totalSheet = 0;
  var diff = 30 * 24 * 60 * 60 * 1000;
  var res = [];
  var sheet = SpreadsheetApp.getActive().getSheets()[totalSheet];
  var datas = sheet.getDataRange().getDisplayValues();

  res.push(datas[0]);

  for (var i = 1; i < datas.length; i++) {
    if(!datas[i][verifyDoneDate]) { res.push(datas[i]); continue; }
    var today = new Date(), verify = new Date(datas[i][verifyDoneDate]);
    if( (today-verify) < diff) { res.push(datas[i]); continue; }
  }

  return res;
}
```

## Set Sheet Contents

```js
function ouSetupSheet() {
  var filterSheet = 1;
  var res = ouFilter();
  var sheet = SpreadsheetApp.getActive().getSheets()[filterSheet];
  sheet.clear();
  for(var i = 0; i < res.length; i++){
    sheet.appendRow(res[i]);
  }
}
```
