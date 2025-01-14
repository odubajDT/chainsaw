# Pod logs

Collecting pod logs can help understand what happened inside one or more pods.

## Configuration

!!! tip "Reference documentation"
    - The full structure of the `PodLogs` resource is documented [here](../reference/apis/chainsaw.v1alpha1.md#chainsaw-kyverno-io-v1alpha1-PodLogs).

### All namespaces

When used with a namespaced resource, it is possible to consider all namespaces in the cluster by setting `namespace: '*'`.

## Usage examples

### Single pod

If a pod `name` is specified, Chainsaw will retrieve logs from this specific pod in the test namespace.

!!! example "Collect pod logs in the test namespace"

    ```yaml
    apiVersion: chainsaw.kyverno.io/v1alpha1
    kind: Test
    metadata:
      name: example
    spec:
      steps:
      - try:
        # ...
        catch:
        - podLogs:
            name: my-pod
        # ...
        finally:
        - podLogs:
            name: my-pod
        # ...
    ```

If a `namespace` is specified, Chainsaw will retrieve logs from this specific pod in the specified namespace.

!!! example "Collect pod logs in a specific namespace"

    ```yaml
    apiVersion: chainsaw.kyverno.io/v1alpha1
    kind: Test
    metadata:
      name: example
    spec:
      steps:
      - try:
        # ...
        catch:
        - podLogs:
            name: my-pod
            namespace: foo
        # ...
        finally:
        - podLogs:
            name: my-pod
            namespace: foo
        # ...
    ```

### All pods

If no pod `name` and `namespace` is specified, Chainsaw will retrieve logs from all pods in the test namespace.

!!! example "Collect all pod logs in the test namespace"

    ```yaml
    apiVersion: chainsaw.kyverno.io/v1alpha1
    kind: Test
    metadata:
      name: example
    spec:
      steps:
      - try:
        # ...
        catch:
        - podLogs: {}
        # ...
        finally:
        - podLogs: {}
        # ...
    ```

On the other hand, if a `namespace` is specified, Chainsaw will retrieve logs from all pods in the specified namespace.

!!! example "Collect all pod logs in a specific namespace"

    ```yaml
    apiVersion: chainsaw.kyverno.io/v1alpha1
    kind: Test
    metadata:
      name: example
    spec:
      steps:
      - try:
        # ...
        catch:
        - podLogs:
            namespace: foo
        # ...
        finally:
        - podLogs:
            namespace: foo
        # ...
    ```

### Label selector

An optional [label selector](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors) can be configured to refine the pods to retrieve logs from.

!!! example "Collect pod logs using a label selector in the test namespace"

    ```yaml
    apiVersion: chainsaw.kyverno.io/v1alpha1
    kind: Test
    metadata:
      name: example
    spec:
      steps:
      - try:
        # ...
        catch:
        - podLogs:
            selector: app=my-app
        # ...
        finally:
        - podLogs:
            selector: app=my-app
        # ...
    ```

If a `namespace` is specified, Chainsaw will retrieve pod logs using the specified namespace.

!!! example "Collect pod logs using a label selector in a specific namespace"

    ```yaml
    apiVersion: chainsaw.kyverno.io/v1alpha1
    kind: Test
    metadata:
      name: example
    spec:
      steps:
      - try:
        # ...
        catch:
        - podLogs:
            selector: app=my-app
            namespace: foo
        # ...
        finally:
        - podLogs:
            selector: app=my-app
            namespace: foo
        # ...
    ```

### Tail

The `tail` field can be used to limit the amount of log lines retrieved when querying pod logs.

!!! note "Default"
    By default, `tail` will be `10` when a label selector is used, and `all` if a pod `name` is specified.

!!! example "Tail example"

    ```yaml
    apiVersion: chainsaw.kyverno.io/v1alpha1
    kind: Test
    metadata:
      name: example
    spec:
      steps:
      - try:
        # ...
        catch:
        - podLogs:
            selector: app=my-app
            namespace: foo
            tail: 30
        # ...
        finally:
        - podLogs:
            selector: app=my-app
            namespace: foo
            tail: 30
        # ...
    ```

### Container

The `container` field can be used to retrieve logs from a specific container in the pod.

!!! note "Default"

    By default logs from all containers will be fetched.

!!! example "Container example"

    ```yaml
    apiVersion: chainsaw.kyverno.io/v1alpha1
    kind: Test
    metadata:
      name: example
    spec:
      steps:
      - try:
        # ...
        catch:
        - podLogs:
            selector: app=my-app
            namespace: foo
            container: nginx
        # ...
        finally:
        - podLogs:
            selector: app=my-app
            namespace: foo
            container: nginx
        # ...
    ```
