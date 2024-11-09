---
title: OAuth auth
layout: default
parent: Middleware
nav_order: 5
---

### OAuth middleware

Example of Google provider

```yaml
  - name: google-oauth
    type: oauth
    paths:
      - /*
    rule:
      clientId: xxx
      clientSecret: xxxx
      # oauth provider google, gitlab, github, amazon, facebook, custom
      provider: google # facebook, gitlab, github, amazon
      redirectUrl: https://example.com/callback/protected
      #RedirectPath is the PATH to redirect users after authentication, e.g: /my-protected-path/dashboard
      redirectPath: /dashboard
      scopes:
        - https://www.googleapis.com/auth/userinfo.email
        - https://www.googleapis.com/auth/userinfo.profile
      state: randomStateString
      jwtSecret: your-strong-jwt-secret | It's optional

```

Example of Authentik provider

```yaml
    - name: oauth-authentik
      type: oauth
      paths:
        - /protected
        - /example-of-oauth
      rule:
        clientId: xxx
        clientSecret: xxx
        # oauth provider google, gitlab, github, amazon, facebook, custom
        provider: custom
        endpoint:
          authUrl: https://authentik.example.com/application/o/authorize/
          tokenUrl: https://authentik.example.com/application/o/token/
          userInfoUrl: https://authentik.example.com/application/o/userinfo/
        redirectUrl: https://example.com/callback
        #RedirectPath is the PATH to redirect users after authentication, e.g: /my-protected-path/dashboard
        redirectPath: ''
        #CookiePath e.g.: /my-protected-path or / || by default is applied on a route path
        cookiePath: "/"
        scopes:
          - email
          - openid
        state: randomStateString
        jwtSecret: your-strong-jwt-secret | It's optional

```
### Access middleware

Access middleware prevents access to a route or specific route path.

Example of access middleware
```yaml
  # The server will return 403
  - name: api-forbidden-paths
    type: access
    ## prevents access paths
    paths:
      - /swagger-ui/*
      - /v2/swagger-ui/*
      - /api-docs/*
      - /internal/*
      - /actuator/*
```
### RateLimit middleware

The RateLimit middleware ensures that services will receive a fair amount of requests, and allows one to define what fair is.

Example of rateLimit middleware
```yaml

```

### Apply middleware on the route

```yaml
  ##### Define routes
  routes:
    - name: Basic auth
      path: /protected
      rewrite: /
      destination: 'https://example.com'
      methods: [POST, PUT, GET]
      healthCheck:
      cors: {}
      middlewares:
        - oauth-authentik
```