**Kube-state-metrics (KSM)** is a Kubernetes add-on that monitors and generates metrics about the state of Kubernetes objects, 
such as deployments, nodes, and pods. It listens to the Kubernetes API server and exposes metrics in Prometheus format, 
making it easy to integrate with Prometheus for monitoring and alerting. KSM focuses on the health and status of Kubernetes resources 
rather than resource usage metrics like CPU and memory, which are provided by the metrics-server. 
