
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
