## How to build?

### Install depends

```bash
make depends
echo export PATH=\${PATH}:/usr/local/go/bin >> ~/.bashrc
reboot
```

### Build CLI

```bash
make subsys
make build
```

## How to test ?

### Run XLB Server

```bash
make run
```

### Run Demo

```bash
make -f Makefile.standalone pipy
make -f Makefile.standalone test-up
make -f Makefile.standalone test-apply-lb
```

### Test

```
make -f Makefile.standalone test
```



### Stop Demo and XLB Server

```bash
make -f Makefile.standalone test-down
make stop
```

