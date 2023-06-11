#!/usr/bin/python
import sys
import pprint
#from pick import pick
# initialize Hive class
from beem import Hive
from beem.discussions import Query, Discussions
stdoutOrigin=sys.stdout 
#sys.stdout = open("/home/youthbrigthfuture/CodeZone/Hivecuration/ContestPosts", "w")
sys.stdout = open("/home/youthbrigthfuture/go/src/github.com/kakaw2016/goscrape/SplinGiveawayScrape/SplinterPosts", "a+")
h = Hive()

title = 'Please choose filter: '
#filters list
options = []
option1 = []
option2 = []
option3 = []
option4 = []
option5 = []
option6 = []
option7 = []
option8 = []

# get index and selected filter name
##option, index = pick(options, title)

#q = Query(limit=100, tag="contest")
q = Query(limit=100, tag="splinterlands")
d = Discussions()
post1 = d.get_discussions('created', q, limit=100)

q2 = Query(limit=100, tag="giveaway")
d2 = Discussions()
post2 = d2.get_discussions('created', q2, limit=100)

q3 = Query(limit=100, tag="leofinance")
d3 = Discussions()
post3 = d3.get_discussions('created', q3, limit=100)

q4 = Query(limit=100, tag="thgaming")
d4 = Discussions()
post4 = d4.get_discussions('created', q4, limit=100)

q5 = Query(limit=100, tag="giveaways")
d5 = Discussions()
post5 = d5.get_discussions('created', q5, limit=100)

q6 = Query(limit=100, tag="spt")
d6 = Discussions()
post6 = d6.get_discussions('created', q6, limit=100)

q7 = Query(limit=100, tag="pgm")
d7 = Discussions()
post7 = d7.get_discussions('created', q7, limit=100)

q8 = Query(limit=100, tag="neoxian")
d8 = Discussions()
post8 = d8.get_discussions('created', q8, limit=100)

# print post list for selected filter
for p in post1:
    option1.append('https://hive.blog/@'+p["author"]+'/'+p["permlink"])
#pprint.pprint(option1)

for p in post2:
    option2.append('https://hive.blog/@'+p["author"]+'/'+p["permlink"])

for p in post3:
    option3.append('https://hive.blog/@'+p["author"]+'/'+p["permlink"])

for p in post4:
    option4.append('https://hive.blog/@'+p["author"]+'/'+p["permlink"])

for p in post5:
    option5.append('https://hive.blog/@'+p["author"]+'/'+p["permlink"])

for p in post6:
    option6.append('https://hive.blog/@'+p["author"]+'/'+p["permlink"])

for p in post7:
    option7.append('https://hive.blog/@'+p["author"]+'/'+p["permlink"])

for p in post8:
    option8.append('https://hive.blog/@'+p["author"]+'/'+p["permlink"])


options = (set(option2) | set(option5) | set(option6) | set(option7)) & (set(option3) | set(option4) | set(option1) | set(option8))

for link in options:
    print(link)

sys.stdout.close()
sys.stdout=stdoutOrigin