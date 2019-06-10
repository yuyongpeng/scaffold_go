module scaffold_go

go 1.12

require (
	github.com/elastic/go-elasticsearch/v8 v8.0.0-20190524105426-8ed748708d21
	github.com/fsnotify/fsnotify v1.4.7
	github.com/go-ini/ini v1.42.0
	github.com/go-sql-driver/mysql v1.4.1
	github.com/hashicorp/hcl v1.0.0
	github.com/inconshreveable/mousetrap v1.0.0
	github.com/jinzhu/gorm v1.9.8
	github.com/jinzhu/inflection v0.0.0-20190603042836-f5c5f50e6090
	github.com/konsorten/go-windows-terminal-sequences v1.0.2
	github.com/magiconair/properties v1.8.1
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mitchellh/mapstructure v1.1.2
	github.com/pelletier/go-toml v1.4.0
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/afero v1.2.2
	github.com/spf13/cast v1.3.0
	github.com/spf13/cobra v0.0.5
	github.com/spf13/jwalterweatherman v1.1.0
	github.com/spf13/pflag v1.0.3
	github.com/spf13/viper v1.4.0
	github.com/urfave/cli v1.20.0
	golang.org/x/sys v0.0.0-20190608050228-5b15430b70e3
	golang.org/x/text v0.3.2
	google.golang.org/appengine v1.6.1
	gopkg.in/yaml.v2 v2.2.2
)

replace (
	github.com/golang/sys => github.com/golang/sys v0.0.0-20190610081024-1e42afee0f76
	golang.org/x/text => github.com/golang/text v0.3.0
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190608022120-eacb66d2a7c3
)
