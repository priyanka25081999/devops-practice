**Informers**

   Informers use the Kubernetes API to learn about changes in the state of a Kubernetes cluster and use that information to maintain a cache (the indexer) of the current cluster state and to inform clients about the changes by calling handler functions. 
   To achieve this, an informer (more specifically: a shared informer) is again a composition of several components.

* The reflector picks up the change via the Kubernetes API and enqueues it in the delta FIFO queue, where deduplication is performed and all changes for the same object are consolidated

* The controller is eventually reading the change from the delta FIFO queue

* The controller updates the indexer

* The controller forwards the change to the processor, which in turns calls all handler functions

* The handler functions are methods of the custom controller who can then inspect the delta, use the updated indexer to retrieve additional information and adjust the cluster state if needed

