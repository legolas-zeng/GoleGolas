#一个words.txt,进行词频统计
#the day is sunny the the
#the sunny is is

#你的脚本应当输出（以词频降序排列）：
#
#the 4
#is 3
#sunny 2
#day 1
############ 通过 xargs 分词 ############
#cat words.txt |xargs -n1|sort -r |uniq -c
#      3 the
#      2 sunny
#      2 is

############ 通过awk 分词 ############

cat words.txt |
awk '{
    for(i=1;i<=NF;i++){
        count[$i]++
    }
} END {
    for(k in count){
        print k" "count[k]
    }
}' |
sort -rnk 2
