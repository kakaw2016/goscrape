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
sys.stdout = open("/home/kakashinaruto/go/src/github.com/kakaw2016/goscrape/ArtCurationHive/Artposts", "a+")


title = 'Please choose filter: '
#filters list
options = []
option1 = []
option2 = []
option22 = []
option3 = []
option4 = []
option5 = []
option6 = []
option7 = []
option8 = []


q = Query(limit=100, tag="art")
d = Discussions(blockchain_instance= h)
post1 = d.get_discussions('created', q, limit=100)

q2 = Query(limit=100, tag="digitalart")
d2 = Discussions(blockchain_instance= h)
post2 = d2.get_discussions('created', q2, limit=100)

q22 = Query(limit=100, tag="music")
d22 = Discussions(blockchain_instance= h)
post22 = d22.get_discussions('created', q22, limit=100)

q3 = Query(limit=100, tag="photography")
d3 = Discussions(blockchain_instance= h)
post3 = d3.get_discussions('created', q3, limit=100)

q4 = Query(limit=100, tag="photo")
d4 = Discussions(blockchain_instance= h)
post4 = d4.get_discussions('created', q4, limit=100)

q5 = Query(limit=100, tag="contest")
d5 = Discussions(blockchain_instance= h)
post5 = d5.get_discussions('created', q5, limit=100)

q6 = Query(limit=100, tag="contests")
d6 = Discussions(blockchain_instance= h)
post6 = d6.get_discussions('created', q6, limit=100)

q7 = Query(limit=100, tag="alive")
d7 = Discussions(blockchain_instance= h)
post7 = d7.get_discussions('created', q7, limit=100)

q8 = Query(limit=100, tag="neoxian")
d8 = Discussions(blockchain_instance= h)
post8 = d8.get_discussions('created', q8, limit=100)

# print post list for selected filter
for p in post1:
    option1.append('https://hive.blog/@'+p["author"]+'/'+p["permlink"])
#pprint.pprint(option1)

for p in post2:
    option2.append('https://hive.blog/@'+p["author"]+'/'+p["permlink"])

for p in post22:
    option22.append('https://hive.blog/@'+p["author"]+'/'+p["permlink"])

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


options = (set(option2) | set(option1) | set(option3) | set(option4) | set(option22)) & (set(option5) | set(option6) | set(option7) | set(option8))

for link in options:
    print(link)

sys.stdout.close()
sys.stdout=stdoutOrigin