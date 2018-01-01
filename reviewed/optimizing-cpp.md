
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
  * It is more efficient to store application-specific information in a separate file than in the big registration database in the Windows system.
* Other databases
  * It may be possible to replace a database by a plain old data file in simple cases
* Graphics
  * GUI
* Other system resources
  * Accessing system devices and using advanced facilities of the operating system can be time consuming because it may involve the loading of several drivers, configuration files and system modules
* Network access
  * These problems should be taken into account when deciding whether to store help files and other resources locally or remotely. If frequent updates are necessary then it may be optimal to mirror the remote data locally.
* Memory access
* Context switches
* Dependency chains
  * In order to take advantage of out-of-order execution, you have to avoid long dependency chains.
* Execution unit throughput 單位吞吐量
  * There is an important distinction between the latency and the throughput of an execution unit.
  * 若運算無相依性則可同可同步執行

# Performance and usability

對於效能與可用性上應注意的項目:

* Big runtime frameworks
* Memory swapping
  * developers  RAM !== end user RAM
* Installation problems
  * The procedures for installation and uninstallation of programs should be standardized and done by the operating system rather than by individual installation tools.
* Automatic updates
* Compatibility problems
  * All software should be tested on different platforms, different screen resolutions, different system color settings and different user access rights.
* Copy protection
* Hardware updating
* Security
* Background services
* Feature bloat
* Take user feedback seriously

# Choosing the optimal algorithm

* Don't use an advanced and complicated algorithm if a simple algorithm can do the job fast enough.
* Boost collection contains well-tested libraries for many common purposes
  * http://www.boost.org/
* The insight you gain by testing and analyzing program performance and studying the bottlenecks can lead to a better understanding of the whole structure of the problem.
  * 通過測試和分析程序性能以及研究瓶頸可以獲得的洞察力可以更好地理解問題的整體結構
* A complete redesign of a program that already works is of course a considerable job, but it may be quite a good investment.
* A redesign can not only improve the performance, it is also likely to lead to a more well-structured program that is easier to maintain.
  * The time you spend on redesigning a program may in fact be less than the time you would have spent fighting with the problems of the original, poorly designed program.

# Development process

For team projects, it is recommended to use a version control tool.

# The efficiency of different C++ constructs

* double precision calculations are just as fast as single precision
* template class is more efficient than a polymorphous class

## Different kinds of variable storage

* Storage on the stack
  * Variables and objects declared inside a function are stored on the stack
  * The stack is a part of memory that is organized in a first-in-last-out fashion.
  * It is used for storing:
    * function return addresses (i.e. where the function was called from),
    * function parameters,
    * local variables,
    * and for saving registers that have to be restored before the function returns.
  * The stack is the `most efficient memory space to store data` because the same range of memory addresses is reused again and again.
  * **all variables and objects should preferably be declared inside the function in which they are used.**
* Global or static storage
  * Variables that are declared outside of any function are called global variables.
  * Global variables are stored in a static part of the memory.
  * The advantage of static data is that it can be initialized to desired values before the program starts.
  * The disadvantage is that the memory space is occupied throughout the whole program execution
  * String constants and floating point constants are stored in static memory in optimized code.
  * Integer constants are usually included as part of the instruction code.
    * can assume that there are no caching problems for integer constants.
* Register storage
  * A register is a small piece of memory inside the CPU used for temporary storage.
  * Variables that are stored in registers are `accessed very fast`.
  * The same register can be used for multiple variables as long as their uses (live ranges) do not overlap.
* Volatile
  * The `volatile` keyword specifies that a variable can be changed by another thread.
  * This prevents the compiler from making optimizations that rely on the assumption that the variable always has the value it was assigned previously in the code.
  * The effect of the keyword `volatile` is that it makes sure the variable is stored in memory rather than in a register and prevents all optimizations on the variable.
    * This can be useful in test situations to avoid that some expression is optimized away.
  * It `doesn't` prevent two threads from attempting to write the variable at the same time.
* Thread-local storage
  * Most compilers can make thread-local storage of static and global variables by using the keyword `__thread` or `__declspec(thread)`.
  * Thread-local storage is `inefficient` because it is accessed through a pointer stored in a thread environment block
  * Thread-local storage should be avoided
  * Variables stored on the stack always belong to the thread in which they are created.
* Far
  * keyword `far` (arrays can also be `huge`)
  * Far storage, far pointers, and far procedures are `inefficient`
  * recommended to use a different operating systems that allows bigger segments
