# Traefik EDNS0/ECS Middleware

## Install

```yaml
additionalArguments:
- "--experimental.plugins.edns0.moduleName=github.com/CumpsD/edns0"
- "--experimental.plugins.edns0.version=v0.0.3"
```

```yaml
apiVersion: traefik.containo.us/v1alpha1
kind: Middleware
metadata:
  name: edns0
spec:
  plugin:
    edns0:
      Prefix: "EDNS0"
```
