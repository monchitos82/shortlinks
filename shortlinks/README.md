# shortlinks

This is a backend service that intends to give you a URL alias functionality where you create (easy to remember) handles and you use them with a prefix.

## How to run?

Make sure you have a **Postgres** database beforehand. You will need it for your backend to be able to query what's existing or just to keep record of your short links.

To start the service, just run using `go mod run` or build a binary if you plan on hosting and automating.

If you feel like turning it into a service or just have a binary, run: `go build`

## How to test?

`CONFIG_FILE=$(pwd)/config.yaml; go test ./tests/... -v`

This would work as well for execution, the default behavior is searching for `config.yaml` where the service runs.


## Examples:

### To create a new shortlink:
  You can have as many short links to the same URL as you want, there's no validation on duplicate sources (there is on handles):
  `curl -X POST http://localhost:8080/shorts -H "Content-Type: application/json" -d '{"handle": "gh", "url": "https://github.com"}'`
  
  As an alternative, you can use the `/updates` endpoint to load a form in your browser, so you have a simple UI to add or update new entries.

  
### To be redirected:
  `curl -L https://localhost:8080/gh` this would redirect your browser (or CURL in this case) to whatever points the `gh` handle, e.g. `gh` redirects to `https://github.com`

To have the best experience for this redirection project, you would want to have this project serving in port 80 and a record in your DNS so you can just use this in your browser like: `https://sl/gh` and be redirected to (github) the site associated to the `gh` handle.

If you want to expose this directly, you are advised to use a proxy from your HTTP server so you don't need to add `CAP_NET_BIND_SERVICE` permissions and add security concerns; you can add this directive to your nginx site:

```
	location / {
            proxy_pass http://<hostname>:8000;
            proxy_set_header Host $host;
            proxy_set_header X-Real-IP $remote_addr;
            proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
            proxy_set_header X-Forwarded-Proto $scheme;
        }
```

### How to add this to my DNS?

Well, you can't in the strict sense, even if you add an SRV entry, the result would be messy. But you can add an entry to a small DHCP/DNS instance like dnsmasq! Or you can simply add the entry to your hosts file. These require minimum effort and are very easy to setup.

```
# dnsmasq.conf entry
address=/sl/<HOST RUNNING [PROXY FOR] SHORTLINKS>
address=/.sl/<HOST RUNNING [PROXY FOR] SHORTLINKS>

# add your DNS downstream so recursive search happens
server=<YOUR DNS IP>
```

### Special endpoints

There are a few endpoints in the project intended to serve some functionality other than redirection, these should not be added as handlers:
* `/shorts` this is the endpoint used to add new entries into the database
* `/urls` this endpoint will return all entries from the database in plain text
* `/updates` this endpoint returns a form to add or update entries into the database
* `/healthcheck` this is a healthcheck, it returns nothing but a 200 code
