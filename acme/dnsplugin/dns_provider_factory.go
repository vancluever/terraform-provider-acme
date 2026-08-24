// Auto-generated file. Do not edit.
package dnsplugin

import (
	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/abion"
	"github.com/go-acme/lego/v5/providers/dns/acmedns"
	"github.com/go-acme/lego/v5/providers/dns/active24"
	"github.com/go-acme/lego/v5/providers/dns/alidns"
	"github.com/go-acme/lego/v5/providers/dns/aliesa"
	"github.com/go-acme/lego/v5/providers/dns/allinkl"
	"github.com/go-acme/lego/v5/providers/dns/alwaysdata"
	"github.com/go-acme/lego/v5/providers/dns/anexia"
	"github.com/go-acme/lego/v5/providers/dns/artfiles"
	"github.com/go-acme/lego/v5/providers/dns/arvancloud"
	"github.com/go-acme/lego/v5/providers/dns/auroradns"
	"github.com/go-acme/lego/v5/providers/dns/autodns"
	"github.com/go-acme/lego/v5/providers/dns/axelname"
	"github.com/go-acme/lego/v5/providers/dns/azion"
	"github.com/go-acme/lego/v5/providers/dns/azuredns"
	"github.com/go-acme/lego/v5/providers/dns/baiducloud"
	"github.com/go-acme/lego/v5/providers/dns/beget"
	"github.com/go-acme/lego/v5/providers/dns/binarylane"
	"github.com/go-acme/lego/v5/providers/dns/bindman"
	"github.com/go-acme/lego/v5/providers/dns/bluecat"
	"github.com/go-acme/lego/v5/providers/dns/bluecatv2"
	"github.com/go-acme/lego/v5/providers/dns/bookmyname"
	"github.com/go-acme/lego/v5/providers/dns/bunny"
	"github.com/go-acme/lego/v5/providers/dns/checkdomain"
	"github.com/go-acme/lego/v5/providers/dns/civo"
	"github.com/go-acme/lego/v5/providers/dns/clouddns"
	"github.com/go-acme/lego/v5/providers/dns/cloudflare"
	"github.com/go-acme/lego/v5/providers/dns/cloudns"
	"github.com/go-acme/lego/v5/providers/dns/cloudru"
	"github.com/go-acme/lego/v5/providers/dns/com35"
	"github.com/go-acme/lego/v5/providers/dns/connbyte"
	"github.com/go-acme/lego/v5/providers/dns/conoha"
	"github.com/go-acme/lego/v5/providers/dns/conohav3"
	"github.com/go-acme/lego/v5/providers/dns/constellix"
	"github.com/go-acme/lego/v5/providers/dns/corenetworks"
	"github.com/go-acme/lego/v5/providers/dns/cpanel"
	"github.com/go-acme/lego/v5/providers/dns/curanet"
	"github.com/go-acme/lego/v5/providers/dns/czechia"
	"github.com/go-acme/lego/v5/providers/dns/dandomain"
	"github.com/go-acme/lego/v5/providers/dns/ddnss"
	"github.com/go-acme/lego/v5/providers/dns/derak"
	"github.com/go-acme/lego/v5/providers/dns/desec"
	"github.com/go-acme/lego/v5/providers/dns/designate"
	"github.com/go-acme/lego/v5/providers/dns/digitalocean"
	"github.com/go-acme/lego/v5/providers/dns/dinahosting"
	"github.com/go-acme/lego/v5/providers/dns/directadmin"
	"github.com/go-acme/lego/v5/providers/dns/dns51"
	"github.com/go-acme/lego/v5/providers/dns/dnscale"
	"github.com/go-acme/lego/v5/providers/dns/dnsexit"
	"github.com/go-acme/lego/v5/providers/dns/dnshomede"
	"github.com/go-acme/lego/v5/providers/dns/dnsimple"
	"github.com/go-acme/lego/v5/providers/dns/dnsla"
	"github.com/go-acme/lego/v5/providers/dns/dnsmadeeasy"
	"github.com/go-acme/lego/v5/providers/dns/dnsservices"
	"github.com/go-acme/lego/v5/providers/dns/dnsupdate"
	"github.com/go-acme/lego/v5/providers/dns/dode"
	"github.com/go-acme/lego/v5/providers/dns/domeneshop"
	"github.com/go-acme/lego/v5/providers/dns/dreamhost"
	"github.com/go-acme/lego/v5/providers/dns/duckdns"
	"github.com/go-acme/lego/v5/providers/dns/dyn"
	"github.com/go-acme/lego/v5/providers/dns/dynadot"
	"github.com/go-acme/lego/v5/providers/dns/dyndnsfree"
	"github.com/go-acme/lego/v5/providers/dns/dynu"
	"github.com/go-acme/lego/v5/providers/dns/easydns"
	"github.com/go-acme/lego/v5/providers/dns/edgecenter"
	"github.com/go-acme/lego/v5/providers/dns/edgedns"
	"github.com/go-acme/lego/v5/providers/dns/edgeone"
	"github.com/go-acme/lego/v5/providers/dns/efficientip"
	"github.com/go-acme/lego/v5/providers/dns/epik"
	"github.com/go-acme/lego/v5/providers/dns/eurodns"
	"github.com/go-acme/lego/v5/providers/dns/euserv"
	"github.com/go-acme/lego/v5/providers/dns/excedo"
	"github.com/go-acme/lego/v5/providers/dns/exec"
	"github.com/go-acme/lego/v5/providers/dns/exoscale"
	"github.com/go-acme/lego/v5/providers/dns/f5xc"
	"github.com/go-acme/lego/v5/providers/dns/fornex"
	"github.com/go-acme/lego/v5/providers/dns/freemyip"
	"github.com/go-acme/lego/v5/providers/dns/gandi"
	"github.com/go-acme/lego/v5/providers/dns/gandiv5"
	"github.com/go-acme/lego/v5/providers/dns/gcloud"
	"github.com/go-acme/lego/v5/providers/dns/gcore"
	"github.com/go-acme/lego/v5/providers/dns/gehirn"
	"github.com/go-acme/lego/v5/providers/dns/gigahostno"
	"github.com/go-acme/lego/v5/providers/dns/glesys"
	"github.com/go-acme/lego/v5/providers/dns/gname"
	"github.com/go-acme/lego/v5/providers/dns/godaddy"
	"github.com/go-acme/lego/v5/providers/dns/gravity"
	"github.com/go-acme/lego/v5/providers/dns/hetzner"
	"github.com/go-acme/lego/v5/providers/dns/hostingde"
	"github.com/go-acme/lego/v5/providers/dns/hostinger"
	"github.com/go-acme/lego/v5/providers/dns/hostingnl"
	"github.com/go-acme/lego/v5/providers/dns/hosttech"
	"github.com/go-acme/lego/v5/providers/dns/hostup"
	"github.com/go-acme/lego/v5/providers/dns/httpnet"
	"github.com/go-acme/lego/v5/providers/dns/httpreq"
	"github.com/go-acme/lego/v5/providers/dns/huaweicloud"
	"github.com/go-acme/lego/v5/providers/dns/hurricane"
	"github.com/go-acme/lego/v5/providers/dns/hyperone"
	"github.com/go-acme/lego/v5/providers/dns/ibmcloud"
	"github.com/go-acme/lego/v5/providers/dns/iijdpf"
	"github.com/go-acme/lego/v5/providers/dns/infoblox"
	"github.com/go-acme/lego/v5/providers/dns/infomaniak"
	"github.com/go-acme/lego/v5/providers/dns/internetbs"
	"github.com/go-acme/lego/v5/providers/dns/inwx"
	"github.com/go-acme/lego/v5/providers/dns/ionos"
	"github.com/go-acme/lego/v5/providers/dns/ionoscloud"
	"github.com/go-acme/lego/v5/providers/dns/ipv64"
	"github.com/go-acme/lego/v5/providers/dns/ispconfig"
	"github.com/go-acme/lego/v5/providers/dns/ispconfigddns"
	"github.com/go-acme/lego/v5/providers/dns/jdcloud"
	"github.com/go-acme/lego/v5/providers/dns/joker"
	"github.com/go-acme/lego/v5/providers/dns/katapult"
	"github.com/go-acme/lego/v5/providers/dns/keyhelp"
	"github.com/go-acme/lego/v5/providers/dns/leaseweb"
	"github.com/go-acme/lego/v5/providers/dns/liara"
	"github.com/go-acme/lego/v5/providers/dns/lightsail"
	"github.com/go-acme/lego/v5/providers/dns/limacity"
	"github.com/go-acme/lego/v5/providers/dns/linode"
	"github.com/go-acme/lego/v5/providers/dns/liquidweb"
	"github.com/go-acme/lego/v5/providers/dns/loopia"
	"github.com/go-acme/lego/v5/providers/dns/luadns"
	"github.com/go-acme/lego/v5/providers/dns/mailinabox"
	"github.com/go-acme/lego/v5/providers/dns/manageengine"
	"github.com/go-acme/lego/v5/providers/dns/manual"
	"github.com/go-acme/lego/v5/providers/dns/metaname"
	"github.com/go-acme/lego/v5/providers/dns/metaregistrar"
	"github.com/go-acme/lego/v5/providers/dns/mijnhost"
	"github.com/go-acme/lego/v5/providers/dns/mittwald"
	"github.com/go-acme/lego/v5/providers/dns/myaddr"
	"github.com/go-acme/lego/v5/providers/dns/mydnsjp"
	"github.com/go-acme/lego/v5/providers/dns/mythicbeasts"
	"github.com/go-acme/lego/v5/providers/dns/namecheap"
	"github.com/go-acme/lego/v5/providers/dns/namedotcom"
	"github.com/go-acme/lego/v5/providers/dns/namesilo"
	"github.com/go-acme/lego/v5/providers/dns/namesurfer"
	"github.com/go-acme/lego/v5/providers/dns/nearlyfreespeech"
	"github.com/go-acme/lego/v5/providers/dns/nederhost"
	"github.com/go-acme/lego/v5/providers/dns/neodigit"
	"github.com/go-acme/lego/v5/providers/dns/netcup"
	"github.com/go-acme/lego/v5/providers/dns/netlify"
	"github.com/go-acme/lego/v5/providers/dns/netnod"
	"github.com/go-acme/lego/v5/providers/dns/nexdns"
	"github.com/go-acme/lego/v5/providers/dns/ngenix"
	"github.com/go-acme/lego/v5/providers/dns/nicmanager"
	"github.com/go-acme/lego/v5/providers/dns/nicru"
	"github.com/go-acme/lego/v5/providers/dns/nifcloud"
	"github.com/go-acme/lego/v5/providers/dns/njalla"
	"github.com/go-acme/lego/v5/providers/dns/nodion"
	"github.com/go-acme/lego/v5/providers/dns/ns1"
	"github.com/go-acme/lego/v5/providers/dns/octenium"
	"github.com/go-acme/lego/v5/providers/dns/omglol"
	"github.com/go-acme/lego/v5/providers/dns/onecloudru"
	"github.com/go-acme/lego/v5/providers/dns/onlinenet"
	"github.com/go-acme/lego/v5/providers/dns/openprovider"
	"github.com/go-acme/lego/v5/providers/dns/opusdns"
	"github.com/go-acme/lego/v5/providers/dns/oraclecloud"
	"github.com/go-acme/lego/v5/providers/dns/otc"
	"github.com/go-acme/lego/v5/providers/dns/ovh"
	"github.com/go-acme/lego/v5/providers/dns/pdns"
	"github.com/go-acme/lego/v5/providers/dns/plesk"
	"github.com/go-acme/lego/v5/providers/dns/pointdns"
	"github.com/go-acme/lego/v5/providers/dns/porkbun"
	"github.com/go-acme/lego/v5/providers/dns/poweradmin"
	"github.com/go-acme/lego/v5/providers/dns/rackspace"
	"github.com/go-acme/lego/v5/providers/dns/rage4"
	"github.com/go-acme/lego/v5/providers/dns/rainyun"
	"github.com/go-acme/lego/v5/providers/dns/rcodezero"
	"github.com/go-acme/lego/v5/providers/dns/regfish"
	"github.com/go-acme/lego/v5/providers/dns/regru"
	"github.com/go-acme/lego/v5/providers/dns/rimuhosting"
	"github.com/go-acme/lego/v5/providers/dns/route53"
	"github.com/go-acme/lego/v5/providers/dns/safedns"
	"github.com/go-acme/lego/v5/providers/dns/sakuracloud"
	"github.com/go-acme/lego/v5/providers/dns/scaleway"
	"github.com/go-acme/lego/v5/providers/dns/scannet"
	"github.com/go-acme/lego/v5/providers/dns/selectel"
	"github.com/go-acme/lego/v5/providers/dns/selectelv2"
	"github.com/go-acme/lego/v5/providers/dns/selfhostde"
	"github.com/go-acme/lego/v5/providers/dns/servercow"
	"github.com/go-acme/lego/v5/providers/dns/shellrent"
	"github.com/go-acme/lego/v5/providers/dns/simply"
	"github.com/go-acme/lego/v5/providers/dns/sonic"
	"github.com/go-acme/lego/v5/providers/dns/spaceship"
	"github.com/go-acme/lego/v5/providers/dns/stackpath"
	"github.com/go-acme/lego/v5/providers/dns/syse"
	"github.com/go-acme/lego/v5/providers/dns/technitium"
	"github.com/go-acme/lego/v5/providers/dns/tele3"
	"github.com/go-acme/lego/v5/providers/dns/tencentcloud"
	"github.com/go-acme/lego/v5/providers/dns/timewebcloud"
	"github.com/go-acme/lego/v5/providers/dns/todaynic"
	"github.com/go-acme/lego/v5/providers/dns/transip"
	"github.com/go-acme/lego/v5/providers/dns/ucloud"
	"github.com/go-acme/lego/v5/providers/dns/ultradns"
	"github.com/go-acme/lego/v5/providers/dns/uniteddomains"
	"github.com/go-acme/lego/v5/providers/dns/variomedia"
	"github.com/go-acme/lego/v5/providers/dns/veesp"
	"github.com/go-acme/lego/v5/providers/dns/vegadns"
	"github.com/go-acme/lego/v5/providers/dns/vercel"
	"github.com/go-acme/lego/v5/providers/dns/versio"
	"github.com/go-acme/lego/v5/providers/dns/vinyldns"
	"github.com/go-acme/lego/v5/providers/dns/virtualname"
	"github.com/go-acme/lego/v5/providers/dns/vkcloud"
	"github.com/go-acme/lego/v5/providers/dns/volcengine"
	"github.com/go-acme/lego/v5/providers/dns/vscale"
	"github.com/go-acme/lego/v5/providers/dns/vultr"
	"github.com/go-acme/lego/v5/providers/dns/wannafind"
	"github.com/go-acme/lego/v5/providers/dns/webnamesca"
	"github.com/go-acme/lego/v5/providers/dns/webnamesru"
	"github.com/go-acme/lego/v5/providers/dns/websupport"
	"github.com/go-acme/lego/v5/providers/dns/wedos"
	"github.com/go-acme/lego/v5/providers/dns/westcn"
	"github.com/go-acme/lego/v5/providers/dns/xinnet"
	"github.com/go-acme/lego/v5/providers/dns/yandex"
	"github.com/go-acme/lego/v5/providers/dns/yandex360"
	"github.com/go-acme/lego/v5/providers/dns/yandexcloud"
	"github.com/go-acme/lego/v5/providers/dns/zilore"
	"github.com/go-acme/lego/v5/providers/dns/zoneedit"
	"github.com/go-acme/lego/v5/providers/dns/zoneee"
	"github.com/go-acme/lego/v5/providers/dns/zonomi"
)

