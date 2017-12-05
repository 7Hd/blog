
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
  * `violation of array bounds` 矩陣溢位
    * A good way to prevent such errors is to replace arrays by well-tested container classes
      * The standard template library (STL) is a useful source of such container classes.
      * but need to `avoid` dynamic memory allocation
    * Text strings are particularly problematic because there may be no certain limit to the length of a string.
      * The old `C-style` method of storing strings in character arrays is `fast` and `efficient`, but `not safe unless the length of each string is checked` before storing.
      * The standard `solution` to this problem is to use string classes, such as `string` or `CString`.
        * `safe` and `flexible`, but quite `inefficient` in large applications.
        * 每次新增/修改都分配新的記憶體空間
          * This can cause the `memory to be fragmented` and involve a high overhead cost of `heap management` and `garbage collection`.
      * A more efficient `solution` that doesn't compromise safety is to store all strings in one memory pool.
  * `Integer overflow`
    * The official C standard says that the behavior of `signed integers` in case of overflow is `"undefined"`.
    * In the case of the `Gnu` compiler
      * There are a number of possible remedies against this problem:
        * check for overflow before it occurs,
        * use unsigned integers - they are guaranteed to wrap around,
        * trap integer overflow with the option `-ftrapv`, but this is extremely inefficient,
        * get a compiler warning for such optimizations with option `-Wstrict-overflow=2`, or
        * make the overflow behavior well-defined with option `-fwrapv` or `-fno-strict-overflow`.
