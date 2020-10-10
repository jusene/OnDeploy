package templates

var GITConf = `
<VirtualHost *:{{ .GITPort }}>
   ServerName _
   SetEnv GIT_PROJECT_ROOT /ddhome/local/gitdata
   SetEnv GIT_HTTP_EXPORT_ALL
   ScriptAlias {{ .GITPath }} /usr/libexec/git-core/git-http-backend/
   <Location />
         AuthType Basic
         AuthName "Git"
         AuthUserFile /etc/httpd/conf/.httpd
         Require valid-user
   </Location>
   <Directory "/usr/libexec/git-core">
      Options ExecCGI Indexes
      Require all granted
   </Directory>
</VirtualHost>
`
