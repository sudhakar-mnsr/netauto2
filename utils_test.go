package cgroups

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

const fedoraMountinfo = `15 35 0:3 / /proc rw,nosuid,nodev,noexec,relatime shared:5 - proc proc rw
16 35 0:14 / /sys rw,nosuid,nodev,noexec,relatime shared:6 - sysfs sysfs rw,seclabel
17 35 0:5 / /dev rw,nosuid shared:2 - devtmpfs devtmpfs rw,seclabel,size=8056484k,nr_inodes=2014121,mode=755
18 16 0:15 / /sys/kernel/security rw,nosuid,nodev,noexec,relatime shared:7 - securityfs securityfs rw
19 16 0:13 / /sys/fs/selinux rw,relatime shared:8 - selinuxfs selinuxfs rw
20 17 0:16 / /dev/shm rw,nosuid,nodev shared:3 - tmpfs tmpfs rw,seclabel
21 17 0:10 / /dev/pts rw,nosuid,noexec,relatime shared:4 - devpts devpts rw,seclabel,gid=5,mode=620,ptmxmode=000
22 35 0:17 / /run rw,nosuid,nodev shared:21 - tmpfs tmpfs rw,seclabel,mode=755
23 16 0:18 / /sys/fs/cgroup rw,nosuid,nodev,noexec shared:9 - tmpfs tmpfs rw,seclabel,mode=755
24 23 0:19 / /sys/fs/cgroup/systemd rw,nosuid,nodev,noexec,relatime shared:10 - cgroup cgroup rw,xattr,release_agent=/usr/lib/systemd/systemd-cgroups-agent,name=systemd
25 16 0:20 / /sys/fs/pstore rw,nosuid,nodev,noexec,relatime shared:20 - pstore pstore rw
26 23 0:21 / /sys/fs/cgroup/cpuset rw,nosuid,nodev,noexec,relatime shared:11 - cgroup cgroup rw,cpuset,clone_children
27 23 0:22 / /sys/fs/cgroup/cpu,cpuacct rw,nosuid,nodev,noexec,relatime shared:12 - cgroup cgroup rw,cpuacct,cpu,clone_children
28 23 0:23 / /sys/fs/cgroup/memory rw,nosuid,nodev,noexec,relatime shared:13 - cgroup cgroup rw,memory,clone_children
29 23 0:24 / /sys/fs/cgroup/devices rw,nosuid,nodev,noexec,relatime shared:14 - cgroup cgroup rw,devices,clone_children
30 23 0:25 / /sys/fs/cgroup/freezer rw,nosuid,nodev,noexec,relatime shared:15 - cgroup cgroup rw,freezer,clone_children
31 23 0:26 / /sys/fs/cgroup/net_cls rw,nosuid,nodev,noexec,relatime shared:16 - cgroup cgroup rw,net_cls,clone_children
32 23 0:27 / /sys/fs/cgroup/blkio rw,nosuid,nodev,noexec,relatime shared:17 - cgroup cgroup rw,blkio,clone_children
33 23 0:28 / /sys/fs/cgroup/perf_event rw,nosuid,nodev,noexec,relatime shared:18 - cgroup cgroup rw,perf_event,clone_children
34 23 0:29 / /sys/fs/cgroup/hugetlb rw,nosuid,nodev,noexec,relatime shared:19 - cgroup cgroup rw,hugetlb,clone_children
35 1 253:2 / / rw,relatime shared:1 - ext4 /dev/mapper/ssd-root--f20 rw,seclabel,data=ordered
36 15 0:30 / /proc/sys/fs/binfmt_misc rw,relatime shared:22 - autofs systemd-1 rw,fd=38,pgrp=1,timeout=300,minproto=5,maxproto=5,direct
37 17 0:12 / /dev/mqueue rw,relatime shared:23 - mqueue mqueue rw,seclabel
38 35 0:31 / /tmp rw shared:24 - tmpfs tmpfs rw,seclabel
39 17 0:32 / /dev/hugepages rw,relatime shared:25 - hugetlbfs hugetlbfs rw,seclabel
40 16 0:7 / /sys/kernel/debug rw,relatime shared:26 - debugfs debugfs rw
41 16 0:33 / /sys/kernel/config rw,relatime shared:27 - configfs configfs rw
42 35 0:34 / /var/lib/nfs/rpc_pipefs rw,relatime shared:28 - rpc_pipefs sunrpc rw
43 15 0:35 / /proc/fs/nfsd rw,relatime shared:29 - nfsd sunrpc rw
45 35 8:17 / /boot rw,relatime shared:30 - ext4 /dev/sdb1 rw,seclabel,data=ordered
46 35 253:4 / /home rw,relatime shared:31 - ext4 /dev/mapper/ssd-home rw,seclabel,data=ordered
47 35 253:5 / /var/lib/libvirt/images rw,noatime,nodiratime shared:32 - ext4 /dev/mapper/ssd-virt rw,seclabel,discard,data=ordered
48 35 253:12 / /mnt/old rw,relatime shared:33 - ext4 /dev/mapper/HelpDeskRHEL6-FedoraRoot rw,seclabel,data=ordered
121 22 0:36 / /run/user/1000/gvfs rw,nosuid,nodev,relatime shared:104 - fuse.gvfsd-fuse gvfsd-fuse rw,user_id=1000,group_id=1000
124 16 0:37 / /sys/fs/fuse/connections rw,relatime shared:107 - fusectl fusectl rw
165 38 253:3 / /tmp/mnt rw,relatime shared:147 - ext4 /dev/mapper/ssd-root rw,seclabel,data=ordered
167 35 253:15 / /var/lib/docker/devicemapper/mnt/aae4076022f0e2b80a2afbf8fc6df450c52080191fcef7fb679a73e6f073e5c2 rw,relatime shared:149 - ext4 /dev/mapper/docker-253:2-425882-aae4076022f0e2b80a2afbf8fc6df450c52080191fcef7fb679a73e6f073e5c2 rw,seclabel,discard,stripe=16,data=ordered
171 35 253:16 / /var/lib/docker/devicemapper/mnt/c71be651f114db95180e472f7871b74fa597ee70a58ccc35cb87139ddea15373 rw,relatime shared:153 - ext4 /dev/mapper/docker-253:2-425882-c71be651f114db95180e472f7871b74fa597ee70a58ccc35cb87139ddea15373 rw,seclabel,discard,stripe=16,data=ordered
175 35 253:17 / /var/lib/docker/devicemapper/mnt/1bac6ab72862d2d5626560df6197cf12036b82e258c53d981fa29adce6f06c3c rw,relatime shared:157 - ext4 /dev/mapper/docker-253:2-425882-1bac6ab72862d2d5626560df6197cf12036b82e258c53d981fa29adce6f06c3c rw,seclabel,discard,stripe=16,data=ordered
179 35 253:18 / /var/lib/docker/devicemapper/mnt/d710a357d77158e80d5b2c55710ae07c94e76d34d21ee7bae65ce5418f739b09 rw,relatime shared:161 - ext4 /dev/mapper/docker-253:2-425882-d710a357d77158e80d5b2c55710ae07c94e76d34d21ee7bae65ce5418f739b09 rw,seclabel,discard,stripe=16,data=ordered
183 35 253:19 / /var/lib/docker/devicemapper/mnt/6479f52366114d5f518db6837254baab48fab39f2ac38d5099250e9a6ceae6c7 rw,relatime shared:165 - ext4 /dev/mapper/docker-253:2-425882-6479f52366114d5f518db6837254baab48fab39f2ac38d5099250e9a6ceae6c7 rw,seclabel,discard,stripe=16,data=ordered
187 35 253:20 / /var/lib/docker/devicemapper/mnt/8d9df91c4cca5aef49eeb2725292aab324646f723a7feab56be34c2ad08268e1 rw,relatime shared:169 - ext4 /dev/mapper/docker-253:2-425882-8d9df91c4cca5aef49eeb2725292aab324646f723a7feab56be34c2ad08268e1 rw,seclabel,discard,stripe=16,data=ordered
191 35 253:21 / /var/lib/docker/devicemapper/mnt/c8240b768603d32e920d365dc9d1dc2a6af46cd23e7ae819947f969e1b4ec661 rw,relatime shared:173 - ext4 /dev/mapper/docker-253:2-425882-c8240b768603d32e920d365dc9d1dc2a6af46cd23e7ae819947f969e1b4ec661 rw,seclabel,discard,stripe=16,data=ordered
195 35 253:22 / /var/lib/docker/devicemapper/mnt/2eb3a01278380bbf3ed12d86ac629eaa70a4351301ee307a5cabe7b5f3b1615f rw,relatime shared:177 - ext4 /dev/mapper/docker-253:2-425882-2eb3a01278380bbf3ed12d86ac629eaa70a4351301ee307a5cabe7b5f3b1615f rw,seclabel,discard,stripe=16,data=ordered
199 35 253:23 / /var/lib/docker/devicemapper/mnt/37a17fb7c9d9b80821235d5f2662879bd3483915f245f9b49cdaa0e38779b70b rw,relatime shared:181 - ext4 /dev/mapper/docker-253:2-425882-37a17fb7c9d9b80821235d5f2662879bd3483915f245f9b49cdaa0e38779b70b rw,seclabel,discard,stripe=16,data=ordered
203 35 253:24 / /var/lib/docker/devicemapper/mnt/aea459ae930bf1de913e2f29428fd80ee678a1e962d4080019d9f9774331ee2b rw,relatime shared:185 - ext4 /dev/mapper/docker-253:2-425882-aea459ae930bf1de913e2f29428fd80ee678a1e962d4080019d9f9774331ee2b rw,seclabel,discard,stripe=16,data=ordered
207 35 253:25 / /var/lib/docker/devicemapper/mnt/928ead0bc06c454bd9f269e8585aeae0a6bd697f46dc8754c2a91309bc810882 rw,relatime shared:189 - ext4 /dev/mapper/docker-253:2-425882-928ead0bc06c454bd9f269e8585aeae0a6bd697f46dc8754c2a91309bc810882 rw,seclabel,discard,stripe=16,data=ordered
211 35 253:26 / /var/lib/docker/devicemapper/mnt/0f284d18481d671644706e7a7244cbcf63d590d634cc882cb8721821929d0420 rw,relatime shared:193 - ext4 /dev/mapper/docker-253:2-425882-0f284d18481d671644706e7a7244cbcf63d590d634cc882cb8721821929d0420 rw,seclabel,discard,stripe=16,data=ordered
215 35 253:27 / /var/lib/docker/devicemapper/mnt/d9dd16722ab34c38db2733e23f69e8f4803ce59658250dd63e98adff95d04919 rw,relatime shared:197 - ext4 /dev/mapper/docker-253:2-425882-d9dd16722ab34c38db2733e23f69e8f4803ce59658250dd63e98adff95d04919 rw,seclabel,discard,stripe=16,data=ordered
219 35 253:28 / /var/lib/docker/devicemapper/mnt/bc4500479f18c2c08c21ad5282e5f826a016a386177d9874c2764751c031d634 rw,relatime shared:201 - ext4 /dev/mapper/docker-253:2-425882-bc4500479f18c2c08c21ad5282e5f826a016a386177d9874c2764751c031d634 rw,seclabel,discard,stripe=16,data=ordered
223 35 253:29 / /var/lib/docker/devicemapper/mnt/7770c8b24eb3d5cc159a065910076938910d307ab2f5d94e1dc3b24c06ee2c8a rw,relatime shared:205 - ext4 /dev/mapper/docker-253:2-425882-7770c8b24eb3d5cc159a065910076938910d307ab2f5d94e1dc3b24c06ee2c8a rw,seclabel,discard,stripe=16,data=ordered
227 35 253:30 / /var/lib/docker/devicemapper/mnt/c280cd3d0bf0aa36b478b292279671624cceafc1a67eaa920fa1082601297adf rw,relatime shared:209 - ext4 /dev/mapper/docker-253:2-425882-c280cd3d0bf0aa36b478b292279671624cceafc1a67eaa920fa1082601297adf rw,seclabel,discard,stripe=16,data=ordered
231 35 253:31 / /var/lib/docker/devicemapper/mnt/8b59a7d9340279f09fea67fd6ad89ddef711e9e7050eb647984f8b5ef006335f rw,relatime shared:213 - ext4 /dev/mapper/docker-253:2-425882-8b59a7d9340279f09fea67fd6ad89ddef711e9e7050eb647984f8b5ef006335f rw,seclabel,discard,stripe=16,data=ordered
235 35 253:32 / /var/lib/docker/devicemapper/mnt/1a28059f29eda821578b1bb27a60cc71f76f846a551abefabce6efd0146dce9f rw,relatime shared:217 - ext4 /dev/mapper/docker-253:2-425882-1a28059f29eda821578b1bb27a60cc71f76f846a551abefabce6efd0146dce9f rw,seclabel,discard,stripe=16,data=ordered
239 35 253:33 / /var/lib/docker/devicemapper/mnt/e9aa60c60128cad1 rw,relatime shared:221 - ext4 /dev/mapper/docker-253:2-425882-e9aa60c60128cad1 rw,seclabel,discard,stripe=16,data=ordered
243 35 253:34 / /var/lib/docker/devicemapper/mnt/5fec11304b6f4713fea7b6ccdcc1adc0a1966187f590fe25a8227428a8df275d-init rw,relatime shared:225 - ext4 /dev/mapper/docker-253:2-425882-5fec11304b6f4713fea7b6ccdcc1adc0a1966187f590fe25a8227428a8df275d-init rw,seclabel,discard,stripe=16,data=ordered
247 35 253:35 / /var/lib/docker/devicemapper/mnt/5fec11304b6f4713fea7b6ccdcc1adc0a1966187f590fe25a8227428a8df275d rw,relatime shared:229 - ext4 /dev/mapper/docker-253:2-425882-5fec11304b6f4713fea7b6ccdcc1adc0a1966187f590fe25a8227428a8df275d rw,seclabel,discard,stripe=16,data=ordered
31 21 0:23 / /DATA/foo_bla_bla rw,relatime - cifs //foo/BLA\040BLA\040BLA/ rw,sec=ntlm,cache=loose,unc=\\foo\BLA BLA BLA,username=my_login,domain=mydomain.com,uid=12345678,forceuid,gid=12345678,forcegid,addr=10.1.30.10,file_mode=0755,dir_mode=0755,nounix,rsize=61440,wsize=65536,actimeo=1`

