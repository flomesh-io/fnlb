#ifndef __LOXILB_DP_H__
#define __LOXILB_DP_H__

#define LOXILB_DP_LOGF           "/var/log/fsmxlb.log"

#ifndef XDP_LL_SEC_DEFAULT
#define XDP_LL_SEC_DEFAULT       "xdp_packet_hook"
#endif

#ifndef TC_LL_SEC_DEFAULT
#define TC_LL_SEC_DEFAULT        "tc_packet_hook0"

enum llb_bpf_mnt_type {
  LL_BPF_MOUNT_NONE = 0,
  LL_BPF_MOUNT_XDP,
  LL_BPF_MOUNT_TC
};

#endif

#include <stdio.h>
#include <stdlib.h>
#include <stddef.h>
#include <stdbool.h>
#include <string.h>
#include <stdint.h>
#include <unistd.h>
#include <errno.h>
#include <assert.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <sys/mman.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <sys/ioctl.h>
#include <net/if.h>
#include <linux/bpf.h>
#include <pthread.h>

#include "../common/common_params.h"
#include "../common/common_user_bpf_xdp.h"
#include "../common/common_libbpf.h"
#include "../common/common_pdi.h"
#include "../common/llb_dpapi.h"
#include "../common/llb_dp_mdi.h"
#include "../common/log.h"

unsigned long long get_os_usecs(void);
unsigned long long get_os_nsecs(void);

struct ebpfcfg {
  int have_mtrace;
  int have_ptrace;
  int no_loader;
  int nodenum;
  int loglevel;
};

void loxilb_set_loglevel(struct ebpfcfg *cfg);
int loxilb_main(struct ebpfcfg *cfg);


#endif
