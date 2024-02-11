module.exports = {
    apps: [
      {
        name: "grabLinkBlurtAfrica",
        script: "./grabLinkBlurtAfrica.js",
        interpreter: "/home/kakashinaruto/.bun/bin/bun",
        cron_restart: "* */11 * * *",
        autorestart: false,
        watch: false,
        max_memory_restart: '1G'

        
      }
    ]
   };