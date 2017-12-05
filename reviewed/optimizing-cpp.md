
# Introduction


# Choosing the optimal platform

## Hardware

## Microprocessor

## OS

## Language

## Compiler

## Function Libraries

## UI Framework

## Overcoming the drawbacks of the C++ language

* Portability 可移植性
  * 以下幾種不利於移植的應另外寫 `module`
    * `direct access to hardware interfaces`
    * `system call`
* Development time 開發時間
  * The development time and maintainability of C++ projects can be improved by consistent modularity and reusable classes.
* Security 安全
  * Standard C++ imple-mentations have **no checking** for `array bounds violations` and `invalid pointers`.
    * It is necessary to adhere to certain programming principles in order to prevent such errors in programs where security matters.
  * `invalid pointers` can be avoided
    * by using references instead of pointers
    * by initializing pointers to zero
    * by setting pointers to zero whenever the objects they point to become invalid
    * by avoiding pointer arithmetics and pointer type casting
      * 避免 `pointer` 運算及型別轉換
    * Avoid the function `scanf`
  * `violation of array bounds`
    * A good way to prevent such errors is to replace arrays by well-tested container classes
      * The standard template library (STL) is a useful source of such container classes.
      * but need to `avoid` dynamic memory allocation
