#!/usr/bin/python
import sys
import pprint
#from pick import pick
# initialize Hive class
from beem import Hive
from beem.discussions import Query, Discussions
from beem.nodelist import NodeList
nodelist = NodeList()
nodelist.update_nodes()
stdoutOrigin=sys.stdout 
#sys.stdout = open("/home/kakashinaruto/CodeZone/Hivecuration/ContestPosts", "w")
sys.stdout = open("/home/kakashinaruto/go/src/github.com/kakaw2016/goscrape/curationAshbagtt/output.txt", "a+")
h = Hive(node=nodelist.get_hive_nodes())

title = 'Please choose filter: '
#filters list
options = []
option1 = []
option2 = []
option3 = []


# get index and selected filter name
##option, index = pick(options, title)

#q = Query(limit=100, tag="contest")
q = Query(limit=100, tag="dynamichivers")
d = Discussions(blockchain_instance=h)
post1 = d.get_discussions('created', q, limit=100)

# print post list for selected filter
for p in post1:
    option1.append('https://hive.blog/@'+p["author"]+'/'+p["permlink"])
#pprint.pprint(option1)

q2 = Query(limit=100, tag="hive-190059")
d2 = Discussions(blockchain_instance=h)
post2 = d2.get_discussions('created', q2, limit=100)

for p in post2:
    option2.append('https://hive.blog/@'+p["author"]+'/'+p["permlink"])



q3 = Query(limit=100, tag="hive-13323")
d3 = Discussions(blockchain_instance=h)
post3 = d3.get_discussions('created', q3, limit=100)
for p in post3:
    option3.append('https://hive.blog/@'+p["author"]+'/'+p["permlink"])



options = (set(option1) | set(option2) | set(option3))

for link in options:
    print(link)

sys.stdout.close()
sys.stdout=stdoutOrigin