* Dynamic memory allocation
  * Dynamic memory allocation is done with the operators `new` and `delete` or with the functions `malloc` and `free`.
    * These operators and functions `consume a significant amount of time`.
  * A part of memory called the `heap` is reserved for dynamic allocation.
  * The heap can easily become fragmented when objects of different sizes are allocated and deallocated in random order.
  * The heap manager can spend a lot of time cleaning up spaces that are no longer used and searching for vacant spaces.
    * This is called `garbage collection`.
  * Dynamic memory allocation also tends to make the code more complicated and error-prone.
* Variables declared inside a class
  * Variables declared inside a class are stored in the order in which they appear in the class declaration.
  * The type of storage is determined where the object of the class is declared
  * An object cannot be stored in a register except in the simplest cases, but its data members can be copied into registers.
  * A class member variable with the `static` modifier will be stored in static memory and will have one and only one instance.
  * Storing variables in a class or structure is a good way of making sure that variables that are used in the same part of the program are also stored near each other.

## Integers variables and operators

* Integer sizes
  * Integers can be different sizes, and they can be `signed` or `unsigned`.
  * standard header file `stdint.h` or `inttypes.h` is available then it is recommended to use that for a portable way of defining integer types of a specific size.
  * it is `inefficient` to use an integer size that is larger than the largest available register size.
  * It is `recommended` to use the default integer size in cases where the size doesn't matter and there is no risk of overflow
    * In large arrays, it may be preferred to use the smallest integer size that is big enough for the specific purpose in order to make better use of the data cache.
  * When considering whether a particular integer size is big enough for a specific purpose,
    * must consider if intermediate calculations can cause overflow.
* Signed versus unsigned integers
  * In most cases, there is no difference in speed between using signed and unsigned integers.
  * But there are a few cases where it matters:
    * Unsigned is faster than signed when you divide (`/`) an integer with a constant
      * This also applies to the modulo operator `%`.
    * Conversion to floating point is faster with signed than with unsigned integers for most instruction sets
    * Overflow behaves differently on signed and unsigned variables.
  * The conversion between signed and unsigned integers is `costless`
  * Be sure not to mix signed and unsigned integers in comparisons, such as `<`.
    * The result of comparing signed with unsigned integers is `ambiguous` and may produce undesired results.
* Integer operators
  * Integer operations are generally very fast.
  * `Multiplication` and `division` take longer time.
    * Integer multiplication takes `11` clock cycles on Pentium 4 processors, and `3 - 4` clock cycles on most other microprocessors.
    * Integer division takes `40 - 80` clock cycles, depending on the microprocessor.
* Increment and decrement operators
  * pre-increment `++i` or post-increment `i++`
  * `x = array[i++]` is more `efficient` than `x = array[++i]`
    * 後者 `x` 取得前須先計算 `++i`

## Floating point variables and operators

* two different types of floating point registers and correspondingly two different types of floating point instructions.
* The `original` method of doing floating point operations involves eight floating point registers organized as a register stack.
  * These registers have long double precision (80 bits).
  * advantages
    * All calculations are done with `long double precision`.
    * Conversions between different precisions take `no extra time`
    * There are intrinsic instructions for `mathematical functions` such as logarithms and trigonometric functions
    * The code is `compact` and `takes little space` in the code cache
  * disadvantages
    * It is difficult for the compiler to make register variables because of the way the register stack is organized.
    * Floating point comparisons are slow
      * unless the Pentium-II or later instruction set is enabled.
    * Conversions between integers and floating point numbers is `inefficient`
    * Division, square root and mathematical functions `take more time` to calculate when long double precision is used.
* A `newer` method of doing floating point operations involves eight or sixteen `vector registers` (XMM or YMM) which can be used for multiple purposes.
  * advantages
    * It is easy to make floating point register variables
    * Vector operations are available for doing `parallel` calculations on vectors of two double precision or four single precision variables in the XMM registers
  * disadvantages
    * Long double precision is not supported
    * The calculation of expressions where operands have `mixed` precision require precision conversion instructions which can be quite `time-consuming`
    * Mathematical functions must use a function library, but this is often faster than the intrinsic hardware functions.
  * In most cases, double precision calculations take no more time than single precision.
    * Long double precision takes only slightly more time.
  * Floating point addition takes `3 - 6` clock cycles, depending on the microprocessor.
    * Multiplication takes `4 - 8` clock cycles. Division takes `14 - 45` clock cycles.
    * Floating point `comparisons` are `inefficient` when the floating point stack registers are used.
    * `Conversions` of float or double to integer `takes a long time` when the floating point stack registers are used.
  * Do `not` mix single and double precision when the XMM registers are used
  * `Avoid` conversions between integers and floating point variables
  * It is strongly recommended to set the flush-to-zero mode unless you have special reasons to use subnormal numbers.
    * ` flush-to-zero` mode
      ```cpp
      // Example 7.5. Set flush-to-zero mode (SSE):
      #include <xmmintrin.h>
      _MM_SET_FLUSH_ZERO_MODE(_MM_FLUSH_ZERO_ON);
      ```
    * `denormals-are-zero` mode
      ```cpp
      // Example 7.6. Set flush-to-zero and denormals-are-zero mode (SSE2):
      #include <xmmintrin.h>
      _mm_setcsr(_mm_getcsr() | 0x8040);
      ```

