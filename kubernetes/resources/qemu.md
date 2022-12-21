**QEMU**

* Kernel based virtual machines: (KVM)
1. It is the fantastic hypervisor baked into the linux kernel. So, hypervisors let you run complete isolated OS or virtual machines inside our actual computer. The Vm's are useful for backup and restore, we can take a snapshot of VM
2. KVM is an open source, baked into linux and free to use.
3. Emulated environments require a software bridge to interact with the hardware, virtualization accesses hardware directly. QEMU is emulated and virtualized software while virtualbox is just the virtualization tool.
4. QEMU supports a wide range of hardware and can make use of the KVM when running a target architecture which is the same as the host architecture.

* Resources : 

1. https://linuxconfig.org/qemu-vs-virtualbox-whats-the-difference
2. https://www.youtube.com/watch?app=desktop&v=Kq849CpGd88
3. https://christitus.com/vm-setup-in-linux/ - Installation guide
