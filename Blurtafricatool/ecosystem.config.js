module.exports = {
    apps: [
      {
        name: 'grabLinkBlurtAfrica',
        script: './grabLinkBlurtAfrica.js',
        interpreter: "/home/kakashinaruto/.bun/bin/bun",
        cron_restart: '15 * * * *',
      }
    ]
   };