## Enums

* An `enum` is simply an integer in disguise.
* Enums are exactly as efficient as integers.
* Enums in header files should therefore have `long and unique enumerator names` or be put `into a namespace`.

## Booleans

* The order of Boolean operands
  * `&&`, `||` 容易造成判斷的項目應放在前面
  * If one operand is `more predictable` than the other, then put the most predictable operand first.
  * If one operand is `faster to calculate` than the other then put the operand that is calculated the fastest first.
* Boolean variables are overdetermined
  * Boolean variables are stored as `8-bit integers` with the value 0 for false and 1 for true.
  * Don't change `&&` to `&` unless you expect the `&&` expression to generate many branch mispredictions.
* Boolean vector operations
  * An integer may be used as a Boolean vector.
  * The operators `&`, `|`, `^`, `~` are useful for Boolean vector operations.

## Pointers and references

* Pointers versus references
  * Pointers and references are `equally` efficient because they are in fact doing the same thing.
  * advantages of using pointers
    * `*` 可清楚顯示其為 Pointers
    * It is possible to do things with pointers that are impossible with references.
      * You can change what a pointer points to and you can do arithmetic operations with pointers.
  * advantages of using references
    * The syntax is simpler when using references
    * References are `safer` to use than pointers because in most cases they are sure to point to a valid address
    * References are useful for copy constructors and overloaded operators
    * Function parameters that are declared as constant references accept expressions as arguments while pointers and non-constant references require a variable.
* Efficiency
  * Accessing a variable or object through a pointer or reference may be just as fast as accessing it directly.
  * disadvantages of using pointers and references
    * it requires an extra register to hold the value of the pointer or reference
* Pointer arithmetic
  * A pointer is in fact an `integer` that holds a memory address.
  * Pointer arithmetic operations are therefore as fast as integer arithmetic operations.
  * Comparing two pointers requires only an integer comparison
  * Calculating the difference between two pointers requires a division, which is slow
  * The object pointed to can be accessed approximately `two clock` cycles after the value of the pointer has been calculated
  * it is recommended to calculate the value of a pointer well before the pointer is used.
    `* x = *(p++)` is more efficient than `x = *(++p)`

## Function pointers

* Calling a function through a function pointer typically takes a `few` clock cycles `more` than calling the function directly if the target address can be predicted

## Member pointers

* In simple cases, a data member pointer simply stores the offset of a data member relative to the beginning of the object, and a member function pointer is simply the address of the member function.
* https://stackoverflow.com/questions/670734/c-pointer-to-class-data-member

## Smart pointers

* A smart pointer is an object that behaves like a pointer
  * It has the special feature that the object it points to is deleted when the pointer is deleted
* The purpose of using smart pointers is to make sure the object is deleted properly and the memory released when the object is no longer used
* `auto_ptr`, `shared_ptr`
* `no` extra cost to accessing an object through a smart pointer
* extra `cost` whenever a smart pointer is `created`, `deleted`, `copied` or `transferred` from one function to another.
  * These costs are higher for shared_ptr than for auto_ptr
* If the same function or class is responsible for creating and deleting the object then you `don't` need a smart pointer.

## Arrays

* An array is implemented simply by storing the elements consecutively in memory.
  * No information about the dimensions of the array is stored.
* multidimensional array
  * The size of all but the first dimension may preferably be a power of 2 if the rows are indexed in a non-sequential order in order to make the address calculation more efficient
  * The same advice applies to arrays of structure or class objects.

## Type conversions

  ```cpp
  int i;  float f;
  f = i;                      // Implicit type conversion
  f = (float)i;               // C-style type casting
  f = float(i);               // Constructor-style type casting
  f = static_cast<float>(i);  // C++ casting operator
  ```
* Signed / unsigned conversion
  * Conversions between signed and unsigned integers simply makes the compiler interpret the bits of the integer in a different way.
  * There is no checking for overflow, and the code takes no extra time.
  * These conversions can be used freely `without any cost` in performance.
