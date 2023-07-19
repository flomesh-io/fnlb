#!/bin/bash
sudo umount /opt/fsmxlb/dp/ >> /dev/null 2>&1
sudo rm -fr /opt/fsmxlb/dp/bpf >> /dev/null 2>&1
sudo mkdir -p /opt/fsmxlb/dp/ >> /dev/null 2>&1
sudo mount -t bpf bpf /opt/fsmxlb/dp/ >> /dev/null 2>&1
