# A8-Diego
Deploying A8 ControlPlane and Register to CF-Diego.

## Installing the A8 Control Plane
BMX Public Diego docker deployment is disabled for security reasons. (See CF documentation). Hence the usual method of runing A8 control plane as a Docker image will not work on Bluemix Public.

To deploy A8 control plane, the registry and controller need to be deployed as CF applications.

This project sets up the Go project structure for push and build on Bluemix of the A8 control plane.

_*WIP* A backend datastore (Redis) can optionally be configured and specified as an installation parameter._

## Installing the a8sidecar in a CF app
The sidecar project needs to run in the same process as the service, since CF does not have the concept of K8 pods.

The same method to load the a8sidecar from the amalgam8 project can be used by a service deployed to CF by putting the a8sidecar install script in the '.profile/' directory.
