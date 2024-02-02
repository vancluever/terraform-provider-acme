//go:generate go-bindata ./templates
package main

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/BurntSushi/toml"
)

// envVarAliases are Terraform-specific environment variables for
// specific providers.
var envVarAliases = map[string]map[string]string{
	"azure": {
		"ARM_CLIENT_ID":       "AZURE_CLIENT_ID",
		"ARM_CLIENT_SECRET":   "AZURE_CLIENT_SECRET",
		"ARM_SUBSCRIPTION_ID": "AZURE_SUBSCRIPTION_ID",
		"ARM_TENANT_ID":       "AZURE_TENANT_ID",
		"ARM_RESOURCE_GROUP":  "AZURE_RESOURCE_GROUP",
		"ARM_OIDC_TOKEN":      "AZURE_OIDC_TOKEN",
		"ARM_USE_OIDC":        "AZURE_USE_OIDC",
	},
	"azuredns": {
		"ARM_CLIENT_ID":       "AZURE_CLIENT_ID",
		"ARM_CLIENT_SECRET":   "AZURE_CLIENT_SECRET",
		"ARM_SUBSCRIPTION_ID": "AZURE_SUBSCRIPTION_ID",
		"ARM_TENANT_ID":       "AZURE_TENANT_ID",
		"ARM_RESOURCE_GROUP":  "AZURE_RESOURCE_GROUP",
		"ARM_OIDC_TOKEN":      "AZURE_OIDC_TOKEN",
		"ARM_USE_OIDC":        "AZURE_USE_OIDC",
	},
}

// providerURLs is a list of providers to override provider pages
// for. Usually this is just used to provide blank links for
// anything that would normally just link back to the provider page
// in lego.
var providerURLs = map[string]string{
	"exec":    "",
	"httpreq": "",
}

// dnsProviderGoTemplate is the template for
// dnsProviderGoTemplateText.
var dnsProviderGoTemplate = template.Must(
	template.New("dns-provider-go-template").Parse(string(MustAsset("templates/dns-provider-go-template.tmpl"))),
)

// dnsProviderDocTemplate is the template for DNS provider
// documentation.
var dnsProviderDocTemplate = template.Must(
	template.New("dns-provider-doc-template").Parse(string(MustAsset("templates/dns-provider-doc-template.tmpl"))),
)

// legoPkgPath is the root lego package path to use.
const legoPkgPath = "github.com/go-acme/lego/v4"

// Type from "go help mod edit"
type pkgInfoGoMod struct {
	Module  pkgInfoModule
	Go      string
	Require []pkgInfoRequire
	Exclude []pkgInfoModule
	Replace []pkgInfoReplace
}

// Type from "go help mod edit"
type pkgInfoModule struct {
	Path    string
	Version string
}

// Type from "go help mod edit"
type pkgInfoRequire struct {
	Path     string
	Version  string
	Indirect bool
}

// Type from "go help mod edit"
type pkgInfoReplace struct {
	Old pkgInfoModule
	New pkgInfoModule
}

type dnsProviderInfo struct {
	Name          string
	URL           string
	Code          string
	GoPkg         string
	Additional    string
	Configuration dnsProviderConfig
	EnvVarAliases map[string]string
}

type dnsProviderConfig struct {
	Credentials map[string]string
	Additional  map[string]string
}

func (c dnsProviderConfig) Present() bool {
	return len(c.Credentials) > 0 || len(c.Additional) > 0
}

// execCommand is a exec.Cmd builder that just sets the error stream
// to stderr.
func execCommand(cmd string, args ...string) *exec.Cmd {
	c := exec.Command(cmd, args...)
	c.Stderr = os.Stderr
	return c
}

