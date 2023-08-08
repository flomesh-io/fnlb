## How to build?

```bash
make depends

make subsys

make build-cli
```

## How to test ?

```bash
make run
make -f Makefile.standalone test-up
make -f Makefile.standalone test-apply-lb
make -f Makefile.standalone test-down
make stop
```

