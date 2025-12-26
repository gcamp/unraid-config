#!/bin/bash
# Start the Management Utility
/usr/local/sbin/emhttp 

# This sets the USB friendly name mapping rule to apply each boot

cp /boot/config/rules.d/99-usb-rules.rules /etc/udev/rules.d/99-usb-rules.rules
chmod 644 /etc/udev/rules.d/99-usb-rules.rules
udevadm control --reload-rules
udevadm trigger --attr-match=subsystem=usb

# Reserve memory for system to prevent total OOM lockup
sysctl -w vm.min_free_kbytes=262144
sysctl -w vm.admin_reserve_kbytes=131072
