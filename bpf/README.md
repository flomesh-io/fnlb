## This README is here for anyone who wants to build fsmxlb ebpf only modules

## Install Dependencies

sudo apt install clang llvm libelf-dev gcc-multilib libpcap-dev  
sudo apt install linux-tools-$(uname -r)  
sudo apt install elfutils dwarves  

## Build libbpf

cd libbpf/src  
sudo make install  
sudo ldconfig  

## Build fsmxlb ebpf

cd -   
make  
