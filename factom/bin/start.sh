nohup /root/bin/factomd -broadcastnum=16 -faulttimeout=120 -startdelay=600 -network=CUSTOM -customnet=fct_community_test -blktime=600 -debugconsole=remotehost:8093 > runlog.txt 2>&1 &
