
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
  * The development time and maintainability of C++ projects can be improved by consistent modularity and reusable classes. (模組化)
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

# Finding the biggest time consumers

使用 CPU `clock cycles` 作為測量單位

## Use a profiler to find hot spots

* Instrumentation
  * compiler 插入額外的程式碼紀錄 function 執行次數與時間
* Debugging
* Time-based sampling 以時間為基礎的採樣
  * 統計每次採樣時在哪個程式區塊
* Event-based sampling
  * 採樣特定 event 發生時所在程式區塊

分析工具可能不准的問題點:

* Coarse time measurement
  * 測量的時間單位太寬
* Execution time too small or too long
* Waiting for user input
* Interference from other processes 受其他程式干擾
* Function addresses are obscured in optimized programs
  * 無法取得對應 function
* Uses debug version of the code
* Jumps between CPU cores
* Poor reproducibility
  * 發生於次數少的特定 event 或不定時的 event

替代方案:

* 中斷點確認流程
* 手動插入 code 測量呼叫次數, 時間
  * 應使用 `#if` 來區隔測量的 code
  * Windows 可使用 `GetTickCount`, `QueryPerformanceCounter` 來取得 millisecond 的時間
  * Windows 可使用 `__rdtsc()` 來取得 CPU time stamp 的時間
    * The time stamp counter becomes `invalid` if a thread jumps between different CPU cores
    * You may have to fix the thread to a specific CPU core during time measurements to avoid this.
      * In Windows, `SetThreadAffinityMask`, in Linux, `sched_setaffinity`

## Most common time-consumers

* Program installation
* Automatic updates
* Program loading
* Dynamic linking and position-independent code
  * Position-independent code is inefficient, especially in 32-bit mode
* File access
* System database
  *  It is more efficient to store application-specific information in a separate file than in the big registration database in the Windows system.
* Other databases
  * It may be possible to replace a database by a plain old data file in simple cases
* Graphics
  * GUI
* Other system resources
  * Accessing system devices and using advanced facilities of the operating system can be time consuming because it may involve the loading of several drivers, configuration files and system modules
* Network access
  * These problems should be taken into account when deciding whether to store help files and other resources locally or remotely. If frequent updates are necessary then it may be optimal to mirror the remote data locally.
* Memory access

