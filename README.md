Proxy
=====

Dead simple proxy app to get access behind firewall.
Why? Sometimes you want can access things that don't have internet access

Use Case
--------
If you want to access things behind a vpc. This can be easily deployed on Cloud Foundry and expose url you can't access outside of network. This is **NOT SAFE** to keep running. It will expose internal access to anything app running on Cloud Foundry can see internally.

Running
-------

```
go get github.com/longnguyen11288/proxy
proxy
```

Usage
----


Going to  https://proxy.apps.10.244.4.10.xip.io/10.244.5.251:49161
 will proxy to 10.244.5.251:49161 and seem transparent that it is going there.
