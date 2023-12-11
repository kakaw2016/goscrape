module.exports = {
    apps: [
      {
        name: 'grabLinkBlurtConnect',
        script: './grabLinkBlurtConnect.js',
        interpreter: "/home/kakashinaruto/.bun/bin/bun",
        cron_restart: '15 * * * *',
      }
    ]
   };