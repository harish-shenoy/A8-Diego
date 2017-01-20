# A8-Diego
Deploying A8 ControlPlane and Register to CF (Tested on Diego, but should work on DEA).

## A8 vs CF go Router
CF goRouter provides service registration and routing. CF container security model prohibits direct communication between service containers. Communication between services must pass through the goRouter.

A8 facilitates a fault tolerant direct communication between service instances by providing a service endpoint registry and routing controller. This functionality is redundant in a CF environment since it is provided by the goRouter.

However Bluemix did provide the A8 control plane via the now deprecated Service Discovery and Service Proxy services.

A8 also provides service tagging in the registry. This facilitates intelligent content based routing. This capability facilitates features such as A/B testing, traffic management, and which service version that are called.

## Installing the A8 Control Plane
BMX Public Diego docker deployment is disabled for security reasons. (See [Docker Security Concerns in a multi-tenant environment]{http://docs.cloudfoundry.org/adminguide/docker.html#multi-tenant}). Hence the usual method of running A8 control plane as a Docker image will not work on Bluemix Public.

To deploy A8 control plane, the registry and controller need to be deployed as CF applications.

This project sets up the Go project structure for push and build on Bluemix of the A8 control plane.

_*WIP* A backend datastore (Redis) can optionally be configured and specified as an installation parameter._

## Installing the a8sidecar in a CF app
The sidecar project needs to run in the same process as the service, since CF does not have the concept of K8 pods.

The method to load the a8sidecar from the amalgam8 project cannot be used by a service deployed to CF because the sidecar install script requires elevated privileges which are not available in the CF container OS stack. This project uses the a8sidecar binary which is built during the control plane build. The sidecar configuration should start the application process and supervise it. An alternative would be to use a multi-buildpack buildpack to configure a container with the application runtime as well as goLang for the sidecar truntime. Such a buildpack will be more complex, but will allow a native build of the sidecar for any supported CF Stack.

### A8 Sidecar configuration
In a CF environment, when registering the service in A8, the host name for the service should be set to the hostname as registered in goRouter, not the host or IP address of the container. If it is not set, the service registry may not be able to distinguish between service instances. This will result in the service being deregistered when one instance fails, even if there is more than one instance executing.

Healthchecks should be against the localhost so that the service count in the registry is correct.