// loadProviders loads all of the provider information from the
// provider TOML files.
func loadProviders() []dnsProviderInfo {
	out, err := execCommand("go", "mod", "edit", "-json").Output()
	if err != nil {
		log.Fatal(err)
	}

	var info pkgInfoGoMod
	if err := json.Unmarshal(out, &info); err != nil {
		log.Fatal(err)
	}

	var version string
	for _, req := range info.Require {
		if req.Path == legoPkgPath {
			version = req.Version
			break
		}
	}

	if version == "" {
		log.Fatalf("package %q not found in go.mod, cannot get version", legoPkgPath)
	}

	out, err = execCommand("go", "env", "GOPATH").Output()
	if err != nil {
		log.Fatal(err)
	}

	pkgDir := filepath.Join(
		strings.TrimSpace(string(out)), "pkg", "mod", strings.ReplaceAll(legoPkgPath, "/", string(os.PathSeparator))+"@"+version)

	// Check to see if this is actually a directory, in case it's not
	// in the cache.
	fi, err := os.Stat(pkgDir)
	if err != nil {
		if os.IsNotExist(err) {
			err := execCommand("go", "mod", "download", legoPkgPath+"@"+version).Run()
			if err != nil {
				log.Fatal(err)
			}

			// Get fileinfo again and fail 100% if still error
			fi, err = os.Stat(pkgDir)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal(err)
		}
	}

	if !fi.Mode().IsDir() {
		log.Fatalf("not a directory: %q", pkgDir)
	}

	// Start loading in the TOML files
	var result []dnsProviderInfo
	rootDir := filepath.Join(pkgDir, "providers", "dns")
	if err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if filepath.Ext(path) != ".toml" {
			return nil
		}

		var p dnsProviderInfo
		_, err = toml.DecodeFile(path, &p)
		if err != nil {
			return err
		}

		// We work off of Go packages to find the metadata, but some
		// packages have different codes (ie: acme-dns for acmedns in Go)
		// so we need to save the provider as the package name.
		p.GoPkg, err = filepath.Rel(rootDir, filepath.Dir(path))
		if err != nil {
			return err
		}

		// Environment variable aliases if we have them (ie: azure)
		if aliases, ok := envVarAliases[p.Code]; ok {
			p.EnvVarAliases = aliases
		}

		// Check for a provider URL override
		if url, ok := providerURLs[p.Code]; ok {
			p.URL = url
		}

		// A couple of docs have hugo template artifacts that could use
		// stripping, just do this for "notice" for now which seems to be
		// the only one that's in use.
		p.Additional = strings.ReplaceAll(
			p.Additional, "{{% notice note %}}\n", "-> **NOTE**: ")
		p.Additional = strings.ReplaceAll(
			p.Additional, "{{% /notice %}}\n", "")

		result = append(result, p)
		return nil
	}); err != nil {
		log.Fatal(err)
	}

	return result
}

// generateGo generates the factory template file.
func generateGo(providers []dnsProviderInfo) {
	b := new(bytes.Buffer)
	if err := dnsProviderGoTemplate.Execute(b, struct {
		PkgPath   string
		Providers []dnsProviderInfo
	}{
		PkgPath:   legoPkgPath,
		Providers: providers,
	}); err != nil {
		log.Fatal(err)
	}

	cmd := execCommand("gofmt")
	cmd.Stdin = b
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(os.Args[2], out, 0666); err != nil {
		log.Fatal(err)
	}
}

// generateProviderDocs generates each of the provider documentation
// pages.
func generateProviderDocs(providers []dnsProviderInfo) {
	for _, provider := range providers {
		b := new(bytes.Buffer)
		if err := dnsProviderDocTemplate.Execute(b, provider); err != nil {
			log.Fatal(err)
		}

		path := filepath.Join(os.Args[2], "dns-providers-"+provider.Code+".md")
		if err := os.WriteFile(path, b.Bytes(), 0666); err != nil {
			log.Fatal(err)
		}

		log.Println("wrote", provider.Code, "documentation to:", path)
	}
}

func main() {
	if len(os.Args) < 3 {
		log.Fatal("usage: generate-dns-providers [go | doc ] PATH")
	}

	providers := loadProviders()
	switch os.Args[1] {
	case "go":
		generateGo(providers)

	case "doc":
		generateProviderDocs(providers)

	default:
		log.Fatal("usage: generate-dns-providers [go | doc] PATH")
	}
}
