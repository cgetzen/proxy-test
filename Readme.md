# Strategy

At first glance, the problem can be almost entirely solved with a webserver like nginx.

However, implementing it in nginx will not allow us to showcase building / deploying docker images. Therefore I will implement this entirely in a programming language.

# Build and Deploy

scripts/build-and-push.sh | xargs scripts/cd.sh

# Enhancement #1
For all options, I would recommend first installing AWS Load Balancer Controller, as the kubernetes in-tree controller is deprecated and lacking features.

Option 1:
- Install external-dns
- Create a service of type LoadBalancer with the following annotations:
    external-dns.alpha.kubernetes.io/hostname: test-subaccount-1.rr.mu.
    service.beta.kubernetes.io/aws-load-balancer-ssl-cert: arn:aws:acm:us-east-1:663118211814:certificate/f77cd7ad-b251-470f-aff3-8a3faa12e4c2

Option 2:
- Create load balancer / certificates / target group infrastructure in terraform
- Use TargetGroupBindings to connect your service to the loadbalancer

Other options include using IP target types over instance target types, which would reduce the load on kube-proxy and enable healthchecking pods at the TG-layer.

# Enhancement #2

## Redirects
Redirects are now controlled through the config/config.yaml file.
It will look like this:
```
redirects:
  page1.html: page2.html
```
Which indicate that page1.html should be 302'd to page2.html

## Load all assets
Previously, each asset was hard-coded. Now, the server will look up and serve everything under  assets/

# Enhancement #3

This is implemented by changing the filename to end with `.tpl`. The contents of the environment variable `TEMPLATE_DATA` will be displayed.

# Enhancement #4
```
kubectl logs proxy-test-68b7c56b4d-pd659
Starting...
Adding redirect from /page1.html to /page2.html
Caching config.html
Caching index.html
Redirect already set up for path page1.html. Will not serve file.
Caching page2.html
```