* Integer size conversion
  * An integer is converted to a longer size by extending the sign-bit if the integer is signed, or by extending with zero-bits if unsigned.
    * This typically takes `one` clock cycle if the source is an arithmetic expression.
* Floating point precision conversion
  * Conversions between `float`, `double` and `long double` `take no extra time` when the floating point register stack is used.
    * It takes between `2 and 15` clock cycles (depending on the processor) when the XMM registers are used.
* Integer to float conversion
  * Conversion of a signed integer to a float or double takes `4 - 16` clock cycles, depending on the processor and the type of registers used
  * Conversion of an unsigned integer takes `longer time`.
    * It is faster to first convert the unsigned integer to a signed integer if there is no risk of overflow
      * `d = (double)(signed int)u`
  * Integer to float conversions can sometimes be avoided by replacing an integer variable by a float variable
    * `i * 2` -> `i * 2.0f`
* Float to integer conversion
  * Conversion of a floating point number to an integer `takes a very long time` unless the SSE2 or later instruction set is enabled
    * Typically, the conversion takes `50 - 100` clock cycles.
  * Possible solutions:
    * Avoid the conversions by using different types of variables.
    * Move the conversions out of the innermost loop by storing intermediate results as floating point.
    * Use 64-bit mode or enable the SSE2 instruction set (requires a microprocessor that supports this)
    * Use `rounding` instead of `truncation` and make a round function using assembly language.
* Pointer type conversion
  * A pointer can be converted to a pointer of a different type, integer
  * These conversions do not produce any extra code.
    * It is simply a matter of interpreting the same bits in a different way or bypassing syntax checks.
* Re-interpreting the type of an object
  * It is possible to make the compiler treat a variable or object as if it had a different type by type-casting its address
  ```cpp
  float x;
  *(int*)&x |= 0x80000000;   // Set sign bit of x
  ```
  * There are a number of dangers to be aware of when type-casting pointers:
    * The trick `violates` the strict aliasing rule of standard C, specifying that two pointers of different types cannot point to the same object (except for char pointers).
    * The trick will `fail` if the object is treated as bigger than it actually is. This above code will fail if an int uses more bits than a float.
    * If you access part of a variable, for example 32 bits of a 64-bit double, then the code will `not be portable` to platforms that use big endian storage.
    * If you access a variable in parts, for example if you write a 64-bit double 32 bits at a time, then the code is likely to execute slower than intended because of a store forwarding delay in the CPU
* Const cast
  * The `const_cast` operator is used for relieving the `const` restriction from a pointer.
    * it doesn't generate any extra code and doesn't take any extra time.
  ```cpp
  const int x = 0; // constant data
  *const_cast<int*>(&x) += 2;
  ```
* Static cast
  * The `static_cast` operator does the same as the C-style type-casting.
* Reinterpret cast
  * The `reinterpret_cast` operator is used for pointer conversions.
  * It does the same as C-style type-casting with a little more syntax check.
  * It does not produce any extra code.
* Dynamic cast
  * The `dynamic_cast` operator is used for converting a pointer to one class to a pointer to another class.
    * It makes a `runtime check` that the conversion is valid.
  * This check makes `dynamic_cast` `more time-consuming` than a simple type casting, but also safer.
  * It may catch programming errors that would otherwise go undetected.
* Converting class objects
  * Conversions involving class objects (rather than pointers to objects) are possible `only` if the programmer has defined a `constructor`, an `overloaded` assignment operator, or an `over-loaded type casting operator` that specifies how to do the conversion.
  * The constructor or overloaded operator is as efficient as a member function

## Branches and switch statements

* The high speed of modern microprocessors is obtained by using a pipeline where instructions are fetched and decoded in several stages before they are executed.
  * the pipeline structure has one big problem:
    * 當 Code 有分支(`if`, `else`)時，無法預先判斷該將哪個分支放入 `pipeline` 可能會造成多花 `10 - 20` clock cycles 執行 fetching, decoding
    * A branch instruction takes typically `0 - 2` clock cycles in the case that the microprocessor has made the right prediction.
    * The time it takes to recover from a branch misprediction is approximately `12 - 25` clock cycles, depending on the processor. This is called the branch misprediction penalty
  * A branch that follows a `simple periodic pattern` can also be predicted quite well if it is inside a loop with few or no other branches.
  * A `for-loop` or `while-loop` is also a kind of branch.
    *  After each iteration it decides whether to repeat or to exit the loop.
    * The maximum loop count that can be predicted perfectly varies `between 9 and 64`, depending on the processor
    * Nested loops are predicted well only on some processors
* A `switch` statements is a kind of branch that can go `more than two ways`.
