
* ldd

```sh
ldd NAME
```

* `LD_PRELOAD=/PATH/TO/LIB`


```sh
LD_PRELOAD=/PATH/TO/LIB ./NAME
```

* wait child thread

```cpp
// signal wait
#include <signal.h>
sigset_t set;
int sig;
sigwait(&set, &sig);

// wait
#include <sys/wait.h>
wait(NULL);

// pause
pause();
```

* top

顯示某個程式 thread 所在 cpu

```sh
top -p <PID> -H

# Fields Management
f

# open: P       = Last Used Cpu (SMP)
<space>

# go back
q
```
