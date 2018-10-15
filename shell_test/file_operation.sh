dir=.
for i in `ls ${dir}`
do
 if[[${i}-eq"a"]]
 then
   echo "${i}"
 fi
done
