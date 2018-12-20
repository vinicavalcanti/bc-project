#!/bin/bash
ROLE="    - rolearn: $1\n      username: system:node:{{EC2PrivateDNSName}}\n      groups:\n    - system:bootstrappers\n    - system:nodes"
kubectl get -n kube-system configmap/aws-auth -o yaml | awk "/mapRoles: \|/{print;print \"$ROLE\";next}1" > /tmp/aws-auth-patch.yml
cat /tmp/aws-auth-patch.yml
#kubectl patch configmap/aws-auth -n kube-system --patch "$(cat /tmp/aws-auth-patch.yml)"
