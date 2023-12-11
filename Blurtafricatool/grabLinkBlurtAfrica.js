const dblurt = require('@beblurt/dblurt');
const fs = require('fs');

var client = new dblurt.Client(
  [
    'https://rpc.beblurt.com',
    'https://rpc.blurt.world',
    'https://blurt-rpc.saboin.com',
    'https://rpc.blurt.live',
  ],
  { timeout: 1500 }
);
 
client.condenser
 .getDiscussions('created', { tag: 'blurtafrica', limit: 100 })
 .then(function (discussions) {
  let data = '';
  for (let i = 0; i < discussions.length; i++) {
    data += discussions[i].url + '\n';
    }
  let newData = data.replace(/.*@/g, "https://blurt.blog/@");
  fs.appendFile('/home/kakashinaruto/go/src/github.com/kakaw2016/goscrape/Blurtafricatool/BlurtConnectLinkScrape.txt', newData, err => {
    if (err) {
      console.error(err);
    }  
    //console.log('Data has been written to file successfully.');
  });
});
