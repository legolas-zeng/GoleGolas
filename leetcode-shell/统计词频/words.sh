#一个words.txt,进行词频统计
#the day is sunny the the
#the sunny is is

#你的脚本应当输出（以词频降序排列）：
#
#the 4
#is 3
#sunny 2
#day 1

cat words.txt |xargs -n1|sort -r |uniq -c
      3 the
      2 sunny
      2 is