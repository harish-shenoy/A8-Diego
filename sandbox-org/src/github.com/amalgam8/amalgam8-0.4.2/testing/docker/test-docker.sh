#!/bin/bash
#
# Copyright 2016 IBM Corporation
#
#   Licensed under the Apache License, Version 2.0 (the "License");
#   you may not use this file except in compliance with the License.
#   You may obtain a copy of the License at
#
#       http://www.apache.org/licenses/LICENSE-2.0
#
#   Unless required by applicable law or agreed to in writing, software
#   distributed under the License is distributed on an "AS IS" BASIS,
#   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#   See the License for the specific language governing permissions and
#   limitations under the License.

set -x
set -o errexit

SCRIPTDIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )

echo "Testing docker-based deployment.."

$SCRIPTDIR/run-controlplane-docker.sh start
sleep 5

docker-compose -f $SCRIPTDIR/gateway.yaml up -d
sleep 5

docker-compose -f $SCRIPTDIR/bookinfo.yaml up -d
sleep 10

# Run the actual test workload
$SCRIPTDIR/../test-scripts/demo_script.sh

echo "Docker tests successful. Cleaning up.."
$SCRIPTDIR/cleanup.sh

sleep 10
