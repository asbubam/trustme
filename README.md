# TrustMe
a simple CLI to manage Trust Relationship of AWS IAM Role  
Trust me. You can add and remove your identity to `Trust Relastionship` easily.

If you want to know about Trust Relationship, then recommend read this article: https://aws.amazon.com/blogs/security/how-to-use-trust-policies-with-iam-roles/

## Install
```shell
$ brew tap asbubam/trustme
$ brew update
$ brew install trustme
$ trustme version
0.0.1
```

## Run TrustMe
* Check current Trust Relationship of IAM Role
```shell
$ trustme tr check --role-name={RoleName}
```

* Check current polices of IAM Role
```shell
$ trustme policy check --role-name={RoleName}
```

* Check your IAM identity.
```shell
$ trustme whoami
```

* Add your identity to Trust Relationship.
```shell
$ trustme tr add --role-name {RoleName}
```

* Assume Role
```shell
$ trustme assume --role-name {RoleName}
```

* Remove your identity from Trust Relationship.
```shell
$ trustme tr remove --role-name {RoleName}
```

