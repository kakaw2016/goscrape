module.exports = {
    apps: [
      {
        name: "grabLinkBlurtConnect",
        script: "./grabLinkBlurtConnect.js",
        interpreter: "/home/kakashinaruto/.bun/bin/bun",
        cron_restart: "* */11 * * *",
        autorestart: false,
        watch: false,
        max_memory_restart: '1G'
      }
    ]
  };