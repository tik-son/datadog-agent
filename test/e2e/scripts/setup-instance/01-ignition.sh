#!/usr/bin/env bash

printf '=%.0s' {0..79} ; echo
set -ex

cd "$(dirname $0)"
ssh-keygen -b 4096 -t rsa -C "datadog" -N "" -f "id_rsa"
SSH_RSA=$(cat id_rsa.pub)

docker run --rm -i quay.io/coreos/fcct:release --pretty --strict <<EOF | tee ignition.json
variant: fcos
version: 1.1.0
passwd:
  users:
    - name: core
      ssh_authorized_keys:
        - "${SSH_RSA}"
systemd:
  units:
    - name: zincati.service
      mask: true
    - name: setup-pupernetes.service
      enabled: true
      contents: |
        [Unit]
        Description=Setup pupernetes
        Wants=network-online.target
        After=network-online.target

        [Service]
        Type=oneshot
        ExecStart=/usr/local/bin/setup-pupernetes
        RemainAfterExit=yes

        [Install]
        WantedBy=multi-user.target
    - name: install-pupernetes-dependencies.service
      enabled: true
      contents: |
        [Unit]
        Description=Install pupernetes dependencies
        Wants=network-online.target
        After=network-online.target

        [Service]
        Type=oneshot
        ExecStart=/usr/bin/rpm-ostree install --idempotent --reboot unzip
        RemainAfterExit=yes

        [Install]
        WantedBy=multi-user.target
    - name: pupernetes.service
      enabled: true
      contents: |
        [Unit]
        Description=Run pupernetes
        Requires=setup-pupernetes.service install-pupernetes-dependencies.service docker.service
        After=setup-pupernetes.service install-pupernetes-dependencies.service docker.service

        [Service]
        Environment=SUDO_USER=core
        ExecStart=/usr/local/bin/pupernetes daemon run /opt/sandbox --kubectl-link /opt/bin/kubectl -v 5 --hyperkube-version 1.10.1 --run-timeout 6h
        Restart=on-failure
        RestartSec=5
        Type=notify
        TimeoutStartSec=600
        TimeoutStopSec=120

        [Install]
        WantedBy=multi-user.target
    - name: terminate.service
      contents: |
        [Unit]
        Description=Trigger a poweroff

        [Service]
        ExecStart=/bin/systemctl poweroff
        Restart=no
    - name: terminate.timer
      enabled: true
      contents: |
        [Timer]
        OnBootSec=7200

        [Install]
        WantedBy=multi-user.target
storage:
  files:
    - path: /usr/local/bin/setup-pupernetes
      mode: 0500
      contents:
        source: "data:,%23%21%2Fbin%2Fbash%20-ex%0Acurl%20-Lf%20--retry%207%20--retry-connrefused%20https%3A%2F%2Fgithub.com%2FDataDog%2Fpupernetes%2Freleases%2Fdownload%2Fv0.11.0%2Fpupernetes%20-o%20%2Fusr%2Flocal%2Fbin%2Fpupernetes%0Asha512sum%20-c%20%2Fusr%2Flocal%2Fshare%2Fpupernetes.sha512sum%0Achmod%20%2Bx%20%2Fusr%2Flocal%2Fbin%2Fpupernetes%0A"
    - path: /usr/local/share/pupernetes.sha512sum
      mode: 0400
      contents:
        source: "data:,fcbf42316b9fbfbf6966b2f010f1bbc5006f7c882fc856d36b5e9f67a323d6b02361a45b88a4b4f7c64ac733078d9fd7d0cf72ef1229697f191b740c9fc95e61%20%2Fusr%2Flocal%2Fbin%2Fpupernetes%0A"
EOF
