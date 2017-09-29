# Ref

1. [[Oracle] 處理中文亂碼 -- NLS_LANG 設定](http://wupinny.blogspot.tw/2010/11/oracle-nlslang.html)
2. [ODAC1120320_x64](http://www.oracle.com/technetwork/database/windows/downloads/index-090165.html)


# Flow

1. Get SQL from PL/SQL
    * Update Table:
      1. right click Table
      2. choose Colums Tab
      3. alter Table
      4. click `View SQL` and copy SQL
      5. click `Apply` to commit
2. Create `xxx.sql` file to `<FOLDER>\SQL` , `<FOLDER>\Tables` or `<FOLDER>\Sequences`.  `<FOLDER>` as blow:
    * `DEV`: DEVDB for development (SAT)
    * `UAT`: TESTDB for user (UAT)
    * `PreProb`: DRDB will copy Prob DB each day.
    * `Prod2`: for product
    * Create Table need `TAFIR` authority or `xxx.sql` file name `xxx` equal to table name.
3. Create `<ID>[_<ORDER_ID>]_<DATE>_<No>.bat`
    ```bat
      @ECHO OFF
      SET NLS_LANG=TRADITIONAL CHINESE_TAIWAN.ZHT16MSWIN950
      CALL ONLINE \path\to\xxx.sql
      @ECHO ON
    ```
4. Execute bat file

# Other

## Create Table

建立 Table `TABLE_NAME` 則 SQL 檔案需命名為 `TABLE_NAME.sql`

## Garbled Text

Set `NLS_LANG` to `TRADITIONAL CHINESE_TAIWAN.AL32UTF8`.