// dnsProviderFactoryFunc is a function that calls a provider's
// constructor and returns the provider interface.
type dnsProviderFactoryFunc func() (challenge.Provider, error)

// dnsProviderFactory is a factory for all of the valid DNS providers
// supported by ACME provider.
var dnsProviderFactory = map[string]dnsProviderFactoryFunc{
	"abion": func() (challenge.Provider, error) {
		p, err := abion.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"acmedns": func() (challenge.Provider, error) {
		p, err := acmedns.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"active24": func() (challenge.Provider, error) {
		p, err := active24.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"alidns": func() (challenge.Provider, error) {
		p, err := alidns.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"aliesa": func() (challenge.Provider, error) {
		p, err := aliesa.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"allinkl": func() (challenge.Provider, error) {
		p, err := allinkl.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"alwaysdata": func() (challenge.Provider, error) {
		p, err := alwaysdata.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"anexia": func() (challenge.Provider, error) {
		p, err := anexia.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"artfiles": func() (challenge.Provider, error) {
		p, err := artfiles.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"arvancloud": func() (challenge.Provider, error) {
		p, err := arvancloud.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"auroradns": func() (challenge.Provider, error) {
		p, err := auroradns.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"autodns": func() (challenge.Provider, error) {
		p, err := autodns.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"axelname": func() (challenge.Provider, error) {
		p, err := axelname.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"azion": func() (challenge.Provider, error) {
		p, err := azion.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"azuredns": func() (challenge.Provider, error) {
		mapEnvironmentVariableValues(map[string]string{
			"ARM_CLIENT_ID":       "AZURE_CLIENT_ID",
			"ARM_CLIENT_SECRET":   "AZURE_CLIENT_SECRET",
			"ARM_RESOURCE_GROUP":  "AZURE_RESOURCE_GROUP",
			"ARM_SUBSCRIPTION_ID": "AZURE_SUBSCRIPTION_ID",
			"ARM_TENANT_ID":       "AZURE_TENANT_ID",
		})
		p, err := azuredns.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"baiducloud": func() (challenge.Provider, error) {
		p, err := baiducloud.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"beget": func() (challenge.Provider, error) {
		p, err := beget.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"binarylane": func() (challenge.Provider, error) {
		p, err := binarylane.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"bindman": func() (challenge.Provider, error) {
		p, err := bindman.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"bluecat": func() (challenge.Provider, error) {
		p, err := bluecat.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"bluecatv2": func() (challenge.Provider, error) {
		p, err := bluecatv2.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"bookmyname": func() (challenge.Provider, error) {
		p, err := bookmyname.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"bunny": func() (challenge.Provider, error) {
		p, err := bunny.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"checkdomain": func() (challenge.Provider, error) {
		p, err := checkdomain.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"civo": func() (challenge.Provider, error) {
		p, err := civo.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"clouddns": func() (challenge.Provider, error) {
		p, err := clouddns.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"cloudflare": func() (challenge.Provider, error) {
		p, err := cloudflare.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"cloudns": func() (challenge.Provider, error) {
		p, err := cloudns.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"cloudru": func() (challenge.Provider, error) {
		p, err := cloudru.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"com35": func() (challenge.Provider, error) {
		p, err := com35.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"connbyte": func() (challenge.Provider, error) {
		p, err := connbyte.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"conoha": func() (challenge.Provider, error) {
		p, err := conoha.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"conohav3": func() (challenge.Provider, error) {
		p, err := conohav3.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"constellix": func() (challenge.Provider, error) {
		p, err := constellix.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"corenetworks": func() (challenge.Provider, error) {
		p, err := corenetworks.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"cpanel": func() (challenge.Provider, error) {
		p, err := cpanel.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"curanet": func() (challenge.Provider, error) {
		p, err := curanet.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"czechia": func() (challenge.Provider, error) {
		p, err := czechia.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"dandomain": func() (challenge.Provider, error) {
		p, err := dandomain.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"ddnss": func() (challenge.Provider, error) {
		p, err := ddnss.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"derak": func() (challenge.Provider, error) {
		p, err := derak.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"desec": func() (challenge.Provider, error) {
		p, err := desec.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"designate": func() (challenge.Provider, error) {
		p, err := designate.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"digitalocean": func() (challenge.Provider, error) {
		p, err := digitalocean.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"dinahosting": func() (challenge.Provider, error) {
		p, err := dinahosting.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"directadmin": func() (challenge.Provider, error) {
		p, err := directadmin.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"dns51": func() (challenge.Provider, error) {
		p, err := dns51.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"dnscale": func() (challenge.Provider, error) {
		p, err := dnscale.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"dnsexit": func() (challenge.Provider, error) {
		p, err := dnsexit.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"dnshomede": func() (challenge.Provider, error) {
		p, err := dnshomede.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"dnsimple": func() (challenge.Provider, error) {
		p, err := dnsimple.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"dnsla": func() (challenge.Provider, error) {
		p, err := dnsla.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"dnsmadeeasy": func() (challenge.Provider, error) {
		p, err := dnsmadeeasy.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"dnsservices": func() (challenge.Provider, error) {
		p, err := dnsservices.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"dnsupdate": func() (challenge.Provider, error) {
		p, err := dnsupdate.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"dode": func() (challenge.Provider, error) {
		p, err := dode.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"domeneshop": func() (challenge.Provider, error) {
		p, err := domeneshop.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"dreamhost": func() (challenge.Provider, error) {
		p, err := dreamhost.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"duckdns": func() (challenge.Provider, error) {
		p, err := duckdns.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"dyn": func() (challenge.Provider, error) {
		p, err := dyn.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"dynadot": func() (challenge.Provider, error) {
		p, err := dynadot.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"dyndnsfree": func() (challenge.Provider, error) {
		p, err := dyndnsfree.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"dynu": func() (challenge.Provider, error) {
		p, err := dynu.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"easydns": func() (challenge.Provider, error) {
		p, err := easydns.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"edgecenter": func() (challenge.Provider, error) {
		p, err := edgecenter.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"edgedns": func() (challenge.Provider, error) {
		p, err := edgedns.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"edgeone": func() (challenge.Provider, error) {
		p, err := edgeone.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"efficientip": func() (challenge.Provider, error) {
		p, err := efficientip.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"epik": func() (challenge.Provider, error) {
		p, err := epik.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"eurodns": func() (challenge.Provider, error) {
		p, err := eurodns.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"euserv": func() (challenge.Provider, error) {
		p, err := euserv.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"excedo": func() (challenge.Provider, error) {
		p, err := excedo.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"exec": func() (challenge.Provider, error) {
		p, err := exec.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"exoscale": func() (challenge.Provider, error) {
		p, err := exoscale.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"f5xc": func() (challenge.Provider, error) {
		p, err := f5xc.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"fornex": func() (challenge.Provider, error) {
		p, err := fornex.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"freemyip": func() (challenge.Provider, error) {
		p, err := freemyip.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"gandi": func() (challenge.Provider, error) {
		p, err := gandi.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"gandiv5": func() (challenge.Provider, error) {
		p, err := gandiv5.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"gcloud": func() (challenge.Provider, error) {
		p, err := gcloud.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"gcore": func() (challenge.Provider, error) {
		p, err := gcore.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"gehirn": func() (challenge.Provider, error) {
		p, err := gehirn.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"gigahostno": func() (challenge.Provider, error) {
		p, err := gigahostno.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"glesys": func() (challenge.Provider, error) {
		p, err := glesys.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"gname": func() (challenge.Provider, error) {
		p, err := gname.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"godaddy": func() (challenge.Provider, error) {
		p, err := godaddy.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"gravity": func() (challenge.Provider, error) {
		p, err := gravity.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"hetzner": func() (challenge.Provider, error) {
		p, err := hetzner.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"hostingde": func() (challenge.Provider, error) {
		p, err := hostingde.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"hostinger": func() (challenge.Provider, error) {
		p, err := hostinger.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"hostingnl": func() (challenge.Provider, error) {
		p, err := hostingnl.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"hosttech": func() (challenge.Provider, error) {
		p, err := hosttech.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"hostup": func() (challenge.Provider, error) {
		p, err := hostup.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"httpnet": func() (challenge.Provider, error) {
		p, err := httpnet.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"httpreq": func() (challenge.Provider, error) {
		p, err := httpreq.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"huaweicloud": func() (challenge.Provider, error) {
		p, err := huaweicloud.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"hurricane": func() (challenge.Provider, error) {
		p, err := hurricane.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"hyperone": func() (challenge.Provider, error) {
		p, err := hyperone.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"ibmcloud": func() (challenge.Provider, error) {
		p, err := ibmcloud.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"iijdpf": func() (challenge.Provider, error) {
		p, err := iijdpf.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"infoblox": func() (challenge.Provider, error) {
		p, err := infoblox.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"infomaniak": func() (challenge.Provider, error) {
		p, err := infomaniak.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"internetbs": func() (challenge.Provider, error) {
		p, err := internetbs.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"inwx": func() (challenge.Provider, error) {
		p, err := inwx.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"ionos": func() (challenge.Provider, error) {
		p, err := ionos.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"ionoscloud": func() (challenge.Provider, error) {
		p, err := ionoscloud.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"ipv64": func() (challenge.Provider, error) {
		p, err := ipv64.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"ispconfig": func() (challenge.Provider, error) {
		p, err := ispconfig.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"ispconfigddns": func() (challenge.Provider, error) {
		p, err := ispconfigddns.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"jdcloud": func() (challenge.Provider, error) {
		p, err := jdcloud.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"joker": func() (challenge.Provider, error) {
		p, err := joker.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"katapult": func() (challenge.Provider, error) {
		p, err := katapult.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"keyhelp": func() (challenge.Provider, error) {
		p, err := keyhelp.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"leaseweb": func() (challenge.Provider, error) {
		p, err := leaseweb.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"liara": func() (challenge.Provider, error) {
		p, err := liara.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"lightsail": func() (challenge.Provider, error) {
		p, err := lightsail.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"limacity": func() (challenge.Provider, error) {
		p, err := limacity.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"linode": func() (challenge.Provider, error) {
		p, err := linode.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"liquidweb": func() (challenge.Provider, error) {
		p, err := liquidweb.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"loopia": func() (challenge.Provider, error) {
		p, err := loopia.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"luadns": func() (challenge.Provider, error) {
		p, err := luadns.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"mailinabox": func() (challenge.Provider, error) {
		p, err := mailinabox.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"manageengine": func() (challenge.Provider, error) {
		p, err := manageengine.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"manual": func() (challenge.Provider, error) {
		p, err := manual.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"metaname": func() (challenge.Provider, error) {
		p, err := metaname.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"metaregistrar": func() (challenge.Provider, error) {
		p, err := metaregistrar.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"mijnhost": func() (challenge.Provider, error) {
		p, err := mijnhost.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"mittwald": func() (challenge.Provider, error) {
		p, err := mittwald.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"myaddr": func() (challenge.Provider, error) {
		p, err := myaddr.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"mydnsjp": func() (challenge.Provider, error) {
		p, err := mydnsjp.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"mythicbeasts": func() (challenge.Provider, error) {
		p, err := mythicbeasts.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"namecheap": func() (challenge.Provider, error) {
		p, err := namecheap.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"namedotcom": func() (challenge.Provider, error) {
		p, err := namedotcom.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"namesilo": func() (challenge.Provider, error) {
		p, err := namesilo.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"namesurfer": func() (challenge.Provider, error) {
		p, err := namesurfer.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"nearlyfreespeech": func() (challenge.Provider, error) {
		p, err := nearlyfreespeech.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"nederhost": func() (challenge.Provider, error) {
		p, err := nederhost.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"neodigit": func() (challenge.Provider, error) {
		p, err := neodigit.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"netcup": func() (challenge.Provider, error) {
		p, err := netcup.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"netlify": func() (challenge.Provider, error) {
		p, err := netlify.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"netnod": func() (challenge.Provider, error) {
		p, err := netnod.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"nexdns": func() (challenge.Provider, error) {
		p, err := nexdns.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"ngenix": func() (challenge.Provider, error) {
		p, err := ngenix.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"nicmanager": func() (challenge.Provider, error) {
		p, err := nicmanager.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"nicru": func() (challenge.Provider, error) {
		p, err := nicru.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"nifcloud": func() (challenge.Provider, error) {
		p, err := nifcloud.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"njalla": func() (challenge.Provider, error) {
		p, err := njalla.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"nodion": func() (challenge.Provider, error) {
		p, err := nodion.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"ns1": func() (challenge.Provider, error) {
		p, err := ns1.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"octenium": func() (challenge.Provider, error) {
		p, err := octenium.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"omglol": func() (challenge.Provider, error) {
		p, err := omglol.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"onecloudru": func() (challenge.Provider, error) {
		p, err := onecloudru.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"onlinenet": func() (challenge.Provider, error) {
		p, err := onlinenet.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"openprovider": func() (challenge.Provider, error) {
		p, err := openprovider.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"opusdns": func() (challenge.Provider, error) {
		p, err := opusdns.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"oraclecloud": func() (challenge.Provider, error) {
		p, err := oraclecloud.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"otc": func() (challenge.Provider, error) {
		p, err := otc.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"ovh": func() (challenge.Provider, error) {
		p, err := ovh.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"pdns": func() (challenge.Provider, error) {
		p, err := pdns.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"plesk": func() (challenge.Provider, error) {
		p, err := plesk.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"pointdns": func() (challenge.Provider, error) {
		p, err := pointdns.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"porkbun": func() (challenge.Provider, error) {
		p, err := porkbun.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"poweradmin": func() (challenge.Provider, error) {
		p, err := poweradmin.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"rackspace": func() (challenge.Provider, error) {
		p, err := rackspace.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"rage4": func() (challenge.Provider, error) {
		p, err := rage4.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"rainyun": func() (challenge.Provider, error) {
		p, err := rainyun.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"rcodezero": func() (challenge.Provider, error) {
		p, err := rcodezero.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"regfish": func() (challenge.Provider, error) {
		p, err := regfish.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"regru": func() (challenge.Provider, error) {
		p, err := regru.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"rimuhosting": func() (challenge.Provider, error) {
		p, err := rimuhosting.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"route53": func() (challenge.Provider, error) {
		p, err := route53.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"safedns": func() (challenge.Provider, error) {
		p, err := safedns.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"sakuracloud": func() (challenge.Provider, error) {
		p, err := sakuracloud.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"scaleway": func() (challenge.Provider, error) {
		p, err := scaleway.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"scannet": func() (challenge.Provider, error) {
		p, err := scannet.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"selectel": func() (challenge.Provider, error) {
		p, err := selectel.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"selectelv2": func() (challenge.Provider, error) {
		p, err := selectelv2.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"selfhostde": func() (challenge.Provider, error) {
		p, err := selfhostde.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"servercow": func() (challenge.Provider, error) {
		p, err := servercow.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"shellrent": func() (challenge.Provider, error) {
		p, err := shellrent.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"simply": func() (challenge.Provider, error) {
		p, err := simply.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"sonic": func() (challenge.Provider, error) {
		p, err := sonic.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"spaceship": func() (challenge.Provider, error) {
		p, err := spaceship.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"stackpath": func() (challenge.Provider, error) {
		p, err := stackpath.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"syse": func() (challenge.Provider, error) {
		p, err := syse.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"technitium": func() (challenge.Provider, error) {
		p, err := technitium.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"tele3": func() (challenge.Provider, error) {
		p, err := tele3.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"tencentcloud": func() (challenge.Provider, error) {
		p, err := tencentcloud.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"timewebcloud": func() (challenge.Provider, error) {
		p, err := timewebcloud.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"todaynic": func() (challenge.Provider, error) {
		p, err := todaynic.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"transip": func() (challenge.Provider, error) {
		p, err := transip.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"ucloud": func() (challenge.Provider, error) {
		p, err := ucloud.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"ultradns": func() (challenge.Provider, error) {
		p, err := ultradns.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"uniteddomains": func() (challenge.Provider, error) {
		p, err := uniteddomains.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"variomedia": func() (challenge.Provider, error) {
		p, err := variomedia.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"veesp": func() (challenge.Provider, error) {
		p, err := veesp.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"vegadns": func() (challenge.Provider, error) {
		p, err := vegadns.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"vercel": func() (challenge.Provider, error) {
		p, err := vercel.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"versio": func() (challenge.Provider, error) {
		p, err := versio.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"vinyldns": func() (challenge.Provider, error) {
		p, err := vinyldns.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"virtualname": func() (challenge.Provider, error) {
		p, err := virtualname.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"vkcloud": func() (challenge.Provider, error) {
		p, err := vkcloud.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"volcengine": func() (challenge.Provider, error) {
		p, err := volcengine.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"vscale": func() (challenge.Provider, error) {
		p, err := vscale.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"vultr": func() (challenge.Provider, error) {
		p, err := vultr.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"wannafind": func() (challenge.Provider, error) {
		p, err := wannafind.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"webnamesca": func() (challenge.Provider, error) {
		p, err := webnamesca.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"webnamesru": func() (challenge.Provider, error) {
		p, err := webnamesru.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"websupport": func() (challenge.Provider, error) {
		p, err := websupport.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"wedos": func() (challenge.Provider, error) {
		p, err := wedos.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"westcn": func() (challenge.Provider, error) {
		p, err := westcn.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"xinnet": func() (challenge.Provider, error) {
		p, err := xinnet.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"yandex": func() (challenge.Provider, error) {
		p, err := yandex.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"yandex360": func() (challenge.Provider, error) {
		p, err := yandex360.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"yandexcloud": func() (challenge.Provider, error) {
		p, err := yandexcloud.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"zilore": func() (challenge.Provider, error) {
		p, err := zilore.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"zoneedit": func() (challenge.Provider, error) {
		p, err := zoneedit.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"zoneee": func() (challenge.Provider, error) {
		p, err := zoneee.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"zonomi": func() (challenge.Provider, error) {
		p, err := zonomi.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
}
