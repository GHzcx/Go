#!/bin/sh
count=`ps -ef |grep "Task_match_photographer"|grep -v "color" | grep -v grep|wc -l`;
`source /etc/profile`
echo $count
if [ $count -eq 0 ]
then
    `nohup php /opt/web/phpc/application/hbg_php_gongyu_script/phpapps/gongyu/public/script.php "Vr\\Task_match_photographer" >>/dev/null &`
fi