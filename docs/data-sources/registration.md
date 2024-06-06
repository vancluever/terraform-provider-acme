# acme_server_url

The `acme_server_url` data source can be used to retrieve the CA server URL
that the provider is currently configured for.

## Example

The following example populates the `server_url` output with the currently
configured CA server URL.

```hcl
provider "acme" {
  server_url = "https://acme-staging-v02.api.letsencrypt.org/directory"
}

data "acme_server_url" "url" {}

output "server_url" {
  value = data.acme_server_url.url.server_url
}
```

#### Argument Reference

This data source takes no arguments.

#### Attribute Reference

The following attributes are exported:

* `id`: the CA server URL that the provider is currently configured for. 
* `server_url`: the CA server URL that the provider is currently configured
  for. Same as `id`.
