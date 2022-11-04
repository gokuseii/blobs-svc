module blobs-svc

go 1.19

require (
	github.com/alecthomas/kingpin v2.2.6+incompatible
	github.com/go-chi/chi v4.1.2+incompatible
	github.com/go-ozzo/ozzo-validation v3.6.0+incompatible
	github.com/google/jsonapi v0.0.0-20200226002910-c8283f632fb7
	github.com/rubenv/sql-migrate v1.2.0
	github.com/spf13/cast v1.4.1
	gitlab.com/distributed_lab/ape v1.7.1
	gitlab.com/distributed_lab/kit v1.11.1
	gitlab.com/distributed_lab/logan v3.8.1+incompatible
	gitlab.com/tokend/go v3.16.0+incompatible
)

require (
	github.com/Masterminds/squirrel v1.5.3 // indirect
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751 // indirect
	github.com/alecthomas/units v0.0.0-20211218093645-b94a6e3cc137 // indirect
	github.com/asaskevich/govalidator v0.0.0-20200108200545-475eaeb16496 // indirect
	github.com/certifi/gocertifi v0.0.0-20200211180108-c7c1fbc02894 // indirect
	github.com/cheekybits/genny v0.0.0-20170328200008-9127e812e1e9 // indirect
	github.com/cilium/ebpf v0.7.0 // indirect
	github.com/cosiner/argv v0.1.0 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/derekparker/trie v0.0.0-20200317170641-1fdf38b7b0e9 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/evalphobia/logrus_sentry v0.8.2 // indirect
	github.com/fatih/structs v1.1.0 // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/getsentry/raven-go v0.2.0 // indirect
	github.com/getsentry/sentry-go v0.7.0 // indirect
	github.com/go-delve/delve v1.9.1 // indirect
	github.com/go-delve/liner v1.2.3-0.20220127212407-d32d89dd2a5d // indirect
	github.com/go-errors/errors v1.0.1 // indirect
	github.com/go-gorp/gorp/v3 v3.0.2 // indirect
	github.com/go-ozzo/ozzo-validation/v4 v4.2.1 // indirect
	github.com/goji/context v0.0.0-20160122015720-68b83f7b0439 // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/google/go-dap v0.6.0 // indirect
	github.com/guregu/null v3.2.0+incompatible // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/jmoiron/sqlx v1.2.0 // indirect
	github.com/lann/builder v0.0.0-20180802200727-47ae307949d0 // indirect
	github.com/lann/ps v0.0.0-20150810152359-62de8c46ede0 // indirect
	github.com/lib/pq v1.10.0 // indirect
	github.com/magiconair/properties v1.8.5 // indirect
	github.com/mattn/go-colorable v0.1.2 // indirect
	github.com/mattn/go-isatty v0.0.9 // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/mitchellh/mapstructure v1.4.1 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/nullstyle/go-xdr v0.0.0-20180726165426-f4c839f75077 // indirect
	github.com/oklog/ulid v1.3.1 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/pelletier/go-toml v1.9.3 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/rs/cors v1.7.0 // indirect
	github.com/russross/blackfriday/v2 v2.0.1 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/spf13/afero v1.6.0 // indirect
	github.com/spf13/cobra v1.2.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.8.1 // indirect
	github.com/stretchr/objx v0.4.0 // indirect
	github.com/stretchr/testify v1.8.0 // indirect
	github.com/subosito/gotenv v1.2.0 // indirect
	github.com/throttled/throttled/v2 v2.7.2 // indirect
	github.com/zenazn/goji v0.9.1-0.20160507202103-64eb34159fe5 // indirect
	gitlab.com/distributed_lab/corer v0.0.0-20171130114150-cbfb46525895 // indirect
	gitlab.com/distributed_lab/figure v2.1.0+incompatible // indirect
	gitlab.com/distributed_lab/json-api-connector v0.2.4 // indirect
	gitlab.com/distributed_lab/lorem v0.2.0 // indirect
	gitlab.com/distributed_lab/running v1.6.0 // indirect
	gitlab.com/distributed_lab/txsub v0.0.0-20171130120140-d7781cbc2319 // indirect
	gitlab.com/distributed_lab/urlval v3.0.0+incompatible // indirect
	gitlab.com/tokend/connectors v0.1.9-rc.4 // indirect
	gitlab.com/tokend/horizon v0.0.0-20211130122149-3181416742eb // indirect
	gitlab.com/tokend/horizon-connector v4.3.1+incompatible // indirect
	gitlab.com/tokend/keypair v0.0.0-20190412110653-b9d7e0c8b312 // indirect
	gitlab.com/tokend/regources v4.9.2-0.20211101124223-04e3b08bddce+incompatible // indirect
	go.starlark.net v0.0.0-20220816155156-cfacd8902214 // indirect
	golang.org/x/arch v0.0.0-20190927153633-4e8777c89be4 // indirect
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
	golang.org/x/net v0.0.0-20220722155237-a158d28d115b // indirect
	golang.org/x/sys v0.0.0-20220811171246-fbc7d0a398ab // indirect
	golang.org/x/text v0.3.7 // indirect
	gopkg.in/ini.v1 v1.62.0 // indirect
	gopkg.in/tylerb/graceful.v1 v1.2.15 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
