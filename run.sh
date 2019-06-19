#!/bin/bash
# sample
# ENVIRONMENT=default ./run.sh esService
#

cmd=`ls -l | grep $1_ | awk '{print $9}'`
echo `pwd`
exec ./${cmd}