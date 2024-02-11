#!/usr/bin/python

from beem import Hive
from beem.nodelist import NodeList
nodelist = NodeList()
nodelist.update_nodes()
nodes = nodelist.get_hive_nodes()
h = Hive(node=nodes)

from beem.discussions import Query, Discussions
import sys
import pprint

stdoutOrigin=sys.stdout 
#sys.stdout = open("/home/kakashinaruto/CodeZone/Hivecuration/ContestPosts", "w")
sys.stdout = open("/home/kakashinaruto/go/src/github.com/kakaw2016/goscrape/SplinGiveawayScrape/SplinterPosts", "a+")


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
d = Discussions(blockchain_instance= h)
post1 = d.get_discussions('created', q, limit=100)

q2 = Query(limit=100, tag="giveaway")
d2 = Discussions(blockchain_instance= h)
post2 = d2.get_discussions('created', q2, limit=100)

q3 = Query(limit=100, tag="leofinance")
d3 = Discussions(blockchain_instance= h)
post3 = d3.get_discussions('created', q3, limit=100)

q4 = Query(limit=100, tag="thgaming")
d4 = Discussions(blockchain_instance= h)
post4 = d4.get_discussions('created', q4, limit=100)

q5 = Query(limit=100, tag="giveaways")
d5 = Discussions(blockchain_instance= h)
post5 = d5.get_discussions('created', q5, limit=100)

q6 = Query(limit=100, tag="spt")
d6 = Discussions(blockchain_instance= h)
post6 = d6.get_discussions('created', q6, limit=100)

q7 = Query(limit=100, tag="pgm")
d7 = Discussions(blockchain_instance= h)
post7 = d7.get_discussions('created', q7, limit=100)

q8 = Query(limit=100, tag="neoxian")
d8 = Discussions(blockchain_instance= h)
post8 = d8.get_discussions('created', q8, limit=100)

# print post list for selected filter
for p1 in post1:
    option1.append('https://hive.blog/@'+p1["author"]+'/'+p1["permlink"])
#pprint.pprint(option1)

for p2 in post2:
    option2.append('https://hive.blog/@'+p2["author"]+'/'+p2["permlink"])

for p3 in post3:
    option3.append('https://hive.blog/@'+p3["author"]+'/'+p3["permlink"])

for p4 in post4:
    option4.append('https://hive.blog/@'+p4["author"]+'/'+p4["permlink"])

for p5 in post5:
    option5.append('https://hive.blog/@'+p5["author"]+'/'+p5["permlink"])

for p6 in post6:
    option6.append('https://hive.blog/@'+p6["author"]+'/'+p6["permlink"])

for p7 in post7:
    option7.append('https://hive.blog/@'+p7["author"]+'/'+p7["permlink"])

for p8 in post8:
    option8.append('https://hive.blog/@'+p8["author"]+'/'+p8["permlink"])


options = (set(option2) | set(option5) | set(option6) | set(option7)) & (set(option3) | set(option4) | set(option1) | set(option8))

for link in options:
    print(link)

sys.stdout.close()
sys.stdout=stdoutOrigin