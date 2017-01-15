# A8-Diego
Deploying A8 ControlPlane and Register to CF-Diego.

## A8 vs CF go Router
CF goRouter provides service registration and routing. CF container security model prohibits direct communication between service containers. Communication between services must pass through the goRouter.

A8 facilitates a fault tolerant direct communication between service instances by providing a service endpoint registry and routing controller. This functionality is redundant in a CF environment since it is provided by the goRouter.

However Bluemix did provide the A8 control plane via the now deprecated Service Discovery and Service Proxy services.

A8 also provides intelligent content based routing. This capability facilitates features such as A/B testing and traffic management.

## Installing the A8 Control Plane
BMX Public Diego docker deployment is disabled for security reasons. (See CF documentation). Hence the usual method of runing A8 control plane as a Docker image will not work on Bluemix Public.

To deploy A8 control plane, the registry and controller need to be deployed as CF applications.

This project sets up the Go project structure for push and build on Bluemix of the A8 control plane.

_*WIP* A backend datastore (Redis) can optionally be configured and specified as an installation parameter._

## Installing the a8sidecar in a CF app
The sidecar project needs to run in the same process as the service, since CF does not have the concept of K8 pods.

The same method to load the a8sidecar from the amalgam8 project can be used by a service deployed to CF by putting the a8sidecar install script in the '.profile/' directory.
