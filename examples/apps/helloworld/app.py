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

import os
from flask import Flask, request
app = Flask(__name__)
healthchecks = 0

@app.route('/hello')
def hello():
    service_version = os.environ.get('A8_SERVICE').split(':')
    version = service_version[1] if len(service_version) == 2 else 'UNVERSIONED'
    return 'Hello version: %s, container: %s\n' % (version, os.environ.get('HOSTNAME'))

@app.route('/health')
def health():
    rc = 200
    # Uncomment the following code to fail every 4, 5,& 6th healthcheck to demonstrate
    # service deregistration and re-registration on healthcheck.
#    global healthchecks
#    healthchecks += 1
#    if healthchecks > 3:
#        rc = 500
#    if healthchecks == 6:
#        healthchecks = 0
    return 'Helloworld is healthy', rc

if __name__ == "__main__":
    app.run(host='0.0.0.0', threaded=True, port=os.environ.get('PORT'))
