#!/bin/bash

if [ "$#" -lt 1 ]; then
  echo "Usage: $0 package1 [package2 ... packageN]"
  exit 1
fi

# Convert the package arguments into a comma-separated list
package_list=$(printf ",%s" "$@")
package_list=${package_list:1}

# Run the Ansible playbook with the package list as a variable
ansible-playbook install-packages.yml -e "package_list=${package_list}"
